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

pub fn config_path() -> Option<PathBuf> {
    ProjectDirs::from("dev", "aethrexal", "cast")
        .map(|proj_dirs| proj_dirs.config_dir().join("config.toml"))
}
