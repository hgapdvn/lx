package lxtypes_test

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/nthanhhai2909/lx/types"
)

// Database represents a simple struct for examples
type Database struct {
	Name string
	Port int
}

func ExampleResultSuccess() {
	// Create a successful Result with a value
	result := lxtypes.ResultSuccess(42)

	// Use Go's idiomatic (value, error) pattern
	if value, err := result.Value(); err == nil {
		fmt.Println("Value:", value)
	}
	// Output:
	// Value: 42
}

func ExampleResultSuccess_struct() {
	// Create a successful Result with a struct
	db := Database{Name: "users", Port: 5432}
	result := lxtypes.ResultSuccess(db)

	if value, err := result.Value(); err == nil {
		fmt.Printf("Database: %s on port %d\n", value.Name, value.Port)
	}
	// Output:
	// Database: users on port 5432
}

func ExampleResultSuccess_pointerStruct() {
	// Create a successful Result with a pointer to struct
	db := &Database{Name: "orders", Port: 3306}
	result := lxtypes.ResultSuccess(db)

	if value, err := result.Value(); err == nil {
		fmt.Printf("Database: %s\n", value.Name)
	}
	// Output:
	// Database: orders
}

func ExampleResultFailure() {
	// Create a failed Result with an error
	result := lxtypes.ResultFailure[int](errors.New("operation failed"))

	// Check for error
	if _, err := result.Value(); err != nil {
		fmt.Println("Error:", err)
	}
	// Output:
	// Error: operation failed
}

func ExampleResultFailure_struct() {
	// Create a failed Result for a struct type
	result := lxtypes.ResultFailure[Database](errors.New("connection failed"))

	if _, err := result.Value(); err != nil {
		fmt.Println("Error:", err)
	}
	// Output:
	// Error: connection failed
}

func ExampleResult_Value() {
	// Success case
	success := lxtypes.ResultSuccess(42)
	value, err := success.Value()
	if err == nil {
		fmt.Println("Success:", value)
	}

	// Failure case
	failure := lxtypes.ResultFailure[int](errors.New("error"))
	_, err = failure.Value()
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Output:
	// Success: 42
	// Error: error
}

func ExampleResult_Value_struct() {
	// Success with struct
	db := Database{Name: "cache", Port: 6379}
	success := lxtypes.ResultSuccess(db)

	if value, err := success.Value(); err == nil {
		fmt.Printf("Connected to %s\n", value.Name)
	}

	// Failure with struct
	failure := lxtypes.ResultFailure[Database](errors.New("timeout"))
	if _, err := failure.Value(); err != nil {
		fmt.Println("Connection failed")
	}
	// Output:
	// Connected to cache
	// Connection failed
}

func ExampleResult_ValueOr() {
	// Success returns original value
	success := lxtypes.ResultSuccess(42)
	fmt.Println(success.ValueOr(0))

	// Failure returns default value
	failure := lxtypes.ResultFailure[int](errors.New("error"))
	fmt.Println(failure.ValueOr(99))
	// Output:
	// 42
	// 99
}

func ExampleResult_ValueOr_struct() {
	// Success with struct
	db := Database{Name: "prod", Port: 5432}
	success := lxtypes.ResultSuccess(db)
	defaultDB := Database{Name: "local", Port: 5432}

	result := success.ValueOr(defaultDB)
	fmt.Println(result.Name)

	// Failure with struct
	failure := lxtypes.ResultFailure[Database](errors.New("error"))
	result2 := failure.ValueOr(defaultDB)
	fmt.Println(result2.Name)
	// Output:
	// prod
	// local
}

func ExampleResult_ValueOr_pointerStruct() {
	// Success with pointer struct
	db := &Database{Name: "prod", Port: 5432}
	success := lxtypes.ResultSuccess(db)
	defaultDB := &Database{Name: "local", Port: 5432}

	result := success.ValueOr(defaultDB)
	fmt.Println(result.Name)

	// Failure with pointer struct
	failure := lxtypes.ResultFailure[*Database](errors.New("error"))
	result2 := failure.ValueOr(defaultDB)
	fmt.Println(result2.Name)
	// Output:
	// prod
	// local
}

// Example showing conversion from Go's (value, error) pattern
func ExampleResult_fromError() {
	// Parse a string to int
	value, err := strconv.Atoi("42")

	// Convert to Result
	var result lxtypes.Result[int]
	if err != nil {
		result = lxtypes.ResultFailure[int](err)
	} else {
		result = lxtypes.ResultSuccess(value)
	}

	// Use the Result
	if v, e := result.Value(); e == nil {
		fmt.Println("Parsed:", v)
	}
	// Output:
	// Parsed: 42
}

// Example showing a practical use case: database operations
func ExampleResult_databaseOperation() {
	// Simulate a database query
	findDatabase := func(name string) lxtypes.Result[Database] {
		if name == "users" {
			return lxtypes.ResultSuccess(Database{Name: "users", Port: 5432})
		}
		return lxtypes.ResultFailure[Database](errors.New("database not found"))
	}

	// Success case
	result := findDatabase("users")
	if db, err := result.Value(); err == nil {
		fmt.Printf("Found: %s on port %d\n", db.Name, db.Port)
	}

	// Failure case with default
	result2 := findDatabase("missing")
	defaultDB := Database{Name: "default", Port: 5432}
	db := result2.ValueOr(defaultDB)
	fmt.Printf("Using: %s\n", db.Name)
	// Output:
	// Found: users on port 5432
	// Using: default
}

// Example showing chaining operations
func ExampleResult_chaining() {
	// Parse and validate
	parsePort := func(s string) lxtypes.Result[int] {
		port, err := strconv.Atoi(s)
		if err != nil {
			return lxtypes.ResultFailure[int](err)
		}
		return lxtypes.ResultSuccess(port)
	}

	validatePort := func(port int) lxtypes.Result[int] {
		if port > 0 && port < 65536 {
			return lxtypes.ResultSuccess(port)
		}
		return lxtypes.ResultFailure[int](errors.New("invalid port"))
	}

	// Parse
	result := parsePort("8080")
	port := result.ValueOr(0)

	// Validate
	validated := validatePort(port)
	finalPort := validated.ValueOr(80)

	fmt.Println("Port:", finalPort)
	// Output:
	// Port: 8080
}

// Example showing error recovery
func ExampleResult_errorRecovery() {
	// Try primary, fallback to secondary
	connectPrimary := func() lxtypes.Result[string] {
		return lxtypes.ResultFailure[string](errors.New("primary unavailable"))
	}

	connectSecondary := func() lxtypes.Result[string] {
		return lxtypes.ResultSuccess("secondary")
	}

	// Try primary
	result := connectPrimary()
	connection := result.ValueOr("")

	// If failed, try secondary
	if connection == "" {
		result2 := connectSecondary()
		connection = result2.ValueOr("fallback")
	}

	fmt.Println("Connected to:", connection)
	// Output:
	// Connected to: secondary
}
