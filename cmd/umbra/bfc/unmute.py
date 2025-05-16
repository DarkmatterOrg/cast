import typer
import os

app = typer.Typer()

@app.command(rich_help_panel="Umbra")
def unmute():
  """
  Unmutes the default device
  """
  os.system("pactl set-sink-mute @DEFAULT_SINK@ 0")

#TODO Add flags for unmuting speaker and mic