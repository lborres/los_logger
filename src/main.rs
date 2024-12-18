use std::{net::TcpStream, time::{Duration, Instant}};

use rand::prelude::*;
use los_logger::{env::Env, time::format_duration, tracer};
use tokio::time::sleep;
use tracing::{ info, trace, warn };

struct Status {
    connected: bool,
    last_disconnect: Option<Instant>
}

impl Status {
    fn check_service(&mut self, target: &str) {
        if self.ping_web(target) {
            if !self.connected {
                let duration = self.last_disconnect.unwrap_or_else(Instant::now).elapsed();
                let formatted_duration = format_duration(duration);
                info!("Connection restored after {}", formatted_duration)
            };
            self.connected = true;
        } else {
            if self.connected {
                warn!("Internet Service Interruption");
                self.last_disconnect = Some(Instant::now());
                self.connected = false;
            }
        }
    }

    fn ping_web(&self, target: &str) -> bool {
        trace!("Pinging: {}", target);
        match TcpStream::connect_timeout(&target.parse().expect("Invalid address"), Duration::from_secs(5)) {
            Ok(_) => true,
            Err(_) => false,
        }
    }
}

#[tokio::main]
async fn main() {
    let env = Env::init();
    let _guard = tracer::init(&env);
    trace!("Initiated Environment Variables\n{:#?}", env);

    info!("Starting LOS Logger");

    let mut status = Status {
        connected: true,
        last_disconnect: None
    };

    let mut rng = thread_rng();
    loop {
        status.check_service(&env.beacon_target);
        let delay = Duration::from_millis(4000 + rng.gen_range(0..3000));
        trace!("Sleeping for {:?}", delay);
        sleep(delay).await;
    }
}
