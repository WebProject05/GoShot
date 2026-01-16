package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func ParseArgs() (*Config, error) {
	cfg := &Config{}

	// Flags
	flag.StringVar(&cfg.Output, "output", "goshot.png", "Output image name")
	flag.StringVar(&cfg.Output, "o", "goshot.png", "Output image name (shorthand)")
	flag.StringVar(&cfg.Theme, "theme", "dracula", "Syntax highlighting theme")
	flag.BoolVar(&cfg.Split, "split", false, "Split output into multiple images")
	flag.BoolVar(&cfg.NoFrame, "no-frame", false, "Disable macOS window frame")
	flag.Float64Var(&cfg.Scale, "scale", 3.0, "Screenshot scale (2 or 3 recommended)")
	flag.IntVar(&cfg.FontSize, "font-size", 16, "Font size for code")

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		return nil, errors.New("No file provided")
	}

	cfg.File = args[0]

	if len(args) >= 2 {
		start, err := strconv.Atoi(args[1])
		if err != nil || start < 1 {
			return nil, fmt.Errorf("Invalid start Line: %s", args[1])
		}
		cfg.StartLine = start
	} else {
		cfg.StartLine = 1
	}

	if len(args) >= 3 {
		end, err := strconv.Atoi(args[2])
		if err != nil || end < cfg.StartLine {
			return nil, fmt.Errorf("Invalid end Line: %s", args[2])
		}
		cfg.EndLine = end
	} else {
		cfg.EndLine = 0
	}

	return cfg, nil
}


func ExitWithError(err error) {
	fmt.Fprintln(os.Stderr, "Error:", err)
	fmt.Fprintln(os.Stderr, "Run `goshot --help` for usage")
	os.Exit(1)
}