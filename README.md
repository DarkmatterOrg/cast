# Cast

[![Version](https://img.shields.io/badge/Version-3.0.0-purple)](https://github.com/DarkmatterOrg/cast/releases/latest)
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

A not so small utility for Linux.

**Work In Progress**

## Installation

### Arch

Not in any repo's atm.

```bash

```

### Build

- pyinstaller is needed

Clone the project

```bash
  git clone https://github.com/DarkmatterOrg/cast
```

Go to the project directory

```bash
  cd cast
```

Create virtual environment

```bash
  uv venv
```

Enter the virtual environment

```bash
  . .venv/bin/activate
```

Install dependencies

```bash
  uv pip install -r requirements.txt
```

Because it'll install the wrong tomli, you're gonna have to change that

```bash
  uv pip uninstall tomli
  uv pip install --no-binary tomli tomli
```

Build

```bash
  ./build-cast.sh
```

## License

[MIT](/LICENSE)
