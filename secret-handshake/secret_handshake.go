package secret

import (
	"strconv"
	"strings"
)

const testVersion = 1

var codes = map[int]string{
	1: "wink",
	2: "double blink",
	3: "close your eyes",
	4: "jump",
}

// Handshake encodes integer value
func Handshake(in uint) []string {
	var result []string
	needReversing := true
	binary := strings.Split(strconv.FormatInt(int64(in), 2), "")
	for i := len(binary) - 1; i >= 0; i-- {
		if binary[len(binary)-i-1] == "1" && i <= 4 {
			if i == 4 {
				needReversing = false
				continue
			}
			result = append(result, codes[i+1])
		}
	}

	if needReversing {
		return reverse(result)
	}
	return result
}

func reverse(slice []string) []string {
	var result []string
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}
	return result
}
