# lxconstraints

Generic type constraints for type-safe Go programming.

## Overview

The `lxconstraints` package provides reusable type constraints that enable writing generic functions with proper type safety. These constraints go beyond Go's built-in `any` and `comparable` to provide fine-grained control over what types your generic functions can accept.

## Installation

Install the lx module (includes all packages):

```bash
go get github.com/nthanhhai2909/lx
```

Then import only what you need:

```go
import "github.com/nthanhhai2909/lx/lxconstraints"
```

## Available Constraints

### Numeric Constraints

#### `Integer`
All signed and unsigned integer types.

**Includes**: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`

```go
func Sum[T lxconstraints.Integer](a, b T) T {
    return a + b
}
```

#### `Signed`
Only signed integer types (can represent negative numbers).

**Includes**: `int`, `int8`, `int16`, `int32`, `int64`

```go
func Abs[T lxconstraints.Signed](n T) T {
    if n < 0 {
        return -n
    }
    return n
}
```

#### `Unsigned`
Only unsigned integer types (always non-negative).

**Includes**: `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`

```go
func HalfOf[T lxconstraints.Unsigned](n T) T {
    return n / 2
}
```

#### `Float`
All floating-point types.

**Includes**: `float32`, `float64`

```go
func Average[T lxconstraints.Float](values []T) T {
    var sum T
    for _, v := range values {
        sum += v
    }
    return sum / T(len(values))
}
```

#### `Complex`
All complex number types.

**Includes**: `complex64`, `complex128`

```go
func Magnitude[T lxconstraints.Complex](c T) float64 {
    return math.Abs(real(c)) + math.Abs(imag(c))
}
```

#### `Number`
All numeric types (integers and floats, but not complex).

**Includes**: All `Integer` and `Float` types

```go
func Min[T lxconstraints.Number](a, b T) T {
    if a < b {
        return a
    }
    return b
}
```

#### `Numeric`
All numeric types including complex numbers.

**Includes**: All `Integer`, `Float`, and `Complex` types

```go
func Double[T lxconstraints.Numeric](n T) T {
    return n + n
}
```

### Ordering and Operations

#### `Ordered`
Types that support comparison operations (`<`, `<=`, `>`, `>=`).

**Includes**: All `Number` types and `string`

```go
func Clamp[T lxconstraints.Ordered](value, min, max T) T {
    if value < min {
        return min
    }
    if value > max {
        return max
    }
    return value
}
```

#### `Addable`
Types that support the `+` operator.

**Includes**: All `Numeric` types and `string` (for concatenation)

```go
func Concat[T lxconstraints.Addable](a, b T) T {
    return a + b
}
```

### Collection Constraints

#### `Slice[T]`
Any slice type with elements of type `T`.

```go
func IsEmpty[T any, S lxconstraints.Slice[T]](s S) bool {
    return len(s) == 0
}

func Reverse[T any, S lxconstraints.Slice[T]](s S) S {
    result := make(S, len(s))
    for i, v := range s {
        result[len(s)-1-i] = v
    }
    return result
}
```

#### `Map[K, V]`
Any map type with keys of type `K` and values of type `V`.

```go
func Keys[K comparable, V any, M lxconstraints.Map[K, V]](m M) []K {
    keys := make([]K, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}
```

#### `Chan[T]`
Any channel type (bidirectional, send-only, or receive-only).

```go
func Drain[T any, C lxconstraints.Chan[T]](ch C) []T {
    var result []T
    for v := range ch {
        result = append(result, v)
    }
    return result
}
```

#### `Pointer[T]`
Any pointer type to `T`.

```go
func DerefOr[T any, P lxconstraints.Pointer[T]](p P, defaultValue T) T {
    if p == nil {
        return defaultValue
    }
    return *p
}
```

## Usage Examples

### Generic Math Function

```go
func Sum[T lxconstraints.Number](values []T) T {
    var total T
    for _, v := range values {
        total += v
    }
    return total
}

// Works with any numeric type
intSum := Sum([]int{1, 2, 3, 4, 5})           // 15
floatSum := Sum([]float64{1.5, 2.5, 3.5})     // 7.5
```

### Generic Sorting

```go
func BubbleSort[T lxconstraints.Ordered](slice []T) {
    n := len(slice)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if slice[j] > slice[j+1] {
                slice[j], slice[j+1] = slice[j+1], slice[j]
            }
        }
    }
}

// Works with numbers and strings
numbers := []int{5, 2, 8, 1, 9}
BubbleSort(numbers)  // [1, 2, 5, 8, 9]

words := []string{"banana", "apple", "cherry"}
BubbleSort(words)  // ["apple", "banana", "cherry"]
```

### Generic Collection Operations

```go
func Filter[T any, S lxconstraints.Slice[T]](s S, predicate func(T) bool) S {
    var result S
    for _, v := range s {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}
```

### Working with Named Types

All constraints support named types using the `~` operator:

```go
type UserID int64
type Score float64

func AddUserIDs[T lxconstraints.Integer](a, b T) T {
    return a + b
}

var id1 UserID = 100
var id2 UserID = 200
result := AddUserIDs(id1, id2)  // Works! result is UserID(300)
```

## Comparison with Standard Library

| lxconstraints | Go stdlib | Notes |
|---------------|-----------|-------|
| `Integer` | - | All integer types |
| `Signed` | - | Only signed integers |
| `Unsigned` | - | Only unsigned integers |
| `Float` | - | All float types |
| `Complex` | - | Complex number types |
| `Number` | - | Integers and floats |
| `Numeric` | - | All numeric types including complex |
| `Ordered` | `cmp.Ordered` (Go 1.21+) | Numbers and strings |
| `Addable` | - | Types supporting `+` operator |
| `Slice[T]` | - | Parameterized slice types |
| `Map[K,V]` | - | Parameterized map types |
| `Chan[T]` | - | All channel types |
| `Pointer[T]` | - | Pointer types |

## Design Principles

1. **Type Safety**: Constraints prevent invalid operations at compile time
2. **Composable**: Constraints can be combined in function signatures
3. **Named Type Support**: The `~` operator allows named types with matching underlying types
4. **Zero Cost**: Constraints are compile-time only with no runtime overhead
5. **Idiomatic**: Follows Go's generics best practices

## Related Packages

- **[lxslices](../lxslices)**: Generic slice operations using these constraints
- **[lxtypes](../lxtypes)**: Functional programming types
- **[lxptrs](../ptrs)**: Pointer utilities

## Contributing

See the main [CONTRIBUTING.md](../CONTRIBUTING.md) for guidelines.

## License

Apache 2.0 - See [LICENSE](../LICENSE) for details.

