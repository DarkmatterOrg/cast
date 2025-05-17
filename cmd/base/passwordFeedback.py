import typer
import os

from typing import Annotated
from utils.logger import warn, success, notice
from utils.checkIfRoot import checkIfRoot


app = typer.Typer()

@app.command(rich_help_panel="Base")
def pwd_fdbk(toggle: Annotated[bool, typer.Option("--enable/--disable", help="Enable or disable password feedback", show_default=False)]):
    """
    Toggles password prompt feedback in terminal, where sudo password prompts will display asterisks when enabled
    """
    checkIfRoot()
    
    pwfeedbackPath = "/etc/sudoers.d/pwfeedback"

    if toggle:
      with open(pwfeedbackPath, "w") as f:
        f.write("Defaults pwfeedback")
        f.close()
      success("Password feedback enabled")
      notice("Please restart your terminal for the changes to take effect")
    else:
      os.remove(pwfeedbackPath)
      success("Password feedback disabled")
      notice("Please restart your terminal for the changes to take effect")