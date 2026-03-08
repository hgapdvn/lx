package lxslices

import (
	"fmt"

	"github.com/nthanhhai2909/lx/lxtypes"
)

// Map applies the given function to each element of the slice and returns a new slice with the results.
// The original slice is not modified.
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, e := range slice {
		result[i] = fn(e)
	}
	return result
}

// FlatMap applies the given function to each element of the slice and returns a new slice with the results.
// The order of the elements in the returned slice is the same as in the original slice.
func FlatMap[T, U any](slice []T, fn func(T) []U) []U {
	var result []U
	for _, e := range slice {
		result = append(result, fn(e)...)
	}
	return result
}

// ForEach applies the given function to each element of the slice.
// It is primarily used for side effects.
func ForEach[T any](slice []T, fn func(T)) {
	for _, e := range slice {
		fn(e)
	}
}

// ForEachIndexed applies the given function to each element of the slice along with its index.
// It is primarily used for side effects when the index is needed.
func ForEachIndexed[T any](slice []T, fn func(int, T)) {
	for i, e := range slice {
		fn(i, e)
	}
}

// Reverse reverses the elements of the slice in-place.
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// GroupBy groups the elements of the slice by the given function.
// The order of the elements in the returned slice is the same as in the original slice.
func GroupBy[T any, K comparable](slice []T, fn func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, e := range slice {
		key := fn(e)
		result[key] = append(result[key], e)
	}
	return result
}

// AssociateBy creates a map from the elements of the slice using the given key-selector function.
// It returns an error if the function produces duplicate keys, acting as a strict map builder.
func AssociateBy[T any, K comparable](slice []T, fn func(T) K) (map[K]T, error) {
	result := make(map[K]T)
	for _, e := range slice {
		key := fn(e)
		if _, exists := result[key]; exists {
			return nil, fmt.Errorf("%w: %v", ErrDuplicateKey, key)
		}
		result[key] = e
	}
	return result, nil
}

// Chunk splits a slice into a slice of consecutive smaller slices (chunks) of the specified size.
// The last chunk may be smaller than the given size if the slice length is not perfectly divisible.
// Returns ErrInvalidSize if size <= 0.
// Returns nil if the input slice is nil.
// Returns an empty slice of slices if the input slice is empty.
//
// Example:
//
//	chunks, err := Chunk([]int{1, 2, 3, 4, 5}, 2)
//	// chunks: [[1, 2], [3, 4], [5]], err: nil
func Chunk[T any](slice []T, size int) ([][]T, error) {
	if size <= 0 {
		return nil, ErrInvalidSize
	}
	if slice == nil {
		return nil, nil
	}
	if len(slice) == 0 {
		return [][]T{}, nil
	}

	chunks := make([][]T, 0, (len(slice)+size-1)/size)
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks, nil
}

// PartitionN splits a slice into N chunks of approximately equal size.
// The earlier chunks will be larger if the slice cannot be split evenly.
// Returns ErrInvalidSize if n <= 0.
// Returns nil if the input slice is nil.
// Returns an empty slice of slices if the input slice is empty.
//
// Example:
//
//	parts, err := PartitionN([]int{1, 2, 3, 4, 5}, 2)
//	// parts: [[1, 2, 3], [4, 5]], err: nil
func PartitionN[T any](slice []T, n int) ([][]T, error) {
	if n <= 0 {
		return nil, ErrInvalidSize
	}
	if slice == nil {
		return nil, nil
	}
	if n == 1 {
		return [][]T{slice}, nil
	}

	length := len(slice)
	if length == 0 {
		return [][]T{}, nil
	}

	chunks := make([][]T, 0, n)
	chunkSize := length / n
	remainder := length % n

	start := 0
	for i := 0; i < n; i++ {
		end := start + chunkSize
		// Distribute the remainder among the first chunks
		if remainder > 0 {
			end++
			remainder--
		}
		if end > length {
			end = length
		}
		// Even if start == end, we should append an empty chunk if length < n
		if start < length {
			chunks = append(chunks, slice[start:end])
		} else {
			chunks = append(chunks, []T{})
		}
		start = end
	}
	return chunks, nil
}

// SplitAt splits a slice at the given index and returns two new slices.
// Returns (slice[:index], slice[index:]).
// If index is <= 0, returns (nil, slice).
// If index >= len(slice), returns (slice, nil).
// If the input slice is nil, it returns (nil, nil).
func SplitAt[T any](slice []T, index int) ([]T, []T) {
	if slice == nil {
		return nil, nil
	}
	if index <= 0 {
		return nil, slice
	}
	if index >= len(slice) {
		return slice, nil
	}
	return slice[:index], slice[index:]
}

// Concat concatenates multiple slices into a single slice.
// Behavior:
// - If no slices are provided, returns nil.
// - If all provided slices are nil, returns nil.
// - If one or more slices are non-nil but total length is zero, returns an empty non-nil slice.
func Concat[T any](slices ...[]T) []T {
	if len(slices) == 0 {
		return nil
	}
	total := 0
	hadNonNil := false
	for _, s := range slices {
		if s != nil {
			hadNonNil = true
		}
		total += len(s)
	}
	if total == 0 {
		if hadNonNil {
			return []T{}
		}
		return nil
	}
	res := make([]T, 0, total)
	for _, s := range slices {
		res = append(res, s...)
	}
	return res
}

// Zip combines two slices into a slice of Pair. The length of the result is the
// minimum of the two input lengths. If both inputs are nil, returns nil. If both
// inputs are empty but non-nil, returns an empty non-nil slice.
func Zip[T any, U any](a []T, b []U) []lxtypes.Pair[T, U] {
	if a == nil && b == nil {
		return nil
	}
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	if n == 0 {
		// at least one input was non-nil (since both-nil handled above), return empty slice
		return []lxtypes.Pair[T, U]{}
	}
	res := make([]lxtypes.Pair[T, U], n)
	for i := 0; i < n; i++ {
		res[i] = lxtypes.Pair[T, U]{First: a[i], Second: b[i]}
	}
	return res
}

// Unzip splits a slice of Pair into two slices. If pairs is nil, returns (nil, nil).
// If pairs is an empty non-nil slice, returns two empty non-nil slices.
func Unzip[T any, U any](pairs []lxtypes.Pair[T, U]) ([]T, []U) {
	if pairs == nil {
		return nil, nil
	}
	n := len(pairs)
	if n == 0 {
		return []T{}, []U{}
	}
	first := make([]T, n)
	second := make([]U, n)
	for i := 0; i < n; i++ {
		first[i] = pairs[i].First
		second[i] = pairs[i].Second
	}
	return first, second
}

// Copy creates a shallow copy of the slice.
// Returns a new slice with the same elements. For nil input, returns nil.
func Copy[T any](slice []T) []T {
	if slice == nil {
		return nil
	}

	result := make([]T, len(slice))
	copy(result, slice)
	return result
}

// Clone is an alias for Copy that creates a shallow copy of the slice.
// Returns a new slice with the same elements. For nil input, returns nil.
func Clone[T any](slice []T) []T {
	return Copy(slice)
}
