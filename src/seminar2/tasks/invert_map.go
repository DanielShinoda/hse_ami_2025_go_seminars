package tasks

func invertMap[K comparable, V comparable](ipt map[K]V) map[V]K {
	inverted := make(map[V]K)

	for key, val := range ipt {
		inverted[val] = key
	}

	return inverted
}
