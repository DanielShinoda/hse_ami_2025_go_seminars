package tasks

// ReverseString переворачивает строку
func ReverseString(s string) string {
	runes := []rune(s)
	ans := ""
	for i := len(runes) - 1; i >= 0; i-- {
		ans += string(runes[i])
	}
	return ans
}
