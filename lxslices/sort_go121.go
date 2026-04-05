//go:build go1.21

package lxslices

import "slices"

// SortBy sorts the slice in-place using the provided less function.
// On Go 1.21+ this uses slices.SortFunc for improved performance.
func SortBy[T any](slice []T, less func(T, T) bool) {
	slices.SortFunc(slice, func(a, b T) int {
		if less(a, b) {
			return -1
		}
		if less(b, a) {
			return 1
		}
		return 0
	})
}

// StableSortBy sorts the slice in-place using the provided less function.
// The order of equal elements is preserved.
// On Go 1.21+ this uses slices.SortStableFunc for improved performance.
func StableSortBy[T any](slice []T, less func(T, T) bool) {
	slices.SortStableFunc(slice, func(a, b T) int {
		if less(a, b) {
			return -1
		}
		if less(b, a) {
			return 1
		}
		return 0
	})
}

// IsSortedBy checks if the slice is sorted according to the given comparator function.
// The comparator should return true if a < b (a should come before b).
// Returns true for empty slices or slices with one element.
// On Go 1.21+ this uses slices.IsSortedFunc.
func IsSortedBy[T any](slice []T, less func(a, b T) bool) bool {
	return slices.IsSortedFunc(slice, func(a, b T) int {
		if less(a, b) {
			return -1
		}
		if less(b, a) {
			return 1
		}
		return 0
	})
}
