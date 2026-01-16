package cmd

import "fmt"

func PrintUsage() {
	fmt.Println(`GoShot — Create beautiful code screenshots from the terminal

Usage:
  goshot <file> [start] [end] [flags]

Examples:
  goshot main.go
  goshot file.cpp 10 120
  goshot file.go 1 400 --split
  goshot file.go --theme monokai -o output.png

Flags:
  -o, --output     Output image name (default: goshot.png)
      --theme      Syntax highlighting theme (default: dracula)
      --split      Split into multiple images
      --no-frame   Disable macOS window frame
      --scale      Screenshot scale (default: 3)
      --font-size  Font size for code (default: 16)
      --help       Show this help message
`)
}
