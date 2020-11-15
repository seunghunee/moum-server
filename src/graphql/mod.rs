use super::models::*;
use diesel::pg::PgConnection;
use diesel::prelude::*;
pub use diesel::r2d2::{ConnectionManager, Pool};
use juniper::FieldResult;

pub struct Ctx {
    pub pool: Pool<ConnectionManager<PgConnection>>,
}

impl juniper::Context for Ctx {}
pub mod mutation;
pub mod query;
