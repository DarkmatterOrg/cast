use crate::args::DevArgs;
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
                println!("{}", "Installing rustup...".yellow());

                Command::new("sh")
                    .args([
                        "-c",
                        "curl --tlsv1.2 -sSf 'https://sh.rustup.rs' | sh -s -- -y",
                    ])
                    .status()
                    .expect("Was unable to install rustup");
            } else if args.remove {
                Command::new("rustup")
                    .args(["self", "uninstall"])
                    .status()
                    .expect("Wasn't able to uninstall rustup");
            } else {
                let res = Command::new("command")
                    .args(["-v", "rupstup", "2>&1 >/dev/null"])
                    .output()
                    .expect("Wasn't able to check if command exists");

                println!("{:?}", res.stdout);
            }
        }

        _ => {
            eprintln!(
                "{}: {} is not a valid option!",
                "WARNING".bold().yellow(),
                args.option.italic().cyan()
            );
            return;
        }
    }
}
