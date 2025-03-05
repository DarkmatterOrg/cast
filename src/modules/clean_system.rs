use std::process::Command;

use crate::utils::status_msg::done;

pub fn clean_system() {
    run_cmd(
        "podman",
        &["image", "prune", "-af"],
        "Failed to prune podman images",
    );
    run_cmd(
        "podman",
        &["volume", "prune", "-f"],
        "Failed to prune podman volumes",
    );
    run_cmd(
        "flatpak",
        &["uninstall", "--unused"],
        "Failed to uninstall unused flatpaks",
    );
    run_cmd(
        "rpm-ostree",
        &["cleanup", "-bm"],
        "Failed to cleanup rpm-ostree",
    );
}

fn run_cmd(cmd: &str, args: &[&str], msg: &str) {
    Command::new(&cmd).args(&*args).spawn().expect(&msg);
    done(format!("{}", &cmd).as_str());
}
