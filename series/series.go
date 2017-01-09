package slice

// All returns a slice of n-sized groupes of string
func All(n int, str string) []string {
	var result []string
	for i := 0; i < len(str)-n+1; i++ {
		result = append(result, str[i:i+n])
	}
	return result
}

// UnsafeFirst return first n-1 characters of a string
func UnsafeFirst(n int, str string) string {
	var result string
	result = str[:n]
	return result
}
