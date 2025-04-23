package goreloaded

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// handleSingleQuotes processes text to properly format single quotes according to typographical rules
// It handles opening and closing quotes by attaching them to the appropriate words
// Returns a new slice with properly formatted quotes
func handleSingleQuotes(s []string) []string {
	// Create a result slice to store the processed words
	result := make([]string, 0, len(s))
	i := 0
	isOpening := 0 // Tracks if we've seen an opening quote

	for i < len(s) {
		// Case 1: Handle standalone quote character
		if s[i] == "'" {
			if i+1 < len(s) && isOpening == 0 {
				// This is an opening quote - attach it to the next word
				s[i+1] = "'" + s[i+1]
				isOpening++
				i++
			} else {
				// This is a closing quote - attach it to the previous word
				if len(result) > 0 {
					result[len(result)-1] = result[len(result)-1] + "'"
				}
				i++
				isOpening = 0
			}
			continue
		}

		// Case 2: Handle word that starts with a quote
		if strings.HasPrefix(s[i], "'") {
			// This is likely an opening quote - keep as is
			result = append(result, s[i])
			isOpening++
			i++
			continue
		}

		// Case 3: Regular word with no quote handling needed
		result = append(result, s[i])
		i++
	}

	return result
}

// CheckPunctuation handles punctuation marks in text by ensuring they're properly attached
// to the preceding words according to standard typographical rules
// Returns a new slice with properly formatted punctuation
func CheckPunctuation(s []string) []string {
	// First, handle single quotes with the specialized function
	s = handleSingleQuotes(s)

	// Process the array from left to right
	i := 0
	for i < len(s) {
		// Skip empty strings
		if strings.TrimSpace(s[i]) == "" {
			s = append(s[:i], s[i+1:]...)
			continue
		}

		// Check if the current word consists only of punctuation marks
		isPunct := true
		for _, char := range strings.TrimSpace(s[i]) {
			if !strings.ContainsRune(".,;:!?'", char) {
				isPunct = false
				break
			}
		}

		if isPunct {
			// If there's a previous word, attach this punctuation to it
			if i > 0 {
				s[i-1] += strings.TrimSpace(s[i])
				// Remove the current punctuation-only element
				s = append(s[:i], s[i+1:]...)
				// Don't increment i since we removed an element
			} else {
				// If it's at the beginning, just move to next word
				i++
			}
		} else {
			// Check if the word starts with punctuation (except for the first word)
			if i > 0 {
				j := 0
				// Count leading punctuation characters
				for j < len(s[i]) && strings.ContainsRune(".,;:!?", rune(s[i][j])) {
					j++
				}

				if j > 0 {
					// Move starting punctuation to the previous word
					s[i-1] += s[i][:j]
					s[i] = s[i][j:]

					// If the current word is now empty, remove it
					if s[i] == "" {
						s = append(s[:i], s[i+1:]...)
						continue
					}
				}
			}

			// Handle embedded punctuation (e.g., "word,next")
			for _, punct := range []string{".", ",", ";", ":", "!", "?"} {
				if strings.Contains(s[i], punct) && !strings.HasPrefix(s[i], punct) && !strings.HasSuffix(s[i], punct) {
					parts := strings.SplitN(s[i], punct, 2)
					if len(parts) == 2 {
						// Replace the current element with the first part + punctuation
						s[i] = parts[0] + punct

						// Insert the second part as a new element
						if i+1 < len(s) {
							s = append(s[:i+1], append([]string{parts[1]}, s[i+1:]...)...)
						} else {
							s = append(s, parts[1])
						}
					}
				}
			}

			i++
		}
	}

	return s
}

// CheckVowel determines if a string starts with a vowel
// Used for handling "a" vs "an" article selection
// Returns true if the string starts with a vowel
func CheckVowel(s string) bool {
	return strings.HasPrefix(strings.ToLower(s), "a") || strings.HasPrefix(strings.ToLower(s), "e") || strings.HasPrefix(strings.ToLower(s), "i") ||
		strings.HasPrefix(strings.ToLower(s), "o") || strings.HasPrefix(strings.ToLower(s), "u") || strings.HasPrefix(strings.ToLower(s), "y")
}

// CheckAorAn corrects the usage of "a" and "an" articles based on whether
// the following word starts with a vowel sound
// Returns a new slice with corrected articles
func CheckAorAn(s []string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == "a" || s[i] == "an" {
			if i+1 < len(s) && s[i] == "a" && (CheckVowel(s[i+1]) || strings.HasPrefix(s[i+1], "h")) {
				// Special handling for 'h' words - some 'h' words take 'an' despite starting with consonant
				if strings.HasPrefix(s[i+1], "h") {
					// Words like "hour", "honest" take "an" because the 'h' is silent
					if strings.HasPrefix(s[i+1], "hour") || strings.HasPrefix(s[i+1], "honest") ||
						strings.HasPrefix(s[i+1], "honor") || strings.HasPrefix(s[i+1], "heir") ||
						strings.HasPrefix(s[i+1], "herb") {
						s[i] = "an"
					}
				} else {
					// Words starting with vowels take "an"
					s[i] = "an"
				}
			} else if i+1 < len(s) && s[i] == "an" && !CheckVowel(s[i+1]) {
				// Words starting with consonants take "a"
				s[i] = "a"
			}
		}
	}

	return s
}

// BinaryToInteger converts a binary string to its decimal representation
// Returns the decimal value as a string
func BinaryToInteger(s string) string {
	num, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strconv.FormatInt(num, 10)
}

// HexadecimalToInteger converts a hexadecimal string to its decimal representation
// Returns the decimal value as a string
func HexadecimalToInteger(n string) string {
	num, err := strconv.ParseInt(n, 16, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strconv.FormatInt(num, 10)
}

// InitCap capitalizes the first letter of each word in a string
// and ensures the rest of each word is lowercase
// Returns the capitalized string
func InitCap(s string) string {
	runes := []rune(s)
	capitalize := true

	for i := 0; i < len(runes); i++ {
		if (runes[i] >= 'A' && runes[i] <= 'Z') || (runes[i] >= 'a' && runes[i] <= 'z') || (runes[i] >= '0' && runes[i] <= '9') {
			if capitalize {
				// Capitalize the first letter of a word
				if runes[i] >= 'a' && runes[i] <= 'z' {
					runes[i] = runes[i] - 32 // Convert to uppercase
				}
				capitalize = false
			} else {
				// Ensure the rest of the word is lowercase
				if runes[i] >= 'A' && runes[i] <= 'Z' {
					runes[i] = runes[i] + 32 // Convert to lowercase
				}
			}
		} else {
			// Non-alphanumeric character - next character should be capitalized
			capitalize = true
		}
	}

	return string(runes)
}

// CheckReq processes special formatting commands in the text
// Handles commands like (cap), (low), (up), (bin), (hex) and their variants
// Returns a new slice with the commands processed and removed
func CheckReq(arr []string) []string {
	for i := 0; i < len(arr); i++ {

		// Handle single-word formatting commands
		if strings.Contains(strings.ToLower(arr[i]), "(cap)") {
			if strings.ToLower(arr[i]) == "(cap)" {
				// Standalone command - apply to previous word
				if i-1 < 0 {
					arr = arr[i+1:] // Remove command if no previous word
				} else {
					arr[i-1] = InitCap(arr[i-1])
					arr = append(arr[:i], arr[i+1:]...) // Remove the command
					i--
				}
			} else {
				// Command attached to the end of the word - apply to current word
				if strings.HasSuffix(strings.ToLower(arr[i]), "(cap)") {
					arr[i] = InitCap(strings.TrimSuffix(strings.ToLower(arr[i]), "(cap)"))
				} else if strings.HasPrefix(strings.ToLower(arr[i]), "(cap)") && i-1 > 0 {
					// Command attached at the beginning of the word - apply to rhe previous word
					arr[i-1] = InitCap(arr[i-1])
					arr[i] = strings.TrimPrefix(strings.ToLower(arr[i]), "(cap)")
				}
				//arr[i] = InitCap(strings.TrimSuffix(strings.ToLower(arr[i]), "(cap)"))
			}
		} else if strings.Contains(strings.ToLower(arr[i]), "(low)") {
			if strings.ToLower(arr[i]) == "(low)" {
				// Standalone command - apply to previous word
				if i-1 < 0 {
					arr = arr[i+1:] // Remove command if no previous word
				} else {
					arr[i-1] = strings.ToLower(arr[i-1])
					arr = append(arr[:i], arr[i+1:]...) // Remove the command
					i--
				}
			} else {
				// Command attached to the end of the word - apply to current word
				if strings.HasSuffix(strings.ToLower(arr[i]), "(low)") {
					arr[i] = strings.ToLower(strings.TrimSuffix(strings.ToLower(arr[i]), "(low)"))
				} else if strings.HasPrefix(strings.ToLower(arr[i]), "(low)") && i-1 > 0 {
					// Command attached at the beginning of the word - apply to rhe previous word
					arr[i-1] = strings.ToLower(arr[i-1])
					arr[i] = strings.TrimPrefix(strings.ToLower(arr[i]), "(low)")
				}
				//arr[i] = strings.ToLower(strings.TrimSuffix(strings.ToLower(arr[i]), "(low)"))
			}
		} else if strings.Contains(strings.ToLower(arr[i]), "(up)") {
			if strings.ToLower(arr[i]) == "(up)" {
				// Standalone command - apply to previous word
				if i-1 < 0 {
					arr = arr[i+1:] // Remove command if no previous word
				} else {
					arr[i-1] = strings.ToUpper(arr[i-1])
					arr = append(arr[:i], arr[i+1:]...) // Remove the command
					i--
				}
			} else {
				// Command attached to the end of the word - apply to current word
				if strings.HasSuffix(strings.ToLower(arr[i]), "(up)") {
					arr[i] = strings.ToUpper(strings.TrimSuffix(strings.ToLower(arr[i]), "(up)"))
				} else if strings.HasPrefix(strings.ToLower(arr[i]), "(up)") && i-1 > 0 {
					// Command attached at the beginning of the word - apply to rhe previous word
					arr[i-1] = strings.ToUpper(arr[i-1])
					arr[i] = strings.TrimPrefix(strings.ToLower(arr[i]), "(up)")
				}
				//arr[i] = strings.ToUpper(strings.TrimSuffix(strings.ToLower(arr[i]), "(up)"))
			}
		} else if strings.Contains(strings.ToLower(arr[i]), "(bin)") {
			if strings.ToLower(arr[i]) == "(bin)" {
				// Standalone command - apply to previous word
				if i-1 < 0 {
					arr = arr[i+1:] // Remove command if no previous word
				} else {
					arr[i-1] = string(BinaryToInteger(arr[i-1]))
					arr = append(arr[:i], arr[i+1:]...) // Remove the command
					i--
				}
			} else {
				// Command attached to the end of the word - apply to current word
				if strings.HasSuffix(strings.ToLower(arr[i]), "(bin)") {
					arr[i] = string(BinaryToInteger(strings.TrimSuffix(strings.ToLower(arr[i]), "(bin)")))
				} else if strings.HasPrefix(strings.ToLower(arr[i]), "(bin)") && i-1 > 0 {
					// Command attached at the beginning of the word - apply to rhe previous word
					arr[i-1] = BinaryToInteger(arr[i-1])
					arr[i] = strings.TrimPrefix(strings.ToLower(arr[i]), "(bin)")
				}
				//arr[i] = string(BinaryToInteger(strings.TrimSuffix(strings.ToLower(arr[i]), "(bin)")))
			}
		} else if strings.Contains(strings.ToLower(arr[i]), "(hex)") {
			if strings.ToLower(arr[i]) == "(hex)" {
				// Standalone command - apply to previous word
				if i-1 < 0 {
					arr = arr[i+1:] // Remove command if no previous word
				} else {
					arr[i-1] = string(HexadecimalToInteger(arr[i-1]))
					arr = append(arr[:i], arr[i+1:]...) // Remove the command
					i--
				}
			} else {
				// Command attached to the end of the word - apply to current word
				if strings.HasSuffix(strings.ToLower(arr[i]), "(hex)") {
					arr[i] = string(HexadecimalToInteger(strings.TrimSuffix(strings.ToLower(arr[i]), "(hex)")))
				} else if strings.HasPrefix(strings.ToLower(arr[i]), "(hex)") && i-1 > 0 {
					// Command attached at the beginning of the word - apply to rhe previous word
					arr[i-1] = HexadecimalToInteger(arr[i-1])
					arr[i] = strings.TrimPrefix(strings.ToLower(arr[i]), "(hex)")
				}
				//arr[i] = string(HexadecimalToInteger(strings.TrimSuffix(strings.ToLower(arr[i]), "(hex)")))
			}
		}

		// Handle multi-word formatting commands with numeric parameter
		// Format: (command,n) where n is the number of words to affect
		if strings.Contains(strings.ToLower(arr[i]), "(cap,") && len(arr[i]) == 5 {
			// Extract the number from the next token
			//back, err := strconv.Atoi(strings.TrimSuffix(arr[i+1][0:], ")"))
			back, err := strconv.Atoi(trimBrackets(arr[i+1]))

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if back < 0 {
				back = -back
			}

			if strings.ToLower(arr[i]) == "(cap," {
				// Standalone command - apply to previous n words
				if i-back < 0 {
					arr = append(arr[:i], arr[i+2:]...) // Remove command if not enough previous words
				} else {
					// Apply capitalization to the specified number of previous words
					for j := 1; j <= back; j++ {
						arr[i-j] = InitCap(arr[i-j])
					}
					// Remove the command tokens
					arr = append(arr[:i], arr[i+2:]...)
					i--
				}
				//[harold, hhh, (cap, 2))]
			}
		} else if strings.Contains(strings.ToLower(arr[i]), "(low,") && len(arr[i]) == 5 {
			back, err := strconv.Atoi(trimBrackets(arr[i+1]))

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if back < 0 {
				back = -back
			}

			if strings.ToLower(arr[i]) == "(low," {
				if i-back < 0 {
					arr = append(arr[:i], arr[i+2:]...)
				} else {
					for j := 1; j <= back; j++ {
						arr[i-j] = strings.ToLower(arr[i-j])
					}
					// Remove the command token after processing all words
					arr = append(arr[:i], arr[i+2:]...)
					i--
				}

			}

		} else if strings.Contains(strings.ToLower(arr[i]), "(up,") && len(arr[i]) == 4 {
			back, err := strconv.Atoi(trimBrackets(arr[i+1]))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if back < 0 {
				back = -back
			}

			if strings.ToLower(arr[i]) == "(up," {
				if i-back < 0 {
					arr = append(arr[:i], arr[i+2:]...)
				} else {
					for j := 1; j <= back; j++ {
						arr[i-j] = strings.ToUpper(arr[i-j])
					}
					// Remove the command token after processing all words
					arr = append(arr[:i], arr[i+2:]...)
					i--
				}

			}
		}
	}

	return arr
}

func trimBrackets(s string) string {
	for strings.Contains(s, ")") {
		s = strings.TrimSuffix(s, ")")
	}

	return s
}
