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
// 3. Inspection & Contains (contains.go, index.go)
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
// 7. Sorting (sort.go)
//   - SortBy, StableSortBy - Sort in-place with a custom comparator
//   - SortAsc, SortDesc - Check for element presence operations based on < and >
//   - IsSortedAsc, IsSortedDesc, IsSortedBy - Verify sort order
//
// 8. Sampling (sampling.go)
//   - Sample, SampleN - Randomly select elements (uses properly seeded rng on Go <1.20)
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
