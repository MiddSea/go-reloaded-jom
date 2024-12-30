package main

import (
	"fmt"

	"github.com/yourusername/reloaders"
)

func main() {
	// Create a string with punctuation marks
	line := "Hello, world! How are you doing? I'm doing well."
	// Call the PunctuationShift function from the reloaders package
	line = reloaders.PunctuationShift(line)
	// Print the modified string
	fmt.Println(line)
}
