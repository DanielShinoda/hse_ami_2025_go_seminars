package tasks

// ReverseString переворачивает строку
func ReverseString(s string) string {
	result := ""
	for _, char := range s {
		result = string(char) + result
	}
	return result
}

