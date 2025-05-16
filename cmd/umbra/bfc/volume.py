import typer
import os

from typing import Annotated

app = typer.Typer()

@app.command(rich_help_panel="Umbra")
def volume(toggle: Annotated[bool, typer.Option("--increase/--decrease", help="Increase or decrease volume", show_default=False)], amount: Annotated[int, typer.Option("--amount", help="Amount to increase or decrease volume by")] = 10):
  """
  Increase or decrease volume on default device
  """

  if toggle:
    os.system(f"pactl set-sink-volume @DEFAULT_SINK@ +{amount}%")
  else:
    os.system(f"pactl set-sink-volume @DEFAULT_SINK@ -{amount}%")