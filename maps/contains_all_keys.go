package lxmaps

// ContainsAllKeys checks if all the specified keys are present in the map.
// If the map is nil, it returns false.
// If the keys slice is empty, it returns true.
// The order of the keys in the input slice is not guaranteed to be the same as in the original map.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	out := ContainsAllKeys(m, "a", "b", "c")
func ContainsAllKeys[K comparable, V any](m map[K]V, keys ...K) bool {
	if len(keys) == 0 {
		return true
	}
	for _, key := range keys {
		if _, ok := m[key]; !ok {
			return false
		}
	}
	return true
}
