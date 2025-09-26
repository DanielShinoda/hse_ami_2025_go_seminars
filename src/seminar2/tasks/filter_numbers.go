package tasks

// FilterNumbers фильтрует числа по условию
func FilterNumbers(numbers []int, predicate func(int) bool) []int {
	var res []int
	for _, num := range numbers {
		if predicate(num) {
			res = append(res, num)
		}
	}
	return res
}
