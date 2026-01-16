GoShot

GoShot is a command-line tool that converts source code into high-quality, shareable screenshots directly from the terminal.

It’s built for developers who want beautiful code images without opening browsers or GUI tools.

What GoShot Does (Current Capabilities)

Syntax highlighting using Chroma

macOS-style window frame

upports programming languages & .txt files

Optional splitting for large code blocks

Automatic language detection

Headless rendering (no visible browser)

Generates high-resolution PNG images

Usage
Basic usage
goshot <file>


Generates a screenshot of the entire file.

Capture a specific line range
goshot <file> 10 120

Specify output filename
goshot <file> -o output.png

Split large code into multiple images
goshot <file> 1 400 --split


Output:

output_1.png
output_2.png
output_3.png


Each image contains a portion of the code with a clean header.

Choose a syntax theme
goshot <file> --theme monokai

Command Syntax
goshot <file> [start] [end] [flags]

Positional arguments
Argument	Description
<file>	Source code file
[start]	Start line (optional)
[end]	End line (optional)
Supported Flags
Flag	Description
-o, --output	Output image name
--theme	Syntax highlighting theme
--split	Split output into multiple images
--help	Show help text

(Other flags are internally supported but not yet exposed or finalized.)

Themes

GoShot supports all themes provided by Chroma.

Commonly used themes

dracula (default)

monokai

github-dark

github-dark-dimmed

tokyonight

nord

one-dark

Example:

goshot main.go --theme dracula

How GoShot Works

Resolves the input file path

Extracts the requested lines

Detects the programming language

Applies syntax highlighting

Wraps the code in a macOS-style window

Renders the output using a headless browser

Saves the result as a PNG image

Everything runs locally.

Image Sharing Note

Some platforms (like WhatsApp) compress images aggressively.

For best quality:

Send generated images as Documents, not Images

Output Behavior
Single image
file.png

Split output
file_1.png
file_2.png
file_3.png


Each image includes:

Filename in the title bar

Part indicator when split is enabled

🛠 Supported Files

Programming language source files (via Chroma)

.txt files

Binary or unsupported file types are rejected.

Project Status

Current state:

Core rendering pipeline: complete

CLI commands: implemented

Split logic: implemented

Loader/progress messages: implemented

Documentation: in progress

Distribution: not yet published