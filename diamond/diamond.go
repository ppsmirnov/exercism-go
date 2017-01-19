package diamond

import (
	"errors"
	"strings"
)

const testVersion = 1

// Gen returns diamond for a given char
func Gen(char byte) (result string, err error) {
	diff := char - 'A'
	if diff > 'Z'-'A' {
		return "", errors.New("Char is out of range")
	}
	var arr []string
	for i := int(diff); int(i) >= 0; i-- {
		str := charTimes(i, ' ') + string(char-byte(i)) + charTimes(int(diff-byte(i)), ' ')
		str += reverseString(str[:len(str)-1]) + "\n"
		arr = append(arr, str)
	}
	arr = append(arr, reverseArray(arr[:len(arr)-1])...)
	result = strings.Join(arr, "")
	return result, nil
}

func charTimes(n int, char byte) (result string) {
	for i := 0; i < n; i++ {
		result += string(char)
	}
	return result
}

func reverseString(str string) (result string) {
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}

func reverseArray(arr []string) (result []string) {
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}
	return result
}
