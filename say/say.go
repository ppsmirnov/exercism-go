package say

var dict = map[uint64]string{
	1:    "one",
	2:    "two",
	3:    "three",
	4:    "four",
	5:    "five",
	6:    "six",
	7:    "seven",
	8:    "eight",
	9:    "nine",
	10:   "ten",
	11:   "eleven",
	12:   "twelve",
	13:   "thirteen",
	14:   "fourteen",
	15:   "fifteen",
	16:   "sixteen",
	17:   "seventeen",
	18:   "eighteen",
	19:   "nineteen",
	20:   "twenty",
	30:   "thirty",
	40:   "forty",
	100:  "hundred",
	1000: "thousand",
	1e6:  "million",
	1e9:  "billion",
	1e12: "trillion",
}

// Say returns string representation of a given number
func Say(n uint64) string {
	if n < 20 {
		result := dict[n]
		return result
	}
	if n < 100 {
		result := dict[n/10*10]
		if len(Say(n%10)) > 0 {
			result += "-" + Say(n%10)
		}
		return result
	}
	var result string
	for i := 1; n >= uint64(i); i *= 10 {
		diff := n % uint64(i)
		result = dict[n/uint64(i)] + " " + dict[uint64(i)]
		if diff > 0 {
			result += " " + Say(n%uint64(i))
		}
	}
	return result
}
