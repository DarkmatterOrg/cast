from typing import Annotated

import typer
import os

from datetime import datetime
from rich.console import Console
from utils.logger import info

app = typer.Typer()
console = Console()

@app.command(rich_help_panel="Base")
def age(
        getEpoch: Annotated[bool, typer.Option("-g", "--getEpoch", help="Get the epoch time instead")] = None,
        epoch: Annotated[int, typer.Option("-e", "--epoch", help="Get system age based on custom epoch", show_default=False)] = None
        ):
  """
  Show the system's age

  [strike]Stolen[/strike] Taken from Alxhr0

  Example:
    [blue on black]cast age --epoch 1742944940[/blue on black]
  """
  if os.path.exists("/bedrock"):
      sysStat = os.stat("/home")
  elif os.path.exists("/ostree"):
      sysStat = os.stat("/ostree")
  else:
      sysStat = os.stat("/")

  if getEpoch:
    info(int(sysStat.st_ctime))
  elif epoch:
      console.print(getInstallTime(epoch))
  else:
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