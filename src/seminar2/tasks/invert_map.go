package tasks

func invertMap[K comparable, V comparable](begin map[K]V) map[V]K {
	result := make(map[V]K)

	for k, v := range begin {
		result[v] = k
	}
	return result
}
