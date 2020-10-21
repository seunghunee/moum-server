mod models {
    #[derive(juniper::GraphQLObject, Queryable)]
    pub struct Article {
        id: i32,
        title: String,
        body: String,
    }

    use super::schema::articles;
    #[derive(juniper::GraphQLInputObject, Insertable)]
    #[table_name = "articles"]
    pub struct AddArticleInput {
        title: String,
        body: String,
    }
}

mod schema;

#[macro_use]
extern crate diesel;
use diesel::pg::PgConnection;
use diesel::prelude::*;
use diesel::r2d2::{ConnectionManager, Pool};
use std::env;

use self::models::*;

use hyper::service::{make_service_fn, service_fn};
use hyper::{Method, Response, Server, StatusCode};
use juniper::{EmptySubscription, FieldResult, RootNode};
use std::net::SocketAddr;
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
}

struct Mutation;
#[juniper::graphql_object(Context = Ctx)]
impl Mutation {
    fn addArticle(ctx: &mut Ctx, input: AddArticleInput) -> FieldResult<bool> {
        use schema::articles;
        let conn = ctx.pool.get().expect("Error: get db pool");
        diesel::insert_into(articles::table)
            .values(input)
            .get_result::<Article>(&conn)
            .expect("Error saving new article");
        Ok(true)
    }
}

type Schema = RootNode<'static, Query, Mutation, EmptySubscription<Ctx>>;

#[tokio::main]
async fn main() {
    let addr = SocketAddr::from(([127, 0, 0, 1], 3000));

    let db_url = env::var("DATABASE_URL").expect("DATABASE_URL must be set");
    let manager = ConnectionManager::new(db_url);
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

    let server = Server::bind(&addr).serve(make_svc);

    if let Err(e) = server.await {
        eprintln!("server error: {}", e);
    }
}
