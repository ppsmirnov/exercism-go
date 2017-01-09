package palindrome

import (
	"fmt"
	"strconv"
)

const testVersion = 1

// Product struct
type Product struct {
	Product        int
	Factorizations [][2]int
}

type products []Product

func (ps products) findProductWithValue(n int) (*Product, error) {
	for i := range ps {
		if ps[i].Product == n {
			return &ps[i], nil
		}
	}
	return &Product{}, fmt.Errorf("Not found")
}

// Products retruns min and max products for input range
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		return Product{}, Product{}, fmt.Errorf("fmin > fmax")
	}
	var ps products
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			try := i * j
			if isPalindrome(try) {
				var product *Product
				product, err := ps.findProductWithValue(try)
				if err == nil {
					product.Factorizations = append(product.Factorizations, [2]int{i, j})
				} else {
					var list [][2]int
					list = append(list, [2]int{i, j})
					product = &Product{try, list}
					ps = append(ps, *product)
				}
				if try <= pmin.Product || pmin.Product == 0 {
					pmin = *product
				}
				if try >= pmax.Product || pmax.Product == 0 {
					pmax = *product
				}
			}
		}
	}
	if len(ps) < 1 {
		return Product{}, Product{}, fmt.Errorf("No palindromes")
	}
	return pmin, pmax, nil
}

func isPalindrome(n int) bool {
	return strconv.Itoa(n) == reverseString(strconv.Itoa(n))
}

func reverseString(str string) string {
	var result string
	for _, r := range str {
		result = string(r) + result
	}
	return result
}
