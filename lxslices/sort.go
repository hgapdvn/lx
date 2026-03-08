package lxslices

import (
	"sort"

	"github.com/nthanhhai2909/lx/lxconstraints"
)

// SortBy sorts the slice in-place using the provided less function.
func SortBy[T any](slice []T, less func(T, T) bool) {
	sort.Slice(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})
}

// StableSortBy sorts the slice in-place using the provided less function.
// The order of equal elements is preserved.
func StableSortBy[T any](slice []T, less func(T, T) bool) {
	sort.SliceStable(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})
}

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

// IsSortedBy checks if the slice is sorted according to the given comparator function.
// The comparator should return true if a < b (a should come before b).
// Returns true for empty slices or slices with one element.
func IsSortedBy[T any](slice []T, less func(a, b T) bool) bool {
	for i := 1; i < len(slice); i++ {
		if less(slice[i], slice[i-1]) {
			return false
		}
	}
	return true
}
