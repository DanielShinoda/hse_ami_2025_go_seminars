package tasks

func invertMap(a map[string]int) map[int]string {
	res := make(map[int]string)
	for key, value := range a {
		res[value] = key

	}
	return res
}
