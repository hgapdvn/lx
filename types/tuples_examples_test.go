package lxtypes_test

import (
	"fmt"
	"strings"

	"github.com/hgapdvn/lx/types"
)

// ============================================================================
// Pair Examples
// ============================================================================

func ExamplePair() {
	// Creating a pair
	p := lxtypes.NewPair(42, "answer")
	fmt.Printf("First: %d, Second: %s\n", p.First, p.Second)

	// Unpacking values
	x, y := p.Values()
	fmt.Printf("Unpacked: %d, %s\n", x, y)
	// Output:
	// First: 42, Second: answer
	// Unpacked: 42, answer
}

func ExamplePair_Swap() {
	p := lxtypes.NewPair(42, "answer")
	swapped := p.Swap()
	fmt.Printf("Original: (%d, %s)\n", p.First, p.Second)
	fmt.Printf("Swapped: (%s, %d)\n", swapped.First, swapped.Second)
	// Output:
	// Original: (42, answer)
	// Swapped: (answer, 42)
}

func ExamplePair_MapFirst() {
	p := lxtypes.NewPair(5, "test")
	doubled := p.MapFirst(func(n int) int { return n * 2 })
	fmt.Printf("Original: (%d, %s)\n", p.First, p.Second)
	fmt.Printf("After MapFirst: (%d, %s)\n", doubled.First, doubled.Second)
	// Output:
	// Original: (5, test)
	// After MapFirst: (10, test)
}

func ExamplePair_MapSecond() {
	p := lxtypes.NewPair(5, "test")
	upper := p.MapSecond(func(s string) string { return strings.ToUpper(s) })
	fmt.Printf("Original: (%d, %s)\n", p.First, p.Second)
	fmt.Printf("After MapSecond: (%d, %s)\n", upper.First, upper.Second)
	// Output:
	// Original: (5, test)
	// After MapSecond: (5, TEST)
}

// ============================================================================
// Triple Examples
// ============================================================================

func ExampleTriple() {
	// Creating a triple
	t := lxtypes.NewTriple(1, "hello", true)
	fmt.Printf("First: %d, Second: %s, Third: %t\n", t.First, t.Second, t.Third)

	// Unpacking values
	x, y, z := t.Values()
	fmt.Printf("Unpacked: %d, %s, %t\n", x, y, z)
	// Output:
	// First: 1, Second: hello, Third: true
	// Unpacked: 1, hello, true
}

func ExampleTriple_ToPair() {
	t := lxtypes.NewTriple(1, "hello", true)
	p := t.ToPair()
	fmt.Printf("Triple: (%d, %s, %t)\n", t.First, t.Second, t.Third)
	fmt.Printf("Pair: (%d, %s)\n", p.First, p.Second)
	// Output:
	// Triple: (1, hello, true)
	// Pair: (1, hello)
}

// ============================================================================
// Quad Examples
// ============================================================================

func ExampleQuad() {
	// Creating a quad
	q := lxtypes.NewQuad(1, "hello", true, 3.14)
	fmt.Printf("Values: %d, %s, %t, %.2f\n", q.First, q.Second, q.Third, q.Fourth)

	// Unpacking values
	w, x, y, z := q.Values()
	fmt.Printf("Unpacked: %d, %s, %t, %.2f\n", w, x, y, z)
	// Output:
	// Values: 1, hello, true, 3.14
	// Unpacked: 1, hello, true, 3.14
}

func ExampleQuad_ToTriple() {
	q := lxtypes.NewQuad(1, "hello", true, 3.14)
	t := q.ToTriple()
	fmt.Printf("Quad: (%d, %s, %t, %.2f)\n", q.First, q.Second, q.Third, q.Fourth)
	fmt.Printf("Triple: (%d, %s, %t)\n", t.First, t.Second, t.Third)
	// Output:
	// Quad: (1, hello, true, 3.14)
	// Triple: (1, hello, true)
}

func ExampleQuad_ToPair() {
	q := lxtypes.NewQuad(1, "hello", true, 3.14)
	p := q.ToPair()
	fmt.Printf("Quad: (%d, %s, %t, %.2f)\n", q.First, q.Second, q.Third, q.Fourth)
	fmt.Printf("Pair: (%d, %s)\n", p.First, p.Second)
	// Output:
	// Quad: (1, hello, true, 3.14)
	// Pair: (1, hello)
}

func ExampleTuple5() {
	// Creating a tuple with 5 different types
	t := lxtypes.NewTuple5(1, "two", true, 4.0, []int{5, 6})
	fmt.Printf("Tuple5: (%d, %s, %t, %.1f, %v)\n", t.V1, t.V2, t.V3, t.V4, t.V5)

	// Unpacking values
	v1, v2, v3, v4, v5 := t.Values()
	fmt.Printf("Unpacked: %d, %s, %t, %.1f, %v\n", v1, v2, v3, v4, v5)
	// Output:
	// Tuple5: (1, two, true, 4.0, [5 6])
	// Unpacked: 1, two, true, 4.0, [5 6]
}

func ExampleNewTuple5() {
	// Use case: Fetching from 5 different services
	type User struct{ Name string }
	type Config struct{ Host string }
	type Stats struct{ Count int }

	data := lxtypes.NewTuple5(
		User{"Alice"},
		[]string{"order1", "order2"},
		Config{"api.example.com"},
		Stats{100},
		map[string]int{"total": 42},
	)

	fmt.Printf("User: %s\n", data.V1.Name)
	fmt.Printf("Orders: %d\n", len(data.V2))
	fmt.Printf("Host: %s\n", data.V3.Host)
	// Output:
	// User: Alice
	// Orders: 2
	// Host: api.example.com
}

func ExampleTuple6() {
	// Creating a tuple with 6 different types
	t := lxtypes.NewTuple6(1, "two", true, 4.0, []int{5}, 'a')
	fmt.Printf("Values: %d, %s, %t, %.1f, %v, %c\n", t.V1, t.V2, t.V3, t.V4, t.V5, t.V6)

	// Unpacking
	v1, v2, v3, v4, v5, v6 := t.Values()
	fmt.Printf("Unpacked: %d, %s, %t, %.1f, %v, %c\n", v1, v2, v3, v4, v5, v6)
	// Output:
	// Values: 1, two, true, 4.0, [5], a
	// Unpacked: 1, two, true, 4.0, [5], a
}

func ExampleTuple7() {
	// Creating a tuple with 7 different types
	t := lxtypes.NewTuple7(1, "two", true, 4.0, []int{5}, 'a', byte(7))
	fmt.Printf("V1=%d, V2=%s, V3=%t, V4=%.1f, V5=%v, V6=%c, V7=%d\n",
		t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7)
	// Output:
	// V1=1, V2=two, V3=true, V4=4.0, V5=[5], V6=a, V7=7
}

func ExampleTuple8() {
	// Creating a tuple with 8 different types
	t := lxtypes.NewTuple8(1, "two", true, 4.0, []int{5}, 'a', byte(7), uint(8))
	fmt.Printf("All 8 values: %d, %s, %t, %.1f, %v, %c, %d, %d\n",
		t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8)
	// Output:
	// All 8 values: 1, two, true, 4.0, [5], a, 7, 8
}

func ExampleTuple8_Values() {
	t := lxtypes.NewTuple8(1, "two", true, 4.0, []int{5}, 'a', byte(7), uint(8))
	v1, v2, v3, v4, v5, v6, v7, v8 := t.Values()
	fmt.Printf("Unpacked: %d, %s, %t, %.1f, %v, %c, %d, %d\n",
		v1, v2, v3, v4, v5, v6, v7, v8)
	// Output:
	// Unpacked: 1, two, true, 4.0, [5], a, 7, 8
}
