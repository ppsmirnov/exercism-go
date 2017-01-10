package allergies

import "fmt"

const testVersion = 1

var allergensTable = map[int]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

// AllergicTo returns whether person with given score allergic to given allergen
func AllergicTo(score uint, subject string) bool {
	for i := 1; i <= int(score); i *= 2 {
		if allergensTable[i] == subject {
			return true
		}
	}
	return false
}

// Allergies returns all allergens for given score
func Allergies(score uint) (allergens []string) {
	binary := fmt.Sprintf("%b", score)
	j := 1
	for i := len(binary) - 1; i >= 0; i-- {
		code := int(binary[i]-'0') * j
		allergen := allergensTable[int(code)]
		if len(allergen) > 0 {
			allergens = append(allergens, allergen)
		}
		j *= 2
	}
	return allergens
}
