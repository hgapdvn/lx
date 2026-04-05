package lxslices

import (
	"math/rand"
	"time"
)

func init() {
	// Seed the global random source for Go <1.20 where the global source is
	// deterministic by default. On Go 1.20+ this is a no-op (auto-seeded).
	rand.Seed(time.Now().UnixNano()) //nolint:staticcheck
}

// Sample returns a random element from the slice and true.
// Returns the zero value of T and false if the slice is empty or nil.
func Sample[T any](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}
	return slice[rand.Intn(len(slice))], true
}

// SampleN returns n random elements from the slice without replacement.
// If n >= len(slice), returns a shuffled copy of the entire slice.
// If n <= 0, returns an empty slice (nil if input is nil, non-nil otherwise).
// Preserves nil vs empty slice semantics: nil input returns nil; non-nil empty
// input returns a non-nil empty slice.
func SampleN[T any](slice []T, n int) []T {
	if slice == nil {
		return nil
	}
	if len(slice) == 0 || n <= 0 {
		return []T{}
	}

	if n >= len(slice) {
		// Return shuffled copy of entire slice
		result := make([]T, len(slice))
		copy(result, slice)
		rand.Shuffle(len(result), func(i, j int) {
			result[i], result[j] = result[j], result[i]
		})
		return result
	}

	// Sample n elements without replacement using a random permutation.
	indices := rand.Perm(len(slice))
	result := make([]T, n)
	for i := 0; i < n; i++ {
		result[i] = slice[indices[i]]
	}
	return result
}
