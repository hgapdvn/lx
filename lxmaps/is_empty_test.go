package lxmaps_test

import (
	"testing"

	"github.com/hgapdvn/lx/lxmaps"
)

func TestIsEmpty_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			expected: true,
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: true,
		},
		{
			name:     "single entry",
			input:    map[string]int{"a": 1},
			expected: false,
		},
		{
			name:     "two entries",
			input:    map[string]int{"a": 1, "b": 2},
			expected: false,
		},
		{
			name:     "many entries",
			input:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.IsEmpty(tt.input)
			if result != tt.expected {
				t.Errorf("IsEmpty() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsEmpty_IntString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			expected: true,
		},
		{
			name:     "empty map",
			input:    map[int]string{},
			expected: true,
		},
		{
			name:     "single entry",
			input:    map[int]string{1: "one"},
			expected: false,
		},
		{
			name:     "multiple entries",
			input:    map[int]string{1: "one", 2: "two", 3: "three"},
			expected: false,
		},
		{
			name:     "negative keys",
			input:    map[int]string{-1: "neg1", -2: "neg2"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.IsEmpty(tt.input)
			if result != tt.expected {
				t.Errorf("IsEmpty() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsEmpty_StringBool(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]bool
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			expected: true,
		},
		{
			name:     "empty map",
			input:    map[string]bool{},
			expected: true,
		},
		{
			name:     "single true",
			input:    map[string]bool{"a": true},
			expected: false,
		},
		{
			name:     "single false",
			input:    map[string]bool{"a": false},
			expected: false,
		},
		{
			name:     "multiple mixed",
			input:    map[string]bool{"a": true, "b": false, "c": true},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.IsEmpty(tt.input)
			if result != tt.expected {
				t.Errorf("IsEmpty() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsEmpty_StringFloat(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]float64
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			expected: true,
		},
		{
			name:     "empty map",
			input:    map[string]float64{},
			expected: true,
		},
		{
			name:     "single entry",
			input:    map[string]float64{"pi": 3.14159},
			expected: false,
		},
		{
			name:     "multiple entries",
			input:    map[string]float64{"pi": 3.14, "e": 2.71, "sqrt2": 1.41},
			expected: false,
		},
		{
			name:     "zero value",
			input:    map[string]float64{"zero": 0.0},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.IsEmpty(tt.input)
			if result != tt.expected {
				t.Errorf("IsEmpty() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsEmpty_StringCustomStruct(t *testing.T) {
	type Item struct {
		Name  string
		Value int
	}

	tests := []struct {
		name     string
		input    map[string]Item
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			expected: true,
		},
		{
			name:     "empty map",
			input:    map[string]Item{},
			expected: true,
		},
		{
			name: "single struct",
			input: map[string]Item{
				"a": {Name: "Alice", Value: 10},
			},
			expected: false,
		},
		{
			name: "multiple structs",
			input: map[string]Item{
				"a": {Name: "Alice", Value: 10},
				"b": {Name: "Bob", Value: 20},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.IsEmpty(tt.input)
			if result != tt.expected {
				t.Errorf("IsEmpty() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsEmpty_StringString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]string
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			expected: true,
		},
		{
			name:     "empty map",
			input:    map[string]string{},
			expected: true,
		},
		{
			name:     "single entry",
			input:    map[string]string{"key": "value"},
			expected: false,
		},
		{
			name:     "empty string value",
			input:    map[string]string{"key": ""},
			expected: false,
		},
		{
			name: "multiple entries",
			input: map[string]string{
				"a": "apple",
				"b": "banana",
				"c": "cherry",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.IsEmpty(tt.input)
			if result != tt.expected {
				t.Errorf("IsEmpty() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsEmpty_IntInt(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]int
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			expected: true,
		},
		{
			name:     "empty map",
			input:    map[int]int{},
			expected: true,
		},
		{
			name:     "single zero entry",
			input:    map[int]int{0: 0},
			expected: false,
		},
		{
			name:     "single entry",
			input:    map[int]int{1: 10},
			expected: false,
		},
		{
			name:     "multiple entries",
			input:    map[int]int{1: 10, 2: 20, 3: 30},
			expected: false,
		},
		{
			name:     "negative keys and values",
			input:    map[int]int{-1: -10, -2: -20},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.IsEmpty(tt.input)
			if result != tt.expected {
				t.Errorf("IsEmpty() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsEmpty_LargeMap(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected bool
	}{
		{
			name:     "1 entry",
			size:     1,
			expected: false,
		},
		{
			name:     "100 entries",
			size:     100,
			expected: false,
		},
		{
			name:     "1000 entries",
			size:     1000,
			expected: false,
		},
		{
			name:     "10000 entries",
			size:     10000,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := make(map[int]int)
			for i := 0; i < tt.size; i++ {
				m[i] = i
			}

			result := lxmaps.IsEmpty(m)
			if result != tt.expected {
				t.Errorf("IsEmpty() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsEmpty_InterfaceTypes(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			expected: true,
		},
		{
			name:     "empty map",
			input:    map[string]interface{}{},
			expected: true,
		},
		{
			name:     "single interface entry",
			input:    map[string]interface{}{"key": "value"},
			expected: false,
		},
		{
			name: "mixed interface types",
			input: map[string]interface{}{
				"str":  "string",
				"int":  42,
				"bool": true,
				"nil":  nil,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.IsEmpty(tt.input)
			if result != tt.expected {
				t.Errorf("IsEmpty() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsEmpty_Consistency_WithSize(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]int
	}{
		{
			name:  "nil map",
			input: nil,
		},
		{
			name:  "empty map",
			input: map[string]int{},
		},
		{
			name:  "single entry",
			input: map[string]int{"a": 1},
		},
		{
			name:  "five entries",
			input: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isEmpty := lxmaps.IsEmpty(tt.input)
			size := lxmaps.Size(tt.input)

			// IsEmpty should be consistent with Size
			expectedEmpty := size == 0
			if isEmpty != expectedEmpty {
				t.Errorf("IsEmpty() = %v, but Size() = %d (expected consistent)", isEmpty, size)
			}
		})
	}
}
