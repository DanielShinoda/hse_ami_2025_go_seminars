package tasks

// CountVowels подсчитывает количество гласных в строке
func CountVowels(s string) int {
	cnt := 0
	for _, value := range s {
		if value == 'a' || value == 'e' || value == 'i' || value == 'o' || value == 'u' || value == 'y' ||
			value == 'A' || value == 'E' || value == 'I' || value == 'O' || value == 'U' || value == 'Y' {

			cnt++
		}
	}
	return cnt
}
