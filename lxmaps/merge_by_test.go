package lxmaps_test

import (
	"testing"

	lxmaps "github.com/hgapdvn/lx/lxmaps"
)

func TestMergeBy_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		m1       map[string]int
		m2       map[string]int
		fn       func(int, int) int
		check    func(map[string]int) bool
		checkMsg string
	}{
		{
			name:     "nil m1",
			m1:       nil,
			m2:       map[string]int{"a": 1, "b": 2},
			fn:       func(a, b int) int { return a + b },
			check:    func(result map[string]int) bool { return len(result) == 2 && result["a"] == 1 && result["b"] == 2 },
			checkMsg: "should handle nil m1",
		},
		{
			name:     "nil m2",
			m1:       map[string]int{"a": 1, "b": 2},
			m2:       nil,
			fn:       func(a, b int) int { return a + b },
			check:    func(result map[string]int) bool { return len(result) == 2 && result["a"] == 1 && result["b"] == 2 },
			checkMsg: "should handle nil m2",
		},
		{
			name:     "both nil",
			m1:       nil,
			m2:       nil,
			fn:       func(a, b int) int { return a + b },
			check:    func(result map[string]int) bool { return len(result) == 0 },
			checkMsg: "should return empty map for both nil",
		},
		{
			name:     "empty m1",
			m1:       map[string]int{},
			m2:       map[string]int{"a": 1},
			fn:       func(a, b int) int { return a + b },
			check:    func(result map[string]int) bool { return len(result) == 1 && result["a"] == 1 },
			checkMsg: "should handle empty m1",
		},
		{
			name:     "empty m2",
			m1:       map[string]int{"a": 1},
			m2:       map[string]int{},
			fn:       func(a, b int) int { return a + b },
			check:    func(result map[string]int) bool { return len(result) == 1 && result["a"] == 1 },
			checkMsg: "should handle empty m2",
		},
		{
			name: "no overlap",
			m1:   map[string]int{"a": 1, "b": 2},
			m2:   map[string]int{"c": 3, "d": 4},
			fn:   func(a, b int) int { return a + b },
			check: func(result map[string]int) bool {
				return len(result) == 4 && result["a"] == 1 && result["b"] == 2 && result["c"] == 3 && result["d"] == 4
			},
			checkMsg: "should merge without conflicts",
		},
		{
			name: "with overlap sum",
			m1:   map[string]int{"a": 1, "b": 2},
			m2:   map[string]int{"b": 3, "c": 4},
			fn:   func(a, b int) int { return a + b },
			check: func(result map[string]int) bool {
				return len(result) == 3 && result["a"] == 1 && result["b"] == 5 && result["c"] == 4
			},
			checkMsg: "should sum overlapping values",
		},
		{
			name: "last value wins",
			m1:   map[string]int{"key": 1},
			m2:   map[string]int{"key": 2},
			fn:   func(a, b int) int { return b },
			check: func(result map[string]int) bool {
				return len(result) == 1 && result["key"] == 2
			},
			checkMsg: "should use last value",
		},
		{
			name: "keep max",
			m1:   map[string]int{"a": 5, "b": 2},
			m2:   map[string]int{"b": 8, "c": 3},
			fn: func(a, b int) int {
				if a > b {
					return a
				}
				return b
			},
			check: func(result map[string]int) bool {
				return len(result) == 3 && result["a"] == 5 && result["b"] == 8 && result["c"] == 3
			},
			checkMsg: "should keep max value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MergeBy(tt.m1, tt.m2, tt.fn)
			if !tt.check(result) {
				t.Errorf("MergeBy() %s", tt.checkMsg)
			}
		})
	}
}

func TestMergeBy_IntString(t *testing.T) {
	tests := []struct {
		name     string
		m1       map[int]string
		m2       map[int]string
		fn       func(string, string) string
		check    func(map[int]string) bool
		checkMsg string
	}{
		{
			name: "concatenate strings",
			m1:   map[int]string{1: "hello", 2: "world"},
			m2:   map[int]string{2: " from", 3: " Go"},
			fn:   func(a, b string) string { return a + b },
			check: func(result map[int]string) bool {
				return len(result) == 3 && result[1] == "hello" && result[2] == "world from" && result[3] == " Go"
			},
			checkMsg: "should concatenate strings",
		},
		{
			name: "keep longer",
			m1:   map[int]string{1: "short", 2: "medium"},
			m2:   map[int]string{2: "longer string", 3: "x"},
			fn: func(a, b string) string {
				if len(a) > len(b) {
					return a
				}
				return b
			},
			check: func(result map[int]string) bool {
				return len(result) == 3 && result[1] == "short" && result[2] == "longer string" && result[3] == "x"
			},
			checkMsg: "should keep longer string",
		},
		{
			name: "nil m1",
			m1:   nil,
			m2:   map[int]string{1: "test"},
			fn:   func(a, b string) string { return a + b },
			check: func(result map[int]string) bool {
				return len(result) == 1 && result[1] == "test"
			},
			checkMsg: "should handle nil m1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MergeBy(tt.m1, tt.m2, tt.fn)
			if !tt.check(result) {
				t.Errorf("MergeBy() %s", tt.checkMsg)
			}
		})
	}
}

func TestMergeBy_StringBool(t *testing.T) {
	tests := []struct {
		name     string
		m1       map[string]bool
		m2       map[string]bool
		fn       func(bool, bool) bool
		check    func(map[string]bool) bool
		checkMsg string
	}{
		{
			name: "AND logic",
			m1:   map[string]bool{"a": true, "b": false, "c": true},
			m2:   map[string]bool{"b": true, "c": true},
			fn:   func(a, b bool) bool { return a && b },
			check: func(result map[string]bool) bool {
				return len(result) == 3 && result["a"] == true && result["b"] == false && result["c"] == true
			},
			checkMsg: "should AND values correctly",
		},
		{
			name: "OR logic",
			m1:   map[string]bool{"a": true, "b": false},
			m2:   map[string]bool{"b": false, "c": true},
			fn:   func(a, b bool) bool { return a || b },
			check: func(result map[string]bool) bool {
				return len(result) == 3 && result["a"] == true && result["b"] == false && result["c"] == true
			},
			checkMsg: "should OR values correctly",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MergeBy(tt.m1, tt.m2, tt.fn)
			if !tt.check(result) {
				t.Errorf("MergeBy() %s", tt.checkMsg)
			}
		})
	}
}

func TestMergeBy_CustomStruct(t *testing.T) {
	type Item struct {
		Name  string
		Count int
	}

	tests := []struct {
		name     string
		m1       map[string]Item
		m2       map[string]Item
		fn       func(Item, Item) Item
		check    func(map[string]Item) bool
		checkMsg string
	}{
		{
			name: "sum counts",
			m1: map[string]Item{
				"item1": {Name: "apple", Count: 5},
				"item2": {Name: "banana", Count: 3},
			},
			m2: map[string]Item{
				"item2": {Name: "banana", Count: 2},
				"item3": {Name: "cherry", Count: 1},
			},
			fn: func(a, b Item) Item {
				return Item{Name: a.Name, Count: a.Count + b.Count}
			},
			check: func(result map[string]Item) bool {
				return len(result) == 3 && result["item1"].Count == 5 && result["item2"].Count == 5 && result["item3"].Count == 1
			},
			checkMsg: "should merge structs",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MergeBy(tt.m1, tt.m2, tt.fn)
			if !tt.check(result) {
				t.Errorf("MergeBy() %s", tt.checkMsg)
			}
		})
	}
}

func TestMergeBy_ComplexMergeLogic(t *testing.T) {
	tests := []struct {
		name     string
		m1       map[string]int
		m2       map[string]int
		fn       func(int, int) int
		check    func(map[string]int) bool
		checkMsg string
	}{
		{
			name: "average",
			m1:   map[string]int{"a": 10, "b": 20},
			m2:   map[string]int{"a": 20, "b": 40},
			fn:   func(a, b int) int { return (a + b) / 2 },
			check: func(result map[string]int) bool {
				return len(result) == 2 && result["a"] == 15 && result["b"] == 30
			},
			checkMsg: "should average on merge",
		},
		{
			name: "product",
			m1:   map[string]int{"a": 2, "b": 3},
			m2:   map[string]int{"a": 4, "b": 5},
			fn:   func(a, b int) int { return a * b },
			check: func(result map[string]int) bool {
				return len(result) == 2 && result["a"] == 8 && result["b"] == 15
			},
			checkMsg: "should multiply on merge",
		},
		{
			name: "max difference",
			m1:   map[string]int{"a": 10, "b": 5},
			m2:   map[string]int{"a": 3, "b": 12},
			fn: func(a, b int) int {
				diff := a - b
				if diff < 0 {
					return -diff
				}
				return diff
			},
			check: func(result map[string]int) bool {
				return len(result) == 2 && result["a"] == 7 && result["b"] == 7
			},
			checkMsg: "should compute absolute difference",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MergeBy(tt.m1, tt.m2, tt.fn)
			if !tt.check(result) {
				t.Errorf("MergeBy() %s", tt.checkMsg)
			}
		})
	}
}

func TestMergeBy_DoesNotModifyInput(t *testing.T) {
	tests := []struct {
		name string
		m1   map[string]int
		m2   map[string]int
	}{
		{
			name: "m1 unchanged after merge",
			m1:   map[string]int{"a": 1, "b": 2},
			m2:   map[string]int{"b": 3, "c": 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m1Copy := make(map[string]int)
			for k, v := range tt.m1 {
				m1Copy[k] = v
			}

			lxmaps.MergeBy(tt.m1, tt.m2, func(a, b int) int { return a + b })

			// Verify m1 is unchanged
			if len(tt.m1) != len(m1Copy) {
				t.Errorf("MergeBy() modified m1 length")
			}
			for k, v := range m1Copy {
				if tt.m1[k] != v {
					t.Errorf("MergeBy() modified m1")
				}
			}
		})
	}
}

func TestMergeBy_StringStringEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		m1       map[string]string
		m2       map[string]string
		fn       func(string, string) string
		check    func(map[string]string) bool
		checkMsg string
	}{
		{
			name: "empty string values",
			m1:   map[string]string{"a": "", "b": "x"},
			m2:   map[string]string{"b": "y", "c": ""},
			fn:   func(a, b string) string { return a + b },
			check: func(result map[string]string) bool {
				return len(result) == 3 && result["a"] == "" && result["b"] == "xy" && result["c"] == ""
			},
			checkMsg: "should handle empty string values",
		},
		{
			name: "all empty strings",
			m1:   map[string]string{"a": "", "b": ""},
			m2:   map[string]string{"b": "", "c": ""},
			fn:   func(a, b string) string { return a + b },
			check: func(result map[string]string) bool {
				return len(result) == 3 && result["a"] == "" && result["b"] == "" && result["c"] == ""
			},
			checkMsg: "should handle all empty strings",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MergeBy(tt.m1, tt.m2, tt.fn)
			if !tt.check(result) {
				t.Errorf("MergeBy() %s", tt.checkMsg)
			}
		})
	}
}
