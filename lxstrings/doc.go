// Package lxstrings provides utilities for working with strings.
//
// This package offers a collection of small, well-tested helpers that
// complement the Go standard library for common string operations. Functions
// are designed to be nil/empty-safe where relevant and to preserve existing
// semantics used across the lx project.
//
// Important notes about semantics in this package:
//
//   - Length returns the byte length of a string (the result of len(s)). Many
//     functions in this package operate on runes (for example SubString,
//     Abbreviate, Reverse) to provide Unicode-aware slicing. Keep this in mind
//     when mixing Length with rune-based operations.
//
//   - IsBlank currently treats the common ASCII whitespace characters
//     (' ', '\n', '\t', '\r') as blank. There is also an IsSpace function
//     which uses unicode.IsSpace to detect Unicode whitespace. Consider which
//     behavior you need when choosing between them.
//
//   - Historical aliases: You will find pairs of functions that are aliases
//     (for example StartBy / StartWith and EndBy / EndWith). Both sets are
//     provided for backward compatibility; we plan to standardize names in a
//     future minor release. For now prefer the "With" family (StartWith/
//     EndWith) for new code.
//
// Example:
//
//	package main
//
//	import (
//	    "fmt"
//	    "github.com/nthanhhai2909/lx/lxstrings"
//	)
//
//	func main() {
//	    s := "  Hello, 世界  "
//	    fmt.Println(lxstrings.TrimSpace(s)) // "Hello, 世界"
//	    fmt.Println(lxstrings.Length(s))    // byte length
//	    fmt.Println(lxstrings.IsBlank("\t\n")) // true
//	}
package lxstrings
