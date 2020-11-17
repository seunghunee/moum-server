pub mod query;
pub use query::Query;

pub mod mutation;
pub use mutation::Mutation;

use diesel::pg::PgConnection;
pub use diesel::r2d2::{ConnectionManager, Pool};

pub struct Ctx {
    pub pool: Pool<ConnectionManager<PgConnection>>,
}

impl juniper::Context for Ctx {}
