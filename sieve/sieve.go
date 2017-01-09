package sieve

func Sieve(n int) []int {
	var result []int
	arr := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if !arr[i-1] {
			result = append(result, i)
			for j := i; i*j <= n; j++ {
				arr[i*j-1] = true
			}
		}
	}
	return result
}
