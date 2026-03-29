package lxtypes_test

import (
	"testing"

	"github.com/hgapdvn/lx/types"
)

// Test with primitive types
func TestEitherLeft(t *testing.T) {
	either := lxtypes.EitherLeft[string, int]("error")

	// Test Left() returns value and true
	left, ok := either.Left()
	if !ok {
		t.Error("Left() should return true for Left-sided Either")
	}
	if left != "error" {
		t.Errorf("Left value = %v, want 'error'", left)
	}

	// Test Right() returns false
	_, ok = either.Right()
	if ok {
		t.Error("Right() should return false for Left-sided Either")
	}
}

func TestEitherRight(t *testing.T) {
	either := lxtypes.EitherRight[string, int](42)

	// Test Right() returns value and true
	right, ok := either.Right()
	if !ok {
		t.Error("Right() should return true for Right-sided Either")
	}
	if right != 42 {
		t.Errorf("Right value = %v, want 42", right)
	}

	// Test Left() returns false
	_, ok = either.Left()
	if ok {
		t.Error("Left() should return false for Right-sided Either")
	}
}

func TestEitherLeftOr(t *testing.T) {
	left := lxtypes.EitherLeft[string, int]("error")
	right := lxtypes.EitherRight[string, int](42)

	if got := left.LeftOr("default"); got != "error" {
		t.Errorf("LeftOr on Left = %v, want 'error'", got)
	}

	if got := right.LeftOr("default"); got != "default" {
		t.Errorf("LeftOr on Right = %v, want 'default'", got)
	}
}

func TestEitherRightOr(t *testing.T) {
	left := lxtypes.EitherLeft[string, int]("error")
	right := lxtypes.EitherRight[string, int](42)

	if got := left.RightOr(0); got != 0 {
		t.Errorf("RightOr on Left = %v, want 0", got)
	}

	if got := right.RightOr(0); got != 42 {
		t.Errorf("RightOr on Right = %v, want 42", got)
	}
}

// Test with struct types
func TestEitherWithStruct(t *testing.T) {
	type ValidationError struct {
		Field   string
		Message string
	}

	type User struct {
		Name string
		Age  int
	}

	// Test with valid user (Right)
	validUser := lxtypes.EitherRight[ValidationError, User](User{Name: "Alice", Age: 30})
	user, ok := validUser.Right()
	if !ok {
		t.Error("Right() should return true for valid user")
	}
	if user.Name != "Alice" || user.Age != 30 {
		t.Errorf("User = %+v, want {Name:Alice Age:30}", user)
	}

	// Test Left() on Right returns false
	_, ok = validUser.Left()
	if ok {
		t.Error("Left() should return false for Right-sided Either")
	}

	// Test with invalid user (Left)
	invalidUser := lxtypes.EitherLeft[ValidationError, User](
		ValidationError{Field: "age", Message: "must be positive"},
	)
	valErr, ok := invalidUser.Left()
	if !ok {
		t.Error("Left() should return true for invalid user")
	}
	if valErr.Field != "age" || valErr.Message != "must be positive" {
		t.Errorf("ValidationError = %+v, want {Field:age Message:must be positive}", valErr)
	}

	// Test Right() on Left returns false
	_, ok = invalidUser.Right()
	if ok {
		t.Error("Right() should return false for Left-sided Either")
	}
}

// Test with pointer types
func TestEitherWithPointer(t *testing.T) {
	type Config struct {
		Host string
		Port int
	}

	// Test with Left pointer
	errMsg := "connection failed"
	leftPtr := lxtypes.EitherLeft[*string, *Config](&errMsg)

	ptr, ok := leftPtr.Left()
	if !ok {
		t.Error("Left() should return true for Left-sided Either with pointer")
	}
	if ptr == nil || *ptr != "connection failed" {
		t.Errorf("Left pointer value = %v, want 'connection failed'", ptr)
	}

	// Test Right() on Left returns false
	_, ok = leftPtr.Right()
	if ok {
		t.Error("Right() should return false for Left-sided Either")
	}

	// Test with Right pointer
	cfg := &Config{Host: "localhost", Port: 8080}
	rightPtr := lxtypes.EitherRight[*string, *Config](cfg)

	cfgPtr, ok := rightPtr.Right()
	if !ok {
		t.Error("Right() should return true for Right-sided Either with pointer")
	}
	if cfgPtr == nil || cfgPtr.Host != "localhost" || cfgPtr.Port != 8080 {
		t.Errorf("Right pointer value = %+v, want {Host:localhost Port:8080}", cfgPtr)
	}

	// Test Left() on Right returns false
	_, ok = rightPtr.Left()
	if ok {
		t.Error("Left() should return false for Right-sided Either")
	}
}

// Test with nested structs
func TestEitherWithNestedStruct(t *testing.T) {
	type Address struct {
		City    string
		Country string
	}

	type Person struct {
		Name    string
		Address Address
	}

	type Error struct {
		Code    int
		Message string
	}

	// Test with Right (Person)
	person := Person{
		Name:    "Bob",
		Address: Address{City: "Tokyo", Country: "Japan"},
	}
	rightEither := lxtypes.EitherRight[Error, Person](person)

	p, ok := rightEither.Right()
	if !ok {
		t.Error("Right() should return true")
	}
	if p.Name != "Bob" || p.Address.City != "Tokyo" {
		t.Errorf("Person = %+v, want {Name:Bob Address:{City:Tokyo Country:Japan}}", p)
	}

	// Test with Left (Error)
	err := Error{Code: 404, Message: "not found"}
	leftEither := lxtypes.EitherLeft[Error, Person](err)

	e, ok := leftEither.Left()
	if !ok {
		t.Error("Left() should return true")
	}
	if e.Code != 404 || e.Message != "not found" {
		t.Errorf("Error = %+v, want {Code:404 Message:not found}", e)
	}
}

// Test OrElse methods with structs
func TestEitherOrElseWithStruct(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	defaultUser := User{ID: 0, Name: "Guest"}
	defaultError := "no error"

	// Test LeftOr with struct
	leftEither := lxtypes.EitherLeft[string, User]("user not found")
	if got := leftEither.LeftOr(defaultError); got != "user not found" {
		t.Errorf("LeftOr = %v, want 'user not found'", got)
	}

	// Test RightOr with struct
	rightEither := lxtypes.EitherRight[string, User](User{ID: 1, Name: "Alice"})
	if got := rightEither.RightOr(defaultUser); got.ID != 1 || got.Name != "Alice" {
		t.Errorf("RightOr = %+v, want {ID:1 Name:Alice}", got)
	}

	// Test RightOr returns default when Left
	if got := leftEither.RightOr(defaultUser); got.ID != 0 || got.Name != "Guest" {
		t.Errorf("RightOr on Left = %+v, want {ID:0 Name:Guest}", got)
	}

	// Test LeftOr returns default when Right
	if got := rightEither.LeftOr(defaultError); got != "no error" {
		t.Errorf("LeftOr on Right = %v, want 'no error'", got)
	}
}

// Test with pointer to struct
func TestEitherWithPointerToStruct(t *testing.T) {
	type Result struct {
		Value int
		Valid bool
	}

	// Test with nil pointer
	var nilPtr *Result
	leftNil := lxtypes.EitherLeft[string, *Result]("invalid")

	str, ok := leftNil.Left()
	if !ok || str != "invalid" {
		t.Errorf("Left() = (%v, %v), want ('invalid', true)", str, ok)
	}

	// Test RightOr with nil default
	rightVal := leftNil.RightOr(nilPtr)
	if rightVal != nil {
		t.Errorf("RightOr with nil = %v, want nil", rightVal)
	}

	// Test with non-nil pointer
	result := &Result{Value: 42, Valid: true}
	rightPtr := lxtypes.EitherRight[string, *Result](result)

	ptr, ok := rightPtr.Right()
	if !ok {
		t.Error("Right() should return true")
	}
	if ptr == nil || ptr.Value != 42 || !ptr.Valid {
		t.Errorf("Right() pointer = %+v, want {Value:42 Valid:true}", ptr)
	}

	// Verify it's the same pointer
	if ptr != result {
		t.Error("Right() should return the same pointer instance")
	}
}
