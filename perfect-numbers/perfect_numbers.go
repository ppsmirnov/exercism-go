package perfect

import "errors"

const testVersion = 1

// Classification type
type Classification string

// ClassificationDeficient is deficient classification
const ClassificationDeficient Classification = "ClassificationDeficient"

// ClassificationPerfect is perfect classification
const ClassificationPerfect Classification = "ClassificationPerfect"

// ClassificationAbundant is abundant classification
const ClassificationAbundant Classification = "ClassificationAbundant"

// ErrOnlyPositive is an error for not positive number
var ErrOnlyPositive = errors.New("Not a positve number")

// Classify returns classification for a given number
func Classify(n uint64) (class Classification, err error) {
	if n <= 0 {
		return "", ErrOnlyPositive
	}
	factors := findFacotrs(n)
	var sum uint64
	for _, factor := range factors {
		sum += uint64(factor)
	}

	if sum < n {
		return ClassificationDeficient, nil
	}

	if sum > n {
		return ClassificationAbundant, nil
	}

	return ClassificationPerfect, nil
}

func findFacotrs(n uint64) (factors []int) {
	for i := 1; uint64(i) <= n/2; i++ {
		if int(n)%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}
