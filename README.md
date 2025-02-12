# Simple Shell in Go

A basic shell-like program in Go that lets you execute commands, change directories (`cd`), view history, and exit. It's simple, functional, and a great way to practice working with Go.

## Features:
- **Prompt**: Shows username and current directory (e.g., `username:/path/to/dir $`).
- **Commands**: Supports `cd`, `exit`, and `history`.
- **Handles Spaces**: `cd "folder name with spaces"` works.
- **History**: Type `history` to see your previous commands.

## Usage:
1. Clone or copy the code.
2. Run with:
   ```bash
   go run main.go
   ```
3. Type commands like:
   ```bash
   cd /path/to/directory
   history
   exit
   ```
##Example
```bash
username:/path/to/dir $ cd "VS Code"
username:/path/to/VS Code $ history
1 cd "VS Code"
2 history
username:/path/to/VS Code $ exit
```

feel free to improve it!
