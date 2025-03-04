use colored::Colorize;

pub fn warning(msg: &str) {
    println!("{}: {}", "WARNING".bold().yellow(), msg);
}

pub fn error(msg: &str) {
    println!("{}: {}", "ERROR".bold().red(), msg);
}
