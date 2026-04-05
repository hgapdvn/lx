package lxmaps

// ContainsBy checks if any key-value pair in the map satisfies the given predicate.
// Returns true if any key-value pair in the map satisfies the predicate, false otherwise.
// For an empty map, returns false.
// Example:
//
//	ContainsBy(map[string]int{"a": 1, "b": 2}, func(k string, v int) bool {
//		return v > 1
//	}) // true
//	ContainsBy(map[string]int{"a": 1, "b": 2}, func(k string, v int) bool {
//		return v > 2
//	}) // false
func ContainsBy[K comparable, V any](m map[K]V, fn func(k K, v V) bool) bool {
	for k, v := range m {
		if fn(k, v) {
			return true
		}
	}
	return false
}
