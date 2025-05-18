import shutil

def isCmdInstalled(cmd: str) -> bool:
  return shutil.which(cmd) is not None