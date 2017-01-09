package foodchain

import "fmt"

const testVersion = 3

var items = []string{"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}
var lines = []string{
	"",
	"It wriggled and jiggled and tickled inside her.",
	"How absurd to swallow a bird!",
	"Imagine that, to swallow a cat!",
	"What a hog, to swallow a dog!",
	"Just opened her throat and swallowed a goat!",
	"I don't know how she swallowed a cow!",
	"She's dead, of course!",
}

// Verse returns one song verse
func Verse(n int) string {
	result := fmt.Sprintf("I know an old lady who swallowed a %s.\n", items[n-1])
	result += lines[n-1]
	if n > 7 {
		return result
	}
	if n > 1 {
		result += "\n"
	}
	for i := n; i >= 2; i-- {
		var line string
		if items[i-2] == "spider" {
			line = fmt.Sprintf("She swallowed the %s to catch the %s that wriggled and jiggled and tickled inside her.\n", items[i-1], items[i-2])
		} else {
			line = fmt.Sprintf("She swallowed the %s to catch the %s.\n", items[i-1], items[i-2])

		}
		result += line
	}
	return result + "I don't know why she swallowed the fly. Perhaps she'll die."
}

// Verses returns verses from start to ende
func Verses(start, end int) string {
	var result string
	for i := start; i <= end; i++ {
		result += Verse(i)
		if i != end {
			result += "\n\n"
		}
	}
	return result
}

// Song returns all song
func Song() string {
	return Verses(1, 8)
}
