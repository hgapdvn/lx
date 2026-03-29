# 🤝 Contributing to lx

Thanks for your interest in contributing to **lx**! 🎉

We're building a comprehensive, high-quality Go utilities library, and we'd love your help. Whether you're fixing a bug, adding a feature, improving documentation, or suggesting ideas, every contribution matters.

This guide will help you get started quickly and ensure a smooth contribution process.

---

## 📋 Table of Contents

- [Code of Conduct](#-code-of-conduct)
- [Ways to Contribute](#-ways-to-contribute)
- [Getting Started](#-getting-started)
- [Development Workflow](#-development-workflow)
- [Coding Standards](#-coding-standards)
- [Testing Requirements](#-testing-requirements)
- [Documentation Guidelines](#-documentation-guidelines)
- [Pull Request Process](#-pull-request-process)
- [Package Design Guidelines](#-package-design-guidelines)

---

## 📜 Code of Conduct

This project adheres to a code of conduct that we expect all contributors to follow. Please be respectful, inclusive, and considerate in all interactions.

**TL;DR**: Be kind, be professional, be constructive.

---

## 💡 Ways to Contribute

There are many ways you can contribute to lx:

### 🐛 Report Bugs
Found a bug? [Open an issue](../../issues/new?labels=bug) with:
- Clear description of the problem
- Steps to reproduce
- Expected vs actual behavior
- Go version and OS

### ✨ Suggest Features
Have an idea? [Open an issue](../../issues/new?labels=enhancement) or [start a discussion](../../discussions) with:
- Use case description
- Proposed API design
- Why it adds value beyond stdlib

### 📝 Improve Documentation
- Fix typos or clarify existing docs
- Add examples to package documentation
- Improve README or guides
- Write blog posts or tutorials

### 🔧 Submit Code
- Fix bugs
- Implement new features
- Add missing tests
- Improve performance

### 🎯 Review Pull Requests
- Test proposed changes
- Provide constructive feedback
- Suggest improvements

---

## 🚀 Getting Started

### Prerequisites

- **Go 1.18+** installed (generics support required)
- **Git** for version control
- A **GitHub account**
- Familiarity with Go generics

### Setup

1. **Fork** the repository on GitHub

2. **Clone** your fork locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/lx.git
   cd lx
   ```

3. **Add upstream** remote:
   ```bash
   git remote add upstream https://github.com/hgapdvn/lx.git
   ```

4. **Verify** everything works:
   ```bash
   go test ./...
   ```

You're ready to contribute! 🎉

---

## 🔄 Development Workflow

### 1. Create a Branch

Always create a new branch from `main`:

```bash
git checkout main
git pull upstream main
git checkout -b feat/your-feature-name
```

#### Branch Naming Convention

Use descriptive names with these prefixes:

| Prefix | Purpose | Example |
|--------|---------|---------|
| `feat/` | New features | `feat/lxmaps-keys-values` |
| `fix/` | Bug fixes | `fix/lxslices-nil-panic` |
| `docs/` | Documentation only | `docs/improve-readme` |
| `test/` | Adding tests | `test/lxstrings-coverage` |
| `refactor/` | Code refactoring | `refactor/lxslices-performance` |
| `chore/` | Tooling, CI, misc | `chore/update-ci` |

### 2. Make Your Changes

- Write clean, idiomatic Go code
- Follow existing code style
- Add comprehensive tests
- Update documentation

### 3. Test Your Changes

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detector
go test -race ./...

# Test specific package
go test ./lxslices/...
```

### 4. Commit Your Changes

We follow **Conventional Commits** for clear, semantic commit messages:

```bash
git add .
git commit -m "feat(lxmaps): add Keys and Values functions"
```

#### Commit Message Format

```
<type>(<scope>): <short description>

[optional body]

[optional footer]
```

**Types**:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation only
- `test`: Adding tests
- `refactor`: Code refactoring
- `perf`: Performance improvement
- `chore`: Tooling, CI, etc.

**Examples**:
```
feat(lxslices): add Chunk function for pagination
fix(lxstrings): handle nil pointer in IsBlank
docs: update installation instructions
test(lxmaps): add edge cases for Merge function
refactor(lxptrs): simplify DerefOr implementation
```

### 5. Push and Create PR

```bash
git push origin feat/your-feature-name
```

Then open a Pull Request on GitHub!

---

## 📏 Coding Standards

### Go Style Guidelines

Follow standard Go conventions:

- Use `gofmt` for formatting
- Run `go vet` to catch common mistakes
- Follow [Effective Go](https://go.dev/doc/effective_go)
- Use meaningful variable names
- Keep functions small and focused

### Package-Specific Guidelines

#### Function Naming

```go
// ✅ Good - clear, descriptive
func ContainsAny[T comparable](slice []T, elems ...T) bool

// ❌ Bad - unclear, abbreviated
func ChkAny[T comparable](s []T, e ...T) bool
```

#### Return Values

Prefer `(value, bool)` for operations that might fail on valid input:

```go
// ✅ Good - signals whether value was found
func First[T any](slice []T) (T, bool)

// ❌ Bad - can't distinguish between zero value and not found
func First[T any](slice []T) T
```

Provide `*Or` variants for convenience:

```go
func First[T any](slice []T) (T, bool)
func FirstOr[T any](slice []T, defaultValue T) T
```

#### Error Handling

```go
// ✅ Good - explicit error handling
func ReadJSON[T any](path string) (T, error)

// Consider Must* variants for convenience
func MustReadJSON[T any](path string) T
```

#### Nil Handling

Always document and handle nil slices/maps gracefully:

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

---

## 🧪 Testing Requirements

**Testing is mandatory for all contributions.**

### Test Coverage

- **Minimum**: 95% coverage for new code
- **Target**: 100% coverage
- Use `go test -cover ./...` to check

### Test Structure

Use table-driven tests:

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
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("Filter() = %v, want %v", result, tt.expected)
            }
        })
    }
}
```

### What to Test

- ✅ Happy path (normal usage)
- ✅ Edge cases (empty, nil, zero values)
- ✅ Boundary conditions (min/max values)
- ✅ Error conditions
- ✅ Different type parameters (for generics)

### Test Best Practices

- Tests should be **deterministic** (no flaky tests)
- Tests should be **fast** (no unnecessary sleeps)
- Tests should be **isolated** (no shared state)
- Use **descriptive test names**
- Add **comments** for complex test cases

---

## 📖 Documentation Guidelines

### Package Documentation

Every package should have a `doc.go` or package comment:

```go
// Package lxmaps provides utilities for working with Go maps.
//
// This package offers functions for common map operations like
// filtering, transforming, and querying maps in a type-safe way.
//
// Example:
//
//     m := map[string]int{"a": 1, "b": 2, "c": 3}
//     keys := lxmaps.Keys(m)  // []string{"a", "b", "c"}
//
package lxmaps
```

### Function Documentation

Every exported function must have:

1. **Brief description** (one line)
2. **Detailed explanation** (if needed)
3. **Parameter descriptions** (if not obvious)
4. **Return value description**
5. **Examples** (in test files)

```go
// Filter returns a new map containing only key-value pairs that satisfy the predicate.
// The order of iteration is not guaranteed.
// Returns nil if the input map is nil.
//
// Example:
//
//     m := map[string]int{"a": 1, "b": 2, "c": 3}
//     even := lxmaps.Filter(m, func(k string, v int) bool {
//         return v%2 == 0
//     })
//     // even: map[string]int{"b": 2}
//
func Filter[K comparable, V any](m map[K]V, predicate func(K, V) bool) map[K]V
```

### README Updates

When adding a new package:
1. Update main README.md
2. Create package-level README.md with examples
3. Update PACKAGE_ROADMAP.md if applicable

---

## 🔍 Pull Request Process

### Before Submitting

- ✅ All tests pass: `go test ./...`
- ✅ Code is formatted: `gofmt -w .`
- ✅ No vet issues: `go vet ./...`
- ✅ Documentation is updated
- ✅ Tests have good coverage
- ✅ Commits follow conventional format

### PR Description Template

```markdown
## Description
Brief description of what this PR does.

## Type of Change
- [ ] Bug fix (non-breaking change fixing an issue)
- [ ] New feature (non-breaking change adding functionality)
- [ ] Breaking change (fix or feature causing existing functionality to change)
- [ ] Documentation update

## Changes Made
- Added X function to lxmaps package
- Fixed nil pointer panic in Y
- Updated documentation for Z

## Testing
- [ ] Added unit tests
- [ ] All tests pass locally
- [ ] Test coverage is >80%

## Checklist
- [ ] Code follows project style guidelines
- [ ] Self-review completed
- [ ] Documentation updated
- [ ] No breaking changes (or documented if unavoidable)

## Related Issues
Closes #123
```

### Review Process

1. **Automated checks** run (tests, linting)
2. **Maintainer review** (usually within 2-3 days)
3. **Feedback incorporated** (if any)
4. **Final approval** and merge

---

## 🎨 Package Design Guidelines

When proposing a new package, ensure it:

### ✅ Should Include

- **Focused scope** - solves one problem well
- **Clear value** - provides more than stdlib
- **Type-safe** - uses generics where appropriate
- **Zero dependencies** - uses only stdlib
- **Well-tested** - comprehensive test coverage
- **Documented** - clear examples and docs

### ❌ Should Avoid

- **Giant "utils"** packages
- **Framework-specific** code
- **External dependencies** (unless absolutely necessary)
- **Trivial wrappers** with no added value
- **Breaking changes** to existing APIs

### Package Structure

**For small packages:**

```
lxpackagename/
├── packagename.go        # Main implementation
├── packagename_test.go   # Tests
├── doc.go               # Package documentation (optional)
├── README.md            # Package-specific docs (optional)
└── examples_test.go     # Example functions (optional)
```

**For large packages** (like lxslices):

```
lxpackagename/
├── feature1.go          # Feature group 1 (e.g., aggregation.go)
├── feature1_test.go     # Tests for feature 1
├── feature2.go          # Feature group 2 (e.g., filter.go)
├── feature2_test.go     # Tests for feature 2
├── errs.go              # Package-specific errors (if needed)
├── doc.go               # Package documentation (optional)
├── README.md            # Package-specific docs (optional)
├── ROADMAP.md           # Future enhancements (optional)
└── examples_test.go     # Example functions (optional)
```

**Guidelines:**
- Split large packages into logical feature groups (one file per feature set)
- Keep test files alongside their implementation files
- Use clear, descriptive filenames that indicate functionality
- Add `doc.go` for comprehensive package documentation
- Include `README.md` for usage examples and quick reference

---

## 🙋 Getting Help

- 💬 **Discussions**: [GitHub Discussions](../../discussions)
- 🐛 **Issues**: [GitHub Issues](../../issues)
- 📧 **Email**: For private matters only

---

## 🙏 Thank You!

Every contribution, no matter how small, helps make `lx` better for everyone. We appreciate your time and effort!

**Happy coding!** 🚀

---

<div align="center">

[⬆ Back to Top](#-contributing-to-lx)

</div>

