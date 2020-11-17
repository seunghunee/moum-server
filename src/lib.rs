mod graphql;
mod models;
mod schema;

pub struct Config {
    pub db_url: String,
    pub port: u16,
}

use std::env;
impl Config {
    pub fn new() -> Result<Config, &'static str> {
        match env::var("DATABASE_URL") {
            Err(_) => Err("DATABASE_URL must be set"),
            Ok(db_url) => Ok(Config { db_url, port: 8080 }),
        }
    }
}

#[macro_use]
extern crate diesel;
use graphql::{Mutation, Query, *};
use hyper::{
    service::{make_service_fn, service_fn},
    Method, Response, Server, StatusCode,
};
use juniper::{EmptySubscription, RootNode};
use std::{error::Error, sync::Arc};
pub async fn run(config: Config) -> Result<(), Box<dyn Error>> {
    let manager = ConnectionManager::new(config.db_url);
    let pool = Pool::new(manager)?;
    let ctx = Arc::new(Ctx { pool });
    let root_node = Arc::new(RootNode::new(Query, Mutation, EmptySubscription::new()));
    let make_svc = make_service_fn(move |_| {
        let ctx = ctx.clone();
        let root_node = root_node.clone();
        async move {
            Ok::<_, hyper::Error>(service_fn(move |req| {
                let ctx = ctx.clone();
                let root_node = root_node.clone();
                async move {
                    match (req.method(), req.uri().path()) {
                        (&Method::GET, "/graphql") | (&Method::POST, "/graphql") => {
                            juniper_hyper::graphql(root_node, ctx, req).await
                        }
                        (&Method::GET, "/playground") => {
                            juniper_hyper::playground("/graphql", None).await
                        }
                        _ => {
                            let mut not_found = Response::default();
                            *not_found.status_mut() = StatusCode::NOT_FOUND;
                            Ok(not_found)
                        }
                    }
                }
            }))
        }
    });

    // TODO: Random port number
    let addr = ([127, 0, 0, 1], config.port).into();
    let server = Server::bind(&addr).serve(make_svc);
    server.await?;

    Ok(())
}
