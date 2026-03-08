package lxslices

import (
	"math/rand"
	"time"
)

// rng is a package-level random source, seeded at init time.
// This ensures proper randomness on Go <1.20 where the global
// math/rand source is deterministic by default.
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// Sample returns a random element from the slice.
// Returns the zero value of T if the slice is empty or nil.
func Sample[T any](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	return slice[rng.Intn(len(slice))]
}

// SampleN returns n random elements from the slice.
// If n >= len(slice), returns a shuffled copy of the entire slice.
// If n <= 0 or slice is empty, returns an empty slice.
// Elements are sampled without replacement (no duplicates).
func SampleN[T any](slice []T, n int) []T {
	if len(slice) == 0 || n <= 0 {
		return []T{}
	}

	if n >= len(slice) {
		// Return shuffled copy of entire slice
		result := make([]T, len(slice))
		copy(result, slice)
		rng.Shuffle(len(result), func(i, j int) {
			result[i], result[j] = result[j], result[i]
		})
		return result
	}

	// Sample n elements without replacement
	result := make([]T, n)
	indices := rng.Perm(len(slice))
	for i := 0; i < n; i++ {
		result[i] = slice[indices[i]]
	}
	return result
}
