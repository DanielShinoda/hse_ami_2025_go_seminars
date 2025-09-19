package tasks

import "strings"

func isVowel(r rune) bool {
	switch strings.ToLower(string(r)) {
	case "a", "e", "i", "o", "u":
		return true
	default:
		return false
	}
}

// CountVowels подсчитывает количество гласных в строке
func CountVowels(str string) int {
	count := 0
	for _, r := range str {
		if isVowel(r) {
			count++
		}
	}
	return count
}
