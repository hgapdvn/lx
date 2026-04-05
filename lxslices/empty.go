package lxslices

// Empty creates and returns an empty slice of the specified generic type T.
func Empty[T any]() []T {
	return []T{}
}

// EmptyInt creates and returns an empty slice of int.
func EmptyInt() []int {
	return Empty[int]()
}

// EmptyInt8 creates and returns an empty slice of int8.
func EmptyInt8() []int8 {
	return Empty[int8]()
}

// EmptyInt16 creates and returns an empty slice of int16.
func EmptyInt16() []int16 {
	return Empty[int16]()
}

// EmptyInt32 creates and returns an empty slice of int32.
func EmptyInt32() []int32 {
	return Empty[int32]()
}

// EmptyInt64 creates and returns an empty slice of int64.
func EmptyInt64() []int64 {
	return Empty[int64]()
}

// EmptyUint creates and returns an empty slice of uint.
func EmptyUint() []uint {
	return Empty[uint]()
}

// EmptyUint8 creates and returns an empty slice of uint8.
func EmptyUint8() []uint8 {
	return Empty[uint8]()
}

// EmptyUint16 creates and returns an empty slice of uint16.
func EmptyUint16() []uint16 {
	return Empty[uint16]()
}

// EmptyUint32 creates and returns an empty slice of uint32.
func EmptyUint32() []uint32 {
	return Empty[uint32]()
}

// EmptyUint64 creates and returns an empty slice of uint64.
func EmptyUint64() []uint64 {
	return Empty[uint64]()
}

// EmptyFloat32 creates and returns an empty slice of float32.
func EmptyFloat32() []float32 {
	return Empty[float32]()
}

// EmptyFloat64 creates and returns an empty slice of float64.
func EmptyFloat64() []float64 {
	return Empty[float64]()
}

// EmptyBool creates and returns an empty slice of bool.
func EmptyBool() []bool {
	return Empty[bool]()
}

// EmptyRune creates and returns an empty slice of rune.
func EmptyRune() []rune {
	return Empty[rune]()
}

// EmptyByte creates and returns an empty slice of byte.
func EmptyByte() []byte {
	return Empty[byte]()
}

// EmptyComplex64 creates and returns an empty slice of complex64.
func EmptyComplex64() []complex64 {
	return Empty[complex64]()
}

// EmptyComplex128 creates and returns an empty slice of complex128.
func EmptyComplex128() []complex128 {
	return Empty[complex128]()
}

// EmptyError creates and returns an empty slice of error.
func EmptyError() []error {
	return Empty[error]()
}

// EmptyString creates and returns an empty slice of string.
func EmptyString() []string {
	return Empty[string]()
}

// IsEmpty returns true if the slice is nil or has zero length.
func IsEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

// IsNotEmpty returns true if the slice has at least one element.
func IsNotEmpty[T any](slice []T) bool {
	return !IsEmpty(slice)
}

// Length returns the number of elements in the slice.
func Length[T any](slice []T) int {
	return len(slice)
}
