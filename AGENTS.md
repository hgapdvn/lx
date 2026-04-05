# AGENTS.md - AI Coding Agent Guide

**Essential knowledge for AI agents working with the lx codebase**

---

## 🎯 Project at a Glance

**lx** is a Go 1.18+ utilities library providing focused, generic-first extensions to the standard library. Every package solves one problem well with zero external dependencies.

- **Core Principle**: Small, focused packages (no giant "utils")
- **Generic-First**: All functions use Go 1.18+ generics for type safety
- **Backwards Compatible**: Targets Go 1.18+ (use `lxslices`/`lxmaps` over stdlib `slices`/`maps` packages from Go 1.21)
- **Production-Ready**: 90%+ test coverage, comprehensive edge case testing

---

## 🏗️ Architecture & Package Layout

### Current Packages (All Stable ✅)

```
lxconstraints/ → Type constraints (Integer, Number, Ordered, Comparable, Addable, Signed, Unsigned, Float)
lxenv/         → Environment variables with typed parsing (Get, GetInt, GetBool, GetDuration)
lxmaps/        → Map operations (Keys, Values, Filter, Merge, Pick, Omit, etc.)
lxptrs/        → Pointer utilities (Ref, Deref, DerefOr, IsNil)
lxslices/      → Slice operations (Filter, Map, Reduce, Chunk, Window, etc.)
lxstrings/     → String utilities (IsBlank, Capitalize, TrimPrefix, etc.)
lxsystems/     → System info (OS detection, environment, paths)
lxtime/        → Time utilities (Days, Hours, Minutes, StartOfDay, EndOfMonth, etc.)
lxtypes/       → Functional & async types (Lazy, Optional, Result, Either, Future, Pair, Triple, etc.)
```

### File Organization Pattern

Each package follows this structure:
```
packagename/
├── main_feature_1.go        # Grouped functions by feature
├── main_feature_1_test.go   # Corresponding tests
├── main_feature_2.go
├── main_feature_2_test.go
├── doc.go                   # Package-level documentation
├── errs.go                  # Package-specific errors (if needed)
└── README.md                # Package-specific examples (if needed)
```

Example: `lxslices/` is organized as:
- `filter.go` (Find, Filter, Partition, Any, All, None, Count)
- `transform.go` (Map, FlatMap, GroupBy, Chunk, Window, Zip, Unzip)
- `manipulation.go` (Append, Prepend, Insert, Remove, Replace, Rotate)
- `aggregation.go` (Reduce, Sum, Average, Min, Max)
- And corresponding `*_test.go` files...

---

## 🎨 Critical Patterns & Conventions

### 1. The (Value, bool) Pattern

**For operations that might fail on valid input**:
```go
func First[T any](slice []T) (T, bool)
func Find[T any](slice []T, predicate func(T) bool) (T, bool)
func Get[K comparable, V any](m map[K]V, key K) (V, bool)
```

**Always provide convenience variants with defaults**:
```go
func First[T any](slice []T) (T, bool)
func FirstOr[T any](slice []T, defaultValue T) T
func MustFirst[T any](slice []T) T  // Panics if not found
```

### 2. Nil-Safe Semantics (CRITICAL)

**Always preserve nil vs empty distinction**:
```go
// ✅ Correct - preserves nil
func Filter[T any](slice []T, predicate func(T) bool) []T {
    if slice == nil {
        return nil  // Not []T{}
    }
    // ... build result
}

// Returns nil if input is nil
// Returns empty slice if input is empty OR no matches
```

**Document nil behavior explicitly** in every function's doc comment.

### 3. Generic Type Constraints

Use these from `lxconstraints`:
```go
Integer    → int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr
Signed     → Signed integer types only (int, int8, int16, int32, int64)
Unsigned   → Unsigned integer types only (uint, uint8, uint16, uint32, uint64, uintptr)
Number     → Integer + float32, float64, complex64, complex128
Float      → float32, float64
Addable    → Types that support the + operator (Number types)
Ordered    → Comparable types with <, >, <=, >= operators
Comparable → Types with == operator (suitable for map keys)
```

**Example usage**:
```go
func Sum[T lxconstraints.Number](slice []T) T
func Sort[T lxconstraints.Ordered](slice []T) []T
func Abs[T lxconstraints.Signed](n T) T
func Unique[T comparable](slice []T) []T
func Map[T, U any](slice []T, fn func(T) U) []U  // Most transformations use 'any'
```

### 4. Functional Patterns for Type Safety

These exist in `lxtypes`:
- `Predicate[T]` → `func(T) bool` (use when documenting predicates)
- `Function[T, U]` → `func(T) U` (for transformations)
- `Optional[T]` → Nullable type wrapper (comma-ok pattern: `Get()` returns `(T, bool)`)
- `Result[T]` → Error handling wrapper (Go pattern: `Value()` returns `(T, error)`)
- `Either[L, R]` → Discriminated union type
- `Lazy[T]` → Deferred or immediate computation with caching
- `Future[T]` → Asynchronous computation with type-safe composition and context support
- `Pair[K, V]`, `Triple[T, U, V]`, `Quad[T, U, V, W]` → Tuple types for 2-4 values
- `Tuple5[T1, T2, ...]` through `Tuple8[...]` → Extended tuples for 5-8 values
- `Ref[T]` → Thread-safe mutable value cell

### 5. Error Handling Pattern

Define errors in `errs.go`:
```go
package lxslices

var (
    ErrDuplicateKey = errors.New("lxslices: duplicate key")
    ErrInvalidSize = errors.New("lxslices: invalid size")
)
```

Return errors for operations that can fail on valid input:
```go
func UniqueGroupBy[T any, K comparable](slice []T, fn func(T) K) (map[K]T, error)
```

### 6. Documentation Template (REQUIRED)

Every exported function must have:
```go
// FunctionName does X and returns Y.
// Details about behavior, edge cases, nil handling.
// Returns (value, found) semantics if applicable.
//
// Example:
//
//     result := FunctionName(input)
//     // result: expected output
//
func FunctionName[T any](param T) T {
```

---

## 🧪 Testing Standards (CRITICAL - 100% Coverage)

### Test File Naming & Organization

- Test files: `*_test.go` (same package declared as `package_name_test`)
- Example test files: `*_examples_test.go` (for runnable documentation examples, declared as `package_name_test`)
- Import the package being tested: `"github.com/hgapdvn/lx/lxslices"`
- Tests use external package import (not `.` import)

Example test file pattern:
```go
package lxslices_test

import "github.com/hgapdvn/lx/lxslices"

// ExampleFilter demonstrates using the Filter function
func ExampleFilter() {
    numbers := []int{1, 2, 3, 4, 5}
    evens := slices.Filter(numbers, func(n int) bool { return n%2 == 0 })
    // evens: [2, 4]
}
```

### Table-Driven Test Pattern (MANDATORY)

```go
func TestFind_Int(t *testing.T) {
    tests := []struct {
        name      string
        slice     []int
        predicate func(int) bool
        expected  struct {
            value int
            found bool
        }
    }{
        {
            name:      "find at beginning",
            slice:     []int{1, 2, 3},
            predicate: func(v int) bool { return v == 1 },
            expected:  struct{value int; found bool}{1, true},
        },
        {
            name:      "nil slice",
            slice:     nil,
            predicate: func(v int) bool { return true },
            expected:  struct{value int; found bool}{0, false},
        },
        {
            name:      "empty slice",
            slice:     []int{},
            predicate: func(v int) bool { return true },
            expected:  struct{value int; found bool}{0, false},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            value, found := slices.Find(tt.slice, tt.predicate)
            // Assert...
        })
    }
}
```

### What MUST Be Tested

✅ **Required**:
- Happy path (normal usage)
- `nil` input
- Empty input
- Single element
- Multiple elements
- Boundary conditions (first, last, min, max)
- Error conditions
- Different type parameters (for generics): test with `int`, `string`, custom structs

❌ **Never skip edge cases** - this is non-negotiable for 100% coverage.

### Coverage Commands

```bash
go test ./...                           # Run all tests
go test -cover ./...                    # Show coverage percentages
go test -coverprofile=coverage.out ./... # Generate coverage file
go tool cover -html=coverage.out        # Open HTML coverage report
```

---

## 🔧 Common Workflows

### Adding a New Function

1. **Check existing patterns** - Review similar functions in the file
2. **Write tests first** - Table-driven test with edge cases
3. **Implement function** - Follow nil safety and constraint patterns
4. **Write documentation** - With example code
5. **Verify coverage** - Must reach 100% on new code
6. **Format & lint**:
   ```bash
   gofmt -w .          # Format
   go vet ./...        # Lint
   go test -race ./... # Race condition detector
   ```

### Adding a New Package

1. Create directory `packagename/`
2. Create files: `main.go`, `main_test.go`, `doc.go`
3. Define package-level doc comment in `doc.go`
4. Start with 2-3 core functions with tests
5. Aim for 90%+ coverage before merge
6. Consider package-specific errors in `errs.go` if needed
7. For complex types or rich APIs, create `*_examples_test.go` files to demonstrate usage patterns (e.g., `types/lazy_examples_test.go`, `constraints/examples_test.go`)

### Creating Example Test Files

Use `*_examples_test.go` files to provide runnable documentation examples:
- Create example functions: `func ExampleFunctionName() { ... }`
- These appear in generated documentation and can be run with `go test`
- Use for demonstrating complex types, functional patterns, or composition
- Showing typical workflows or patterns

**When to use**:
- Complex packages with many functions (types, constraints)
- Demonstrating function composition or method chaining
- Showing typical workflows or patterns

**When not needed**:
- Simple utilities with straightforward usage
- Already well-documented with inline examples
- Few exported functions

### Running Tests for Development

```bash
# Single package
go test -v ./lxslices/

# All with coverage
go test -cover ./...

# Specific test
go test -run TestFilter_Int ./lxslices/

# With race detector (catches concurrency bugs)
go test -race ./...
```

---

## 📋 Key Files to Know

| File | Purpose |
|------|---------|
| `.github/copilot-instructions.md` | Comprehensive coding guidelines (read first!) |
| `CONTRIBUTING.md` | Contribution process and standards |
| `PACKAGE_ROADMAP.md` | Future packages and strategic direction |
| `go.mod` | Module definition: `module github.com/hgapdvn/lx` |
| `slices/doc.go` | Example of package organization documentation |
| `slices/filter.go` | Example of core functions pattern |
| `slices/filter_test.go` | Example of comprehensive testing pattern |
| `time/days.go` | Example of simple utility functions (lxtime package) |
| `types/doc.go` | Example of complex package documentation with categories |
| `types/lazy_examples_test.go` | Example of runnable documentation test files |
| `constraints/examples_test.go` | Example of constraint usage patterns |

---

## 🚨 Common Pitfalls to Avoid

### 1. Breaking Nil Semantics
```go
// ❌ WRONG
func Filter[T any](slice []T, predicate func(T) bool) []T {
    result := []T{}  // Returns non-nil empty slice
    // ...
}

// ✅ CORRECT
func Filter[T any](slice []T, predicate func(T) bool) []T {
    if slice == nil {
        return nil
    }
    var result []T  // Will be nil if nothing added
    // ...
}
```

### 2. Inadequate Testing
Every function needs tests for: nil, empty, single, multiple, boundaries, errors. Check coverage with `go test -cover ./...`.

### 3. Inconsistent Naming
```go
// ❌ INCONSISTENT
First, FindLast, GetElement, LastOne

// ✅ CONSISTENT
First, Last, Find, Get, Min, Max
```

### 4. Missing Documentation
No doc comments = no way to know nil behavior, return semantics, or edge cases. Every exported function MUST have docs with example code.

### 5. Generic Constraints Misuse
```go
// ❌ WRONG - Sum needs Number constraint
func Sum[T any](slice []T) T { /* ... */ }

// ✅ CORRECT
func Sum[T lxconstraints.Number](slice []T) T { /* ... */ }
```

### 6. Not Testing Multiple Types
Generics need testing across type parameters:
```go
func TestFilter_Int(t *testing.T) { /* ... */ }
func TestFilter_String(t *testing.T) { /* ... */ }
func TestFilter_Struct(t *testing.T) { /* ... */ }
```

### 7. External Dependencies
❌ The project has **ZERO external dependencies** - only stdlib. Never add imports outside `import "..."` packages.

---

## 🔄 Cross-Package Patterns

### Import Pattern

```go
import (
    "github.com/hgapdvn/lx/lxslices"      // Import actual package directory names (no lx prefix)
    "github.com/hgapdvn/lx/lxconstraints" // Not /lxconstraints
    "github.com/hgapdvn/lx/lxtypes"       // Not /lxtypes
)

// Use as: slices.Filter(...)
// Package names have lx prefix when used: lxslices, lxmaps, lxtypes, etc.
```

### Version Compatibility: Go 1.18+ Baseline
- `slices` package uses build tags for Go 1.21+ vs older versions
  - `sort_go121.go` (uses stdlib `slices.SortFunc`)
  - `sort_legacy.go` (uses `sort.Slice`)
- Always test with `-race` detector for concurrent safety

### Map vs Slice Operations
- `lxmaps` mirrors `lxslices` patterns but for maps:
  - `Keys()`, `Values()`, `Entries()` (like `slices.Map`)
  - `Filter()` (like `slices.Filter`)
  - `Merge()`, `Intersect()`, `Union()`, `Difference()`

---

## 📖 Documentation & Examples

### Package-Level Doc (in `doc.go`)
Describe categories of functions with examples:
```go
// Package lxslices provides generic slice utilities.
//
// The package is organized into several key categories:
//
// 1. Transformation (transform.go)
//   - Map, FlatMap - Transform elements
//   - GroupBy, Chunk - Batch operations
//
// 2. Filtering (filter.go)
//   - Filter, Find - Query elements
//   - Any, All, None - Predicate checks
```

### Function Doc (above each function)
```go
// FirstOr returns the first element or a default value if the slice is empty or nil.
//
// Example:
//
//     value := lxslices.FirstOr([]int{2, 3}, 1)
//     // value: 2
//
func FirstOr[T any](slice []T, defaultValue T) T {
```

---

## ✅ Quick Checklist for New Code

- [ ] Function uses appropriate generic constraints
- [ ] Nil behavior documented and tested
- [ ] Table-driven tests for all cases (nil, empty, single, multiple)
- [ ] Tests cover different type parameters
- [ ] 100% test coverage on new code
- [ ] Function documentation with example
- [ ] No external dependencies
- [ ] Follows naming conventions (e.g., `First`, `FirstOr`, `MustFirst`)
- [ ] Consistent with similar functions in same/other packages
- [ ] All tests pass: `go test ./...`
- [ ] Format & lint: `gofmt -w .` and `go vet ./...`

---

## 📞 References

- **Full Guidelines**: `.github/copilot-instructions.md`
- **Contribution Guide**: `CONTRIBUTING.md`
- **Package Planning**: `PACKAGE_ROADMAP.md`
- **Module**: `github.com/hgapdvn/lx`
- **Go Version**: 1.18+
- **License**: Apache 2.0

---

**Last Updated**: April 4, 2026  
**For**: AI Coding Agents (Copilot, Cursor, Claude, etc.)
