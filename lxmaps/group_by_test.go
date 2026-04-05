package lxmaps_test

import (
	"testing"

	lxmaps "github.com/hgapdvn/lx/lxmaps"
)

func TestGroupBy_StringIntByEvenOdd(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		fn       func(string, int) string
		check    func(map[string]map[string]int) bool
		checkMsg string
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(k string, v int) string { return "any" },
			check:    func(result map[string]map[string]int) bool { return result == nil },
			checkMsg: "should return nil",
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			fn:       func(k string, v int) string { return "any" },
			check:    func(result map[string]map[string]int) bool { return len(result) == 0 },
			checkMsg: "should return empty map",
		},
		{
			name:  "group by even/odd",
			input: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			fn: func(k string, v int) string {
				if v%2 == 0 {
					return "even"
				}
				return "odd"
			},
			check: func(result map[string]map[string]int) bool {
				return len(result) == 2 &&
					len(result["even"]) == 2 &&
					len(result["odd"]) == 2 &&
					result["even"]["b"] == 2 &&
					result["even"]["d"] == 4 &&
					result["odd"]["a"] == 1 &&
					result["odd"]["c"] == 3
			},
			checkMsg: "should group even and odd values correctly",
		},
		{
			name:  "single group",
			input: map[string]int{"a": 2, "b": 4, "c": 6},
			fn: func(k string, v int) string {
				if v%2 == 0 {
					return "even"
				}
				return "odd"
			},
			check: func(result map[string]map[string]int) bool {
				return len(result) == 1 && len(result["even"]) == 3
			},
			checkMsg: "should create single group",
		},
		{
			name:  "group by key length",
			input: map[string]int{"a": 1, "ab": 2, "abc": 3, "x": 4, "xy": 5},
			fn: func(k string, v int) string {
				switch len(k) {
				case 1:
					return "short"
				case 2:
					return "medium"
				default:
					return "long"
				}
			},
			check: func(result map[string]map[string]int) bool {
				return len(result) == 3 &&
					len(result["short"]) == 2 &&
					len(result["medium"]) == 2 &&
					len(result["long"]) == 1
			},
			checkMsg: "should group by key length",
		},
		{
			name:  "group by value range",
			input: map[string]int{"a": 5, "b": 15, "c": 25, "d": 35, "e": 45},
			fn: func(k string, v int) string {
				if v < 20 {
					return "low"
				} else if v < 40 {
					return "mid"
				}
				return "high"
			},
			check: func(result map[string]map[string]int) bool {
				return len(result) == 3 &&
					len(result["low"]) == 2 &&
					len(result["mid"]) == 2 &&
					len(result["high"]) == 1
			},
			checkMsg: "should group by value ranges",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.GroupBy(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("GroupBy() %s", tt.checkMsg)
			}
		})
	}
}

func TestGroupBy_IntStringByLength(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		fn       func(int, string) int
		check    func(map[int]map[int]string) bool
		checkMsg string
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(k int, v string) int { return 0 },
			check:    func(result map[int]map[int]string) bool { return result == nil },
			checkMsg: "should return nil",
		},
		{
			name:     "empty map",
			input:    map[int]string{},
			fn:       func(k int, v string) int { return len(v) },
			check:    func(result map[int]map[int]string) bool { return len(result) == 0 },
			checkMsg: "should return empty map",
		},
		{
			name: "group by string length",
			input: map[int]string{
				1: "a",
				2: "ab",
				3: "abc",
				4: "ab",
				5: "abcde",
			},
			fn: func(k int, v string) int { return len(v) },
			check: func(result map[int]map[int]string) bool {
				return len(result) == 4 &&
					len(result[1]) == 1 &&
					len(result[2]) == 2 &&
					len(result[3]) == 1 &&
					len(result[5]) == 1 &&
					result[2][2] == "ab" &&
					result[2][4] == "ab"
			},
			checkMsg: "should group by string length",
		},
		{
			name: "group by key value",
			input: map[int]string{
				1: "x", 2: "x", 3: "y", 4: "y", 5: "z",
			},
			fn: func(k int, v string) int {
				if k%2 == 0 {
					return 0
				}
				return 1
			},
			check: func(result map[int]map[int]string) bool {
				return len(result) == 2 &&
					len(result[0]) == 2 &&
					len(result[1]) == 3
			},
			checkMsg: "should group by key parity",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.GroupBy(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("GroupBy() %s", tt.checkMsg)
			}
		})
	}
}

func TestGroupBy_StringBoolByValue(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]bool
		fn       func(string, bool) bool
		check    func(map[bool]map[string]bool) bool
		checkMsg string
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(k string, v bool) bool { return v },
			check:    func(result map[bool]map[string]bool) bool { return result == nil },
			checkMsg: "should return nil",
		},
		{
			name:     "empty map",
			input:    map[string]bool{},
			fn:       func(k string, v bool) bool { return v },
			check:    func(result map[bool]map[string]bool) bool { return len(result) == 0 },
			checkMsg: "should return empty map",
		},
		{
			name: "group by value",
			input: map[string]bool{
				"a": true,
				"b": false,
				"c": true,
				"d": false,
			},
			fn: func(k string, v bool) bool { return v },
			check: func(result map[bool]map[string]bool) bool {
				return len(result) == 2 &&
					len(result[true]) == 2 &&
					len(result[false]) == 2 &&
					result[true]["a"] == true &&
					result[true]["c"] == true &&
					result[false]["b"] == false &&
					result[false]["d"] == false
			},
			checkMsg: "should group by boolean value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.GroupBy(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("GroupBy() %s", tt.checkMsg)
			}
		})
	}
}

func TestGroupBy_CustomStruct(t *testing.T) {
	type Item struct {
		Name     string
		Category string
		Price    float64
	}

	tests := []struct {
		name     string
		input    map[string]Item
		fn       func(string, Item) string
		check    func(map[string]map[string]Item) bool
		checkMsg string
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(k string, v Item) string { return v.Category },
			check:    func(result map[string]map[string]Item) bool { return result == nil },
			checkMsg: "should return nil",
		},
		{
			name: "group by category",
			input: map[string]Item{
				"item1": {Name: "Apple", Category: "Fruit", Price: 1.5},
				"item2": {Name: "Carrot", Category: "Vegetable", Price: 0.8},
				"item3": {Name: "Banana", Category: "Fruit", Price: 0.6},
				"item4": {Name: "Lettuce", Category: "Vegetable", Price: 1.2},
			},
			fn: func(k string, v Item) string { return v.Category },
			check: func(result map[string]map[string]Item) bool {
				return len(result) == 2 &&
					len(result["Fruit"]) == 2 &&
					len(result["Vegetable"]) == 2 &&
					result["Fruit"]["item1"].Name == "Apple" &&
					result["Fruit"]["item3"].Name == "Banana" &&
					result["Vegetable"]["item2"].Name == "Carrot" &&
					result["Vegetable"]["item4"].Name == "Lettuce"
			},
			checkMsg: "should group struct by category",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.GroupBy(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("GroupBy() %s", tt.checkMsg)
			}
		})
	}
}

func TestGroupBy_StringStringByFirstChar(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]string
		fn       func(string, string) string
		check    func(map[string]map[string]string) bool
		checkMsg string
	}{
		{
			name: "group by first character",
			input: map[string]string{
				"k1": "apple",
				"k2": "apricot",
				"k3": "banana",
				"k4": "avocado",
				"k5": "blueberry",
			},
			fn: func(k string, v string) string {
				if len(v) == 0 {
					return ""
				}
				return string(v[0])
			},
			check: func(result map[string]map[string]string) bool {
				return len(result) == 2 &&
					len(result["a"]) == 3 &&
					len(result["b"]) == 2 &&
					result["a"]["k1"] == "apple" &&
					result["a"]["k2"] == "apricot" &&
					result["a"]["k4"] == "avocado" &&
					result["b"]["k3"] == "banana" &&
					result["b"]["k5"] == "blueberry"
			},
			checkMsg: "should group by first character",
		},
		{
			name:     "empty map",
			input:    map[string]string{},
			fn:       func(k string, v string) string { return "any" },
			check:    func(result map[string]map[string]string) bool { return len(result) == 0 },
			checkMsg: "should handle empty map",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.GroupBy(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("GroupBy() %s", tt.checkMsg)
			}
		})
	}
}

func TestGroupBy_IntIntByQuotient(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]int
		fn       func(int, int) int
		check    func(map[int]map[int]int) bool
		checkMsg string
	}{
		{
			name: "group by integer division",
			input: map[int]int{
				1: 5, 2: 15, 3: 7, 4: 20, 5: 3, 6: 22,
			},
			fn: func(k int, v int) int {
				return v / 10
			},
			check: func(result map[int]map[int]int) bool {
				return len(result) == 3 &&
					len(result[0]) == 3 && // 5, 7, 3
					len(result[1]) == 1 && // 15
					len(result[2]) == 2 // 20, 22
			},
			checkMsg: "should group by value/10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.GroupBy(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("GroupBy() %s", tt.checkMsg)
			}
		})
	}
}

func TestGroupBy_LargeMap(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		check    func(map[int]map[int]int) bool
		checkMsg string
	}{
		{
			name: "large map with 1000 entries grouped by modulo 10",
			size: 1000,
			check: func(result map[int]map[int]int) bool {
				// Should have 10 groups (0-9)
				if len(result) != 10 {
					return false
				}
				// Each group should have 100 entries
				for i := 0; i < 10; i++ {
					if len(result[i]) != 100 {
						return false
					}
				}
				return true
			},
			checkMsg: "should group 1000 entries into 10 groups of 100",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := make(map[int]int)
			for i := 0; i < tt.size; i++ {
				m[i] = i
			}

			result := lxmaps.GroupBy(m, func(k int, v int) int {
				return v % 10
			})

			if !tt.check(result) {
				t.Errorf("GroupBy() %s", tt.checkMsg)
			}
		})
	}
}
