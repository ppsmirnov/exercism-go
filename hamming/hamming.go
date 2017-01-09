package hamming

import "errors"

const testVersion = 5

func Distance(a, b string) (int, error) {
	var diff int
	if len(a) != len(b) {
		return -1, errors.New("strings: missmatched length")
	}
	for i, r := range a {
		if r != []rune(b)[i] {
			diff++
		}
	}
	return diff, nil
}
