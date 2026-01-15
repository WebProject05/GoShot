package renderer

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)



func HighlightToHTML(
	lines []string,
	language string,
	theme string,
) (string, error) {
	code := strings.Join(lines, "\n")

	var lexer chroma.Lexer

	if language == "plaintext" {
		lexer = lexers.Fallback
	} else {
		lexer = lexers.Get(language)
	}

	if lexer == nil {
		return "", fmt.Errorf("No lexer found for the language: %s", language)
	}

	iterator, err := lexer.Tokenise(nil, code)
	if err != nil {
		return "", err
	}

	// Style selection
	style := styles.Get(theme)
	if style == nil {
		style = styles.Fallback
	}

	// Formatting the HTML of the extracted code
	formatter := html.New(
		html.WithLineNumbers(false),
		html.WithClasses(false),  // This is for the css
	)

	var buf bytes.Buffer

	err = formatter.Format(&buf, style, iterator)

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}