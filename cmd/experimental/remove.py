import typer
import shutil
import subprocess
import os

from utils.logger import info, notice, error
from rich.console import Console

app = typer.Typer()
console = Console()

@app.command(rich_help_panel="Experimentals")
def remove(pkg: str):
  """
  Removes a package
  """
  #TODO Support for more package managers

  if shutil.which("pacman"):
    info(f"Removing {pkg}")
    try:
      os.system(f"sudo pacman -Rs --noconfirm {pkg}")
    except:
      error(f"Failed to remove {pkg}")
      raise typer.Exit(code=1)
  else:
    notice("Could not find supported package manager")