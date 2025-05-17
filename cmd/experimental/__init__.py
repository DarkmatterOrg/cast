import typer

from cmd.experimental.install import app as install
from cmd.experimental.remove import app as remove
from cmd.experimental.search import app as search
from cmd.experimental.test import app as test
from cmd.experimental.update import app as update

app = typer.Typer()

app.add_typer(install)
app.add_typer(remove)
app.add_typer(search)
app.add_typer(test)
app.add_typer(update)