package pascal

// Triangle return Pascal's tringle for n li
func Triangle(n int) [][]int {
	var result [][]int
	for i := 0; i < n; i++ {
		result = append(result, row(i))
	}
	return result
}

func row(n int) []int {
	var result []int
	for i := 0; i <= n; i++ {
		result = append(result, number(n, i))
	}
	return result
}

func fact(n int) int {
	if n < 2 {
		return 1
	}
	n *= fact(n - 1)
	return n
}

func number(n, r int) int {
	return fact(n) / (fact(r) * fact(n-r))
}
