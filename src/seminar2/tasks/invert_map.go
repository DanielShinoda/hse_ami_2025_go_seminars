package tasks

func invertMap[K comparable, V comparable](m map[K]V) map[V]K {
	res := make(map[V]K)
	for k, v := range m {
		res[v] = k
	}
	return res
}
