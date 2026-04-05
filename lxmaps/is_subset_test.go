package lxmaps_test

import (
	"testing"

	"github.com/hgapdvn/lx/lxmaps"
)

func TestIsSubset_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		subset   map[string]int
		superset map[string]int
		expected bool
	}{
		{
			name:     "nil subset nil superset",
			subset:   nil,
			superset: nil,
			expected: true,
		},
		{
			name:     "empty subset nil superset",
			subset:   map[string]int{},
			superset: nil,
			expected: true,
		},
		{
			name:     "empty subset empty superset",
			subset:   map[string]int{},
			superset: map[string]int{},
			expected: true,
		},
		{
			name:     "nil subset non-empty superset",
			subset:   nil,
			superset: map[string]int{"a": 1},
			expected: true,
		},
		{
			name:     "empty subset non-empty superset",
			subset:   map[string]int{},
			superset: map[string]int{"a": 1, "b": 2},
			expected: true,
		},
		{
			name:     "non-empty subset nil superset",
			subset:   map[string]int{"a": 1},
			superset: nil,
			expected: false,
		},
		{
			name:     "identical single entry",
			subset:   map[string]int{"a": 1},
			superset: map[string]int{"a": 1},
			expected: true,
		},
		{
			name:     "identical multiple entries",
			subset:   map[string]int{"a": 1, "b": 2, "c": 3},
			superset: map[string]int{"a": 1, "b": 2, "c": 3},
			expected: true,
		},
		{
			name:     "subset is true subset",
			subset:   map[string]int{"a": 1, "b": 2},
			superset: map[string]int{"a": 1, "b": 2, "c": 3},
			expected: true,
		},
		{
			name:     "subset with extra key",
			subset:   map[string]int{"a": 1, "b": 2, "d": 4},
			superset: map[string]int{"a": 1, "b": 2, "c": 3},
			expected: false,
		},
		{
			name:     "subset with different value",
			subset:   map[string]int{"a": 1, "b": 99},
			superset: map[string]int{"a": 1, "b": 2, "c": 3},
			expected: false,
		},
		{
			name:     "superset is smaller",
			subset:   map[string]int{"a": 1, "b": 2, "c": 3},
			superset: map[string]int{"a": 1, "b": 2},
			expected: false,
		},
		{
			name:     "disjoint sets",
			subset:   map[string]int{"x": 10, "y": 20},
			superset: map[string]int{"a": 1, "b": 2},
			expected: false,
		},
		{
			name:     "single matching key",
			subset:   map[string]int{"a": 1},
			superset: map[string]int{"a": 1, "b": 2, "c": 3},
			expected: true,
		},
		{
			name:     "zero values",
			subset:   map[string]int{"a": 0, "b": 0},
			superset: map[string]int{"a": 0, "b": 0, "c": 1},
			expected: true,
		},
		{
			name:     "negative values",
			subset:   map[string]int{"a": -1, "b": -2},
			superset: map[string]int{"a": -1, "b": -2, "c": 3},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.IsSubset(tt.subset, tt.superset)
			if result != tt.expected {
				t.Errorf("IsSubset() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsSubset_IntString(t *testing.T) {
	tests := []struct {
		name     string
		subset   map[int]string
		superset map[int]string
		expected bool
	}{
		{
			name:     "empty subset",
			subset:   map[int]string{},
			superset: map[int]string{1: "a", 2: "b"},
			expected: true,
		},
		{
			name:     "string subset true",
			subset:   map[int]string{1: "hello", 2: "world"},
			superset: map[int]string{1: "hello", 2: "world", 3: "go"},
			expected: true,
		},
		{
			name:     "string subset false value mismatch",
			subset:   map[int]string{1: "hello", 2: "universe"},
			superset: map[int]string{1: "hello", 2: "world", 3: "go"},
			expected: false,
		},
		{
			name:     "string subset false key missing",
			subset:   map[int]string{1: "hello", 4: "test"},
			superset: map[int]string{1: "hello", 2: "world"},
			expected: false,
		},
		{
			name:     "empty string values",
			subset:   map[int]string{1: "", 2: "x"},
			superset: map[int]string{1: "", 2: "x", 3: "y"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.IsSubset(tt.subset, tt.superset)
			if result != tt.expected {
				t.Errorf("IsSubset() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsSubset_StringBool(t *testing.T) {
	tests := []struct {
		name     string
		subset   map[string]bool
		superset map[string]bool
		expected bool
	}{
		{
			name:     "nil subset",
			subset:   nil,
			superset: map[string]bool{"a": true},
			expected: true,
		},
		{
			name:     "true values match",
			subset:   map[string]bool{"a": true, "b": true},
			superset: map[string]bool{"a": true, "b": true, "c": false},
			expected: true,
		},
		{
			name:     "false values match",
			subset:   map[string]bool{"a": false, "b": false},
			superset: map[string]bool{"a": false, "b": false, "c": true},
			expected: true,
		},
		{
			name:     "value mismatch true vs false",
			subset:   map[string]bool{"a": true},
			superset: map[string]bool{"a": false, "b": true},
			expected: false,
		},
		{
			name:     "key missing",
			subset:   map[string]bool{"x": true},
			superset: map[string]bool{"a": true, "b": false},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.IsSubset(tt.subset, tt.superset)
			if result != tt.expected {
				t.Errorf("IsSubset() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsSubset_StringFloat(t *testing.T) {
	tests := []struct {
		name     string
		subset   map[string]float64
		superset map[string]float64
		expected bool
	}{
		{
			name:     "float values match",
			subset:   map[string]float64{"pi": 3.14, "e": 2.71},
			superset: map[string]float64{"pi": 3.14, "e": 2.71, "sqrt2": 1.41},
			expected: true,
		},
		{
			name:     "float value mismatch",
			subset:   map[string]float64{"pi": 3.15},
			superset: map[string]float64{"pi": 3.14, "e": 2.71},
			expected: false,
		},
		{
			name:     "zero float",
			subset:   map[string]float64{"zero": 0.0},
			superset: map[string]float64{"zero": 0.0, "one": 1.0},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.IsSubset(tt.subset, tt.superset)
			if result != tt.expected {
				t.Errorf("IsSubset() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsSubset_StringString(t *testing.T) {
	tests := []struct {
		name     string
		subset   map[string]string
		superset map[string]string
		expected bool
	}{
		{
			name:     "empty subset empty superset",
			subset:   map[string]string{},
			superset: map[string]string{},
			expected: true,
		},
		{
			name:     "word pairs match",
			subset:   map[string]string{"a": "apple", "b": "banana"},
			superset: map[string]string{"a": "apple", "b": "banana", "c": "cherry"},
			expected: true,
		},
		{
			name:     "word pairs mismatch",
			subset:   map[string]string{"a": "ant"},
			superset: map[string]string{"a": "apple", "b": "banana"},
			expected: false,
		},
		{
			name:     "all empty strings",
			subset:   map[string]string{"a": "", "b": ""},
			superset: map[string]string{"a": "", "b": "", "c": "x"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.IsSubset(tt.subset, tt.superset)
			if result != tt.expected {
				t.Errorf("IsSubset() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsSubset_IntInt(t *testing.T) {
	tests := []struct {
		name     string
		subset   map[int]int
		superset map[int]int
		expected bool
	}{
		{
			name:     "nil both",
			subset:   nil,
			superset: nil,
			expected: true,
		},
		{
			name:     "numbers subset",
			subset:   map[int]int{1: 10, 2: 20},
			superset: map[int]int{1: 10, 2: 20, 3: 30},
			expected: true,
		},
		{
			name:     "negative keys",
			subset:   map[int]int{-1: 100, -2: 200},
			superset: map[int]int{-1: 100, -2: 200, 0: 0},
			expected: true,
		},
		{
			name:     "zero key and value",
			subset:   map[int]int{0: 0},
			superset: map[int]int{0: 0, 1: 1},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.IsSubset(tt.subset, tt.superset)
			if result != tt.expected {
				t.Errorf("IsSubset() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsSubset_LargeMap(t *testing.T) {
	tests := []struct {
		name     string
		check    func() bool
		checkMsg string
	}{
		{
			name: "large subset is subset of large superset",
			check: func() bool {
				subset := make(map[string]int)
				superset := make(map[string]int)

				// Create subset with 100 entries with unique keys
				for i := 0; i < 100; i++ {
					key := "k_" + string(rune(i/26+97)) + "_" + string(rune(i%26+97))
					subset[key] = i
					superset[key] = i // Also add to superset
				}

				// Add 100 more entries to superset
				for i := 100; i < 200; i++ {
					key := "k_" + string(rune(i/26+97)) + "_" + string(rune(i%26+97))
					superset[key] = i
				}

				return lxmaps.IsSubset(subset, superset)
			},
			checkMsg: "should handle large maps",
		},
		{
			name: "large subset not subset of smaller superset",
			check: func() bool {
				subset := make(map[string]int)
				superset := make(map[string]int)

				// Create subset with 100 entries
				for i := 0; i < 100; i++ {
					key := "key_" + string(rune(i))
					subset[key] = i
				}

				// Create superset with only 50 entries
				for i := 0; i < 50; i++ {
					key := "key_" + string(rune(i))
					superset[key] = i
				}

				return !lxmaps.IsSubset(subset, superset)
			},
			checkMsg: "should return false for large subset of small superset",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsSubset() %s", tt.checkMsg)
			}
		})
	}
}

func TestIsSubset_PropertyBased(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "every map is subset of itself",
			check: func() bool {
				m := map[string]int{"a": 1, "b": 2, "c": 3}
				return lxmaps.IsSubset(m, m)
			},
		},
		{
			name: "empty map is subset of any map",
			check: func() bool {
				empty := map[string]int{}
				m := map[string]int{"a": 1, "b": 2}
				return lxmaps.IsSubset(empty, m)
			},
		},
		{
			name: "subset of subset is subset of original",
			check: func() bool {
				// a is subset of b, b is subset of c, so a is subset of c
				a := map[string]int{"x": 1}
				b := map[string]int{"x": 1, "y": 2}
				c := map[string]int{"x": 1, "y": 2, "z": 3}

				return lxmaps.IsSubset(a, b) && lxmaps.IsSubset(b, c) && lxmaps.IsSubset(a, c)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsSubset() property-based test failed")
			}
		})
	}
}
