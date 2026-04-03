package lxmaps_test

import (
	"testing"

	lxmaps "github.com/hgapdvn/lx/maps"
)

func TestReduce_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		fn       func(int, string, int) int
		initial  int
		expected int
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(acc int, k string, v int) int { return acc + v },
			initial:  100,
			expected: 100,
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			fn:       func(acc int, k string, v int) int { return acc + v },
			initial:  50,
			expected: 50,
		},
		{
			name:     "single entry",
			input:    map[string]int{"a": 10},
			fn:       func(acc int, k string, v int) int { return acc + v },
			initial:  0,
			expected: 10,
		},
		{
			name:     "sum values",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			fn:       func(acc int, k string, v int) int { return acc + v },
			initial:  0,
			expected: 6,
		},
		{
			name:     "product values",
			input:    map[string]int{"a": 2, "b": 3, "c": 4},
			fn:       func(acc int, k string, v int) int { return acc * v },
			initial:  1,
			expected: 24,
		},
		{
			name:     "count entries",
			input:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			fn:       func(acc int, k string, v int) int { return acc + 1 },
			initial:  0,
			expected: 4,
		},
		{
			name:  "filter sum even only",
			input: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			fn: func(acc int, k string, v int) int {
				if v%2 == 0 {
					return acc + v
				}
				return acc
			},
			initial:  0,
			expected: 6,
		},
		{
			name:     "negative values",
			input:    map[string]int{"a": -10, "b": -20, "c": -30},
			fn:       func(acc int, k string, v int) int { return acc + v },
			initial:  0,
			expected: -60,
		},
		{
			name:     "zero values",
			input:    map[string]int{"a": 0, "b": 0, "c": 0},
			fn:       func(acc int, k string, v int) int { return acc + v },
			initial:  0,
			expected: 0,
		},
		{
			name:     "sum using both key and value",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			fn:       func(acc int, k string, v int) int { return acc + len(k)*v },
			initial:  0,
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Reduce(tt.input, tt.fn, tt.initial)
			if result != tt.expected {
				t.Errorf("Reduce() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestReduce_IntString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		fn       func(string, int, string) string
		initial  string
		expected string
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(acc string, k int, v string) string { return acc + v },
			initial:  "start",
			expected: "start",
		},
		{
			name:     "empty map",
			input:    map[int]string{},
			fn:       func(acc string, k int, v string) string { return acc + v },
			initial:  "empty",
			expected: "empty",
		},
		{
			name:     "single entry concatenate",
			input:    map[int]string{1: "one"},
			fn:       func(acc string, k int, v string) string { return acc + v },
			initial:  "",
			expected: "one",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Reduce(tt.input, tt.fn, tt.initial)
			if result != tt.expected {
				t.Errorf("Reduce() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestReduce_StringBool(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]bool
		fn       func(bool, string, bool) bool
		initial  bool
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(acc bool, k string, v bool) bool { return acc && v },
			initial:  true,
			expected: true,
		},
		{
			name:     "empty map",
			input:    map[string]bool{},
			fn:       func(acc bool, k string, v bool) bool { return acc && v },
			initial:  true,
			expected: true,
		},
		{
			name:     "all true AND logic",
			input:    map[string]bool{"a": true, "b": true, "c": true},
			fn:       func(acc bool, k string, v bool) bool { return acc && v },
			initial:  true,
			expected: true,
		},
		{
			name:     "one false AND logic",
			input:    map[string]bool{"a": true, "b": false, "c": true},
			fn:       func(acc bool, k string, v bool) bool { return acc && v },
			initial:  true,
			expected: false,
		},
		{
			name:     "all false OR logic",
			input:    map[string]bool{"a": false, "b": false, "c": false},
			fn:       func(acc bool, k string, v bool) bool { return acc || v },
			initial:  false,
			expected: false,
		},
		{
			name:     "one true OR logic",
			input:    map[string]bool{"a": false, "b": true, "c": false},
			fn:       func(acc bool, k string, v bool) bool { return acc || v },
			initial:  false,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Reduce(tt.input, tt.fn, tt.initial)
			if result != tt.expected {
				t.Errorf("Reduce() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestReduce_StringCustomStruct(t *testing.T) {
	type Item struct {
		Name  string
		Value int
	}

	tests := []struct {
		name     string
		input    map[string]Item
		fn       func(int, string, Item) int
		initial  int
		expected int
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(acc int, k string, v Item) int { return acc + v.Value },
			initial:  0,
			expected: 0,
		},
		{
			name:     "empty map",
			input:    map[string]Item{},
			fn:       func(acc int, k string, v Item) int { return acc + v.Value },
			initial:  100,
			expected: 100,
		},
		{
			name: "sum struct values",
			input: map[string]Item{
				"item1": {Name: "A", Value: 10},
				"item2": {Name: "B", Value: 20},
				"item3": {Name: "C", Value: 30},
			},
			fn:       func(acc int, k string, v Item) int { return acc + v.Value },
			initial:  0,
			expected: 60,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Reduce(tt.input, tt.fn, tt.initial)
			if result != tt.expected {
				t.Errorf("Reduce() = %d, want %d", result, tt.expected)
			}
		})
	}
}
