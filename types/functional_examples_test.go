package lxtypes_test

import (
	"fmt"
	"strings"

	"github.com/hgapdvn/lx/types"
)

// ============================================================================
// Predicate Examples
// ============================================================================

func ExamplePredicate() {
	isPositive := lxtypes.Predicate[int](func(n int) bool {
		return n > 0
	})

	fmt.Println(isPositive(5))
	fmt.Println(isPositive(-3))
	// Output:
	// true
	// false
}

func ExamplePredicate_And() {
	isEven := lxtypes.Predicate[int](func(n int) bool { return n%2 == 0 })
	isPositive := lxtypes.Predicate[int](func(n int) bool { return n > 0 })

	isEvenAndPositive := isEven.And(isPositive)

	fmt.Println(isEvenAndPositive(4))
	fmt.Println(isEvenAndPositive(-2))
	// Output:
	// true
	// false
}

func ExamplePredicate_Or() {
	isEven := lxtypes.Predicate[int](func(n int) bool { return n%2 == 0 })
	isNegative := lxtypes.Predicate[int](func(n int) bool { return n < 0 })

	isEvenOrNegative := isEven.Or(isNegative)

	fmt.Println(isEvenOrNegative(4))  // even
	fmt.Println(isEvenOrNegative(-3)) // negative
	fmt.Println(isEvenOrNegative(3))  // neither
	// Output:
	// true
	// true
	// false
}

func ExamplePredicate_Negate() {
	isEven := lxtypes.Predicate[int](func(n int) bool { return n%2 == 0 })
	isOdd := isEven.Negate()

	fmt.Println(isOdd(3))
	fmt.Println(isOdd(4))
	// Output:
	// true
	// false
}

// ============================================================================
// BiPredicate Examples
// ============================================================================

func ExampleBiPredicate() {
	inRange := lxtypes.BiPredicate[int, int](func(value, max int) bool {
		return value >= 0 && value <= max
	})

	fmt.Println(inRange(5, 10))
	fmt.Println(inRange(15, 10))
	// Output:
	// true
	// false
}

func ExampleBiPredicate_And() {
	inRange := lxtypes.BiPredicate[int, int](func(value, max int) bool {
		return value >= 0 && value <= max
	})
	lessThan := lxtypes.BiPredicate[int, int](func(a, b int) bool {
		return a < b
	})

	// Both conditions must be true
	validAndLess := inRange.And(lessThan)

	fmt.Println(validAndLess(5, 10)) // in range AND less than
	fmt.Println(validAndLess(10, 5)) // in range but NOT less than
	// Output:
	// true
	// false
}

func ExampleBiPredicate_Or() {
	equals := lxtypes.BiPredicate[int, int](func(a, b int) bool {
		return a == b
	})
	bothEven := lxtypes.BiPredicate[int, int](func(a, b int) bool {
		return a%2 == 0 && b%2 == 0
	})

	// Either condition can be true
	equalsOrBothEven := equals.Or(bothEven)

	fmt.Println(equalsOrBothEven(5, 5)) // equal
	fmt.Println(equalsOrBothEven(4, 6)) // both even
	fmt.Println(equalsOrBothEven(3, 5)) // neither
	// Output:
	// true
	// true
	// false
}

func ExampleBiPredicate_Negate() {
	equals := lxtypes.BiPredicate[int, int](func(a, b int) bool {
		return a == b
	})

	notEquals := equals.Negate()

	fmt.Println(notEquals(5, 3))
	fmt.Println(notEquals(5, 5))
	// Output:
	// true
	// false
}

// ============================================================================
// Consumer Examples
// ============================================================================

func ExampleConsumer() {
	printStr := lxtypes.Consumer[string](func(s string) {
		fmt.Println(s)
	})

	printStr("Hello, World!")
	// Output:
	// Hello, World!
}

func ExampleConsumer_AndThen() {
	var results []string

	append1 := lxtypes.Consumer[string](func(s string) {
		results = append(results, s)
	})
	append2 := lxtypes.Consumer[string](func(s string) {
		results = append(results, strings.ToUpper(s))
	})

	combined := append1.AndThen(append2)
	combined("hello")

	fmt.Println(results)
	// Output:
	// [hello HELLO]
}

// ============================================================================
// BiConsumer Examples
// ============================================================================

func ExampleBiConsumer() {
	printPair := lxtypes.BiConsumer[string, int](func(label string, value int) {
		fmt.Printf("%s: %d\n", label, value)
	})

	printPair("Count", 42)
	// Output:
	// Count: 42
}

func ExampleBiConsumer_AndThen() {
	results := []string{}

	append1 := lxtypes.BiConsumer[string, string](func(a, b string) {
		results = append(results, a+b)
	})
	append2 := lxtypes.BiConsumer[string, string](func(a, b string) {
		results = append(results, b+a)
	})

	combined := append1.AndThen(append2)
	combined("A", "B")

	fmt.Println(results)
	// Output:
	// [AB BA]
}

// ============================================================================
// Function Examples
// ============================================================================

func ExampleFunction() {
	toString := lxtypes.Function[int, string](func(n int) string {
		return fmt.Sprintf("Number: %d", n)
	})

	fmt.Println(toString(42))
	// Output:
	// Number: 42
}

func ExampleFunction_AndThen() {
	// Create a Function
	double := lxtypes.Function[int, int](func(n int) int { return n * 2 })

	// Chain with AndThen
	addTen := func(n int) int { return n + 10 }
	doubleThenAddTen := double.AndThen(addTen)

	fmt.Println(doubleThenAddTen(5)) // (5 * 2) + 10 = 20
	// Output:
	// 20
}

func ExampleFunction_Compose() {
	// Create a Function
	double := lxtypes.Function[int, int](func(n int) int { return n * 2 })

	// Chain with Compose
	addOne := func(n int) int { return n + 1 }
	addOneThenDouble := double.Compose(addOne)

	fmt.Println(addOneThenDouble(5)) // (5 + 1) * 2 = 12
	// Output:
	// 12
}

// ============================================================================
// BiFunction Examples
// ============================================================================

func ExampleBiFunction() {
	concat := lxtypes.BiFunction[string, string, string](func(a, b string) string {
		return a + b
	})

	fmt.Println(concat("Hello, ", "World!"))
	// Output:
	// Hello, World!
}

func ExampleBiFunction_AndThen() {
	// BiFunction that adds two numbers
	add := lxtypes.BiFunction[int, int, int](func(a, b int) int {
		return a + b
	})

	// Chain with AndThen to double the result
	addThenDouble := add.AndThen(func(n int) int { return n * 2 })

	fmt.Println(addThenDouble(3, 4)) // (3 + 4) * 2 = 14
	// Output:
	// 14
}

// ============================================================================
// Supplier Examples
// ============================================================================

func ExampleSupplier() {
	counter := 0
	getNext := lxtypes.Supplier[int](func() int {
		counter++
		return counter
	})

	fmt.Println(getNext())
	fmt.Println(getNext())
	fmt.Println(getNext())
	// Output:
	// 1
	// 2
	// 3
}

// ============================================================================
// BinaryOperator Examples
// ============================================================================

func ExampleBinaryOperator() {
	sum := lxtypes.BinaryOperator[int](func(a, b int) int {
		return a + b
	})

	fmt.Println(sum(10, 20))
	// Output:
	// 30
}

// ============================================================================
// Comparator Examples
// ============================================================================

func ExampleComparator() {
	intComparator := lxtypes.Comparator[int](func(a, b int) int {
		return a - b
	})

	fmt.Println(intComparator(5, 10) < 0) // 5 < 10
	fmt.Println(intComparator(10, 5) > 0) // 10 > 5
	fmt.Println(intComparator(5, 5) == 0) // 5 == 5
	// Output:
	// true
	// true
	// true
}

func ExampleComparator_Reversed() {
	ascending := lxtypes.Comparator[int](func(a, b int) int {
		return a - b
	})
	descending := ascending.Reversed()

	fmt.Println(descending(5, 10) > 0) // 5 > 10 in descending order
	fmt.Println(descending(10, 5) < 0) // 10 < 5 in descending order
	// Output:
	// true
	// true
}

func ExampleComparator_ThenComparing() {
	type Person struct {
		Name string
		Age  int
	}

	byName := lxtypes.Comparator[Person](func(a, b Person) int {
		if a.Name < b.Name {
			return -1
		}
		if a.Name > b.Name {
			return 1
		}
		return 0
	})

	byAge := lxtypes.Comparator[Person](func(a, b Person) int {
		return a.Age - b.Age
	})

	byNameThenAge := byName.ThenComparing(byAge)

	alice25 := Person{"Alice", 25}
	alice30 := Person{"Alice", 30}
	bob30 := Person{"Bob", 30}

	fmt.Println(byNameThenAge(alice25, alice30) < 0) // Alice 25 < Alice 30
	fmt.Println(byNameThenAge(alice30, bob30) < 0)   // Alice < Bob
	// Output:
	// true
	// true
}
