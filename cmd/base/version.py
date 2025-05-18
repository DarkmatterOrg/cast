import typer
import requests

from rich.console import Console
from utils.logger import info, notice
from castValues import VERSION

console = Console()
app = typer.Typer()

@app.command(rich_help_panel="Base")
def version():
  """
  Print the version number of cast
  """
  info(f"[italic cyan]{VERSION}[/italic cyan]")

  with console.status("Checking for updates..."):
    resJson = requests.get("https://api.github.com/repos/DarkmatterOrg/cast/releases").json()
  
  for res in resJson:
    if res["tag_name"] > VERSION and not res["prerelease"] and not res["draft"]:
      notice(f"[green]New version available[/green]: {res['tag_name']}")