package lxmaps

// IsSubset checks if subset is a subset of superset.
// A map is a subset of another if all key-value pairs in subset exist in superset with the same values.
// Returns true if subset is nil or empty, or if all pairs match.
// Returns false if superset is nil or if any pair in subset doesn't match superset.
//
// Example:
//
//	subset := map[string]int{"a": 1, "b": 2}
//	superset := map[string]int{"a": 1, "b": 2, "c": 3}
//	IsSubset(subset, superset)
//	// true
//
//	subset2 := map[string]int{"a": 1, "b": 99}
//	IsSubset(subset2, superset)
//	// false (b has different value)
func IsSubset[K comparable, V comparable](subset, superset map[K]V) bool {
	// Empty or nil subset is subset of any map
	if len(subset) == 0 {
		return true
	}

	// Nil superset cannot contain non-empty subset
	if superset == nil {
		return false
	}

	// Check if all key-value pairs in subset exist in superset
	for k, v := range subset {
		if val, exists := superset[k]; !exists || val != v {
			return false
		}
	}
	return true
}
