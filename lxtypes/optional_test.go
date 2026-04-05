package lxtypes_test

import (
	"testing"

	"github.com/hgapdvn/lx/lxtypes"
)

// Test types
type Person struct {
	Name string
	Age  int
}

func TestOptionalOf(t *testing.T) {
	t.Run("integer", func(t *testing.T) {
		opt := lxtypes.OptionalOf(42)

		value, ok := opt.Get()
		if !ok {
			t.Error("Expected Get to return true")
		}
		if value != 42 {
			t.Errorf("Get() = %v, want 42", value)
		}
	})

	t.Run("string", func(t *testing.T) {
		opt := lxtypes.OptionalOf("hello")

		value, ok := opt.Get()
		if !ok {
			t.Error("Expected Get to return true")
		}
		if value != "hello" {
			t.Errorf("Get() = %v, want hello", value)
		}
	})

	t.Run("struct", func(t *testing.T) {
		person := Person{Name: "Alice", Age: 30}
		opt := lxtypes.OptionalOf(person)

		value, ok := opt.Get()
		if !ok {
			t.Error("Expected Get to return true")
		}
		if value.Name != "Alice" || value.Age != 30 {
			t.Errorf("Get() = %+v, want %+v", value, person)
		}
	})

	t.Run("pointer of struct", func(t *testing.T) {
		person := &Person{Name: "Bob", Age: 25}
		opt := lxtypes.OptionalOf(person)

		value, ok := opt.Get()
		if !ok {
			t.Error("Expected Get to return true")
		}
		if value.Name != "Bob" || value.Age != 25 {
			t.Errorf("Get() = %+v, want %+v", value, person)
		}
	})

	t.Run("zero value", func(t *testing.T) {
		// Even zero values should be present
		opt := lxtypes.OptionalOf(0)

		value, ok := opt.Get()
		if !ok {
			t.Error("Expected Get to return true for zero value")
		}
		if value != 0 {
			t.Errorf("Get() = %v, want 0", value)
		}
	})
}

func TestOptionalEmpty(t *testing.T) {
	t.Run("integer", func(t *testing.T) {
		opt := lxtypes.OptionalEmpty[int]()

		value, ok := opt.Get()
		if ok {
			t.Error("Expected Get to return false for empty")
		}
		if value != 0 {
			t.Errorf("Get() returned %v, want 0 for empty", value)
		}
	})

	t.Run("string", func(t *testing.T) {
		opt := lxtypes.OptionalEmpty[string]()

		value, ok := opt.Get()
		if ok {
			t.Error("Expected Get to return false for empty")
		}
		if value != "" {
			t.Errorf("Get() returned %v, want empty string", value)
		}
	})

	t.Run("struct", func(t *testing.T) {
		opt := lxtypes.OptionalEmpty[Person]()

		value, ok := opt.Get()
		if ok {
			t.Error("Expected Get to return false for empty")
		}
		if value.Name != "" || value.Age != 0 {
			t.Errorf("Get() returned %+v, want zero Person", value)
		}
	})

	t.Run("pointer of struct", func(t *testing.T) {
		opt := lxtypes.OptionalEmpty[*Person]()

		value, ok := opt.Get()
		if ok {
			t.Error("Expected Get to return false for empty")
		}
		if value != nil {
			t.Errorf("Get() returned %+v, want nil", value)
		}
	})
}

func TestOptionalOrElse(t *testing.T) {
	t.Run("present value returns original", func(t *testing.T) {
		present := lxtypes.OptionalOf(42)
		if got := present.OrElse(0); got != 42 {
			t.Errorf("Of.OrElse(0) = %v, want 42", got)
		}
	})

	t.Run("empty returns default", func(t *testing.T) {
		empty := lxtypes.OptionalEmpty[int]()
		if got := empty.OrElse(99); got != 99 {
			t.Errorf("Empty.OrElse(99) = %v, want 99", got)
		}
	})

	t.Run("struct with present", func(t *testing.T) {
		person := Person{Name: "Alice", Age: 30}
		present := lxtypes.OptionalOf(person)
		defaultPerson := Person{Name: "Default", Age: 0}

		got := present.OrElse(defaultPerson)
		if got.Name != "Alice" {
			t.Errorf("OrElse() = %+v, want %+v", got, person)
		}
	})

	t.Run("struct with empty", func(t *testing.T) {
		empty := lxtypes.OptionalEmpty[Person]()
		defaultPerson := Person{Name: "Default", Age: 0}

		got := empty.OrElse(defaultPerson)
		if got.Name != "Default" {
			t.Errorf("OrElse() = %+v, want %+v", got, defaultPerson)
		}
	})

	t.Run("pointer of struct with present", func(t *testing.T) {
		person := &Person{Name: "Alice", Age: 30}
		present := lxtypes.OptionalOf(person)
		defaultPerson := &Person{Name: "Default", Age: 0}

		got := present.OrElse(defaultPerson)
		if got.Name != "Alice" {
			t.Errorf("OrElse() = %+v, want %+v", got, person)
		}
	})

	t.Run("pointer of struct with empty", func(t *testing.T) {
		empty := lxtypes.OptionalEmpty[*Person]()
		defaultPerson := &Person{Name: "Default", Age: 0}

		got := empty.OrElse(defaultPerson)
		if got.Name != "Default" {
			t.Errorf("OrElse() = %+v, want %+v", got, defaultPerson)
		}
	})
}

func TestOptionalOrElseGet(t *testing.T) {
	t.Run("present value doesn't call function", func(t *testing.T) {
		present := lxtypes.OptionalOf(42)
		called := false
		fn := func() int {
			called = true
			return 0
		}

		if got := present.OrElseGet(fn); got != 42 {
			t.Errorf("Of.OrElseGet(...) = %v, want 42", got)
		}
		if called {
			t.Error("Function should not be called when value is present")
		}
	})

	t.Run("empty calls function", func(t *testing.T) {
		empty := lxtypes.OptionalEmpty[int]()
		called := false
		fn := func() int {
			called = true
			return 99
		}

		if got := empty.OrElseGet(fn); got != 99 {
			t.Errorf("Empty.OrElseGet(...) = %v, want 99", got)
		}
		if !called {
			t.Error("Function should be called when value is empty")
		}
	})

	t.Run("struct with empty", func(t *testing.T) {
		empty := lxtypes.OptionalEmpty[Person]()
		fn := func() Person {
			return Person{Name: "Computed", Age: 42}
		}

		got := empty.OrElseGet(fn)
		if got.Name != "Computed" || got.Age != 42 {
			t.Errorf("OrElseGet() = %+v, want {Computed 42}", got)
		}
	})

	t.Run("pointer of struct with empty", func(t *testing.T) {
		empty := lxtypes.OptionalEmpty[*Person]()
		fn := func() *Person {
			return &Person{Name: "Computed", Age: 42}
		}

		got := empty.OrElseGet(fn)
		if got.Name != "Computed" || got.Age != 42 {
			t.Errorf("OrElseGet() = %+v, want {Computed 42}", got)
		}
	})
}

func TestOptionalOfNullable(t *testing.T) {
	t.Run("non-nil pointer creates present", func(t *testing.T) {
		value := 42
		opt := lxtypes.OptionalOfNullable(&value)

		got, ok := opt.Get()
		if !ok {
			t.Error("Expected OfNullable with non-nil to be present")
		}
		if got != 42 {
			t.Errorf("OfNullable value = %v, want 42", got)
		}
	})

	t.Run("nil pointer creates empty", func(t *testing.T) {
		var nilPtr *int
		opt := lxtypes.OptionalOfNullable(nilPtr)

		_, ok := opt.Get()
		if ok {
			t.Error("Expected OfNullable with nil to be empty")
		}
	})

	t.Run("struct pointer non-nil", func(t *testing.T) {
		person := Person{Name: "Charlie", Age: 35}
		opt := lxtypes.OptionalOfNullable(&person)

		got, ok := opt.Get()
		if !ok {
			t.Error("Expected OfNullable with non-nil struct to be present")
		}
		if got.Name != "Charlie" || got.Age != 35 {
			t.Errorf("OfNullable value = %+v, want %+v", got, person)
		}
	})

	t.Run("struct pointer nil", func(t *testing.T) {
		var nilPtr *Person
		opt := lxtypes.OptionalOfNullable(nilPtr)

		_, ok := opt.Get()
		if ok {
			t.Error("Expected OfNullable with nil struct to be empty")
		}
	})

	t.Run("pointer of pointer struct", func(t *testing.T) {
		person := &Person{Name: "Diana", Age: 28}
		opt := lxtypes.OptionalOfNullable(&person)

		got, ok := opt.Get()
		if !ok {
			t.Error("Expected OfNullable with non-nil pointer to be present")
		}
		if got.Name != "Diana" || got.Age != 28 {
			t.Errorf("OfNullable value = %+v, want %+v", got, person)
		}
	})
}

func TestOptionalChaining(t *testing.T) {
	t.Run("chain with present value", func(t *testing.T) {
		opt := lxtypes.OptionalOf(42)
		result := opt.OrElse(0) + 10
		if result != 52 {
			t.Errorf("Expected 52, got %v", result)
		}
	})

	t.Run("chain with empty value", func(t *testing.T) {
		opt := lxtypes.OptionalEmpty[int]()
		result := opt.OrElse(0) + 10
		if result != 10 {
			t.Errorf("Expected 10, got %v", result)
		}
	})

	t.Run("struct modification", func(t *testing.T) {
		opt := lxtypes.OptionalOf(Person{Name: "Eve", Age: 20})
		person := opt.OrElse(Person{Name: "Unknown", Age: 0})
		person.Age += 1

		if person.Age != 21 {
			t.Errorf("Expected Age 21, got %v", person.Age)
		}
	})

	t.Run("pointer struct modification", func(t *testing.T) {
		original := &Person{Name: "Frank", Age: 40}
		opt := lxtypes.OptionalOf(original)
		person := opt.OrElse(&Person{Name: "Unknown", Age: 0})
		person.Age += 1

		// Should modify the original
		if original.Age != 41 {
			t.Errorf("Expected Age 41, got %v", original.Age)
		}
	})
}

func TestOptionalCommaOkPattern(t *testing.T) {
	t.Run("present value with comma-ok", func(t *testing.T) {
		opt := lxtypes.OptionalOf(42)

		if value, ok := opt.Get(); ok {
			if value != 42 {
				t.Errorf("Expected value 42, got %v", value)
			}
		} else {
			t.Error("Expected ok=true for present value")
		}
	})

	t.Run("empty value with comma-ok", func(t *testing.T) {
		opt := lxtypes.OptionalEmpty[int]()

		if value, ok := opt.Get(); ok {
			t.Errorf("Expected ok=false, but got value %v", value)
		}
	})

	t.Run("struct with comma-ok", func(t *testing.T) {
		opt := lxtypes.OptionalOf(Person{Name: "Grace", Age: 45})

		if person, ok := opt.Get(); ok {
			if person.Name != "Grace" {
				t.Errorf("Expected Grace, got %v", person.Name)
			}
		} else {
			t.Error("Expected ok=true for present struct")
		}
	})
}
