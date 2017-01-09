package diffsquares

import "math"

//SquareOfSums return square of sums
func SquareOfSums(n int) int {
	var result int
	for i := 1; i <= n; i++ {
		result += i
	}
	return int(math.Pow(float64(result), 2.0))
}

// SumOfSquares returns sum of squares
func SumOfSquares(n int) int {
	var result int
	for i := 1; i <= n; i++ {
		result += int(math.Pow(float64(i), 2.0))
	}
	return result
}

// Difference returns SquareOfSums - SumOfSquares
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
