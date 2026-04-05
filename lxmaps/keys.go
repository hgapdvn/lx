package lxmaps

// Keys returns a slice of keys from a map.
// If the map is nil, it returns an empty slice.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	out := Keys(m)
//	// out: []string{"a", "b", "c"}
//
//	m2 := map[string]int{}
//	out2 := Keys(m2)
//	// out2: []string{}
//
//	m3 := map[string]int{"a": 1, "b": 2, "c": 3}
//	out3 := Keys(m3, m3)
//	// out3: []string{"a", "b", "c"}
func Keys[K comparable, V any](in ...map[K]V) []K {
	if len(in) == 0 {
		return nil
	}
	size := 0
	for i := range in {
		size += len(in[i])
	}

	keys := make([]K, 0, size)
	for _, m := range in {
		for k := range m {
			keys = append(keys, k)
		}
	}
	return keys
}
