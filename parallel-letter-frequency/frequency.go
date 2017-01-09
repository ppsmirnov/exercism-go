package letter

import "fmt"

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(arr []string) FreqMap {
	m := FreqMap{}

	ch := make(chan FreqMap)

	for _, str := range arr {
		go func(s string) {
			ch <- Frequency(s)
		}(str)
	}

	for m2 := range ch {
		fmt.Println(m2)
		// for r, v := range m2 {
		// 	m[r] += v
		// }
	}
	return m
}
