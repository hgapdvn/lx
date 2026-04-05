package lxmaps

// MergeBy merges two maps into a single map using a custom merge function for conflicts.
// When a key exists in both maps, the merge function is called with the value from m1 and m2.
// The result of the merge function becomes the new value in the result map.
// If m1 is nil, an empty map is created. If m2 is nil, returns a clone of m1.
//
// Example:
//
//	m1 := map[string]int{"a": 1, "b": 2}
//	m2 := map[string]int{"b": 3, "c": 4}
//	result := MergeBy(m1, m2, func(existing, new int) int {
//		return existing + new
//	})
//	// result: map[string]int{"a": 1, "b": 5, "c": 4}
func MergeBy[K comparable, V any](m1, m2 map[K]V, fn func(V, V) V) map[K]V {
	var out map[K]V
	if m1 == nil {
		out = make(map[K]V)
	} else {
		out = Clone(m1)
	}
	for k, v := range m2 {
		if existing, ok := out[k]; ok {
			out[k] = fn(existing, v)
		} else {
			out[k] = v
		}
	}
	return out
}
