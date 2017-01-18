package diamond

import "strings"

const testVersion = 1

func Gen(char byte) (result string, err error) {
	diff := char - 'A'
	var arr []string
	for i := int(diff); int(i) >= 0; i-- {
		str := charTimes(i, '1') + string(char-byte(i)) + charTimes(int(diff-byte(i)), '1')
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
