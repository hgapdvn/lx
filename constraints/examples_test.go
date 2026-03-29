package lxconstraints_test

import (
	"fmt"

	"github.com/hgapdvn/lx/constraints"
)

// Helper functions for examples
func integerSum[T lxconstraints.Integer](a, b T) T {
	return a + b
}

func numberSum[T lxconstraints.Number](values []T) T {
	var total T
	for _, v := range values {
		total += v
	}
	return total
}

func orderedMin[T lxconstraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func addableAdd[T lxconstraints.Addable](a, b T) T {
	return a + b
}

func signedAbs[T lxconstraints.Signed](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func unsignedHalf[T lxconstraints.Unsigned](n T) T {
	return n / 2
}

func floatAverage[T lxconstraints.Float](values []T) T {
	var sum T
	for _, v := range values {
		sum += v
	}
	return sum / T(len(values))
}

func sliceIsEmpty[T any, S lxconstraints.Slice[T]](s S) bool {
	return len(s) == 0
}

func mapKeys[K comparable, V any, M lxconstraints.Map[K, V]](m M) []K {
	result := make([]K, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

func pointerDeref[T any, P lxconstraints.Pointer[T]](p P, defaultValue T) T {
	if p == nil {
		return defaultValue
	}
	return *p
}

func ExampleInteger() {
	// Sum function works with any integer type
	fmt.Println(integerSum(5, 3))                  // int
	fmt.Println(integerSum(int64(10), int64(20)))  // int64
	fmt.Println(integerSum(uint8(100), uint8(50))) // uint8
	// Output:
	// 8
	// 30
	// 150
}

func ExampleNumber() {
	// Generic sum function that works with all numeric types
	integers := []int{1, 2, 3, 4, 5}
	fmt.Println(numberSum(integers))

	floats := []float64{1.5, 2.5, 3.5}
	fmt.Println(numberSum(floats))
	// Output:
	// 15
	// 7.5
}

func ExampleOrdered() {
	// Generic min function that works with ordered types
	fmt.Println(orderedMin(5, 3))              // int
	fmt.Println(orderedMin(2.5, 3.7))          // float64
	fmt.Println(orderedMin("banana", "apple")) // string
	// Output:
	// 3
	// 2.5
	// apple
}

func ExampleAddable() {
	// Generic add function that works with addable types
	fmt.Println(addableAdd(10, 20))                       // int addition
	fmt.Println(addableAdd(1.5, 2.5))                     // float addition
	fmt.Println(addableAdd("Hello, ", "World!"))          // string concatenation
	fmt.Println(addableAdd(complex(1, 2), complex(3, 4))) // complex addition
	// Output:
	// 30
	// 4
	// Hello, World!
	// (4+6i)
}

func ExampleSigned() {
	// Function that requires signed integers (can be negative)
	fmt.Println(signedAbs(-5))
	fmt.Println(signedAbs(int32(10)))
	fmt.Println(signedAbs(int64(-42)))
	// Output:
	// 5
	// 10
	// 42
}

func ExampleUnsigned() {
	// Function that works only with unsigned integers
	fmt.Println(unsignedHalf(uint(100)))
	fmt.Println(unsignedHalf(uint64(200)))
	// Output:
	// 50
	// 100
}

func ExampleFloat() {
	// Function specific to floating point numbers
	values := []float64{10.0, 20.0, 30.0, 40.0}
	fmt.Println(floatAverage(values))
	// Output:
	// 25
}

func ExampleSlice() {
	// Generic function that works with any slice type
	fmt.Println(sliceIsEmpty([]int{}))
	fmt.Println(sliceIsEmpty([]int{1, 2, 3}))
	fmt.Println(sliceIsEmpty([]string{"hello"}))
	// Output:
	// true
	// false
	// false
}

func ExampleMap() {
	// Generic function that works with any map type
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println(len(mapKeys(m)))
	// Output:
	// 2
}

func ExamplePointer() {
	// Generic function that works with pointers
	val := 42
	ptr := &val
	var nilPtr *int

	fmt.Println(pointerDeref(ptr, 0))
	fmt.Println(pointerDeref(nilPtr, 99))
	// Output:
	// 42
	// 99
}
