package main

import (
	"fmt"
	"log"

	"goshot/extractor"
	"goshot/renderer"
)

func main() {
	filename := "extractor.go"
	startLine := 1
	endLine := 25
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

	html, err := renderer.HighlightToHTML(result.Lines, lang, theme)
	if err != nil {
		log.Fatal(err)
	}

	err = renderer.RenderToImage(html, filename, outputImage)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("GoShot screenshot generated: %s (%s lines %d→%d)\n",
		outputImage, lang, result.StartLine, result.EndLine)
}
