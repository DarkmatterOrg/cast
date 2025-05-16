import typer
import os

from utils.logger import info

app = typer.Typer()

@app.command(rich_help_panel="Umbra")
def gglobal(name: str, email: str):
  """
  Git global config
  """
  info("Setting name...")
  os.system(f"git config --global user.name \"{name}\"")
  info("Setting email...")
  os.system(f"git config --global user.email \"{email}\"")