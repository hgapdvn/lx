package lxmaps

// Filter creates a new map containing only the entries for which the predicate returns true.
// Returns nil if input is nil. The order of entries in the returned map is not guaranteed.
//
// Filter is functionally equivalent to [PickBy].
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	out := Filter(m, func(k string, v int) bool { return v > 1 })
//	// out: map[string]int{"b": 2, "c": 3}
func Filter[K comparable, V any](m map[K]V, predicate func(k K, v V) bool) map[K]V {
	return PickBy(m, predicate)
}
