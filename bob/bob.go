package bob

import (
	"regexp"
	"strings"
)

const testVersion = 2

// Hey returns Bos's answer
func Hey(in string) string {
	reg, _ := regexp.Compile(`[^\p{L}]`)
	onlyLetters := reg.ReplaceAllString(in, "")
	in = strings.TrimSpace(in)

	if strings.ToUpper(in) == in && len(onlyLetters) != 0 {
		return "Whoa, chill out!"
	}

	if strings.HasSuffix(in, "?") {
		return "Sure."
	}

	if len(in) == 0 {
		return "Fine. Be that way!"
	}

	return "Whatever."
}
