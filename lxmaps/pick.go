package lxmaps

// Pick returns a new map containing only keys present in m among those listed.
// Keys in keys that are not in m are ignored (no zero-value entries are added).
// For nil m, returns nil. For an empty non-nil m, returns a new empty map.
// When keys is empty, returns a new empty non-nil map.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2}
//	Pick(m, "a", "c") // map[string]int{"a": 1}
func Pick[K comparable, V any](m map[K]V, keys ...K) map[K]V {
	if m == nil {
		return nil
	}
	if len(keys) == 0 {
		return make(map[K]V)
	}
	out := make(map[K]V, len(keys))
	for _, k := range keys {
		if v, ok := m[k]; ok {
			out[k] = v
		}
	}
	return out
}
