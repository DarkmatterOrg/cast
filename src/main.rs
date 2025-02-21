mod args;
mod modules;

use args::Commands;
use clap::Parser;
use colored::Colorize;
//use std::path::Path;

use crate::modules::{
    auto_update::auto_update, bios::bios, clean_system::clean_system, toggle_password_feedback::toggle_password_feedback, update::update,
};

const VERSION: &str = clap::crate_version!();
fn main() {
    //if !Path::new("/usr/bin/nebula").exists() {
    //    panic!("{}", "Nebula isn't present.".bold().red());
    //}

    let cli = args::Cli::parse();

    match cli.command {
        Commands::Version => {
            println!("{}: v{}", "Cast".bold().purple(), VERSION.blue());
        }

        Commands::AutoUpdate(args) => {
            if is_root() {
                auto_update(&args);
            } else {
                println!("{}: Please run this command with sudo.", "WARNING".bold().yellow());
            }
        }

        Commands::Bios => {
            bios();
        }

        Commands::CleanSystem => {
            clean_system();
        }

        Commands::TogglePasswordFeedback(args) => {
            //TODO toggle_password_feedback(&args);
        }

        Commands::Update(args) => {
            update(&args);
        }
    }
}

fn is_root() -> bool {
    unsafe { libc::geteuid() == 0 }
}
