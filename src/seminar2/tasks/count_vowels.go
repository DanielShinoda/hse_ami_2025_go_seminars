package tasks

// CountVowels подсчитывает количество гласных в строке
func CountVowels(s string) int {
	count := 0
	for _, r := range s {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}
	return count
}
