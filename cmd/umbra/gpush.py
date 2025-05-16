import typer
import shutil
import os

from rich.prompt import Prompt
from utils.logger import warn, info, success

app = typer.Typer()

@app.command(rich_help_panel="Umbra")
def gpush():
  """
  Adds everything, commits and pushes
  """

  if shutil.which("git") is None:
    warn("Git is not installed")
    raise typer.Exit()

  info("Adding everything...")
  os.system("git add .")

  commit = Prompt.ask("[bold magenta]Commit message[/bold magenta]")

  if not commit:
    warn("Commit message cannot be empty")
    raise typer.Exit()

  info("Committing...")
  os.system(f"git commit -m \"{commit}\"")
  
  info("Pushing...")
  try:
    os.system("git push")
  except:
    warn("Push failed")
    raise typer.Exit(code=1)
  
  success("Pushed")