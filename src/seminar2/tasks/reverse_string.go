package tasks

// ReverseString переворачивает строку
func ReverseString(str string) string {
	s := ""

	for _, ch := range str {
		s = string(ch) + s
	}

	return s
}
