import typer
import shutil
import os
import subprocess

from typing_extensions import Annotated

from utils.logger import info, notice, error, success, warn, debug
from utils.checkIfRoot import checkIfRoot
from rich.console import Console

app = typer.Typer()
console = Console()

@app.command(rich_help_panel="Experimentals")
def install(pkg: str, verbose: Annotated[bool, typer.Option("-V", "--verbose", help="Show more output", show_default=False)] = None):
  """
  Installs a package

  Example:
  [blue on black]cast install firefox[/blue on black]
  """
  #TODO Support for more package managers

  checkIfRoot()

  if shutil.which("pacman"):
    if shutil.which("yay"):
      info("Using yay")
      installPkgCmd(f"yay -Syy --noconfirm {pkg}", pkg, verbose, True)
    elif shutil.which("paru"):
      info("Using paru")
      installPkgCmd(f"paru -Syy --noconfirm {pkg}", pkg, verbose)
    else:
      info("Using pacman")
      installPkgCmd(f"pacman -Syy --noconfirm {pkg}", pkg, verbose)
  else:
    notice("Could not find supported package manager")

def archCheckNoPkg(pkg, output: str):
  if f"No AUR package found for {pkg}" in output or f"error: target not found: '{pkg}'" in output:
    warn(f"No package found for {pkg}")
    raise typer.Exit(code=1)

def installPkgCmd(cmdToRun, pkg: str, verbose: bool, isYay: bool = False):
  try:
    if verbose:
      if isYay:
        info("No realtime messages cause yay sucks, the output will be shown soon...")
        output = subprocess.run(cmdToRun, shell=True, check=True, capture_output=True)
        print(output.stdout.decode("utf-8"))
        if f"No AUR package found for {pkg}" in output.stderr.decode("utf-8"):
          raise Exception(output.stderr.decode("utf-8"))
      else:
        subprocess.run(cmdToRun, shell=True, check=True)
    else:
      with console.status(f"Installing {pkg}..."):
        if isYay:
          output = subprocess.run(cmdToRun, shell=True, check=True, capture_output=True)
          if f"No AUR package found for {pkg}" in output.stderr.decode("utf-8"):
            raise Exception(output.stderr.decode("utf-8"))
        else:
          subprocess.run(cmdToRun, stdout=open(os.devnull, "wb"), stderr=open(os.devnull, "wb"), shell=True, check=True)
  except:
    error(f"Failed to install {pkg}")
  else:
    success(f"Installed {pkg}")