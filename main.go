package main

import (
	"log"
	"os"
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
	resultTxt := processTxt(sampleTxt)
	err = writeResult(resultOutFile, resultTxt)
	checkError(err)
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
		return nil	// return true if 3 arguments
	default:
		err = fmt.Errorf("too many arguments")
		return err // return false and error message if more than 3 arguments
	}
}

func readSampleFile(filename string) (content string, err error){
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

func processTxt(txt string) string {
	// split text into words ignoring multiple white spaces.
	words := strings.Fields(txt)  
	quoteCount := 0
	for i, word := range words {
		lastWordIndex := len(words)-1
		islastWord := i == lastWordIndex
		switch word {
		case "A", "a":
			print(">isLastWord:,", islastWord, "|a or A: ", i,"len(words)", len(words),"<")
			if !islastWord && strings.Contains("aeiouhAEIOUH", string(words[i+1][0])) { // check next word in range
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
			if strings.Contains("aeiouhAEIOUH", string(words[i+1][0])) {
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
		case "(hex)":
			words[i-1] = strconv.Itoa(Hex2Dec(words[i-1]))
		case "(cap)":
			Cap(&words[i-1])
		Distributed version control systems (DVCS)case "(low)":
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
func readSampleFile(filename string) string {
	content, _ := os.ReadFile(filename)
	return string(content)
}
*/
