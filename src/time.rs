use std::time::Duration;

use chrono::Local;
use tracing_subscriber::fmt::time::FormatTime;

pub struct CustomTimer;

impl FormatTime for CustomTimer {
    fn format_time(&self, w: &mut tracing_subscriber::fmt::format::Writer<'_>) -> std::fmt::Result {
        let now = Local::now();
        write!(w, "{}", now.format("%Y-%m-%d %H:%M:%S%.3f"))
    }
}

pub fn format_duration(duration: Duration) -> String {
    let days = duration.as_secs() / 86400;
    let hours = (duration.as_secs() % 86400) / 3600;
    let minutes = (duration.as_secs() % 3600) / 60;
    let seconds = duration.as_secs() % 60;
    let millis = duration.subsec_millis();

    let mut output = String::new();

    if days > 0 {
        output.push_str(&format!("{}d", days));
    }
    if hours > 0 || !output.is_empty() {
        output.push_str(&format!("{}h", hours));
    }
    if minutes > 0 || !output.is_empty() {
        output.push_str(&format!("{}m", minutes));
    }
    if seconds > 0 || !output.is_empty() {
        output.push_str(&format!("{}.{:03}s", seconds, millis));
    }

    output.trim_end().to_string()
}
