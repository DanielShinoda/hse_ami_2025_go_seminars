package tasks

// ReverseString переворачивает строку
func ReverseString(source string) string {
	runes := []rune(source)
	for i := 0; i < len(runes)/2; i++ {
		tmp := runes[i]
		runes[i] = runes[len(runes)-i-1]
		runes[len(runes)-i-1] = tmp
	}
	result := ""
	for _, r := range runes {
		result += string(r)
	}
	return result
}
