package binary

import (
	"errors"
	"math"
)

const testVersion = 2

// ParseBinary returns decimal integer for a given binary string
func ParseBinary(binary string) (result int, err error) {
	var power float64
	for i := len(binary) - 1; i >= 0; i-- {
		diff := binary[i] - '0'
		if diff > 1 || diff < 0 {
			return result, errors.New("Bad input")
		}
		result += int(diff) * int(math.Pow(2, power))
		power++
	}
	return result, err
}
