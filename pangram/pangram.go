package pangram

import (
	"regexp"
	"strings"
)

const testVersion = 1

// IsPangram computes weather input string is pangram
func IsPangram(str string) bool {
	letterSet := make(map[rune]bool)
	reg, _ := regexp.Compile("[^A-za-z]+")

	str = strings.ToLower(reg.ReplaceAllString(str, ""))
	for _, r := range str {
		letterSet[r] = true
	}

	return len(letterSet) == 26
}
