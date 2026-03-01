package tasks

func invertMap[K comparable, V comparable](d map[K]V) map[V]K {
	ans := make(map[V]K)
	for k,v :=range d {
		ans[v]=k
	}
	return ans

}
