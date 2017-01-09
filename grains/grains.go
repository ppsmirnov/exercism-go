package grains

import "fmt"

const testVersion = 1

// Square return naumber of grains on a single square
func Square(n int) (uint64, error) {
	var result uint64 = 1

	if n < 1 || n > 64 {
		return 0, fmt.Errorf("Invalid input")
	}

	for i := 1; i < n; i++ {
		result *= 2
	}

	return result, nil
}

// Total returns total number of grains on a chess board
func Total() uint64 {
	var result uint64
	for i := 1; i < 65; i++ {
		square, _ := Square(i)
		result += square
	}
	return result
}
