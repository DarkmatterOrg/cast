use crate::utils::status_msg::error;
use std::process::Command;

pub fn run_command(command: &str, arg: &str) {
    let full_command = format!("{} {}", command, arg);

    let status = Command::new("/bin/sh")
        .args(["-c", full_command.as_str()])
        .status()
        .expect("Failed to run the command");

    if !status.success() {
        error("Failed to run the command.")
    }
}
