package wordcount

import (
	"regexp"
	"strings"
)

const testVersion = 3

// Frequency type
type Frequency map[string]int

// WordCount returns count of every word in a given phrase
func WordCount(phrase string) Frequency {
	f := make(Frequency)
	reg, _ := regexp.Compile(`[\n@".,\/#!$%\^&\*;:{}=\-_()]`)
	phrase = strings.ToLower(phrase)
	phrase = reg.ReplaceAllString(phrase, " ")

	for _, word := range strings.Split(phrase, " ") {
		w := strings.Trim(word, " ")
		w = strings.Trim(w, "'")
		if len(w) > 0 {
			f[w]++
		}
	}
	return f
}
