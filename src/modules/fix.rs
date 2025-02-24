use crate::args::FixesArgs;
use colored::Colorize;
use std::process::Command;

pub fn fix(args: &FixesArgs) {
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
}
