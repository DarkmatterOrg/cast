import tomli
import typer
import os
import pwd

from pathlib import Path
from typing import Any, Dict

DEFAULT_CONFIG = {
  "insult": False,
  "modules": {
    "experimental": False,
    "fixes": False,
    "packageManager": False,
    "horizon": {
      "enabled": False,
      "distrobox": True
    },
    "umbra": {
      "enabled": False,
      "bfc": True,
      "git": True
    }
  }
}

def getUserHome() -> Path:
  if sudoUser := os.getenv("SUDO_USER"):
    return Path(f"/home/{sudoUser}")
  else:
    return Path.home()

CONFIG_PATH = Path(getUserHome() / ".config/cast/config.toml")

def deepMerge(default: Dict[str, Any], user: Dict[str, Any]) -> Dict[str, Any]:
  """Deeply merge two dictionaries, user settings override defaults."""
  merged = default.copy()
  for key, value in user.items():
    if key in merged:
      if isinstance(merged[key], dict) and isinstance(value, dict):
        merged[key] = deepMerge(merged[key], value)
      else:
        merged[key] = value
    else:
      merged[key] = value
  return merged

def loadConfig() -> Dict[str, Any]:
  """Load and merge configuration from file or defaults."""
  config = DEFAULT_CONFIG.copy()
  if CONFIG_PATH.exists():
    try:
      with open(CONFIG_PATH, "rb") as f:
        user_config = tomli.load(f)
      config = deepMerge(config, user_config)
    except tomli.TOMLDecodeError as e:
      typer.echo(f"Error parsing config file: {e}", err=True)
      raise typer.Exit(code=1)
  return config