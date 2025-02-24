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
pub struct PasswordFeedbackArgs {
    /// Enable asterisks in password
    #[arg(short, long)]
    pub enable: bool,
    /// Disable asterisks in password
    #[arg(short, long)]
    pub disable: bool,
}

#[derive(Args)]
pub struct DevArgs {
    /// Values: rustup
    pub option: String,
    /// Install the utility
    #[arg(short, long)]
    pub install: bool,
    /// Remove the utility
    #[arg(short, long)]
    pub remove: bool,
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
    /// Install different programming utilities
    Dev(DevArgs),
    ///Toggles password prompt feedback in terminal, where sudo password prompts will display asterisks when enabled
    PasswordFeedback(PasswordFeedbackArgs),
    /// Run an update on user, system or both
    Update(UpdateArgs),
}
