package lxslices

// Window returns all contiguous sub-slices (sliding windows) of the specified size.
// Each window is a sub-slice of the original slice and shares underlying memory.
// Returns ErrInvalidSize if size <= 0.
// Returns nil if the input slice is nil.
// Returns an empty slice if the input is empty or len(slice) < size.
//
// Example:
//
//	windows, err := Window([]int{1, 2, 3, 4}, 2)
//	// windows: [[1, 2], [2, 3], [3, 4]], err: nil
func Window[T any](slice []T, size int) ([][]T, error) {
	if size <= 0 {
		return nil, ErrInvalidSize
	}
	if slice == nil {
		return nil, nil
	}
	if len(slice) < size {
		return [][]T{}, nil
	}
	count := len(slice) - size + 1
	result := make([][]T, count)
	for i := 0; i < count; i++ {
		result[i] = slice[i : i+size]
	}
	return result, nil
}

// WindowFunc applies the given function to each sliding window of the specified size
// and returns a slice of results.
// Returns ErrInvalidSize if size <= 0.
// Returns nil if the input slice is nil.
// Returns an empty slice if the input is empty or len(slice) < size.
//
// Example:
//
//	sums, err := WindowFunc([]int{1, 2, 3, 4}, 2, func(w []int) int {
//	    return w[0] + w[1]
//	})
//	// sums: [3, 5, 7], err: nil
func WindowFunc[T, U any](slice []T, size int, fn func([]T) U) ([]U, error) {
	if size <= 0 {
		return nil, ErrInvalidSize
	}
	if slice == nil {
		return nil, nil
	}
	if len(slice) < size {
		return []U{}, nil
	}
	count := len(slice) - size + 1
	result := make([]U, count)
	for i := 0; i < count; i++ {
		result[i] = fn(slice[i : i+size])
	}
	return result, nil
}
