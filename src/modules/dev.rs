use crate::args::DevArgs;
use crate::config;
use crate::utils::status_msg::{done, info, notice};
use crate::utils::{is_cmd_installed::is_cmd_installed, status_msg::warning};
use colored::Colorize;
use std::process::Command;

pub fn dev(args: &DevArgs) {
    let is_insult_enabled = config::get::get_config().insults;

    if args.install && args.remove {
        if is_insult_enabled {
            warning(
                "What the fuck are you doing? You can't use --install and --remove at the same time!",
            );
        } else {
            warning("You can't use both --install and --remove at the same time.");
        }

        return;
    }

    match args.option.as_str() {
        "rustup" => {
            if args.install {
                if is_cmd_installed("rustup") {
                    if is_insult_enabled {
                        notice("Are you dumb? Rustup is already installed...");
                    } else {
                        notice("Rustup is already installed, nothing todo.");
                    }

                    return;
                }

                Command::new("sh")
                    .args([
                        "-c",
                        "curl --tlsv1.2 -sSf 'https://sh.rustup.rs' | sh -s -- -y",
                    ])
                    .status()
                    .expect("Was unable to install rustup");

                done("Rustup got installed, please restart the terminal to start using it.");
            } else if args.remove {
                if !is_cmd_installed("rustup") {
                    if is_insult_enabled {
                        notice("Are you dumb? Rustup isn't even installed...");
                    } else {
                        notice("Rustup is not installed, nothing todo.");
                    }
                    return;
                }

                Command::new("rustup")
                    .args(["self", "uninstall"])
                    .status()
                    .expect("Wasn't able to uninstall rustup");

                done("Rustup got removed, please restart the terminal to start using it.");
            } else {
                if is_cmd_installed("rustup") {
                    info(format!("Rustup is currently: {}", "installed".bold().green()).as_str());
                } else {
                    info(format!("Rustup is currently: {}", "not installed".bold().red()).as_str());
                }
            }
        }

        _ => {
            if is_insult_enabled {
                warning(
                    format!(
                        "What the fuck do you want me to do with, {}? It's not a valid option!",
                        args.option.italic().cyan()
                    )
                    .as_str(),
                );
            } else {
                warning(format!("{} is not a valid option!", args.option.italic().cyan()).as_str());
            }
            return;
        }
    }
}
