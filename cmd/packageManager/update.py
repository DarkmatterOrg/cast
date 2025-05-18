import typer
import subprocess

import utils.isCmdInstalled as isCmdInstalled
from utils.logger import info, notice, error, success
from utils.checkIfRoot import checkIfRoot
from rich.console import Console

app = typer.Typer()
console = Console()

@app.command(rich_help_panel="Experimentals")
def update():
  """
  Update system
  """

  checkIfRoot()

  confirm = typer.confirm("Are you sure you want to update using Cast?")
  if confirm:
    foundPM = False
    supportedPMs = ["pacman", "apt", "apk", "dnf", "eopkg"]

    for pm in supportedPMs:
      if isCmdInstalled(pm):
        foundPM = True
        match pm:
          case "pacman":
            updatePkg(f"pacman -Syyu --noconfirm")
          case "apt":
            updatePkg(f"apt update -y && apt upgrade -y")
          case "apk":
            updatePkg(f"apk update")
          case "dnf":
            updatePkg(f"dnf upgrade -y")
          case "eopkg":
            updatePkg(f"eopkg upgrade -y")
          case _:
            error("This should never happen")
    
    if not foundPM:
      notice("Could not find supported package manager")

def updatePkg(cmdToRun: str):
  try:
    subprocess.run(cmdToRun, shell=True, check=True)
  except:
    error("Failed to update system")
    raise typer.Exit(code=1)
  else:
    success("System updated")