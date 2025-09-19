package tasks

// CountVowels подсчитывает количество гласных в строке
func CountVowels(str string) int {
	c := 0
	for _, ch := range str {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			c++
		}
	}

	return c
}
