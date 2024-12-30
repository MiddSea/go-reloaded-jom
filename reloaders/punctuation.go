package helpers

import (
	"regexp"
)

// PunctuationFormat formats the punctuation marks in a string
func PunctuationFormat(line string) string {
	// Regular expression to match any spaces around single punctuation marks
	spaces := regexp.MustCompile(`\s*([.,!?:;])`)
	// Regular expression to match any spaces between single quotes
	spBtwnSglQuotes := regexp.MustCompile(`'\s*(.*?)\s*'`)
	// Regular expression to match any spaces around single punctuation marks before an alphabetical character
	spSglPunc := regexp.MustCompile(`([.,!?:;])([[:alpha:]])`)

	// Replace the spaces with one space after the punctuation mark
	line = spaces.ReplaceAllString(line, "$1")

	// Replace the spaces between single quotes with no spaces
	line = spBtwnSglQuotes.ReplaceAllString(line, "'$1'")

	// Replace the spaces with one space after the punctuation mark before an alphabetical character
	line = re3.ReplaceAllString(line, "$1 $2")

	return line
}
