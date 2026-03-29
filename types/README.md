# lxtypes

Generic type definitions for functional programming, optional values, and error handling in Go.

## Overview

The `lxtypes` package provides reusable, type-safe generic type definitions inspired by functional programming patterns from Java and Rust. These types enable cleaner, more composable code when working with functions, optional values, and error handling.

**Includes:**
- **Functional types**: Predicate, Consumer, Function, Supplier, and their binary variants
- **Optional values**: Optional[T] (Java-style Optional)
- **Error handling**: Result[T] (specialized for Go's error type)
- **Binary choice**: Either[L, R] (general-purpose union type)
- **Tuple types**: Pair, Triple, Quad for multi-value returns
- **Lazy evaluation**: Lazy[T] for deferred computation with caching
- **Async operations**: Future[T] for concurrent computations with composability
- **Mutable state**: Ref[T] for thread-safe shared mutable values

## Installation

Install the lx module (includes all packages):

```bash
go get github.com/nthanhhai2909/lx
```

Then import only what you need:

```go
import "github.com/nthanhhai2909/lx/lxtypes"
```

## Type Definitions

### Basic Functional Types

#### Predicate[T]
Tests a condition on an input value.

```go
isPositive := lxtypes.Predicate[int](func(n int) bool {
    return n > 0
})

fmt.Println(isPositive(5))   // true
fmt.Println(isPositive(-3))  // false
```

**Methods:**
- `And(other Predicate[T]) Predicate[T]` - Logical AND composition
- `Or(other Predicate[T]) Predicate[T]` - Logical OR composition
- `Negate() Predicate[T]` - Logical negation

#### Consumer[T]
Performs an operation on an input without returning a result.

```go
printValue := lxtypes.Consumer[string](func(s string) {
    fmt.Println(s)
})

printValue("Hello, World!")
```

**Methods:**
- `AndThen(after Consumer[T]) Consumer[T]` - Sequential composition

#### Function[T, U]
Transforms an input of type T to an output of type U.

```go
toUpper := lxtypes.Function[string, string](func(s string) string {
    return strings.ToUpper(s)
})

result := toUpper("hello")  // "HELLO"
```

#### Supplier[T]
Provides a value with no input.

```go
counter := 0
getNext := lxtypes.Supplier[int](func() int {
    counter++
    return counter
})

fmt.Println(getNext())  // 1
fmt.Println(getNext())  // 2
```

### Binary Types

#### BiPredicate[T, U]
Tests a condition on two input values.

```go
equals := lxtypes.BiPredicate[int, int](func(a, b int) bool {
    return a == b
})

fmt.Println(equals(5, 5))  // true
```

**Methods:**
- `And(other BiPredicate[T, U]) BiPredicate[T, U]` - Logical AND composition
- `Or(other BiPredicate[T, U]) BiPredicate[T, U]` - Logical OR composition
- `Negate() BiPredicate[T, U]` - Logical negation

#### BiConsumer[T, U]
Performs an operation on two inputs without returning a result.

```go
printPair := lxtypes.BiConsumer[string, int](func(s string, n int) {
    fmt.Printf("%s: %d\n", s, n)
})

printPair("count", 42)  // count: 42
```

#### BiFunction[T, U, R]
Transforms two inputs to a single output.

```go
concat := lxtypes.BiFunction[string, string, string](func(a, b string) string {
    return a + b
})

result := concat("Hello, ", "World!")  // "Hello, World!"
```

### Operator Types

#### UnaryOperator[T]
Transforms a value to the same type.

```go
square := lxtypes.UnaryOperator[int](func(n int) int {
    return n * n
})

result := square(5)  // 25
```

#### BinaryOperator[T]
Combines two values of the same type.

```go
sum := lxtypes.BinaryOperator[int](func(a, b int) int {
    return a + b
})

result := sum(3, 7)  // 10
```

### Comparison

#### Comparator[T]
Compares two values for ordering.

```go
intComparator := lxtypes.Comparator[int](func(a, b int) int {
    if a < b {
        return -1
    }
    if a > b {
        return 1
    }
    return 0
})

result := intComparator(3, 5)  // -1 (3 < 5)
```

**Methods:**
- `Reversed() Comparator[T]` - Reverse the ordering
- `ThenComparing(other Comparator[T]) Comparator[T]` - Lexicographic comparison

## Optional and Error Handling Types

### Optional[T]

Represents a value that may or may not be present. Inspired by Java's `Optional<T>` but adapted to Go's idiomatic comma-ok pattern.

An `Optional` is either:
- **Present**: Contains a value (created with `OptionalOf` or `OptionalOfNullable`)
- **Empty**: Contains no value (created with `OptionalEmpty`)

#### Creating Optionals

```go
// Present value
opt := lxtypes.OptionalOf(42)

// Empty value
empty := lxtypes.OptionalEmpty[int]()

// From nullable pointer (safe nil handling)
value := 42
opt1 := lxtypes.OptionalOfNullable(&value)  // OptionalOf(42)

var nilPtr *int
opt2 := lxtypes.OptionalOfNullable(nilPtr)  // OptionalEmpty()
```

#### Checking Presence with Comma-Ok Pattern

```go
opt := lxtypes.OptionalOf(42)

// Use Go's idiomatic comma-ok pattern
if value, ok := opt.Get(); ok {
    fmt.Println("Has value:", value)  // Has value: 42
} else {
    fmt.Println("No value")
}
```

#### Working with Structs

```go
type User struct {
    Name  string
    Email string
    Age   int
}

// With struct values
user := User{Name: "Alice", Email: "alice@example.com", Age: 30}
opt := lxtypes.OptionalOf(user)

if u, ok := opt.Get(); ok {
    fmt.Printf("User: %s, Email: %s\n", u.Name, u.Email)
}

// With struct pointers
userPtr := &User{Name: "Bob", Email: "bob@example.com", Age: 25}
opt2 := lxtypes.OptionalOf(userPtr)

if u, ok := opt2.Get(); ok {
    fmt.Printf("User: %s\n", u.Name)
}

// Safe nil handling with OptionalOfNullable
var nilUser *User
opt3 := lxtypes.OptionalOfNullable(nilUser)  // Empty Optional

defaultUser := User{Name: "Guest", Email: "guest@example.com"}
user = opt3.OrElse(defaultUser)  // Returns defaultUser
```

#### Safe Access with Defaults

```go
// Static default
opt := lxtypes.OptionalOf(42)
value := opt.OrElse(0)  // 42

empty := lxtypes.OptionalEmpty[int]()
value2 := empty.OrElse(99)  // 99

// Computed default (lazy evaluation)
value3 := empty.OrElseGet(func() int {
    return expensiveComputation()
})
```

#### Practical Example: Database Lookup

```go
// Simulate a database lookup that might return nil
func FindUserByID(id int) *User {
    // Database lookup...
    if found {
        return &user
    }
    return nil
}

// Safe handling with Optional
userPtr := FindUserByID(123)
opt := lxtypes.OptionalOfNullable(userPtr)

// Use comma-ok pattern
if user, ok := opt.Get(); ok {
    fmt.Printf("Found: %s\n", user.Name)
} else {
    fmt.Println("User not found")
}

// Or use default
defaultUser := User{Name: "Guest", Email: "guest@example.com"}
user := opt.OrElse(defaultUser)
```

**Methods:**
- `Get() (T, bool)` - Returns (value, true) if present, or (zero, false) if empty (comma-ok pattern)
- `OrElse(T) T` - Get value or default
- `OrElseGet(func() T) T` - Get value or computed default (lazy evaluation)

**Use Cases:**
- Safe dictionary/map lookups
- Database query results that might not exist
- Configuration values that are optional
- Eliminating nil pointer panics
- API responses that may be absent
- File operations that may fail

### Result[T]

Represents the result of an operation that may succeed or fail. Adapted to Go's idiomatic (value, error) pattern.

A `Result` is either:
- **Success**: Contains a value (created with `ResultSuccess`)
- **Failure**: Contains an error (created with `ResultFailure`)

#### Creating Results

```go
// Success with value
result := lxtypes.ResultSuccess(42)

// Failure with error
result := lxtypes.ResultFailure[int](errors.New("operation failed"))
```

#### Checking Success with Value-Error Pattern

```go
result := divide(10, 2)

// Use Go's idiomatic (value, error) pattern
if value, err := result.Value(); err == nil {
    fmt.Println("Success:", value)
} else {
    fmt.Println("Error:", err)
}
```

#### Working with Structs

```go
type Config struct {
    Host string
    Port int
}

// With struct values
config := Config{Host: "localhost", Port: 8080}
result := lxtypes.ResultSuccess(config)

if cfg, err := result.Value(); err == nil {
    fmt.Printf("Config: %s:%d\n", cfg.Host, cfg.Port)
}

// With struct pointers
configPtr := &Config{Host: "example.com", Port: 443}
result2 := lxtypes.ResultSuccess(configPtr)

if cfg, err := result2.Value(); err == nil {
    fmt.Printf("Config: %s\n", cfg.Host)
}

// Failure with default
result3 := lxtypes.ResultFailure[Config](errors.New("config not found"))
defaultConfig := Config{Host: "default", Port: 80}
config = result3.ValueOr(defaultConfig)  // Returns defaultConfig
```

#### Safe Access with Defaults

```go
// Success returns original value
success := lxtypes.ResultSuccess(42)
value := success.ValueOr(0)  // 42

// Failure returns default value
failure := lxtypes.ResultFailure[int](errors.New("error"))
value2 := failure.ValueOr(99)  // 99
```

#### Practical Example: API Call

```go
// Simulate an API call
func FetchUser(id int) lxtypes.Result[User] {
    // Make API call...
    if err != nil {
        return lxtypes.ResultFailure[User](err)
    }
    return lxtypes.ResultSuccess(user)
}

// Handle the result
result := FetchUser(123)

// Use value-error pattern
if user, err := result.Value(); err == nil {
    fmt.Printf("User: %s\n", user.Name)
} else {
    fmt.Printf("Error: %v\n", err)
}

// Or use default
defaultUser := User{Name: "Guest"}
user := result.ValueOr(defaultUser)
```

#### Converting from Go's (value, error) Pattern

```go
// Standard library function
value, err := strconv.Atoi("42")

// Convert to Result
var result lxtypes.Result[int]
if err != nil {
    result = lxtypes.ResultFailure[int](err)
} else {
    result = lxtypes.ResultSuccess(value)
}

// Now use Result methods
finalValue := result.ValueOr(0)
```

**Methods:**
- `Value() (T, error)` - Returns (value, nil) if success, or (zero, err) if failure (Go's idiomatic pattern)
- `ValueOr(T) T` - Get value or default (no error checking needed)

**Use Cases:**
- Wrapping functions that return (value, error)
- Standardizing error handling across your codebase
- Avoiding nil pointer panics with default values
- Database operations that may fail
- File I/O operations
- Network requests
- Configuration loading
- Chainable error propagation

### Either[L, R]

Represents a value of one of two possible types (a disjoint union). 
For general-purpose binary choice where both alternatives are valid values.

An `Either` is either:
- **Left**: Contains a left value (often used for errors/failures)
- **Right**: Contains a right value (often used for success/normal cases)

**Note**: Both Left and Right are equally valid - this is different from Result which is specifically for error handling with Go's `error` type.

#### Creating Eithers

```go
// Left
either := lxtypes.EitherLeft[string, int]("error")

// Right
either := lxtypes.EitherRight[string, int](42)
```

#### Accessing Values

```go
// Check and access Left
if left, ok := either.Left(); ok {
    fmt.Println("Left value:", left)
}

// Check and access Right
if right, ok := either.Right(); ok {
    fmt.Println("Right value:", right)
}

// Pattern matching style
if left, ok := either.Left(); ok {
    fmt.Println("Error:", left)
} else if right, ok := either.Right(); ok {
    fmt.Println("Success:", right)
}
```

#### Safe Access with Defaults

```go
// Get value or default
leftVal := either.LeftOr("default")
rightVal := either.RightOr(0)
```

**Methods:**
- `Left() (L, bool)` - Returns left value and true if Left, or zero value and false if Right
- `Right() (R, bool)` - Returns right value and true if Right, or zero value and false if Left
- `LeftOr(L) L` - Returns left value or default if Right
- `RightOr(R) R` - Returns right value or default if Left

**Use Cases:**
- Validation with custom error types
- Parsing that returns one of two types
- Union types before pattern matching
- Polymorphic return values
- Polymorphic return values

**Real-World Example**:
```go
type ValidationError struct {
    Field   string
    Message string
}

type User struct {
    Name string
    Age  int
}

func ValidateUser(name string, age int) lxtypes.Either[ValidationError, User] {
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

// Usage
result := ValidateUser("Alice", 30)
if user, ok := result.Right(); ok {
    fmt.Printf("Valid user: %s, age %d\n", user.Name, user.Age)
} else if validationErr, ok := result.Left(); ok {
    fmt.Printf("Validation error in %s: %s\n", validationErr.Field, validationErr.Message)
}
```

### When to Use Which?

| Type | Use When | Example |
|------|----------|---------|
| **Optional[T]** | Value might be absent | `FindUser(id) Optional[User]` |
| **Result[T]** | Operation might fail with Go's error | `ReadFile(path) Result[[]byte]` |
| **Either[L, R]** | Need custom error types or general binary choice | `Either[ValidationError, User]` |


## Tuple Types

### Pair[T, U]
A generic two-element tuple.

```go
// Create a pair
p := lxtypes.NewPair(42, "answer")
fmt.Printf("First: %d, Second: %s\n", p.First, p.Second)

// Unpack values
x, y := p.Values()

// Swap elements
swapped := p.Swap()  // Pair[string, int]

// Transform elements
doubled := p.MapFirst(func(n int) int { return n * 2 })
upper := p.MapSecond(func(s string) string { return strings.ToUpper(s) })
```

**Methods:**
- `Values() (T, U)` - Unpack the pair into separate values
- `Swap() Pair[U, T]` - Create a new pair with swapped elements
- `MapFirst(func(T) T) Pair[T, U]` - Transform the first element
- `MapSecond(func(U) U) Pair[T, U]` - Transform the second element

**Use Cases:**
- Return multiple values from functions
- Zip operations (combine two slices)
- Key-value pairs
- Coordinate pairs

### Triple[T, U, V]
A generic three-element tuple.

```go
// Create a triple
t := lxtypes.NewTriple(1, "hello", true)
fmt.Printf("Values: %d, %s, %t\n", t.First, t.Second, t.Third)

// Unpack values
x, y, z := t.Values()

// Convert to pair (drops third element)
p := t.ToPair()
```

**Methods:**
- `Values() (T, U, V)` - Unpack the triple into separate values
- `ToPair() Pair[T, U]` - Convert to a Pair, discarding the third element

**Use Cases:**
- RGB color values
- 3D coordinates
- Database query results with multiple columns
- Function results with status, value, and metadata

### Quad[T, U, V, W]
A generic four-element tuple.

```go
// Create a quad
q := lxtypes.NewQuad(1, "hello", true, 3.14)
fmt.Printf("Values: %d, %s, %t, %.2f\n", q.First, q.Second, q.Third, q.Fourth)

// Unpack values
w, x, y, z := q.Values()

// Convert to smaller tuples
p := q.ToPair()    // Pair with first two elements
t := q.ToTriple()  // Triple with first three elements
```

**Methods:**
- `Values() (T, U, V, W)` - Unpack the quad into separate values
- `ToPair() Pair[T, U]` - Convert to a Pair, discarding last two elements
- `ToTriple() Triple[T, U, V]` - Convert to a Triple, discarding the fourth element

**Use Cases:**
- RGBA color values
- Complex return values
- Database rows with multiple typed columns
- Configuration tuples

### Tuple5[T1, T2, T3, T4, T5]
A generic five-element tuple for combining five different types.

```go
// Create a tuple with 5 values
t := lxtypes.NewTuple5(1, "two", true, 4.0, []int{5, 6})
fmt.Printf("Values: %d, %s, %t, %.1f, %v\n", t.V1, t.V2, t.V3, t.V4, t.V5)

// Unpack values
v1, v2, v3, v4, v5 := t.Values()

// Real-world example: Combining service responses
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
```

**Methods:**
- `Values() (T1, T2, T3, T4, T5)` - Unpack the tuple into separate values

**Use Cases:**
- Combining results from 5 different services
- Complex dashboard data aggregation
- Multi-source data fetching with FutureJoin5

### Tuple6[T1, T2, T3, T4, T5, T6]
A generic six-element tuple.

```go
t := lxtypes.NewTuple6(1, "two", true, 4.0, []int{5}, 'a')
v1, v2, v3, v4, v5, v6 := t.Values()
```

**Methods:**
- `Values() (T1, T2, T3, T4, T5, T6)` - Unpack the tuple into separate values

**Use Cases:**
- Combining results from 6 different services
- Complex data pipelines
- Multi-dimensional data structures

### Tuple7[T1, T2, T3, T4, T5, T6, T7]
A generic seven-element tuple.

```go
t := lxtypes.NewTuple7(1, "two", true, 4.0, []int{5}, 'a', byte(7))
v1, v2, v3, v4, v5, v6, v7 := t.Values()
```

**Methods:**
- `Values() (T1, T2, T3, T4, T5, T6, T7)` - Unpack the tuple into separate values

### Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]
A generic eight-element tuple for combining eight different types.

```go
t := lxtypes.NewTuple8(1, "two", true, 4.0, []int{5}, 'a', byte(7), uint(8))
v1, v2, v3, v4, v5, v6, v7, v8 := t.Values()

fmt.Printf("All 8 values: %d, %s, %t, %.1f, %v, %c, %d, %d\n",
    t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8)
```

**Methods:**
- `Values() (T1, T2, T3, T4, T5, T6, T7, T8)` - Unpack the tuple into separate values

**Use Cases:**
- Combining results from 8 different services
- Maximum flexibility for parallel data fetching
- Complex microservice orchestration

**Note:** Tuples 5-8 use a `V1, V2, V3...` naming convention instead of `First, Second, Third...` for consistency and clarity when dealing with many values.

## Lazy Evaluation

### Lazy[T]

Represents a value that may be computed immediately (eager) or on first access (deferred). Provides a unified interface for both strategies with automatic caching and thread-safe deferred evaluation.

**Key Features:**
- **Deferred computation**: Compute values only when needed
- **Automatic caching**: Results cached after first computation
- **Thread-safe**: Uses `sync.Once` for safe concurrent access
- **Error handling**: Supports computations that may fail
- **Unified interface**: Treat eager and deferred values uniformly

#### Creating Lazy Values

##### LazyEager - Immediate Value

Wraps an already-computed value with no deferred computation.

```go
// Wrap an immediate value
lazy := lxtypes.LazyEager(42)
value, _ := lazy.Get()  // Returns immediately: 42

// Check if evaluated (always true for eager)
fmt.Println(lazy.IsEvaluated())  // true
```

##### LazyEagerOrError - Wrap Existing Results

Convert an existing `(value, error)` result into a Lazy.

```go
result, err := someFunction()
lazy := lxtypes.LazyEagerOrError(result, err)

// Later access
value, err := lazy.Get()  // Returns result and err immediately
```

##### LazyDeferred - Deferred Computation

Compute value lazily on first access. The computation function is called at most once.

```go
// Create lazy computation
expensive := lxtypes.LazyDeferred(func() (int, error) {
    // Expensive operation (database query, file read, etc.)
    time.Sleep(time.Second)
    return 42, nil
})

fmt.Println(expensive.IsEvaluated())  // false (not computed yet)

// Compute on first access
value, err := expensive.Get()  // Takes 1 second
fmt.Println(value)                    // 42

// Subsequent calls return cached value
value2, _ := expensive.Get()  // Returns instantly
fmt.Println(expensive.IsEvaluated())  // true
```

#### Methods

##### Get() (T, error)

Returns the value, computing it if necessary. For deferred values, the computation happens on first call and the result is cached.

```go
lazy := lxtypes.LazyDeferred(func() (string, error) {
    return "computed", nil
})

value, err := lazy.Get()
if err != nil {
    // Handle error
}
fmt.Println(value)  // "computed"
```

##### MustGet() T

Returns the value, panicking if computation fails. Useful when you're certain the computation will succeed.

```go
lazy := lxtypes.LazyDeferred(func() (int, error) {
    return 42, nil
})

value := lazy.MustGet()  // 42 (panics if error)
```

##### IsEvaluated() bool

Returns true if the value has been computed (for deferred) or was provided immediately (for eager).

```go
lazy := lxtypes.LazyDeferred(func() (int, error) {
    return 42, nil
})

fmt.Println(lazy.IsEvaluated())  // false
lazy.Get()
fmt.Println(lazy.IsEvaluated())  // true
```

#### Use Cases

**1. Expensive Resource Initialization**

```go
// Database connection pool that's only created when needed
dbPool := lxtypes.LazyDeferred(func() (*sql.DB, error) {
    return sql.Open("postgres", connectionString)
})

// In fast path, database never initialized
if cachedResult, ok := cache.Get(key); ok {
    return cachedResult
}

// In slow path, database initialized on first access
db, err := dbPool.Get()
if err != nil {
    return err
}
result := db.Query(...)
```

**2. Configuration Loading**

```go
type Config struct {
    APIKey string
    Endpoint string
}

// Config only loaded when actually needed
config := lxtypes.LazyDeferred(func() (Config, error) {
    return loadConfigFromFile("config.json")
})

// Conditional usage
if needsAPI {
    cfg := config.MustGet()
    callAPI(cfg.APIKey, cfg.Endpoint)
}
```

**3. Conditional Computations**

```go
// Expensive fallback that's only computed if primary fails
fallbackData := lxtypes.LazyDeferred(func() ([]byte, error) {
    return fetchFromSlowBackup()
})

// Try primary source
data, err := fetchFromPrimary()
if err != nil {
    // Fallback computation only happens here
    data, err = fallbackData.Get()
}
```

**4. Thread-Safe Singletons**

```go
var logger = lxtypes.LazyDeferred(func() (*Logger, error) {
    return NewLogger("app.log")
})

// Safe to call from multiple goroutines
func Log(msg string) {
    log := logger.MustGet()  // Initialized exactly once
    log.Write(msg)
}
```

**5. Caching Expensive Computations**

```go
// Compute once, reuse many times
fibonacci := lxtypes.LazyDeferred(func() ([]int, error) {
    // Generate first 100 fibonacci numbers
    fib := make([]int, 100)
    fib[0], fib[1] = 0, 1
    for i := 2; i < 100; i++ {
        fib[i] = fib[i-1] + fib[i-2]
    }
    return fib, nil
})

// Use multiple times without recomputation
for i := 0; i < 10; i++ {
    fib, _ := fibonacci.Get()  // Only computed once
    fmt.Println(fib[i])
}
```

#### Error Handling

Lazy supports operations that may fail. Errors are cached along with values.

```go
// Computation that may fail
lazy := lxtypes.LazyDeferred(func() (Data, error) {
    data, err := loadData()
    if err != nil {
        return Data{}, fmt.Errorf("failed to load: %w", err)
    }
    return data, nil
})

// Handle error on access
data, err := lazy.Get()
if err != nil {
    log.Printf("Error: %v", err)
    return
}

// Or use MustGet if failure should panic
data := lazy.MustGet()
```

#### Thread Safety

`LazyDeferred` is thread-safe and ensures the computation function is called exactly once, even with concurrent access.

```go
lazy := lxtypes.LazyDeferred(func() (int, error) {
    time.Sleep(100 * time.Millisecond)
    return 42, nil
})

// Launch 100 concurrent goroutines
var wg sync.WaitGroup
for i := 0; i < 100; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        value, _ := lazy.Get()  // Function called exactly once
        fmt.Println(value)
    }()
}
wg.Wait()
```

#### Eager vs Deferred Comparison

| Feature | LazyEager | LazyDeferred |
|---------|-----------|--------------|
| Computation timing | Immediate | On first `Get()` |
| Performance | Instant access | First access has delay |
| Use case | Wrapping existing values | Expensive computations |
| `IsEvaluated()` | Always `true` | `false` until first `Get()` |
| Thread safety | N/A (immutable) | Thread-safe with `sync.Once` |
| Memory | Stores value immediately | Stores function until evaluation |

**When to use LazyEager:**
- Wrapping existing computation results
- Converting `(value, error)` pairs to Lazy interface
- No benefit from deferring (value already available)

**When to use LazyDeferred:**
- Expensive computations (database queries, file I/O, API calls)
- Conditional usage (may not be needed in all code paths)
- Singleton initialization
- Caching computed values across multiple accesses

## Async Operations

### Future[T]

Represents a value that will be available in the future through asynchronous computation. Futures support sequential composition through `FutureThen`, parallel execution with `FutureAll` and `FutureJoin*`, and context-aware cancellation.

**Key Features:**
- **Hot start**: Computation begins immediately when future is created
- **Lock-free**: Efficient channel-based synchronization
- **Context-aware**: Respects context cancellation and timeouts
- **Type transformations**: Chain operations with different types via `FutureThen`
- **Parallel composition**: Combine multiple futures with `FutureAll` and `FutureJoin`
- **Error propagation**: Errors flow through transformation chains

#### Creating Futures

##### FutureDo - Async Computation

Execute a function asynchronously. The computation starts immediately in a background goroutine.

```go
// Start async computation (hot start)
future := lxtypes.FutureDo(func() (int, error) {
    // Simulate expensive operation (API call, database query, etc.)
    time.Sleep(100 * time.Millisecond)
    return 42, nil
})

// Do other work while computation runs...

// Get the result (blocks until ready)
ctx := context.Background()
result, err := future.Get(ctx)
if err != nil {
    // Handle error
}
fmt.Println(result)  // 42
```

##### FutureOf - Immediate Value

Create a future that's already completed with a value. No goroutine is started.

```go
future := lxtypes.FutureOf(100)

ctx := context.Background()
result, _ := future.Get(ctx)  // Returns immediately
fmt.Println(result)  // 100
```

##### FutureError - Immediate Error

Create a future that's already completed with an error.

```go
future := lxtypes.FutureError[int](errors.New("failed"))

ctx := context.Background()
_, err := future.Get(ctx)  // Returns immediately
fmt.Println(err)  // failed
```

#### Sequential Composition with FutureThen

Chain dependent operations that transform types. Each step runs only after the previous succeeds.

```go
// Step 1: Get user ID
userIdFuture := lxtypes.FutureDo(func() (int, error) {
    return 123, nil
})

// Step 2: Fetch user details using ID (int -> User)
userFuture := lxtypes.FutureThen(userIdFuture, func(id int) (User, error) {
    time.Sleep(50 * time.Millisecond)
    return User{ID: id, Name: "Alice"}, nil
})

// Step 3: Fetch user's card (User -> Card)
cardFuture := lxtypes.FutureThen(userFuture, func(user User) (Card, error) {
    time.Sleep(50 * time.Millisecond)
    return Card{UserID: user.ID, Number: "****1234"}, nil
})

ctx := context.Background()
card, err := cardFuture.Get(ctx)
// Total time: ~100ms (sequential), card contains result
```

**Error Propagation:**

If any step in the chain fails, the error propagates and subsequent transformations don't execute.

```go
future := lxtypes.FutureDo(func() (int, error) {
    return 0, errors.New("initial error")
})

// This transformation won't execute
future2 := lxtypes.FutureThen(future, func(n int) (string, error) {
    return "won't run", nil
})

ctx := context.Background()
_, err := future2.Get(ctx)
fmt.Println(err)  // initial error
```

#### Parallel Execution

##### FutureAll - Same Type

Execute multiple futures of the same type concurrently and combine results into a slice.

```go
// Start 3 parallel operations
service1 := lxtypes.FutureDo(func() (Data, error) {
    return fetchFromService1()  // 100ms
})
service2 := lxtypes.FutureDo(func() (Data, error) {
    return fetchFromService2()  // 100ms
})
service3 := lxtypes.FutureDo(func() (Data, error) {
    return fetchFromService3()  // 100ms
})

// Combine all results
allData := lxtypes.FutureAll(service1, service2, service3)

ctx := context.Background()
results, err := allData.Get(ctx)
// Total time: ~100ms (parallel, not 300ms)
// results: []Data{data1, data2, data3}

// Transform combined results
response := lxtypes.FutureThen(allData, func(data []Data) (Response, error) {
    return combineData(data), nil
})
```

**Error Handling:**

If any future fails, `FutureAll` returns the first error encountered. All futures continue executing in the background.

```go
f1 := lxtypes.FutureDo(func() (int, error) { return 1, nil })
f2 := lxtypes.FutureDo(func() (int, error) { return 0, errors.New("failed") })
f3 := lxtypes.FutureDo(func() (int, error) { return 3, nil })

allFuture := lxtypes.FutureAll(f1, f2, f3)

ctx := context.Background()
_, err := allFuture.Get(ctx)
fmt.Println(err)  // failed
```

##### FutureJoin2 - Two Different Types

Combine two futures of different types into a `Pair`.

```go
userFuture := lxtypes.FutureDo(func() (User, error) {
    return fetchUser()  // 50ms
})

configFuture := lxtypes.FutureDo(func() (Config, error) {
    return fetchConfig()  // 50ms
})

// Join into Pair[User, Config]
combined := lxtypes.FutureJoin2(userFuture, configFuture)

ctx := context.Background()
result, err := combined.Get(ctx)
// result.First = User, result.Second = Config
```

##### FutureJoin3 - Three Different Types

Combine three futures into a `Triple`.

```go
userFuture := lxtypes.FutureDo(func() (User, error) {
    return fetchUser()
})

configFuture := lxtypes.FutureDo(func() (Config, error) {
    return fetchConfig()
})

statsFuture := lxtypes.FutureDo(func() (Stats, error) {
    return fetchStats()
})

// Join into Triple[User, Config, Stats]
combined := lxtypes.FutureJoin3(userFuture, configFuture, statsFuture)

ctx := context.Background()
result, err := combined.Get(ctx)
// result.First = User, result.Second = Config, result.Third = Stats
```

##### FutureJoin4 - Four Different Types

Combine four futures into a `Quad`.

```go
joined := lxtypes.FutureJoin4(future1, future2, future3, future4)

ctx := context.Background()
result, err := joined.Get(ctx)
// result.First, result.Second, result.Third, result.Fourth
```

##### FutureJoin5 - Five Different Types

Combine five futures of different types into a `Tuple5`. Perfect for coordinating multiple microservices.

```go
// Fetch from 5 different services
userFuture := lxtypes.FutureDo(func() (User, error) {
    return fetchUser()
})
ordersFuture := lxtypes.FutureDo(func() ([]Order, error) {
    return fetchOrders()
})
paymentFuture := lxtypes.FutureDo(func() (Payment, error) {
    return fetchPayment()
})
inventoryFuture := lxtypes.FutureDo(func() (Inventory, error) {
    return fetchInventory()
})
recommendationsFuture := lxtypes.FutureDo(func() ([]Product, error) {
    return fetchRecommendations()
})

// Combine all 5 futures
combined := lxtypes.FutureJoin5(userFuture, ordersFuture, paymentFuture, inventoryFuture, recommendationsFuture)

ctx := context.Background()
result, err := combined.Get(ctx)
// Access: result.V1 (User), result.V2 ([]Order), result.V3 (Payment), result.V4 (Inventory), result.V5 ([]Product)

// Transform into dashboard response
response := lxtypes.FutureThen(combined, func(data lxtypes.Tuple5[User, []Order, Payment, Inventory, []Product]) (Dashboard, error) {
    return Dashboard{
        User:            data.V1,
        Orders:          data.V2,
        Payment:         data.V3,
        Inventory:       data.V4,
        Recommendations: data.V5,
    }, nil
})
```

##### FutureJoin6 - Six Different Types

Combine six futures into a `Tuple6`.

```go
combined := lxtypes.FutureJoin6(f1, f2, f3, f4, f5, f6)
result, err := combined.Get(ctx)
// Access: result.V1, result.V2, result.V3, result.V4, result.V5, result.V6
```

##### FutureJoin7 - Seven Different Types

Combine seven futures into a `Tuple7`.

```go
combined := lxtypes.FutureJoin7(f1, f2, f3, f4, f5, f6, f7)
result, err := combined.Get(ctx)
// Access: result.V1 through result.V7
```

##### FutureJoin8 - Eight Different Types

Combine eight futures into a `Tuple8`. Maximum flexibility for complex parallel operations.

```go
combined := lxtypes.FutureJoin8(f1, f2, f3, f4, f5, f6, f7, f8)
result, err := combined.Get(ctx)
// Access: result.V1 through result.V8
```

##### FutureAny - First successful result

Execute multiple futures in parallel and return the first successfully completed
result (the first child whose error is nil). This is useful when you have
multiple sources or fallbacks and you only care about the first successful
response (similar to JavaScript's Promise.any semantics).

Example:

```go
// Try three sources in parallel and take the first successful value
f1 := lxtypes.FutureDo(func() (string, error) {
    time.Sleep(50 * time.Millisecond)
    return "slow", nil
})

f2 := lxtypes.FutureDo(func() (string, error) {
    // Fast success
    return "fast", nil
})

f3 := lxtypes.FutureDo(func() (string, error) {
    return "", errors.New("failed")
})

any := lxtypes.FutureAny(f1, f2, f3)
ctx := context.Background()
val, err := any.Get(ctx)
if err != nil {
    // All sources failed
}
fmt.Println(val) // "fast"
```

Behavior & semantics:

- Returns as soon as a child future completes successfully (err == nil).
- If all child futures complete with non-nil errors, `FutureAny` returns the
  first encountered error.
- If no futures are provided, `FutureAny` returns a failed future immediately
  (error: "lxtypes: no futures provided").
- Child futures are not cancelled when `FutureAny` returns; they continue
  running in the background. The returned future's `Get(ctx)` respects the
  provided context (cancellation and timeouts) for that call only.

Error handling:

- `FutureAny` returns the first successful value; when no success occurs it
  returns the first error observed (consistent with the project's `FutureAll`
  behavior of returning a first error). If you prefer aggregated errors or
  different empty-input semantics we can provide variants.

##### Which One to Use?

| Scenario | Function | Return Type | Example |
|----------|----------|-------------|---------|
| **All same type** | `FutureAll` | `Future[[]T]` | Fetch from N database shards, all returning `[]Record` |
| **2 different types** | `FutureJoin2` | `Future[Pair[T, U]]` | Fetch `User` and `Config` in parallel |
| **3 different types** | `FutureJoin3` | `Future[Triple[T, U, V]]` | Fetch `User`, `Config`, and `Stats` in parallel |
| **4 different types** | `FutureJoin4` | `Future[Quad[T, U, V, W]]` | Fetch `User`, `Config`, `Stats`, and `Permissions` |
| **5 different types** | `FutureJoin5` | `Future[Tuple5[...]]` | Fetch from 5 different microservices |
| **6 different types** | `FutureJoin6` | `Future[Tuple6[...]]` | Fetch from 6 different services |
| **7 different types** | `FutureJoin7` | `Future[Tuple7[...]]` | Fetch from 7 different services |
| **8 different types** | `FutureJoin8` | `Future[Tuple8[...]]` | Maximum 8 different services in one call |
| **9+ different types** | Nest `FutureJoin*` calls | Multiple tuples | Combine smaller tuples into larger structures |

**Example - Fetching from 5 services (Dashboard Use Case):**

```go
// Start all 5 service calls in parallel
userFuture := lxtypes.FutureDo(func() (User, error) {
    return fetchUserService()
})
ordersFuture := lxtypes.FutureDo(func() ([]Order, error) {
    return fetchOrderService()
})
paymentFuture := lxtypes.FutureDo(func() (Payment, error) {
    return fetchPaymentService()
})
inventoryFuture := lxtypes.FutureDo(func() (Inventory, error) {
    return fetchInventoryService()
})
recommendationsFuture := lxtypes.FutureDo(func() ([]Product, error) {
    return fetchRecommendationService()
})

// Combine all 5 - all execute in parallel
allData := lxtypes.FutureJoin5(userFuture, ordersFuture, paymentFuture, inventoryFuture, recommendationsFuture)

// Transform into response
response := lxtypes.FutureThen(allData, func(data lxtypes.Tuple5[User, []Order, Payment, Inventory, []Product]) (DashboardResponse, error) {
    return DashboardResponse{
        User:            data.V1,
        Orders:          data.V2,
        Payment:         data.V3,
        Inventory:       data.V4,
        Recommendations: data.V5,
    }, nil
})

ctx := context.Background()
result, err := response.Get(ctx)
```

**Example - Mixing same and different types:**

```go
// Start 2 same-type futures
f1 := lxtypes.FutureDo(func() (int, error) {
    return 1, nil
})
f2 := lxtypes.FutureDo(func() (int, error) {
    return 2, nil
})

// Start 1 different-type future
f3 := lxtypes.FutureDo(func() (string, error) {
    return "three", nil
})

// Combine first two with FutureAll (same type)
allInts := lxtypes.FutureAll(f1, f2)

// Combine all three with FutureJoin3 (mixed types)
allMixed := lxtypes.FutureJoin3(f1, f2, f3)

// Access results
ints, _ := allInts.Get(ctx)  // []int{1, 2}
mixed, _ := allMixed.Get(ctx)  // Triple[int, int, string]
```

## Mutable State

### Ref[T]

`Ref[T]` is a thread-safe mutable value cell. It wraps a single value of any type and protects concurrent access with a read-write mutex. Unlike `Lazy[T]` (compute-once) or `Future[T]` (async computation), `Ref[T]` is designed for **shared mutable state** that changes over time.

#### Creating a Ref

```go
counter := lxtypes.NewRef(0)
config  := lxtypes.NewRef(AppConfig{Host: "localhost", Port: 8080})
```

#### Methods

| Method | Signature | Description |
|--------|-----------|-------------|
| `Get` | `Get() T` | Returns the current value (read-lock) |
| `Set` | `Set(value T)` | Replaces the value (write-lock) |
| `Update` | `Update(fn func(T) T)` | Atomically transforms the value (write-lock) |

#### Use Case 1 — Shared counter across goroutines

```go
counter := lxtypes.NewRef(0)

var wg sync.WaitGroup
for i := 0; i < 100; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        counter.Update(func(v int) int { return v + 1 })
    }()
}
wg.Wait()

fmt.Println(counter.Get())  // 100
```

#### Use Case 2 — Hot-reloadable config / shared state

```go
type AppConfig struct {
    Host    string
    Timeout time.Duration
}

cfg := lxtypes.NewRef(AppConfig{Host: "localhost", Timeout: 5 * time.Second})

// Reload on signal (in another goroutine)
go func() {
    newCfg := loadConfigFromFile()
    cfg.Set(newCfg)
}()

// Use the current config safely
current := cfg.Get()
fmt.Println("Connecting to", current.Host)
```

#### Use Case 3 — Capturing old value inside Update

```go
var auditLog []string

counter := lxtypes.NewRef(0)

counter.Update(func(v int) int {
    // Capture the old value for auditing before returning the new value
    auditLog = append(auditLog, fmt.Sprintf("changed from %d", v))
    return v + 1
})

fmt.Println(counter.Get())    // 1
fmt.Println(auditLog[0])      // changed from 0
```

#### Comparison: Ref[T] vs Lazy[T] vs Future[T]

| Feature | `Ref[T]` | `Lazy[T]` | `Future[T]` |
|---------|----------|-----------|-------------|
| **Mutable** | ✅ Yes | ❌ No (compute-once) | ❌ No (compute-once) |
| **Thread-safe** | ✅ Yes | ✅ Yes | ✅ Yes |
| **Purpose** | Shared mutable state | Deferred / cached value | Async computation |
| **Concurrency** | Multiple reads/writes | Safe single computation | Background goroutine |
| **Get** | Returns current value | Computes or returns cache | Blocks until done |

#### When to Use Ref[T]

- You need **shared mutable state** that multiple goroutines read and write.
- You want to **atomically transform** a value without external locking.
- You need a **hot-reloadable** configuration or feature flag that changes at runtime.
- You are building a **counter**, **accumulator**, or any value that evolves over time.

Prefer `Lazy[T]` when the value is computed once and never changed. Prefer `Future[T]` when the value is the result of a single asynchronous operation. Use `Ref[T]` when the value genuinely changes over time and must be safe for concurrent access.
