package lxmaps

// Get retrieves the value associated with the key from the map.
// Returns (value, true) if the key exists in the map.
// Returns (zero value, false) if the key does not exist or the map is nil.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2}
//	if value, ok := Get(m, "a"); ok {
//	    // value: 1
//	}
func Get[K comparable, V any](m map[K]V, key K) (V, bool) {
	v, ok := m[key]
	return v, ok
}
