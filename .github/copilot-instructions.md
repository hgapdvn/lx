# AI Instructions for lx Project

> Comprehensive guide for AI assistants working with the lx Go utilities library

**Last Updated**: February 28, 2026

---

## 📋 Project Overview

**Project Name**: lx  
**Repository**: github.com/hgapdvn/lx  
**Language**: Go 1.18+  
**License**: Apache 2.0
**Purpose**: Small, focused extensions to Go's standard library

### Core Philosophy

1. **Small & Focused**: Each package solves one problem well (no giant "utils" packages)
2. **Generic-First**: Built with Go 1.18+ generics for type safety
3. **Retro-Compatibility**: Supports Go 1.18+. Always favor `lxslices` and `lxmaps` over the newer `slices` and `maps` standard library packages (introduced in Go 1.21) to ensure backwards compatibility.
4. **Zero Dependencies**: Uses only Go standard library
5. **Idiomatic Go**: Follows Go conventions and best practices
6. **Well-Tested**: Comprehensive test coverage (>80% target, 90%+ ideal)
7. **Nil-Safe**: Thoughtful handling of nil values and edge cases

---

## 🏗️ Project Structure

```
lx/
├── lxconstraints/    # Generic type constraints (Integer, Number, Ordered, etc.)
├── lxptrs/           # Pointer utilities (Ref, Deref, etc.)
├── lxslices/         # Slice operations (Filter, Map, Reduce, etc.)
├── lxstrings/        # String utilities (IsBlank, Capitalize, etc.)
├── lxsystems/        # System information (OS detection, paths, etc.)
├── lxtuples/         # Tuple types (Pair, Triple, etc.)
├── lxtypes/          # Type utilities (Lazy, etc.)
└── (future packages) # See PACKAGE_ROADMAP.md
```

### Key Files
- `README.md` - Main project documentation
- `CONTRIBUTING.md` - Contribution guidelines (502 lines, comprehensive)
- `CODE_OF_CONDUCT.md` - Community standards
- `SECURITY.md` - Security vulnerability reporting
- `PACKAGE_ROADMAP.md` - Strategic planning (15 proposed packages)
- `LAUNCH_CHECKLIST.md` - Deployment guide

---

## 🎯 Coding Standards & Conventions

### 1. Function Naming

**✅ DO**:
```go
func ContainsAny[T comparable](slice []T, elems ...T) bool
func GetOrDefault[T any](value *T, defaultValue T) T
func IsBlank(str string) bool
```

**❌ DON'T**:
```go
func ChkAny[T comparable](s []T, e ...T) bool  // Too abbreviated
func get[T any](v *T, def T) T                 // Not exported
func checkIfBlank(str string) bool             // Too verbose
```

### 2. Return Value Patterns

**For operations that might fail on valid input**:
```go
// Returns (value, found)
func First[T any](slice []T) (T, bool)
func Min[T Ordered](slice []T) (T, bool)
func Get[T any](slice []T, index int) (T, bool)
```

**Provide convenience variants with defaults**:
```go
func First[T any](slice []T) (T, bool)
func FirstOr[T any](slice []T, defaultValue T) T
```

**For operations that should never fail**:
```go
func Filter[T any](slice []T, predicate func(T) bool) []T
func Map[T, U any](slice []T, fn func(T) U) []U
```

### 3. Nil Handling

**Always document nil behavior**:
```go
// Filter returns a new slice containing only elements that satisfy predicate.
// Returns nil if input is nil, empty slice if no elements match.
func Filter[T any](slice []T, predicate func(T) bool) []T {
    if slice == nil {
        return nil
    }
    // ...
}
```

**Preserve nil vs empty semantics**:
```go
// ✅ Good - preserves nil
if slice == nil {
    return nil
}

// ❌ Bad - loses nil information
return []T{}
```

### 4. Error Handling

**Use errors for operations that can fail**:
```go
func UniqueGroupBy[T any, K comparable](slice []T, fn func(T) K) (map[K]T, error)
```

**Define package-specific errors**:
```go
// In errs.go
var (
    ErrDuplicateKey = errors.New("lxslices: duplicate key")
)
```

**Consider Must* variants for convenience**:
```go
func Parse(s string) (T, error)
func MustParse(s string) T  // Panics on error
```

### 5. Generic Type Parameters

**Use descriptive single letters**:
```go
T - generic type
U - second generic type
K - key type (maps)
V - value type (maps)
```

**Use constraints appropriately**:
```go
func Sum[T lxconstraints.Number](slice []T) T
func Sort[T lxconstraints.Ordered](slice []T) []T
func Contains[T comparable](slice []T, elem T) bool
```

---

## 🧪 Testing Standards

### Test Structure

**Use table-driven tests**:
```go
func TestFilter(t *testing.T) {
    tests := []struct {
        name      string
        input     []int
        predicate func(int) bool
        expected  []int
    }{
        {
            name:      "filter even numbers",
            input:     []int{1, 2, 3, 4, 5},
            predicate: func(n int) bool { return n%2 == 0 },
            expected:  []int{2, 4},
        },
        {
            name:      "empty slice",
            input:     []int{},
            predicate: func(n int) bool { return true },
            expected:  []int{},
        },
        {
            name:      "nil slice",
            input:     nil,
            predicate: func(n int) bool { return true },
            expected:  nil,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Filter(tt.input, tt.predicate)
            // Assert results...
        })
    }
}
```

### Test Coverage Requirements

- **Minimum**: 100% coverage for new code
- **Target**: 100% coverage
- **Check with**: `go test -cover ./...`
- **View HTML**: `go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out`

### What to Test

✅ **Must Test**:
- Happy path (normal usage)
- Edge cases (empty, nil, zero values)
- Boundary conditions (first, last, min, max)
- Error conditions
- Different type parameters (for generics)

✅ **Test File Organization**:
```go
// Test different types
TestFilter_Int(t *testing.T)
TestFilter_String(t *testing.T)
TestFilter_Struct(t *testing.T)

// Test edge cases
TestFilter_EmptySlice(t *testing.T)
TestFilter_NilSlice(t *testing.T)
TestFilter_SingleElement(t *testing.T)
```

---

## 📖 Documentation Standards

### Package Documentation

**Every package needs**:
```go
// Package lxslices provides utilities for working with Go slices.
//
// This package offers functions for common slice operations like
// filtering, mapping, reducing, and querying slices in a type-safe way
// using Go generics.
//
// Example:
//
//     numbers := []int{1, 2, 3, 4, 5}
//     evens := lxslices.Filter(numbers, func(n int) bool {
//         return n%2 == 0
//     })
//     // evens: [2, 4]
//
package lxslices
```

### Function Documentation

**Template**:
```go
// FunctionName does X with Y and returns Z.
// Additional details about behavior.
// Returns (value, found) semantics or error handling details.
//
// Example:
//
//     result := FunctionName(input)
//     // result: expected output
//
func FunctionName[T any](param T) T {
    // ...
}
```

**Real Example**:
```go
// Filter returns a new slice containing only the elements of the original slice that satisfy the predicate.
// The order of the elements in the returned slice is the same as in the original slice.
// Returns nil if input is nil, empty slice if no elements match.
//
// Example:
//
//     numbers := []int{1, 2, 3, 4, 5}
//     evens := lxslices.Filter(numbers, func(n int) bool {
//         return n%2 == 0
//     })
//     // evens: [2, 4]
//
func Filter[T any](slice []T, predicate func(T) bool) []T
```

---

## 🔧 Development Workflow

### 1. Branch Naming

Use prefixes:
- `feat/` - New features (`feat/lxmaps-keys-values`)
- `fix/` - Bug fixes (`fix/lxslices-nil-panic`)
- `docs/` - Documentation (`docs/improve-readme`)
- `test/` - Adding tests (`test/lxstrings-coverage`)
- `refactor/` - Code refactoring (`refactor/lxslices-performance`)
- `chore/` - Tooling, CI (`chore/update-ci`)

### 2. Commit Messages

**Format**: `<type>(<scope>): <description>`

**Types**: `feat`, `fix`, `docs`, `test`, `refactor`, `perf`, `chore`

**Examples**:
```
feat(lxslices): add Chunk function for pagination
fix(lxstrings): handle nil pointer in IsBlank
docs: update installation instructions
test(lxmaps): add edge cases for Merge function
refactor(lxptrs): simplify DerefOr implementation
```

### 3. File Organization

**For new packages**:
```
lxpackagename/
├── packagename.go          # Main functions
├── packagename_test.go     # Tests
├── types.go               # Types/structs (if needed)
├── constants.go           # Constants (if needed)
├── errors.go              # Package-specific errors
├── doc.go                 # Package documentation
└── README.md              # Package-specific docs
```

**For large packages** (like lxslices):
```
lxslices/
├── aggregation.go          # Reduce, Sum, Min, Max, Average
├── aggregation_test.go     # Tests for aggregation
├── filter.go               # Find, Filter, Any, All, None
├── filter_test.go          # Tests for filter
├── manipulation.go         # Append, Insert, Remove, etc.
├── manipulation_test.go    # Tests for manipulation
└── ...
```

### 4. Adding New Functions

**Steps**:
1. Check if function fits existing file or needs new file
2. Write function signature and documentation
3. Write comprehensive tests (80%+ coverage)
4. Implement function
5. Run tests: `go test ./packagename/`
6. Check coverage: `go test -cover ./packagename/`
7. Format code: `gofmt -w .`
8. Lint: `go vet ./...`

---

## 🎨 Package Design Guidelines

### When Creating New Packages

**✅ Should Include**:
- Focused scope (solves one problem well)
- Clear value beyond stdlib
- Type-safe with generics where appropriate
- Zero dependencies (stdlib only)
- Comprehensive tests (>80% coverage)
- Clear documentation with examples

**❌ Should Avoid**:
- Giant "utils" packages
- Framework-specific code
- External dependencies (unless essential)
- Trivial wrappers with no value
- Breaking changes to existing APIs

### API Design Patterns

**Consistency across packages**:
```go
// Contains pattern - check existence
Contains(collection, element) bool
ContainsAny(collection, elements...) bool
ContainsAll(collection, elements...) bool
ContainsFunc(collection, predicate) bool

// Get pattern - retrieve with default
Get(collection, key) (value, bool)
GetOr(collection, key, default) value
MustGet(collection, key) value

// OrDefault pattern - fallback values
FirstOr(slice, default) value
LastOr(slice, default) value
DerefOr(ptr, default) value
```

---

## 🔍 Common Patterns & Examples

### Pattern 1: Safe Operations with (value, bool)

```go
// Find first matching element
if value, found := lxslices.First(slice); found {
    fmt.Println(value)
}

// Or use default
value := lxslices.FirstOr(slice, defaultValue)
```

### Pattern 2: Functional Transformations

```go
// Chain operations
result := lxslices.Filter(
    lxslices.Map(input, transform),
    predicate,
)

// Or step by step
transformed := lxslices.Map(input, transform)
filtered := lxslices.Filter(transformed, predicate)
```

### Pattern 3: Generic Type Constraints

```go
// Use appropriate constraints
func Sum[T lxconstraints.Number](slice []T) T
func Max[T lxconstraints.Ordered](slice []T) (T, bool)
func Unique[T comparable](slice []T) []T
func Map[T, U any](slice []T, fn func(T) U) []U
```

### Pattern 4: Nil Safety

```go
// Always check and document nil behavior
func Process[T any](slice []T) []T {
    if slice == nil {
        return nil  // Preserve nil
    }
    // Process...
}
```

---

## ⚠️ Common Pitfalls to Avoid

### 1. Breaking Nil Semantics

```go
// ❌ Bad - loses nil information
func Filter[T any](slice []T, predicate func(T) bool) []T {
    result := []T{}  // Always returns non-nil
    // ...
}

// ✅ Good - preserves nil
func Filter[T any](slice []T, predicate func(T) bool) []T {
    if slice == nil {
        return nil
    }
    var result []T  // Will be nil if nothing added
    // ...
}
```

### 2. Not Testing Edge Cases

```go
// ✅ Always test:
- nil input
- empty input
- single element
- multiple elements
- boundary values (min/max)
- error conditions
```

### 3. Inconsistent Naming

```go
// ❌ Bad - inconsistent
func GetFirstElement[T any](slice []T) (T, bool)
func FindLastElement[T any](slice []T) (T, bool)

// ✅ Good - consistent
func First[T any](slice []T) (T, bool)
func Last[T any](slice []T) (T, bool)
```

### 4. Missing Documentation

```go
// ❌ Bad - no docs
func Filter[T any](slice []T, predicate func(T) bool) []T {
    // ...
}

// ✅ Good - comprehensive docs
// Filter returns a new slice containing only elements that satisfy the predicate.
// Returns nil if input is nil, empty slice if no elements match.
func Filter[T any](slice []T, predicate func(T) bool) []T {
    // ...
}
```

---

## 🤖 AI Assistant Guidelines

### When Implementing New Features

1. **Check existing patterns** - Look at similar functions in the codebase
2. **Follow naming conventions** - Use consistent names across packages
3. **Write tests first** - Or at least alongside implementation
4. **Document thoroughly** - Every exported function needs docs
5. **Handle nil carefully** - Document and test nil behavior
6. **Use appropriate constraints** - Number, Ordered, comparable, any
7. **Consider Or variants** - For functions returning (value, bool)
8. **Check coverage** - Aim for 90%+ coverage

### When Investigating CI Pipeline Failures

1. **Verify cross-platform compatibility** - Tests that manipulate or read OS environment variables (like `os.Getenv`) or file paths may behave differently or uniquely conflict with system variables (e.g., the built-in `USERNAME` variable on Windows vs Linux). Capture snapshots of standard OS state before assertions.
2. **Check line endings & formatting** - Differences in CRLF (`\r\n`) vs LF (`\n`) checked out by Git on Windows runners will cause strict formatters like `gofmt` to instantly fail. Ensure line-endings are normalized (e.g. via `.gitattributes`).
3. **Reproduce strictly** - Always attempt to reproduce linting errors locally (e.g., running `golangci-lint run ./...`) and consider building/testing natively for the target OS (e.g., `GOOS=windows go build ./...`) to catch OS-specific flaws.
4. **Isolate environment modifications** - In tests, heavily prefer using `t.Setenv` over `os.Setenv` to prevent cross-contamination and race conditions between concurrent tests.

### When Reviewing Code

1. **API consistency** - Does it match existing patterns?
2. **Nil safety** - Is nil behavior documented and tested?
3. **Test coverage** - Are edge cases covered?
4. **Documentation** - Is it clear and includes examples?
5. **Generic constraints** - Are they appropriate?
6. **Performance** - Are there unnecessary allocations?
7. **Idiomatic Go** - Does it follow Go conventions?

### When Adding Tests

1. **Use table-driven tests** - Consistent with existing tests
2. **Test multiple types** - Especially for generic functions
3. **Cover edge cases** - nil, empty, single, multiple
4. **Descriptive names** - `TestFilter_EmptySlice` not `TestFilter1`
5. **Check coverage** - Run `go test -cover ./packagename/`

### When Writing Documentation

1. **One-line summary** - What does it do?
2. **Detailed behavior** - How does it handle edge cases?
3. **Return value semantics** - What does it return?
4. **Example code** - Show usage
5. **Related functions** - Link to similar functions

---

## 📞 Getting Help

- **Documentation**: Start with README.md and CONTRIBUTING.md
- **Examples**: Look at existing functions in lxslices
- **Testing**: See existing test files for patterns
- **Questions**: Open a GitHub Discussion
- **Bugs**: Open a GitHub Issue with reproduction steps

---

## ✅ Quick Reference

### Before Committing

```bash
# Run all tests
go test ./...

# Check coverage
go test -cover ./...

# Format code
gofmt -w .

# Lint code
go vet ./...

# Run with race detector
go test -race ./...
```

### Common Commands

```bash
# Test specific package
go test ./lxslices/

# Test with verbose output
go test -v ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Benchmark
go test -bench=. ./packagename/
```

---

## 🎯 Summary for AI Assistants

When working with the lx project:

1. **Philosophy**: Small, focused, idiomatic, type-safe, zero-dependency
2. **Quality**: 100% test coverage, comprehensive edge case testing
3. **Consistency**: Follow existing patterns in lxslices and other packages
4. **Documentation**: Every function needs clear docs with examples
5. **Nil Safety**: Always document and test nil behavior
6. **Generics**: Use appropriate constraints (Number, Ordered, comparable)
7. **Returns**: Use (value, bool) for potentially failing operations
8. **Testing**: Table-driven tests with descriptive names
9. **Naming**: Consistent with Go conventions and existing packages
10. **Roadmap**: Check PACKAGE_ROADMAP.md for planned features

**Most Important**: This is a high-quality, production-ready library. Every contribution should maintain these standards!

---

<div align="center">

**Last Updated**: February 28, 2026  
**Version**: 0.1.0 (Beta)  
**Maintainer**: github.com/nthanhhai2909

</div>