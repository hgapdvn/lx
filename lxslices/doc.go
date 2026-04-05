// Package lxslices provides a comprehensive set of generic utility functions
// for operating on Go slices. It aims to offer functional programming
// paradigms and common slice manipulations in a type-safe manner using Go 1.18+ generics.
//
// The package is organized into several key categories:
//
// 1. Transformation (transform.go)
//   - Map, FlatMap - Transform slice elements
//   - ForEach, ForEachIndexed - Iterate over elements for side effects
//   - GroupBy - Group elements by a key
//   - AssociateBy - Create a strict map from elements using a key builder
//   - Chunk - Split a slice into smaller appropriately sized batches
//   - Window, WindowFunc - Sliding window operations
//   - Reverse - Reverses a slice in-place
//   - Concat - Joins multiple slices
//   - Zip, Unzip - Combine and split paired slices
//   - Copy, Clone - Shallow copy slices
//
// 2. Filtering (filter.go)
//   - Filter - Keep elements matching a predicate
//   - Partition - Split elements by a predicate
//   - Find, FindLast - Locate specific elements
//   - Any, All, None - Check predicates across elements
//   - Count - Count elements matching a predicate
//
// 3. Inspection & Contains (contains_key.go, index.go)
//   - Contains, ContainsAny, ContainsAll, ContainsFunc - Check for element presence
//   - Index, LastIndex, IndexFunc - Find element indices
//   - First, Last, Get - Safely retrieve elements
//   - MinIndex, MaxIndex - Find extrema indices
//
// 4. Manipulation (manipulation.go)
//   - Append, Prepend, Insert - Add elements
//   - Remove, RemoveAt, RemoveAll, RemoveFunc - Remove elements
//   - Replace, ReplaceAt - Swap elements
//   - RotateLeft, RotateRight - Shift elements circularly in-place
//
// 5. Aggregation (aggregation.go)
//   - Reduce - Accumulate values
//   - Sum, Average - Math operations on numeric slices
//   - Min, Max - Find extrema values
//
// 6. Set Operations (set.go)
//   - Unique - Remove duplicates
//   - Intersection, Union, Difference - Set arithmetic on slices
//
// 7. Sorting (sort.go, sort_go121.go, sort_legacy.go)
//   - SortBy, StableSortBy - Sort in-place with a custom comparator
//     (uses slices.SortFunc on Go 1.21+, sort.Slice on older versions)
//   - SortAsc, SortDesc - Sort in ascending or descending order
//   - IsSortedAsc, IsSortedDesc, IsSortedBy - Verify sort order
//
// 8. Sampling (sampling.go)
//   - Sample - Randomly select a single element; returns (T, bool)
//   - SampleN - Randomly select n distinct elements
//
// 9. Take/Drop (take_drop.go)
//   - Take, TakeLast, TakeWhile - Select elements from the start or while a condition holds
//   - Drop, DropLast, DropWhile - Skip elements from the start or while a condition holds
//
// 10. Comparison (comparison.go)
//   - Equal, EqualFunc - Compare two slices for equality
//   - StartsWith, EndsWith - Check prefix or suffix subsequences
//
// 11. Generation (generation.go)
//   - Repeat, RepeatSlice - Produce slices by repeating a value or slice
//   - Range, RangeStep - Produce numeric ranges
//
// # Nil vs Empty Slice Semantics
//
// This package follows a consistent policy: a nil input slice produces a nil
// output slice; a non-nil but empty input slice produces a non-nil empty output
// slice (unless the function explicitly documents otherwise). Functions that
// mutate the slice in-place (Reverse, RotateLeft, RotateRight, SortBy, etc.)
// are excluded from this policy because they do not return a new slice.
//
// Note on Mutation:
// Most functions in this package (like Map, Filter, Remove, Append) return
// entirely new slices and do not modify the input slice.
// Functions that explicitly mutate the slice in-place include:
//   - Reverse
//   - RotateLeft, RotateRight
//   - SortBy, StableSortBy, SortAsc, SortDesc
//
// For usage examples see the accompanying *_test.go files.
package lxslices
