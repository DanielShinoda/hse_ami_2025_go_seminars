package tasks

func invertMap[K comparable, V comparable](ourMap map[K]V) map[V]K {
	new_map := make(map[V]K)
	for key, value := range ourMap {
		new_map[value] = key
	}
	return new_map
}
