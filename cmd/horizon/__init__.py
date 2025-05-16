import typer

from utils.config import load_config

from cmd.horizon.dboxUbuntu import app as dbox_ubuntu

config = load_config()
app = typer.Typer()

if config["modules"]["horizon"]["distrobox"]:
  app.add_typer(dbox_ubuntu)