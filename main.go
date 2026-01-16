package main

import (
	"fmt"
	"log"
	"path/filepath"

	"goshot/cmd"
	"goshot/extractor"
	"goshot/internal/split"
	"goshot/renderer"
)

const MaxLinesPerImage = 80 // default split size (used only when --split is enabled)

// normalizeOutputName ensures a valid PNG output name
func normalizeOutputName(output string) (base string, ext string) {
	ext = filepath.Ext(output)

	// Default to .png if no extension
	if ext == "" {
		return output, ".png"
	}

	// Enforce PNG only (v1)
	if ext != ".png" {
		log.Fatalf("Unsupported output format: %s (only .png supported)", ext)
	}

	base = output[:len(output)-len(ext)]
	return
}

func main() {
	cfg, err := cmd.ParseArgs()
	if err != nil {
		cmd.ExitWithError(err)
	}

	result, err := extractor.ExtractCode(cfg.File, cfg.StartLine, cfg.EndLine)
	if err != nil {
		log.Fatal(err)
	}

	path, err := extractor.ResolveFilePath(cfg.File)
	if err != nil {
		log.Fatal(err)
	}

	lang, err := extractor.DetectLanguage(path)
	if err != nil {
		log.Fatal(err)
	}

	var chunks [][]string
	if cfg.Split {
		chunks = split.ChunkLines(result.Lines, MaxLinesPerImage)
	} else {
		chunks = [][]string{result.Lines}
	}

	totalParts := len(chunks)
	baseName := filepath.Base(cfg.File)

	baseOutput, ext := normalizeOutputName(cfg.Output)

	for i, chunk := range chunks {
		part := i + 1

		// Output filename
		output := baseOutput + ext
		if totalParts > 1 {
			output = fmt.Sprintf("%s_%d%s", baseOutput, part, ext)
		}

		// Header title
		header := baseName
		if totalParts > 1 {
			header = fmt.Sprintf("%s • Part %d / %d", baseName, part, totalParts)
		}

		fmt.Printf("[%d/%d] Generating image...\n", part, totalParts)

		html, err := renderer.HighlightToHTML(chunk, lang, cfg.Theme)
		if err != nil {
			log.Fatal(err)
		}

		err = renderer.RenderToImage(html, header, output)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf(
		"✔ Done. Generated %d image(s) (%s lines %d→%d)\n",
		totalParts,
		lang,
		result.StartLine,
		result.EndLine,
	)
}
