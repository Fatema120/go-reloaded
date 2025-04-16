package goreloaded

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
func CheckPunctuation(s []string) string {

}
*/

func CheckVowel(s string) bool {
	return strings.HasPrefix(s, "a") || strings.HasPrefix(s, "e") || strings.HasPrefix(s, "i") ||
		strings.HasPrefix(s, "o") || strings.HasPrefix(s, "u") || strings.HasPrefix(s, "y")
}

func CheckAorAn(s []string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == "a" || s[i] == "an" {
			if i+1 < len(s) && s[i] == "a" && (CheckVowel(s[i+1]) || strings.HasPrefix(s[i+1], "h")) {
				if strings.HasPrefix(s[i+1], "h") {
					if s[i+1] == "hour" || s[i+1] == "honest" || s[i+1] == "honor" || s[i+1] == "heir" || s[i+1] == "herb" {
						s[i] = "an"
					}
				} else {
					s[i] = "an"
				}
			} else if i+1 < len(s) && s[i] == "an" && !CheckVowel(s[i+1]) {
				s[i] = "a"
			}
		}
	}

	return s
}

func BinaryToInteger(s string) string {
	num, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strconv.FormatInt(num, 10)
}

func HexadecimalToInteger(n string) string {
	num, err := strconv.ParseInt(n, 16, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strconv.FormatInt(num, 10)
}

func InitCap(s string) string {
	runes := []rune(s)
	capitalize := true

	for i := 0; i < len(runes); i++ {
		if (runes[i] >= 'A' && runes[i] <= 'Z') || (runes[i] >= 'a' && runes[i] <= 'z') || (runes[i] >= '0' && runes[i] <= '9') {
			if capitalize {
				if runes[i] >= 'a' && runes[i] <= 'z' {
					runes[i] = runes[i] - 32
				}
				capitalize = false
			} else {
				if runes[i] >= 'A' && runes[i] <= 'Z' {
					runes[i] = runes[i] + 32
				}
			}
		} else {
			capitalize = true
		}
	}

	return string(runes)
}

func CheckReq(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		if strings.Contains(arr[i], "(cap)") {
			if arr[i] == "(cap)" {
				if i-1 < 0 {
					arr = arr[i+1:]
				} else {
					arr[i-1] = InitCap(arr[i-1])
					arr = append(arr[:i], arr[i+1:]...)
					i--
				}
			} else {
				arr[i] = InitCap(strings.TrimSuffix(arr[i], "(cap)"))
			}
		} else if strings.Contains(arr[i], "(low)") {
			if arr[i] == "(low)" {
				if i-1 < 0 {
					arr = arr[i+1:]
				} else {
					arr[i-1] = strings.ToLower(arr[i-1])
					arr = append(arr[:i], arr[i+1:]...)
					i--
				}
			} else {
				arr[i] = strings.ToLower(strings.TrimSuffix(arr[i], "(low)"))
			}
		} else if strings.Contains(arr[i], "(up)") {
			if arr[i] == "(up)" {
				if i-1 < 0 {
					arr = arr[i+1:]
				} else {
					arr[i-1] = strings.ToUpper(arr[i-1])
					arr = append(arr[:i], arr[i+1:]...)
					i--
				}
			} else {
				arr[i] = strings.ToUpper(strings.TrimSuffix(arr[i], "(up)"))
			}
		} else if strings.Contains(arr[i], "(bin)") {
			if arr[i] == "(bin)" {
				if i-1 < 0 {
					arr = arr[i+1:]
				} else {
					arr[i-1] = string(BinaryToInteger(arr[i-1]))
					arr = append(arr[:i], arr[i+1:]...)
					i--
				}
			} else {
				arr[i] = string(BinaryToInteger(strings.TrimSuffix(arr[i], "(bin)")))
			}
		} else if strings.Contains(arr[i], "(hex)") {
			if arr[i] == "(hex)" {
				if i-1 < 0 {
					arr = arr[i+1:]
				} else {
					arr[i-1] = string(HexadecimalToInteger(arr[i-1]))
					arr = append(arr[:i], arr[i+1:]...)
					i--
				}
			} else {
				arr[i] = string(HexadecimalToInteger(strings.TrimSuffix(arr[i], "(hex)")))
			}
		}


		if strings.Contains(arr[i], "(cap,") {
			back, err := strconv.Atoi(strings.TrimSuffix(arr[i+1][0:], ")"))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if arr[i] == "(cap," {
				if i-back < 0 {
					arr = append(arr[:i], arr[i+2:]...)
				} else {
					for j := 1; j <= back; j++ {
						arr[i-j] = InitCap(arr[i-j])
					}

					arr = append(arr[:i], arr[i+2:]...)
					i--
				}

			} else {
				arr[i] = InitCap(strings.TrimSuffix(arr[i], arr[i][len(arr[i])-5:]))
				for j := 1; j <= back-1; j++ {
					arr[i-j] = InitCap(arr[i-j])
				}

				arr = append(arr[:i+1], arr[i+2:]...) 
			}

		} else if strings.Contains(arr[i], "(low,") {
			back, err := strconv.Atoi(strings.TrimSuffix(arr[i+1][0:], ")"))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if arr[i] == "(low," {
				if i-back < 0 {
					arr = append(arr[:i], arr[i+2:]...)
				} else {
					for j := 1; j <= back; j++ {
						arr[i-j] = strings.ToLower(arr[i-j])
					}

					arr = append(arr[:i], arr[i+2:]...)
					i--
				}

			} else {
				arr[i] = strings.ToLower(strings.TrimSuffix(arr[i], arr[i][len(arr[i])-5:])) 
				for j := 1; j <= back-1; j++ {
					arr[i-j] = InitCap(arr[i-j])
				}

				arr = append(arr[:i+1], arr[i+2:]...) 
			}

		} else if strings.Contains(arr[i], "(up,") {
			back, err := strconv.Atoi(strings.TrimSuffix(arr[i+1][0:], ")")) 
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if arr[i] == "(up," {
				if i-back < 0 {
					arr = append(arr[:i], arr[i+2:]...)
				} else {
					for j := 1; j <= back; j++ {
						arr[i-j] = strings.ToUpper(arr[i-j])
					}

					arr = append(arr[:i], arr[i+2:]...)
					i--
				}

			} else {
				arr[i] = strings.ToUpper(strings.TrimSuffix(arr[i], arr[i][len(arr[i])-4:])) 
				for j := 1; j <= back-1; j++ {
					arr[i-j] = InitCap(arr[i-j])
				}

				arr = append(arr[:i+1], arr[i+2:]...) 
			}
		}
	}

	return arr
}
