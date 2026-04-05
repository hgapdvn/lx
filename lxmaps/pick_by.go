package lxmaps

// PickBy returns a new map containing only the entries for which the predicate returns true.
// If the input map is nil, PickBy returns nil.
// The order of the entries in the returned map is not guaranteed to be the same as in the original map.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	out := PickBy(m, func(k string, v int) bool { return v > 1 })
//	// out: map[string]int{"b": 2, "c": 3}
func PickBy[K comparable, V any](m map[K]V, predicate func(K, V) bool) map[K]V {
	if m == nil {
		return nil
	}
	out := make(map[K]V)
	for k, v := range m {
		if predicate(k, v) {
			out[k] = v
		}
	}
	return out
}
