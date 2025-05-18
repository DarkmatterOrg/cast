import typer

from cmd.experimental.remove import app as remove
from cmd.experimental.test import app as test
from cmd.experimental.update import app as update

app = typer.Typer()

app.add_typer(remove)
app.add_typer(test)
app.add_typer(update)