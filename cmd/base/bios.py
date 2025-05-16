import typer
import os

from utils.logger import notice

app = typer.Typer()

@app.command(rich_help_panel="Base")
def bios():
  """
  Boot into this device's BIOS/UEFI screen
  """
  if not os.path.exists("/sys/firmware/efi"):
    notice("Rebooting to legacy BIOS from OS is not supported")
    raise typer.Exit()
  
  confirm = typer.confirm("Are you sure you want to reboot to BIOS?")
  if confirm:
    os.system("systemctl reboot --firmware-setup")