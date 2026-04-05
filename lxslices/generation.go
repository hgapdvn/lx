package lxslices

import "github.com/hgapdvn/lx/lxconstraints"

// Repeat creates a new slice with the given value repeated n times.
// Returns nil if n is less than or equal to 0.
//
// Example:
//
//	result := lxslices.Repeat("hello", 3)
//	// result: ["hello", "hello", "hello"]
func Repeat[T any](value T, n int) []T {
	if n <= 0 {
		return nil
	}
	result := make([]T, n)
	for i := 0; i < n; i++ {
		result[i] = value
	}
	return result
}

// RepeatSlice creates a new slice by repeating the given slice n times.
// Returns nil if the input slice is nil or n is less than or equal to 0.
// Returns an empty slice if the input slice is empty and n is greater than 0.
//
// Example:
//
//	result := lxslices.RepeatSlice([]int{1, 2}, 3)
//	// result: [1, 2, 1, 2, 1, 2]
func RepeatSlice[T any](slice []T, n int) []T {
	if slice == nil || n <= 0 {
		return nil
	}
	if len(slice) == 0 {
		return []T{}
	}
	result := make([]T, 0, len(slice)*n)
	for i := 0; i < n; i++ {
		result = append(result, slice...)
	}
	return result
}

// Range creates a slice of integers from start (inclusive) to end (exclusive).
// If start >= end, returns nil.
//
// Example:
//
//	result := lxslices.Range(1, 5)
//	// result: [1, 2, 3, 4]
//
//	result := lxslices.Range(5, 5)
//	// result: nil
func Range[T lxconstraints.Integer](start, end T) []T {
	if start >= end {
		return nil
	}
	size := int(end - start)
	result := make([]T, size)
	for i := 0; i < size; i++ {
		result[i] = start + T(i)
	}
	return result
}

// RangeStep creates a slice of integers from start (inclusive) towards end (exclusive)
// with the given step size.
// Returns nil if:
//   - step is 0
//   - step is positive and start >= end
//   - step is negative and start <= end
//
// Example:
//
//	result := lxslices.RangeStep(0, 10, 2)
//	// result: [0, 2, 4, 6, 8]
//
//	result := lxslices.RangeStep(10, 0, -2)
//	// result: [10, 8, 6, 4, 2]
func RangeStep[T lxconstraints.Integer](start, end, step T) []T {
	if step == 0 {
		return nil
	}
	if step > 0 && start >= end {
		return nil
	}
	if step < 0 && start <= end {
		return nil
	}

	// Calculate size
	var size int
	if step > 0 {
		size = int((end - start + step - 1) / step)
	} else {
		size = int((end - start + step + 1) / step)
	}

	result := make([]T, size)
	value := start
	for i := 0; i < size; i++ {
		result[i] = value
		value += step
	}
	return result
}
