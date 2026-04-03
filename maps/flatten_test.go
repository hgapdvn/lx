package lxmaps_test

import (
	"sort"
	"testing"

	lxmaps "github.com/hgapdvn/lx/maps"
)

func TestFlatten_StringInt(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string][]int
		checkFunc func([]int) bool
		checkMsg  string
	}{
		{
			name:      "nil map",
			input:     nil,
			checkFunc: func(result []int) bool { return result == nil },
			checkMsg:  "should return nil",
		},
		{
			name:      "empty map",
			input:     map[string][]int{},
			checkFunc: func(result []int) bool { return len(result) == 0 },
			checkMsg:  "should return empty slice",
		},
		{
			name: "single key single element",
			input: map[string][]int{
				"a": {1},
			},
			checkFunc: func(result []int) bool {
				return len(result) == 1 && result[0] == 1
			},
			checkMsg: "should flatten single element",
		},
		{
			name: "single key multiple elements",
			input: map[string][]int{
				"a": {1, 2, 3},
			},
			checkFunc: func(result []int) bool {
				return len(result) == 3 && result[0] == 1 && result[1] == 2 && result[2] == 3
			},
			checkMsg: "should flatten multiple elements",
		},
		{
			name: "multiple keys multiple elements",
			input: map[string][]int{
				"a": {1, 2},
				"b": {3, 4},
				"c": {5},
			},
			checkFunc: func(result []int) bool {
				return len(result) == 5 && containsAll(result, []int{1, 2, 3, 4, 5})
			},
			checkMsg: "should flatten all elements",
		},
		{
			name: "empty slices mixed",
			input: map[string][]int{
				"a": {1},
				"b": {},
				"c": {2, 3},
			},
			checkFunc: func(result []int) bool {
				return len(result) == 3 && containsAll(result, []int{1, 2, 3})
			},
			checkMsg: "should handle empty slices",
		},
		{
			name: "all empty slices",
			input: map[string][]int{
				"a": {},
				"b": {},
				"c": {},
			},
			checkFunc: func(result []int) bool {
				return len(result) == 0
			},
			checkMsg: "should return empty slice for all empty slices",
		},
		{
			name: "large values",
			input: map[string][]int{
				"a": {1000000, 2000000},
				"b": {3000000},
			},
			checkFunc: func(result []int) bool {
				return len(result) == 3 && containsAll(result, []int{1000000, 2000000, 3000000})
			},
			checkMsg: "should handle large values",
		},
		{
			name: "negative values",
			input: map[string][]int{
				"a": {-1, -2},
				"b": {0, 1},
			},
			checkFunc: func(result []int) bool {
				return len(result) == 4 && containsAll(result, []int{-1, -2, 0, 1})
			},
			checkMsg: "should handle negative values",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Flatten(tt.input)
			if !tt.checkFunc(result) {
				t.Errorf("Flatten() %s", tt.checkMsg)
			}
		})
	}
}

func TestFlatten_IntString(t *testing.T) {
	tests := []struct {
		name      string
		input     map[int][]string
		checkFunc func([]string) bool
		checkMsg  string
	}{
		{
			name:      "nil map",
			input:     nil,
			checkFunc: func(result []string) bool { return result == nil },
			checkMsg:  "should return nil",
		},
		{
			name:      "empty map",
			input:     map[int][]string{},
			checkFunc: func(result []string) bool { return len(result) == 0 },
			checkMsg:  "should return empty slice",
		},
		{
			name: "flatten strings",
			input: map[int][]string{
				1: {"apple", "apricot"},
				2: {"banana"},
				3: {"cherry", "clementine"},
			},
			checkFunc: func(result []string) bool {
				return len(result) == 5 && containsAllStrings(result, []string{"apple", "apricot", "banana", "cherry", "clementine"})
			},
			checkMsg: "should flatten all strings",
		},
		{
			name: "empty strings",
			input: map[int][]string{
				1: {"", "a"},
				2: {""},
			},
			checkFunc: func(result []string) bool {
				return len(result) == 3 && containsAllStrings(result, []string{"", "a", ""})
			},
			checkMsg: "should handle empty strings",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Flatten(tt.input)
			if !tt.checkFunc(result) {
				t.Errorf("Flatten() %s", tt.checkMsg)
			}
		})
	}
}

func TestFlatten_StringBool(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string][]bool
		checkFunc func([]bool) bool
		checkMsg  string
	}{
		{
			name:      "nil map",
			input:     nil,
			checkFunc: func(result []bool) bool { return result == nil },
			checkMsg:  "should return nil",
		},
		{
			name: "flatten bools",
			input: map[string][]bool{
				"a": {true, false},
				"b": {true},
			},
			checkFunc: func(result []bool) bool {
				if len(result) != 3 {
					return false
				}
				trueCount := 0
				falseCount := 0
				for _, v := range result {
					if v {
						trueCount++
					} else {
						falseCount++
					}
				}
				return trueCount == 2 && falseCount == 1
			},
			checkMsg: "should flatten boolean slices",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Flatten(tt.input)
			if !tt.checkFunc(result) {
				t.Errorf("Flatten() %s", tt.checkMsg)
			}
		})
	}
}

func TestFlatten_StringFloat(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string][]float64
		checkFunc func([]float64) bool
		checkMsg  string
	}{
		{
			name: "flatten floats",
			input: map[string][]float64{
				"a": {3.14, 2.71},
				"b": {1.41},
			},
			checkFunc: func(result []float64) bool {
				return len(result) == 3
			},
			checkMsg: "should flatten float slices",
		},
		{
			name: "zero floats",
			input: map[string][]float64{
				"a": {0.0},
				"b": {1.0, -1.0},
			},
			checkFunc: func(result []float64) bool {
				return len(result) == 3
			},
			checkMsg: "should handle zero values",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Flatten(tt.input)
			if !tt.checkFunc(result) {
				t.Errorf("Flatten() %s", tt.checkMsg)
			}
		})
	}
}

func TestFlatten_CustomStruct(t *testing.T) {
	type Item struct {
		Name  string
		Value int
	}

	tests := []struct {
		name      string
		input     map[string][]Item
		checkFunc func([]Item) bool
		checkMsg  string
	}{
		{
			name:      "nil map",
			input:     nil,
			checkFunc: func(result []Item) bool { return result == nil },
			checkMsg:  "should return nil",
		},
		{
			name: "flatten structs",
			input: map[string][]Item{
				"group1": {{Name: "A", Value: 1}, {Name: "B", Value: 2}},
				"group2": {{Name: "C", Value: 3}},
			},
			checkFunc: func(result []Item) bool {
				return len(result) == 3
			},
			checkMsg: "should flatten struct slices",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Flatten(tt.input)
			if !tt.checkFunc(result) {
				t.Errorf("Flatten() %s", tt.checkMsg)
			}
		})
	}
}

func TestFlatten_LargeMap(t *testing.T) {
	tests := []struct {
		name       string
		numKeys    int
		elemPerKey int
		checkFunc  func([]int) bool
		checkMsg   string
	}{
		{
			name:       "100 keys 10 elements each",
			numKeys:    100,
			elemPerKey: 10,
			checkFunc: func(result []int) bool {
				return len(result) == 1000
			},
			checkMsg: "should flatten 1000 total elements",
		},
		{
			name:       "1000 keys 1 element each",
			numKeys:    1000,
			elemPerKey: 1,
			checkFunc: func(result []int) bool {
				return len(result) == 1000
			},
			checkMsg: "should flatten 1000 total elements",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := make(map[string][]int)
			for i := 0; i < tt.numKeys; i++ {
				key := "key_" + string(rune(i%10)) + "_" + string(rune(i/10))
				slice := make([]int, tt.elemPerKey)
				for j := 0; j < tt.elemPerKey; j++ {
					slice[j] = i*tt.elemPerKey + j
				}
				m[key] = slice
			}

			result := lxmaps.Flatten(m)
			if !tt.checkFunc(result) {
				t.Errorf("Flatten() %s", tt.checkMsg)
			}
		})
	}
}

func TestFlatten_PreservesElements(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string][]int
		checkFunc func([]int) bool
		checkMsg  string
	}{
		{
			name: "preserves all elements from single key",
			input: map[string][]int{
				"a": {1, 2, 3, 4, 5},
			},
			checkFunc: func(result []int) bool {
				if len(result) != 5 {
					return false
				}
				for i, v := range result {
					if v != i+1 {
						return false
					}
				}
				return true
			},
			checkMsg: "should preserve order within single slice",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Flatten(tt.input)
			if !tt.checkFunc(result) {
				t.Errorf("Flatten() %s", tt.checkMsg)
			}
		})
	}
}

func TestFlatten_AllValuesPresent(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string][]int
		expected []int
	}{
		{
			name: "all values present",
			input: map[string][]int{
				"a": {1, 2},
				"b": {3, 4},
				"c": {5},
			},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name: "all values with duplicates",
			input: map[string][]int{
				"a": {1, 1},
				"b": {2, 2},
			},
			expected: []int{1, 1, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Flatten(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("Flatten() length = %d, want %d", len(result), len(tt.expected))
			}
			// Check all expected values are present (order may vary)
			if !containsAll(result, tt.expected) {
				t.Errorf("Flatten() missing some values")
			}
		})
	}
}

// Helper function to check if all elements are present
func containsAll(result []int, expected []int) bool {
	if len(result) != len(expected) {
		return false
	}
	resultCopy := make([]int, len(result))
	copy(resultCopy, result)
	sort.Ints(resultCopy)

	expectedCopy := make([]int, len(expected))
	copy(expectedCopy, expected)
	sort.Ints(expectedCopy)

	for i, v := range resultCopy {
		if v != expectedCopy[i] {
			return false
		}
	}
	return true
}

// Helper function for strings
func containsAllStrings(result []string, expected []string) bool {
	if len(result) != len(expected) {
		return false
	}
	resultCopy := make([]string, len(result))
	copy(resultCopy, result)
	sort.Strings(resultCopy)

	expectedCopy := make([]string, len(expected))
	copy(expectedCopy, expected)
	sort.Strings(expectedCopy)

	for i, v := range resultCopy {
		if v != expectedCopy[i] {
			return false
		}
	}
	return true
}
