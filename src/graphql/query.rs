use diesel::prelude::*;
use juniper::FieldResult;

use super::Ctx;
use crate::models::Article;

pub struct Query;

#[juniper::graphql_object(Context = Ctx)]
impl Query {
    fn articles(ctx: &Ctx) -> FieldResult<Vec<Article>> {
        use crate::schema::articles;
        let conn = ctx.pool.get().expect("Error pool get");
        let result = articles::table
            .limit(5)
            .load::<Article>(&conn)
            .expect("Error loading articles");
        Ok(result)
    }
    fn article(ctx: &Ctx, title: String) -> FieldResult<Article> {
        use crate::schema::articles;
        let conn = ctx.pool.get().expect("Error pool get");
        let result = articles::table
            .filter(articles::title.eq(&title))
            .first::<Article>(&conn)
            .expect("Error loading article");
        Ok(result)
    }
}
