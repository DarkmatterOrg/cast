use crate::utils::{
    is_cmd_installed::is_cmd_installed, run_command::run_command, status_msg::error,
};

pub fn dbox_ubuntu() {
    if !is_cmd_installed("distrobox") {
        error("Distrobox isn't installed.");
        return;
    }

    run_command(
        "distrobox-create",
        format!("--image quay.io/toolbx/ubuntu-toolbox:latest --name ubuntu").as_str(),
    );
}

pub fn dbox_ubuntu_with_pkgs(pkgs: &str) {
    if !is_cmd_installed("distrobox") {
        error("Distrobox isn't installed.");
        return;
    }

    run_command(
        "distrobox-create",
        format!(
            "--image quay.io/toolbx/ubuntu-toolbox:latest --name ubuntu --additional-packages '{}'",
            pkgs
        )
        .as_str(),
    );
}
