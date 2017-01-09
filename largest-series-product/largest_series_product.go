package lsproduct

import "fmt"

const testVersion = 4

// LargestSeriesProduct returns  largest product for a contiguous substring of digits of length span
func LargestSeriesProduct(digits string, span int) (int, error) {
	var result int
	if span > len(digits) || span < 0 {
		return -1, fmt.Errorf("Inccorect input")
	}
	for i := 0; i < len(digits)-span+1; i++ {
		product := 1
		for j := 0; j < span; j++ {
			digit := digits[i+j] - '0'
			if digit > 9 {
				return -1, fmt.Errorf("Input contains not a digit")
			}
			product *= int(digit)
		}
		if product > result {
			result = product
		}
	}
	return result, nil
}
