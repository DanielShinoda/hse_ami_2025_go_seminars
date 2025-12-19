package tasks

import (
	"fmt"
	"sort"
)

func invertMap[K comparable, V comparable](source map[K]V) map[V]K {
	result := make(map[V]K, len(source))
	keys := make([]K, 0, len(source))

	for k := range source {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return fmt.Sprintf("%v", keys[i]) < fmt.Sprintf("%v", keys[j])
	})

	for _, k := range keys {
		result[source[k]] = k
	}

	return result
}
