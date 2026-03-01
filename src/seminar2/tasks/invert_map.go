package tasks

func invertMap[K comparable, V comparable](source map[K]V) map[V]K {
	inverted := make(map[V]K)
	for key, value := range source {
		inverted[value] = key
	}
	return inverted
}
