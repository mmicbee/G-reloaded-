package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Step 1: Validate terminal arguments
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Error: Invalid arguments.")
		fmt.Println("Usage: go run . <input_file> <output_file>")
		return
	}

	inputFile := args[0]
	outputFile := args[1]

	// Step 2: Read the input file using the Go File System API
	inputData, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	// Step 3: Core execution - process the raw text string
	rawText := string(inputData)
	processedText := ProcessText(rawText)

	// Step 4: Write the completed string to the output file
	err = os.WriteFile(outputFile, []byte(processedText), 0644)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		return
	}
}
