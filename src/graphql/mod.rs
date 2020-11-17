pub mod query;
pub use query::Query;

pub mod mutation;
pub use mutation::Mutation;

use super::models::*;
use diesel::pg::PgConnection;
use diesel::prelude::*;
pub use diesel::r2d2::{ConnectionManager, Pool};
use juniper::FieldResult;

pub struct Ctx {
    pub pool: Pool<ConnectionManager<PgConnection>>,
}

impl juniper::Context for Ctx {}
