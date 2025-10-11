package tasks

func invertMap[K comparable, V comparable](m map[K]V) map[V]K {
	inverted := make(map[V]K)
	for key, value := range m {
		inverted[value] = key
	}
	return inverted
}
