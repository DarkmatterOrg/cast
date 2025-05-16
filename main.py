import typer

from utils.config import load_config
from utils.logger import importantWarn

from cmd.base import app as base
from cmd.experimental import app as experimental
from cmd.fixes import app as fix
from cmd.horizon import app as horizon
from cmd.umbra import app as umbra
from cmd.umbra.bfc import app as bfc

config = load_config()
app = typer.Typer()

app.add_typer(base)

if config["modules"]["experimental"]:
  app.add_typer(experimental)

if config["modules"]["fixes"]:
  app.add_typer(fix, name="fix", rich_help_panel="Base", help="Different fixes for various things")

if config["modules"]["horizon"]["enabled"]:
  app.add_typer(horizon)

if config["modules"]["umbra"]["enabled"]:
  app.add_typer(umbra)

  if config["modules"]["umbra"]["bfc"]:
    app.add_typer(bfc, name="bfc", rich_help_panel="Umbra", help="Commands to use from a stream deck with Bitfocus Companion")

if __name__ == "__main__":
  if config["modules"]["experimental"]:
    if config["insult"]:
      importantWarn("Either you're fucking stupid or just curious. Either way; experimental features are enabled, use at your own risk")
    else:
      importantWarn("Experimental features are enabled, use at your own risk")
      
  app()