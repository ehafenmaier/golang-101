package main

import (
	"fmt"
	"io"
	"os"
)

type fileWriter struct {}

func main() {
	// Store the name of the file passed in to a separate variable
	fileName := os.Args[1]

	// Open the file as read only
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Create a fileWriter
	fw := fileWriter{}

	// Use io.Copy to copy file contents to new fileWriter
	io.Copy(fw, file)
}

// Function with fileWriter receiver that adheres to io.Writer interface
func (fileWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("The file contents is this many bytes:", len(bs))

	return len(bs), nil
}