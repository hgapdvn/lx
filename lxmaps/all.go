package lxmaps

// All checks whether all entries in the map satisfy the given predicate.
// Returns true if the map is nil, empty, or all entries satisfy the predicate.
// Returns false if any entry does not satisfy the predicate.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := All(m, func(k string, v int) bool {
//		return v > 0
//	})
//	// result: true
func All[K comparable, V any](m map[K]V, fn func(K, V) bool) bool {
	for k, v := range m {
		if !fn(k, v) {
			return false
		}
	}
	return true
}
