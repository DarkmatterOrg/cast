use clap::{Arg, Args, Command, Parser, Subcommand};
use std::path::Path;

#[derive(Parser)]
#[command(about, long_about = None, arg_required_else_help = true)]
pub struct Cli {
    #[clap(subcommand)]
    pub command: Commands,
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
pub struct DevUtilArgs {
    pub install: Option<String>,
    pub remove: Option<String>,
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
pub struct UpdateArgs {
    #[arg(short, long)]
    pub system: bool,
    #[arg(short, long)]
    pub user: bool,
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
    /// Get the path where the config is
    Config,
    /// Install different programming utilities
    DevUtil(DevUtilArgs),
    /// Different fixes for various things
    Fix(FixesArgs),
    ///Toggles password prompt feedback in terminal, where sudo password prompts will display asterisks when enabled
    PasswordFeedback(PasswordFeedbackArgs),
    /// Run an update on user, system or both
    Update(UpdateArgs),
}

pub fn build_commands() -> Command {
    // Start by defining the basic command
    let mut cmd = Command::new("cast")
        .about("User utility tool for Darkmatter")
        .arg_required_else_help(true);

    if Path::new("/usr/share/horizon").exists() {
        cmd = cmd
            .subcommand(Command::new("dbox-ubuntu").about("Create a Ubuntu distrobox container"))
            .subcommand(
                Command::new("dbox-ubuntu-with-pkgs")
                    .arg(
                        Arg::new("pkgs")
                            .long("pkgs")
                            .help("List of packages to install")
                            .required(true)
                            .value_name("Packages"),
                    )
                    .about("Create a Ubuntu distrobox + user specified packages"),
            )
            .subcommand(
                Command::new("install-nix")
                    .about("Install Nix using the DeterminateSystems nix-installer"),
            );
    }

    if Path::new("/usr/share/nova").exists() {
        cmd = cmd
            .subcommand(Command::new("enable-tailscale").about("Enable Tailscale"))
            .subcommand(Command::new("disable-tailscale").about("Disable Tailscale"))
    }

    cmd = cmd
        .subcommand(Command::new("version").about("Show CLI version"))
        .subcommand(
            Command::new("auto-update")
                .about("Toggle the auto-updater")
                .arg_required_else_help(true)
                .arg(
                    Arg::new("status")
                        .short('s')
                        .long("status")
                        .help("Get current status of the auto-updater")
                        .num_args(0),
                )
                .arg(
                    Arg::new("enable")
                        .short('e')
                        .long("enable")
                        .help("Enable auto-updater")
                        .num_args(0),
                )
                .arg(
                    Arg::new("disable")
                        .short('d')
                        .long("disable")
                        .help("Disable auto-updater")
                        .num_args(0),
                ),
        )
        .subcommand(Command::new("bios").about("Boot into this device's BIOS/UEFI screen"))
        .subcommand(Command::new("clean-system").about("Clean up old and unused system files"))
        .subcommand(Command::new("config").about("Get the config path"))
        .subcommand(
            Command::new("dev-util")
                .about("Install different programming utilities")
                .arg(
                    Arg::new("install")
                        .short('i')
                        .long("install")
                        .help("Install the specified utility")
                        .value_name("UTILITY") // The name shown in help messages
                        .num_args(1) // Require exactly one argument
                        .required(false),
                )
                .arg(
                    Arg::new("remove")
                        .short('r')
                        .long("remove")
                        .help("Remove the specified utility")
                        .value_name("UTILITY")
                        .num_args(1)
                        .required(false),
                ),
        )
        .subcommand(
            Command::new("fix")
                .about("Different fixes for various things")
                .arg_required_else_help(true)
                .arg(
                    Arg::new("proton_hang")
                        .long("proton-hang")
                        .help("Kills all processes related to wine and proton")
                        .num_args(0),
                )
                .arg(
                    Arg::new("gmod")
                        .long("gmod")
                        .help("Patch GMod's 64-bit beta to work properly on Linux")
                        .num_args(0),
                )
                .arg(
                    Arg::new("discord")
                        .long("discord")
                        .help("Fix Discord flatpak RPC")
                        .num_args(0),
                )
                .arg(
                    Arg::new("vesktop")
                        .long("vesktop")
                        .help("Fix Vesktop flatpak RPC")
                        .num_args(0),
                ),
        )
        .subcommand(
            Command::new("password-feedback")
                .about("Toggle password feedback in the terminal")
                .arg_required_else_help(true)
                .arg(
                    Arg::new("enable")
                        .short('e')
                        .long("enable")
                        .help("Enable asterisks in password")
                        .num_args(0),
                )
                .arg(
                    Arg::new("disable")
                        .short('d')
                        .long("disable")
                        .help("Disable asterisks in password")
                        .num_args(0),
                ),
        )
        .subcommand(
            Command::new("update")
                .about("Run an update on user, system, or both")
                .arg_required_else_help(true)
                .arg(
                    Arg::new("system")
                        .short('s')
                        .long("system")
                        .help("Update the system")
                        .num_args(0),
                )
                .arg(
                    Arg::new("user")
                        .short('u')
                        .long("user")
                        .help("Update the user environment")
                        .num_args(0),
                ),
        );

    cmd
}
