use std::path::Path;
use colored::Colorize;
use std::process::Command;

pub fn bios() {
  if !Path::new("/sys/firmware/efi").exists() {
    eprintln!("{}", "Rebooting to legacy BIOS from OS is not supported.".bold().red());
    return;
  } else {
    Command::new("systemctl")
      .args(["reboot", "--firmware-setup"])
      .spawn()
      .expect("Failed to run reboot");
  }
}
