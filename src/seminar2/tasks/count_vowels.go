package tasks

import "unicode"

// CountVowels подсчитывает количество гласных в строке

func CountVowels(s string) int {
	cnt := 0
	for _, i := range s {
		lowreg := unicode.ToLower(i)
		switch lowreg {
		case 'a', 'e', 'i', 'o', 'u':
			cnt++
		}
	}
	return cnt
}
