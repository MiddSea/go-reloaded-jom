// from https://pkg.go.dev/os#Exit
package main

import (
	"fmt"
	"os"
)

func main() {
	show := func(key string) {
		val, ok := os.LookupEnv(key)
		if !ok {
			fmt.Printf("%s not set\n", key)
		} else {
			fmt.Printf("%s=%s\n", key, val)
		}
	}

	os.Setenv("SOME_KEY", "value")
	os.Setenv("EMPTY_KEY", "")

	show("SOME_KEY")
	show("EMPTY_KEY")
	show("MISSING_KEY")

	env := os.Environ()
	for _, e := range env {
		fmt.Println(e)
	}
	show("HOME")
}
