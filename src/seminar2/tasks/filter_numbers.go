package tasks

func FilterNumbers(numbers []int, predicate func(int) bool) []int {
	res := []int{}
	for _, n := range numbers {
		if predicate(n) {
			res = append(res, n)
		}
	}
	return res
}
