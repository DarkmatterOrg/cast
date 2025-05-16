import typer
import shutil
import subprocess
import os

from utils.logger import info, notice, error
from rich.console import Console

app = typer.Typer()
console = Console()

@app.command(rich_help_panel="Experimentals")
def install(pkg: str):
  """
  Installs a package
  """
  #TODO Support for more package managers

  if shutil.which("pacman"):
    if shutil.which("yay"):
      info("Installing with yay")
      try:
        os.system(f"yay -Syy --noconfirm {pkg}")
      except:
        error(f"Failed to install {pkg}")
        raise typer.Exit(code=1)
    elif shutil.which("paru"):
      info("Installing with paru")
      try:
        os.system(f"paru -Syy --noconfirm {pkg}")
      except:
        error(f"Failed to install {pkg}")
        raise typer.Exit(code=1)
    else:
      info("Installing with pacman")
      try:
        os.system(f"sudo pacman -Syy --noconfirm {pkg}")
      except:
        error(f"Failed to install {pkg}")
        raise typer.Exit(code=1)
  else:
    notice("Could not find supported package manager")