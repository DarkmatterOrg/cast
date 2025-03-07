use crate::{
    args::AutoUpdateArgs,
    utils::{
        image_check_utils::get_image_type,
        status_msg::{error, info},
    },
};
use colored::Colorize;
use std::process::{Command, Stdio};

pub fn auto_update(args: &AutoUpdateArgs) {
    if args.status {
        let mut current_status = "Disabled".red().bold();

        let current_status_result = Command::new("systemctl")
            .args(["is-enabled", get_timer_file_name()])
            .stdout(Stdio::piped())
            .output()
            .unwrap();

        let get_current_status = String::from_utf8(current_status_result.stdout).unwrap();

        if get_current_status.contains("enabled") {
            current_status = "Enabled".green().bold();
        }

        info(format!("Automatic updates are currently: {}", current_status).as_str());
    }

    if args.enable && args.disable {
        eprintln!(
            "{}",
            "You can't use both --enable and --disable at the same time."
                .bold()
                .red()
        );
        return;
    }

    if args.enable {
        Command::new("systemctl")
            .args(["enable", get_timer_file_name()])
            .spawn()
            .expect("Failed to enable auto-updates");
        info(format!("Auto Updater have been: {}", "enabled".bold().green()).as_str());
    }

    if args.disable {
        Command::new("systemctl")
            .args(["disable", get_timer_file_name()])
            .spawn()
            .expect("Failed to disable auto-updates");
        info(format!("Auto Updater have been: {}", "disabled".bold().red()).as_str());
    }
}

fn get_timer_file_name() -> &'static str {
    let image_type = get_image_type();

    if image_type.contains("umbra") {
        return "umbra-update.timer";
    } else if image_type.contains("nova") {
        return "nova-update.timer";
    } else if image_type.contains("aster") || image_type.contains("arcturus") {
        return "horizon-update.timer";
    } else {
        error("No correct image type!");
        panic!();
    }
}
