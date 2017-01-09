package triangle

import (
	"math"
	"sort"
)

const testVersion = 3

const (
	NaT = "not a triangle"
	Equ = "equilateral"
	Iso = "isosceles"
	Sca = "scalene"
)

// Kind represents triangle type
type Kind string

// KindFromSides calculates triangle type
func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	sort.Float64s(sides)

	if sides[2] > sides[0]+sides[1] || !uncorrect(sides) {
		return NaT
	}

	if a == b && a == c {
		return Equ
	}

	if sides[0] != sides[1] && sides[1] != sides[2] && sides[0] != sides[2] {
		return Sca
	}

	return Iso
}

func uncorrect(sides []float64) bool {
	for _, side := range sides {
		if side <= 0 || math.IsNaN(side) || math.IsInf(side, 0) {
			return false
		}
	}
	return true
}
