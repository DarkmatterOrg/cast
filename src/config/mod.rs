pub mod get;
pub mod set;

use directories::ProjectDirs;
use serde::{Deserialize, Serialize};
use std::path::PathBuf;

#[derive(Serialize, Deserialize, Default)]
#[serde(default)]
pub struct Config {
    pub insults: bool,
    // Add new fields here with #[serde(default)] attribute
}

pub fn config_path() -> PathBuf {
    if let Some(proj_dirs) = ProjectDirs::from("dev", "aethrexal", "cast") {
        return proj_dirs.config_dir().join("config.toml");

        // Linux:   /home/alice/.config/cast
        // Windows: C:\Users\Alice\AppData\Roaming\aethrexal\cast
        // macOS:   /Users/Alice/Library/Application Support/dev.aethrexal.cast
    } else {
        return Default::default();
    }
}
