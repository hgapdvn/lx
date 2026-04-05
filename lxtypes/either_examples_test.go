package lxtypes_test

import (
	"fmt"

	"github.com/hgapdvn/lx/lxtypes"
)

func ExampleEitherLeft() {
	either := lxtypes.EitherLeft[string, int]("error occurred")

	// Check with Left()
	if left, ok := either.Left(); ok {
		fmt.Println("Left:", left)
	}

	// Right() returns false
	if _, ok := either.Right(); !ok {
		fmt.Println("Not a Right")
	}
	// Output:
	// Left: error occurred
	// Not a Right
}

func ExampleEitherRight() {
	either := lxtypes.EitherRight[string, int](42)

	// Check with Right()
	if right, ok := either.Right(); ok {
		fmt.Println("Right:", right)
	}

	// Left() returns false
	if _, ok := either.Left(); !ok {
		fmt.Println("Not a Left")
	}
	// Output:
	// Right: 42
	// Not a Left
}

func ExampleEither_pattern() {
	// Pattern matching style
	either := lxtypes.EitherRight[string, int](42)

	if left, ok := either.Left(); ok {
		fmt.Println("Left:", left)
	} else if right, ok := either.Right(); ok {
		fmt.Println("Right:", right)
	}
	// Output:
	// Right: 42
}

func ExampleEither_LeftOr() {
	left := lxtypes.EitherLeft[string, int]("error")
	right := lxtypes.EitherRight[string, int](42)

	fmt.Println(left.LeftOr("default"))
	fmt.Println(right.LeftOr("default"))
	// Output:
	// error
	// default
}

func ExampleEither_RightOr() {
	left := lxtypes.EitherLeft[string, int]("error")
	right := lxtypes.EitherRight[string, int](42)

	fmt.Println(left.RightOr(0))
	fmt.Println(right.RightOr(0))
	// Output:
	// 0
	// 42
}

func ExampleEither_validation() {
	type ValidationError struct {
		Field   string
		Message string
	}

	type User struct {
		Name string
		Age  int
	}

	validateUser := func(name string, age int) lxtypes.Either[ValidationError, User] {
		if age < 0 {
			return lxtypes.EitherLeft[ValidationError, User](
				ValidationError{Field: "age", Message: "must be non-negative"},
			)
		}
		if name == "" {
			return lxtypes.EitherLeft[ValidationError, User](
				ValidationError{Field: "name", Message: "cannot be empty"},
			)
		}
		return lxtypes.EitherRight[ValidationError, User](User{Name: name, Age: age})
	}

	result := validateUser("Alice", 30)
	if user, ok := result.Right(); ok {
		fmt.Printf("Valid user: %s, age %d\n", user.Name, user.Age)
	} else if err, ok := result.Left(); ok {
		fmt.Printf("Validation error in %s: %s\n", err.Field, err.Message)
	}
	// Output:
	// Valid user: Alice, age 30
}

func ExampleEither_withStruct() {
	type Config struct {
		Host string
		Port int
	}

	// Success case
	success := lxtypes.EitherRight[string, Config](Config{Host: "localhost", Port: 8080})

	if cfg, ok := success.Right(); ok {
		fmt.Printf("Config: %s:%d\n", cfg.Host, cfg.Port)
	}

	// Error case
	failure := lxtypes.EitherLeft[string, Config]("invalid config")

	if err, ok := failure.Left(); ok {
		fmt.Println("Error:", err)
	}
	// Output:
	// Config: localhost:8080
	// Error: invalid config
}
