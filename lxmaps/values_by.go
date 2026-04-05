package lxmaps

// ValuesBy returns a slice of values from the map m that satisfy the predicate.
// Returns nil if input is nil. Returns an empty slice if no elements match.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	out := ValuesBy(m, func(v int) bool { return v > 1 })
//	// out: []int{2, 3}
func ValuesBy[K comparable, V any](m map[K]V, predicate func(v V) bool) []V {
	if m == nil {
		return nil
	}
	out := make([]V, 0, len(m))
	for _, v := range m {
		if predicate(v) {
			out = append(out, v)
		}
	}
	return out
}
