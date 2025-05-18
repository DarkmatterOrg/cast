import typer
import shutil
import os

from utils.logger import info, notice, error, warn
from utils.checkIfRoot import checkIfRoot
from rich.console import Console

app = typer.Typer()
console = Console()

@app.command(rich_help_panel="Experimentals")
def search(pkg: str):
  """
  Search for a package

  Example:
  [blue on black]cast search firefox[/blue on black]
  """
  #TODO Support for more package managers

  checkIfRoot()

  if shutil.which("pacman"):
    if shutil.which("yay"):
      info("Searching with yay")
      try:
        os.system(f"yay -Ss {pkg}")
      except:
        error(f"Failed to search for {pkg}")
        raise typer.Exit(code=1)
    elif shutil.which("paru"):
      info("Searching with paru")
      try:
        os.system(f"paru -Ss {pkg}")
      except:
        error(f"Failed to search for {pkg}")
        raise typer.Exit(code=1)
    else:
      info("Searching with pacman")
      try:
        os.system(f"pacman -Ss {pkg}")
      except:
        error(f"Failed to search for {pkg}")
        raise typer.Exit(code=1)
  else:
    notice("Could not find supported package manager")