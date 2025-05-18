import typer

from cmd.experimental.test import app as test
from cmd.experimental.update import app as update

app = typer.Typer()

app.add_typer(test)
app.add_typer(update)