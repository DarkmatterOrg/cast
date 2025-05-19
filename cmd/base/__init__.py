import typer

from cmd.base.age import app as age
from cmd.base.authors import app as authors
from cmd.base.bios import app as bios
from cmd.base.changelog import app as changelog
from cmd.base.config import app as config
from cmd.base.passwordFeedback import app as passwordFeedback
from cmd.base.version import app as version

app = typer.Typer()

app.add_typer(age)
app.add_typer(authors)
app.add_typer(bios)
app.add_typer(changelog)
app.add_typer(config)
app.add_typer(passwordFeedback)
app.add_typer(version)