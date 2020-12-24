use juniper::{graphql_interface, graphql_object, GraphQLInputObject, GraphQLObject, ID};

// TODO: 근본적인 문제해결 (https://github.com/graphql-rust/juniper/issues/814)
#[graphql_interface(for = Article)]
pub trait Node {
    fn id(&self) -> ID;
}

#[derive(Queryable)]
pub struct Article {
    id: uuid::Uuid,
    title: String,
    body: String,
}
#[graphql_object(impl = NodeValue)]
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
#[graphql_interface]
impl Node for Article {
    fn id(&self) -> ID {
        self.id.to_string().into()
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
