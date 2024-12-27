package main

// go-reloaded stuff
// simple version

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

var words []string

// var textRunes []rune

const QUOTE = "'"
const SPACE = " "
const EMPTY_STRING = ""
const STRING_EX string = "abc"
const Q_OPEN = true
const Q_CLOSE = false

// var punctuation = []rune {'.', ',', '!', '?', ':', ';', '(',
/*, ')', '[', ']', '{', '}', '<', '>',
'/', '\\', '|', '`', '~', '@', '#', '$',
'%', '^', '&', '*', '-', '+', '=', '_', '"' */
//					}

const PUNCTUATION string = ".,!?:;" // + "()"

// var vowels_h  = []rune  {'a', 'e', 'i', 'o', 'u', 'h', 'A', 'E', 'I', 'O', 'U', 'H'}

const VOWELS_H string = "aeiouhAEIOUH"

func main() {
	errStr := ""
	args := os.Args
	numArgs := len(args) - 1 // 1st is program name, 2nd is input file, 3rd is output file
	// print("numArgs:", numArgs)
	// print("args:", args)
	switch {
	case numArgs < 2:
		errStr = "less than 2 arguments"
		print(errStr)
	case numArgs > 2:
		errStr = "more than 2 arguments"
		print(errStr)
	case numArgs == 2:
		sampleInputFile := args[1]
		// filePointer, err := os.Open(sampleInput)
		inDat, err := os.ReadFile(sampleInputFile)
		check(err)
		resultOutputFile := args[2]
		// fp, err := os.Open(resultOutput) // os.Open only read
		// using os.OpenFile instead
		rOutputFile, err := os.OpenFile(resultOutputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644) // open the file write-only.
		check(err)
		defer rOutputFile.Close()

		// separate into words slice using fields to deal with multiple spaces
		words = strings.Fields(string(inDat))

		print("words after fields:", words, "\n")

		// process Words
		words, err := processWords(words)
		check(err)

		// TO DO remove RETURNS
		// join words back into string
		outStr := strings.Join(words, " ")
		// write to output file what a change
		rOutputBytes, err := rOutputFile.WriteString(outStr)
		check(err)
		fmt.Println("Finished! Written ", rOutputBytes, " bytes.")
	}
}

// basic error handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// BULK of string replace happens here
func processWords(wrds []string) (oWrds []string, err error) {

	wrds, err = processQuotes(wrds)
	check(err)
	// delete empty words
	// in quotes

	wrds, err = processPunctuation(wrds)
	check(err)
	// delete empty words
	// in puctuation

	// delete empty words (probaly not needed)
	// if wd == "" {
	// words = slices.Delete(words, i)
	//}
	//

	// TO DO:
	wrds, err = processAandAn(wrds)
	check(err)
	// delete empty words
	// in AandA

	// TO DO:
	// Process bin, hex,
	// and single of multiple words with
	// commands Cap(), Up(), Low()
	wrds, err = processCommands(wrds)
	check(err)

	return wrds, nil
}

// processQuotesInPutString processes the input string to handle single quotes and spaces.
// It iterates through the input string, identifying single quotes and spaces to determine
// the start and end indices of quoted sections.
//
// Parameters:
// sInput (string): The input string to be processed.
//
// Returns:
// sOutput (string): The processed output string.
// err (error): An error value if any error occurs during processing.
func processQuotes(wrds []string) (outWrds []string, err error) {
	quoteCount := 0
	openQuote := Q_CLOSE

	fmt.Printf("sOutputR: %v\n", wrds)
	// fmt.Printf("i: %cr: %c| %v\n", string(sOutputR))
	for iWd, wd := range wrds {
		// TO DO:
		if strings.HasPrefix(wd, QUOTE) && iWd > 0 && openQuote == Q_OPEN {
			// add quote to end of previous word
			wrds[iWd-1] = addSuffixToString(wrds[iWd-1], QUOTE)
			// trim quote from beginning of current word
			wd = cutStringPrefix(wd, QUOTE)
			// update quoteCount to reflect new quote added to previous word
			quoteCount++
			if wd == EMPTY_STRING {
				// delete empty word
				wrds = slices.Delete(wrds, iWd, iWd+1)
			}
		} 

		// count quotes
		quoteCount += strings.Count(wd, QUOTE)
		// if quoteCount is odd, then there is an open quote
		openQuote = quoteCount%2 == 1
		// if wd has quote at end and openQuote is Q_CLOSE
		// move quote to next word
		// trim quote from end of current word
		if strings.HasSuffix(wd, QUOTE) && iWd < len(wrds)-1 && openQuote == Q_CLOSE {
			wrds[iWd+1] = addPrefixToString(QUOTE, wrds[iWd+1])
			wd = deleteRuneFromString(wd, len(wd)-1)
		}

	}

	// if it is openQuote = Q_OPEN and current wd contains
	/* for i, r := range wd {
			// remember quote index if followed by space and open quote
			if isQuote(r) {
				quoteCount++
				if quoteCount%2 == 1 {
					openQuote = Q_OPEN
				} else {
					openQuote = Q_CLOSE
				}
			}

			// mv Q_CLOSE quote from  beginning of current word to end of previous word
			if  openQuote == Q_CLOSE && i == 0 && !outOfWordsBounds(wrds, iWd, wrds[iWd-1]) {
				wrds[iWd-1] = addRuneSuffixToString(QUOTE, wd)
				wd = deleteRuneFromString(wd, i)
			}

			if isSpace(r) && i > 0 && isQuote(outWords[i-1]) && openQuote {
				openQuoteIndex = i - 1
				mvQuote = true
			}

			// remember first space directly after end of last word
			if isSpace(r) && i > 0 && !isSpace(sOutputR[i-1]) {
				spcAfterLastWordIndex = i
			}

			if isQuote(r) && i > 0 && isSpace(outWords[i-1]) && !openQuote {
				outWords[i] = SPACE
				outWords[spcAfterLastWordIndex] = QUOTE
			}

			// if not a space and previous was a space & openQuote & mvQuote
			// move quote to previous space
			if !isSpace(r) && i > 0 && isSpace(outWords[i-1]) && openQuote && mvQuote {
				outWords[i-1] = QUOTE
				outWords[openQuoteIndex] = SPACE
				mvQuote = false
			}

			// print current state
			fmt.Printf("i: %2dr: %c| %v\n", i, r, string(outWords))
			// TO DO: remove debug print
		}
	}

	if quoteCount%2 == 1 {
		return "", errors.New("quotes are not in pairs")
		}
		// fmt.Printf("sOutputR: %v\n", string(outWords))
		return outWords, nil */
	return wrds, nil
	}

func cutStringPrefix(wd, pref string) string {
	wd = strings.TrimPrefix(wd, pref)
	return wd
}
func processCommands(wrds []string) ([]string, error) {
	panic("unimplemented")
}

func addSuffixToString(s, suff string) string {
	panic("unimplemented")
}
func deleteRuneFromString(s string, i int) string {
	panic("unimplemented")
}

func addPrefixToString(s, pref string) (string) {
	s = pref + s
	return s
}



// if word is an "a" or "A" and is followed by a word starting
// with a vowel or "h"
// then append 'n' after "a/a" of current word
func processAandAn(wrds []string) (oWrds []string, err error) {
	// check bounds
	lastWrd := len(wrds)-1
    for i := 0; i <= lastWrd - 1; i++ { // 
	if (wrds[i] == "a" || wrds[i] == "A") && 
	(i <= lastWrd && strings.IndexAny(wrds[i+1], VOWELS_H) == 0) {
	    Owrds[i] = Owrds[i] + "n"
		// TO DO: check if this is correct 2024-12-27_19-13 
	    }
    }
	
	// is vowel of "h" {
		

	return Owrds, errors.New("index out of bounds")
	}


// outputOfBounds(i, i-5) or outOfBounds(i, i+1)
func outOfWordsBounds(wrds []string, i, b int) (bool, error) {
	if i < 0 || b < 0 || i >= len(wrds) || b >= len(wrds) {
		return true, errors.New("index out of []Words bounds")
		// error message could be tidied up
	}
	return false, nil
}

/*func outOfStringBounds(wd string, i, b int) (bool, error) {
	if i < 0 || b < 0 || i >= len(wd) || b >= len(wd) {
		return true, errors.New("index out of string's bounds")
		// error message could be tidied up
	}
	return false, nil
}*/

//func processQuotes(wdIn string, qtOpenI bool, num_word int) (wdOut string, qtOpenR bool) {
//	// go through string
//	wdOut = ""
//	for pos, r := range []rune(wdIn) {
//		// if qtOpenI && r == SNGL_QUOTE {
//		if isQuote(r) { // TO DO
//			qtOpenI = !qtOpenI
//		} else if !qtOpenI && pos == 0 {
//			wdOut += string(r)
//			// words[4], words[3] = words[3], words[4]
//			// words[num_word]
//		}
//		if isQuote(r) && pos == len(wdIn)-1 {
//			if len(words) >= 2 {
//				words[num_word+1] = addRunePrefixToString(
//					QUOTE,
//					words[num_word+1],
//				)
//			}
//		}
//	}
//	return wdOut, qtOpenI
//}

func isQuote(r rune) bool {
	return r == QUOTE
}

/* func addRunePrefixToString(r rune, s string) string {
	cpRs := make([]rune, len(s)+1)
	cpRs[0] = r
	copy(cpRs[1:], []rune(s))
	return string(cpRs)
} */

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}

// processPunctuation processes the input string to handle punctuation.
// It iterates through the input string, identifying punctuation characters and spaces to determine
// the start and end indices of punctuation sections and eliminate spaces between them
func processPunctuation(wrds []string) (outWords []string, err error) {
	// tag first space character after non-space character
	// if it's punctuation move character to position of first space character after non-space character
	spcAfterLastWordIndex := -1

	wasSpace := false
	outWords := []string(wrds) // copied as rune array for safe easy
	// indexing and access to previous runes
	// var puncToMove rune

	fmt.Printf("puncutation: >>%v<<, punc outWords: %v\n", string(outWords))
	for i, r := range outWords {
		// set spcAfterLastWordIndex to index of first space after non-space character
		if isSpace(r) && i > 0 && !isSpace(outWords[i-1]) {
			spcAfterLastWordIndex = i
			wasSpace = true
		}
		// TO DO: is this correct??
		if !isSpace(r) && i > 0 && isSpace(outWords[i-1]) && wasSpace {
			outWords[i-1] = r
			outWords[i] = ' '
			wasSpace = false
		}

		// move all PUNCTUATION to left after non-space character

		// move spaces along until reach current i index
		// TO DO:
		//for isSpace(r) && {
		// move spaces along until reach current i index
		//}

		fmt.Printf("punc i: %2dr: %c| %v\n", i, r, string(outWords))
	}
	// TODO: check for punctuation at end of string
	// outWords = string(outWords)
	return outWords, nil
	// return outWords, nil
	// err := errors.New("punctuation processing failed")
	// return "", err
	// return sOutput, nil

}

func isPunctuation(r rune) bool {
	if slices.Contains(punctuation, r) {
		return true
	}
	return false
}

/*

func main() {
	f := func(c rune) bool {
		return !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc("(bin  foo1;bar2,baz3)", f))
}

*/
