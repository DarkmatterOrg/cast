import typer

from utils.config import loadConfig

from cmd.horizon.dboxUbuntu import app as dbox_ubuntu

config = loadConfig()
app = typer.Typer()

if config["modules"]["horizon"]["distrobox"]:
  app.add_typer(dbox_ubuntu)