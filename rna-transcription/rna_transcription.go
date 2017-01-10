package strand

const testVersion = 3

var dict = map[rune]rune{
	'C': 'G',
	'G': 'C',
	'T': 'A',
	'A': 'U',
}

// ToRNA translates DNA string to RNA one
func ToRNA(DNA string) (RNA string) {
	for _, r := range DNA {
		RNA += string(dict[r])
	}
	return RNA
}
