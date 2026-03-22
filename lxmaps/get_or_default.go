package lxmaps

// GetOrDefault retrieves the value associated with the key from the map.
// Returns the value if the key exists in the map.
// Returns defaultValue if the key does not exist or the map is nil.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2}
//	value := GetOrDefault(m, "c", 999)
//	// value: 999
func GetOrDefault[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if v, ok := m[key]; ok {
		return v
	}
	return defaultValue
}
