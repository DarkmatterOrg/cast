import typer
import tomli

app = typer.Typer()

@app.command(rich_help_panel="Base")
def version():
    """
    Print the version number of cast
    """
    with open("pyproject.toml", "rb") as f:
        pyproject = tomli.load(f)
        typer.echo(f"{pyproject["project"]["version"]}")