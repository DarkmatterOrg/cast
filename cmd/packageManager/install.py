import typer
import os
import subprocess

from typing_extensions import Annotated

from utils.logger import notice, error, success
from utils.checkIfRoot import checkIfRoot
from utils.isCmdInstalled import isCmdInstalled
from rich.console import Console

app = typer.Typer()
console = Console()

@app.command(rich_help_panel="Package Manager")
def install(pkg: str, verbose: Annotated[bool, typer.Option("-v", "--verbose", help="Show more output", show_default=False)] = None):
  """
  Installs a package

  Example:
  [blue on black]cast install firefox[/blue on black]
  """

  checkIfRoot()

  foundPM = False
  supportedPMs = ["pacman", "apt", "apk", "dnf", "emerge", "xbps-install", "scratch", "eopkg"]

  for pm in supportedPMs:
    if isCmdInstalled(pm):
      foundPM = True
      match pm:
        case "pacman":
          installPkgCmd(f"pacman -Syy --noconfirm {pkg}", pkg, verbose)
        case "apt":
          installPkgCmd(f"apt install -y {pkg}", pkg, verbose)
        case "apk":
          installPkgCmd(f"apk add {pkg}", pkg, verbose)
        case "dnf":
          installPkgCmd(f"dnf install -y {pkg}", pkg, verbose)
        case "emerge":
          installPkgCmd(f"emerge {pkg}", pkg, verbose)
        case "xbps-install":
          installPkgCmd(f"xbps-install -S -y {pkg}", pkg, verbose)
        case "scratch":
          installPkgCmd(f"scratch install {pkg}", pkg, verbose)
        case "eopkg":
          installPkgCmd(f"eopkg install -y {pkg}", pkg, verbose)
        case _:
          error("This should never happen")
          raise typer.Exit(code=1)

  if not foundPM:
    notice("Could not find supported package manager")

def installPkgCmd(cmdToRun, pkg: str, verbose: bool):
  try:
    if verbose:
      subprocess.run(cmdToRun, shell=True, check=True)
    else:
      with console.status(f"Installing {pkg}..."):
        subprocess.run(cmdToRun, stdout=open(os.devnull, "wb"), stderr=open(os.devnull, "wb"), shell=True, check=True)
  except:
    error(f"Failed to install {pkg}")
  else:
    success(f"Installed {pkg}")