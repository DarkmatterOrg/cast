use crate::args::UpdateArgs;
use colored::Colorize;
use std::process::Command;

pub fn update(args: &UpdateArgs){
  let update_option = args.option.to_lowercase();

  if update_option == "user" {
    update_user();
  } else if update_option == "system" {
    update_system();
  } else if update_option == "both" {
    update_system();
    update_user();
  } else {
    eprintln!("{}: {} is not a valid option!", "ERROR".bold().red(), update_option.cyan().italic());
    return;
  }
}

fn update_user(){
  let mut update_user_cmd = Command::new("nebula")
  .arg("update-system")
  .spawn()
  .expect("Failed to start updating the user");

  let status = update_user_cmd.wait().expect("Failed to update the user");

  if !status.success() {
    eprintln!(
      "{}: Failed to update the user!",
      "ERROR".bold().red()
    );
  }
}

fn update_system(){
  let mut update_system_cmd = Command::new("sudo")
  .args(["nebula", "update-system"])
  .spawn()
  .expect("Failed to start updating the system");

  let status = update_system_cmd.wait().expect("Failed to update the system");

  if !status.success() {
    eprintln!(
      "{}: Failed to update the system!",
      "ERROR".bold().red()
    );
  }
}
