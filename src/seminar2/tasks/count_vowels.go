package tasks

func CountVowels(s1 string) int {
	s2 := "aeiouyAEIOUY"
	count := 0
	for _, ch1 := range s1 {
		for _, ch2 := range s2 {
			if ch1 == ch2 {
				count++
			}
		} 
	}
	return count
}
