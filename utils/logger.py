import datetime

from rich.console import Console

DATE = datetime.datetime.now().strftime("%H:%M:%S")

console = Console(highlight=False)

def info(message: str):
    console.print(f"[white]{DATE}[/white] [bold blue]INFO[/bold blue] {message}")

def notice(message: str):
    console.print(f"[white]{DATE}[/white] [bold cyan]NOTICE[/bold cyan] {message}")

def warn(message: str):
    console.print(f"[white]{DATE}[/white] [bold yellow]WARNING[/bold yellow] {message}")

def importantWarn(message: str):
    console.print(f"[white]{DATE}[/white] [bold yellow]WARNING[/bold yellow] [bold black on yellow]{message}[/bold black on yellow]")

def error(message: str):
    console.print(f"[white]{DATE}[/white] [bold red]ERROR[/bold red] {message}")

def fatal(message: str):
    console.print(f"[white]{DATE}[/white] [bold red]FATAL[/bold red] {message}")

def debug(message: str):
    console.print(f"[white]{DATE}[/white] [bold magenta]DEBUG[/bold magenta] {message}")

def success(message: str):
    console.print(f"[white]{DATE}[/white] [bold green]SUCCESS[/bold green] {message}")
