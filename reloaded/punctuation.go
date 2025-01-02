package reloaded

import (
	"regexp"
	"strings"
)

// PunctuationFormat formats the punctuation marks in a string
func PunctuationFormatTerse(line string) string {
	// (1 was 1) Regular expression to match any spaces before single punctuation marks
//	spBfPunc := regexp.MustCompile(`[ \t]*(?:[.,!?:;])`)

// (1 was 1) Regular expression to match any spaces before any num punctuation marks
	spBfPunc := regexp.MustCompile(`[ \t]+?([.,!?:;]+?)`)

	// (2 was 3) Regular expression to match any spaces around single punctuation marks before an alphabetical character
//	spSglPunc := regexp.MustCompile(`([.,!?:;])([[:alpha:]])`)

	// (2 was 3) Regular expression to match any spaces around single punctuation marks before an alphabetical character
//	spSglPunc := regexp.MustCompile(`(?:[ \t])([.,!?:;])([[:alpha:][:digit:]])`)

// (2 was 3) Regular expression to match any spaces around single punctuation marks before an alphabetical character
//	spSglPunc := regexp.MustCompile(`(?:[ \t])([.,!?:;])([[:alpha:][:digit:]])`)
	
// (2 was 3) Regular expression to match any spaces around single punctuation marks before an alphabetical character

	// (3 was 2) Regular expression to match any spaces between single quotes
	//spBtwnSglQuotes := regexp.MustCompile(`'[ \t]*?([\S]*?)[ \t]*?'`)
	
	// (1 was 1) Replace the spaces with one space after the punctuation mark
//	line = spBfPunc.ReplaceAllString(line, "$1")

	// (1 was 1) Replace the spaces with one space after the punctuation mark
	line = spBfPunc.ReplaceAllString(line, "$1 ")

	// (2 was 3) Replace the spaces with one space after the punctuation mark before an alphabetical character
//	line = spSglPunc.ReplaceAllString(line, "$1 $2")

	// (3 was 2) Replace the spaces between single quotes with no spaces
	//line = spBtwnSglQuotes.ReplaceAllString(line, "'$1'")
	
	return line
}

func PunctuationFormatDetailed(line string) string {
    // First, handle multiple spaces and trim the line
    line = strings.TrimSpace(line)

    // Handle consecutive punctuation marks without spaces between them
    // Examples: "??!", "...", "?!"
    // We don't want spaces between consecutive punctuation
    consecutivePunctuation := regexp.MustCompile(`[ \t]*([.,:;!?]{2,})[ \t]*`)
    line = consecutivePunctuation.ReplaceAllString(line, "$1")

    // Handle colons before quotes and their spacing
    // Example: "said:" -> "said: "
    colonQuote := regexp.MustCompile(`(\w+)[ \t]*:[ \t]*'`)
    line = colonQuote.ReplaceAllString(line, "$1: '")

    // Handle spaces around individual punctuation marks
    // Ensures one space after punctuation unless it's followed by a quote or another punctuation
    // Removes spaces before punctuation
   // singlePunctuation := regexp.MustCompile(`[ \t]*([.,!?])[ \t]*(?![.,!?'])`)
   singlePunctuation := regexp.MustCompile(`[ \t]*([.,!?])[ \t]*([.,!?'])`)
 //  line = singlePunctuation.ReplaceAllString(line, "$1 $2")

    // Handle quoted text with proper spacing
    // Matches quotes that might have a colon before them
    // Group 1: Optional preceding text and colon
    // Group 2: The actual quoted content
    quotes := regexp.MustCompile(`(?:')[ \t]*(.*?)[ \t]*(?:')`)
    line = quotes.ReplaceAllStringFunc(line, func(match string) string {
        // Remove extra spaces within the quote
        inner := quotes.FindStringSubmatch(match)[1]
        // Apply punctuation formatting within the quote
        inner = singlePunctuation.ReplaceAllString(inner, "$1 ")
        inner = consecutivePunctuation.ReplaceAllString(inner, "$1")
        return "'" + strings.TrimSpace(inner) + "'"
    })

    // Clean up any double spaces that might have been created
    multipleSpaces := regexp.MustCompile(`[ \t]{2,}`)
    line = multipleSpaces.ReplaceAllString(line, " ")

    // Final trim to remove any trailing spaces
    return strings.TrimSpace(line)
}

func PunctuationRegEx(line string) string {
      // First, handle multiple spaces and trim the line
      // line = strings.TrimSpace(line)
	  // 
      trimSpaceTab := regexp.MustCompile(`(?m)(^[ \t]*)|([\t ]*$)`)
	  line = trimSpaceTab.ReplaceAllString(line, "")

      // Handle multiple sequential punctuation (like ..., ?!, etc.)
      // First combine them without spaces
      multiPunc := regexp.MustCompile(`[ \t]*([.,!?]+)[ \t]*`)
      line = multiPunc.ReplaceAllString(line, "$1")
      
      // Handle colons before quotes
      //colonQuote := regexp.MustCompile(`([\w]+)[ \t]*:[ \t]*'`)
      // added round brackets to include ) in the word
	  colonQuote := regexp.MustCompile(`([\w)]+)[ \t]*:[ \t]*'`)
      line = colonQuote.ReplaceAllString(line, "$1: '")
      
      // Handle quoted text
      quotes := regexp.MustCompile(`'([^']*)'`)
      line = quotes.ReplaceAllStringFunc(line, func(match string) string {
          inner := quotes.FindStringSubmatch(match)[1]
          // Clean up spaces inside quotes
          // inner = strings.TrimSpace(inner)
		  // use triSpaceTab instead of strings.TrimSpace
		  inner = trimSpaceTab.ReplaceAllString(inner, "")

          // Fix punctuation inside quotes
          inner = multiPunc.ReplaceAllString(inner, "$1")
          return "'" + inner + "'"
      })
      
      // Process single punctuation marks
      // We'll do this in multiple passes since we can't use lookahead
      
      // First, remove spaces before punctuation
      spaceBeforePunc := regexp.MustCompile(`[ \t]+([.,!?])`)
      line = spaceBeforePunc.ReplaceAllString(line, "$1")
      
      // Then add space after punctuation, but we'll clean up extras later
      spaceAfterPunc := regexp.MustCompile(`([.,!?])([^.,!?' \t])`)
      line = spaceAfterPunc.ReplaceAllString(line, "$1 $2")
      
      // Clean up any multiple spaces that might have been created
      multipleSpaces := regexp.MustCompile(`[ \t]{2,}`)
      line = multipleSpaces.ReplaceAllString(line, " ")
      
      // Final trim to remove any trailing spaces
      // return strings.TrimSpace(line)
	  // use triSpaceTab instead of strings.TrimSpace
	  line = trimSpaceTab.ReplaceAllString(line, "")
	  return line
	}

func AtoAnRegEx(line string) string {
	// Regular expression an 'a' or 'A' followed by any vowel (with optional single quotes)
//	aBfvowel_H := regexp.MustCompile(`([ \t]'?[aA])([ \t]'?[aeiouAEIOUhh])`)
	aBfvowel_H := regexp.MustCompile(`(^|[ \t]'?[aA])([ \t]'?[aeiouAEIOUhH])`)

	// replace 'a' -> 'an' and 'A' with 'An' if next word starts with a vowel or H
	line = aBfvowel_H.ReplaceAllString(line, `${1}n${2}`)

	return line
}