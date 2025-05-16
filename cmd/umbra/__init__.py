import typer

from utils.config import load_config
from cmd.umbra.gglobal import app as gglobal
from cmd.umbra.gpush import app as gpush

config = load_config()
app = typer.Typer()

if config["modules"]["umbra"]["git"]:
  app.add_typer(gglobal)
  app.add_typer(gpush)