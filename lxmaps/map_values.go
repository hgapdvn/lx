package lxmaps

// MapValues transforms the values of a map using a transformation function.
// Returns a new map with original keys and transformed values.
// If the input map is nil, returns nil.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := MapValues(m, func(v int) string {
//		return fmt.Sprintf("num_%d", v)
//	})
//	// result: map[string]string{"a": "num_1", "b": "num_2", "c": "num_3"}
func MapValues[K comparable, V, U any](m map[K]V, fn func(V) U) map[K]U {
	if m == nil {
		return nil
	}

	result := make(map[K]U, len(m))
	for k, v := range m {
		result[k] = fn(v)
	}
	return result
}
