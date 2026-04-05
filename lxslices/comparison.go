package lxslices

// Equal checks if two slices are equal (same length and same elements in the same order).
// Two nil slices are considered equal.
// A nil slice is not equal to an empty slice.
//
// Example:
//
//	Equal([]int{1, 2, 3}, []int{1, 2, 3}) // true
//	Equal([]int{1, 2, 3}, []int{1, 2, 4}) // false
//	Equal([]int{}, []int{})                // true
//	Equal([]int(nil), []int(nil))          // true
//	Equal([]int(nil), []int{})             // false
func Equal[T comparable](a, b []T) bool {
	// Handle nil cases
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	// Check length
	if len(a) != len(b) {
		return false
	}

	// Compare elements
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// EqualFunc checks if two slices are equal using a custom comparison function.
// Two nil slices are considered equal.
// A nil slice is not equal to an empty slice.
// The eq function should return true if two elements are considered equal.
//
// Example:
//
//	type Person struct { Name string; Age int }
//	a := []Person{{"Alice", 30}, {"Bob", 25}}
//	b := []Person{{"Alice", 31}, {"Bob", 26}}
//	EqualFunc(a, b, func(p1, p2 Person) bool {
//	    return p1.Name == p2.Name
//	}) // true (comparing by name only)
func EqualFunc[T any](a, b []T, eq func(T, T) bool) bool {
	// Handle nil cases
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	// Check length
	if len(a) != len(b) {
		return false
	}

	// Compare elements using custom function
	for i := range a {
		if !eq(a[i], b[i]) {
			return false
		}
	}

	return true
}

// StartsWith checks if the slice starts with the given prefix.
// Returns true if prefix is nil or empty.
// Returns false if slice is nil and prefix is not nil/empty.
//
// Example:
//
//	StartsWith([]int{1, 2, 3, 4}, []int{1, 2})    // true
//	StartsWith([]int{1, 2, 3, 4}, []int{2, 3})    // false
//	StartsWith([]int{1, 2, 3}, []int{})           // true
//	StartsWith([]int{1, 2}, []int{1, 2, 3})       // false
func StartsWith[T comparable](slice, prefix []T) bool {
	// Empty or nil prefix always matches
	if len(prefix) == 0 {
		return true
	}

	// Nil slice cannot start with non-empty prefix
	if slice == nil {
		return false
	}

	// Prefix longer than slice cannot match
	if len(prefix) > len(slice) {
		return false
	}

	// Compare prefix elements
	for i := range prefix {
		if slice[i] != prefix[i] {
			return false
		}
	}

	return true
}

// EndsWith checks if the slice ends with the given suffix.
// Returns true if suffix is nil or empty.
// Returns false if slice is nil and suffix is not nil/empty.
//
// Example:
//
//	EndsWith([]int{1, 2, 3, 4}, []int{3, 4})    // true
//	EndsWith([]int{1, 2, 3, 4}, []int{2, 3})    // false
//	EndsWith([]int{1, 2, 3}, []int{})           // true
//	EndsWith([]int{3, 4}, []int{1, 2, 3, 4})    // false
func EndsWith[T comparable](slice, suffix []T) bool {
	// Empty or nil suffix always matches
	if len(suffix) == 0 {
		return true
	}

	// Nil slice cannot end with non-empty suffix
	if slice == nil {
		return false
	}

	// Suffix longer than slice cannot match
	if len(suffix) > len(slice) {
		return false
	}

	// Compare suffix elements from the end
	offset := len(slice) - len(suffix)
	for i := range suffix {
		if slice[offset+i] != suffix[i] {
			return false
		}
	}

	return true
}

// HasPrefix is an alias for StartsWith.
// Checks if the slice starts with the given prefix.
//
// Example:
//
//	HasPrefix([]int{1, 2, 3, 4}, []int{1, 2})    // true
func HasPrefix[T comparable](slice, prefix []T) bool {
	return StartsWith(slice, prefix)
}

// HasSuffix is an alias for EndsWith.
// Checks if the slice ends with the given suffix.
//
// Example:
//
//	HasSuffix([]int{1, 2, 3, 4}, []int{3, 4})    // true
func HasSuffix[T comparable](slice, suffix []T) bool {
	return EndsWith(slice, suffix)
}
