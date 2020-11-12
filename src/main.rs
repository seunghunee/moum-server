mod models {
    #[derive(Queryable)]
    pub struct Article {
        id: uuid::Uuid,
        title: String,
        body: String,
    }
    #[juniper::graphql_object]
    impl Article {
        fn id(&self) -> juniper::ID {
            self.id.to_string().into()
        }
        fn title(&self) -> &str {
            self.title.as_str()
        }
        fn body(&self) -> &str {
            self.body.as_str()
        }
    }

    use super::schema::articles;
    #[derive(juniper::GraphQLInputObject, Insertable)]
    #[table_name = "articles"]
    pub struct AddArticleInput {
        title: String,
        body: String,
    }

    pub struct AddArticlePayload {
        pub article: Article,
    }
    #[juniper::graphql_object]
    impl AddArticlePayload {
        fn article(&self) -> &Article {
            &self.article
        }
    }

    #[derive(juniper::GraphQLInputObject)]
    pub struct UpdateArticleInput {
        pub id: juniper::ID,
        pub title: String,
        pub body: String,
    }
    pub struct UpdateArticlePayload {
        pub article: Article,
    }
    #[juniper::graphql_object]
    impl UpdateArticlePayload {
        fn article(&self) -> &Article {
            &self.article
        }
    }

    #[derive(juniper::GraphQLInputObject)]
    pub struct DeleteArticleInput {
        pub id: juniper::ID,
    }
    #[derive(juniper::GraphQLObject)]
    pub struct DeleteArticlePayload {
        pub deleted_id: juniper::ID,
    }
}

mod schema;

#[macro_use]
extern crate diesel;
use diesel::pg::PgConnection;
use diesel::prelude::*;
use diesel::r2d2::{ConnectionManager, Pool};

use self::models::*;

use hyper::service::{make_service_fn, service_fn};
use hyper::{Method, Response, Server, StatusCode};
use juniper::{EmptySubscription, FieldResult, RootNode};
use std::sync::Arc;

struct Ctx {
    pool: Pool<ConnectionManager<PgConnection>>,
}

impl juniper::Context for Ctx {}

struct Query;
#[juniper::graphql_object(Context = Ctx)]
impl Query {
    fn articles(ctx: &Ctx) -> FieldResult<Vec<Article>> {
        use schema::articles;
        let conn = ctx.pool.get().expect("Error pool get");
        let result = articles::table
            .limit(5)
            .load::<Article>(&conn)
            .expect("Error loading articles");
        Ok(result)
    }
    fn article(ctx: &Ctx, title: String) -> FieldResult<Article> {
        use schema::articles;
        let conn = ctx.pool.get().expect("Error pool get");
        let result = articles::table
            .filter(articles::title.eq(&title))
            .first::<Article>(&conn)
            .expect("Error loading article");
        Ok(result)
    }
}

struct Mutation;
#[juniper::graphql_object(Context = Ctx)]
impl Mutation {
    fn add_article(ctx: &mut Ctx, input: AddArticleInput) -> FieldResult<AddArticlePayload> {
        use schema::articles;
        let conn = ctx.pool.get().expect("Error: get db pool");
        let article = diesel::insert_into(articles::table)
            .values(input)
            .get_result::<Article>(&conn)
            .expect("Error saving new article");
        Ok(AddArticlePayload { article })
    }
    fn update_article(
        ctx: &mut Ctx,
        input: UpdateArticleInput,
    ) -> FieldResult<UpdateArticlePayload> {
        use schema::articles;
        let conn = ctx.pool.get().expect("Error: get db pool");
        let article =
            diesel::update(articles::table.find(uuid::Uuid::parse_str(&input.id).unwrap()))
                .set((
                    articles::title.eq(input.title),
                    articles::body.eq(input.body),
                ))
                .get_result::<Article>(&conn)
                .expect("Error saving new article");
        Ok(UpdateArticlePayload { article })
    }
    fn delete_article(
        ctx: &mut Ctx,
        input: DeleteArticleInput,
    ) -> FieldResult<DeleteArticlePayload> {
        use schema::articles::dsl::*;
        let conn = ctx.pool.get().expect("Error: get db pool");
        diesel::delete(articles.find(uuid::Uuid::parse_str(&input.id).unwrap()))
            .execute(&conn)
            .expect("Error deleting article");
        Ok(DeleteArticlePayload {
            deleted_id: input.id,
        })
    }
}

type Schema = RootNode<'static, Query, Mutation, EmptySubscription<Ctx>>;

use std::process;

#[tokio::main]
async fn main() {
    let config = Config::new().unwrap_or_else(|err| {
        eprintln!("{}", err);
        process::exit(1);
    });

    let manager = ConnectionManager::new(config.db_url);
    let pool = Pool::new(manager).expect("Error pool");
    let ctx = Arc::new(Ctx { pool });
    let root_node = Arc::new(Schema::new(Query, Mutation, EmptySubscription::new()));
    let make_svc = make_service_fn(move |_| {
        let ctx = ctx.clone();
        let root_node = root_node.clone();
        async move {
            Ok::<_, hyper::Error>(service_fn(move |req| {
                let ctx = ctx.clone();
                let root_node = root_node.clone();
                async move {
                    match (req.method(), req.uri().path()) {
                        (&Method::GET, "/graphql") | (&Method::POST, "/graphql") => {
                            juniper_hyper::graphql(root_node, ctx, req).await
                        }
                        (&Method::GET, "/playground") => {
                            juniper_hyper::playground("/graphql", None).await
                        }
                        _ => {
                            let mut not_found = Response::default();
                            *not_found.status_mut() = StatusCode::NOT_FOUND;
                            Ok(not_found)
                        }
                    }
                }
            }))
        }
    });

    // TODO: Random port number
    let addr = ([127, 0, 0, 1], config.port).into();
    let server = Server::bind(&addr).serve(make_svc);

    if let Err(e) = server.await {
        eprintln!("moum: {}", e);
    }
}

struct Config {
    db_url: String,
    port: u16,
}

use std::env;
impl Config {
    fn new() -> Result<Config, &'static str> {
        match env::var("DATABASE_URL") {
            Err(_) => Err("moum: DATABASE_URL must be set"),
            Ok(db_url) => Ok(Config { db_url, port: 8080 }),
        }
    }
}
