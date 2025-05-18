import typer

from utils.config import loadConfig
from cmd.umbra.gglobal import app as gglobal
from cmd.umbra.gpush import app as gpush

config = loadConfig()
app = typer.Typer()

if config["modules"]["umbra"]["git"]:
  app.add_typer(gglobal)
  app.add_typer(gpush)