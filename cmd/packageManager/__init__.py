import typer

from cmd.packageManager.search import app as search

app = typer.Typer()

app.add_typer(search)