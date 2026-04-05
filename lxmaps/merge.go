package lxmaps

// Merge merges multiple maps into a single map.
// If the maps are nil, it returns an empty map.
// If the maps are empty, it returns an empty map.
// If the maps have duplicate keys, the value from the last map will be used.
// The order of the maps in the input slice is not guaranteed to be the same as in the original maps.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	out := Merge(m, m)
//	// out: map[string]int{"a": 1, "b": 2, "c": 3}
func Merge[K comparable, V any](in ...map[K]V) map[K]V {
	out := make(map[K]V)
	for _, m := range in {
		for k, v := range m {
			out[k] = v
		}
	}
	return out
}
