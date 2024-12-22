package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	// "errors"
	"fmt"
)

func main() {
	err := checkArgs(os.Args)
	checkError(err)

	sampleInFile := os.Args[1]
	resultOutFile := os.Args[2]

	sampleTxt, err := readSampleFile(sampleInFile)
	checkError(err)

	resultTxt, err := processTxt(sampleTxt)
	checkError(err)

	err = writeResult(resultOutFile, resultTxt)
	checkError(err)
	fmt.Println("Conversion successful. Result saved in:", resultOutFile)

}

func checkArgs(args []string) (err error) {
	switch len(args) {
	case 1:
		err = fmt.Errorf("no input / output files given")
		return err // return false and error message if less than 2 arguments
	case 2:
		err = fmt.Errorf("only input, no output file given")
		return err // return false and error message if less than 3 arguments
	case 3:
		return nil // return true if 3 arguments
	default:
		err = fmt.Errorf("too many arguments")
		return err // return false and error message if more than 3 arguments
	}
}

func readSampleFile(filename string) (content string, err error) {
	var contentB []byte
	contentB, err = os.ReadFile(filename)
	if err != nil {
		log.Panicf("failed to readSample  file: %v", err) // needs log package
	}
	return string(contentB), err

}

func writeResult(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644) //
	if err != nil {
		log.Panicf("failed to writeResult to file: %v", err) // needs log package
	}
	return err
}

func checkError(err error) {
	if err != nil {
		log.Panicf("error: %v", err)
	}
}

func processTxt(txt string) (oTxt string, err error) {
	// split text into words ignoring multiple white spaces.
	words := strings.Fields(txt)
	// quoteCount := 0

	words, err = processQuotes(words)
	checkError(err)

	words, err = processPunctuation(words)
	checkError(err)

	words, err = processAorAn(words)
	checkError(err)

	// process commands bin, hex,
	// and cap, low, up for single words multiple words.
	words, err = processCommands(words)
	checkError(err)

	return strings.Join(words, " "), nil
}
func processQuotes(words []string) ([]string, error) {
	countQuote := 0
	openQuote := false
	print("processing quotes...")
	for i := 0; i < len(words); i++ {
		print("\n\tword ", i)
		word := words[i]
		print(":", word, " ")
		openQuote = countQuote%2 == 1
		if word == "one" {
			words = slices.Delete(words, i, i+1)
		}
		if strings.HasPrefix(word, "'") && openQuote {
			countQuote++
			word = strings.TrimPrefix(word, "'")
			if word == "" {
				words = slices.Delete(words, i, i+1)
			}

			print(" countQuote:", countQuote)
			prevWord := words[i-2]
			print(" prevWord:", countQuote)

			words[i-2] = prevWord + "'"
			countQuote += strings.Count(word, "'")
		}
		if strings.HasSuffix(word, "'") && !openQuote {
			words[i] = strings.TrimSuffix(word, "'")
			if words[i] == "" {
				words = slices.Delete(words, i, i+1)
				i--
			}
		}
	}
	print("Processing quotes...")
	return words, nil
}

// if first character of word is punctuation take all punctuation characters that are a part of the
// prefix of the word and add them to the previous word
// and remove the punctuation from the current word.
func processPunctuation(words []string) ([]string, error) {

	/*for i, word := range words {

	if strings.Contains(".,!?:;", string(word[0])) {
		if i > 0 {
			words[i-1] += string(word[0])
		} else {
			words[i-1] += string(word[0])
		}
		words[i] = word[1:]
	}
				countQuote++}  */
	print("Processing punctuation...")
	return words, nil
}
func processAorAn(words []string) ([]string, error) {
	for i, word := range words {
		if (word == "A" || word == "a") && i < len(words)-1 {
			nextWord := words[i+1]
			if strings.ContainsAny(string(nextWord[0]), "aeiouhAEIOUH") {
				words[i] = word + "n"
			}
		}
	}
	print("Processing punctuation...")
	return words, nil
}

/*func returnPrefixPunctuation(word string) string {
	punctuation := ""
	for _, c := range word {
		if strings.Contains(".,!?:;", string(c)) {
			punctuation += string(c)
		} else {
			break
		}
	}
	return punctuation
}*/

func processCommands(words []string) ([]string, error) {
	// panic("UNIMPLEMENTED")
	print("UNIMPLEMENTED_")
	print("Processing commands...")
	return words, nil
}

/*	for i, word := range words {
		lastWordIndex := len(words) - 1
		islastWord := i == lastWordIndex
		switch word {
		case "'":
			quoteCount++
			if quoteCount%2 == 1 {
				words[i+1] = word + words[i+1]
			} else if strings.Contains(".,!?:;", words[i-1]) {
				words[i-2] += word
			} else {
				words[i-1] += word
			}
		case "A", "a":
			print(">isLastWord:", islastWord, "|a or A: ", i, "len(words)", len(words), "<")
			if !islastWord && strings.Contains("aeiouhAEIOUH", string(words[i+1][0])) { // check next word in range
				words[i] += "n"
			}
		case "(bin)":
			words[i-1] = strconv.Itoa(Bin2Dec(words[i-1]))
		case "(hex)":
			words[i-1] = strconv.Itoa(Hex2Dec(words[i-1]))
		case "(cap)":
			Cap(&words[i-1])
		case "(low)":
			Lower(&words[i-1])
		case "(up)":
			Upper(&words[i-1])
		case "(cap,":
			digit := extractDigit(words[i+1])
			applyFunc2Many(Cap, digit, i-1, words)
		case "(low,":
			digit := extractDigit(words[i+1])
			applyFunc2Many(Lower, digit, i-1, words)
		case "(up,":
			digit := extractDigit(words[i+1])
			applyFunc2Many(Upper, digit, i-1, words)
		default:
			if strings.Contains("...,!?:;", string(word[0])) {
				if string(words[i-1][1]) == ")" {
					words[i-3] += string(word[0])
				} else {
					words[i-1] += string(word[0])
				}
				words[i] = word[1:]
			}
		}
	}
	return strings.Join(removeCommands(words), " ")
}

/*
/*	for i, word := range words {
package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	processedContent := processText(readSampleFile("./sample.txt"))
	println(processedContent)
	os.WriteFile("./result.txt", []byte(processedContent), 0644)
}

func processText(text string) string {
	words := strings.Fields(text)
	quoteCount := 0
	for i, word := range words {
		switch word {
		case "A", "a":
			if stringsProcessing.Contains("aeiouhAEIOUH", string(words[i+1][0])) {
				words[i] += "n"
			}
		case "'":
			quoteCount++
			if quoteCount%2 == 1 {
				words[i+1] = word + words[i+1]
			} else if strings.Contains("...,!?:;", words[i-1]) {
				words[i-2] += word
			} else {
				words[i-1] += word
			}
		case "(bin)":
			words[i-1] = strconv.Itoa(Bin2Dec(words[i-1]))
		case "(hex)":Processing
			words[i-1] = strconv.Itoa(Hex2Dec(words[i-1]))
		case "(cap)":d
			Cap(&words[i-1])
		case "(low)":
			Lower(&words[i-1])
		case "(up)":
			Upper(&words[i-1])
		case "(cap,":
			digit := extractDigit(words[i+1])
			applyFunc2Many(Cap, digit, i-1, words)
		case "(low,":
			digit := extractDigit(words[i+1])
			applyFunc2Many(Lower, digit, i-1, words)
		case "(up,":
			digit := extractDigit(words[i+1])
			applyFunc2Many(Upper, digit, i-1, words)
		default:
			if strings.Contains("...,!?:;", string(word[0])) {
				if string(words[i-1][1]) == ")" {
					words[i-3] += string(word[0])
				} else {
					words[i-1] += string(word[0])

				words[i] = word[1:]
			}
		}
	}
	return strings.Join(removeCommands(words), " ")
}
*/

func extractDigit(word string) int {
	myint, _ := strconv.Atoi(word[:len(word)-1])
	return myint
}

// assuming start is i-1 and d is number extracted from 3) for example.
// apply function f to the last d words starting from start
func applyFunc2Many(f func(*string), d int, start int, words []string) {
	for i := start; i > start-d; i-- {
		f(&words[i])
	}
}

func removeCommands(words []string) []string {
	result := []string{}
	removeNext := false
	for _, word := range words {
		if word == "(cap)" || word == "(low)" || word == "(up)" || word == "(bin)" || word == "(hex)" {
			continue
		}
		if word == "(cap)," || word == "(low)," || word == "(up)," || word == "(bin)," || word == "(hex)," {
			result[len(result)-1] += ","
			continue
		}
		if word == "(cap," || word == "(low," || word == "(up," {
			removeNext = true
			continue
		}
		if removeNext {
			removeNext = false
			continue
		}
		if len(word) > 1 && OnlyPunctuation(word) {
			result[len(result)-1] += word
			continue
		}
		if strings.Contains(".,!?:;'", word) { // remove punctuation
			continue
		}
		result = append(result, word)
	}
	return result
}

func OnlyPunctuation(word string) bool {
	for _, c := range word {
		if !strings.Contains(".,!?:;'", string(c)) {
			return false
		}
	}
	return true
}

// convert binary number to digital
func Bin2Dec(binary string) int {
	result, _ := strconv.ParseInt(binary, 2, 64)
	return int(result)
}

// convert hex number to digital
func Hex2Dec(hex string) int {
	result, _ := strconv.ParseInt(hex, 16, 64)
	return int(result)
}

func Cap(str *string) {
	if len(*str) > 0 {
		*str = strings.ToUpper((*str)[:1]) + (*str)[1:]
	}
}

func Lower(str *string) {
	*str = strings.ToLower(*str)
}

func Upper(str *string) {
	*str = strings.ToUpper(*str)
}

/*
func readSampleFile(filename sProcessingtring) string {
	content, _ := os.ReadFile(filename)
	return string(content)
}
*/
