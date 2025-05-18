import typer

from castValues import AUTHORS
from rich import print
from rich.panel import Panel


app = typer.Typer()

@app.command(rich_help_panel="Base")
def authors():
  """
  The Authors of Cast
  """
  for author in AUTHORS:
    print(Panel.fit(f"[bold]Name[/bold]: {author['name']}\n[bold]Email[/bold]: {author['email']}\n[bold]Website[/bold]: [blue]{author['website']}[/blue]", border_style="magenta"))