package lxslices

import (
	"github.com/nthanhhai2909/lx/lxconstraints"
	"github.com/nthanhhai2909/lx/types"
)

// Index returns the index of the first instance of elem in slice, or -1 if elem is not present in slice.
func Index[T comparable](slice []T, elem T) int {
	for i, e := range slice {
		if e == elem {
			return i
		}
	}
	return -1
}

// IndexFunc returns the index of the first element in slice for which predicate returns true, or -1 if none do.
func IndexFunc[T any](slice []T, predicate func(T) bool) int {
	for i, e := range slice {
		if predicate(e) {
			return i
		}
	}
	return -1
}

// LastIndex returns the index of the last instance of elem in slice, or -1 if elem is not present in slice.
func LastIndex[T comparable](slice []T, elem T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == elem {
			return i
		}
	}
	return -1
}

// LastIndexFunc returns the index of the last element in slice for which predicate returns true, or -1 if none do.
func LastIndexFunc[T any](slice []T, predicate func(T) bool) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(slice[i]) {
			return i
		}
	}
	return -1
}

// MinIndex returns the index of the minimum element in the slice, or -1 if the slice is empty.
func MinIndex[T lxconstraints.Ordered](slice []T) (int, bool) {
	if len(slice) == 0 {
		return -1, false
	}
	idx := 0
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[idx] {
			idx = i
		}
	}
	return idx, true
}

// MaxIndex returns the index of the maximum element in the slice, or -1 if the slice is empty.
func MaxIndex[T lxconstraints.Ordered](slice []T) (int, bool) {
	if len(slice) == 0 {
		return -1, false
	}
	idx := 0
	for i := 1; i < len(slice); i++ {
		if slice[i] > slice[idx] {
			idx = i
		}
	}
	return idx, true
}

// First returns the first element of the slice and true, or the zero value and false if the slice is empty.
func First[T any](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}
	return slice[0], true
}

// Last returns the last element of the slice and true, or the zero value and false if the slice is empty.
func Last[T any](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}
	return slice[len(slice)-1], true
}

// Get returns the element at index and true, or the zero value and false if the index is out of bounds.
func Get[T any](slice []T, index int) (T, bool) {
	if index < 0 || index >= len(slice) {
		var zero T
		return zero, false
	}
	return slice[index], true
}

// BinarySearch performs binary search on a sorted slice to find the target element.
// Returns the index where target is found, or -1 if not found.
// The slice must be sorted in ascending order, otherwise the behavior is undefined.
// Returns -1 if the slice is nil or empty.
//
// Example:
//
//	sorted := []int{1, 3, 5, 7, 9}
//	idx := BinarySearch(sorted, 5)
//	// idx: 2
//	idx = BinarySearch(sorted, 4)
//	// idx: -1 (not found)
func BinarySearch[T lxconstraints.Ordered](slice []T, target T) int {
	if len(slice) == 0 {
		return -1
	}

	left, right := 0, len(slice)-1
	for left <= right {
		mid := left + (right-left)/2
		if slice[mid] == target {
			return mid
		}
		if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// BinarySearchFunc performs binary search on a slice using a custom comparator function.
// The comparator should return a negative integer, zero, or a positive integer as the
// target is less than, equal to, or greater than the element being compared.
// Returns the index where target is found, or -1 if not found.
// The slice must be sorted according to the comparator, otherwise the behavior is undefined.
// Returns -1 if the slice is nil or empty.
//
// Example:
//
//	type Person struct { Name string; Age int }
//	people := []Person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 35}}
//	cmp := lxtypes.Comparator[Person](func(target, elem Person) int {
//	    if target.Age < elem.Age { return -1 }
//	    if target.Age > elem.Age { return 1 }
//	    return 0
//	})
//	idx := BinarySearchFunc(people, Person{"", 30}, cmp)
//	// idx: 1
func BinarySearchFunc[T any](slice []T, target T, comparator lxtypes.Comparator[T]) int {
	if len(slice) == 0 {
		return -1
	}

	left, right := 0, len(slice)-1
	for left <= right {
		mid := left + (right-left)/2
		cmp := comparator(target, slice[mid])
		if cmp == 0 {
			return mid
		}
		if cmp > 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
