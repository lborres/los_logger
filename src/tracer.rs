use std::fs::OpenOptions;

use chrono::Local;
use crate::{env::Env, time::CustomTimer};
use tracing::{level_filters::LevelFilter, Level};
use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt, Layer};

pub fn init(env: &Env) -> tracing_appender::non_blocking::WorkerGuard {
    let file_timestamp = Local::now().format("%y%m%d").to_string();
    let filename = format!("{}{}_{}", &env.log_dir, file_timestamp, &env.log_file_name);
    let log_file = OpenOptions::new()
        .create(true)
        .append(true)
        .open(&filename)
        .expect("Failed to create log file");

    let (non_blocking, guard) = tracing_appender::non_blocking(log_file);

    let file_layer = tracing_subscriber::fmt::layer()
        .with_timer(CustomTimer)
        .with_writer(non_blocking)
        .with_target(false)
        .with_level(false)
        .with_ansi(false)
        // INFO: log file will never need to contain debug info
        .with_filter(LevelFilter::from_level(Level::INFO));

    let console_layer = tracing_subscriber::fmt::layer()
        .with_timer(CustomTimer)
        .with_target(false)
        .with_level(false)
        .with_filter(LevelFilter::from_level(env.log_level));

    tracing_subscriber::registry()
        .with(file_layer)
        .with(console_layer)
        .init();

    // INFO: WorkerGuard implements Drop
    // return guard in main so that it will not be dropped
    guard
}
