use super::{config_path, get};
use std::fs;

pub fn initialize_config() {
    let path = config_path();

    if !path.as_os_str().is_empty() {
        // Create config directory if needed
        if let Some(parent) = path.parent() {
            fs::create_dir_all(parent).expect("Failed to create config directory");
        }

        // Merge existing config with defaults
        let config = get::get_config();

        // Write updated config
        let toml = toml::to_string_pretty(&config).expect("Failed to serialize config");
        fs::write(&path, toml).expect("Failed to write config");
    }
}
