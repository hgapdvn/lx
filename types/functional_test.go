package lxtypes_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/types"
)

// Predicate Tests
func TestPredicate(t *testing.T) {
	isEven := lxtypes.Predicate[int](func(n int) bool {
		return n%2 == 0
	})

	if !isEven(4) {
		t.Error("Expected 4 to be even")
	}
	if isEven(3) {
		t.Error("Expected 3 to be odd")
	}
}

func TestPredicateAnd(t *testing.T) {
	isEven := lxtypes.Predicate[int](func(n int) bool { return n%2 == 0 })
	isPositive := lxtypes.Predicate[int](func(n int) bool { return n > 0 })
	isEvenAndPositive := isEven.And(isPositive)

	tests := []struct {
		input int
		want  bool
	}{
		{4, true},
		{-2, false},
		{3, false},
		{-3, false},
	}

	for _, tt := range tests {
		if got := isEvenAndPositive(tt.input); got != tt.want {
			t.Errorf("isEvenAndPositive(%d) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestPredicateOr(t *testing.T) {
	isEven := lxtypes.Predicate[int](func(n int) bool { return n%2 == 0 })
	isNegative := lxtypes.Predicate[int](func(n int) bool { return n < 0 })

	isEvenOrNegative := isEven.Or(isNegative)

	tests := []struct {
		input int
		want  bool
	}{
		{4, true},
		{-3, true},
		{-2, true},
		{3, false},
	}

	for _, tt := range tests {
		if got := isEvenOrNegative(tt.input); got != tt.want {
			t.Errorf("isEvenOrNegative(%d) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestPredicateNegate(t *testing.T) {
	isEven := lxtypes.Predicate[int](func(n int) bool { return n%2 == 0 })
	isOdd := isEven.Negate()

	if !isOdd(3) {
		t.Error("Expected 3 to be odd")
	}
	if isOdd(4) {
		t.Error("Expected 4 to be even")
	}
}

// BiPredicate Tests
func TestBiPredicate(t *testing.T) {
	equals := lxtypes.BiPredicate[int, int](func(a, b int) bool {
		return a == b
	})

	if !equals(5, 5) {
		t.Error("Expected 5 to equal 5")
	}
	if equals(5, 3) {
		t.Error("Expected 5 not to equal 3")
	}
}

func TestBiPredicateAnd(t *testing.T) {
	inRange := lxtypes.BiPredicate[int, int](func(value, max int) bool {
		return value >= 0 && value <= max
	})
	lessThan := lxtypes.BiPredicate[int, int](func(a, b int) bool {
		return a < b
	})

	// Both conditions must be true
	combined := inRange.And(lessThan)

	tests := []struct {
		name string
		a, b int
		want bool
	}{
		{"valid in range and less", 5, 10, true},
		{"not less than", 10, 5, false},
		{"negative value", -5, 10, false},
		{"equal values", 5, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combined(tt.a, tt.b); got != tt.want {
				t.Errorf("combined(%d, %d) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestBiPredicateOr(t *testing.T) {
	equals := lxtypes.BiPredicate[int, int](func(a, b int) bool {
		return a == b
	})
	bothEven := lxtypes.BiPredicate[int, int](func(a, b int) bool {
		return a%2 == 0 && b%2 == 0
	})

	// Either condition can be true
	combined := equals.Or(bothEven)

	tests := []struct {
		name string
		a, b int
		want bool
	}{
		{"equal", 5, 5, true},
		{"both even", 4, 6, true},
		{"equal and even", 4, 4, true},
		{"neither", 3, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combined(tt.a, tt.b); got != tt.want {
				t.Errorf("combined(%d, %d) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestBiPredicateNegate(t *testing.T) {
	equals := lxtypes.BiPredicate[int, int](func(a, b int) bool {
		return a == b
	})

	notEquals := equals.Negate()

	tests := []struct {
		name string
		a, b int
		want bool
	}{
		{"equal values", 5, 5, false},
		{"different values", 5, 3, true},
		{"negative equal", -5, -5, false},
		{"negative different", -5, 5, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := notEquals(tt.a, tt.b); got != tt.want {
				t.Errorf("notEquals(%d, %d) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// BiConsumer Tests
func TestBiConsumer(t *testing.T) {
	result := ""
	concat := lxtypes.BiConsumer[string, string](func(a, b string) {
		result = a + b
	})

	concat("Hello, ", "World!")
	if result != "Hello, World!" {
		t.Errorf("Expected 'Hello, World!', got %q", result)
	}
}

func TestBiConsumerAndThen(t *testing.T) {
	results := []string{}
	append1 := lxtypes.BiConsumer[string, string](func(a, b string) {
		results = append(results, a+b)
	})
	append2 := lxtypes.BiConsumer[string, string](func(a, b string) {
		results = append(results, b+a)
	})

	combined := append1.AndThen(append2)
	combined("A", "B")

	if len(results) != 2 || results[0] != "AB" || results[1] != "BA" {
		t.Errorf("Expected [AB, BA], got %v", results)
	}
}

// Consumer Tests
func TestConsumer(t *testing.T) {
	sum := 0
	addToSum := lxtypes.Consumer[int](func(n int) {
		sum += n
	})

	addToSum(5)
	if sum != 5 {
		t.Errorf("Expected sum to be 5, got %d", sum)
	}
}

func TestConsumerAndThen(t *testing.T) {
	results := []int{}
	append1 := lxtypes.Consumer[int](func(n int) {
		results = append(results, n)
	})
	append2 := lxtypes.Consumer[int](func(n int) {
		results = append(results, n*2)
	})

	combined := append1.AndThen(append2)
	combined(5)

	if len(results) != 2 || results[0] != 5 || results[1] != 10 {
		t.Errorf("Expected [5, 10], got %v", results)
	}
}

// Function Tests
func TestFunction(t *testing.T) {
	double := lxtypes.Function[int, int](func(n int) int {
		return n * 2
	})

	if got := double(5); got != 10 {
		t.Errorf("Expected 10, got %d", got)
	}
}

func TestFunctionAndThen(t *testing.T) {
	// Test AndThen method on Function type
	double := lxtypes.Function[int, int](func(n int) int { return n * 2 })
	addTen := func(n int) int { return n + 10 }

	// AndThen: double first, then addTen
	// double.AndThen(addTen)(5) = addTen(double(5)) = addTen(10) = 20
	doubleThenAddTen := double.AndThen(addTen)

	if got := doubleThenAddTen(5); got != 20 {
		t.Errorf("Expected 20, got %d", got)
	}

	// Chain multiple AndThen
	multiplyThree := func(n int) int { return n * 3 }
	chained := double.AndThen(addTen).AndThen(multiplyThree)
	// (5 * 2) = 10, (10 + 10) = 20, (20 * 3) = 60

	if got := chained(5); got != 60 {
		t.Errorf("Expected 60, got %d", got)
	}
}

func TestFunctionCompose(t *testing.T) {
	// Test Compose method on Function type
	double := lxtypes.Function[int, int](func(n int) int { return n * 2 })
	addOne := func(n int) int { return n + 1 }

	// Compose: addOne first, then double
	// double.Compose(addOne)(5) = double(addOne(5)) = double(6) = 12
	addOneThenDouble := double.Compose(addOne)

	if got := addOneThenDouble(5); got != 12 {
		t.Errorf("Expected 12, got %d", got)
	}

	// Chain multiple Compose
	subtractTwo := func(n int) int { return n - 2 }
	chained := double.Compose(addOne).Compose(subtractTwo)
	// (5 - 2) = 3, (3 + 1) = 4, (4 * 2) = 8

	if got := chained(5); got != 8 {
		t.Errorf("Expected 8, got %d", got)
	}
}

// BiFunction Tests
func TestBiFunction(t *testing.T) {
	add := lxtypes.BiFunction[int, int, int](func(a, b int) int {
		return a + b
	})

	if got := add(3, 4); got != 7 {
		t.Errorf("Expected 7, got %d", got)
	}
}

func TestBiFunctionAndThen(t *testing.T) {
	// BiFunction that adds two numbers
	add := lxtypes.BiFunction[int, int, int](func(a, b int) int {
		return a + b
	})

	// Chain with AndThen to double the result
	double := func(n int) int { return n * 2 }
	addThenDouble := add.AndThen(double)

	if got := addThenDouble(3, 4); got != 14 {
		t.Errorf("Expected 14, got %d", got) // (3 + 4) * 2 = 14
	}

	// Chain multiple AndThen operations
	addTen := func(n int) int { return n + 10 }
	chained := add.AndThen(double).AndThen(addTen)

	if got := chained(3, 4); got != 24 {
		t.Errorf("Expected 24, got %d", got) // (3 + 4) * 2 + 10 = 24
	}
}

// Supplier Tests
func TestSupplier(t *testing.T) {
	counter := 0
	incrementingSupplier := lxtypes.Supplier[int](func() int {
		counter++
		return counter
	})

	if got := incrementingSupplier(); got != 1 {
		t.Errorf("Expected 1, got %d", got)
	}
	if got := incrementingSupplier(); got != 2 {
		t.Errorf("Expected 2, got %d", got)
	}
}

// UnaryOperator Tests
func TestUnaryOperator(t *testing.T) {
	square := lxtypes.UnaryOperator[int](func(n int) int {
		return n * n
	})

	if got := square(5); got != 25 {
		t.Errorf("Expected 25, got %d", got)
	}
}

// BinaryOperator Tests
func TestBinaryOperator(t *testing.T) {
	multiply := lxtypes.BinaryOperator[int](func(a, b int) int {
		return a * b
	})

	if got := multiply(3, 4); got != 12 {
		t.Errorf("Expected 12, got %d", got)
	}
}

// Comparator Tests
func TestComparator(t *testing.T) {
	intComparator := lxtypes.Comparator[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	if got := intComparator(3, 5); got >= 0 {
		t.Errorf("Expected negative, got %d", got)
	}
	if got := intComparator(5, 3); got <= 0 {
		t.Errorf("Expected positive, got %d", got)
	}
	if got := intComparator(4, 4); got != 0 {
		t.Errorf("Expected 0, got %d", got)
	}
}

func TestComparatorReversed(t *testing.T) {
	intComparator := lxtypes.Comparator[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	reversed := intComparator.Reversed()

	if got := reversed(3, 5); got <= 0 {
		t.Errorf("Expected positive, got %d", got)
	}
	if got := reversed(5, 3); got >= 0 {
		t.Errorf("Expected negative, got %d", got)
	}
}

func TestComparatorThenComparing(t *testing.T) {
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

	p1 := Person{"Alice", 30}
	p2 := Person{"Alice", 25}
	p3 := Person{"Bob", 30}

	if got := byNameThenAge(p1, p2); got <= 0 {
		t.Errorf("Expected positive (Alice 30 > Alice 25), got %d", got)
	}
	if got := byNameThenAge(p1, p3); got >= 0 {
		t.Errorf("Expected negative (Alice < Bob), got %d", got)
	}
}
