use crate::args::FixesArgs;
use colored::Colorize;
use std::fs::File;
use std::io::prelude::*;
use std::process::Command;
use std::{env, fs};

pub fn fix(args: &FixesArgs) {
    if args.discord {
        println!("{}", "Fixing Discord Flatpak RPC...".bold().yellow());

        fix_rpc("com.discordapp.Discord", "discord", "Discord");

        println!("{}", "Done!".bold().green());
    }

    if args.gmod {
        println!("{}", "Fixing gmod...".bold().yellow());

        Command::new("mkdir")
            .args(["-p", "/tmp/patch-gmod"])
            .status()
            .expect("Weren't able to create the tmp dir");

        Command::new("sh")
            .args(["-c", "wget $(curl -s https://api.github.com/repos/solsticegamestudios/GModCEFCodecFix/releases/latest | jq -r '.assets[] | select(.name | test(\"GModCEFCodecFix-Linux\")) | .browser_download_url') -P /tmp/patch-gmod"])
            .status()
            .expect("Wasn't able to fix GMod");

        Command::new("chmod")
            .args(["+x", "/tmp/patch-gmod/GModCEFCodecFix-Linux"])
            .status()
            .expect("Wasn't able to chmod");

        Command::new("/tmp/patch-gmod/GModCEFCodecFix-Linux")
            .status()
            .expect("Not able to run script");

        Command::new("rm")
            .args(["-rf", "/tmp/patch-gmod"])
            .status()
            .expect("Not able to remove /tmp/patch-gmod");

        println!("{}", "Done!".bold().green());
    }

    if args.proton_hang {
        println!("{}", "Fixing proton hang...".bold().yellow());

        let protoncore = [
            "pv-bwrap",
            "pressure-vessel",
            "reaper",
            "explorer.exe",
            "rpcss.exe",
            "plugplay.exe",
            "services.exe",
            "svchost.exe",
            "winedevice.exe",
            "winedevice.exe",
            "wineserver",
        ];

        for x in protoncore {
            Command::new("killall")
                .args(["-9", x])
                .status()
                .expect("Was unable to kill");
        }

        println!("{}", "Done!".bold().green());
    }

    if args.vesktop {
        println!("{}", "Fixing Vesktop Flatpak RPC...".bold().yellow());

        fix_rpc("dev.vencord.Vesktop", "vesktop", "Vesktop");

        println!("{}", "Done!".bold().green());
    }
}

fn fix_rpc(long_client: &str, short_client: &str, pretty_name: &str) {
    let user = env::var("USER").unwrap();
    let tmp_dir = format!("/home/{}/.config/user-tmpfiles.d", &user);

    match fs::create_dir(&tmp_dir) {
        Ok(_) => {
            println!("{}", "Directory created...".italic().yellow());
        }
        Err(_) => {
            println!(
                "{}",
                "Directory already exists, continuing...".italic().yellow()
            );
        }
    }

    let mut tmp_file =
        File::create(format!("{}/discord-rpc.conf", tmp_dir)).expect("Wasn't able to create file");
    tmp_file
        .write_all(
            format!(
                "L %t/discord-ipc-0 - - - - app/{}/discord-ipc-0",
                long_client
            )
            .as_bytes(),
        )
        .expect("Wasn't able to write to file");

    Command::new("systemctl")
        .args([
            "enable",
            "--user",
            "--now",
            "systemd-tmpfiles-setup.service",
        ])
        .status()
        .expect("Unable to start the service");

    Command::new("ln")
        .args([
            "-s",
            format!(
                "/run/user/1000/.flatpak/{}/xdg-run/discord-ipc-0",
                long_client
            )
            .as_str(),
            "/run/user/1000/discord-ipc-0",
        ])
        .status()
        .expect("Unable to link");

    println!("{}", "Need to use sudo...".bold().blue());

    Command::new("sudo")
        .args(["flatpak", "override", "--filesystem=xdg-run/discord-ipc-*"])
        .status()
        .expect("Unable to override discord-ipc");

    Command::new("sudo")
        .args([
            "flatpak",
            "override",
            format!("--filesystem=xdg-run/.flatpak/{}:create", long_client).as_str(),
        ])
        .status()
        .expect(format!("Unable to override {}", &pretty_name).as_str());

    println!("{}", "Adding fix to autostart...".italic().blue());

    Command::new("mkdir")
        .args(["-pv", format!("/home/{}/.config/autostart", user).as_str()])
        .status()
        .expect("Unable to make autostart dir");

    let mut autostart_file = File::create(format!(
        "/home/{}/.config/autostart/{}-rpc.desktop",
        user, short_client
    ))
    .expect("Wasn't able to create file");
    autostart_file
      .write_all(format!("[Desktop Entry]\nCategories=Utility;\nComment=Symlink for {pretty_name} RPC\nIcon={short_client}\nName={pretty_name}RPC\nStartupNotify=true\nTerminal=false\nType=Application\nExec=ln -s /run/user/1000/.flatpak/{long_client}/xdg-run/discord-ipc-0 /run/user/1000/discord-ipc-0").as_bytes())
      .expect("Wasn't able to write to file");
}
