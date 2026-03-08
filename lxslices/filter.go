package lxslices

// Find returns the first element that satisfies the predicate.
// It returns the element and true if found, otherwise it returns the zero value of T and false.
func Find[T any](slice []T, predicate func(T) bool) (T, bool) {
	for _, e := range slice {
		if predicate(e) {
			return e, true
		}
	}
	var zero T
	return zero, false
}

// FindLast returns the last element that satisfies the predicate.
// It returns the element and true if found, otherwise it returns the zero value of T and false.
func FindLast[T any](slice []T, predicate func(T) bool) (T, bool) {
	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(slice[i]) {
			return slice[i], true
		}
	}
	var zero T
	return zero, false
}

// Filter returns a new slice containing only the elements that satisfy the predicate.
// The original slice is not modified.
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, e := range slice {
		if predicate(e) {
			result = append(result, e)
		}
	}
	return result
}

// Any returns true if at least one element satisfies the predicate.
// It is equivalent to the any() function in Python.
func Any[T any](slice []T, predicate func(T) bool) bool {
	for _, e := range slice {
		if predicate(e) {
			return true
		}
	}
	return false
}

// All returns true if all elements satisfy the predicate.
// It is equivalent to the all() function in Python.
func All[T any](slice []T, predicate func(T) bool) bool {
	for _, e := range slice {
		if !predicate(e) {
			return false
		}
	}
	return true
}

// None returns true if no element satisfies the predicate.
// It is equivalent to the none() function in Python.
func None[T any](slice []T, predicate func(T) bool) bool {
	return !Any(slice, predicate)
}

// Count returns the number of elements in the slice that satisfy the predicate.
func Count[T any](slice []T, predicate func(T) bool) int {
	count := 0
	for _, e := range slice {
		if predicate(e) {
			count++
		}
	}
	return count
}
