package lxmaps_test

import (
	"testing"

	"github.com/hgapdvn/lx/lxmaps"
)

func TestContainsAllValues_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		values   []int
		expected bool
	}{
		{
			name:     "nil map with single value",
			input:    nil,
			values:   []int{1},
			expected: false,
		},
		{
			name:     "nil map with multiple values",
			input:    nil,
			values:   []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "empty map with single value",
			input:    map[string]int{},
			values:   []int{1},
			expected: false,
		},
		{
			name:     "empty map with multiple values",
			input:    map[string]int{},
			values:   []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "single entry no values match",
			input:    map[string]int{"a": 1},
			values:   []int{2},
			expected: false,
		},
		{
			name:     "single entry single value matches",
			input:    map[string]int{"a": 1},
			values:   []int{1},
			expected: true,
		},
		{
			name:     "single entry multiple values only one matches",
			input:    map[string]int{"a": 1},
			values:   []int{1, 2},
			expected: false,
		},
		{
			name:     "multiple entries no values match",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			values:   []int{4, 5, 6},
			expected: false,
		},
		{
			name:     "multiple entries first value matches only",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			values:   []int{1, 4, 5},
			expected: false,
		},
		{
			name:     "multiple entries all values match",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			values:   []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "map has more values than checking",
			input:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			values:   []int{1, 3},
			expected: true,
		},
		{
			name:     "map has more values but missing one check value",
			input:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			values:   []int{1, 3, 5},
			expected: false,
		},
		{
			name:     "empty values list",
			input:    map[string]int{"a": 1, "b": 2},
			values:   []int{},
			expected: true,
		},
		{
			name:     "duplicate values in input list all present",
			input:    map[string]int{"a": 1, "b": 2},
			values:   []int{1, 1, 1},
			expected: true,
		},
		{
			name:     "duplicate values in input list with missing value",
			input:    map[string]int{"a": 1, "b": 2},
			values:   []int{1, 3, 3},
			expected: false,
		},
		{
			name:     "zero value exists",
			input:    map[string]int{"a": 0, "b": 1},
			values:   []int{0},
			expected: true,
		},
		{
			name:     "negative value exists",
			input:    map[string]int{"a": -1, "b": -2},
			values:   []int{-1},
			expected: true,
		},
		{
			name:     "map with duplicate values checking all present",
			input:    map[string]int{"a": 5, "b": 5, "c": 5},
			values:   []int{5},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.ContainsAllValues(tt.input, tt.values...)
			if result != tt.expected {
				t.Errorf("ContainsAllValues() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestContainsAllValues_IntString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		values   []string
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			values:   []string{"a", "b", "c"},
			expected: false,
		},
		{
			name:     "empty map",
			input:    map[int]string{},
			values:   []string{"a", "b", "c"},
			expected: false,
		},
		{
			name:     "no values match",
			input:    map[int]string{1: "a", 2: "b", 3: "c"},
			values:   []string{"d", "e", "f"},
			expected: false,
		},
		{
			name:     "first value matches only",
			input:    map[int]string{1: "a", 2: "b", 3: "c"},
			values:   []string{"a", "d", "e"},
			expected: false,
		},
		{
			name:     "all values match",
			input:    map[int]string{1: "a", 2: "b", 3: "c"},
			values:   []string{"a", "b", "c"},
			expected: true,
		},
		{
			name:     "subset of values match",
			input:    map[int]string{1: "a", 2: "b", 3: "c"},
			values:   []string{"a", "b"},
			expected: true,
		},
		{
			name:     "single value matches",
			input:    map[int]string{1: "a"},
			values:   []string{"a"},
			expected: true,
		},
		{
			name:     "single value does not match",
			input:    map[int]string{1: "a"},
			values:   []string{"b"},
			expected: false,
		},
		{
			name:     "empty values list",
			input:    map[int]string{1: "a", 2: "b"},
			values:   []string{},
			expected: true,
		},
		{
			name:     "empty string value exists",
			input:    map[int]string{1: "", 2: "a"},
			values:   []string{""},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.ContainsAllValues(tt.input, tt.values...)
			if result != tt.expected {
				t.Errorf("ContainsAllValues() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestContainsAllValues_StringBool(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]bool
		values   []bool
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			values:   []bool{true, false},
			expected: false,
		},
		{
			name:     "empty map",
			input:    map[string]bool{},
			values:   []bool{true},
			expected: false,
		},
		{
			name:     "true value exists",
			input:    map[string]bool{"a": true, "b": false},
			values:   []bool{true},
			expected: true,
		},
		{
			name:     "false value exists",
			input:    map[string]bool{"a": true, "b": false},
			values:   []bool{false},
			expected: true,
		},
		{
			name:     "both values exist, checking both",
			input:    map[string]bool{"a": true, "b": false},
			values:   []bool{true, false},
			expected: true,
		},
		{
			name:     "both values exist, checking true only",
			input:    map[string]bool{"a": true, "b": false},
			values:   []bool{true},
			expected: true,
		},
		{
			name:     "only true exists, checking false",
			input:    map[string]bool{"a": true, "c": true},
			values:   []bool{false},
			expected: false,
		},
		{
			name:     "only false exists, checking true",
			input:    map[string]bool{"a": false, "b": false},
			values:   []bool{true},
			expected: false,
		},
		{
			name:     "empty values list",
			input:    map[string]bool{"a": true, "b": false},
			values:   []bool{},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.ContainsAllValues(tt.input, tt.values...)
			if result != tt.expected {
				t.Errorf("ContainsAllValues() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestContainsAllValues_StringString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]string
		values   []string
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			values:   []string{"hello", "world"},
			expected: false,
		},
		{
			name:     "empty map",
			input:    map[string]string{},
			values:   []string{"hello", "world"},
			expected: false,
		},
		{
			name:     "no values match",
			input:    map[string]string{"a": "apple", "b": "banana"},
			values:   []string{"cherry", "date"},
			expected: false,
		},
		{
			name:     "one value matches only",
			input:    map[string]string{"a": "apple", "b": "banana"},
			values:   []string{"apple", "cherry"},
			expected: false,
		},
		{
			name:     "multiple values match",
			input:    map[string]string{"a": "apple", "b": "banana", "c": "cherry"},
			values:   []string{"apple", "banana", "cherry"},
			expected: true,
		},
		{
			name:     "subset of values match",
			input:    map[string]string{"a": "apple", "b": "banana", "c": "cherry"},
			values:   []string{"apple", "banana"},
			expected: true,
		},
		{
			name:     "case sensitive no match",
			input:    map[string]string{"a": "Apple"},
			values:   []string{"apple"},
			expected: false,
		},
		{
			name:     "unicode values all present",
			input:    map[string]string{"a": "こんにちは", "b": "さようなら"},
			values:   []string{"こんにちは"},
			expected: true,
		},
		{
			name:     "unicode values not all present",
			input:    map[string]string{"a": "こんにちは"},
			values:   []string{"こんにちは", "さようなら"},
			expected: false,
		},
		{
			name:     "empty values list",
			input:    map[string]string{"a": "apple", "b": "banana"},
			values:   []string{},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.ContainsAllValues(tt.input, tt.values...)
			if result != tt.expected {
				t.Errorf("ContainsAllValues() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// BenchmarkContainsAllValues benchmarks the ContainsAllValues function
func BenchmarkContainsAllValues(b *testing.B) {
	m := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5,
		"f": 6, "g": 7, "h": 8, "i": 9, "j": 10,
	}
	values := []int{2, 4, 6}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lxmaps.ContainsAllValues(m, values...)
	}
}
