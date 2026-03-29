package lxmaps

// Values returns a slice of values from a map.
// If the maps are nil, it returns an empty slice.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	out := Values(m)
//	// out: []int{1, 2, 3}
func Values[K comparable, V any](in ...map[K]V) []V {
	if len(in) == 0 {
		return nil
	}

	size := 0
	for i := range in {
		size += len(in[i])
	}

	values := make([]V, 0, size)
	for _, m := range in {
		for _, v := range m {
			values = append(values, v)
		}
	}

	return values
}
