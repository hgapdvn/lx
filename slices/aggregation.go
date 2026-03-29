package lxslices

import "github.com/hgapdvn/lx/constraints"

// Reduce applies the given function to each element of the slice and returns the result.
// The order of the elements in the returned slice is the same as in the original slice.
func Reduce[T, U any](slice []T, fn func(accumulator U, element T) U, initial U) U {
	result := initial
	for _, e := range slice {
		result = fn(result, e)
	}
	return result
}

// Sum returns the sum of all elements in the slice.
func Sum[T lxconstraints.Number](slice []T) T {
	var total T
	for _, v := range slice {
		total += v
	}
	return total
}

// Min returns (value, found) for safety on empty slices
func Min[T lxconstraints.Ordered](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}
	res := slice[0]
	for _, v := range slice[1:] {
		if v < res {
			res = v
		}
	}
	return res, true
}

// Max returns (value, found) for safety on empty slices
func Max[T lxconstraints.Ordered](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}
	res := slice[0]
	for _, v := range slice[1:] {
		if v > res {
			res = v
		}
	}
	return res, true
}

// Average returns the arithmetic mean of the slice as a float64 and a boolean
// indicating whether a value was found. For an empty slice it returns 0 and false.
func Average[T lxconstraints.Number](slice []T) (float64, bool) {
	if len(slice) == 0 {
		return 0, false
	}
	var total float64
	for _, v := range slice {
		total += float64(v)
	}
	return total / float64(len(slice)), true
}

// Median returns the median value of the slice and a boolean indicating whether
// a value was found. For an empty slice it returns 0 and false.
// The function creates a sorted copy of the slice to find the median.
// For even-length slices, it returns the average of the two middle elements.
func Median[T lxconstraints.Number](slice []T) (float64, bool) {
	if len(slice) == 0 {
		return 0, false
	}

	// Create a copy to avoid modifying the original
	sorted := Clone(slice)
	SortAsc(sorted)

	n := len(sorted)
	if n%2 == 1 {
		// Odd length: return middle element
		return float64(sorted[n/2]), true
	}
	// Even length: return average of two middle elements
	mid1 := float64(sorted[n/2-1])
	mid2 := float64(sorted[n/2])
	return (mid1 + mid2) / 2, true
}

// Mode returns the most frequent element in the slice and a boolean indicating
// whether a value was found. For an empty slice it returns the zero value and false.
// If multiple elements have the same highest frequency, it returns the first one encountered.
func Mode[T comparable](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}

	// Count occurrences
	counts := make(map[T]int, len(slice))
	for _, v := range slice {
		counts[v]++
	}

	// Find the element with highest count
	var mode T
	maxCount := 0
	for v, count := range counts {
		if count > maxCount {
			maxCount = count
			mode = v
		}
	}

	return mode, true
}

// MinMax returns both the minimum and maximum values in a single pass through the slice.
// Returns (min, max, found). For an empty slice, it returns zero values and false.
// This is more efficient than calling Min and Max separately.
func MinMax[T lxconstraints.Ordered](slice []T) (T, T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, zero, false
	}

	min := slice[0]
	max := slice[0]

	for _, v := range slice[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return min, max, true
}
