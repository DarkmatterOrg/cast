import typer
import shutil
import os
import subprocess

from typing_extensions import Annotated
from utils.checkIfRoot import checkIfRoot
from utils.logger import info, notice, error, success
from rich.console import Console

app = typer.Typer()
console = Console()

@app.command(rich_help_panel="Experimentals")
def remove(pkg: str, verbose: Annotated[bool, typer.Option("-v", "--verbose", help="Show more output", show_default=False)] = None):
  """
  Removes a package

  Example:
  [blue on black]cast remove firefox[/blue on black]
  """
  #TODO Support for more package managers

  checkIfRoot()

  if shutil.which("pacman"):
    info(f"Removing {pkg}")
    removePkg(f"pacman -Rs --noconfirm {pkg}", pkg, verbose)
  else:
    notice("Could not find supported package manager")

def removePkg(cmdTorRun: str, pkg: str, verbose: bool):
  try:
    if verbose:
      subprocess.run(cmdTorRun, shell=True, check=True)
    else:
      with console.status(f"Removing..."):
        subprocess.run(cmdTorRun, stdout=open(os.devnull, "wb"), shell=True, check=True)
  except:
    error(f"Failed to remove {pkg}")
  else:
    success(f"Removed {pkg}")