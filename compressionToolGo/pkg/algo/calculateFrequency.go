package algo

import (
	"unicode"
)

func CalculateTheFrequency(lines []byte) map[rune]int {
	m := make(map[rune]int)

	for _, line := range lines {
		runes := []rune(string(line))
		for i:=0;i< len(runes);i++{
		r:= runes[i]
		if ! unicode.IsSpace(r){
 			_, err := m[r]
			if !err{
				m[r]=0
			}
			m[r]++;
		}
	}
	}

	return m
}