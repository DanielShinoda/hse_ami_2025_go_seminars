package tasks

// ReverseString переворачивает строку
func ReverseString(g string) string {
	s := ""
	for i := len([]rune(g)) - 1; i >= 0; i-- {
		s += string([]rune(g)[i])
	}
	return s
}
