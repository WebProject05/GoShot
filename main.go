package main

import (
	"fmt"
	"log"
	"path/filepath"

	"goshot/extractor"
	"goshot/internal/split"
	"goshot/renderer"
)

const MaxLinesPerImage = 80   // Can be modified using the commands (later)

func main() {
	filename := `file.cpp`
	startLine := 1
	endLine := 300
	theme := "dracula"
	outputImage := "output.png"

	result, err := extractor.ExtractCode(filename, startLine, endLine)
	if err != nil {
		log.Fatal(err)
	}

	path, err := extractor.ResolveFilePath(filename)
	if err != nil {
		log.Fatal(err)
	}

	lang, err := extractor.DetectLanguage(path)
	if err != nil {
		log.Fatal(err)
	}

	chunks := split.ChunkLines(result.Lines, MaxLinesPerImage)
	totalParts := len(chunks)

	baseName := filepath.Base(filename)

	for i, chunk := range chunks {
		part := i + 1

		output := outputImage
		if totalParts > 1 {
			ext := filepath.Ext(outputImage)
			name := outputImage[:len(outputImage)-len(ext)]
			output = fmt.Sprintf("%s_%d%s", name, part, ext)
		}

		header := baseName
		if totalParts > 1 {
			header = fmt.Sprintf("%s • Part %d / %d", baseName, part, totalParts)
		}

		fmt.Printf("[%d/%d] Generating image...\n", part, totalParts)

		html, err := renderer.HighlightToHTML(chunk, lang, theme)
		if err != nil {
			log.Fatal(err)
		}

		err = renderer.RenderToImage(html, header, output)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf(
		"Done. Generated %d image(s) (%s lines %d→%d)\n",
		totalParts,
		lang,
		result.StartLine,
		result.EndLine,
	)
}
