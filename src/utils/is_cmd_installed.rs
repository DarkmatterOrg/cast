use which::which;

pub fn is_cmd_installed(cmd: &str) -> bool {
    match which(cmd) {
        Ok(_) => {
            return true;
        }
        Err(_) => {
            return false;
        }
    };
}
