package main

import (
	"fmt"
	"os"
)

func matchBrackets(str string) bool {
	runes := []rune(str)
	var openBrkts []rune
	// contains opened brackaets
	//lastOpened := -1 //possibly don't need
	for _, r := range runes {
		switch r {
		case '(', '[', '{':
			openBrkts = append(openBrkts, r)
			//lastOpened++
		case ')':
			if len(openBrkts) >  0 && openBrkts[len(openBrkts)-1] != '(' {
				return false
			}
			openBrkts = openBrkts[:len(openBrkts)-1]
		case ']':
			if len(openBrkts) >  0 && openBrkts[len(openBrkts)-1] != '[' {
				return false
			}
			openBrkts = openBrkts[:len(openBrkts)-1]
		case '}':
			if len(openBrkts) >  0  && openBrkts[len(openBrkts)-1] != '{' {
				return false
			}
			openBrkts = openBrkts[:len(openBrkts)-1]
		}
	}
	return len(openBrkts) == 0
}

func main() {
	// for all arguments
	for _, arg := range os.Args[1:] {
		if matchBrackets(arg) {
			fmt.Printf("OK")
		} else {
			fmt.Printf("Error")
		}
	}
}
