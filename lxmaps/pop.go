package lxmaps

// Pop removes and returns the value for the given key from the map.
// Returns (value, true) if the key was found and removed, (zero value, false) otherwise.
// Also returns (zero value, false) if the map is nil.
// The map is modified in-place.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	val, found := Pop(m, "b")
//	// val: 2, found: true
//	// m is now map[string]int{"a": 1, "c": 3}
//
//	val, found := Pop(m, "d")
//	// val: 0, found: false
//	// m unchanged
//
//	var m2 map[string]int
//	val, found := Pop(m2, "a")
//	// val: 0, found: false (nil map)
func Pop[K comparable, V any](m map[K]V, key K) (V, bool) {
	val, ok := m[key]
	if ok {
		delete(m, key)
	}
	return val, ok
}
