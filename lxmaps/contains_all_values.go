package lxmaps

// ContainsAllValues checks if all the specified values are present in the map.
// If the map is nil, it returns false.
// If the values slice is empty, it returns true.
// The order of the values in the input slice is not guaranteed to be the same as in the original map.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	out := ContainsAllValues(m, 1, 2, 3)
//	// out: true
func ContainsAllValues[K comparable, V comparable](m map[K]V, values ...V) bool {
	if len(values) == 0 {
		return true
	}
	for _, value := range values {
		found := false
		for _, v := range m {
			if v == value {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
