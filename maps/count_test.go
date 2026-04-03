package lxmaps_test

import (
	"testing"

	lxmaps "github.com/hgapdvn/lx/maps"
)

func TestCount_StringInt(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]int
		predicate func(string, int) bool
		expected  int
	}{
		{
			name:      "nil map",
			input:     nil,
			predicate: func(k string, v int) bool { return true },
			expected:  0,
		},
		{
			name:      "empty map",
			input:     map[string]int{},
			predicate: func(k string, v int) bool { return true },
			expected:  0,
		},
		{
			name:      "single entry matches",
			input:     map[string]int{"a": 10},
			predicate: func(k string, v int) bool { return v > 5 },
			expected:  1,
		},
		{
			name:      "single entry no match",
			input:     map[string]int{"a": 10},
			predicate: func(k string, v int) bool { return v < 5 },
			expected:  0,
		},
		{
			name:      "count all entries",
			input:     map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(k string, v int) bool { return true },
			expected:  3,
		},
		{
			name:      "count no entries",
			input:     map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(k string, v int) bool { return false },
			expected:  0,
		},
		{
			name:      "count by value condition",
			input:     map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			predicate: func(k string, v int) bool { return v > 2 },
			expected:  2,
		},
		{
			name:      "count even values",
			input:     map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			predicate: func(k string, v int) bool { return v%2 == 0 },
			expected:  2,
		},
		{
			name:      "count by key length",
			input:     map[string]int{"x": 1, "abc": 2, "defg": 3},
			predicate: func(k string, v int) bool { return len(k) > 1 },
			expected:  2,
		},
		{
			name:      "count by both key and value",
			input:     map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			predicate: func(k string, v int) bool { return len(k) == 1 && v > 2 },
			expected:  2,
		},
		{
			name:      "count negative values",
			input:     map[string]int{"a": -5, "b": 10, "c": -3, "d": 0},
			predicate: func(k string, v int) bool { return v < 0 },
			expected:  2,
		},
		{
			name:      "count zero values",
			input:     map[string]int{"a": 0, "b": 0, "c": 1, "d": 2},
			predicate: func(k string, v int) bool { return v == 0 },
			expected:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Count(tt.input, tt.predicate)
			if result != tt.expected {
				t.Errorf("Count() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestCount_IntString(t *testing.T) {
	tests := []struct {
		name      string
		input     map[int]string
		predicate func(int, string) bool
		expected  int
	}{
		{
			name:      "nil map",
			input:     nil,
			predicate: func(k int, v string) bool { return true },
			expected:  0,
		},
		{
			name:      "empty map",
			input:     map[int]string{},
			predicate: func(k int, v string) bool { return true },
			expected:  0,
		},
		{
			name:      "count by key",
			input:     map[int]string{1: "a", 2: "b", 3: "c", 4: "d"},
			predicate: func(k int, v string) bool { return k > 2 },
			expected:  2,
		},
		{
			name:      "count by value length",
			input:     map[int]string{1: "a", 2: "ab", 3: "abc", 4: "abcd"},
			predicate: func(k int, v string) bool { return len(v) > 2 },
			expected:  2,
		},
		{
			name:      "count non-empty strings",
			input:     map[int]string{1: "hello", 2: "", 3: "world", 4: ""},
			predicate: func(k int, v string) bool { return v != "" },
			expected:  2,
		},
		{
			name:      "count by key and value length",
			input:     map[int]string{1: "a", 2: "bb", 3: "ccc", 4: "dddd"},
			predicate: func(k int, v string) bool { return k == len(v) },
			expected:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Count(tt.input, tt.predicate)
			if result != tt.expected {
				t.Errorf("Count() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestCount_StringBool(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]bool
		predicate func(string, bool) bool
		expected  int
	}{
		{
			name:      "nil map",
			input:     nil,
			predicate: func(k string, v bool) bool { return v },
			expected:  0,
		},
		{
			name:      "empty map",
			input:     map[string]bool{},
			predicate: func(k string, v bool) bool { return v },
			expected:  0,
		},
		{
			name:      "count true values",
			input:     map[string]bool{"a": true, "b": false, "c": true},
			predicate: func(k string, v bool) bool { return v },
			expected:  2,
		},
		{
			name:      "count false values",
			input:     map[string]bool{"a": true, "b": false, "c": true, "d": false},
			predicate: func(k string, v bool) bool { return !v },
			expected:  2,
		},
		{
			name:      "count by key and value",
			input:     map[string]bool{"a": true, "b": false, "c": true, "d": false},
			predicate: func(k string, v bool) bool { return len(k) == 1 && v },
			expected:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Count(tt.input, tt.predicate)
			if result != tt.expected {
				t.Errorf("Count() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestCount_StringCustomStruct(t *testing.T) {
	type Item struct {
		Name   string
		Value  int
		Active bool
	}

	tests := []struct {
		name      string
		input     map[string]Item
		predicate func(string, Item) bool
		expected  int
	}{
		{
			name:      "nil map",
			input:     nil,
			predicate: func(k string, v Item) bool { return v.Active },
			expected:  0,
		},
		{
			name:      "empty map",
			input:     map[string]Item{},
			predicate: func(k string, v Item) bool { return v.Active },
			expected:  0,
		},
		{
			name: "count active items",
			input: map[string]Item{
				"item1": {Name: "A", Value: 10, Active: true},
				"item2": {Name: "B", Value: 20, Active: false},
				"item3": {Name: "C", Value: 30, Active: true},
			},
			predicate: func(k string, v Item) bool { return v.Active },
			expected:  2,
		},
		{
			name: "count by value field",
			input: map[string]Item{
				"item1": {Name: "A", Value: 10, Active: true},
				"item2": {Name: "B", Value: 20, Active: false},
				"item3": {Name: "C", Value: 30, Active: true},
			},
			predicate: func(k string, v Item) bool { return v.Value > 15 },
			expected:  2,
		},
		{
			name: "count by name length",
			input: map[string]Item{
				"a": {Name: "Alice", Value: 10, Active: true},
				"b": {Name: "Bob", Value: 20, Active: false},
				"c": {Name: "Charlie", Value: 30, Active: true},
			},
			predicate: func(k string, v Item) bool { return len(v.Name) > 3 },
			expected:  2,
		},
		{
			name: "count by multiple conditions",
			input: map[string]Item{
				"a": {Name: "Alice", Value: 10, Active: true},
				"b": {Name: "Bob", Value: 20, Active: false},
				"c": {Name: "Charlie", Value: 30, Active: true},
				"d": {Name: "Dave", Value: 5, Active: true},
			},
			predicate: func(k string, v Item) bool { return v.Active && v.Value > 10 },
			expected:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Count(tt.input, tt.predicate)
			if result != tt.expected {
				t.Errorf("Count() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestCount_IntInt_LargeMap(t *testing.T) {
	tests := []struct {
		name      string
		size      int
		predicate func(int, int) int
		expected  int
	}{
		{
			name: "count even keys in large map",
			size: 100,
			predicate: func(k int, v int) int {
				if k%2 == 0 {
					return 1
				}
				return 0
			},
			expected: 50,
		},
		{
			name: "count where value > key",
			size: 50,
			predicate: func(k int, v int) int {
				if v > k {
					return 1
				}
				return 0
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := make(map[int]int)
			for i := 1; i <= tt.size; i++ {
				m[i] = i
			}

			count := 0
			for k, v := range m {
				count += tt.predicate(k, v)
			}

			if count != tt.expected {
				t.Errorf("Count() = %d, want %d", count, tt.expected)
			}
		})
	}
}

func TestCount_StringStringWithKeyFilter(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]string
		predicate func(string, string) bool
		expected  int
	}{
		{
			name:      "nil map",
			input:     nil,
			predicate: func(k string, v string) bool { return k == v },
			expected:  0,
		},
		{
			name:      "empty map",
			input:     map[string]string{},
			predicate: func(k string, v string) bool { return k == v },
			expected:  0,
		},
		{
			name: "count where key equals value",
			input: map[string]string{
				"a": "a",
				"b": "x",
				"c": "c",
				"d": "y",
			},
			predicate: func(k string, v string) bool { return k == v },
			expected:  2,
		},
		{
			name: "count same first letter",
			input: map[string]string{
				"a": "apple",
				"b": "banana",
				"c": "cat",
				"d": "elephant",
			},
			predicate: func(k string, v string) bool {
				return len(v) > 0 && string(v[0]) == k
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Count(tt.input, tt.predicate)
			if result != tt.expected {
				t.Errorf("Count() = %d, want %d", result, tt.expected)
			}
		})
	}
}
