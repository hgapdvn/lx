package lxmaps

// ContainsKey checks if the map contains the specified key.
// If the map is nil, it returns false.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	out := ContainsKey(m, "a")
//	// out: true
//
// The value associated with the key is not checked, only the key's presence.
func ContainsKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}
