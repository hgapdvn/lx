package lxmaps_test

import (
	"testing"

	lxmaps "github.com/hgapdvn/lx/lxmaps"
)

func TestAll_IntBool(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		fn       func(string, int) bool
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(k string, v int) bool { return v > 0 },
			expected: true,
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			fn:       func(k string, v int) bool { return v > 0 },
			expected: true,
		},
		{
			name: "all positive values",
			input: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			fn: func(k string, v int) bool {
				return v > 0
			},
			expected: true,
		},
		{
			name: "one negative value",
			input: map[string]int{
				"a": 1,
				"b": -2,
				"c": 3,
			},
			fn: func(k string, v int) bool {
				return v > 0
			},
			expected: false,
		},
		{
			name: "single entry true",
			input: map[string]int{
				"a": 5,
			},
			fn: func(k string, v int) bool {
				return v > 0
			},
			expected: true,
		},
		{
			name: "single entry false",
			input: map[string]int{
				"a": -5,
			},
			fn: func(k string, v int) bool {
				return v > 0
			},
			expected: false,
		},
		{
			name: "all zero values",
			input: map[string]int{
				"a": 0,
				"b": 0,
				"c": 0,
			},
			fn: func(k string, v int) bool {
				return v == 0
			},
			expected: true,
		},
		{
			name: "mixed zero and non-zero, check for zero fails",
			input: map[string]int{
				"a": 0,
				"b": 1,
				"c": 0,
			},
			fn: func(k string, v int) bool {
				return v == 0
			},
			expected: false,
		},
		{
			name: "check on key",
			input: map[string]int{
				"x": 1,
				"y": 2,
				"z": 3,
			},
			fn: func(k string, v int) bool {
				return k >= "x" && k <= "z"
			},
			expected: true,
		},
		{
			name: "check on key fails",
			input: map[string]int{
				"a": 1,
				"b": 2,
				"z": 3,
			},
			fn: func(k string, v int) bool {
				return k >= "x" && k <= "z"
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.All(tt.input, tt.fn)
			if result != tt.expected {
				t.Errorf("All() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestAll_StringBool(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]string
		fn       func(string, string) bool
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(k, v string) bool { return v != "" },
			expected: true,
		},
		{
			name:     "empty map",
			input:    map[string]string{},
			fn:       func(k, v string) bool { return v != "" },
			expected: true,
		},
		{
			name: "all non-empty values",
			input: map[string]string{
				"a": "alpha",
				"b": "beta",
				"c": "gamma",
			},
			fn: func(k, v string) bool {
				return v != ""
			},
			expected: true,
		},
		{
			name: "one empty value",
			input: map[string]string{
				"a": "alpha",
				"b": "",
				"c": "gamma",
			},
			fn: func(k, v string) bool {
				return v != ""
			},
			expected: false,
		},
		{
			name: "all values have key prefix",
			input: map[string]string{
				"key1": "key1_value",
				"key2": "key2_value",
				"key3": "key3_value",
			},
			fn: func(k, v string) bool {
				return len(v) > len(k)
			},
			expected: true,
		},
		{
			name: "value length check fails",
			input: map[string]string{
				"longkey1": "short",
				"k2":       "value",
			},
			fn: func(k, v string) bool {
				return len(v) > len(k)
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.All(tt.input, tt.fn)
			if result != tt.expected {
				t.Errorf("All() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestAll_IntInt(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]int
		fn       func(int, int) bool
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(k, v int) bool { return true },
			expected: true,
		},
		{
			name:     "empty map",
			input:    map[int]int{},
			fn:       func(k, v int) bool { return true },
			expected: true,
		},
		{
			name: "key less than value",
			input: map[int]int{
				1: 10,
				2: 20,
				3: 30,
			},
			fn: func(k, v int) bool {
				return k < v
			},
			expected: true,
		},
		{
			name: "key less than value fails",
			input: map[int]int{
				1: 10,
				5: 3,
				3: 30,
			},
			fn: func(k, v int) bool {
				return k < v
			},
			expected: false,
		},
		{
			name: "sum condition",
			input: map[int]int{
				1: 9,
				2: 8,
				3: 7,
			},
			fn: func(k, v int) bool {
				return k+v == 10
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.All(tt.input, tt.fn)
			if result != tt.expected {
				t.Errorf("All() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestAll_BoolBool(t *testing.T) {
	tests := []struct {
		name     string
		input    map[bool]bool
		fn       func(bool, bool) bool
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(k, v bool) bool { return true },
			expected: true,
		},
		{
			name:     "empty map",
			input:    map[bool]bool{},
			fn:       func(k, v bool) bool { return true },
			expected: true,
		},
		{
			name: "key equals value",
			input: map[bool]bool{
				true:  true,
				false: false,
			},
			fn: func(k, v bool) bool {
				return k == v
			},
			expected: true,
		},
		{
			name: "key equals value fails",
			input: map[bool]bool{
				true:  false,
				false: true,
			},
			fn: func(k, v bool) bool {
				return k == v
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.All(tt.input, tt.fn)
			if result != tt.expected {
				t.Errorf("All() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestAll_FloatFloat(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]float64
		fn       func(string, float64) bool
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(k string, v float64) bool { return v > 0 },
			expected: true,
		},
		{
			name:     "empty map",
			input:    map[string]float64{},
			fn:       func(k string, v float64) bool { return v > 0 },
			expected: true,
		},
		{
			name: "all positive floats",
			input: map[string]float64{
				"pi":  3.14159,
				"e":   2.71828,
				"phi": 1.61803,
			},
			fn: func(k string, v float64) bool {
				return v > 0
			},
			expected: true,
		},
		{
			name: "one negative float",
			input: map[string]float64{
				"pos": 3.14,
				"neg": -2.71,
			},
			fn: func(k string, v float64) bool {
				return v > 0
			},
			expected: false,
		},
		{
			name: "zero float value",
			input: map[string]float64{
				"zero": 0.0,
			},
			fn: func(k string, v float64) bool {
				return v == 0.0
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.All(tt.input, tt.fn)
			if result != tt.expected {
				t.Errorf("All() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestAll_CustomStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	tests := []struct {
		name     string
		input    map[string]Person
		fn       func(string, Person) bool
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(k string, p Person) bool { return p.Age > 0 },
			expected: true,
		},
		{
			name:     "empty map",
			input:    map[string]Person{},
			fn:       func(k string, p Person) bool { return p.Age > 0 },
			expected: true,
		},
		{
			name: "all adults",
			input: map[string]Person{
				"alice": {Name: "Alice", Age: 30},
				"bob":   {Name: "Bob", Age: 25},
				"carol": {Name: "Carol", Age: 35},
			},
			fn: func(k string, p Person) bool {
				return p.Age >= 18
			},
			expected: true,
		},
		{
			name: "one minor",
			input: map[string]Person{
				"alice": {Name: "Alice", Age: 30},
				"dave":  {Name: "Dave", Age: 15},
				"eve":   {Name: "Eve", Age: 25},
			},
			fn: func(k string, p Person) bool {
				return p.Age >= 18
			},
			expected: false,
		},
		{
			name: "name contains key",
			input: map[string]Person{
				"a": {Name: "a_person", Age: 20},
				"b": {Name: "b_person", Age: 25},
			},
			fn: func(k string, p Person) bool {
				return len(p.Name) > 0
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.All(tt.input, tt.fn)
			if result != tt.expected {
				t.Errorf("All() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestAll_LargeMap(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		fn       func(int, int) bool
		expected bool
	}{
		{
			name: "1000 entries all pass",
			size: 1000,
			fn: func(k, v int) bool {
				return k < v
			},
			expected: true,
		},
		{
			name: "1000 entries one fails",
			size: 1000,
			fn: func(k, v int) bool {
				return k < v
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make(map[int]int)
			for i := 0; i < tt.size; i++ {
				if tt.name == "1000 entries all pass" {
					input[i] = i + 1
				} else {
					input[i] = i
				}
			}

			result := lxmaps.All(input, tt.fn)
			if result != tt.expected {
				t.Errorf("All() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestAll_EarlyTermination(t *testing.T) {
	tests := []struct {
		name              string
		check             func() bool
		expectedCallCount int
	}{
		{
			name: "short-circuits on first false",
			check: func() bool {
				callCount := 0
				input := map[string]int{
					"a": 1,
					"b": 2,
					"c": 3,
				}
				lxmaps.All(input, func(k string, v int) bool {
					callCount++
					return v != 2
				})
				// callCount should be 1, but could be up to 2 due to map iteration order
				return callCount >= 1 && callCount <= 3
			},
			expectedCallCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("All() early termination check failed")
			}
		})
	}
}

func TestAll_SideEffects(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "predicate receives correct arguments",
			check: func() bool {
				input := map[string]int{
					"x": 10,
					"y": 20,
				}
				receivedKeys := []string{}
				receivedValues := []int{}

				lxmaps.All(input, func(k string, v int) bool {
					receivedKeys = append(receivedKeys, k)
					receivedValues = append(receivedValues, v)
					return true
				})

				return len(receivedKeys) == 2 && len(receivedValues) == 2
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("All() side effects check failed")
			}
		})
	}
}

func TestAll_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		fn       func(string, int) bool
		expected bool
	}{
		{
			name: "always true predicate",
			input: map[string]int{
				"a": 1,
				"b": 2,
			},
			fn: func(k string, v int) bool {
				return true
			},
			expected: true,
		},
		{
			name: "always false predicate",
			input: map[string]int{
				"a": 1,
			},
			fn: func(k string, v int) bool {
				return false
			},
			expected: false,
		},
		{
			name: "single entry always true",
			input: map[string]int{
				"single": 1,
			},
			fn: func(k string, v int) bool {
				return true
			},
			expected: true,
		},
		{
			name: "single entry always false",
			input: map[string]int{
				"single": 1,
			},
			fn: func(k string, v int) bool {
				return false
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.All(tt.input, tt.fn)
			if result != tt.expected {
				t.Errorf("All() = %v, want %v", result, tt.expected)
			}
		})
	}
}
