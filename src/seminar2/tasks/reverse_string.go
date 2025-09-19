package tasks

// ReverseString переворачивает строку
func ReverseString(input string) string {
	slice := []rune(input)
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return string(slice)
}
