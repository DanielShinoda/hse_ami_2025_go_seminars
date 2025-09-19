package tasks

import "slices"

// FilterNumbers фильтрует числа по условию
func FilterNumbers(numbers []int, predicate func(int) bool) []int {
	numbers = slices.DeleteFunc(numbers, func(el int) bool { return !predicate(el) })
	return numbers
}
