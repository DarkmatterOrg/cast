use crate::args::AutoUpdateArgs;
use colored::Colorize;
use std::process::{Command, Stdio};

pub fn toggle_password_feedback(args: &AutoUpdateArgs) {
  if args.status {
    let mut current_status = "Disabled".red().bold();

    let current_status_result = Command::new("systemctl")
      .args(["is-enabled", "umbra-update.timer"])
      .stdout(Stdio::piped())
      .output()
      .unwrap();
    let get_current_status = String::from_utf8(current_status_result.stdout).unwrap();

    if get_current_status.contains("enabled") {
      current_status = "Enabled".green().bold();
    }

    println!("Automatic updates are currently: {}", current_status);
  }

  if args.enable && args.disable {
    eprintln!("{}", "You can't use both --enable and --disable at the same time.".bold().red());
    return;
  }

  if args.enable {
    Command::new("systemctl")
      .args(["enable", "umbra-update.timer"])
      .spawn()
      .expect("Failed to enable auto-updates");
    println!("Auto Updater have been: {}", "enabled".bold().green());
  }

  if args.disable {
    Command::new("systemctl")
      .args(["disable", "umbra-update.timer"])
      .spawn()
      .expect("Failed to disable auto-updates");
    println!("Auto Updater have been: {}", "disabled".bold().red());
  }
}
