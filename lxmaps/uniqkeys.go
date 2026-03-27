package lxmaps

// UniqKeys returns a slice of unique keys from a map.
// If the maps are nil, it returns an empty slice.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	out := UniqKeys(m)
//	// out: []string{"a", "b", "c"}
//
//	m2 := map[string]int{}
//	out2 := UniqKeys(m2)
//	// out2: []string{}
//
//	m3 := map[string]int{"a": 1, "b": 2, "c": 3}
//	out3 := UniqKeys(m3, m3)
//	// out3: []string{"a", "b", "c"}
//
//	m4 := map[string]int{"a": 1, "b": 2, "c": 3}
//	out4 := UniqKeys(m4, m4)
//	// out4: []string{"a", "b", "c"}
//
//	m5 := map[string]int{"a": 1, "b": 2, "c": 3}
//	out5 := UniqKeys(m5, m5)
//	// out5: []string{"a", "b", "c"}
//
//	m6 := map[string]int{"a": 1, "b": 2, "c": 3}
//	out6 := UniqKeys(m6, m6)
//	// out6: []string{"a", "b", "c"}
//
//	m7 := map[string]int{"a": 1, "b": 2, "c": 3}
//	out7 := UniqKeys(m7, m7)
//	// out7: []string{"a", "b", "c"}
//
//	m8 := map[string]int{"a": 1, "b": 2, "c": 3}
//	out8 := UniqKeys(m8, m8)
//	// out8: []string{"a", "b", "c"}
//
//	m9 := map[string]int{"a": 1, "b": 2, "c": 3}
//	out9 := UniqKeys(m9, m9)
//	// out9: []string{"a", "b", "c"}
//
//	m10 := map[string]int{"a": 1, "b": 2, "c": 3}
//	out10 := UniqKeys(m10, m10)
//	// out10: []string{"a", "b", "c"}
//
//	m11 := map[string]int{"a": 1, "b": 2, "c": 3}
//	out11 := UniqKeys(m11, m11)
//	// out11: []string{"a", "b", "c"}	
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
	return Keys(seen)
}
