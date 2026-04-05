package lxtypes

// Result represents the outcome of an operation that may succeed or fail.
// It provides a type-safe way to handle operations that can error without using exceptions.
//
// A Result can be created in two ways:
//   - ResultSuccess(value) - Creates a Result containing a success value
//   - ResultFailure(err) - Creates a Result containing an error
//
// The Value method returns (value, nil) if successful, or (zero, error) if failed.
// This follows Go's idiomatic (value, error) pattern.
//
// Example:
//
//	// Success case
//	result := lxtypes.ResultSuccess(42)
//	if value, err := result.Value(); err == nil {
//	    fmt.Println("Success:", value)  // Success: 42
//	}
//
//	// Failure case
//	result2 := lxtypes.ResultFailure[int](errors.New("failed"))
//	if value, err := result2.Value(); err != nil {
//	    fmt.Println("Error:", err)  // Error: failed
//	}
//
//	// Use with default value
//	value := result2.ValueOr(99)  // 99
type Result[T any] interface {
	// Value returns the success value and nil if successful,
	// or zero value and the error if failed.
	// This follows Go's idiomatic (value, error) pattern for result handling.
	//
	// Example:
	//
	//	result := lxtypes.ResultSuccess(42)
	//	if value, err := result.Value(); err == nil {
	//	    fmt.Println(value)  // 42
	//	}
	Value() (T, error)

	// ValueOr returns the success value if successful, or the provided default value if failed.
	// Use this when you want a simple fallback without checking the error.
	//
	// Example:
	//
	//	failure := lxtypes.ResultFailure[int](errors.New("error"))
	//	value := failure.ValueOr(99)  // 99
	ValueOr(defaultValue T) T
}

// ResultSuccess creates a successful Result containing the given value.
// The result will return (value, nil) from Value().
//
// Example:
//
//	result := lxtypes.ResultSuccess(42)
//	value, err := result.Value()  // value=42, err=nil
func ResultSuccess[T any](value T) Result[T] {
	return successResult[T]{value}
}

// ResultFailure creates a failed Result containing the given error.
// The result will return (zero, err) from Value().
//
// Example:
//
//	result := lxtypes.ResultFailure[int](errors.New("operation failed"))
//	value, err := result.Value()  // value=0, err=error("operation failed")
func ResultFailure[T any](err error) Result[T] {
	return failureResult[T]{err: err}
}

// -------------------------------------- Success Result implementation --------------------------------------
type successResult[T any] struct {
	value T
}

func (s successResult[T]) Value() (T, error) {
	return s.value, nil
}

func (s successResult[T]) ValueOr(defaultValue T) T {
	return s.value
}

// -------------------------------------- Failure Result implementation --------------------------------------

type failureResult[T any] struct {
	err error
}

func (f failureResult[T]) Value() (T, error) {
	var zero T
	return zero, f.err
}

func (f failureResult[T]) ValueOr(defaultValue T) T {
	return defaultValue
}
