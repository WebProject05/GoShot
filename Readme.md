# GoShot

A command-line tool that converts source code into high-quality, shareable screenshots directly from the terminal.

Built for developers who want beautiful code images without opening browsers or GUI tools.

##  Features

- **Syntax Highlighting** — Powered by Chroma with multiple theme options
- **macOS-style Window Frame** — Professional appearance for your code
- **Language Support** — All programming languages & .txt files  
- **Code Splitting** — Automatically split large files into multiple images
- **Language Detection** — Automatic syntax detection based on file type
- **Headless Rendering** — No visible browser window required
- **High-Resolution Output** — Generate sharp PNG images for any resolution

##  Quick Start

### Basic Usage

Generate a screenshot of an entire file:

`bash
goshot <file>
`

### Capture a Specific Range

`bash
goshot <file> 10 120
`

### Specify Output Filename

`bash
goshot <file> -o output.png
`

### Split Large Code Blocks

`bash
goshot <file> 1 400 --split
`

Output:
- output_1.png
- output_2.png
- output_3.png

Each image contains a portion of the code with a clean header.

### Choose a Syntax Theme

`bash
goshot <file> --theme monokai
`

##  Command Reference

### Syntax

`
goshot <file> [start] [end] [flags]
`

### Positional Arguments

| Argument | Description |
|----------|-------------|
| <file> | Source code file |
| [start] | Start line (optional) |
| [end] | End line (optional) |

### Supported Flags

| Flag | Description |
|------|-------------|
| -o, --output | Output image name |
| --theme | Syntax highlighting theme |
| --split | Split output into multiple images |
| --help | Show help text |

*Note: Other flags are internally supported but not yet exposed or finalized.*

##  Themes

GoShot supports all themes provided by [Chroma](https://github.com/alecthomas/chroma).

### Popular Themes

- dracula (default)
- monokai
- github-dark
- github-dark-dimmed
- 	okyonight
- 
ord
- one-dark

### Example

`bash
goshot main.go --theme dracula
`

##  How It Works

GoShot follows this pipeline:

1. Resolves the input file path
2. Extracts the requested lines
3. Detects the programming language
4. Applies syntax highlighting
5. Wraps the code in a macOS-style window
6. Renders the output using a headless browser
7. Saves the result as a PNG image

Everything runs locally.

##  Supported File Types

- Programming language source files (via Chroma)
- .txt files

Binary or unsupported file types are rejected.

##  Image Sharing Tips

Some platforms (like WhatsApp) compress images aggressively.

For best quality:
- Send generated images as **Documents**, not Images

##  Output Behavior

### Single Image

`
file.png
`

### Split Output

`
file_1.png
file_2.png
file_3.png
`

Each image includes:
- Filename in the title bar
- Part indicator when split is enabled

##  Project Status

| Component | Status |
|-----------|--------|
| Core rendering pipeline |  Complete |
| CLI commands |  Complete |
| Split logic |  Complete |
| Loader/progress messages |  Complete |
| Documentation |  In Progress |
| Distribution |  Not Yet Published |
