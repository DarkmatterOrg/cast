use crate::args::PasswordFeedbackArgs;
use colored::Colorize;
use std::fs;
use std::io::Write;

pub fn toggle_password_feedback(args: &PasswordFeedbackArgs) {
    if args.enable && args.disable {
        eprintln!(
            "{}",
            "You can't use both --enable and --disable at the same time."
                .bold()
                .red()
        );
        return;
    }

    if args.enable {
        let mut file = fs::File::create("/etc/sudoers.d/enable-pwfeedback")
            .expect("Wasn't able to create file");
        file.write_all(b"Defaults pwfeedback")
            .expect("Wasn't able to write to file");

        println!(
            "Password feedback is now {}! Restart terminal to see changes.",
            "enabled".bold().green()
        );
    }

    if args.disable {
        let res = fs::remove_file("/etc/sudoers.d/enable-pwfeedback");

        match res {
            Ok(_) => println!(
                "Password feedback is now {}! Restart terminal to see changes.",
                "disabled".bold().red()
            ),
            Err(_) => eprintln!(
                "{}",
                "File didn't exist or wasn't able to be removed"
                    .bold()
                    .red()
            ),
        }
    }
}
