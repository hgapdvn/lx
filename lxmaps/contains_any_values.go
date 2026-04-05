package lxmaps

// ContainsAnyValues checks if any of the specified values are present in the map.
// If the map is nil, it returns false.
// If the values slice is empty, it returns false.
// The order of the values in the input slice is not guaranteed to be the same as in the original map.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	out := ContainsAnyValues(m, 1, 2, 3)
func ContainsAnyValues[K comparable, V comparable](m map[K]V, values ...V) bool {
	for _, value := range values {
		for _, v := range m {
			if v == value {
				return true
			}
		}
	}
	return false
}
