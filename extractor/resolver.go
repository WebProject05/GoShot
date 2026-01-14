package extractor

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func validatePath(path string) (string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	if info.IsDir() {
		return "", errors.New("Provided path is a directory, not a file")
	}

	return filepath.Abs(path)
}

func ResolveFilePath(input string) (string, error) {
	// If input already looks like the path
	if strings.ContainsAny(input, `/\`) {
		return validatePath(input)
	}

	matches := []string{} // Keep track of all the dupicate fils if exist

	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	err = filepath.WalkDir(cwd, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Base(path) == input {
			matches = append(matches, path)
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	if len(matches) > 1 {
		msg := fmt.Sprintf(
			"Multiple files named %s found: \n", input,
		)

		for _, m := range matches {
			msg += " - " + m + "\n"
		}

		msg += "Please provide the exact path."
		return "", errors.New(msg)
	}

	return filepath.Abs(matches[0])

}
