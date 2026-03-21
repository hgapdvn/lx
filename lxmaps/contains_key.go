package lxmaps

// ContainsKey returns true if the map contains the key.
// The value associated with the key is not checked, only the key's presence.
func ContainsKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}
