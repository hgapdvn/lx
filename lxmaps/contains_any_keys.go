package lxmaps

// ContainsAnyKeys checks if any of the specified keys are present in the map.
// If the map is nil, it returns false.
// If the keys slice is empty, it returns false.
// The order of the keys in the input slice is not guaranteed to be the same as in the original map.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	out := ContainsAnyKeys(m, "a", "b", "c")
func ContainsAnyKeys[K comparable, V any](m map[K]V, keys ...K) bool {
	for _, key := range keys {
		if _, ok := m[key]; ok {
			return true
		}
	}
	return false
}
