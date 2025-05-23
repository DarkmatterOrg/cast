import typer
import shutil
import os

from utils.logger import info, notice, error, success
from utils.checkIfRoot import checkIfRoot
from rich.console import Console

app = typer.Typer()
console = Console()

@app.command(rich_help_panel="Experimentals")
def update():
  """
  Update system
  """
  #TODO Support for more package managers

  checkIfRoot()

  if shutil.which("pacman"):
    try:
      os.system("pacman -Syyu --noconfirm")
    except:
      error("Failed to update system")
      raise typer.Exit(code=1)
    else:
      success("System updated")
  else:
    notice("Could not find supported package manager")