import typer

from rich.console import Console
from utils.config import loadConfig,CONFIG_PATH
from utils.logger import info

app = typer.Typer()
console = Console()

@app.command(rich_help_panel="Base")
def config():
  """
  Show the current configuration and where it's located
  """
  config = loadConfig()
  console.print(config)

  if CONFIG_PATH.exists():
    info(f"Config file found at {CONFIG_PATH}")
  else:
    info(f"Config should be created at {CONFIG_PATH}")