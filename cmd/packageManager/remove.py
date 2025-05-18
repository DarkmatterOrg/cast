import typer
import shutil
import os
import subprocess

from typing_extensions import Annotated
from utils.checkIfRoot import checkIfRoot
from utils.logger import notice, error, success
from utils.isCmdInstalled import isCmdInstalled
from rich.console import Console

app = typer.Typer()
console = Console()

@app.command(rich_help_panel="Package Manager")
def remove(pkg: str, verbose: Annotated[bool, typer.Option("-v", "--verbose", help="Show more output", show_default=False)] = None):
  """
  Removes a package

  Example:
  [blue on black]cast remove firefox[/blue on black]
  """

  checkIfRoot()

  foundPM = False
  supportedPMs = ["pacman", "apt", "apk", "dnf", "emerge", "xbps-install", "scratch", "eopkg"]

  for pm in supportedPMs:
    if isCmdInstalled(pm):
      foundPM = True
      match pm:
        case "pacman":
          removePkg(f"pacman -Rs --noconfirm {pkg}", pkg, verbose)
        case "apt":
          removePkg(f"apt remove -y {pkg}", pkg, verbose)
        case "apk":
          removePkg(f"apk del {pkg}", pkg, verbose)
        case "dnf":
          removePkg(f"dnf remove -y {pkg}", pkg, verbose)
        case "emerge":
          removePkg(f"emerge {pkg}", pkg, verbose)
        case "xbps-install":
          removePkg(f"xbps-remove -y {pkg}", pkg, verbose)
        case "scratch":
          removePkg(f"scratch remove {pkg}", pkg, verbose)
        case "eopkg":
          removePkg(f"eopkg remove -y {pkg}", pkg, verbose)
        case _:
          error("This should never happen")
          raise typer.Exit(code=1)
  
  if not foundPM:
    notice("Could not find supported package manager")

def removePkg(cmdTorRun: str, pkg: str, verbose: bool):
  try:
    if verbose:
      subprocess.run(cmdTorRun, shell=True, check=True)
    else:
      with console.status(f"Removing..."):
        subprocess.run(cmdTorRun, stdout=open(os.devnull, "wb"), stderr=open(os.devnull, "wb"), shell=True, check=True)
  except:
    error(f"Failed to remove {pkg}")
  else:
    success(f"Removed {pkg}")