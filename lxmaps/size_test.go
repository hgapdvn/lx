package lxmaps_test

import (
	"testing"

	"github.com/hgapdvn/lx/lxmaps"
)

func TestSize_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected int
	}{
		{
			name:     "nil map",
			input:    nil,
			expected: 0,
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: 0,
		},
		{
			name:     "single entry",
			input:    map[string]int{"a": 1},
			expected: 1,
		},
		{
			name:     "two entries",
			input:    map[string]int{"a": 1, "b": 2},
			expected: 2,
		},
		{
			name:     "three entries",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			expected: 3,
		},
		{
			name:     "ten entries",
			input:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Size(tt.input)
			if result != tt.expected {
				t.Errorf("Size() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestSize_IntString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		expected int
	}{
		{
			name:     "nil map",
			input:    nil,
			expected: 0,
		},
		{
			name:     "empty map",
			input:    map[int]string{},
			expected: 0,
		},
		{
			name:     "single entry",
			input:    map[int]string{1: "one"},
			expected: 1,
		},
		{
			name:     "multiple entries",
			input:    map[int]string{1: "one", 2: "two", 3: "three"},
			expected: 3,
		},
		{
			name:     "negative keys",
			input:    map[int]string{-1: "neg1", -2: "neg2", 0: "zero"},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Size(tt.input)
			if result != tt.expected {
				t.Errorf("Size() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestSize_StringBool(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]bool
		expected int
	}{
		{
			name:     "nil map",
			input:    nil,
			expected: 0,
		},
		{
			name:     "empty map",
			input:    map[string]bool{},
			expected: 0,
		},
		{
			name:     "all true",
			input:    map[string]bool{"a": true, "b": true, "c": true},
			expected: 3,
		},
		{
			name:     "all false",
			input:    map[string]bool{"a": false, "b": false},
			expected: 2,
		},
		{
			name:     "mixed",
			input:    map[string]bool{"a": true, "b": false, "c": true, "d": false},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Size(tt.input)
			if result != tt.expected {
				t.Errorf("Size() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestSize_StringFloat(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]float64
		expected int
	}{
		{
			name:     "nil map",
			input:    nil,
			expected: 0,
		},
		{
			name:     "empty map",
			input:    map[string]float64{},
			expected: 0,
		},
		{
			name:     "single entry",
			input:    map[string]float64{"pi": 3.14159},
			expected: 1,
		},
		{
			name:     "multiple entries",
			input:    map[string]float64{"pi": 3.14159, "e": 2.71828, "sqrt2": 1.41421},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Size(tt.input)
			if result != tt.expected {
				t.Errorf("Size() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestSize_StringCustomStruct(t *testing.T) {
	type Item struct {
		Name  string
		Value int
	}

	tests := []struct {
		name     string
		input    map[string]Item
		expected int
	}{
		{
			name:     "nil map",
			input:    nil,
			expected: 0,
		},
		{
			name:     "empty map",
			input:    map[string]Item{},
			expected: 0,
		},
		{
			name: "single struct",
			input: map[string]Item{
				"a": {Name: "Alice", Value: 10},
			},
			expected: 1,
		},
		{
			name: "multiple structs",
			input: map[string]Item{
				"a": {Name: "Alice", Value: 10},
				"b": {Name: "Bob", Value: 20},
				"c": {Name: "Charlie", Value: 30},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Size(tt.input)
			if result != tt.expected {
				t.Errorf("Size() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestSize_LargeMap(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{
			name:     "100 entries",
			size:     100,
			expected: 100,
		},
		{
			name:     "1000 entries",
			size:     1000,
			expected: 1000,
		},
		{
			name:     "10000 entries",
			size:     10000,
			expected: 10000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := make(map[int]int)
			for i := 0; i < tt.size; i++ {
				m[i] = i
			}

			result := lxmaps.Size(m)
			if result != tt.expected {
				t.Errorf("Size() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestSize_StringString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]string
		expected int
	}{
		{
			name:     "nil map",
			input:    nil,
			expected: 0,
		},
		{
			name:     "empty map",
			input:    map[string]string{},
			expected: 0,
		},
		{
			name:     "single string",
			input:    map[string]string{"key": "value"},
			expected: 1,
		},
		{
			name: "multiple strings",
			input: map[string]string{
				"a": "apple",
				"b": "banana",
				"c": "cherry",
				"d": "date",
				"e": "elderberry",
			},
			expected: 5,
		},
		{
			name: "empty string values",
			input: map[string]string{
				"a": "",
				"b": "value",
				"c": "",
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Size(tt.input)
			if result != tt.expected {
				t.Errorf("Size() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestSize_IntInt(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]int
		expected int
	}{
		{
			name:     "nil map",
			input:    nil,
			expected: 0,
		},
		{
			name:     "empty map",
			input:    map[int]int{},
			expected: 0,
		},
		{
			name:     "single entry",
			input:    map[int]int{1: 10},
			expected: 1,
		},
		{
			name:     "multiple entries",
			input:    map[int]int{1: 10, 2: 20, 3: 30, 4: 40, 5: 50},
			expected: 5,
		},
		{
			name:     "negative keys",
			input:    map[int]int{-1: 10, -2: 20, 0: 0, 1: 10},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Size(tt.input)
			if result != tt.expected {
				t.Errorf("Size() = %d, want %d", result, tt.expected)
			}
		})
	}
}
