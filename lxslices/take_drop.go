package lxslices

import "github.com/nthanhhai2909/lx/lxtypes"

// Take returns the first n elements from the slice.
// If n is greater than the slice length, returns the entire slice.
// If n is less than or equal to 0, returns an empty slice.
// Returns nil if the input slice is nil.
//
// Example:
//
//	Take([]int{1, 2, 3, 4, 5}, 3) // []int{1, 2, 3}
//	Take([]int{1, 2}, 5)          // []int{1, 2}
//	Take([]int{1, 2, 3}, 0)       // []int{}
func Take[T any](slice []T, n int) []T {
	if slice == nil {
		return nil
	}

	if n <= 0 {
		return []T{}
	}

	if n >= len(slice) {
		return Clone(slice)
	}

	return Clone(slice[:n])
}

// TakeLast returns the last n elements from the slice.
// If n is greater than the slice length, returns the entire slice.
// If n is less than or equal to 0, returns an empty slice.
// Returns nil if the input slice is nil.
//
// Example:
//
//	TakeLast([]int{1, 2, 3, 4, 5}, 3) // []int{3, 4, 5}
//	TakeLast([]int{1, 2}, 5)          // []int{1, 2}
//	TakeLast([]int{1, 2, 3}, 0)       // []int{}
func TakeLast[T any](slice []T, n int) []T {
	if slice == nil {
		return nil
	}

	if n <= 0 {
		return []T{}
	}

	if n >= len(slice) {
		return Clone(slice)
	}

	return Clone(slice[len(slice)-n:])
}

// TakeWhile returns elements from the start of the slice while the predicate returns true.
// Stops at the first element where predicate returns false.
// Returns nil if the input slice is nil.
// Returns an empty slice if the predicate returns false for the first element.
//
// Example:
//
//	isEven := lxtypes.Predicate[int](func(n int) bool { return n%2 == 0 })
//	TakeWhile([]int{2, 4, 6, 7, 8}, isEven) // []int{2, 4, 6}
func TakeWhile[T any](slice []T, predicate lxtypes.Predicate[T]) []T {
	if slice == nil {
		return nil
	}

	for i, v := range slice {
		if !predicate(v) {
			if i == 0 {
				return []T{}
			}
			return Clone(slice[:i])
		}
	}

	return Clone(slice)
}

// Drop returns the slice without the first n elements.
// If n is greater than or equal to the slice length, returns an empty slice.
// If n is less than or equal to 0, returns a copy of the entire slice.
// Returns nil if the input slice is nil.
//
// Example:
//
//	Drop([]int{1, 2, 3, 4, 5}, 2) // []int{3, 4, 5}
//	Drop([]int{1, 2}, 5)          // []int{}
//	Drop([]int{1, 2, 3}, 0)       // []int{1, 2, 3}
func Drop[T any](slice []T, n int) []T {
	if slice == nil {
		return nil
	}

	if n <= 0 {
		return Clone(slice)
	}

	if n >= len(slice) {
		return []T{}
	}

	return Clone(slice[n:])
}

// DropLast returns the slice without the last n elements.
// If n is greater than or equal to the slice length, returns an empty slice.
// If n is less than or equal to 0, returns a copy of the entire slice.
// Returns nil if the input slice is nil.
//
// Example:
//
//	DropLast([]int{1, 2, 3, 4, 5}, 2) // []int{1, 2, 3}
//	DropLast([]int{1, 2}, 5)          // []int{}
//	DropLast([]int{1, 2, 3}, 0)       // []int{1, 2, 3}
func DropLast[T any](slice []T, n int) []T {
	if slice == nil {
		return nil
	}

	if n <= 0 {
		return Clone(slice)
	}

	if n >= len(slice) {
		return []T{}
	}

	return Clone(slice[:len(slice)-n])
}

// DropWhile returns the slice after skipping elements from the start while the predicate returns true.
// Starts including elements from the first element where predicate returns false.
// Returns nil if the input slice is nil.
// Returns an empty slice if the predicate returns true for all elements.
//
// Example:
//
//	isEven := lxtypes.Predicate[int](func(n int) bool { return n%2 == 0 })
//	DropWhile([]int{2, 4, 6, 7, 8}, isEven) // []int{7, 8}
func DropWhile[T any](slice []T, predicate lxtypes.Predicate[T]) []T {
	if slice == nil {
		return nil
	}

	for i, v := range slice {
		if !predicate(v) {
			return Clone(slice[i:])
		}
	}

	return []T{}
}
