package raindrops

import "strconv"

const testVersion = 2

func Convert(num int) string {
	var result string
	if num%3 == 0 {
		result += "Pling"
	}
	if num%5 == 0 {
		result += "Plang"
	}
	if num%7 == 0 {
		result += "Plong"
	}

	if len(result) == 0 {
		result = strconv.Itoa(num)
	}

	return result
}
