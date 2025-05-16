

pyinstaller --onefile \
  --hidden-import typer \
  --hidden-import tomli \
  --collect-submodules typer \
  --collect-submodules tomli \
  --name cast \
  --path .venv/lib/python3.13/site-packages \
  main.py

# pip install --no-binary tomli tomli