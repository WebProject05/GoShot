package main

import (
	"fmt"

	"goshot/extractor"
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
}
