package lxmaps

// Update updates the value for a key in the map using an update function.
// The update function receives the current value and a boolean indicating whether the key exists.
// If the key exists, the function receives (currentValue, true).
// If the key doesn't exist, the function receives (zero value, false).
// The map is modified in-place and also returned for chaining.
// If m is nil, a new map is created.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2}
//	Update(m, "a", func(v int, exists bool) int {
//		if exists {
//			return v + 10
//		}
//		return 100
//	})
//	// m: map[string]int{"a": 11, "b": 2}
//
//	Update(m, "c", func(v int, exists bool) int {
//		if exists {
//			return v + 10
//		}
//		return 100
//	})
//	// m: map[string]int{"a": 11, "b": 2, "c": 100}
func Update[K comparable, V any](m map[K]V, key K, fn func(V, bool) V) map[K]V {
	if m == nil {
		m = make(map[K]V)
	}

	val, exists := m[key]
	m[key] = fn(val, exists)
	return m
}
