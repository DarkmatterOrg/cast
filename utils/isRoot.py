import os

def isRoot():
  return os.getuid() == 0