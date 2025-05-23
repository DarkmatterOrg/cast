import typer
import os
import shutil

from typing_extensions import Annotated
from utils.logger import warn
from utils.config import load_config

app = typer.Typer()

@app.command(rich_help_panel="Horizon")
def dbox_ubuntu(pkgs: Annotated[str, typer.Option("--pkgs", help="Comma separated list of packages to install")] = None):
  """
  Create a Ubuntu distrobox container
  """

  if shutil.which("distrobox") is None:
    config = load_config()

    if config["insult"]:
      warn("How the fuck do you expect me to do this when you don't even have distrobox installed")
    else:
      warn("Distrobox is not installed")
      
    raise typer.Exit()

  if pkgs:
    pkgs = pkgs.replace(",", " ")

  if pkgs:
    os.system(f"distrobox-create --image quay.io/toolbx/ubuntu-toolbox:latest --name ubuntu --additional-packages \"{pkgs}\"")
  else:
    os.system("distrobox-create --image quay.io/toolbx/ubuntu-toolbox:latest --name ubuntu")