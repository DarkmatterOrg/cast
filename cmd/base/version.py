import typer

from castValues import VERSION

app = typer.Typer()

@app.command(rich_help_panel="Base")
def version():
    """
    Print the version number of cast
    """
    typer.echo(VERSION)