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
