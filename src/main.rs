mod args;
mod config;
mod modules;
mod utils;

use args::AutoUpdateArgs;
use args::DevUtilArgs;
use args::FixesArgs;
use args::PasswordFeedbackArgs;
use args::UpdateArgs;
use args::build_commands;
use colored::Colorize;
use modules::horizon::nix::install_nix;
use modules::nova::tailscale::disable_tailscale;
use modules::nova::tailscale::enable_tailscale;

use crate::modules::{
    auto_update::auto_update,
    bios::bios,
    clean_system::clean_system,
    dev::dev,
    fix::fix,
    horizon::dbox::{dbox_ubuntu, dbox_ubuntu_with_pkgs},
    toggle_password_feedback::toggle_password_feedback,
    update::update,
};

use crate::utils::{
    image_check_utils::is_correct_image,
    is_root::is_root,
    status_msg::{error, info, notice},
};

const VERSION: &str = clap::crate_version!();
fn main() {
    if !is_correct_image() {
        error("Cast can only be used on Umbra, Nova or Horizon based images.");
        return;
    }

    config::set::initialize_config();

    let matches = build_commands().get_matches();

    match matches.subcommand() {
        Some(("version", _)) => {
            println!("{}: v{}", "Cast".bold().purple(), VERSION.blue());
        }
        Some(("auto-update", args)) => {
            let auto_update_args = AutoUpdateArgs {
                status: args.get_one::<bool>("status").copied().unwrap_or(false),
                enable: args.get_one::<bool>("enable").copied().unwrap_or(false),
                disable: args.get_one::<bool>("disable").copied().unwrap_or(false),
            };

            if is_root() {
                auto_update(&auto_update_args);
            } else {
                notice("Please run this command with sudo.");
            }
        }
        Some(("bios", _)) => {
            bios();
        }
        Some(("clean-system", _)) => {
            clean_system();
        }
        Some(("config", _)) => {
            info(format!("Config can be found at: {:?}", config::config_path()).as_str());
        }
        Some(("dev-util", args)) => {
            let dev_args = DevUtilArgs {
                install: args.get_one::<String>("install").cloned(),
                remove: args.get_one::<String>("remove").cloned(),
            };

            dev(&dev_args);
        }
        Some(("fix", args)) => {
            let fix_args = FixesArgs {
                proton_hang: args
                    .get_one::<bool>("proton_hang")
                    .copied()
                    .unwrap_or(false),
                gmod: args.get_one::<bool>("gmod").copied().unwrap_or(false),
                discord: args.get_one::<bool>("discord").copied().unwrap_or(false),
                vesktop: args.get_one::<bool>("vesktop").copied().unwrap_or(false),
            };

            fix(&fix_args); // Pass the converted args to the fix function
        }
        Some(("password-feedback", args)) => {
            let feedback_args = PasswordFeedbackArgs {
                enable: args.get_one::<bool>("enable").copied().unwrap_or(false),
                disable: args.get_one::<bool>("disable").copied().unwrap_or(false),
            };

            if is_root() {
                toggle_password_feedback(&feedback_args);
            } else {
                notice("Please run this command with sudo.");
            }
        }
        Some(("update", args)) => {
            let update_args = UpdateArgs {
                system: args.get_one::<bool>("system").copied().unwrap_or(false),
                user: args.get_one::<bool>("user").copied().unwrap_or(false),
            };

            update(&update_args);
        }
        Some(("dbox-ubuntu", _)) => {
            dbox_ubuntu();
        }
        Some(("dbox-ubuntu-with-pkgs", args)) => {
            if let Some(pkgs) = args.get_one::<String>("pkgs") {
                dbox_ubuntu_with_pkgs(pkgs);
            } else {
                error("Packages are required, but none were provided.");
            }
        }
        Some(("install-nix", _)) => {
            install_nix();
        }
        Some(("enable-tailscale", _)) => {
            enable_tailscale();
        }
        Some(("disable-tailscale", _)) => {
            disable_tailscale();
        }
        None => {
            println!("No command provided.");
        }
        _ => {
            println!("Invalid command.");
        }
    }
}
