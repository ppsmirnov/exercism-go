package acronym

import (
	"regexp"
	"strings"
	"unicode"
)

const testVersion = 1

func abbreviate(str string) string {
	var result string
	reg, _ := regexp.Compile("[^A-Za-z]+")
	str = reg.ReplaceAllString(str, " ")
	for i, r := range str {
		if !unicode.IsSpace([]rune(str)[i]) &&
			(unicode.IsUpper(r) && !unicode.IsUpper([]rune(str)[i+1])) ||
			(i > 0 && unicode.IsSpace([]rune(str)[i-1])) {
			result += strings.ToUpper(string(r))
		}
	}
	return result
}
