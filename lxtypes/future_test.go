package lxtypes

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// ========================================
// FutureDo Tests
// ========================================

func TestFutureDo_Success(t *testing.T) {
	tests := []struct {
		name     string
		fn       func() (int, error)
		expected int
	}{
		{
			name: "immediate success",
			fn: func() (int, error) {
				return 42, nil
			},
			expected: 42,
		},
		{
			name: "zero value",
			fn: func() (int, error) {
				return 0, nil
			},
			expected: 0,
		},
		{
			name: "delayed computation",
			fn: func() (int, error) {
				time.Sleep(10 * time.Millisecond)
				return 100, nil
			},
			expected: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			future := FutureDo(tt.fn)
			ctx := context.Background()

			result, err := future.Get(ctx)
			if err != nil {
				t.Errorf("Get() returned unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Get() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFutureDo_Error(t *testing.T) {
	expectedErr := errors.New("computation failed")

	future := FutureDo(func() (int, error) {
		return 0, expectedErr
	})

	ctx := context.Background()
	result, err := future.Get(ctx)

	if err != expectedErr {
		t.Errorf("Get() error = %v, want %v", err, expectedErr)
	}
	if result != 0 {
		t.Errorf("Get() result = %v, want 0", result)
	}
}

func TestFutureDo_ContextCancellation(t *testing.T) {
	future := FutureDo(func() (int, error) {
		time.Sleep(100 * time.Millisecond)
		return 42, nil
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	result, err := future.Get(ctx)

	if err != context.Canceled {
		t.Errorf("Get() error = %v, want %v", err, context.Canceled)
	}
	if result != 0 {
		t.Errorf("Get() result = %v, want 0", result)
	}
}

func TestFutureDo_ContextTimeout(t *testing.T) {
	future := FutureDo(func() (int, error) {
		time.Sleep(200 * time.Millisecond)
		return 42, nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	result, err := future.Get(ctx)

	if err != context.DeadlineExceeded {
		t.Errorf("Get() error = %v, want %v", err, context.DeadlineExceeded)
	}
	if result != 0 {
		t.Errorf("Get() result = %v, want 0", result)
	}
}

func TestFutureDo_MultipleGets(t *testing.T) {
	callCount := int32(0)
	future := FutureDo(func() (int, error) {
		atomic.AddInt32(&callCount, 1)
		return 42, nil
	})

	ctx := context.Background()

	// Call Get multiple times
	for i := 0; i < 5; i++ {
		result, err := future.Get(ctx)
		if err != nil {
			t.Errorf("Get() iteration %d returned error: %v", i, err)
		}
		if result != 42 {
			t.Errorf("Get() iteration %d = %v, want 42", i, result)
		}
	}

	// Verify function was called only once
	if atomic.LoadInt32(&callCount) != 1 {
		t.Errorf("function called %d times, want 1", atomic.LoadInt32(&callCount))
	}
}

func TestFutureDo_ConcurrentGets(t *testing.T) {
	callCount := int32(0)
	future := FutureDo(func() (int, error) {
		atomic.AddInt32(&callCount, 1)
		time.Sleep(50 * time.Millisecond)
		return 42, nil
	})

	ctx := context.Background()
	var wg sync.WaitGroup
	numGoroutines := 10

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result, err := future.Get(ctx)
			if err != nil {
				t.Errorf("Get() returned error: %v", err)
			}
			if result != 42 {
				t.Errorf("Get() = %v, want 42", result)
			}
		}()
	}

	wg.Wait()

	// Verify function was called only once
	if atomic.LoadInt32(&callCount) != 1 {
		t.Errorf("function called %d times, want 1", atomic.LoadInt32(&callCount))
	}
}

// ========================================
// FutureOf Tests
// ========================================

func TestFutureOf(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected interface{}
	}{
		{
			name:     "integer",
			value:    42,
			expected: 42,
		},
		{
			name:     "string",
			value:    "hello",
			expected: "hello",
		},
		{
			name:     "zero int",
			value:    0,
			expected: 0,
		},
		{
			name:     "empty string",
			value:    "",
			expected: "",
		},
		{
			name:     "struct",
			value:    struct{ Name string }{"Alice"},
			expected: struct{ Name string }{"Alice"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			switch v := tt.value.(type) {
			case int:
				future := FutureOf(v)
				result, err := future.Get(ctx)
				if err != nil {
					t.Errorf("Get() returned error: %v", err)
				}
				if result != tt.expected.(int) {
					t.Errorf("Get() = %v, want %v", result, tt.expected)
				}
			case string:
				future := FutureOf(v)
				result, err := future.Get(ctx)
				if err != nil {
					t.Errorf("Get() returned error: %v", err)
				}
				if result != tt.expected.(string) {
					t.Errorf("Get() = %v, want %v", result, tt.expected)
				}
			case struct{ Name string }:
				future := FutureOf(v)
				result, err := future.Get(ctx)
				if err != nil {
					t.Errorf("Get() returned error: %v", err)
				}
				if result != tt.expected.(struct{ Name string }) {
					t.Errorf("Get() = %v, want %v", result, tt.expected)
				}
			}
		})
	}
}

func TestFutureOf_ImmediateCompletion(t *testing.T) {
	future := FutureOf(42)
	ctx := context.Background()

	// Should return immediately without blocking
	start := time.Now()
	result, err := future.Get(ctx)
	duration := time.Since(start)

	if err != nil {
		t.Errorf("Get() returned error: %v", err)
	}
	if result != 42 {
		t.Errorf("Get() = %v, want 42", result)
	}
	if duration > 10*time.Millisecond {
		t.Errorf("Get() took %v, should be immediate", duration)
	}
}

// ========================================
// FutureError Tests
// ========================================

func TestFutureError(t *testing.T) {
	expectedErr := errors.New("test error")
	future := FutureError[int](expectedErr)

	ctx := context.Background()
	result, err := future.Get(ctx)

	if err != expectedErr {
		t.Errorf("Get() error = %v, want %v", err, expectedErr)
	}
	if result != 0 {
		t.Errorf("Get() result = %v, want 0", result)
	}
}

func TestFutureError_ImmediateCompletion(t *testing.T) {
	expectedErr := errors.New("test error")
	future := FutureError[string](expectedErr)

	ctx := context.Background()

	// Should return immediately without blocking
	start := time.Now()
	result, err := future.Get(ctx)
	duration := time.Since(start)

	if err != expectedErr {
		t.Errorf("Get() error = %v, want %v", err, expectedErr)
	}
	if result != "" {
		t.Errorf("Get() result = %q, want empty string", result)
	}
	if duration > 10*time.Millisecond {
		t.Errorf("Get() took %v, should be immediate", duration)
	}
}

// ========================================
// FutureThen Tests
// ========================================

func TestFutureThen_Success(t *testing.T) {
	// Create a chain: int -> int -> string
	future := FutureDo(func() (int, error) {
		return 5, nil
	})

	future2 := FutureThen(future, func(n int) (int, error) {
		return n * 2, nil
	})

	future3 := FutureThen(future2, func(n int) (string, error) {
		return "result: " + string(rune('0'+n)), nil
	})

	ctx := context.Background()
	result, err := future3.Get(ctx)

	if err != nil {
		t.Errorf("Get() returned error: %v", err)
	}
	// 5 * 2 = 10, '0' + 10 = ':'
	expected := "result: :"
	if result != expected {
		t.Errorf("Get() = %q, want %q", result, expected)
	}
}

func TestFutureThen_ParentError(t *testing.T) {
	expectedErr := errors.New("parent error")

	future := FutureDo(func() (int, error) {
		return 0, expectedErr
	})

	callCount := int32(0)
	future2 := FutureThen(future, func(n int) (string, error) {
		atomic.AddInt32(&callCount, 1)
		return "should not be called", nil
	})

	ctx := context.Background()
	result, err := future2.Get(ctx)

	if err != expectedErr {
		t.Errorf("Get() error = %v, want %v", err, expectedErr)
	}
	if result != "" {
		t.Errorf("Get() result = %q, want empty string", result)
	}
	if atomic.LoadInt32(&callCount) != 0 {
		t.Errorf("transform function called, should not be called when parent errors")
	}
}

func TestFutureThen_TransformError(t *testing.T) {
	expectedErr := errors.New("transform error")

	future := FutureDo(func() (int, error) {
		return 42, nil
	})

	future2 := FutureThen(future, func(n int) (string, error) {
		return "", expectedErr
	})

	ctx := context.Background()
	result, err := future2.Get(ctx)

	if err != expectedErr {
		t.Errorf("Get() error = %v, want %v", err, expectedErr)
	}
	if result != "" {
		t.Errorf("Get() result = %q, want empty string", result)
	}
}

func TestFutureThen_ContextCancellation(t *testing.T) {
	future := FutureDo(func() (int, error) {
		time.Sleep(100 * time.Millisecond)
		return 42, nil
	})

	future2 := FutureThen(future, func(n int) (string, error) {
		return "transformed", nil
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	result, err := future2.Get(ctx)

	if err != context.Canceled {
		t.Errorf("Get() error = %v, want %v", err, context.Canceled)
	}
	if result != "" {
		t.Errorf("Get() result = %q, want empty string", result)
	}
}

func TestFutureThen_LongChain(t *testing.T) {
	// Create a long chain of transformations
	future := FutureDo(func() (int, error) {
		return 1, nil
	})

	// Chain multiple transformations
	for i := 0; i < 10; i++ {
		future = FutureThen(future, func(n int) (int, error) {
			return n + 1, nil
		})
	}

	ctx := context.Background()
	result, err := future.Get(ctx)

	if err != nil {
		t.Errorf("Get() returned error: %v", err)
	}
	if result != 11 {
		t.Errorf("Get() = %v, want 11", result)
	}
}

// ========================================
// FutureAll Tests
// ========================================

func TestFutureAll_AllSuccess(t *testing.T) {
	f1 := FutureDo(func() (int, error) {
		time.Sleep(10 * time.Millisecond)
		return 1, nil
	})
	f2 := FutureDo(func() (int, error) {
		time.Sleep(20 * time.Millisecond)
		return 2, nil
	})
	f3 := FutureDo(func() (int, error) {
		time.Sleep(5 * time.Millisecond)
		return 3, nil
	})

	allFuture := FutureAll(f1, f2, f3)

	ctx := context.Background()
	result, err := allFuture.Get(ctx)

	if err != nil {
		t.Errorf("Get() returned error: %v", err)
	}
	if len(result) != 3 {
		t.Errorf("Get() returned %d results, want 3", len(result))
	}
	if result[0] != 1 || result[1] != 2 || result[2] != 3 {
		t.Errorf("Get() = %v, want [1 2 3]", result)
	}
}

func TestFutureAll_OneError(t *testing.T) {
	expectedErr := errors.New("future 2 failed")

	f1 := FutureDo(func() (int, error) {
		return 1, nil
	})
	f2 := FutureDo(func() (int, error) {
		return 0, expectedErr
	})
	f3 := FutureDo(func() (int, error) {
		return 3, nil
	})

	allFuture := FutureAll(f1, f2, f3)

	ctx := context.Background()
	result, err := allFuture.Get(ctx)

	if err != expectedErr {
		t.Errorf("Get() error = %v, want %v", err, expectedErr)
	}
	if result != nil {
		t.Errorf("Get() result = %v, want nil", result)
	}
}

func TestFutureAll_MultipleErrors(t *testing.T) {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	f1 := FutureDo(func() (int, error) {
		return 0, err1
	})
	f2 := FutureDo(func() (int, error) {
		return 0, err2
	})

	allFuture := FutureAll(f1, f2)

	ctx := context.Background()
	_, err := allFuture.Get(ctx)

	// Should return first error (err1 or err2)
	if err != err1 && err != err2 {
		t.Errorf("Get() error = %v, want %v or %v", err, err1, err2)
	}
}

func TestFutureAll_Empty(t *testing.T) {
	allFuture := FutureAll[int]()

	ctx := context.Background()
	result, err := allFuture.Get(ctx)

	if err != nil {
		t.Errorf("Get() returned error: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Get() returned %d results, want 0", len(result))
	}
}

func TestFutureAll_ContextCancellation(t *testing.T) {
	f1 := FutureDo(func() (int, error) {
		time.Sleep(100 * time.Millisecond)
		return 1, nil
	})
	f2 := FutureDo(func() (int, error) {
		time.Sleep(200 * time.Millisecond)
		return 2, nil
	})

	allFuture := FutureAll(f1, f2)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	result, err := allFuture.Get(ctx)

	if err != context.DeadlineExceeded {
		t.Errorf("Get() error = %v, want %v", err, context.DeadlineExceeded)
	}
	if result != nil {
		t.Errorf("Get() result = %v, want nil", result)
	}
}

// ========================================
// FutureJoin2 Tests
// ========================================

func TestFutureJoin2_Success(t *testing.T) {
	f1 := FutureDo(func() (int, error) {
		time.Sleep(10 * time.Millisecond)
		return 42, nil
	})
	f2 := FutureDo(func() (string, error) {
		time.Sleep(20 * time.Millisecond)
		return "hello", nil
	})

	joined := FutureJoin2(f1, f2)

	ctx := context.Background()
	result, err := joined.Get(ctx)

	if err != nil {
		t.Errorf("Get() returned error: %v", err)
	}
	if result.First != 42 {
		t.Errorf("result.First = %v, want 42", result.First)
	}
	if result.Second != "hello" {
		t.Errorf("result.Second = %v, want hello", result.Second)
	}
}

func TestFutureJoin2_FirstError(t *testing.T) {
	expectedErr := errors.New("first error")

	f1 := FutureDo(func() (int, error) {
		return 0, expectedErr
	})
	f2 := FutureDo(func() (string, error) {
		return "hello", nil
	})

	joined := FutureJoin2(f1, f2)

	ctx := context.Background()
	result, err := joined.Get(ctx)

	if err != expectedErr {
		t.Errorf("Get() error = %v, want %v", err, expectedErr)
	}
	if result.First != 0 || result.Second != "" {
		t.Errorf("Get() result = %v, want zero values", result)
	}
}

func TestFutureJoin2_SecondError(t *testing.T) {
	expectedErr := errors.New("second error")

	f1 := FutureDo(func() (int, error) {
		return 42, nil
	})
	f2 := FutureDo(func() (string, error) {
		return "", expectedErr
	})

	joined := FutureJoin2(f1, f2)

	ctx := context.Background()
	result, err := joined.Get(ctx)

	if err != expectedErr {
		t.Errorf("Get() error = %v, want %v", err, expectedErr)
	}
	if result.First != 0 || result.Second != "" {
		t.Errorf("Get() result = %v, want zero values", result)
	}
}

// ========================================
// FutureJoin3 Tests
// ========================================

func TestFutureJoin3_Success(t *testing.T) {
	f1 := FutureDo(func() (int, error) {
		return 1, nil
	})
	f2 := FutureDo(func() (string, error) {
		return "two", nil
	})
	f3 := FutureDo(func() (bool, error) {
		return true, nil
	})

	joined := FutureJoin3(f1, f2, f3)

	ctx := context.Background()
	result, err := joined.Get(ctx)

	if err != nil {
		t.Errorf("Get() returned error: %v", err)
	}
	if result.First != 1 {
		t.Errorf("result.First = %v, want 1", result.First)
	}
	if result.Second != "two" {
		t.Errorf("result.Second = %v, want two", result.Second)
	}
	if result.Third != true {
		t.Errorf("result.Third = %v, want true", result.Third)
	}
}

func TestFutureJoin3_Error(t *testing.T) {
	expectedErr := errors.New("error in third")

	f1 := FutureDo(func() (int, error) {
		return 1, nil
	})
	f2 := FutureDo(func() (string, error) {
		return "two", nil
	})
	f3 := FutureDo(func() (bool, error) {
		return false, expectedErr
	})

	joined := FutureJoin3(f1, f2, f3)

	ctx := context.Background()
	_, err := joined.Get(ctx)

	if err != expectedErr {
		t.Errorf("Get() error = %v, want %v", err, expectedErr)
	}
}

// ========================================
// FutureJoin4 Tests
// ========================================

func TestFutureJoin4_Success(t *testing.T) {
	f1 := FutureDo(func() (int, error) {
		return 1, nil
	})
	f2 := FutureDo(func() (string, error) {
		return "two", nil
	})
	f3 := FutureDo(func() (bool, error) {
		return true, nil
	})
	f4 := FutureDo(func() (float64, error) {
		return 4.0, nil
	})

	joined := FutureJoin4(f1, f2, f3, f4)

	ctx := context.Background()
	result, err := joined.Get(ctx)

	if err != nil {
		t.Errorf("Get() returned error: %v", err)
	}
	if result.First != 1 {
		t.Errorf("result.First = %v, want 1", result.First)
	}
	if result.Second != "two" {
		t.Errorf("result.Second = %v, want two", result.Second)
	}
	if result.Third != true {
		t.Errorf("result.Third = %v, want true", result.Third)
	}
	if result.Fourth != 4.0 {
		t.Errorf("result.Fourth = %v, want 4.0", result.Fourth)
	}
}

func TestFutureJoin4_Error(t *testing.T) {
	expectedErr := errors.New("error in fourth")

	f1 := FutureDo(func() (int, error) {
		return 1, nil
	})
	f2 := FutureDo(func() (string, error) {
		return "two", nil
	})
	f3 := FutureDo(func() (bool, error) {
		return true, nil
	})
	f4 := FutureDo(func() (float64, error) {
		return 0, expectedErr
	})

	joined := FutureJoin4(f1, f2, f3, f4)

	ctx := context.Background()
	_, err := joined.Get(ctx)

	if err != expectedErr {
		t.Errorf("Get() error = %v, want %v", err, expectedErr)
	}
}

// ========================================
// Real-World Scenario Tests
// ========================================

type User struct {
	ID   int
	Name string
}

type Config struct {
	APIKey string
	Host   string
}

type Stats struct {
	Count int
	Total float64
}

func TestRealWorld_SequentialChain(t *testing.T) {
	// Simulate: getUserId -> fetchUser -> fetchUserStats

	// Step 1: Get user ID
	userIdFuture := FutureDo(func() (int, error) {
		time.Sleep(10 * time.Millisecond)
		return 123, nil
	})

	// Step 2: Fetch user
	userFuture := FutureThen(userIdFuture, func(id int) (User, error) {
		time.Sleep(10 * time.Millisecond)
		return User{ID: id, Name: "Alice"}, nil
	})

	// Step 3: Fetch stats
	statsFuture := FutureThen(userFuture, func(user User) (Stats, error) {
		time.Sleep(10 * time.Millisecond)
		return Stats{Count: 10, Total: 100.0}, nil
	})

	ctx := context.Background()
	stats, err := statsFuture.Get(ctx)

	if err != nil {
		t.Errorf("Get() returned error: %v", err)
	}
	if stats.Count != 10 || stats.Total != 100.0 {
		t.Errorf("Get() = %v, want {10 100.0}", stats)
	}
}

func TestRealWorld_ParallelFetch(t *testing.T) {
	// Simulate: fetch user, config, and stats in parallel

	userFuture := FutureDo(func() (User, error) {
		time.Sleep(30 * time.Millisecond)
		return User{ID: 1, Name: "Alice"}, nil
	})

	configFuture := FutureDo(func() (Config, error) {
		time.Sleep(20 * time.Millisecond)
		return Config{APIKey: "key123", Host: "api.example.com"}, nil
	})

	statsFuture := FutureDo(func() (Stats, error) {
		time.Sleep(10 * time.Millisecond)
		return Stats{Count: 5, Total: 50.0}, nil
	})

	// Combine all three
	combined := FutureJoin3(userFuture, configFuture, statsFuture)

	ctx := context.Background()
	start := time.Now()
	result, err := combined.Get(ctx)
	duration := time.Since(start)

	if err != nil {
		t.Errorf("Get() returned error: %v", err)
	}
	if result.First.Name != "Alice" {
		t.Errorf("User name = %v, want Alice", result.First.Name)
	}
	if result.Second.APIKey != "key123" {
		t.Errorf("Config APIKey = %v, want key123", result.Second.APIKey)
	}
	if result.Third.Count != 5 {
		t.Errorf("Stats count = %v, want 5", result.Third.Count)
	}

	// Should complete in ~30ms (longest task), not 60ms (sum of all tasks)
	if duration > 50*time.Millisecond {
		t.Errorf("parallel execution took %v, expected ~30ms", duration)
	}
}

func TestRealWorld_FallbackOnError(t *testing.T) {
	// Try primary source, fallback to secondary on error

	primaryFuture := FutureDo(func() (string, error) {
		return "", errors.New("primary failed")
	})

	secondaryFuture := FutureDo(func() (string, error) {
		time.Sleep(10 * time.Millisecond)
		return "fallback data", nil
	})

	ctx := context.Background()

	// Try primary first
	data, err := primaryFuture.Get(ctx)
	if err != nil {
		// Fallback to secondary
		data, err = secondaryFuture.Get(ctx)
	}

	if err != nil {
		t.Errorf("Fallback failed: %v", err)
	}
	if data != "fallback data" {
		t.Errorf("data = %v, want 'fallback data'", data)
	}
}

func TestRealWorld_TransformCombinedResults(t *testing.T) {
	// Fetch multiple sources in parallel, then transform combined result

	data1 := FutureDo(func() (int, error) {
		return 10, nil
	})
	data2 := FutureDo(func() (int, error) {
		return 20, nil
	})
	data3 := FutureDo(func() (int, error) {
		return 30, nil
	})

	// Combine into slice
	allData := FutureAll(data1, data2, data3)

	// Transform: sum all values
	sumFuture := FutureThen(allData, func(values []int) (int, error) {
		sum := 0
		for _, v := range values {
			sum += v
		}
		return sum, nil
	})

	ctx := context.Background()
	result, err := sumFuture.Get(ctx)

	if err != nil {
		t.Errorf("Get() returned error: %v", err)
	}
	if result != 60 {
		t.Errorf("sum = %v, want 60", result)
	}
}

// ========================================
// Benchmark Tests
// ========================================

func BenchmarkFutureDo_Simple(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		future := FutureDo(func() (int, error) {
			return 42, nil
		})
		_, _ = future.Get(ctx)
	}
}

func BenchmarkFutureDo_ConcurrentGets(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		future := FutureDo(func() (int, error) {
			return 42, nil
		})

		var wg sync.WaitGroup
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				_, _ = future.Get(ctx)
			}()
		}
		wg.Wait()
	}
}

func BenchmarkFutureThen_Chain(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		future := FutureDo(func() (int, error) {
			return 1, nil
		})

		future = FutureThen(future, func(n int) (int, error) {
			return n + 1, nil
		})
		future = FutureThen(future, func(n int) (int, error) {
			return n * 2, nil
		})
		future = FutureThen(future, func(n int) (int, error) {
			return n - 1, nil
		})

		_, _ = future.Get(ctx)
	}
}

func BenchmarkFutureAll_Parallel(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		futures := make([]Future[int], 10)
		for j := 0; j < 10; j++ {
			val := j
			futures[j] = FutureDo(func() (int, error) {
				return val, nil
			})
		}

		allFuture := FutureAll(futures...)
		_, _ = allFuture.Get(ctx)
	}
}

func TestFutureJoin5_Success(t *testing.T) {
	f1 := FutureDo(func() (int, error) {
		return 1, nil
	})
	f2 := FutureDo(func() (string, error) {
		return "two", nil
	})
	f3 := FutureDo(func() (bool, error) {
		return true, nil
	})
	f4 := FutureDo(func() (float64, error) {
		return 4.0, nil
	})
	f5 := FutureDo(func() ([]int, error) {
		return []int{5, 6, 7}, nil
	})

	joined := FutureJoin5(f1, f2, f3, f4, f5)

	ctx := context.Background()
	result, err := joined.Get(ctx)

	if err != nil {
		t.Errorf("Get() returned error: %v", err)
	}
	if result.V1 != 1 {
		t.Errorf("result.V1 = %v, want 1", result.V1)
	}
	if result.V2 != "two" {
		t.Errorf("result.V2 = %v, want two", result.V2)
	}
	if result.V3 != true {
		t.Errorf("result.V3 = %v, want true", result.V3)
	}
	if result.V4 != 4.0 {
		t.Errorf("result.V4 = %v, want 4.0", result.V4)
	}
	if len(result.V5) != 3 || result.V5[0] != 5 {
		t.Errorf("result.V5 = %v, want [5 6 7]", result.V5)
	}
}

func TestFutureJoin5_Error(t *testing.T) {
	expectedErr := errors.New("error in third")

	f1 := FutureDo(func() (int, error) {
		return 1, nil
	})
	f2 := FutureDo(func() (string, error) {
		return "two", nil
	})
	f3 := FutureDo(func() (bool, error) {
		return false, expectedErr
	})
	f4 := FutureDo(func() (float64, error) {
		return 4.0, nil
	})
	f5 := FutureDo(func() ([]int, error) {
		return []int{5}, nil
	})

	joined := FutureJoin5(f1, f2, f3, f4, f5)

	ctx := context.Background()
	_, err := joined.Get(ctx)

	if err != expectedErr {
		t.Errorf("Get() error = %v, want %v", err, expectedErr)
	}
}

func TestFutureJoin5_Parallel(t *testing.T) {
	// Each future takes 50ms
	// If sequential: 250ms total
	// If parallel: ~50ms total

	start := time.Now()

	f1 := FutureDo(func() (int, error) {
		time.Sleep(50 * time.Millisecond)
		return 1, nil
	})
	f2 := FutureDo(func() (string, error) {
		time.Sleep(50 * time.Millisecond)
		return "two", nil
	})
	f3 := FutureDo(func() (bool, error) {
		time.Sleep(50 * time.Millisecond)
		return true, nil
	})
	f4 := FutureDo(func() (float64, error) {
		time.Sleep(50 * time.Millisecond)
		return 4.0, nil
	})
	f5 := FutureDo(func() ([]int, error) {
		time.Sleep(50 * time.Millisecond)
		return []int{5}, nil
	})

	joined := FutureJoin5(f1, f2, f3, f4, f5)

	ctx := context.Background()
	_, err := joined.Get(ctx)

	duration := time.Since(start)

	if err != nil {
		t.Errorf("Get() returned error: %v", err)
	}

	// Should complete in ~50ms (parallel), not 250ms (sequential)
	if duration > 100*time.Millisecond {
		t.Errorf("parallel execution took %v, expected ~50ms", duration)
	}
}

func TestFutureJoin6_Success(t *testing.T) {
	f1 := FutureDo(func() (int, error) {
		return 1, nil
	})
	f2 := FutureDo(func() (string, error) {
		return "two", nil
	})
	f3 := FutureDo(func() (bool, error) {
		return true, nil
	})
	f4 := FutureDo(func() (float64, error) {
		return 4.0, nil
	})
	f5 := FutureDo(func() ([]int, error) {
		return []int{5}, nil
	})
	f6 := FutureDo(func() (rune, error) {
		return 'a', nil
	})

	joined := FutureJoin6(f1, f2, f3, f4, f5, f6)

	ctx := context.Background()
	result, err := joined.Get(ctx)

	if err != nil {
		t.Errorf("Get() returned error: %v", err)
	}
	if result.V1 != 1 {
		t.Errorf("result.V1 = %v, want 1", result.V1)
	}
	if result.V2 != "two" {
		t.Errorf("result.V2 = %v, want two", result.V2)
	}
	if result.V3 != true {
		t.Errorf("result.V3 = %v, want true", result.V3)
	}
	if result.V4 != 4.0 {
		t.Errorf("result.V4 = %v, want 4.0", result.V4)
	}
	if len(result.V5) != 1 || result.V5[0] != 5 {
		t.Errorf("result.V5 = %v, want [5]", result.V5)
	}
	if result.V6 != 'a' {
		t.Errorf("result.V6 = %v, want 'a'", result.V6)
	}
}

func TestFutureJoin6_Error(t *testing.T) {
	expectedErr := errors.New("error in sixth")

	f1 := FutureDo(func() (int, error) {
		return 1, nil
	})
	f2 := FutureDo(func() (string, error) {
		return "two", nil
	})
	f3 := FutureDo(func() (bool, error) {
		return true, nil
	})
	f4 := FutureDo(func() (float64, error) {
		return 4.0, nil
	})
	f5 := FutureDo(func() ([]int, error) {
		return []int{5}, nil
	})
	f6 := FutureDo(func() (rune, error) {
		return 0, expectedErr
	})

	joined := FutureJoin6(f1, f2, f3, f4, f5, f6)

	ctx := context.Background()
	_, err := joined.Get(ctx)

	if err != expectedErr {
		t.Errorf("Get() error = %v, want %v", err, expectedErr)
	}
}

func TestFutureJoin7_Success(t *testing.T) {
	f1 := FutureDo(func() (int, error) {
		return 1, nil
	})
	f2 := FutureDo(func() (string, error) {
		return "two", nil
	})
	f3 := FutureDo(func() (bool, error) {
		return true, nil
	})
	f4 := FutureDo(func() (float64, error) {
		return 4.0, nil
	})
	f5 := FutureDo(func() ([]int, error) {
		return []int{5}, nil
	})
	f6 := FutureDo(func() (rune, error) {
		return 'a', nil
	})
	f7 := FutureDo(func() (byte, error) {
		return 7, nil
	})

	joined := FutureJoin7(f1, f2, f3, f4, f5, f6, f7)

	ctx := context.Background()
	result, err := joined.Get(ctx)

	if err != nil {
		t.Errorf("Get() returned error: %v", err)
	}
	if result.V1 != 1 {
		t.Errorf("result.V1 = %v, want 1", result.V1)
	}
	if result.V2 != "two" {
		t.Errorf("result.V2 = %v, want two", result.V2)
	}
	if result.V3 != true {
		t.Errorf("result.V3 = %v, want true", result.V3)
	}
	if result.V4 != 4.0 {
		t.Errorf("result.V4 = %v, want 4.0", result.V4)
	}
	if len(result.V5) != 1 || result.V5[0] != 5 {
		t.Errorf("result.V5 = %v, want [5]", result.V5)
	}
	if result.V6 != 'a' {
		t.Errorf("result.V6 = %v, want 'a'", result.V6)
	}
	if result.V7 != 7 {
		t.Errorf("result.V7 = %v, want 7", result.V7)
	}
}

func TestFutureJoin7_Error(t *testing.T) {
	expectedErr := errors.New("error in seventh")

	f1 := FutureDo(func() (int, error) {
		return 1, nil
	})
	f2 := FutureDo(func() (string, error) {
		return "two", nil
	})
	f3 := FutureDo(func() (bool, error) {
		return true, nil
	})
	f4 := FutureDo(func() (float64, error) {
		return 4.0, nil
	})
	f5 := FutureDo(func() ([]int, error) {
		return []int{5}, nil
	})
	f6 := FutureDo(func() (rune, error) {
		return 'a', nil
	})
	f7 := FutureDo(func() (byte, error) {
		return 0, expectedErr
	})

	joined := FutureJoin7(f1, f2, f3, f4, f5, f6, f7)

	ctx := context.Background()
	_, err := joined.Get(ctx)

	if err != expectedErr {
		t.Errorf("Get() error = %v, want %v", err, expectedErr)
	}
}

func TestFutureJoin8_Success(t *testing.T) {
	f1 := FutureDo(func() (int, error) {
		return 1, nil
	})
	f2 := FutureDo(func() (string, error) {
		return "two", nil
	})
	f3 := FutureDo(func() (bool, error) {
		return true, nil
	})
	f4 := FutureDo(func() (float64, error) {
		return 4.0, nil
	})
	f5 := FutureDo(func() ([]int, error) {
		return []int{5}, nil
	})
	f6 := FutureDo(func() (rune, error) {
		return 'a', nil
	})
	f7 := FutureDo(func() (byte, error) {
		return 7, nil
	})
	f8 := FutureDo(func() (uint, error) {
		return 8, nil
	})

	joined := FutureJoin8(f1, f2, f3, f4, f5, f6, f7, f8)

	ctx := context.Background()
	result, err := joined.Get(ctx)

	if err != nil {
		t.Errorf("Get() returned error: %v", err)
	}
	if result.V1 != 1 {
		t.Errorf("result.V1 = %v, want 1", result.V1)
	}
	if result.V2 != "two" {
		t.Errorf("result.V2 = %v, want two", result.V2)
	}
	if result.V3 != true {
		t.Errorf("result.V3 = %v, want true", result.V3)
	}
	if result.V4 != 4.0 {
		t.Errorf("result.V4 = %v, want 4.0", result.V4)
	}
	if len(result.V5) != 1 || result.V5[0] != 5 {
		t.Errorf("result.V5 = %v, want [5]", result.V5)
	}
	if result.V6 != 'a' {
		t.Errorf("result.V6 = %v, want 'a'", result.V6)
	}
	if result.V7 != 7 {
		t.Errorf("result.V7 = %v, want 7", result.V7)
	}
	if result.V8 != 8 {
		t.Errorf("result.V8 = %v, want 8", result.V8)
	}
}

func TestFutureJoin8_Error(t *testing.T) {
	expectedErr := errors.New("error in eighth")

	f1 := FutureDo(func() (int, error) {
		return 1, nil
	})
	f2 := FutureDo(func() (string, error) {
		return "two", nil
	})
	f3 := FutureDo(func() (bool, error) {
		return true, nil
	})
	f4 := FutureDo(func() (float64, error) {
		return 4.0, nil
	})
	f5 := FutureDo(func() ([]int, error) {
		return []int{5}, nil
	})
	f6 := FutureDo(func() (rune, error) {
		return 'a', nil
	})
	f7 := FutureDo(func() (byte, error) {
		return 7, nil
	})
	f8 := FutureDo(func() (uint, error) {
		return 0, expectedErr
	})

	joined := FutureJoin8(f1, f2, f3, f4, f5, f6, f7, f8)

	ctx := context.Background()
	_, err := joined.Get(ctx)

	if err != expectedErr {
		t.Errorf("Get() error = %v, want %v", err, expectedErr)
	}
}

// ========================================
// FutureAny Tests
// ========================================

func TestFutureAny_SuccessEarly(t *testing.T) {
	f1 := FutureDo(func() (int, error) {
		// slow success
		time.Sleep(50 * time.Millisecond)
		return 1, nil
	})
	f2 := FutureDo(func() (int, error) {
		// fast success
		return 2, nil
	})
	f3 := FutureDo(func() (int, error) {
		// error
		return 0, errors.New("failed")
	})

	any := FutureAny(f1, f2, f3)
	ctx := context.Background()
	start := time.Now()
	val, err := any.Get(ctx)
	dur := time.Since(start)

	if err != nil {
		t.Fatalf("Get() returned unexpected error: %v", err)
	}
	if val != 2 {
		t.Fatalf("Get() = %v, want 2", val)
	}
	// Should return quickly because f2 completes immediately
	if dur > 20*time.Millisecond {
		t.Fatalf("FutureAny took too long: %v", dur)
	}
}

func TestFutureAny_SuccessLater(t *testing.T) {
	f1 := FutureDo(func() (int, error) {
		return 0, errors.New("err1")
	})
	f2 := FutureDo(func() (int, error) {
		time.Sleep(30 * time.Millisecond)
		return 0, errors.New("err2")
	})
	f3 := FutureDo(func() (int, error) {
		time.Sleep(50 * time.Millisecond)
		return 3, nil
	})

	any := FutureAny(f1, f2, f3)
	ctx := context.Background()
	val, err := any.Get(ctx)

	if err != nil {
		t.Fatalf("Get() returned unexpected error: %v", err)
	}
	if val != 3 {
		t.Fatalf("Get() = %v, want 3", val)
	}
}

func TestFutureAny_AllError(t *testing.T) {
	err1 := errors.New("e1")
	err2 := errors.New("e2")

	f1 := FutureDo(func() (int, error) { return 0, err1 })
	f2 := FutureDo(func() (int, error) { return 0, err2 })

	any := FutureAny(f1, f2)
	ctx := context.Background()
	_, err := any.Get(ctx)

	if err == nil {
		t.Fatalf("Get() expected error, got nil")
	}
	// Expect one of the provided errors (first encountered)
	if err != err1 && err != err2 {
		t.Fatalf("Get() error = %v, want %v or %v", err, err1, err2)
	}
}

func TestFutureAny_Empty(t *testing.T) {
	any := FutureAny[int]()
	ctx := context.Background()
	_, err := any.Get(ctx)
	if err == nil {
		t.Fatalf("Get() expected error for empty futures, got nil")
	}
}
