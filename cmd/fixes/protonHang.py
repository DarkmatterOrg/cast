import typer
import subprocess

from utils.logger import info

app = typer.Typer()

@app.command(rich_help_panel="Fixes", short_help="Kills all processes related to wine and proton.")
def proton_hang():
  """
  Kills all processes related to wine and proton. This forces it to restart next time you launch the game (you might still have to press STOP in steam to kill the game binary)
  """
  protonCore = [
    "pv-bwrap",
    "pressure-vessel",
    "reaper",
    "explorer.exe",
    "rpcss.exe",
    "plugplay.exe",
    "services.exe",
    "svchost.exe",
    "winedevice.exe",
    "winedevice.exe",
    "wineserver"
  ]

  for process in protonCore:
    try:
      result = subprocess.getoutput(f"killall -9 {process}")
      info(f"{result}")
    except:
      pass