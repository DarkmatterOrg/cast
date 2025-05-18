import typer
import shutil
import subprocess

from utils.logger import info, notice, error
from utils.checkIfRoot import checkIfRoot
from rich.console import Console

app = typer.Typer()
console = Console()

@app.command(rich_help_panel="Package Manager")
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
      searchPkg(f"yay -Ss {pkg}", pkg)
    elif shutil.which("paru"):
      info("Searching with paru")
      searchPkg(f"paru -Ss {pkg}", pkg)
    else:
      info("Searching with pacman")
      searchPkg(f"pacman -Ss {pkg}", pkg)
  else:
    notice("Could not find supported package manager")

def searchPkg(cmdToRun: str, pkg: str):
  try:
    subprocess.run(cmdToRun, shell=True, check=True)
  except:
    error(f"Failed to search for {pkg}")
    raise typer.Exit(code=1)