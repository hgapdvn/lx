package lxmaps

// ContainsAllValues returns true if all of the values are present in the map.
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
