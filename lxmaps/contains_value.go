package lxmaps

// ContainsValue checks if the map contains the specified value.
// If the map is nil, it returns false.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	out := ContainsValue(m, 1)
//	// out: true
func ContainsValue[K comparable, V comparable](m map[K]V, in V) bool {
	for _, v := range m {
		if v == in {
			return true
		}
	}
	return false
}
