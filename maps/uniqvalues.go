package lxmaps

// UniqValues returns a slice of unique values from a map.
// If the maps are nil, it returns an empty slice.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	out := UniqValues(m)
//	// out: []int{1, 2, 3}
//

func UniqValues[K comparable, V comparable](in ...map[K]V) []V {
	if len(in) == 0 {
		return nil
	}

	uniqValueMap := make(map[V]struct{})
	for _, m := range in {
		for _, v := range m {
			uniqValueMap[v] = struct{}{}
		}
	}

	uniqValues := make([]V, 0, len(uniqValueMap))
	for v := range uniqValueMap {
		uniqValues = append(uniqValues, v)
	}

	return uniqValues
}
