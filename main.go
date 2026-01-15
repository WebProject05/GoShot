package main

import (
	"fmt"
	"goshot/extractor"
	"goshot/renderer"
	"os"
)

func main() {
	filename := "extractor.go"
	startLine := 1
	endLine := 40

	result, err := extractor.ExtractCode(filename, startLine, endLine)
	if err != nil {
		panic(err)
	}

	fmt.Println("Total lines in file:", result.TotalLines)
	fmt.Printf("Extracted lines (%d → %d):\n\n", result.StartLine, result.EndLine)

	for i, line := range result.Lines {
		fmt.Printf("%4d | %s\n", result.StartLine+i, line)
	}

	fmt.Println("\nExtractor working perfectly.")

	path, err := extractor.ResolveFilePath(filename)
	if err != nil {
		panic(err)
	}

	lang, err := extractor.DetectLanguage(path)
	if err != nil {
		panic(err)
	}

	fmt.Println("Detected language:", lang)


	html, err := renderer.HighlightToHTML(
		result.Lines,
		lang,
		"dracula",
	)

	if err != nil {
		panic(err)
	}

	os.WriteFile("output.html", []byte(html), 0644)
	fmt.Println("Html code has been generated.")

}
