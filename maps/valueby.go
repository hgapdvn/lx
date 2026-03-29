package lxmaps

// ValueBy transforms the values of the input map using the provided function fn.
//
// Behavior:
//   - If m is nil, ValueBy returns an empty (non-nil) map (consistent with KeyBy behavior).
//   - The returned map has the same keys as the input map, with each value replaced by fn(k, v).
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2}
//	out := ValueBy(m, func(k string, v int) string { return fmt.Sprintf("%s:%d", k, v) })
//	// out: map[string]string{"a":"a:1", "b":"b:2"}
func ValueBy[K comparable, V any, U any](m map[K]V, fn func(K, V) U) map[K]U {
	out := make(map[K]U, len(m))
	if m == nil {
		// return empty non-nil map to match KeyBy behavior
		return out
	}
	for k, v := range m {
		out[k] = fn(k, v)
	}
	return out
}
