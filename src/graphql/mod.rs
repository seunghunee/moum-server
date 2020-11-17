pub mod query;
pub use query::Query;

pub mod mutation;
pub use mutation::Mutation;

use diesel::{
    pg::PgConnection,
    r2d2::{ConnectionManager, Pool},
};
use std::error::Error;

pub struct Ctx {
    pool: Pool<ConnectionManager<PgConnection>>,
}

impl juniper::Context for Ctx {}
impl Ctx {
    pub fn new(db_url: &str) -> Result<Ctx, Box<dyn Error>> {
        let manager = ConnectionManager::new(db_url);
        let pool = Pool::new(manager)?;

        Ok(Ctx { pool })
    }
}
