use crate::args::UpdateArgs;
use crate::utils::status_msg::error;
use std::process::Command;

pub fn update(args: &UpdateArgs) {
    if args.system {
        update_system();
    }

    if args.user {
        update_user();
    }
}

fn update_user() {
    let mut update_user_cmd = Command::new("nebula")
        .arg("update-system")
        .spawn()
        .expect("Failed to start updating the user");

    let status = update_user_cmd.wait().expect("Failed to update the user");

    if !status.success() {
        error("Failed to update the user!");
    }
}

fn update_system() {
    let mut update_system_cmd = Command::new("sudo")
        .args(["nebula", "update-system"])
        .spawn()
        .expect("Failed to start updating the system");

    let status = update_system_cmd
        .wait()
        .expect("Failed to update the system");

    if !status.success() {
        error("Failed to update the system!");
    }
}
