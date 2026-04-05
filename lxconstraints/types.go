package lxconstraints

// Package lxconstraints provides reusable type constraints for generic programming.
//
// These constraints enable type-safe generic functions across the lx library
// and can be used in your own code for better type safety and code reuse.
//
// Example:
//
//	func Sum[T lxconstraints.Number](values []T) T {
//	    var sum T
//	    for _, v := range values {
//	        sum += v
//	    }
//	    return sum
//	}

// Integer represents all signed and unsigned integer types.
// The ~ operator allows named types with these underlying types.
//
// Includes: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Signed represents all signed integer types.
//
// Includes: int, int8, int16, int32, int64
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned represents all unsigned integer types.
//
// Includes: uint, uint8, uint16, uint32, uint64, uintptr
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Float represents all floating point types.
//
// Includes: float32, float64
type Float interface {
	~float32 | ~float64
}

// Complex represents all complex number types.
//
// Includes: complex64, complex128
type Complex interface {
	~complex64 | ~complex128
}

// Number represents all numeric types (integers and floats).
// Useful for arithmetic operations that work with any numeric type.
//
// Includes: All Integer and Float types
type Number interface {
	Integer | Float
}

// Numeric represents all numeric types including complex numbers.
// Use this when you need to support complex arithmetic.
//
// Includes: All Integer, Float, and Complex types
type Numeric interface {
	Integer | Float | Complex
}

// Ordered represents all types that support ordering operations (<, <=, >, >=).
// This includes all numeric types and strings.
//
// Includes: All Number types and string
type Ordered interface {
	Number | ~string
}

// Addable represents types that support the + operator.
// This includes all numeric types and strings (for concatenation).
//
// Includes: All Numeric types and string
type Addable interface {
	Numeric | ~string
}

// Slice represents any slice type.
// Useful for functions that work with slices regardless of element type.
//
// Example:
//
//	func IsEmpty[T any, S lxconstraints.Slice[T]](s S) bool {
//	    return len(s) == 0
//	}
type Slice[T any] interface {
	~[]T
}

// Map represents any map type.
// Useful for functions that work with maps regardless of key/value types.
//
// Example:
//
//	func IsEmpty[K comparable, V any, M lxconstraints.Map[K, V]](m M) bool {
//	    return len(m) == 0
//	}
type Map[K comparable, V any] interface {
	~map[K]V
}

// Chan represents any channel type.
// Useful for functions that work with channels regardless of direction.
//
// Example:
//
//	func Send[T any, C lxconstraints.Chan[T]](ch C, value T) {
//	    ch <- value
//	}
type Chan[T any] interface {
	~chan T | ~chan<- T | ~<-chan T
}

// Pointer represents any pointer type.
// Useful for functions that need to work with pointers generically.
//
// Example:
//
//	func IsNil[T any, P lxconstraints.Pointer[T]](p P) bool {
//	    return p == nil
//	}
type Pointer[T any] interface {
	*T
}
