<div align="center">

# 🚀 lx

[![Go Version](https://img.shields.io/badge/Go-1.18%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/nthanhhai2909/lx)](https://goreportcard.com/report/github.com/nthanhhai2909/lx)
[![GoDoc](https://pkg.go.dev/badge/github.com/nthanhhai2909/lx)](https://pkg.go.dev/github.com/nthanhhai2909/lx)

**Small, focused extensions to Go's standard library**

*Lightweight • Idiomatic • Type-safe • Zero Dependencies*

[Features](#-features) • [Packages](#-packages) • [Installation](#-installation) • [Examples](#-quick-examples) • [Roadmap](#-roadmap)

</div>

---

## 💡 Why lx?

Working with Go's standard library is great, but sometimes you need just a *little* more. `lx` provides carefully crafted utilities that:

- ✨ **Enhance productivity** without adding bloat
- 🎯 **Stay focused** - small, single-purpose packages (no giant "utils")
- 🔒 **Type-safe** with Go 1.18+ generics
- 📦 **Zero external dependencies** - just pure Go
- 🧪 **Battle-tested** with comprehensive test coverage
- 📖 **Well-documented** with examples and clear APIs

Think of it as the **missing pieces** of Go's stdlib - nothing more, nothing less.

---

## ✨ Features

| Feature | Description |
|---------|-------------|
| 🎭 **Generic-first** | Built from the ground up with Go generics for type safety |
| 🪶 **Lightweight** | Each package is intentionally small and focused |
| 🔄 **Consistent APIs** | Predictable naming patterns across all packages |
| ⚡ **Zero overhead** | No reflection tricks, just pure Go performance |
| 🧪 **100% tested** | Comprehensive test coverage with edge cases |
| 📚 **Clear docs** | Every function has examples and clear documentation |
| 🔧 **Nil-safe** | Thoughtful handling of nil values and edge cases |

---

## 📦 Packages

### Core Packages

| Package | Description | Status | Examples |
|---------|-------------|--------|----------|
| [`lxslices`](slices) | Slice operations (filter, map, reduce, etc.) | ✅ Stable | [View](slices#examples) |
| [`lxstrings`](./lxstrings) | String utilities (blank checks, case conversion, etc.) | ✅ Stable | [View](./lxstrings#examples) |
| [`lxptrs`](ptrs) | Pointer helpers (ref, deref, safe operations) | ✅ Stable | [View](ptrs#examples) |
| [`lxtypes`](types) | Functional type definitions (Predicate, Function, Optional, Result, Either, Ref, Future, Lazy, Tuples, etc.) | ✅ Stable | [View](types#examples) |
| [`lxtuples`](./lxtuples) | Tuple types (Pair, Triple, Quad) | ✅ Stable | [View](./lxtuples#examples) |
| [`lxsystems`](./lxsystems) | System information (OS, paths, environment) | ✅ Stable | [View](./lxsystems#examples) |
| [`lxconstraints`](./lxconstraints) | Generic type constraints | ✅ Stable | [View](./lxconstraints#examples) |

### 🚧 Coming Soon

See our [**Package Roadmap**](./PACKAGE_ROADMAP.md) for planned packages:

- 🔥 **Phase 1**: `lxmaps`, `lxmath`, `lxerrors`, `lxio`
- 🟡 **Phase 2**: `lxtime`, `lxjson`, `lxhttp`, `lxregex`
- 🟢 **Phase 3**: `lxcontext`, `lxrand`, `lxvalidate`, `lxcrypto`
- And more based on community feedback!

> 💡 **Want to contribute?** Check out [open issues](../../issues) or suggest new packages!

---

## 🚀 Installation

```bash
go get github.com/nthanhhai2909/lx
```

**Requirements**: Go 1.18 or higher (generics support required)

> 💡 **Note**: Installing `lx` downloads all packages, but you only import what you need:
> ```go
> import "github.com/nthanhhai2909/lx/lxslices"  // Only use slices
> import "github.com/nthanhhai2909/lx/lxstrings" // Only use strings
> ```
> Go's compiler will only include imported packages in your binary.

---

## ⚡ Quick Examples

### Working with Slices

```go
import "github.com/nthanhhai2909/lx/lxslices"

// Filter even numbers
numbers := []int{1, 2, 3, 4, 5, 6}
evens := lxslices.Filter(numbers, func(n int) bool {
    return n%2 == 0
})
// evens: [2, 4, 6]

// Transform with Map
doubled := lxslices.Map(numbers, func(n int) int {
    return n * 2
})
// doubled: [2, 4, 6, 8, 10, 12]

// Aggregate with Reduce
sum := lxslices.Sum(numbers)
// sum: 21

// Find elements
first, found := lxslices.Find(numbers, func(n int) bool {
    return n > 3
})
// first: 4, found: true

// Chunk for pagination
chunks := lxslices.Chunk(numbers, 2)
// chunks: [[1, 2], [3, 4], [5, 6]]
```

### String Utilities

```go
import "github.com/nthanhhai2909/lx/lxstrings"

// Check for blank strings (empty or whitespace)
if lxstrings.IsBlank("   ") {
    // Handle blank input
}

// Safe string operations
lxstrings.Capitalize("hello world")  // "Hello world"
lxstrings.Abbreviate("Long text...", 10)  // "Long te..."
lxstrings.Reverse("golang")  // "gnalog"
lxstrings.ContainsIgnoreCase("Hello", "HELLO")  // true
```

### Pointer Helpers

```go
import "github.com/nthanhhai2909/lx/lxptrs"

// Create pointers inline
value := lxptrs.Ref(42)  // *int pointing to 42

// Safe dereference
str := lxptrs.Deref(stringPtr)  // panics if nil

// Or use with default
str := lxptrs.DerefOr(stringPtr, "default")  // "default" if nil
```

### System Information

```go
import "github.com/nthanhhai2909/lx/lxsystems"

// Get system info
lxsystems.OSName    // "darwin", "linux", "windows"
lxsystems.NumCPU()  // 8
lxsystems.IsLinux() // true/false

// Safe path operations
home := lxsystems.UserHomeDirOr("/tmp")
temp := lxsystems.TempDir()
```

---

## 📚 Documentation

Each package has comprehensive documentation with examples:

- **GoDoc**: [pkg.go.dev/github.com/nthanhhai2909/lx](https://pkg.go.dev/github.com/nthanhhai2909/lx)
- **Package READMEs**: Check each package directory for detailed docs
- **Examples**: Every function includes usage examples in tests

---

## 🗺️ Roadmap

We have an ambitious but focused roadmap! See [**PACKAGE_ROADMAP.md**](./PACKAGE_ROADMAP.md) for:

- 📋 15 planned packages across 4 tiers
- 🎯 Implementation priority matrix
- 📅 Timeline and phases
- 💡 Package design principles
- 🤝 How to contribute

**Next up**: `lxmaps` for comprehensive map operations! 🎉

---

## 🤝 Contributing

We ❤️ contributions! Here's how you can help:

- 🐛 **Report bugs** via [GitHub Issues](../../issues)
- 💡 **Suggest features** or new packages
- 📝 **Improve documentation**
- 🔧 **Submit pull requests**

See [**CONTRIBUTING.md**](./CONTRIBUTING.md) for detailed guidelines.

### Good First Issues

Looking to contribute? Check out issues labeled [`good first issue`](../../issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22) - perfect for newcomers!

---

## 🏆 Design Philosophy

`lx` follows these principles:

| Principle | What it means |
|-----------|---------------|
| **🎯 Focused** | Each package solves one problem well |
| **📖 Readable** | Code should be self-documenting |
| **🔒 Safe** | Thoughtful nil handling and error messages |
| **⚡ Fast** | No unnecessary allocations or reflection |
| **🧪 Tested** | Comprehensive test coverage (>80%) |
| **🔄 Consistent** | Predictable APIs across all packages |
| **📦 Minimal** | Zero external dependencies |

---

## 📊 Project Status

| Metric | Status |
|--------|--------|
| **Stability** | Beta - API may change |
| **Test Coverage** | >85% across all packages |
| **Go Version** | 1.18+ required (generics) |
| **Dependencies** | Zero (stdlib only) |
| **License** | Apache 2.0 |

---

## 💬 Community & Support

- 💬 **Discussions**: [GitHub Discussions](../../discussions)
- 🐛 **Issues**: [GitHub Issues](../../issues)
- ⭐ **Star us** if you find this useful!
- 📣 **Share** with your Go developer friends

---

## 🙏 Acknowledgments

Inspired by excellent libraries like:
- [samber/lo](https://github.com/samber/lo) - Lodash-style utilities
- [golang.org/x/exp](https://pkg.go.dev/golang.org/x/exp) - Experimental packages
- [spf13/cast](https://github.com/spf13/cast) - Type conversion

`lx` builds on these ideas with a focus on minimalism, zero dependencies, and Go-idiomatic APIs.

---

## 📄 License

Apache License 2.0 - see [LICENSE](LICENSE) for details.

---

<div align="center">

**Made with ❤️ for the Go community**

[⬆ Back to Top](#-lx)

</div>

