package lxtypes

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// ========================================
// LazyEager Tests
// ========================================

func TestLazyEager_Get(t *testing.T) {
	tests := []struct {
		name          string
		value         interface{}
		expectedValue interface{}
		expectedError error
	}{
		{
			name:          "integer value",
			value:         42,
			expectedValue: 42,
			expectedError: nil,
		},
		{
			name:          "string value",
			value:         "hello",
			expectedValue: "hello",
			expectedError: nil,
		},
		{
			name:          "zero value",
			value:         0,
			expectedValue: 0,
			expectedError: nil,
		},
		{
			name:          "empty string",
			value:         "",
			expectedValue: "",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.value.(type) {
			case int:
				lazy := LazyEager(v)
				result, err := lazy.Get()
				if result != tt.expectedValue {
					t.Errorf("Get() = %v, want %v", result, tt.expectedValue)
				}
				if err != tt.expectedError {
					t.Errorf("Get() error = %v, want %v", err, tt.expectedError)
				}
			case string:
				lazy := LazyEager(v)
				result, err := lazy.Get()
				if result != tt.expectedValue {
					t.Errorf("Get() = %v, want %v", result, tt.expectedValue)
				}
				if err != tt.expectedError {
					t.Errorf("Get() error = %v, want %v", err, tt.expectedError)
				}
			}
		})
	}
}

func TestLazyEager_MustGet(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		lazy := LazyEager(123)
		result := lazy.MustGet()
		if result != 123 {
			t.Errorf("MustGet() = %v, want 123", result)
		}
	})

	t.Run("multiple types", func(t *testing.T) {
		intLazy := LazyEager(42)
		if intLazy.MustGet() != 42 {
			t.Error("int MustGet() failed")
		}

		strLazy := LazyEager("test")
		if strLazy.MustGet() != "test" {
			t.Error("string MustGet() failed")
		}

		type Person struct {
			Name string
			Age  int
		}
		personLazy := LazyEager(Person{Name: "Alice", Age: 30})
		person := personLazy.MustGet()
		if person.Name != "Alice" || person.Age != 30 {
			t.Error("struct MustGet() failed")
		}
	})
}

func TestLazyEager_IsEvaluated(t *testing.T) {
	lazy := LazyEager(100)

	// Should always be true for eager values
	if !lazy.IsEvaluated() {
		t.Error("IsEvaluated() should return true for eager values")
	}

	// Should still be true after Get
	_, _ = lazy.Get()
	if !lazy.IsEvaluated() {
		t.Error("IsEvaluated() should still be true after Get()")
	}
}

// ========================================
// LazyEagerOrError Tests
// ========================================

func TestLazyEagerOrError_Success(t *testing.T) {
	lazy := LazyEagerOrError(42, nil)

	value, err := lazy.Get()
	if value != 42 {
		t.Errorf("Get() = %v, want 42", value)
	}
	if err != nil {
		t.Errorf("Get() error = %v, want nil", err)
	}
}

func TestLazyEagerOrError_WithError(t *testing.T) {
	expectedErr := errors.New("test error")
	lazy := LazyEagerOrError(0, expectedErr)

	value, err := lazy.Get()
	if value != 0 {
		t.Errorf("Get() = %v, want 0", value)
	}
	if err != expectedErr {
		t.Errorf("Get() error = %v, want %v", err, expectedErr)
	}

	// Error should be consistent across calls
	_, err2 := lazy.Get()
	if err2 != expectedErr {
		t.Errorf("Second Get() error = %v, want %v", err2, expectedErr)
	}
}

func TestLazyEagerOrError_MustGet_Panic(t *testing.T) {
	expectedErr := errors.New("panic error")
	lazy := LazyEagerOrError(0, expectedErr)

	defer func() {
		r := recover()
		if r == nil {
			t.Error("MustGet() should have panicked")
		}
		if r != expectedErr {
			t.Errorf("panic value = %v, want %v", r, expectedErr)
		}
	}()

	lazy.MustGet() // Should panic
}

func TestLazyEagerOrError_IsEvaluated(t *testing.T) {
	lazy := LazyEagerOrError(42, nil)

	// Should always be true for eager values
	if !lazy.IsEvaluated() {
		t.Error("IsEvaluated() should return true for eager values")
	}
}

// ========================================
// LazyDeferred Tests
// ========================================

func TestLazyDeferred_Get(t *testing.T) {
	tests := []struct {
		name          string
		fn            func() (int, error)
		expectedValue int
		expectedError error
	}{
		{
			name: "successful computation",
			fn: func() (int, error) {
				return 42, nil
			},
			expectedValue: 42,
			expectedError: nil,
		},
		{
			name: "computation with error",
			fn: func() (int, error) {
				return 0, errors.New("computation failed")
			},
			expectedValue: 0,
			expectedError: errors.New("computation failed"),
		},
		{
			name: "zero value",
			fn: func() (int, error) {
				return 0, nil
			},
			expectedValue: 0,
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lazy := LazyDeferred(tt.fn)

			value, err := lazy.Get()
			if value != tt.expectedValue {
				t.Errorf("Get() = %v, want %v", value, tt.expectedValue)
			}
			if tt.expectedError != nil && err == nil {
				t.Errorf("Get() error = nil, want error")
			}
			if tt.expectedError == nil && err != nil {
				t.Errorf("Get() error = %v, want nil", err)
			}
			if tt.expectedError != nil && err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("Get() error = %v, want %v", err, tt.expectedError)
			}
		})
	}
}

func TestLazyDeferred_GetCalledOnce(t *testing.T) {
	callCount := 0

	lazy := LazyDeferred(func() (int, error) {
		callCount++
		return 42, nil
	})

	if callCount != 0 {
		t.Errorf("Function called before Get(), count = %d", callCount)
	}

	// First call
	value1, err1 := lazy.Get()
	if callCount != 1 {
		t.Errorf("After first Get(), call count = %d, want 1", callCount)
	}
	if value1 != 42 || err1 != nil {
		t.Errorf("First Get() = (%v, %v), want (42, nil)", value1, err1)
	}

	// Second call (should not execute function again)
	value2, err2 := lazy.Get()
	if callCount != 1 {
		t.Errorf("After second Get(), call count = %d, want 1", callCount)
	}
	if value2 != 42 || err2 != nil {
		t.Errorf("Second Get() = (%v, %v), want (42, nil)", value2, err2)
	}

	// Third call
	value3, _ := lazy.Get()
	if callCount != 1 {
		t.Errorf("After third Get(), call count = %d, want 1", callCount)
	}
	if value3 != 42 {
		t.Errorf("Third Get() = %v, want 42", value3)
	}
}

func TestLazyDeferred_ErrorCaching(t *testing.T) {
	callCount := 0
	expectedErr := errors.New("test error")

	lazy := LazyDeferred(func() (string, error) {
		callCount++
		return "", expectedErr
	})

	// First call
	_, err1 := lazy.Get()
	if callCount != 1 {
		t.Errorf("After first Get(), call count = %d, want 1", callCount)
	}
	if err1 == nil || err1.Error() != expectedErr.Error() {
		t.Errorf("First Get() error = %v, want %v", err1, expectedErr)
	}

	// Second call (should return cached error)
	_, err2 := lazy.Get()
	if callCount != 1 {
		t.Errorf("After second Get(), call count = %d, want 1", callCount)
	}
	if err2 == nil || err2.Error() != expectedErr.Error() {
		t.Errorf("Second Get() error = %v, want %v", err2, expectedErr)
	}
}

func TestLazyDeferred_MustGet(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		lazy := LazyDeferred(func() (int, error) {
			return 99, nil
		})

		result := lazy.MustGet()
		if result != 99 {
			t.Errorf("MustGet() = %v, want 99", result)
		}
	})

	t.Run("panic on error", func(t *testing.T) {
		expectedErr := errors.New("must panic")
		lazy := LazyDeferred(func() (int, error) {
			return 0, expectedErr
		})

		defer func() {
			r := recover()
			if r == nil {
				t.Error("MustGet() should have panicked")
			}
			if r != expectedErr {
				t.Errorf("panic value = %v, want %v", r, expectedErr)
			}
		}()

		lazy.MustGet() // Should panic
		t.Error("Should not reach here")
	})
}

func TestLazyDeferred_IsEvaluated(t *testing.T) {
	lazy := LazyDeferred(func() (int, error) {
		return 42, nil
	})

	// Should be false before Get
	if lazy.IsEvaluated() {
		t.Error("IsEvaluated() should be false before Get()")
	}

	// Call Get
	_, _ = lazy.Get()

	// Should be true after Get
	if !lazy.IsEvaluated() {
		t.Error("IsEvaluated() should be true after Get()")
	}
}

func TestLazyDeferred_IsEvaluated_AfterError(t *testing.T) {
	lazy := LazyDeferred(func() (int, error) {
		return 0, errors.New("error")
	})

	if lazy.IsEvaluated() {
		t.Error("IsEvaluated() should be false before Get()")
	}

	_, _ = lazy.Get()

	// Should be true even if computation failed
	if !lazy.IsEvaluated() {
		t.Error("IsEvaluated() should be true after Get(), even with error")
	}
}

// ========================================
// Concurrency Tests
// ========================================

func TestLazyDeferred_ConcurrentAccess(t *testing.T) {
	callCount := int32(0)

	lazy := LazyDeferred(func() (int, error) {
		atomic.AddInt32(&callCount, 1)
		time.Sleep(10 * time.Millisecond) // Simulate work
		return 42, nil
	})

	// Launch 100 concurrent goroutines
	const goroutines = 100
	var wg sync.WaitGroup
	wg.Add(goroutines)

	results := make([]int, goroutines)
	errs := make([]error, goroutines)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			defer wg.Done()
			results[idx], errs[idx] = lazy.Get()
		}(i)
	}

	wg.Wait()

	// Function should be called exactly once
	if atomic.LoadInt32(&callCount) != 1 {
		t.Errorf("Function called %d times, want 1", callCount)
	}

	// All results should be the same
	for i := 0; i < goroutines; i++ {
		if results[i] != 42 {
			t.Errorf("Goroutine %d: Get() = %v, want 42", i, results[i])
		}
		if errs[i] != nil {
			t.Errorf("Goroutine %d: Get() error = %v, want nil", i, errs[i])
		}
	}

	// Should be evaluated after concurrent access
	if !lazy.IsEvaluated() {
		t.Error("IsEvaluated() should be true after concurrent Get()")
	}
}

func TestLazyDeferred_ConcurrentMustGet(t *testing.T) {
	callCount := int32(0)

	lazy := LazyDeferred(func() (string, error) {
		atomic.AddInt32(&callCount, 1)
		return "concurrent", nil
	})

	const goroutines = 50
	var wg sync.WaitGroup
	wg.Add(goroutines)

	results := make([]string, goroutines)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			defer wg.Done()
			results[idx] = lazy.MustGet()
		}(i)
	}

	wg.Wait()

	// Function should be called exactly once
	if atomic.LoadInt32(&callCount) != 1 {
		t.Errorf("Function called %d times, want 1", callCount)
	}

	// All results should be the same
	for i := 0; i < goroutines; i++ {
		if results[i] != "concurrent" {
			t.Errorf("Goroutine %d: MustGet() = %v, want 'concurrent'", i, results[i])
		}
	}
}

// ========================================
// Type Variety Tests
// ========================================

func TestLazy_DifferentTypes(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		lazy := LazyDeferred(func() (int, error) {
			return 42, nil
		})
		value, err := lazy.Get()
		if value != 42 || err != nil {
			t.Errorf("Get() = (%v, %v), want (42, nil)", value, err)
		}
	})

	t.Run("string", func(t *testing.T) {
		lazy := LazyDeferred(func() (string, error) {
			return "test", nil
		})
		value, err := lazy.Get()
		if value != "test" || err != nil {
			t.Errorf("Get() = (%v, %v), want (test, nil)", value, err)
		}
	})

	t.Run("slice", func(t *testing.T) {
		lazy := LazyDeferred(func() ([]int, error) {
			return []int{1, 2, 3}, nil
		})
		value, err := lazy.Get()
		if len(value) != 3 || err != nil {
			t.Errorf("Get() = (%v, %v), want ([1 2 3], nil)", value, err)
		}
	})

	t.Run("map", func(t *testing.T) {
		lazy := LazyDeferred(func() (map[string]int, error) {
			return map[string]int{"a": 1}, nil
		})
		value, err := lazy.Get()
		if value["a"] != 1 || err != nil {
			t.Errorf("Get() = (%v, %v), want (map[a:1], nil)", value, err)
		}
	})

	t.Run("struct", func(t *testing.T) {
		type Data struct {
			ID   int
			Name string
		}
		lazy := LazyDeferred(func() (Data, error) {
			return Data{ID: 1, Name: "test"}, nil
		})
		value, err := lazy.Get()
		if value.ID != 1 || value.Name != "test" || err != nil {
			t.Errorf("Get() = (%v, %v), want ({1 test}, nil)", value, err)
		}
	})

	t.Run("pointer", func(t *testing.T) {
		type Node struct {
			Value int
		}
		lazy := LazyDeferred(func() (*Node, error) {
			return &Node{Value: 100}, nil
		})
		value, err := lazy.Get()
		if value == nil || value.Value != 100 || err != nil {
			t.Errorf("Get() failed for pointer type")
		}
	})
}

// ========================================
// Edge Case Tests
// ========================================

func TestLazyDeferred_PanicInFunction(t *testing.T) {
	lazy := LazyDeferred(func() (int, error) {
		panic("computation panic")
	})

	defer func() {
		r := recover()
		if r == nil {
			t.Error("Expected panic to propagate")
		}
		if r != "computation panic" {
			t.Errorf("panic value = %v, want 'computation panic'", r)
		}
	}()

	_, _ = lazy.Get()
}

func TestLazyDeferred_LongRunningComputation(t *testing.T) {
	start := time.Now()

	lazy := LazyDeferred(func() (string, error) {
		time.Sleep(50 * time.Millisecond)
		return "done", nil
	})

	// First Get should take time
	value, err := lazy.Get()
	elapsed := time.Since(start)

	if value != "done" || err != nil {
		t.Errorf("Get() = (%v, %v), want (done, nil)", value, err)
	}
	if elapsed < 50*time.Millisecond {
		t.Errorf("Get() took %v, should take at least 50ms", elapsed)
	}

	// Second Get should be instant
	start = time.Now()
	value2, _ := lazy.Get()
	elapsed2 := time.Since(start)

	if value2 != "done" {
		t.Errorf("Second Get() = %v, want done", value2)
	}
	if elapsed2 > 10*time.Millisecond {
		t.Errorf("Second Get() took %v, should be instant (< 10ms)", elapsed2)
	}
}

func TestLazyDeferred_NilReturn(t *testing.T) {
	lazy := LazyDeferred(func() (*string, error) {
		return nil, nil
	})

	value, err := lazy.Get()
	if value != nil {
		t.Errorf("Get() = %v, want nil", value)
	}
	if err != nil {
		t.Errorf("Get() error = %v, want nil", err)
	}
}

// ========================================
// Comparison Tests
// ========================================

func TestLazy_EagerVsDeferred(t *testing.T) {
	// Both should behave the same after evaluation
	eager := LazyEager(42)
	deferred := LazyDeferred(func() (int, error) {
		return 42, nil
	})

	eagerVal, eagerErr := eager.Get()
	deferredVal, deferredErr := deferred.Get()

	if eagerVal != deferredVal {
		t.Errorf("Values differ: eager=%v, deferred=%v", eagerVal, deferredVal)
	}
	if eagerErr != deferredErr {
		t.Errorf("Errors differ: eager=%v, deferred=%v", eagerErr, deferredErr)
	}

	// Both should be evaluated
	if !eager.IsEvaluated() {
		t.Error("Eager should be evaluated")
	}
	if !deferred.IsEvaluated() {
		t.Error("Deferred should be evaluated after Get()")
	}
}

// ========================================
// Benchmark Tests
// ========================================

func BenchmarkLazyEager_Get(b *testing.B) {
	lazy := LazyEager(42)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = lazy.Get()
	}
}

func BenchmarkLazyDeferred_FirstGet(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lazy := LazyDeferred(func() (int, error) {
			return 42, nil
		})
		_, _ = lazy.Get()
	}
}

func BenchmarkLazyDeferred_CachedGet(b *testing.B) {
	lazy := LazyDeferred(func() (int, error) {
		return 42, nil
	})
	_, _ = lazy.Get() // Prime the cache

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = lazy.Get()
	}
}

func BenchmarkLazyDeferred_ExpensiveComputation(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lazy := LazyDeferred(func() ([]int, error) {
			result := make([]int, 1000)
			for j := 0; j < 1000; j++ {
				result[j] = j * j
			}
			return result, nil
		})
		_, _ = lazy.Get()
	}
}

func BenchmarkLazyDeferred_ConcurrentAccess(b *testing.B) {
	lazy := LazyDeferred(func() (int, error) {
		time.Sleep(1 * time.Millisecond)
		return 42, nil
	})

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = lazy.Get()
		}
	})
}

// ========================================
// Interface Implementation Tests
// ========================================

func TestLazy_InterfaceImplementation(t *testing.T) {
	// Ensure both types implement Lazy interface
	var _ Lazy[int] = LazyEager(42)
	var _ Lazy[int] = LazyDeferred(func() (int, error) { return 42, nil })
	var _ Lazy[int] = LazyEagerOrError(42, nil)
}

// ========================================
// Real-World Scenario Tests
// ========================================

func TestLazy_ConfigurationLoading(t *testing.T) {
	// Simulate lazy configuration loading
	type Config struct {
		Host string
		Port int
	}

	loadCount := 0
	configLoader := LazyDeferred(func() (Config, error) {
		loadCount++
		// Simulate loading from file/env
		return Config{Host: "localhost", Port: 8080}, nil
	})

	// Configuration not loaded yet
	if loadCount != 0 {
		t.Errorf("Config loaded too early, count = %d", loadCount)
	}
	if configLoader.IsEvaluated() {
		t.Error("Config should not be evaluated yet")
	}

	// Load configuration when needed
	config, err := configLoader.Get()
	if err != nil {
		t.Errorf("Get() error = %v, want nil", err)
	}
	if config.Host != "localhost" || config.Port != 8080 {
		t.Errorf("Get() = %+v, want {localhost 8080}", config)
	}
	if loadCount != 1 {
		t.Errorf("Config loaded %d times, want 1", loadCount)
	}

	// Subsequent access uses cached config
	config2, _ := configLoader.Get()
	if config2.Host != config.Host || config2.Port != config.Port {
		t.Error("Cached config differs from original")
	}
	if loadCount != 1 {
		t.Errorf("Config loaded %d times, want 1", loadCount)
	}
}

func TestLazy_OptionalComputation(t *testing.T) {
	// Simulate computation that may or may not be needed
	expensive := LazyDeferred(func() ([]byte, error) {
		// Simulate expensive operation (e.g., reading large file)
		data := make([]byte, 1000)
		for i := range data {
			data[i] = byte(i % 256)
		}
		return data, nil
	})

	// In fast path, computation is never triggered
	fastPath := true
	if fastPath {
		// Do something that doesn't need the expensive computation
		if expensive.IsEvaluated() {
			t.Error("Expensive computation should not be triggered in fast path")
		}
	}

	// In slow path, computation happens
	fastPath = false
	if !fastPath {
		data, err := expensive.Get()
		if err != nil || len(data) != 1000 {
			t.Errorf("Slow path failed: data len=%d, err=%v", len(data), err)
		}
		if !expensive.IsEvaluated() {
			t.Error("Should be evaluated after Get()")
		}
	}
}

// ========================================
// Mixed Eager/Deferred Tests
// ========================================

func TestLazy_MixedUsage(t *testing.T) {
	// Create a slice of Lazy values (some eager, some deferred)
	lazies := []Lazy[int]{
		LazyEager(1),
		LazyDeferred(func() (int, error) { return 2, nil }),
		LazyEagerOrError(3, nil),
		LazyDeferred(func() (int, error) { return 4, nil }),
		LazyEager(5),
	}

	// Process all of them uniformly
	sum := 0
	for _, lazy := range lazies {
		value := lazy.MustGet()
		sum += value
	}

	if sum != 15 {
		t.Errorf("Sum = %d, want 15", sum)
	}

	// All should be evaluated now
	for i, lazy := range lazies {
		if !lazy.IsEvaluated() {
			t.Errorf("Lazy[%d] should be evaluated", i)
		}
	}
}

// ========================================
// Memory and Performance Tests
// ========================================

func TestLazyDeferred_MemoryEfficiency(t *testing.T) {
	// Create many lazy values
	const count = 1000
	lazies := make([]Lazy[int], count)

	for i := 0; i < count; i++ {
		val := i
		lazies[i] = LazyDeferred(func() (int, error) {
			return val * val, nil
		})
	}

	// Only evaluate half of them
	for i := 0; i < count/2; i++ {
		lazies[i].Get()
	}

	// Check evaluation status
	evaluatedCount := 0
	for _, lazy := range lazies {
		if lazy.IsEvaluated() {
			evaluatedCount++
		}
	}

	if evaluatedCount != count/2 {
		t.Errorf("Evaluated count = %d, want %d", evaluatedCount, count/2)
	}
}

// ========================================
// Edge Case: Zero Values
// ========================================

func TestLazy_ZeroValues(t *testing.T) {
	t.Run("zero int", func(t *testing.T) {
		lazy := LazyEager(0)
		value, err := lazy.Get()
		if value != 0 || err != nil {
			t.Errorf("Get() = (%v, %v), want (0, nil)", value, err)
		}
	})

	t.Run("empty string", func(t *testing.T) {
		lazy := LazyEager("")
		value, err := lazy.Get()
		if value != "" || err != nil {
			t.Errorf("Get() = (%v, %v), want ('', nil)", value, err)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var nilSlice []int
		lazy := LazyEager(nilSlice)
		value, err := lazy.Get()
		if value != nil || err != nil {
			t.Errorf("Get() = (%v, %v), want (nil, nil)", value, err)
		}
	})

	t.Run("nil pointer", func(t *testing.T) {
		var nilPtr *string
		lazy := LazyEager(nilPtr)
		value, err := lazy.Get()
		if value != nil || err != nil {
			t.Errorf("Get() = (%v, %v), want (nil, nil)", value, err)
		}
	})
}

// ========================================
// MustGet Panic Recovery Tests
// ========================================

func TestLazyDeferred_MustGet_PanicRecovery(t *testing.T) {
	tests := []struct {
		name        string
		lazy        Lazy[int]
		shouldPanic bool
		panicValue  interface{}
	}{
		{
			name: "success - no panic",
			lazy: LazyDeferred(func() (int, error) {
				return 42, nil
			}),
			shouldPanic: false,
		},
		{
			name: "error - panics",
			lazy: LazyDeferred(func() (int, error) {
				return 0, errors.New("test error")
			}),
			shouldPanic: true,
			panicValue:  errors.New("test error"),
		},
		{
			name:        "eager error - panics",
			lazy:        LazyEagerOrError(0, errors.New("eager error")),
			shouldPanic: true,
			panicValue:  errors.New("eager error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					r := recover()
					if r == nil {
						t.Error("MustGet() should have panicked")
					}
					// Check error message matches
					if err, ok := r.(error); ok {
						expectedErr := tt.panicValue.(error)
						if err.Error() != expectedErr.Error() {
							t.Errorf("panic error = %v, want %v", err, expectedErr)
						}
					}
				}()
			}

			result := tt.lazy.MustGet()
			if tt.shouldPanic {
				t.Error("Should not reach here")
			}
			if !tt.shouldPanic && result != 42 {
				t.Errorf("MustGet() = %v, want 42", result)
			}
		})
	}
}

// ========================================
// IsEvaluated Consistency Tests
// ========================================

func TestLazy_IsEvaluatedConsistency(t *testing.T) {
	t.Run("eager always evaluated", func(t *testing.T) {
		lazy := LazyEager(100)

		// Check multiple times
		for i := 0; i < 5; i++ {
			if !lazy.IsEvaluated() {
				t.Errorf("IsEvaluated() call %d returned false", i)
			}
		}
	})

	t.Run("deferred evaluated after Get", func(t *testing.T) {
		lazy := LazyDeferred(func() (int, error) {
			return 42, nil
		})

		// Before Get
		if lazy.IsEvaluated() {
			t.Error("IsEvaluated() should be false before Get()")
		}

		// After Get
		_, _ = lazy.Get()
		for i := 0; i < 5; i++ {
			if !lazy.IsEvaluated() {
				t.Errorf("IsEvaluated() call %d returned false after Get()", i)
			}
		}
	})

	t.Run("deferred evaluated after MustGet", func(t *testing.T) {
		lazy := LazyDeferred(func() (int, error) {
			return 42, nil
		})

		// Before MustGet
		if lazy.IsEvaluated() {
			t.Error("IsEvaluated() should be false before MustGet()")
		}

		// After MustGet
		lazy.MustGet()
		if !lazy.IsEvaluated() {
			t.Error("IsEvaluated() should be true after MustGet()")
		}
	})
}
