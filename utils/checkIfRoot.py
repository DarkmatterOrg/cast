import typer
import os

from utils.logger import warn

def checkIfRoot():
  if os.getuid() != 0:
    warn("You need to run this command with sudo")
    raise typer.Exit(code=1)