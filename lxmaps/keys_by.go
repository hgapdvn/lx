package lxmaps

// KeysBy returns a slice of keys of the map m for which the predicate returns true.
// The order of the keys in the returned slice is not guaranteed to be the same as in the original map.
// If the map is nil, it returns an empty slice.
// // Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	greaterThan := func(k string, v int) bool { return v > 1 }
//	keys := KeysBy(m, greaterThan) // keys will contain ["b", "c"]
func KeysBy[K comparable, V any](m map[K]V, predicate func(K, V) bool) []K {
	if m == nil {
		return []K{}
	}
	keys := make([]K, 0, len(m))
	for k, v := range m {
		if predicate(k, v) {
			keys = append(keys, k)
		}
	}
	return keys
}
