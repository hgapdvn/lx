package lxmaps

// Flatten flattens a map of slices into a single slice containing all elements.
// Returns nil if the input map is nil, an empty slice if the map is empty.
// The order of elements is not guaranteed due to map iteration order.
//
// Example:
//
//	m := map[string][]int{
//		"a": {1, 2},
//		"b": {3, 4},
//		"c": {5},
//	}
//	result := Flatten(m)
//	// result: []int{1, 2, 3, 4, 5} (order may vary due to map iteration)
func Flatten[K comparable, V any](m map[K][]V) []V {
	if m == nil {
		return nil
	}

	// First pass: calculate total size
	totalSize := 0
	for _, slice := range m {
		totalSize += len(slice)
	}

	if totalSize == 0 {
		return []V{}
	}

	// Second pass: collect all elements
	result := make([]V, 0, totalSize)
	for _, slice := range m {
		result = append(result, slice...)
	}
	return result
}
