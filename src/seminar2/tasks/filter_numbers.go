package tasks

// FilterNumbers фильтрует числа по условию
func FilterNumbers(numbers []int, predicate func(int) bool) []int {
	new_numbers := make([]int, 0, len(numbers))
	for _, num := range numbers {
		if predicate(num) {
			new_numbers = append(new_numbers, num)
		}
	}
	return new_numbers
}
