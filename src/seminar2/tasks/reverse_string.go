package tasks

// ReverseString переворачивает строку
func ReverseString(s string) string {
	result := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		result = append(result, s[len(s)-i-1])
	}
	return string(result)
}
