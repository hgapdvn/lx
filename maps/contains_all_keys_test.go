package lxmaps_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/maps"
)

func TestContainsAllKeys_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		keys     []string
		expected bool
	}{
		{
			name:     "nil map with single key",
			input:    nil,
			keys:     []string{"a"},
			expected: false,
		},
		{
			name:     "nil map with multiple keys",
			input:    nil,
			keys:     []string{"a", "b", "c"},
			expected: false,
		},
		{
			name:     "empty map with single key",
			input:    map[string]int{},
			keys:     []string{"a"},
			expected: false,
		},
		{
			name:     "empty map with multiple keys",
			input:    map[string]int{},
			keys:     []string{"a", "b", "c"},
			expected: false,
		},
		{
			name:     "single entry no keys match",
			input:    map[string]int{"a": 1},
			keys:     []string{"b"},
			expected: false,
		},
		{
			name:     "single entry single key matches",
			input:    map[string]int{"a": 1},
			keys:     []string{"a"},
			expected: true,
		},
		{
			name:     "single entry multiple keys only one matches",
			input:    map[string]int{"a": 1},
			keys:     []string{"a", "b"},
			expected: false,
		},
		{
			name:     "multiple entries no keys match",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			keys:     []string{"d", "e", "f"},
			expected: false,
		},
		{
			name:     "multiple entries first key matches only",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			keys:     []string{"a", "d", "e"},
			expected: false,
		},
		{
			name:     "multiple entries all keys match",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			keys:     []string{"a", "b", "c"},
			expected: true,
		},
		{
			name:     "map has more keys than checking",
			input:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			keys:     []string{"a", "c"},
			expected: true,
		},
		{
			name:     "map has more keys but missing one check key",
			input:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			keys:     []string{"a", "c", "e"},
			expected: false,
		},
		{
			name:     "empty keys list",
			input:    map[string]int{"a": 1, "b": 2},
			keys:     []string{},
			expected: true,
		},
		{
			name:     "duplicate keys in input list all present",
			input:    map[string]int{"a": 1, "b": 2},
			keys:     []string{"a", "a", "a"},
			expected: true,
		},
		{
			name:     "duplicate keys in input list with missing key",
			input:    map[string]int{"a": 1, "b": 2},
			keys:     []string{"a", "c", "c"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.ContainsAllKeys(tt.input, tt.keys...)
			if result != tt.expected {
				t.Errorf("ContainsAllKeys() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestContainsAllKeys_IntString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		keys     []int
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			keys:     []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "empty map",
			input:    map[int]string{},
			keys:     []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "no keys match",
			input:    map[int]string{1: "a", 2: "b", 3: "c"},
			keys:     []int{4, 5, 6},
			expected: false,
		},
		{
			name:     "first key matches only",
			input:    map[int]string{1: "a", 2: "b", 3: "c"},
			keys:     []int{1, 4, 5},
			expected: false,
		},
		{
			name:     "all keys match",
			input:    map[int]string{1: "a", 2: "b", 3: "c"},
			keys:     []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "subset of keys match",
			input:    map[int]string{1: "a", 2: "b", 3: "c"},
			keys:     []int{1, 2},
			expected: true,
		},
		{
			name:     "single key matches",
			input:    map[int]string{1: "a"},
			keys:     []int{1},
			expected: true,
		},
		{
			name:     "single key does not match",
			input:    map[int]string{1: "a"},
			keys:     []int{2},
			expected: false,
		},
		{
			name:     "empty keys list",
			input:    map[int]string{1: "a", 2: "b"},
			keys:     []int{},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.ContainsAllKeys(tt.input, tt.keys...)
			if result != tt.expected {
				t.Errorf("ContainsAllKeys() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestContainsAllKeys_BoolFloat64(t *testing.T) {
	tests := []struct {
		name     string
		input    map[bool]float64
		keys     []bool
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			keys:     []bool{true, false},
			expected: false,
		},
		{
			name:     "empty map",
			input:    map[bool]float64{},
			keys:     []bool{true},
			expected: false,
		},
		{
			name:     "true key exists only",
			input:    map[bool]float64{true: 1.5},
			keys:     []bool{true},
			expected: true,
		},
		{
			name:     "false key exists only",
			input:    map[bool]float64{false: 2.5},
			keys:     []bool{false},
			expected: true,
		},
		{
			name:     "both keys exist, checking both",
			input:    map[bool]float64{true: 1.5, false: 2.5},
			keys:     []bool{true, false},
			expected: true,
		},
		{
			name:     "both keys exist, checking true only",
			input:    map[bool]float64{true: 1.5, false: 2.5},
			keys:     []bool{true},
			expected: true,
		},
		{
			name:     "only true exists, checking false",
			input:    map[bool]float64{true: 1.5},
			keys:     []bool{false},
			expected: false,
		},
		{
			name:     "only false exists, checking true",
			input:    map[bool]float64{false: 2.5},
			keys:     []bool{true},
			expected: false,
		},
		{
			name:     "empty keys list",
			input:    map[bool]float64{true: 1.5, false: 2.5},
			keys:     []bool{},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.ContainsAllKeys(tt.input, tt.keys...)
			if result != tt.expected {
				t.Errorf("ContainsAllKeys() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// BenchmarkContainsAllKeys benchmarks the ContainsAllKeys function
func BenchmarkContainsAllKeys(b *testing.B) {
	m := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5,
		"f": 6, "g": 7, "h": 8, "i": 9, "j": 10,
	}
	keys := []string{"b", "d", "f"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lxmaps.ContainsAllKeys(m, keys...)
	}
}
