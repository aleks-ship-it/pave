# pave — make any CLI tool runnable from anywhere in your terminal

Stop copying binaries to `/usr/local/bin` or messing with PATH. `pave` symlinks your CLIs into a single directory and manages them for you.

## Quick install

**Windows:**

```powershell
irm https://github.com/aleks-ship-it/pave/releases/latest/download/install.ps1 | iex
```

**Linux/macOS:**

```bash
curl -fsSL https://github.com/aleks-ship-it/pave/releases/latest/download/install.sh | bash
```

## Usage

### `pave link` — make a CLI globally accessible

```bash
pave link --name mycli --path /path/to/mycli
```

Flags: `--name` (required), `--path` (required), `--verbose`, `--dry-run`

### `pave unlink` — remove a linked CLI

```bash
pave unlink --name mycli
```

Flags: `--name` (required), `--verbose`, `--dry-run`

### `pave list` — show all linked CLIs

```bash
pave list
```

Flags: `--verbose`

### `pave status` — check a specific link

```bash
pave status --name mycli
```

Flags: `--name`, `--verbose`

### `pave generate` — create install scripts for your tool

```bash
pave generate --name mytool --repo aleks/mytool --bin mytool
```

Flags: `--name` (required), `--repo` (required), `--bin` (required), `--out`, `--verbose`, `--dry-run`

## How it works

**Linux:** Symlinks go to `~/.local/bin`. Add it to PATH: `export PATH=~/.local/bin:$PATH`

**macOS:** Symlinks go to `~/bin`. Add it to PATH: `export PATH=~/bin:$PATH`

**Windows:** Creates `.cmd` wrapper scripts in `%LOCALAPPDATA%\pave\bin` (no admin required). Add to PATH or `pave` warns you.

## Release your own tool

Generate cross-platform install scripts that download from GitHub releases:

```bash
pave generate \
  --name mycli \
  --repo username/mycli \
  --bin mycli \
  --out ./dist
```

This creates `install.sh` (Unix/macOS) and `install.ps1` (Windows) that users can run to install your tool.

## Built with

[Go](https://golang.org) · [Cobra](https://github.com/spf13/cobra)

## License

[![MIT License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
