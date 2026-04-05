package lxmaps_test

import (
	"fmt"
	"testing"

	"github.com/hgapdvn/lx/lxmaps"
)

func TestDifference_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		m1       map[string]int
		m2       map[string]int
		wantNil  bool
		expected map[string]int
	}{
		{
			name:     "m1 nil",
			m1:       nil,
			m2:       map[string]int{"a": 1},
			wantNil:  true,
			expected: nil,
		},
		{
			name:     "m2 nil",
			m1:       map[string]int{"a": 1},
			m2:       nil,
			wantNil:  false,
			expected: map[string]int{"a": 1},
		},
		{
			name:     "both nil",
			m1:       nil,
			m2:       nil,
			wantNil:  true,
			expected: nil,
		},
		{
			name:     "both empty",
			m1:       map[string]int{},
			m2:       map[string]int{},
			wantNil:  false,
			expected: map[string]int{},
		},
		{
			name:     "m1 empty, m2 non-empty",
			m1:       map[string]int{},
			m2:       map[string]int{"a": 1},
			wantNil:  false,
			expected: map[string]int{},
		},
		{
			name:     "m1 non-empty, m2 empty",
			m1:       map[string]int{"a": 1, "b": 2},
			m2:       map[string]int{},
			wantNil:  false,
			expected: map[string]int{"a": 1, "b": 2},
		},
		{
			name:     "no common keys",
			m1:       map[string]int{"a": 1, "b": 2},
			m2:       map[string]int{"c": 3, "d": 4},
			wantNil:  false,
			expected: map[string]int{"a": 1, "b": 2},
		},
		{
			name:     "one key in m2",
			m1:       map[string]int{"a": 1, "b": 2, "c": 3},
			m2:       map[string]int{"b": 99},
			wantNil:  false,
			expected: map[string]int{"a": 1, "c": 3},
		},
		{
			name:     "all keys in m2",
			m1:       map[string]int{"a": 1, "b": 2, "c": 3},
			m2:       map[string]int{"a": 10, "b": 20, "c": 30},
			wantNil:  false,
			expected: map[string]int{},
		},
		{
			name:     "multiple keys in m2, m2 has extra",
			m1:       map[string]int{"a": 1, "b": 2, "c": 3},
			m2:       map[string]int{"a": 10, "b": 20, "d": 40},
			wantNil:  false,
			expected: map[string]int{"c": 3},
		},
		{
			name:     "preserves m1 values",
			m1:       map[string]int{"x": 100, "y": 200},
			m2:       map[string]int{"y": 999},
			wantNil:  false,
			expected: map[string]int{"x": 100},
		},
		{
			name:     "single entry, not in m2",
			m1:       map[string]int{"key": 42},
			m2:       map[string]int{"other": 88},
			wantNil:  false,
			expected: map[string]int{"key": 42},
		},
		{
			name:     "single entry, in m2",
			m1:       map[string]int{"key": 42},
			m2:       map[string]int{"key": 88},
			wantNil:  false,
			expected: map[string]int{},
		},
		{
			name:     "zero value in m1",
			m1:       map[string]int{"a": 0, "b": 1},
			m2:       map[string]int{"b": 1},
			wantNil:  false,
			expected: map[string]int{"a": 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Difference(tt.m1, tt.m2)

			if tt.wantNil {
				if result != nil {
					t.Errorf("expected nil, got %v", result)
				}
				return
			}

			if result == nil {
				t.Errorf("expected non-nil, got nil")
				return
			}

			if len(result) != len(tt.expected) {
				t.Errorf("length mismatch: got %d, expected %d", len(result), len(tt.expected))
			}

			for k, v := range tt.expected {
				if resultV, exists := result[k]; !exists {
					t.Errorf("key %q not found in result", k)
				} else if resultV != v {
					t.Errorf("value mismatch for key %q: got %d, expected %d", k, resultV, v)
				}
			}
		})
	}
}

func TestDifference_IntString(t *testing.T) {
	tests := []struct {
		name     string
		m1       map[int]string
		m2       map[int]string
		expected map[int]string
	}{
		{
			name:     "simple case",
			m1:       map[int]string{1: "a", 2: "b", 3: "c"},
			m2:       map[int]string{2: "x", 3: "y", 4: "z"},
			expected: map[int]string{1: "a"},
		},
		{
			name:     "negative keys",
			m1:       map[int]string{-1: "a", 0: "b", 1: "c"},
			m2:       map[int]string{-1: "x", 1: "y"},
			expected: map[int]string{0: "b"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Difference(tt.m1, tt.m2)

			if len(result) != len(tt.expected) {
				t.Errorf("length mismatch: got %d, expected %d", len(result), len(tt.expected))
			}

			for k, v := range tt.expected {
				if resultV, exists := result[k]; !exists {
					t.Errorf("key %d not found in result", k)
				} else if resultV != v {
					t.Errorf("value mismatch for key %d: got %q, expected %q", k, resultV, v)
				}
			}
		})
	}
}

func TestDifference_BoolFloat64(t *testing.T) {
	tests := []struct {
		name     string
		m1       map[bool]float64
		m2       map[bool]float64
		expected map[bool]float64
	}{
		{
			name:     "boolean keys with floats",
			m1:       map[bool]float64{true: 1.5, false: 2.5},
			m2:       map[bool]float64{true: 3.5},
			expected: map[bool]float64{false: 2.5},
		},
		{
			name:     "all keys in m2",
			m1:       map[bool]float64{true: 1.1, false: 2.2},
			m2:       map[bool]float64{true: 3.3, false: 4.4},
			expected: map[bool]float64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Difference(tt.m1, tt.m2)

			if len(result) != len(tt.expected) {
				t.Errorf("length mismatch: got %d, expected %d", len(result), len(tt.expected))
			}

			for k, v := range tt.expected {
				if resultV, exists := result[k]; !exists {
					t.Errorf("key %v not found in result", k)
				} else if resultV != v {
					t.Errorf("value mismatch for key %v: got %f, expected %f", k, resultV, v)
				}
			}
		})
	}
}

func TestDifference_DoesNotModifyInputs(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}

	originalM1 := make(map[string]int)
	for k, v := range m1 {
		originalM1[k] = v
	}
	originalM2 := make(map[string]int)
	for k, v := range m2 {
		originalM2[k] = v
	}

	_ = lxmaps.Difference(m1, m2)

	for k, v := range originalM1 {
		if m1[k] != v {
			t.Errorf("m1 was modified: key %q changed from %d to %d", k, v, m1[k])
		}
	}

	for k, v := range originalM2 {
		if m2[k] != v {
			t.Errorf("m2 was modified: key %q changed from %d to %d", k, v, m2[k])
		}
	}
}

func TestDifference_StructValue(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	tests := []struct {
		name     string
		m1       map[string]Person
		m2       map[string]Person
		expected map[string]Person
	}{
		{
			name: "struct values preserved from m1",
			m1: map[string]Person{
				"alice":   {Name: "Alice", Age: 30},
				"bob":     {Name: "Bob", Age: 25},
				"charlie": {Name: "Charlie", Age: 35},
			},
			m2: map[string]Person{
				"bob": {Name: "Bob", Age: 99},
				"eve": {Name: "Eve", Age: 28},
			},
			expected: map[string]Person{
				"alice":   {Name: "Alice", Age: 30},
				"charlie": {Name: "Charlie", Age: 35},
			},
		},
		{
			name: "empty struct fields",
			m1: map[string]Person{
				"x": {Name: "", Age: 0},
				"y": {Name: "Y", Age: 1},
			},
			m2: map[string]Person{
				"y": {Name: "Y", Age: 99},
			},
			expected: map[string]Person{
				"x": {Name: "", Age: 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Difference(tt.m1, tt.m2)
			if len(result) != len(tt.expected) {
				t.Errorf("length mismatch: got %d, expected %d", len(result), len(tt.expected))
			}
			for k, v := range tt.expected {
				if resultV, exists := result[k]; !exists {
					t.Errorf("key %q not found in result", k)
				} else if resultV != v {
					t.Errorf("value mismatch for key %q: got %+v, expected %+v", k, resultV, v)
				}
			}
		})
	}
}

func TestDifference_UnicodeKeys(t *testing.T) {
	tests := []struct {
		name     string
		m1       map[string]int
		m2       map[string]int
		expected map[string]int
	}{
		{
			name: "unicode characters",
			m1: map[string]int{
				"こんにちは":   1,
				"世界":      2,
				"English": 3,
			},
			m2: map[string]int{
				"こんにちは":  99,
				"другой": 77,
			},
			expected: map[string]int{
				"世界":      2,
				"English": 3,
			},
		},
		{
			name: "emoji keys",
			m1: map[string]int{
				"😊": 1,
				"🚀": 2,
				"🎉": 3,
			},
			m2: map[string]int{
				"🚀":  99,
				"❤️": 77,
			},
			expected: map[string]int{
				"😊": 1,
				"🎉": 3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Difference(tt.m1, tt.m2)
			if len(result) != len(tt.expected) {
				t.Errorf("length mismatch: got %d, expected %d", len(result), len(tt.expected))
			}
			for k, v := range tt.expected {
				if resultV, exists := result[k]; !exists {
					t.Errorf("key %q not found in result", k)
				} else if resultV != v {
					t.Errorf("value mismatch for key %q: got %d, expected %d", k, resultV, v)
				}
			}
		})
	}
}

func TestDifference_SpecialCharKeys(t *testing.T) {
	tests := []struct {
		name     string
		m1       map[string]int
		m2       map[string]int
		expected map[string]int
	}{
		{
			name: "special character keys",
			m1: map[string]int{
				"!@#$%": 1,
				"^&*()": 2,
				"<>?":   3,
			},
			m2: map[string]int{
				"^&*()": 99,
				"~`":    77,
			},
			expected: map[string]int{
				"!@#$%": 1,
				"<>?":   3,
			},
		},
		{
			name: "whitespace keys",
			m1: map[string]int{
				" ":       1,
				"\t":      2,
				"\n":      3,
				"  key  ": 4,
			},
			m2: map[string]int{
				"\n":      88,
				"  key  ": 77,
			},
			expected: map[string]int{
				" ":  1,
				"\t": 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Difference(tt.m1, tt.m2)
			if len(result) != len(tt.expected) {
				t.Errorf("length mismatch: got %d, expected %d", len(result), len(tt.expected))
			}
			for k, v := range tt.expected {
				if resultV, exists := result[k]; !exists {
					t.Errorf("key %q not found in result", k)
				} else if resultV != v {
					t.Errorf("value mismatch for key %q: got %d, expected %d", k, resultV, v)
				}
			}
		})
	}
}

func TestDifference_CaseSensitive(t *testing.T) {
	m1 := map[string]int{
		"Key": 1,
		"key": 2,
		"KEY": 3,
	}
	m2 := map[string]int{
		"key": 99,
	}

	result := lxmaps.Difference(m1, m2)

	expected := map[string]int{
		"Key": 1,
		"KEY": 3,
	}

	if len(result) != len(expected) {
		t.Errorf("length mismatch: got %d, expected %d", len(result), len(expected))
	}

	for k, v := range expected {
		if resultV, exists := result[k]; !exists {
			t.Errorf("key %q not found", k)
		} else if resultV != v {
			t.Errorf("value mismatch for key %q: got %d, expected %d", k, resultV, v)
		}
	}
}

func TestDifference_LargeMap(t *testing.T) {
	m1 := make(map[string]int)
	m2 := make(map[string]int)

	// Create large maps with 10000 entries
	for i := 0; i < 10000; i++ {
		m1[fmt.Sprintf("key_%d", i)] = i
		if i%3 == 0 {
			m2[fmt.Sprintf("key_%d", i)] = i * 2
		}
	}

	result := lxmaps.Difference(m1, m2)

	// Should have approximately 2/3 of entries (those not in m2)
	expectedCount := (10000 * 2) / 3
	if result == nil {
		t.Error("expected non-nil result")
	}
	// Allow some variance due to modulo distribution
	if len(result) < expectedCount-100 || len(result) > expectedCount+100 {
		t.Errorf("unexpected result size: got %d, expected ~%d", len(result), expectedCount)
	}

	// Verify values are from m1, not m2
	for k, v := range result {
		m1Val, _ := m1[k]
		if v != m1Val {
			t.Errorf("value not from m1 for key %q: got %d, m1 has %d", k, v, m1Val)
		}
		// Verify key is not in m2
		if _, inM2 := m2[k]; inM2 {
			t.Errorf("key %q should not be in m2", k)
		}
	}
}

func BenchmarkDifference(b *testing.B) {
	m1 := make(map[string]int)
	m2 := make(map[string]int)

	for i := 0; i < 1000; i++ {
		m1[fmt.Sprintf("key_%d", i)] = i
		if i%2 == 0 {
			m2[fmt.Sprintf("key_%d", i)] = i
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lxmaps.Difference(m1, m2)
	}
}
