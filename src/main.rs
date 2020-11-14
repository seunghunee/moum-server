use std::process;

use moum_server::Config;

#[tokio::main]
async fn main() {
    let config = Config::new().unwrap_or_else(|e| {
        eprintln!("moum: {}", e);
        process::exit(1);
    });

    if let Err(e) = moum_server::run(config).await {
        eprintln!("moum: {}", e);
        process::exit(1);
    };
}
