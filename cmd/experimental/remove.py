import typer
import shutil
import os
import subprocess

from utils.checkIfRoot import checkIfRoot
from utils.logger import info, notice, error, success
from rich.console import Console

app = typer.Typer()
console = Console()

@app.command(rich_help_panel="Experimentals")
def remove(pkg: str):
  """
  Removes a package

  Example:
  [blue on black]cast remove firefox[/blue on black]
  """
  #TODO Support for more package managers

  checkIfRoot()

  if shutil.which("pacman"):
    info(f"Removing {pkg}")
    try:
      with console.status("Removing..."):
        subprocess.getoutput(f"pacman -Rs --noconfirm {pkg}")
    except:
      error(f"Failed to remove {pkg}")
      raise typer.Exit(code=1)
    else:
      success(f"Removed {pkg}")
  else:
    notice("Could not find supported package manager")