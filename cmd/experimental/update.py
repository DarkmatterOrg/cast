import typer
import shutil
import subprocess
import os

from utils.logger import info, notice, error
from rich.console import Console

app = typer.Typer()
console = Console()

@app.command(rich_help_panel="Experimentals")
def update():
  """
  Update system
  """
  #TODO Support for more package managers

  if shutil.which("pacman"):
    try:
      os.system("sudo pacman -Syyu --noconfirm")
    except:
      error("Failed to update system")
      raise typer.Exit(code=1)
  else:
    notice("Could not find supported package manager")