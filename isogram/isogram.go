package isogram

import (
	"regexp"
	"strings"
)

const testVersion = 1

// IsIsogram returns true if string is isogram
func IsIsogram(str string) bool {
	var dict = make(map[rune]int)
	reg, _ := regexp.Compile(`[^\p{L}]`)
	onlyLetters := reg.ReplaceAllString(str, "")
	for _, r := range strings.ToLower(onlyLetters) {
		if dict[r] > 0 {
			return false
		}
		dict[r]++
	}
	return true
}
