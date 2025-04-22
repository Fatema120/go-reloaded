package goreloaded

import (
	"strings"
)

func TextEditor(s string) string {
	// Split the input string into individual words
	words := strings.Split(s, " ")

	// Process formatting commands like (cap), (up), (low), (bin), (hex)
	words = CheckReq(words)

	// Handle punctuation marks to ensure proper spacing and attachment
	words = CheckPunctuation(words)

	// Process formatting commands like (cap), (up), (low), (bin), (hex) if there was existing commands with no spaces e.g., (cap,2)
	words = CheckReq(words)

	// Correct usage of "a" vs "an" based on following words
	words = CheckAorAn(words)

	// Join the processed words back into a single string with spaces
	return strings.Join(words, " ")
}
