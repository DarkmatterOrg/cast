import typer
import os

app = typer.Typer()

@app.command(rich_help_panel="Umbra")
def mute():
  """
  Mutes the default device
  """
  os.system("pactl set-sink-mute @DEFAULT_SINK@ 1")

#TODO Add flags for muting speaker and mic