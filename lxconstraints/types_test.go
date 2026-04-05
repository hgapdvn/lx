package lxconstraints_test

import (
	"testing"

	"github.com/hgapdvn/lx/lxconstraints"
)

// Helper functions for testing constraints

func integerAdd[T lxconstraints.Integer](v T) T {
	return v + 1
}

func signedSub[T lxconstraints.Signed](v T) T {
	return v - 1
}

func unsignedAdd[T lxconstraints.Unsigned](v T) T {
	return v + 1
}

func floatMul[T lxconstraints.Float](v T) T {
	return v * 2
}

func complexDouble[T lxconstraints.Complex](v T) T {
	return v + v
}

func testSum[T lxconstraints.Number](a, b T) T {
	return a + b
}

func testDouble[T lxconstraints.Numeric](v T) T {
	return v + v
}

func testMin[T lxconstraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func testAdd[T lxconstraints.Addable](a, b T) T {
	return a + b
}

func sliceLength[T any, S lxconstraints.Slice[T]](s S) int {
	return len(s)
}

func mapSize[K comparable, V any, M lxconstraints.Map[K, V]](m M) int {
	return len(m)
}

func pointerIsNil[T any, P lxconstraints.Pointer[T]](p P) bool {
	return p == nil
}

// TestInteger verifies Integer constraint accepts all integer types
func TestInteger(t *testing.T) {
	if got := integerAdd(int(5)); got != 6 {
		t.Errorf("int: expected 6, got %v", got)
	}
	if got := integerAdd(int64(5)); got != 6 {
		t.Errorf("int64: expected 6, got %v", got)
	}
	if got := integerAdd(uint(5)); got != 6 {
		t.Errorf("uint: expected 6, got %v", got)
	}
	if got := integerAdd(uint64(5)); got != 6 {
		t.Errorf("uint64: expected 6, got %v", got)
	}
}

// TestSigned verifies Signed constraint accepts only signed integers
func TestSigned(t *testing.T) {
	if got := signedSub(int(5)); got != 4 {
		t.Errorf("int: expected 4, got %v", got)
	}
	if got := signedSub(int64(5)); got != 4 {
		t.Errorf("int64: expected 4, got %v", got)
	}
}

// TestUnsigned verifies Unsigned constraint accepts only unsigned integers
func TestUnsigned(t *testing.T) {
	if got := unsignedAdd(uint(5)); got != 6 {
		t.Errorf("uint: expected 6, got %v", got)
	}
	if got := unsignedAdd(uint64(5)); got != 6 {
		t.Errorf("uint64: expected 6, got %v", got)
	}
}

// TestFloat verifies Float constraint accepts float types
func TestFloat(t *testing.T) {
	if got := floatMul(float32(2.5)); got != 5.0 {
		t.Errorf("float32: expected 5.0, got %v", got)
	}
	if got := floatMul(float64(2.5)); got != 5.0 {
		t.Errorf("float64: expected 5.0, got %v", got)
	}
}

// TestComplex verifies Complex constraint accepts complex types
func TestComplex(t *testing.T) {
	if got := complexDouble(complex64(1 + 2i)); got != complex64(2+4i) {
		t.Errorf("complex64: expected (2+4i), got %v", got)
	}
	if got := complexDouble(complex128(1 + 2i)); got != complex128(2+4i) {
		t.Errorf("complex128: expected (2+4i), got %v", got)
	}
}

// TestNumber verifies Number constraint accepts integers and floats
func TestNumber(t *testing.T) {
	if got := testSum(5, 3); got != 8 {
		t.Errorf("int: expected 8, got %v", got)
	}
	if got := testSum(2.5, 1.5); got != 4.0 {
		t.Errorf("float64: expected 4.0, got %v", got)
	}
}

// TestNumeric verifies Numeric constraint accepts all numeric types including complex
func TestNumeric(t *testing.T) {
	if got := testDouble(5); got != 10 {
		t.Errorf("int: expected 10, got %v", got)
	}
	if got := testDouble(2.5); got != 5.0 {
		t.Errorf("float64: expected 5.0, got %v", got)
	}
	if got := testDouble(complex128(1 + 2i)); got != complex128(2+4i) {
		t.Errorf("complex128: expected (2+4i), got %v", got)
	}
}

// TestOrdered verifies Ordered constraint accepts numbers and strings
func TestOrdered(t *testing.T) {
	if got := testMin(5, 3); got != 3 {
		t.Errorf("int: expected 3, got %v", got)
	}
	if got := testMin(2.5, 3.5); got != 2.5 {
		t.Errorf("float64: expected 2.5, got %v", got)
	}
	if got := testMin("banana", "apple"); got != "apple" {
		t.Errorf("string: expected 'apple', got %v", got)
	}
}

// TestAddable verifies Addable constraint accepts numeric types and strings
func TestAddable(t *testing.T) {
	if got := testAdd(5, 3); got != 8 {
		t.Errorf("int: expected 8, got %v", got)
	}
	if got := testAdd(2.5, 1.5); got != 4.0 {
		t.Errorf("float64: expected 4.0, got %v", got)
	}
	if got := testAdd("hello", " world"); got != "hello world" {
		t.Errorf("string: expected 'hello world', got %v", got)
	}
	if got := testAdd(complex128(1+2i), complex128(3+4i)); got != complex128(4+6i) {
		t.Errorf("complex128: expected (4+6i), got %v", got)
	}
}

// TestSlice verifies Slice constraint works with slice types
func TestSlice(t *testing.T) {
	intSlice := []int{1, 2, 3}
	if got := sliceLength(intSlice); got != 3 {
		t.Errorf("[]int: expected length 3, got %d", got)
	}

	stringSlice := []string{"a", "b"}
	if got := sliceLength(stringSlice); got != 2 {
		t.Errorf("[]string: expected length 2, got %d", got)
	}

	emptySlice := []float64{}
	if got := sliceLength(emptySlice); got != 0 {
		t.Errorf("[]float64: expected length 0, got %d", got)
	}
}

// TestMap verifies Map constraint works with map types
func TestMap(t *testing.T) {
	intMap := map[string]int{"a": 1, "b": 2}
	if got := mapSize(intMap); got != 2 {
		t.Errorf("map[string]int: expected size 2, got %d", got)
	}

	emptyMap := map[int]string{}
	if got := mapSize(emptyMap); got != 0 {
		t.Errorf("map[int]string: expected size 0, got %d", got)
	}
}

// TestPointer verifies Pointer constraint works with pointer types
func TestPointer(t *testing.T) {
	var nilPtr *int
	if !pointerIsNil(nilPtr) {
		t.Error("expected nil pointer to be nil")
	}

	val := 42
	ptr := &val
	if pointerIsNil(ptr) {
		t.Error("expected non-nil pointer to not be nil")
	}
}

// TestNamedTypes verifies constraints work with named types
func TestNamedTypes(t *testing.T) {
	type MyInt int
	type MyFloat float64
	type MyString string

	var a, b MyInt = 5, 3
	if got := testSum(a, b); got != 8 {
		t.Errorf("MyInt: expected 8, got %v", got)
	}

	var c, d MyFloat = 2.5, 1.5
	if got := testSum(c, d); got != 4.0 {
		t.Errorf("MyFloat: expected 4.0, got %v", got)
	}

	var s1, s2 MyString = "banana", "apple"
	if got := testMin(s1, s2); got != "apple" {
		t.Errorf("MyString: expected 'apple', got %v", got)
	}
}
