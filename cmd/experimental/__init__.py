import typer

from cmd.experimental.test import app as test

app = typer.Typer()

app.add_typer(test)