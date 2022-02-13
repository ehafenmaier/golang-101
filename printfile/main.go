package main

import (
	"fmt"
	"os"
)

func main() {
	// Store the name of the file passed in to a separate variable
	fileName := os.Args[1]

	// Use the ReadFile function to get the contents of
	// the file as a byte slice
	contents, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	// Print out the byte slice contents as a string
	fmt.Println(string(contents))
}