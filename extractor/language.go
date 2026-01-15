package extractor

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/alecthomas/chroma/lexers"
)

func DetectLanguage(path string) (string, error) {
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(path), "."))

	// Explicitly allow plain text
	if ext == "txt" {
		return "plaintext", nil
	}

	// Ask Chroma to match a lexer based on filename / extension
	lexer := lexers.Match(path)
	if lexer == nil {
		return "", errors.New("unsupported file type: no lexer found")
	}

	lang := lexer.Config().Name

	// Chroma fallback lexer → reject it
	if lang == "Text" {
		return "", errors.New("unsupported file type: not a programming language")
	}

	return lang, nil
}
