package tasks

// FilterNumbers фильтрует числа по условию
func FilterNumbers(numbers []int, predicate func(int) bool) []int {
	result := []int{}
	for _, it := range numbers {
		if predicate(it) {
			result = append(result, it)
		}
	}
	return result
}
