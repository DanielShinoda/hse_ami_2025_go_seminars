package tasks

// CountVowels подсчитывает количество гласных в строке
func CountVowels(s string) int {
	all := map[rune]bool{
		'a': true,
		'A': true,
		'e': true,
		'E': true,
		'i': true,
		'I': true,
		'o': true,
		'O': true,
		'u': true,
		'U': true,
	}
	total := 0
	for _, c := range s {
		if _, ok := all[c]; ok {
			total++
		}
	}

	return total
}
