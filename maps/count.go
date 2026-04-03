package lxmaps

// Count returns the number of entries in the map that satisfy the predicate.
// If the map is nil, returns 0.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	count := Count(m, func(k string, v int) bool { return v > 1 })
//	// count: 2 (entries "b" and "c")
func Count[K comparable, V any](m map[K]V, predicate func(K, V) bool) int {
	count := 0
	for k, v := range m {
		if predicate(k, v) {
			count++
		}
	}
	return count
}
