package lxmaps_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/maps"
)

func TestContainsAnyKey_StringInt(t *testing.T) {
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
			name:     "single entry multiple keys with one match",
			input:    map[string]int{"a": 1},
			keys:     []string{"b", "a", "c"},
			expected: true,
		},
		{
			name:     "multiple entries no keys match",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			keys:     []string{"d", "e", "f"},
			expected: false,
		},
		{
			name:     "multiple entries first key matches",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			keys:     []string{"a", "d", "e"},
			expected: true,
		},
		{
			name:     "multiple entries middle key matches",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			keys:     []string{"d", "b", "e"},
			expected: true,
		},
		{
			name:     "multiple entries last key matches",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			keys:     []string{"d", "e", "c"},
			expected: true,
		},
		{
			name:     "multiple entries multiple keys match",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			keys:     []string{"a", "b", "c"},
			expected: true,
		},
		{
			name:     "map with single key no match",
			input:    map[string]int{"key": 10},
			keys:     []string{"other"},
			expected: false,
		},
		{
			name:     "map with single key match",
			input:    map[string]int{"key": 10},
			keys:     []string{"key"},
			expected: true,
		},
		{
			name:     "empty keys list",
			input:    map[string]int{"a": 1, "b": 2},
			keys:     []string{},
			expected: false,
		},
		{
			name:     "duplicate keys in input list",
			input:    map[string]int{"a": 1, "b": 2},
			keys:     []string{"a", "a", "a"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.ContainsAnyKeys(tt.input, tt.keys...)
			if result != tt.expected {
				t.Errorf("ContainsAnyKeys() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestContainsAnyKey_IntString(t *testing.T) {
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
			name:     "first key matches",
			input:    map[int]string{1: "a", 2: "b", 3: "c"},
			keys:     []int{1, 4, 5},
			expected: true,
		},
		{
			name:     "middle key matches",
			input:    map[int]string{1: "a", 2: "b", 3: "c"},
			keys:     []int{4, 2, 5},
			expected: true,
		},
		{
			name:     "last key matches",
			input:    map[int]string{1: "a", 2: "b", 3: "c"},
			keys:     []int{4, 5, 3},
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
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.ContainsAnyKeys(tt.input, tt.keys...)
			if result != tt.expected {
				t.Errorf("ContainsAnyKeys() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestContainsAnyKey_BoolFloat64(t *testing.T) {
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
			name:     "true key exists",
			input:    map[bool]float64{true: 1.5, false: 2.5},
			keys:     []bool{true},
			expected: true,
		},
		{
			name:     "false key exists",
			input:    map[bool]float64{true: 1.5, false: 2.5},
			keys:     []bool{false},
			expected: true,
		},
		{
			name:     "both keys exist",
			input:    map[bool]float64{true: 1.5, false: 2.5},
			keys:     []bool{true, false},
			expected: true,
		},
		{
			name:     "true key exists but looking for false",
			input:    map[bool]float64{true: 1.5},
			keys:     []bool{false},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.ContainsAnyKeys(tt.input, tt.keys...)
			if result != tt.expected {
				t.Errorf("ContainsAnyKeys() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
