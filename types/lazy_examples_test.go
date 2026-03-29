package lxtypes_test

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hgapdvn/lx/types"
)

// ExampleLazyEager demonstrates creating a Lazy with an immediate value.
func ExampleLazyEager() {
	// Create a lazy wrapper around an immediate value
	lazy := lxtypes.LazyEager(42)

	// Check if it's already evaluated
	fmt.Println("Is evaluated:", lazy.IsEvaluated())

	// Get the value (returns immediately)
	value, err := lazy.Get()
	fmt.Println("Value:", value)
	fmt.Println("Error:", err)

	// Output:
	// Is evaluated: true
	// Value: 42
	// Error: <nil>
}

// ExampleLazyEager_string demonstrates LazyEager with string type.
func ExampleLazyEager_string() {
	lazy := lxtypes.LazyEager("hello world")

	value := lazy.MustGet()
	fmt.Println(value)

	// Output:
	// hello world
}

// ExampleLazyEagerOrError demonstrates wrapping an existing result with error.
func ExampleLazyEagerOrError() {
	// Simulate a function that returns (value, error)
	computeValue := func() (string, error) {
		return "success", nil
	}

	// Wrap the result in a Lazy
	result, err := computeValue()
	lazy := lxtypes.LazyEagerOrError(result, err)

	value, err := lazy.Get()
	fmt.Println("Value:", value)
	fmt.Println("Error:", err)

	// Output:
	// Value: success
	// Error: <nil>
}

// ExampleLazyEagerOrError_withError demonstrates wrapping an error result.
func ExampleLazyEagerOrError_withError() {
	// Simulate a function that returns an error
	computeValue := func() (int, error) {
		return 0, errors.New("computation failed")
	}

	// Wrap the error result in a Lazy
	result, err := computeValue()
	lazy := lxtypes.LazyEagerOrError(result, err)

	_, err = lazy.Get()
	fmt.Println("Error:", err)

	// Output:
	// Error: computation failed
}

// ExampleLazyDeferred demonstrates deferred computation.
func ExampleLazyDeferred() {
	// Create a lazy computation
	lazy := lxtypes.LazyDeferred(func() (int, error) {
		// Expensive computation
		return 42, nil
	})

	// Check if it's evaluated before first access
	fmt.Println("Is evaluated before Get:", lazy.IsEvaluated())

	// Compute the value on first call
	value, err := lazy.Get()
	fmt.Println("Value:", value)
	fmt.Println("Error:", err)

	// Check if it's evaluated after Get
	fmt.Println("Is evaluated after Get:", lazy.IsEvaluated())

	// Second call returns cached value
	value2, _ := lazy.Get()
	fmt.Println("Cached value:", value2)

	// Output:
	// Is evaluated before Get: false
	// Value: 42
	// Error: <nil>
	// Is evaluated after Get: true
	// Cached value: 42
}

// ExampleLazyDeferred_expensiveComputation demonstrates lazy evaluation benefit.
func ExampleLazyDeferred_expensiveComputation() {
	callCount := 0

	// Create a lazy computation that tracks how many times it's called
	lazy := lxtypes.LazyDeferred(func() (string, error) {
		callCount++
		// Simulate expensive operation
		time.Sleep(10 * time.Millisecond)
		return "computed", nil
	})

	fmt.Println("Call count before:", callCount)

	// First call - computation happens
	value1, _ := lazy.Get()
	fmt.Println("First call:", value1)
	fmt.Println("Call count after first:", callCount)

	// Second call - returns cached value
	value2, _ := lazy.Get()
	fmt.Println("Second call:", value2)
	fmt.Println("Call count after second:", callCount)

	// Output:
	// Call count before: 0
	// First call: computed
	// Call count after first: 1
	// Second call: computed
	// Call count after second: 1
}

// ExampleLazyDeferred_withError demonstrates error handling in deferred computation.
func ExampleLazyDeferred_withError() {
	lazy := lxtypes.LazyDeferred(func() (int, error) {
		return 0, errors.New("computation failed")
	})

	// First call returns error
	_, err := lazy.Get()
	fmt.Println("First error:", err)

	// Second call returns the same cached error
	_, err = lazy.Get()
	fmt.Println("Second error:", err)

	// IsEvaluated is still true even though computation failed
	fmt.Println("Is evaluated:", lazy.IsEvaluated())

	// Output:
	// First error: computation failed
	// Second error: computation failed
	// Is evaluated: true
}

// ExampleLazy_MustGet demonstrates the MustGet panic behavior.
func ExampleLazy_MustGet() {
	// Success case
	success := lxtypes.LazyEager(100)
	value := success.MustGet()
	fmt.Println("Success value:", value)

	// Error case (would panic in real code)
	// failed := lxtypes.LazyEagerOrError(0, errors.New("failed"))
	// failed.MustGet() // This would panic!

	// Output:
	// Success value: 100
}

// ExampleLazy_unifiedInterface demonstrates treating eager and deferred lazies uniformly.
func ExampleLazy_unifiedInterface() {
	// Function that works with any Lazy
	process := func(name string, lazy lxtypes.Lazy[string]) {
		value, err := lazy.Get()
		if err != nil {
			fmt.Printf("%s: error - %v\n", name, err)
		} else {
			fmt.Printf("%s: %s\n", name, value)
		}
	}

	// Works with eager values
	eager := lxtypes.LazyEager("immediate")
	process("eager", eager)

	// Works with deferred computations
	deferred := lxtypes.LazyDeferred(func() (string, error) {
		return "computed", nil
	})
	process("deferred", deferred)

	// Output:
	// eager: immediate
	// deferred: computed
}

// ExampleLazyDeferred_databaseConnection demonstrates a realistic use case.
func ExampleLazyDeferred_databaseConnection() {
	// Simulate expensive database connection that's only created when needed
	dbConnection := lxtypes.LazyDeferred(func() (*mockDB, error) {
		// Simulate connection delay
		time.Sleep(5 * time.Millisecond)
		return &mockDB{name: "production-db"}, nil
	})

	fmt.Println("DB connection created:", dbConnection.IsEvaluated())

	// Connection only happens when actually needed
	db, err := dbConnection.Get()
	if err != nil {
		fmt.Println("Connection failed:", err)
		return
	}

	fmt.Println("Connected to:", db.name)
	fmt.Println("DB connection created:", dbConnection.IsEvaluated())

	// Output:
	// DB connection created: false
	// Connected to: production-db
	// DB connection created: true
}

// ExampleLazyDeferred_transform demonstrates transforming lazy values.
func ExampleLazyDeferred_transform() {
	// Create a lazy computation
	lazy := lxtypes.LazyDeferred(func() (string, error) {
		return "hello", nil
	})

	// Transform the result (this example shows manual transformation)
	value, err := lazy.Get()
	if err == nil {
		transformed := strings.ToUpper(value)
		fmt.Println(transformed)
	}

	// Output:
	// HELLO
}

// ExampleLazyDeferred_conditionalEvaluation demonstrates lazy evaluation benefit.
func ExampleLazyDeferred_conditionalEvaluation() {
	expensiveComputation := lxtypes.LazyDeferred(func() (int, error) {
		fmt.Println("Computing...")
		return 42, nil
	})

	condition := false
	if condition {
		// This branch never executes, so computation never happens
		value := expensiveComputation.MustGet()
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Skipped expensive computation")
	}

	fmt.Println("Was computed:", expensiveComputation.IsEvaluated())

	// Output:
	// Skipped expensive computation
	// Was computed: false
}

// ExampleLazyDeferred_multipleTypes demonstrates Lazy with different types.
func ExampleLazyDeferred_multipleTypes() {
	// Integer lazy
	intLazy := lxtypes.LazyDeferred(func() (int, error) {
		return 42, nil
	})

	// String lazy
	strLazy := lxtypes.LazyDeferred(func() (string, error) {
		return "lazy string", nil
	})

	// Struct lazy
	type Person struct {
		Name string
		Age  int
	}
	personLazy := lxtypes.LazyDeferred(func() (Person, error) {
		return Person{Name: "Alice", Age: 30}, nil
	})

	intVal, intErr := intLazy.Get()
	strVal, strErr := strLazy.Get()
	personVal, personErr := personLazy.Get()

	if intErr == nil {
		fmt.Println("Int:", intVal)
	}
	if strErr == nil {
		fmt.Println("String:", strVal)
	}
	if personErr == nil {
		fmt.Printf("Person: %s, %d\n", personVal.Name, personVal.Age)
	}

	// Output:
	// Int: 42
	// String: lazy string
	// Person: Alice, 30
}

// mockDB is a helper type for examples
type mockDB struct {
	name string
}
