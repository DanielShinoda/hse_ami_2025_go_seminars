package tasks

// ReverseString переворачивает строку
func ReverseString(s string) string {
	result := make([]rune, 0, len(s))
	for i := 0; i < len(s); i++ {
		result = append(result, rune(s[len(s)-i-1]))
	}
	return string(result)
}
