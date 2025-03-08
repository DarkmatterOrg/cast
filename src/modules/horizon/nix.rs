use crate::utils::run_command::run_command;

pub fn install_nix() {
    // Install Nix
    run_command(
        "curl",
        format!(
            "--proto '=https' --tlsv1.2 -sSf -L https://install.determinate.systems/nix | sh -s -- install"
        ).as_str(),
    );
}
