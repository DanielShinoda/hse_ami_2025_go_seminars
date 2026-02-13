package tasks

// FilterNumbers фильтрует числа по условию
func FilterNumbers(numbers []int, predicate func(int) bool) []int {
	var ans []int
	for _, value := range numbers {
		if predicate(value) {
			ans = append(ans, value)
		}
	}
	return ans
}
