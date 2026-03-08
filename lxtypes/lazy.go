package lxtypes

import (
	"sync"
	"sync/atomic"
)

// Package lxtypes provides common type utilities and abstractions.
//
// This package offers generic type constructs like lazy evaluation,
// optional values, and other type-safe wrappers.

// Lazy represents a value that may be computed immediately or on first access.
// It provides a unified interface for both eager (immediate) and deferred (lazy) values.
//
// Lazy is safe for concurrent use when using LazyDeferred, as it ensures
// the computation function is called at most once via sync.Once.
//
// Example (Deferred):
//
//	expensive := lxtypes.LazyDeferred(func() (int, error) {
//	    // Expensive computation
//	    time.Sleep(time.Second)
//	    return 42, nil
//	})
//	value, _ := expensive.Get() // Computed here
//	value2, _ := expensive.Get() // Returns cached value
//
// Example (Eager):
//
//	immediate := lxtypes.LazyEager(42)
//	value, _ := immediate.Get() // Returns immediately
type Lazy[T any] interface {
	// Get returns the value, computing it if necessary.
	// For eager values, returns immediately. For deferred values,
	// computes on first call and caches the result.
	Get() (T, error)

	// MustGet returns the value, computing it if necessary.
	// Panics if Get() returns an error.
	MustGet() T

	// IsEvaluated returns true if the value has been computed or was provided eagerly.
	// Always returns true for eager values.
	IsEvaluated() bool
}

// LazyEager creates a Lazy that wraps an already-computed value.
// No computation is performed - Get() immediately returns the provided value.
//
// Example:
//
//	immediate := lxtypes.LazyEager(42)
//	value, _ := immediate.Get() // Returns 42 immediately
func LazyEager[T any](value T) Lazy[T] {
	return eagerLazy[T]{value: value}
}

// LazyEagerOrError creates a Lazy that wraps an already-computed value or error.
// Useful for converting existing results into the Lazy interface.
//
// Example:
//
//	result, err := someFunction()
//	lazy := lxtypes.LazyEagerOrError(result, err)
//	value, err := lazy.Get() // Returns result and err immediately
func LazyEagerOrError[T any](value T, err error) Lazy[T] {
	return eagerLazy[T]{value: value, err: err}
}

// LazyDeferred creates a Lazy that computes its value on first access.
// The computation function is called at most once, even with concurrent access.
//
// Example:
//
//	expensive := lxtypes.LazyDeferred(func() (int, error) {
//	    // Expensive computation
//	    return 42, nil
//	})
//	value, _ := expensive.Get() // Computed here
//	value2, _ := expensive.Get() // Returns cached value
func LazyDeferred[T any](fn func() (T, error)) Lazy[T] {
	return &deferredLazy[T]{fn: fn}
}

// ----------------------------------- Eager Lazy -----------------------------------
type eagerLazy[T any] struct {
	value T
	err   error
}

func (e eagerLazy[T]) Get() (T, error) {
	return e.value, e.err
}

func (e eagerLazy[T]) MustGet() T {
	if e.err != nil {
		panic(e.err)
	}
	return e.value
}

func (e eagerLazy[T]) IsEvaluated() bool {
	return true
}

// ----------------------------------- Deferred Lazy -----------------------------------
type deferredLazy[T any] struct {
	fn        func() (T, error)
	once      sync.Once
	cache     T
	err       error
	evaluated int32 // 0 for false, 1 for true
}

func (d *deferredLazy[T]) Get() (T, error) {
	d.once.Do(func() {
		d.cache, d.err = d.fn()
		atomic.StoreInt32(&d.evaluated, 1)
	})
	return d.cache, d.err
}

func (d *deferredLazy[T]) MustGet() T {
	value, err := d.Get()
	if err != nil {
		panic(err)
	}
	return value
}

func (d *deferredLazy[T]) IsEvaluated() bool {
	return atomic.LoadInt32(&d.evaluated) == 1
}
