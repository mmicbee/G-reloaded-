package main

import (
	"fmt"
	"os"
	"text-editor/processor"
)

func main() {
	if len(os.Args)!= 3 {
		fmt.Println("Usage: go run. input.txt output.txt")
		os.Exit(1)
	}

	input, err := os.ReadFile(os.Args[1])
	if err!= nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	result := processor.Process(string(input))
	
	err = os.WriteFile(os.Args[2], []byte(result), 0644)
	if err!= nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}