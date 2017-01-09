package protein

const testVersion = 1

var codes = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromCodon returns protein for given codon
func FromCodon(codon string) string {
	return codes[codon]
}

// FromRNA returns slice of proteins for given rna string
func FromRNA(rna string) (result []string) {
	if len(rna)%3 != 0 {
		return result
	}

	for i := 0; i < len(rna); i += 3 {
		var codon string
		for j := 0; j < 3; j++ {
			codon += string(rna[i+j])
		}
		protein := FromCodon(codon)
		if protein == "STOP" {
			break
		}
		result = append(result, protein)
	}
	return result
}
