import typer
import os

from datetime import datetime
from rich.console import Console

app = typer.Typer()
console = Console()

@app.command(rich_help_panel="Base")
def age():
  """
  Show the system's age

  ~~Stolen~~ Taken from Alxhr0
  """
  if os.path.exists("/bedrock"):
      sysStat = os.stat("/home")
  elif os.path.exists("/ostree"):
      sysStat = os.stat("/ostree")
  else:
      sysStat = os.stat("/")

  console.print(getInstallTime(sysStat.st_ctime))

def getInstallTime(installed_stamp: float) -> str:
    installed = datetime.fromtimestamp(installed_stamp)
    now = datetime.now()
    diff = now - installed

    years = diff.days // 365
    months = (diff.days % 365) // 30
    days = (diff.days % 365) % 30
    hours = diff.seconds // 3600
    minutes = (diff.seconds % 3600) // 60
    seconds = diff.seconds % 60

    colors = ["red", "yellow", "green", "cyan", "blue", "magenta"]
    timeComponents = [
        (years, "Year"),
        (months, "Month"),
        (days, "Day"),
        (hours, "Hour"),
        (minutes, "Minute"),
        (seconds, "Second")
    ]
    
    parts = []
    for idx, (value, unit) in enumerate(timeComponents):
        if value > 0:
            color = colors[idx]
            plural = "s" if value != 1 else ""
            parts.append(
                f"[bold {color}]{value}[/bold {color}] [{color}]{unit}{plural}[/{color}]"
            )
    
    if not parts:
        return "[bold]0[/][magenta] Seconds[/magenta]"

    return " ".join(parts[:6])