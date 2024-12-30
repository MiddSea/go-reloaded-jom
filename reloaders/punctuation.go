package helpers

import (
	"regexp"
)

// PunctuationFormat formats the punctuation marks in a string
func PunctuationShift(line string) string {
	// Regular expression to match any spaces around single punctuation marks
	spacesBfPunc := regexp.MustCompile(`\s*([.,!?:;])`)
	// Regular expression to match any spaces between single quotes
	spBtwnSglQuotes := regexp.MustCompile(`'\s*(.*?)\s*'`)
	// Regular expression to match any spaces around single punctuation marks before an alphabetical character
	puncAlpha := regexp.MustCompile(`([.,!?:;])([[:alpha:]])`)

	// Replace the spaces with one space after the punctuation mark
	line = spacesBfPunc.ReplaceAllString(line, "$1")

	// Replace the spaces between single quotes with no spaces
	line = spBtwnSglQuotes.ReplaceAllString(line, "'$1'")

	// Replace the spaces with one space after the punctuation mark before an alphabetical character
	line = puncAlpha.ReplaceAllString(line, "$1 $2")

	return line
}
