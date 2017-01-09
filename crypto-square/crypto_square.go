package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

const testVersion = 2

func Encode(str string) string {
	normalized := strings.ToLower(reg.ReplaceAllString(str, ""))
	reg, _ := regexp.Compile(`[^\p{L}\d]`)
	l := int(math.Ceil(math.Sqrt(float64(len(normalized)))))
	cipherText := make([]string, l)
	for i := 0; i < l; i++ {
		for j := i; j < len(normalized); j += l {
			cipherText[i] += string(normalized[j])
		}
	}
	return strings.Join(cipherText, " ")
}
