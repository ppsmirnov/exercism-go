package anagram

import (
	"math/rand"
	"strings"
)

// Detect retunrns a slice of anagrams for given subject and candidates
func Detect(subject string, candidates []string) (anograms []string) {
	subject = strings.ToLower(subject)
	sortedSubject := sortSlice([]byte(subject))
	for _, c := range candidates {
		c := strings.ToLower(c)
		if subject == strings.ToLower(c) {
			continue
		}

		if string(sortedSubject) == string(sortSlice([]byte(c))) {
			anograms = append(anograms, c)
		}
	}
	return anograms
}

func sortSlice(slice []byte) (sorted []byte) {
	if len(slice) < 2 {
		return slice
	}

	pivotID := rand.Intn(len(slice))

	var less, more []byte
	for i, b := range slice {
		if i == pivotID {
			continue
		}
		if b < slice[pivotID] {
			less = append(less, b)
		} else {
			more = append(more, b)
		}
	}

	sorted = append(sortSlice(less), slice[pivotID])
	sorted = append(sorted, sortSlice(more)...)
	return sorted
}
