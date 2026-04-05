package lxtypes

import (
	"sync"
	"sync/atomic"
	"testing"
)

// ========================================
// NewRef Tests
// ========================================

func TestNewRef(t *testing.T) {
	tests := []struct {
		name     string
		expected int
	}{
		{name: "positive integer", expected: 42},
		{name: "zero", expected: 0},
		{name: "negative integer", expected: -10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRef(tt.expected)
			if got := r.Get(); got != tt.expected {
				t.Errorf("NewRef(%v).Get() = %v, want %v", tt.expected, got, tt.expected)
			}
		})
	}
}

// ========================================
// Ref.Set Tests
// ========================================

func TestRef_Set(t *testing.T) {
	t.Run("replace integer value", func(t *testing.T) {
		r := NewRef(1)
		r.Set(99)
		if got := r.Get(); got != 99 {
			t.Errorf("After Set(99), Get() = %v, want 99", got)
		}
	})

	t.Run("replace string value", func(t *testing.T) {
		r := NewRef("hello")
		r.Set("world")
		if got := r.Get(); got != "world" {
			t.Errorf("After Set(\"world\"), Get() = %v, want \"world\"", got)
		}
	})

	t.Run("set same value", func(t *testing.T) {
		r := NewRef(42)
		r.Set(42)
		if got := r.Get(); got != 42 {
			t.Errorf("After Set(42) on existing 42, Get() = %v, want 42", got)
		}
	})

	t.Run("multiple sets", func(t *testing.T) {
		r := NewRef(0)
		for i := 1; i <= 5; i++ {
			r.Set(i)
		}
		if got := r.Get(); got != 5 {
			t.Errorf("After multiple sets, Get() = %v, want 5", got)
		}
	})
}

// ========================================
// Ref.Update Tests
// ========================================

func TestRef_Update(t *testing.T) {
	t.Run("increment integer", func(t *testing.T) {
		r := NewRef(0)
		r.Update(func(v int) int { return v + 1 })
		if got := r.Get(); got != 1 {
			t.Errorf("After Update(+1), Get() = %v, want 1", got)
		}
	})

	t.Run("transform string", func(t *testing.T) {
		r := NewRef("hello")
		r.Update(func(v string) string { return v + " world" })
		if got := r.Get(); got != "hello world" {
			t.Errorf("After Update, Get() = %v, want \"hello world\"", got)
		}
	})

	t.Run("multiple updates", func(t *testing.T) {
		r := NewRef(0)
		for i := 0; i < 5; i++ {
			r.Update(func(v int) int { return v + 1 })
		}
		if got := r.Get(); got != 5 {
			t.Errorf("After 5 increments, Get() = %v, want 5", got)
		}
	})

	t.Run("identity update", func(t *testing.T) {
		r := NewRef(42)
		r.Update(func(v int) int { return v })
		if got := r.Get(); got != 42 {
			t.Errorf("After identity Update, Get() = %v, want 42", got)
		}
	})
}

// ========================================
// Ref.ZeroValue Tests
// ========================================

func TestRef_ZeroValue(t *testing.T) {
	t.Run("int zero value", func(t *testing.T) {
		r := NewRef(0)
		if got := r.Get(); got != 0 {
			t.Errorf("NewRef(0).Get() = %v, want 0", got)
		}
	})

	t.Run("string zero value", func(t *testing.T) {
		r := NewRef("")
		if got := r.Get(); got != "" {
			t.Errorf("NewRef(\"\").Get() = %v, want \"\"", got)
		}
	})

	t.Run("bool zero value", func(t *testing.T) {
		r := NewRef(false)
		if got := r.Get(); got != false {
			t.Errorf("NewRef(false).Get() = %v, want false", got)
		}
	})

	t.Run("float64 zero value", func(t *testing.T) {
		r := NewRef(0.0)
		if got := r.Get(); got != 0.0 {
			t.Errorf("NewRef(0.0).Get() = %v, want 0.0", got)
		}
	})
}

// ========================================
// Ref with Different Types Tests
// ========================================

func TestRef_DifferentTypes(t *testing.T) {
	t.Run("string type", func(t *testing.T) {
		r := NewRef("initial")
		r.Set("updated")
		if got := r.Get(); got != "updated" {
			t.Errorf("string Ref.Get() = %v, want \"updated\"", got)
		}
	})

	t.Run("bool type", func(t *testing.T) {
		r := NewRef(false)
		r.Set(true)
		if got := r.Get(); !got {
			t.Errorf("bool Ref.Get() = %v, want true", got)
		}
	})

	t.Run("struct type", func(t *testing.T) {
		type Config struct {
			Host string
			Port int
		}
		r := NewRef(Config{Host: "localhost", Port: 8080})
		r.Set(Config{Host: "prod.example.com", Port: 443})
		got := r.Get()
		if got.Host != "prod.example.com" || got.Port != 443 {
			t.Errorf("struct Ref.Get() = %v, want {prod.example.com 443}", got)
		}
	})

	t.Run("slice type", func(t *testing.T) {
		r := NewRef([]int{1, 2, 3})
		r.Update(func(v []int) []int {
			// Create a new slice to avoid aliasing; append may return a new backing array
			result := make([]int, len(v), len(v)+1)
			copy(result, v)
			return append(result, 4)
		})
		got := r.Get()
		if len(got) != 4 || got[3] != 4 {
			t.Errorf("slice Ref.Get() = %v, want [1 2 3 4]", got)
		}
	})

	t.Run("map type", func(t *testing.T) {
		r := NewRef(map[string]int{"a": 1})
		r.Update(func(v map[string]int) map[string]int {
			// Return a new map to avoid aliasing issues
			newMap := make(map[string]int, len(v)+1)
			for k, val := range v {
				newMap[k] = val
			}
			newMap["b"] = 2
			return newMap
		})
		got := r.Get()
		if got["a"] != 1 || got["b"] != 2 {
			t.Errorf("map Ref.Get() = %v, want map[a:1 b:2]", got)
		}
	})
}

// ========================================
// Concurrent Tests
// ========================================

func TestRef_Concurrent(t *testing.T) {
	const goroutines = 100

	r := NewRef(0)
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			r.Update(func(v int) int { return v + 1 })
		}()
	}

	wg.Wait()

	if got := r.Get(); got != goroutines {
		t.Errorf("After %d concurrent increments, Get() = %v, want %d", goroutines, got, goroutines)
	}
}

func TestRef_ConcurrentSetAndGet(t *testing.T) {
	r := NewRef(0)
	var wg sync.WaitGroup
	const goroutines = 50

	// Concurrent readers and writers
	wg.Add(goroutines * 2)
	for i := 0; i < goroutines; i++ {
		go func(val int) {
			defer wg.Done()
			r.Set(val)
		}(i)
		go func() {
			defer wg.Done()
			_ = r.Get() // just read, race detector will catch data races
		}()
	}

	wg.Wait()
	// No assertion on final value since writes are concurrent, but no race should occur
}

// ========================================
// Update Captures Old Value Tests
// ========================================

func TestRef_UpdateCapturesOldValue(t *testing.T) {
	r := NewRef(10)
	var capturedOld int

	r.Update(func(v int) int {
		capturedOld = v // capture old value before transforming
		return v * 2
	})

	if capturedOld != 10 {
		t.Errorf("Update closure captured old value = %v, want 10", capturedOld)
	}
	if got := r.Get(); got != 20 {
		t.Errorf("After Update(*2), Get() = %v, want 20", got)
	}
}

func TestRef_UpdateCapturesOldValue_String(t *testing.T) {
	r := NewRef("hello")
	var capturedOld string

	r.Update(func(v string) string {
		capturedOld = v
		return v + " world"
	})

	if capturedOld != "hello" {
		t.Errorf("Update closure captured old value = %v, want \"hello\"", capturedOld)
	}
	if got := r.Get(); got != "hello world" {
		t.Errorf("After Update, Get() = %v, want \"hello world\"", got)
	}
}

// ========================================
// Atomic Counter Tests
// ========================================

func TestRef_AtomicCounter(t *testing.T) {
	// Verify using atomic operations that concurrent Update is race-free
	var atomicCount int64
	r := NewRef(0)

	const goroutines = 100
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			r.Update(func(v int) int {
				atomic.AddInt64(&atomicCount, 1)
				return v + 1
			})
		}()
	}

	wg.Wait()

	if got := r.Get(); got != goroutines {
		t.Errorf("Ref counter = %v, want %d", got, goroutines)
	}
	if atomicCount != goroutines {
		t.Errorf("atomic counter = %v, want %d", atomicCount, goroutines)
	}
}
