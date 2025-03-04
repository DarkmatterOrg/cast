use clap::{Args, Parser, Subcommand};

/// Umrba's custom CLI
#[derive(Parser)]
#[command(long_about = None, arg_required_else_help = true)]
pub struct Cli {
    #[clap(subcommand)]
    pub command: Commands,
}

#[derive(Args)]
#[command(arg_required_else_help = true)]
pub struct UpdateArgs {
    /// User, System or Both
    #[arg(short, long)]
    pub option: String,
}

#[derive(Args)]
#[command(arg_required_else_help = true)]
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
#[command(arg_required_else_help = true)]
pub struct PasswordFeedbackArgs {
    /// Enable asterisks in password
    #[arg(short, long)]
    pub enable: bool,
    /// Disable asterisks in password
    #[arg(short, long)]
    pub disable: bool,
}

#[derive(Args)]
#[command(arg_required_else_help = true)]
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

#[derive(Args)]
#[command(arg_required_else_help = true)]
pub struct FixesArgs {
    /// Kills all processes related to wine and proton. This forces it to restart next time you launch the game (you might still have to press STOP in steam to kill the game binary)
    #[arg(short, long)]
    pub proton_hang: bool,
    /// Patch GMod's 64-bit beta to work properly on Linux (https://github.com/solsticegamestudios/GModCEFCodecFix)
    #[arg(short, long)]
    pub gmod: bool,
    /// Fix Discord flatpak RPC
    #[arg(short, long)]
    pub discord: bool,
    /// Fix Vesktop flatpak RPC
    #[arg(short, long)]
    pub vesktop: bool,
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
    /// Different fixes for various things
    Fix(FixesArgs),
    ///Toggles password prompt feedback in terminal, where sudo password prompts will display asterisks when enabled
    PasswordFeedback(PasswordFeedbackArgs),
    /// Run an update on user, system or both
    Update(UpdateArgs),
}
