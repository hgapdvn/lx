package lxslices

import "github.com/hgapdvn/lx/lxconstraints"

// SortAsc sorts the slice in-place in ascending order.
func SortAsc[T lxconstraints.Ordered](slice []T) {
	SortBy(slice, func(a, b T) bool { return a < b })
}

// SortDesc sorts the slice in-place in descending order.
func SortDesc[T lxconstraints.Ordered](slice []T) {
	SortBy(slice, func(a, b T) bool { return a > b })
}

// IsSortedAsc checks if the slice is sorted in ascending order using the < operator.
// Returns true for empty slices or slices with one element.
func IsSortedAsc[T lxconstraints.Ordered](slice []T) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			return false
		}
	}
	return true
}

// IsSortedDesc checks if the slice is sorted in descending order using the > operator.
// Returns true for empty slices or slices with one element.
func IsSortedDesc[T lxconstraints.Ordered](slice []T) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] < slice[i] {
			return false
		}
	}
	return true
}
