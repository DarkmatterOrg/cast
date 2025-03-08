use crate::args::DevUtilArgs;
use crate::config;
use crate::utils::status_msg::{done, notice};
use crate::utils::{is_cmd_installed::is_cmd_installed, status_msg::warning};
use colored::Colorize;
use std::process::Command;

pub fn dev(args: &DevUtilArgs) {
    let is_insult_enabled = config::get::get_config().insults;

    // Ensure install and remove are not used at the same time
    if args.install.is_some() && args.remove.is_some() {
        if is_insult_enabled {
            warning(
                "What the fuck are you doing? You can't use --install and --remove at the same time!",
            );
        } else {
            warning("You can't use both --install and --remove at the same time.");
        }
        return;
    }

    // Handle `--install` case
    if let Some(package) = &args.install {
        match package.as_str() {
            "rustup" => {
                if is_cmd_installed("rustup") {
                    if is_insult_enabled {
                        notice("Are you dumb? Rustup is already installed...");
                    } else {
                        notice("Rustup is already installed, nothing to do.");
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
            }
            _ => {
                if is_insult_enabled {
                    warning(&format!(
                        "What the fuck do you want me to do with, {}? It's not a valid option!",
                        package.italic().cyan()
                    ));
                } else {
                    warning(&format!(
                        "{} is not a valid option!",
                        package.italic().cyan()
                    ));
                }
            }
        }
        return;
    }

    // Handle `--remove` case
    if let Some(package) = &args.remove {
        match package.as_str() {
            "rustup" => {
                if !is_cmd_installed("rustup") {
                    if is_insult_enabled {
                        notice("Are you dumb? Rustup isn't even installed...");
                    } else {
                        notice("Rustup is not installed, nothing to do.");
                    }
                    return;
                }

                Command::new("rustup")
                    .args(["self", "uninstall"])
                    .status()
                    .expect("Wasn't able to uninstall rustup");

                done("Rustup got removed, please restart the terminal.");
            }
            _ => {
                if is_insult_enabled {
                    warning(&format!(
                        "What the fuck do you want me to do with, {}? It's not a valid option!",
                        package.italic().cyan()
                    ));
                } else {
                    warning(&format!(
                        "{} is not a valid option!",
                        package.italic().cyan()
                    ));
                }
            }
        }
        return;
    }
}
