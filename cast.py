import os

import typer

from utils.config import loadConfig
from utils.logger import importantWarn, debug

from cmd.base import app as base
from cmd.experimental import app as experimental
from cmd.fixes import app as fix
from cmd.horizon import app as horizon
from cmd.packageManager import app as packageManager
from cmd.umbra import app as umbra
from cmd.umbra.bfc import app as bfc

config = loadConfig()
app = typer.Typer(add_completion=False, rich_markup_mode="rich", no_args_is_help=True)

app.add_typer(base)

if config["modules"]["experimental"] or os.getenv("CastIsDebug"):
  app.add_typer(experimental)

if config["modules"]["fixes"] or os.getenv("CastIsDebug"):
  app.add_typer(fix, name="fix", rich_help_panel="Base", help="Different fixes for various things")

if config["modules"]["horizon"]["enabled"] or os.getenv("CastIsDebug"):
  app.add_typer(horizon)

if config["modules"]["packageManager"] or os.getenv("CastIsDebug"):
  app.add_typer(packageManager)

if config["modules"]["umbra"]["enabled"] or os.getenv("CastIsDebug"):
  app.add_typer(umbra)

  if config["modules"]["umbra"]["bfc"] or os.getenv("CastIsDebug"):
    app.add_typer(bfc, name="bfc", rich_help_panel="Umbra", help="Commands to use from a stream deck with Bitfocus Companion")

if __name__ == "__main__":
  if config["modules"]["experimental"]:
    if config["insult"]:
      importantWarn("Either you're fucking stupid or just curious. Either way; experimental features are enabled, use at your own risk")
    else:
      importantWarn("Experimental features are enabled, use at your own risk")

  if os.getenv("CastIsDebug"):
    debug("Cast is currently in Debug mode!")
      
  app()