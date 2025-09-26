package tasks

// FilterNumbers фильтрует числа по условию
func FilterNumbers(numbers []int, predicate func(int) bool) []int {
	var result []int
	for i := 0; i < len(numbers); i++ {
		if predicate(numbers[i]) {
			result = append(result, numbers[i])
		}
	}
	return result
}
