package main

import (
	"bufio"
	"fmt"
	"os"
)


// Count the total number of lines the File has
func CountLines(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("Error Opening file %w", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}

	return lines, scanner.Err()

}


// Extract the Lines of the file based on the request
func ExtractLines(filePath string, start int, end int) (string, error) {
	totalLines, err := CountLines(filePath)
	if err != nil {
		return "", fmt.Errorf("Error %w", err)
	}

	if start > totalLines {
		return "", fmt.Errorf("Start Line %d exceeds the total lines %d", start, totalLines)
	}
	
	if start < 1 {
		return "", fmt.Errorf("Start line must be greater than one.")
	}

	if end < start {
		return "", fmt.Errorf("End line must be greater than start.")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("Failed to open file %w", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var (
		lineNum int = 1
		result  string
		found   bool
	)

	for scanner.Scan() {
		if lineNum >= start && lineNum <= end {
			result += scanner.Text() + "\n"
			found = true
		}

		if lineNum > end {
			break
		}

		lineNum++
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("Failing while reading file %w", err)
	}

	if !found {
		return "", fmt.Errorf("Line range %d-%d is out of file bounds", start, end)
	}

	return result, nil
}

func main() {
	code, err := ExtractLines("C:\\Users\\chsan\\OneDrive\\Desktop\\CodeLabs\\College_Labs\\fileEmployee.cpp", 1010, 1000)
	if err != nil {
		panic(err)
	}
	fmt.Println(code)
}
