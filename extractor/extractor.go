package extractor

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type ExtractResult struct {
	Lines      []string
	StartLine  int
	EndLine    int
	TotalLines int
}

func countLines(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}

func ExtractCode(filename string, start int, end int) (*ExtractResult, error) {
	if start <= 0 {
		return nil, errors.New("Starting Line must be greater than or equal to 1")
	}

	if end <= 0 {
		return nil, errors.New("Ending Line must be greater than one.")
	}

	if start > end {
		return nil, errors.New("Start Line cannot be greater than EndLine")
	}

	path, err := ResolveFilePath(filename)
	if err != nil {
		return nil, err
	}

	totalLines, err := countLines(path)
	if err != nil {
		return nil, err
	}

	if start > totalLines {
		return nil, fmt.Errorf(
			"Start Line (%d) exceeds total lines in file (%d)",
			start, totalLines,
		)
	}

	// Clamp the end line
	if end > totalLines {
		fmt.Printf("Warning: Endline (%d) exceeds file lenght (%d). Clamping to %d. \n",
			end, totalLines, totalLines,
		)
		end = totalLines
	}

	// Extract the Lines
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	current := 0
	lines := []string{}

	for scanner.Scan() {
		current++

		if current < start {
			continue
		}

		if current > end {
			break
		}

		lines = append(lines, scanner.Text()) // Adding the lines to the lines array
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(lines) == 0 {
		return nil, errors.New("No code extracted")
	}

	return &ExtractResult{
		Lines:      lines,
		StartLine:  start,
		EndLine:    end,
		TotalLines: totalLines,
	}, nil
}
