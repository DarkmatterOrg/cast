import typer

from cmd.umbra.bfc.mute import app as mute
from cmd.umbra.bfc.unmute import app as unmute
from cmd.umbra.bfc.volume import app as volume

app = typer.Typer()

app.add_typer(mute)
app.add_typer(unmute)
app.add_typer(volume)