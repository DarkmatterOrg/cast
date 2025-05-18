import typer

from cmd.packageManager.install import app as install
from cmd.packageManager.remove import app as remove
from cmd.packageManager.search import app as search
from cmd.packageManager.update import app as update

app = typer.Typer()

app.add_typer(install)
app.add_typer(remove)
app.add_typer(search)
app.add_typer(update)