package lxtypes

// Optional represents a value that may or may not be present.
// It provides a type-safe way to handle optional values without using nil pointers.
//
// An Optional can be created in three ways:
//   - OptionalOf(value) - Creates an Optional containing a value
//   - OptionalEmpty() - Creates an empty Optional with no value
//   - OptionalOfNullable(ptr) - Creates an Optional from a pointer (empty if nil)
//
// The Get method returns (value, true) if present, or (zero, false) if empty.
// This follows Go's idiomatic comma-ok pattern.
//
// Example:
//
//	// Check if value is present
//	opt := lxtypes.OptionalOf(42)
//	if value, ok := opt.Get(); ok {
//	    fmt.Println("Value:", value)  // Value: 42
//	}
//
//	// Use with default value
//	empty := lxtypes.OptionalEmpty[int]()
//	value := empty.OrElse(99)  // 99
//
//	// Safe nil handling with pointers
//	var ptr *string
//	opt2 := lxtypes.OptionalOfNullable(ptr)
//	value2 := opt2.OrElse("default")  // "default"
type Optional[T any] interface {
	// Get returns the contained value and true if present, or zero value and false if empty.
	// This follows Go's idiomatic comma-ok pattern for optional values.
	//
	// Example:
	//
	//	opt := lxtypes.OptionalOf(42)
	//	if value, ok := opt.Get(); ok {
	//	    fmt.Println(value)  // 42
	//	}
	Get() (T, bool)

	// OrElse returns the contained value if present, or the provided default value if empty.
	//
	// Example:
	//
	//	empty := lxtypes.OptionalEmpty[int]()
	//	value := empty.OrElse(99)  // 99
	OrElse(defaultValue T) T

	// OrElseGet returns the contained value if present, or calls fn and returns its result if empty.
	// Use this when computing the default value is expensive.
	//
	// Example:
	//
	//	empty := lxtypes.OptionalEmpty[int]()
	//	value := empty.OrElseGet(func() int {
	//	    return expensiveComputation()
	//	})
	OrElseGet(fn func() T) T
}

// OptionalOf creates an Optional containing the given value.
// The value is always present in the returned Optional.
//
// Example:
//
//	opt := lxtypes.OptionalOf(42)
//	value, ok := opt.Get()  // value=42, ok=true
func OptionalOf[T any](value T) Optional[T] {
	return valueOptional[T]{value: value}
}

// OptionalOfNullable creates an Optional from a pointer.
// Returns a present Optional if the pointer is non-nil, otherwise returns an empty Optional.
// This is useful for safely converting nullable pointers to Optional values.
//
// Example:
//
//	value := 42
//	opt1 := lxtypes.OptionalOfNullable(&value)  // Present with value 42
//	value1, ok1 := opt1.Get()  // value1=42, ok1=true
//
//	var nilPtr *int
//	opt2 := lxtypes.OptionalOfNullable(nilPtr)  // Empty
//	value2, ok2 := opt2.Get()  // value2=0, ok2=false
func OptionalOfNullable[T any](ptr *T) Optional[T] {
	if ptr == nil {
		return OptionalEmpty[T]()
	}
	return OptionalOf(*ptr)
}

// OptionalEmpty creates an empty Optional with no value.
// Use this to represent the absence of a value in a type-safe way.
//
// Example:
//
//	empty := lxtypes.OptionalEmpty[int]()
//	value, ok := empty.Get()  // value=0, ok=false
//	defaultValue := empty.OrElse(99)  // 99
func OptionalEmpty[T any]() Optional[T] {
	return emptyOptional[T]{}
}

// ----------------------------------- Value Optional implementation -----------------------------------
type valueOptional[T any] struct {
	value T
}

func (v valueOptional[T]) Get() (T, bool) {
	return v.value, true
}

func (v valueOptional[T]) OrElse(defaultValue T) T {
	return v.value
}

func (v valueOptional[T]) OrElseGet(fn func() T) T {
	return v.value
}

// ----------------------------------- Empty Optional implementation -----------------------------------
type emptyOptional[T any] struct{}

func (e emptyOptional[T]) Get() (T, bool) {
	var zero T
	return zero, false
}

func (e emptyOptional[T]) OrElse(defaultValue T) T {
	return defaultValue
}

func (e emptyOptional[T]) OrElseGet(fn func() T) T {
	return fn()
}
