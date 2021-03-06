use diesel::prelude::*;
use juniper::FieldResult;

use super::Ctx;
use crate::models::Article;

pub struct Query;

#[juniper::graphql_object(Context = Ctx)]
impl Query {
    fn articles(ctx: &Ctx) -> FieldResult<Vec<Article>> {
        use crate::schema::articles::dsl::*;
        let conn = ctx.pool.get()?;
        let result = articles.limit(5).load::<Article>(&conn)?;
        Ok(result)
    }
    fn article(ctx: &Ctx, title: String) -> FieldResult<Article> {
        use crate::schema::articles;
        let conn = ctx.pool.get()?;
        let result = articles::table
            .filter(articles::title.eq(&title))
            .first::<Article>(&conn)?;
        Ok(result)
    }
}
