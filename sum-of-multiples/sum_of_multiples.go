package summultiples

// MultipleSummer returns sum of all the multiples o
func MultipleSummer(multiples ...int) func(limit int) int {
	return func(limit int) int {
		if len(multiples) == 0 || limit < 1 {
			return 0
		}
		ch := make(chan []int)
		for _, m := range multiples {
			go func(m, limit int) {
				ch <- findMultiples(m, limit)
			}(m, limit)
		}
		var concatedSlices []int
		for _ = range multiples {
			concatedSlices = concatSlices(concatedSlices, <-ch)
		}
		result := sumOfSlice(concatedSlices)
		return result
	}
}

func findMultiples(n, limit int) []int {
	var result []int

	for i := n; i < limit; i += n {
		result = append(result, i)
	}
	return result
}

func sumOfSlice(slice []int) int {
	var result int
	for _, i := range slice {
		result += i
	}
	return result
}

func concatSlices(slices ...[]int) []int {
	result := slices[0]
	for i := 1; i < len(slices); i++ {
		for _, el := range slices[i] {
			if pos(result, el) == -1 {
				result = append(result, el)
			}
		}
	}
	return result
}

func pos(slice []int, value int) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}
