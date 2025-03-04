use colored::Colorize;

pub fn done(msg: &str) {
    println!("{}: {}", "DONE".bold().green(), msg);
}

pub fn info(msg: &str) {
    println!("{}: {}", "INFO".bold().cyan(), msg);
}

pub fn notice(msg: &str) {
    println!("{}: {}", "NOTICE".bold().blue(), msg);
}

pub fn warning(msg: &str) {
    println!("{}: {}", "WARNING".bold().yellow(), msg);
}

pub fn error(msg: &str) {
    println!("{}: {}", "ERROR".bold().red(), msg);
}
