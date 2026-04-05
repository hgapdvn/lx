//go:build !go1.21

package lxslices

import "sort"

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
