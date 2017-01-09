package pythagorean

// Triplet contains numbers of triplet
type Triplet [3]int

func (t Triplet) sum() int {
	var result int
	for _, i := range t {
		result += i
	}
	return result
}

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	var result []Triplet
	for i := min; i <= max; i++ {
		for j := i; j <= max; j++ {
			for k := j; k <= max; k++ {
				if i*i+j*j == k*k {
					result = append(result, Triplet{i, j, k})
				}
			}
		}
	}
	return result
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	var result []Triplet
	for _, t := range Range(1, p) {
		if t.sum() == p {
			result = append(result, t)
		}
	}
	return result
}
