package lxmaps

// UniqKeys returns a slice of unique keys from a map.
// If the maps are nil, it returns an empty slice.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	out := UniqKeys(m)
//	// out: []string{"a", "b", "c"}
func UniqKeys[K comparable, V any](in ...map[K]V) []K {
	if len(in) == 0 {
		return nil
	}

	seen := make(map[K]struct{})
	for _, m := range in {
		for k := range m {
			seen[k] = struct{}{}
		}
	}
	result := make([]K, 0, len(seen))
	for k := range seen {
		result = append(result, k)
	}
	return result
}
