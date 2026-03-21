package lxslices

// Unique returns a new slice with the unique elements of the original slice.
// The order of the elements in the returned slice is the same as in the original slice.
func Unique[T comparable](slice []T) []T {
	if slice == nil {
		return nil
	}
	seen := make(map[T]struct{})
	result := make([]T, 0, len(slice))
	for _, e := range slice {
		if _, ok := seen[e]; !ok {
			seen[e] = struct{}{}
			result = append(result, e)
		}
	}
	return result
}

// Difference returns a new slice containing elements that are in slice1 but not in slice2.
// The order of elements from slice1 is preserved. Duplicates in slice1 are removed.
// Returns nil if slice1 is nil. Returns an empty non-nil slice if slice1 is empty but non-nil.
func Difference[T comparable](slice1, slice2 []T) []T {
	if slice1 == nil {
		return nil
	}
	if len(slice1) == 0 {
		return slice1
	}

	m := make(map[T]struct{}, len(slice2))
	for _, e := range slice2 {
		m[e] = struct{}{}
	}
	result := make([]T, 0, len(slice1))
	seen := make(map[T]struct{})
	for _, e := range slice1 {
		if _, found := m[e]; !found {
			if _, ok := seen[e]; !ok {
				seen[e] = struct{}{}
				result = append(result, e)
			}
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

// Intersection returns a new slice containing elements that appear in both slice1 and slice2.
// The order of elements from slice1 is preserved. Duplicates in slice1 are removed.
func Intersection[T comparable](slice1, slice2 []T) []T {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil
	}
	m := make(map[T]struct{}, len(slice2))
	for _, e := range slice2 {
		m[e] = struct{}{}
	}
	var result []T
	seen := make(map[T]struct{})
	for _, e := range slice1 {
		if _, found := m[e]; found {
			if _, ok := seen[e]; !ok {
				seen[e] = struct{}{}
				result = append(result, e)
			}
		}
	}
	return result
}

// Union returns a new slice containing the union of elements from slice1 and slice2.
// The order of the elements in the returned slice is such that elements from slice1
// come first, followed by elements from slice2 that were not in slice1.
// This function does not modify the original slices.
func Union[T comparable](slice1, slice2 []T) []T {
	// if both nil, return nil
	if slice1 == nil && slice2 == nil {
		return nil
	}
	seen := make(map[T]struct{})
	var result []T
	// append unique elements from slice1 (preserve order)
	for _, e := range slice1 {
		if _, ok := seen[e]; !ok {
			seen[e] = struct{}{}
			result = append(result, e)
		}
	}
	// append elements from slice2 that haven't been seen
	for _, e := range slice2 {
		if _, ok := seen[e]; !ok {
			seen[e] = struct{}{}
			result = append(result, e)
		}
	}
	// if result is empty but one of inputs was a non-nil empty slice, return empty slice
	if len(result) == 0 {
		// if either input is a non-nil empty slice, return empty slice (not nil)
		if len(slice1) == 0 || len(slice2) == 0 {
			return []T{}
		}
		// otherwise both were nil or had no unique elements; returning nil is fine
		return nil
	}
	return result
}
