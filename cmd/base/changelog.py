import typer
import requests

from rich.console import Console
from castValues import VERSION
from rich import print
from rich.panel import Panel
from rich.markdown import Markdown
from utils.logger import warn

console = Console()
app = typer.Typer()

@app.command(rich_help_panel="Base")
def changelog():
  """
  Show the changelog for the current version
  """
  if VERSION < "3.0.0":
    warn("Changelog is only available for versions >= 3.0.0")
    raise typer.Exit(code=1)

  with console.status("Getting changelog..."):
    resJson = requests.get("https://api.github.com/repos/DarkmatterOrg/cast/releases").json()
  
  for res in resJson:
    if res["tag_name"] == VERSION and not res["prerelease"] and not res["draft"]:
      changelogBody = res["body"]
      if not changelogBody:
        print(Panel("No changelog available for this version", title=f"[bold purple]{VERSION}[/bold purple]", border_style="magenta"))
      else:
        print(Panel.fit(Markdown(changelogBody), title=f"[bold purple]{VERSION}[/bold purple]", border_style="magenta"))