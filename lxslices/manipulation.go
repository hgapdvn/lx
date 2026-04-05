package lxslices

// Append appends the provided elements to the end of the slice and returns the result.
func Append[T any](slice []T, elems ...T) []T {
	if len(elems) == 0 {
		return slice
	}
	return append(slice, elems...)
}

// Prepend adds the provided elements to the beginning of the slice and returns the result.
func Prepend[T any](slice []T, elems ...T) []T {
	if len(elems) == 0 {
		return slice
	}
	// Preallocate exact capacity for efficiency
	res := make([]T, 0, len(elems)+len(slice))
	res = append(res, elems...)
	res = append(res, slice...)
	return res
}

// Insert inserts elem at the specified index and returns the new slice.
// If index < 0, it is clamped to 0 (insertion at the beginning).
// If index >= len(slice), the element is appended to the end.
// Inserting into a nil slice returns a new non-nil slice containing just elem.
func Insert[T any](slice []T, index int, elem T) []T {
	if index < 0 {
		index = 0
	}
	if index >= len(slice) {
		return append(append([]T{}, slice...), elem)
	}
	result := make([]T, len(slice)+1)
	copy(result, slice[:index])
	result[index] = elem
	copy(result[index+1:], slice[index:])
	return result
}

// Remove removes the first occurrence of value from the slice and returns the result.
// If the value is not present the original slice is returned unchanged.
func Remove[T comparable](slice []T, value T) []T {
	for i, v := range slice {
		if v == value {
			return RemoveAt(slice, i)
		}
	}
	return slice
}

// RemoveAt removes the element at index and returns the resulting slice.
// If index is out of bounds the original slice is returned unchanged.
func RemoveAt[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	// Use append to build the new slice without the element at index
	return append(append([]T{}, slice[:index]...), slice[index+1:]...)
}

// RemoveAll removes all occurrences of value from the slice and returns the result.
func RemoveAll[T comparable](slice []T, value T) []T {
	if len(slice) == 0 {
		return slice
	}
	res := make([]T, 0, len(slice))
	for _, v := range slice {
		if v != value {
			res = append(res, v)
		}
	}
	// If nothing removed, return original slice to preserve nil vs empty semantics
	if len(res) == len(slice) {
		return slice
	}
	return res
}

// RemoveFunc removes elements for which predicate returns true and returns the result.
func RemoveFunc[T any](slice []T, predicate func(T) bool) []T {
	if len(slice) == 0 {
		return slice
	}
	res := make([]T, 0, len(slice))
	for _, v := range slice {
		if !predicate(v) {
			res = append(res, v)
		}
	}
	if len(res) == len(slice) {
		return slice
	}
	return res
}

// RemoveDuplicates is an alias for Unique.
func RemoveDuplicates[T comparable](slice []T) []T {
	return Unique(slice)
}

// Replace replaces all occurrences of old with new and returns the resulting slice.
func Replace[T comparable](slice []T, old, new T) []T {
	if len(slice) == 0 {
		return slice
	}
	res := make([]T, len(slice))
	copy(res, slice)
	changed := false
	for i, v := range res {
		if v == old {
			res[i] = new
			changed = true
		}
	}
	if !changed {
		return slice
	}
	return res
}

// ReplaceAt sets the element at index to new and returns the resulting slice.
// If index is out of bounds the original slice is returned unchanged.
func ReplaceAt[T any](slice []T, index int, new T) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	res := make([]T, len(slice))
	copy(res, slice)
	res[index] = new
	return res
}

// RotateLeft rotates the slice to the left by k positions in-place.
// For k <= 0, it rotates to the right by -k positions.
func RotateLeft[T any](slice []T, k int) {
	n := len(slice)
	if n == 0 || k == 0 {
		return
	}

	// Normalize k to [0, n)
	k = k % n
	if k < 0 {
		k += n
	}
	if k == 0 {
		return
	}

	// Use the reversal algorithm:
	// 1. Reverse [0, k)
	// 2. Reverse [k, n)
	// 3. Reverse [0, n)
	Reverse(slice[:k])
	Reverse(slice[k:])
	Reverse(slice)
}

// RotateRight rotates the slice to the right by k positions in-place.
// Equivalent to RotateLeft(slice, -k).
func RotateRight[T any](slice []T, k int) {
	RotateLeft(slice, -k)
}
