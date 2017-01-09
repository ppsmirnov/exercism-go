package etl

import "strings"

// Transform returns new scrabble score system
func Transform(old map[int][]string) map[string]int {
	result := make(map[string]int)
	for k, v := range old {
		for _, str := range v {
			result[strings.ToLower(str)] = k
		}
	}
	return result
}
