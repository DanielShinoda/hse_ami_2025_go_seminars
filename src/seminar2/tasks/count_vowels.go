package tasks

// CountVowels подсчитывает количество гласных в строке
func CountVowels(s string) int {
	all := map[rune]struct{}{
		'a': {},
		'A': {},
		'e': {},
		'E': {},
		'i': {},
		'I': {},
		'o': {},
		'O': {},
		'u': {},
		'U': {},
	}
	total := 0
	for _, c := range s {
		if _, ok := all[c]; ok {
			total++
		}
	}

	return total
}
