package tasks

func invertMap[K comparable, V comparable](input map[K]V) map[V]K {
	a := make(map[V]K)
	for k, v := range input {
		a[v] = k
	}
	return a
}
