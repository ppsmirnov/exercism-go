package luhn

import (
	"regexp"
	"strconv"
)

func checkSum(str string) int {
	var sum int
	var arr []int
	reg, _ := regexp.Compile(`\D`)
	onlyDigits := reg.ReplaceAllString(str, "")
	if len(onlyDigits) == 0 {
		return -1
	}
	for i, j := len(onlyDigits), 0; i > 0; i-- {
		if j%2 == 0 {
			sum += int(onlyDigits[i-1] - '0')
			arr = append(arr, int(onlyDigits[i-1]-'0'))
		} else {
			newValue := int(onlyDigits[i-1]-'0') * 2
			if newValue >= 10 {
				newValue = (newValue - 9)
			}
			sum += newValue
			arr = append(arr, newValue)
		}
		j++
	}
	return sum
}

// Valid checks if string is Luhn valid
func Valid(str string) bool {
	return checkSum(str)%10 == 0
}

// AddCheck returns correct strings
func AddCheck(str string) string {
	for i := 0; i <= 9; i++ {
		try := str + strconv.Itoa(i)
		if Valid(try) {
			return try
		}
	}
	return str
}
