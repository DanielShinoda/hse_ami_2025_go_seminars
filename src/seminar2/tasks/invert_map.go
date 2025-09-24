package tasks

func invertMap[K comparable, V comparable](m map[K]V) map[V]K {
	invt := make(map[V]K)
	for k, val := range m {
		invt[val] = k
	}
	return invt
}