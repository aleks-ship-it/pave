# Pave CLI - Usage Guide

Pave is a CLI tool that helps you manage symbolic links and generate installation scripts for your custom tools.

## Table of Contents
- [Installation](#installation)
- [Global Flags](#global-flags)
- [Commands](#commands)
  - [pave link](#pave-link)
  - [pave unlink](#pave-unlink)
  - [pave list](#pave-list)
  - [pave status](#pave-status)
  - [pave generate](#pave-generate)
- [Common Workflows](#common-workflows)

---

## Installation

### Add Pave to PATH (Windows)

Run once to enable `pave` from any directory:

```cmd
pave link --name pave --path "C:\Users\aleks\Desktop\pave\pave.exe"
```

Then restart your terminal. The following directory is automatically added to your PATH:
```
C:\Users\aleks\AppData\Roaming\pave\bin
```

---

## Global Flags

These flags work with all commands:

| Flag | Shorthand | Description |
|------|-----------|-------------|
| `--verbose` | - | Enable verbose output with detailed information |
| `--dry-run` | - | Preview changes without applying them |
| `--version` | `-V` | Print version information |
| `--help` | `-h` | Show help for any command |

**Example:**
```cmd
pave --version
pave --help
pave link --help
```

---

## Commands

### `pave link`

Create a symbolic link or wrapper script to make a CLI tool accessible from anywhere.

**Syntax:**
```cmd
pave link --name <name> --path <path> [--verbose] [--dry-run]
```

**Flags:**
| Flag | Required | Description |
|------|----------|-------------|
| `--name` | Yes | Name of the link (how you'll call it) |
| `--path` | Yes | Full path to the executable |

**Examples:**

```cmd
# Link a simple CLI tool
pave link --name my-cli --path "C:\Users\aleks\Desktop\my-cli.exe"

# Link with verbose output
pave link --name my-tool --path "C:\Tools\my-tool.exe" --verbose

# Preview what would happen (dry-run)
pave link --name test --path "C:\path\to\test.exe" --dry-run

# Link a batch script
pave link --name deploy --path "C:\Scripts\deploy.bat"
```

**After linking, you can run your tool from anywhere:**
```cmd
my-cli
my-tool
deploy
```

---

### `pave unlink`

Remove a previously created symbolic link or wrapper script.

**Syntax:**
```cmd
pave unlink --name <name> [--verbose] [--dry-run]
```

**Flags:**
| Flag | Required | Description |
|------|----------|-------------|
| `--name` | Yes | Name of the link to remove |

**Examples:**

```cmd
# Remove a linked tool
pave unlink --name my-cli

# Preview removal (dry-run)
pave unlink --name my-cli --dry-run

# Remove with verbose output
pave unlink --name my-tool --verbose
```

---

### `pave list`

List all managed symbolic links with their status.

**Syntax:**
```cmd
pave list [--verbose]
```

**Examples:**

```cmd
# List all linked tools
pave list

# List with verbose output
pave list --verbose
```

**Output example:**
```
Managed links:
  my-cli -> C:\Users\aleks\Desktop\my-cli.exe [valid]
  deploy -> C:\Scripts\deploy.bat [valid]
```

---

### `pave status`

Show the status of a specific link or all links.

**Syntax:**
```cmd
pave status [--name <name>] [--verbose]
```

**Flags:**
| Flag | Required | Description |
|------|----------|-------------|
| `--name` | No | Name of the link to check (omit to show all) |

**Examples:**

```cmd
# Check status of a specific tool
pave status --name my-cli

# Check all links (same as pave list)
pave status

# Verbose status check
pave status --name my-cli --verbose
```

**Output example:**
```
Name: my-cli
Path: C:\Users\aleks\Desktop\my-cli.exe
Target: C:\Users\aleks\AppData\Roaming\pave\bin\my-cli
Status: valid
```

---

### `pave generate`

Generate installation scripts (shell and PowerShell) for distributing your CLI tool.

**Syntax:**
```cmd
pave generate --name <name> --repo <repo> --bin <bin> [--out <dir>] [--verbose] [--dry-run]
```

**Flags:**
| Flag | Required | Description |
|------|----------|-------------|
| `--name` | Yes | Name of the tool |
| `--repo` | Yes | GitHub repository (format: `owner/repo`) |
| `--bin` | Yes | Binary name (without extension) |
| `--out` | No | Output directory (default: current directory) |

**Examples:**

```cmd
# Generate install scripts for a tool
pave generate --name mycli --repo aleks/mycli --bin mycli

# Generate to a specific directory
pave generate --name mycli --repo aleks/mycli --bin mycli --out "C:\Dist"

# Preview the scripts without writing files (dry-run)
pave generate --name mycli --repo aleks/mycli --bin mycli --dry-run

# Generate with verbose output
pave generate --name mycli --repo aleks/mycli --bin mycli --verbose
```

**Generated files:**
- `install.sh` - Shell script for Linux/macOS installation
- `install.ps1` - PowerShell script for Windows installation

---

## Common Workflows

### Workflow 1: Make a Custom CLI Globally Accessible

```cmd
# 1. Create your CLI tool
echo @echo off > C:\Tools\my-awesome-cli.cmd
echo echo Hello from my-awesome-cli >> C:\Tools\my-awesome-cli.cmd

# 2. Link it with pave
pave link --name mycli --path "C:\Tools\my-awesome-cli.cmd"

# 3. Run it from anywhere
mycli
```

### Workflow 2: Check and Manage Your Linked Tools

```cmd
# List all linked tools
pave list

# Check if a specific tool is working
pave status --name mycli

# Remove a tool you no longer need
pave unlink --name mycli
```

### Workflow 3: Generate Installation Scripts for Distribution

```cmd
# Generate scripts for your GitHub release
pave generate \
  --name mytool \
  --repo aleks/mytool \
  --bin mytool \
  --out ./dist

# This creates:
# ./dist/install.sh
# ./dist/install.ps1
```

### Workflow 4: Test Before Applying (Dry Run)

```cmd
# See what would happen without making changes
pave link --name test --path "C:\path\to\test.exe" --dry-run

# Check what generate would produce
pave generate --name test --repo user/test --bin test --dry-run
```

---

## Tips

1. **Use `--verbose`** when troubleshooting to see detailed output
2. **Use `--dry-run`** to preview changes before applying them
3. **Windows note**: On Windows, pave creates `.cmd` wrapper scripts instead of symlinks to avoid requiring administrator privileges
4. **PATH requirement**: The directory `C:\Users\aleks\AppData\Roaming\pave\bin` must be in your PATH for linked tools to work
5. **Registry location**: Link information is stored in `%LOCALAPPDATA%\pave\links.json`

---

## Troubleshooting

### "command not found" when running a linked tool
**Solution:** Restart your terminal after linking, or run:
```cmd
set PATH=%PATH%;C:\Users\aleks\AppData\Roaming\pave\bin
```

### "target path does not exist" error
**Solution:** Use the full Windows path (with drive letter), not WSL/Linux paths:
```cmd
# Wrong
pave link --name test --path /mnt/c/path/to/test.exe

# Right
pave link --name test --path "C:\path\to\test.exe"
```

### Link shows as "broken" in status
**Solution:** The original executable was moved or deleted. Either update the link or remove it:
```cmd
pave unlink --name broken-tool
```
