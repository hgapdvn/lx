package lxmaps

// Intersect returns a new map containing only the entries from m1 whose keys also exist in m2.
// The values from m1 are preserved in the result.
//
// For nil input, returns nil. For empty maps, the result will contain only
// the keys that exist in both maps.
//
// Example:
//
//	Intersect(map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"b": 99, "c": 88, "d": 77})
//	// Returns: map[string]int{"b": 2, "c": 3}
func Intersect[K comparable, V any](m1, m2 map[K]V) map[K]V {
	if m1 == nil || m2 == nil {
		return nil
	}

	result := make(map[K]V)
	for k, v := range m1 {
		if _, exists := m2[k]; exists {
			result[k] = v
		}
	}
	return result
}
