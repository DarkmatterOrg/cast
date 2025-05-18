import typer
import subprocess
import os
import stat
import time
import shutil

from rich.console import Console
from utils.logger import info, success

app = typer.Typer()
console = Console()

@app.command(rich_help_panel="Fixes", short_help="Patch GMod")
def gmod():
  """
  Patch GMod's 64-bit beta to work properly on Linux (https://github.com/solsticegamestudios/GModCEFCodecFix)
  """
  info("Creating /tmp/patch-gmod")
  if not os.path.exists("/tmp/patch-gmod"):
    os.mkdir("/tmp/patch-gmod")
    info("Created /tmp/patch-gmod")

  with console.status("Downloading patch"):
    subprocess.getoutput("wget $(curl -s https://api.github.com/repos/solsticegamestudios/GModCEFCodecFix/releases/latest | jq -r '.assets[] | select(.name | test(\"GModCEFCodecFix-Linux\")) | .browser_download_url') -P /tmp/patch-gmod")
  info("Patch downloaded")

  info("Giving the patch run permissions")
  os.chmod("/tmp/patch-gmod/GModCEFCodecFix-Linux", stat.S_IRWXU)

  info("Patching gmod")
  os.system("/tmp/patch-gmod/GModCEFCodecFix-Linux")

  time.sleep(4)

  info("Removing the patch")
  shutil.rmtree("/tmp/patch-gmod")

  success("Fixed Gmod")
    