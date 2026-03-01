package tasks

// CountVowels подсчитывает количество гласных в строке
func CountVowels(input string) int {
	var count int
	for _, char := range input {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
			count++
		}
	}
	return count
}
