use crate::args::DevArgs;
use crate::utils::{is_cmd_installed::is_cmd_installed, status_msg::warning};
use colored::Colorize;
use std::process::Command;

pub fn dev(args: &DevArgs) {
    if args.install && args.remove {
        eprintln!(
            "{}",
            "You can't use both --install and --remove at the same time."
                .bold()
                .red()
        );
        return;
    }

    match args.option.as_str() {
        "rustup" => {
            if args.install {
                if is_cmd_installed("rustup") {
                    println!(
                        "{}",
                        "Rustup is already installed, nothing todo.".bold().yellow()
                    );
                    return;
                }

                Command::new("sh")
                    .args([
                        "-c",
                        "curl --tlsv1.2 -sSf 'https://sh.rustup.rs' | sh -s -- -y",
                    ])
                    .status()
                    .expect("Was unable to install rustup");

                println!(
                    "{}",
                    "Rustup got installed, please restart the terminal to start using it."
                        .italic()
                        .yellow()
                );
            } else if args.remove {
                if !is_cmd_installed("rustup") {
                    println!(
                        "{}",
                        "Rustup is not installed, nothing todo.".bold().yellow()
                    );
                    return;
                }

                Command::new("rustup")
                    .args(["self", "uninstall"])
                    .status()
                    .expect("Wasn't able to uninstall rustup");

                println!(
                    "{}",
                    "Rustup got removed, please restart the terminal to start using it."
                        .italic()
                        .yellow()
                );
            } else {
                if is_cmd_installed("rustup") {
                    println!("Rustup is currently: {}", "installed".bold().green());
                } else {
                    println!("Rustup is currently: {}", "not installed".bold().red());
                }
            }
        }

        _ => {
            warning(format!("{} is not a valid option!", args.option.italic().cyan()).as_str());
            return;
        }
    }
}
