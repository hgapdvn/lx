package lxmaps

// KeyBy transforms the keys of the input map using the provided function fn.
//
// Behavior:
//   - If m is nil, KeyBy returns an empty (non-nil) map (consistent with ValueBy behavior).
//   - The returned map has the same values as the input map, with each key replaced by fn(k, v).
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2}
//	out := KeyBy(m, func(k string, v int) string { return k + "_x" })
//	// out: map[string]int{"a_x": 1, "b_x": 2}
func KeyBy[K comparable, J comparable, V any](m map[K]V, fn func(K, V) J) map[J]V {
	out := make(map[J]V, len(m))
	for k, v := range m {
		out[fn(k, v)] = v
	}
	return out
}
