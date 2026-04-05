package lxtypes_test

import (
	"fmt"

	"github.com/hgapdvn/lx/lxtypes"
)

// Person represents a simple struct for examples
type User struct {
	Name  string
	Email string
	Age   int
}

func ExampleOptionalOf() {
	// Create an Optional with a value
	opt := lxtypes.OptionalOf(42)

	// Use comma-ok pattern to check and get value
	if value, ok := opt.Get(); ok {
		fmt.Println("Value:", value)
	}
	// Output:
	// Value: 42
}

func ExampleOptionalOf_struct() {
	// Create an Optional with a struct
	user := User{Name: "Alice", Email: "alice@example.com", Age: 30}
	opt := lxtypes.OptionalOf(user)

	if value, ok := opt.Get(); ok {
		fmt.Printf("User: %s, Email: %s\n", value.Name, value.Email)
	}
	// Output:
	// User: Alice, Email: alice@example.com
}

func ExampleOptionalOf_pointerStruct() {
	// Create an Optional with a pointer to struct
	user := &User{Name: "Bob", Email: "bob@example.com", Age: 25}
	opt := lxtypes.OptionalOf(user)

	if value, ok := opt.Get(); ok {
		fmt.Printf("User: %s, Age: %d\n", value.Name, value.Age)
	}
	// Output:
	// User: Bob, Age: 25
}

func ExampleOptionalEmpty() {
	// Create an empty Optional
	opt := lxtypes.OptionalEmpty[int]()

	// Check if value is present
	if _, ok := opt.Get(); !ok {
		fmt.Println("No value present")
	}

	// Use default value
	value := opt.OrElse(99)
	fmt.Println("Value:", value)
	// Output:
	// No value present
	// Value: 99
}

func ExampleOptionalEmpty_struct() {
	// Create an empty Optional for a struct
	opt := lxtypes.OptionalEmpty[User]()

	// Use default struct
	defaultUser := User{Name: "Guest", Email: "guest@example.com", Age: 0}
	user := opt.OrElse(defaultUser)
	fmt.Printf("User: %s\n", user.Name)
	// Output:
	// User: Guest
}

func ExampleOptionalOfNullable() {
	// Non-nil pointer creates a present Optional
	value := 42
	opt1 := lxtypes.OptionalOfNullable(&value)

	if v, ok := opt1.Get(); ok {
		fmt.Println("Value:", v)
	}

	// Nil pointer creates an empty Optional
	var nilPtr *int
	opt2 := lxtypes.OptionalOfNullable(nilPtr)

	if _, ok := opt2.Get(); !ok {
		fmt.Println("No value from nil pointer")
	}
	// Output:
	// Value: 42
	// No value from nil pointer
}

func ExampleOptionalOfNullable_struct() {
	// Non-nil struct pointer
	user := User{Name: "Charlie", Email: "charlie@example.com", Age: 35}
	opt1 := lxtypes.OptionalOfNullable(&user)

	if v, ok := opt1.Get(); ok {
		fmt.Printf("User: %s\n", v.Name)
	}

	// Nil struct pointer
	var nilPtr *User
	opt2 := lxtypes.OptionalOfNullable(nilPtr)

	if _, ok := opt2.Get(); !ok {
		fmt.Println("No user found")
	}
	// Output:
	// User: Charlie
	// No user found
}

func ExampleOptional_Get() {
	// Present value
	present := lxtypes.OptionalOf(42)
	if value, ok := present.Get(); ok {
		fmt.Println("Present:", value)
	}

	// Empty value
	empty := lxtypes.OptionalEmpty[int]()
	if value, ok := empty.Get(); ok {
		fmt.Println("Empty:", value)
	} else {
		fmt.Println("Empty: no value")
	}
	// Output:
	// Present: 42
	// Empty: no value
}

func ExampleOptional_OrElse() {
	// Present Optional returns its value
	present := lxtypes.OptionalOf(42)
	fmt.Println(present.OrElse(0))

	// Empty Optional returns default value
	empty := lxtypes.OptionalEmpty[int]()
	fmt.Println(empty.OrElse(99))
	// Output:
	// 42
	// 99
}

func ExampleOptional_OrElse_struct() {
	// Present Optional with struct
	user := User{Name: "Diana", Email: "diana@example.com", Age: 28}
	present := lxtypes.OptionalOf(user)
	defaultUser := User{Name: "Guest", Email: "guest@example.com", Age: 0}

	result := present.OrElse(defaultUser)
	fmt.Println(result.Name)

	// Empty Optional with struct
	empty := lxtypes.OptionalEmpty[User]()
	result2 := empty.OrElse(defaultUser)
	fmt.Println(result2.Name)
	// Output:
	// Diana
	// Guest
}

func ExampleOptional_OrElseGet() {
	// Present Optional doesn't call the function
	present := lxtypes.OptionalOf(42)
	result1 := present.OrElseGet(func() int {
		fmt.Println("Computing default...")
		return 0
	})
	fmt.Println(result1)

	// Empty Optional calls the function
	empty := lxtypes.OptionalEmpty[int]()
	result2 := empty.OrElseGet(func() int {
		fmt.Println("Computing default...")
		return 99
	})
	fmt.Println(result2)
	// Output:
	// 42
	// Computing default...
	// 99
}

func ExampleOptional_OrElseGet_struct() {
	// Empty Optional with expensive struct creation
	empty := lxtypes.OptionalEmpty[User]()

	user := empty.OrElseGet(func() User {
		fmt.Println("Creating default user...")
		return User{Name: "Guest", Email: "guest@example.com", Age: 0}
	})
	fmt.Println(user.Name)
	// Output:
	// Creating default user...
	// Guest
}

// Example showing a practical use case: database lookup
func ExampleOptional_databaseLookup() {
	// Simulate a database lookup that might return nil
	findUserByID := func(id int) *User {
		if id == 1 {
			return &User{Name: "Alice", Email: "alice@example.com", Age: 30}
		}
		return nil
	}

	// Safe handling with Optional
	userPtr := findUserByID(1)
	opt1 := lxtypes.OptionalOfNullable(userPtr)
	user1 := opt1.OrElse(User{Name: "Unknown", Email: "", Age: 0})
	fmt.Printf("User 1: %s\n", user1.Name)

	// Non-existent user
	userPtr2 := findUserByID(999)
	opt2 := lxtypes.OptionalOfNullable(userPtr2)
	user2 := opt2.OrElse(User{Name: "Unknown", Email: "", Age: 0})
	fmt.Printf("User 2: %s\n", user2.Name)
	// Output:
	// User 1: Alice
	// User 2: Unknown
}

// Example showing chaining with comma-ok pattern
func ExampleOptional_chaining() {
	// Find and process a value if present
	opt := lxtypes.OptionalOf(User{Name: "Eve", Email: "eve@example.com", Age: 20})

	if user, ok := opt.Get(); ok {
		// Process the user
		user.Age += 1
		fmt.Printf("%s is now %d years old\n", user.Name, user.Age)
	}
	// Output:
	// Eve is now 21 years old
}
