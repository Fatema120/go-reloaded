package main

import (
	"fmt"
	"goreloaded"
	"os"
	"strings"
)

func main() {
	// Check arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if !strings.HasSuffix(inputFile, ".txt") || !strings.HasSuffix(outputFile, ".txt") {
		fmt.Println("Error: Required .txt file")
		return
	}

	// Read input file
	inputData, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	// Process the text
	processedText := goreloaded.TextEditor(string(inputData))

	// Write to output file
	err = os.WriteFile(outputFile, []byte(processedText), 0644)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		return
	}

	fmt.Println("Processing complete. Output saved to:", outputFile)
}
