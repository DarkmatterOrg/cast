import typer

from utils.logger import info, notice, error, success, warn
from rich import print
from rich.panel import Panel
app = typer.Typer()

@app.command(rich_help_panel="Experimentals")
def test():
  """
  This is just here so I can test some code before adding it to other commands, just to see how it works or looks.

  Will most likely do nothing at all
  """
  print(Panel("Hello, [red]World[/red]!"))