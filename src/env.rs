use std::env;

use tracing::Level;

#[derive(Debug)]
pub struct Env {
    pub log_level: Level,
    pub beacon_target: String,
    pub log_dir: String,
    pub log_file_name: String,
    pub log_file_loc: String,
}

impl Env {
    pub fn init() -> Self {
        dotenv::dotenv().ok();
        let log_level: Level = match env::var("LOG_LEVEL").unwrap_or_else(|_| {return "INFO".to_string()}).as_str() {
            "TRACE" => Level::TRACE,
            "DEBUG" => Level::DEBUG,
            "INFO" => Level::INFO,
            "WARN" => Level::WARN,
            "ERROR" => Level::ERROR,
            _ => Level::INFO
        };
        let beacon_target = env::var("BEACON_TARGET").expect("Missing required env variable");
        let log_dir = env::var("LOG_DIR").expect("Missing required env variable");
        let log_file_name = env::var("LOG_FILE_NAME").expect("Missing required env variable");
        let log_file_loc = env::var("LOG_FILE_LOC").expect("Missing required env variable");

        Self {
            log_level,
            beacon_target,
            log_dir,
            log_file_name,
            log_file_loc,
        }
    }
}
