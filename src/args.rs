use clap::{Args, Parser, Subcommand};

/// Umrba's custom CLI
#[derive(Parser)]
#[command(long_about = None, arg_required_else_help = true)]
pub struct Cli {
  #[clap(subcommand)]
  pub command: Commands,
}

#[derive(Args)]
pub struct UpdateArgs {
  /// User, System or Both
  #[arg(short, long)]
  pub option: String,
}

#[derive(Args)]
pub struct AutoUpdateArgs {
  /// Get current status
  #[arg(short, long)]
  pub status: bool,
  /// Enable auto updates
  #[arg(short, long)]
  pub enable: bool,
  /// Disable auto updates
  #[arg(short, long)]
  pub disable: bool,
}

#[derive(Args)]
pub struct TogglePasswordArgs {
  /// Get current status
  #[arg(short, long)]
  pub status: bool,
  /// Enable asterisks in password
  #[arg(short, long)]
  pub enable: bool,
  /// Disable asterisks in password
  #[arg(short, long)]
  pub disable: bool,
}

#[derive(Subcommand)]
pub enum Commands {
  /// Show CLI version
  Version,
  /// Toggle the auto-updater
  AutoUpdate(AutoUpdateArgs),
  /// Boot into this device's BIOS/UEFI screen
  Bios,
  ///Clean up old up unused podman images, volumes, flatpak packages and rpm-ostree content
  CleanSystem,
  ///Toggles password prompt feedback in terminal, where sudo password prompts will display asterisks when enabled
  TogglePassword(TogglePasswordArgs),
  /// Run an update on user, system or both
  Update(UpdateArgs),
}
