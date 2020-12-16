use juniper::{graphql_object, GraphQLInputObject, GraphQLObject, ID};

#[derive(Queryable)]
pub struct Article {
    id: uuid::Uuid,
    title: String,
    body: String,
}
#[graphql_object]
impl Article {
    fn id(&self) -> ID {
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
#[derive(GraphQLInputObject, Insertable)]
#[table_name = "articles"]
pub struct AddArticleInput {
    title: String,
    body: String,
}

pub struct AddArticlePayload {
    pub article: Article,
}
#[graphql_object]
impl AddArticlePayload {
    fn article(&self) -> &Article {
        &self.article
    }
}

#[derive(GraphQLInputObject)]
pub struct UpdateArticleInput {
    pub id: ID,
    pub title: String,
    pub body: String,
}
pub struct UpdateArticlePayload {
    pub article: Article,
}
#[graphql_object]
impl UpdateArticlePayload {
    fn article(&self) -> &Article {
        &self.article
    }
}

#[derive(GraphQLInputObject)]
pub struct DeleteArticleInput {
    pub id: ID,
}
#[derive(GraphQLObject)]
pub struct DeleteArticlePayload {
    pub deleted_id: ID,
}
