use super::{Config, config_path};
use crate::utils::status_msg::warning;
use std::fs;

pub fn get_config() -> Config {
    let path = match config_path() {
        Some(p) => p,
        None => return Config::default(),
    };

    // Return default config if file doesn't exist
    if !path.exists() {
        return Config::default();
    }

    // Read and parse config
    match fs::read_to_string(&path) {
        Ok(contents) => toml::from_str(&contents).unwrap_or_else(|_| {
            warning("Invalid config file, using defaults");
            Config::default()
        }),
        Err(_) => {
            warning("Failed to read config file, using defaults");
            Config::default()
        }
    }
}
