import typer

from cmd.fixes.gmod import app as gmod
from cmd.fixes.protonHang import app as proton_hang

app = typer.Typer()

app.add_typer(gmod)
app.add_typer(proton_hang)