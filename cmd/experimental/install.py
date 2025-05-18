import typer
import shutil
import os
import subprocess

from typing_extensions import Annotated

from utils.logger import info, notice, error, success, warn
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
      info("Installing with yay")
      try:
        if verbose:
          os.system(f"yay -Syy --noconfirm {pkg}")
        else:
          with console.status("Installing..."):
            output = subprocess.getoutput(f"yay -Syy --noconfirm {pkg}")

            if f"No AUR package found for {pkg}" in output or f"error: target not found: '{pkg}'" in output:
              warn(f"No package found for {pkg}")
              raise typer.Exit(code=1)
      except:
        error(f"Failed to install {pkg}")
        raise typer.Exit(code=1)
      else:
        success(f"Installed {pkg}")
    elif shutil.which("paru"):
      info("Installing with paru")
      try:
        if verbose:
          os.system(f"paru -Syy --noconfirm {pkg}")
        else:
          with console.status("Installing..."):
            output = subprocess.getoutput(f"paru -Syy --noconfirm {pkg}")

            if f"No AUR package found for {pkg}" in output or f"error: target not found: '{pkg}'" in output:
              warn(f"No package found for {pkg}")
              raise typer.Exit(code=1)
      except:
        error(f"Failed to install {pkg}")
        raise typer.Exit(code=1)
      else:
        success(f"Installed {pkg}")
    else:
      info("Installing with pacman")
      try:
        if verbose:
          os.system(f"pacman -Syy --noconfirm {pkg}")
        else:
          with console.status("Installing..."):
            output =subprocess.getoutput(f"pacman -Syy --noconfirm {pkg}")

            if f"No AUR package found for {pkg}" in output or f"error: target not found: '{pkg}'" in output:
              warn(f"No package found for {pkg}")
              raise typer.Exit(code=1)
      except:
        error(f"Failed to install {pkg}")
        raise typer.Exit(code=1)
      else:
        success(f"Installed {pkg}")
  else:
    notice("Could not find supported package manager")