package lxmaps

// Difference returns a new map containing entries from m1 whose keys do not exist in m2.
// This performs a set difference operation: m1 - m2.
//
// For nil m1, returns nil. If m1 is empty or all its keys exist in m2,
// returns an empty map.
//
// Example:
//
//	Difference(map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"b": 99, "c": 88, "d": 77})
//	// Returns: map[string]int{"a": 1}
func Difference[K comparable, V any](m1, m2 map[K]V) map[K]V {
	if m1 == nil {
		return nil
	}

	result := make(map[K]V)
	for k, v := range m1 {
		if _, exists := m2[k]; !exists {
			result[k] = v
		}
	}
	return result
}
