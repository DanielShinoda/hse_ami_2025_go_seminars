package tasks

func invertMap[K comparable, V comparable](lolo map[K]V) map[V]K {
	res := make(map[V]K)
	for k, v := range lolo {
		res[v] = k
	}
	return res
}
