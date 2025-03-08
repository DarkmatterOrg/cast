use crate::utils::run_command::run_command;

pub fn enable_tailscale() {
    run_command("systemctl", format!("enable --now tailscale").as_str());
}

pub fn disable_tailscale() {
    run_command("systemctl", format!("disable --now tailscale").as_str());
}
