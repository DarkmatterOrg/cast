mod args;
mod config;
mod modules;
mod utils;

use args::Commands;
use clap::Parser;
use colored::Colorize;
use std::path::Path;

use crate::modules::{
    auto_update::auto_update, bios::bios, clean_system::clean_system, dev::dev, fix::fix,
    toggle_password_feedback::toggle_password_feedback, update::update,
};

use crate::utils::{
    is_root::is_root,
    status_msg::{error, info, notice},
};

const VERSION: &str = clap::crate_version!();
fn main() {
    if !Path::new("/usr/share/umbra").exists() {
        error("Cast can only be used on Umbra.");
        return;
    }

    config::set::initialize_config();

    let cli = args::Cli::parse();

    match cli.command {
        Commands::Version => {
            println!("{}: v{}", "Cast".bold().purple(), VERSION.blue());
        }

        Commands::AutoUpdate(args) => {
            if is_root() {
                auto_update(&args);
            } else {
                notice("Please run this command with sudo.");
            }
        }

        Commands::Bios => {
            bios();
        }

        Commands::CleanSystem => {
            clean_system();
        }

        Commands::Config => {
            info(format!("Config can be found at: {:?}", config::config_path()).as_str());
        }

        Commands::DevUtil(args) => {
            dev(&args);
        }

        Commands::Fix(args) => {
            fix(&args);
        }

        Commands::PasswordFeedback(args) => {
            if is_root() {
                toggle_password_feedback(&args);
            } else {
                notice("Please run this command with sudo.");
            }
        }

        Commands::Update(args) => {
            update(&args);
        }
    }
}
