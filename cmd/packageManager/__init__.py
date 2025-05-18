import typer

from cmd.packageManager.install import app as install
from cmd.packageManager.search import app as search

app = typer.Typer()

app.add_typer(install)
app.add_typer(search)