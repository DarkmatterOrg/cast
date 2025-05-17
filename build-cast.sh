

pyinstaller --onefile \
  --hidden-import typer \
  --hidden-import tomli \
  --collect-submodules typer \
  --collect-submodules tomli \
  --name cast \
  --path .venv/lib/python3.13/site-packages \
  cast.py

# pip install --no-binary tomli tomli