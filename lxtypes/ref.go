package lxtypes

import "sync"

// Ref represents a thread-safe mutable value cell.
// It provides safe concurrent read and write access to a single value
// using a read-write mutex internally.
//
// A Ref can be created with:
//   - NewRef(value) - Creates a Ref holding the given initial value
//
// Example:
//
//	counter := lxtypes.NewRef(0)
//	counter.Update(func(v int) int { return v + 1 })
//	fmt.Println(counter.Get())  // 1
type Ref[T any] interface {
	// Get returns the current value.
	Get() T

	// Set replaces the current value with the given value.
	Set(value T)

	// Update atomically applies fn to the current value and stores the result.
	// fn is called under the write lock, so it is safe for concurrent use.
	//
	// Example:
	//
	//  ref := lxtypes.NewRef(0)
	//  ref.Update(func(v int) int { return v + 1 })
	Update(fn func(T) T)
}

// NewRef creates a new Ref holding the given initial value.
//
// Example:
//
//	ref := lxtypes.NewRef(42)
//	fmt.Println(ref.Get())  // 42
func NewRef[T any](value T) Ref[T] {
	return &ref[T]{value: value}
}

type ref[T any] struct {
	mu    sync.RWMutex
	value T
}

func (r *ref[T]) Get() T {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.value
}

func (r *ref[T]) Set(value T) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.value = value
}

func (r *ref[T]) Update(fn func(T) T) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.value = fn(r.value)
}
