use diesel::prelude::*;
use juniper::FieldResult;

use super::Ctx;
use crate::models::{
    AddArticleInput, AddArticlePayload, Article, DeleteArticleInput, DeleteArticlePayload,
    UpdateArticleInput, UpdateArticlePayload,
};

pub struct Mutation;

#[juniper::graphql_object(Context = Ctx)]
impl Mutation {
    fn add_article(ctx: &mut Ctx, input: AddArticleInput) -> FieldResult<AddArticlePayload> {
        use crate::schema::articles;
        let conn = ctx.pool.get()?;
        let article = diesel::insert_into(articles::table)
            .values(input)
            .get_result::<Article>(&conn)?;
        Ok(AddArticlePayload { article })
    }
    fn update_article(
        ctx: &mut Ctx,
        input: UpdateArticleInput,
    ) -> FieldResult<UpdateArticlePayload> {
        use crate::schema::articles;
        let conn = ctx.pool.get()?;
        let article = diesel::update(articles::table.find(uuid::Uuid::parse_str(&input.id)?))
            .set((
                articles::title.eq(input.title),
                articles::body.eq(input.body),
            ))
            .get_result::<Article>(&conn)?;
        Ok(UpdateArticlePayload { article })
    }
    fn delete_article(
        ctx: &mut Ctx,
        input: DeleteArticleInput,
    ) -> FieldResult<DeleteArticlePayload> {
        use crate::schema::articles::dsl::*;
        let conn = ctx.pool.get()?;
        diesel::delete(articles.find(uuid::Uuid::parse_str(&input.id)?)).execute(&conn)?;
        Ok(DeleteArticlePayload {
            deleted_id: input.id,
        })
    }
}
