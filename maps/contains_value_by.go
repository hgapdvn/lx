package lxmaps

// ContainsValueBy checks if the map contains a value that satisfies the predicate.
// If the map is nil, it returns false.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	out := ContainsValueBy(m, func(v int) bool { return v > 1 })
//	// out: true
func ContainsValueBy[K comparable, V any](m map[K]V, predicate func(v V) bool) bool {
	for _, v := range m {
		if predicate(v) {
			return true
		}
	}
	return false
}
