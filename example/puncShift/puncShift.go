package main

import (
	"fmt"

	reloaders "reloaders/reloaders"
)

func main() {
	// Create a string with punctuation marks
	line := "Hello,   world !  How are you doing? I'm d .. . .  'oing ' well."
	fmt.Println(line)
	// Call the PunctuationShift function from the reloaders package
	line = reloaders.PunctuationShift(line)
	// Print the modified string
	fmt.Println(line)
}
