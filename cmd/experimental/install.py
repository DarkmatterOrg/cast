import typer
import shutil
import os
import subprocess

from utils.logger import info, notice, error, success
from utils.checkIfRoot import checkIfRoot
from rich.console import Console

app = typer.Typer()
console = Console()

@app.command(rich_help_panel="Experimentals")
def install(pkg: str):
  """
  Installs a package
  """
  #TODO Support for more package managers

  checkIfRoot()

  if shutil.which("pacman"):
    if shutil.which("yay"):
      info("Installing with yay")
      try:
        with console.status("Installing..."):
          subprocess.getoutput(f"yay -Syy --noconfirm {pkg}")
      except:
        error(f"Failed to install {pkg}")
        raise typer.Exit(code=1)
      else:
        success(f"Installed {pkg}")
    elif shutil.which("paru"):
      info("Installing with paru")
      try:
        with console.status("Installing..."):
          subprocess.getoutput(f"paru -Syy --noconfirm {pkg}")
      except:
        error(f"Failed to install {pkg}")
        raise typer.Exit(code=1)
      else:
        success(f"Installed {pkg}")
    else:
      info("Installing with pacman")
      try:
        with console.status("Installing..."):
          subprocess.getoutput(f"pacman -Syy --noconfirm {pkg}")
      except:
        error(f"Failed to install {pkg}")
        raise typer.Exit(code=1)
      else:
        success(f"Installed {pkg}")
  else:
    notice("Could not find supported package manager")