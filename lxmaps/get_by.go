package lxmaps

// GetBy returns the value of the first map entry for which predicate returns true.
// If the map is nil or empty, or no entry matches, it returns the zero value of V and false.
// Map iteration order is not defined; when multiple entries match, any one of their values may be returned.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2}
//	v, ok := GetBy(m, func(k string, v int) bool { return k == "a" })
//	// v: 1, ok: true
func GetBy[K comparable, V any](m map[K]V, predicate func(k K, v V) bool) (V, bool) {
	var zero V
	if len(m) == 0 {
		return zero, false
	}
	for k, v := range m {
		if predicate(k, v) {
			return v, true
		}
	}
	return zero, false
}
