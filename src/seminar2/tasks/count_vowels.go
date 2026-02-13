package tasks

import (
	"strings"
)

const vowels = "aeiouAEIOU"

func CountVowels(s string) int {
	count := 0
	for _, i := range s {
		if strings.ContainsRune(vowels, i) {
			count++
		}
	}
	return count
}
