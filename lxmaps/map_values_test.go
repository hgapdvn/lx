package lxmaps_test

import (
	"testing"

	"github.com/hgapdvn/lx/lxmaps"
)

func TestMapValues_IntString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		fn       func(int) string
		check    func(map[string]string) bool
		checkMsg string
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(v int) string { return "x" },
			check:    func(result map[string]string) bool { return result == nil },
			checkMsg: "should return nil for nil input",
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			fn:       func(v int) string { return "x" },
			check:    func(result map[string]string) bool { return len(result) == 0 },
			checkMsg: "should return empty map for empty input",
		},
		{
			name:  "single entry",
			input: map[string]int{"a": 1},
			fn: func(v int) string {
				if v == 1 {
					return "one"
				}
				return "other"
			},
			check: func(result map[string]string) bool {
				return len(result) == 1 && result["a"] == "one"
			},
			checkMsg: "should transform single value",
		},
		{
			name: "multiple entries",
			input: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			fn: func(v int) string {
				switch v {
				case 1:
					return "one"
				case 2:
					return "two"
				case 3:
					return "three"
				default:
					return "other"
				}
			},
			check: func(result map[string]string) bool {
				return len(result) == 3 &&
					result["a"] == "one" &&
					result["b"] == "two" &&
					result["c"] == "three"
			},
			checkMsg: "should transform multiple values",
		},
		{
			name: "negative values",
			input: map[string]int{
				"neg":  -5,
				"pos":  10,
				"zero": 0,
			},
			fn: func(v int) string {
				if v < 0 {
					return "negative"
				} else if v == 0 {
					return "zero"
				}
				return "positive"
			},
			check: func(result map[string]string) bool {
				return len(result) == 3 &&
					result["neg"] == "negative" &&
					result["pos"] == "positive" &&
					result["zero"] == "zero"
			},
			checkMsg: "should handle negative values",
		},
		{
			name: "zero values",
			input: map[string]int{
				"x": 0,
				"y": 1,
			},
			fn: func(v int) string { return "value_" + string(rune(v+48)) },
			check: func(result map[string]string) bool {
				return len(result) == 2 && result["x"] != ""
			},
			checkMsg: "should handle zero values",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MapValues(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("MapValues() %s", tt.checkMsg)
			}
		})
	}
}

func TestMapValues_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		fn       func(string) int
		check    func(map[int]int) bool
		checkMsg string
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(v string) int { return 0 },
			check:    func(result map[int]int) bool { return result == nil },
			checkMsg: "should return nil",
		},
		{
			name:     "empty map",
			input:    map[int]string{},
			fn:       func(v string) int { return len(v) },
			check:    func(result map[int]int) bool { return len(result) == 0 },
			checkMsg: "should return empty map",
		},
		{
			name: "string to length",
			input: map[int]string{
				1: "a",
				2: "abc",
				3: "ab",
			},
			fn: func(v string) int { return len(v) },
			check: func(result map[int]int) bool {
				return len(result) == 3 &&
					result[1] == 1 &&
					result[2] == 3 &&
					result[3] == 2
			},
			checkMsg: "should map strings to their lengths",
		},
		{
			name: "string to hash",
			input: map[int]string{
				1: "key1",
				2: "key2",
			},
			fn: func(v string) int {
				sum := 0
				for _, ch := range v {
					sum += int(ch)
				}
				return sum
			},
			check: func(result map[int]int) bool {
				return len(result) == 2
			},
			checkMsg: "should hash string values",
		},
		{
			name: "zero int result",
			input: map[int]string{
				1: "zero",
				2: "one",
			},
			fn: func(v string) int { return 0 },
			check: func(result map[int]int) bool {
				return len(result) == 2
			},
			checkMsg: "should handle zero as result value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MapValues(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("MapValues() %s", tt.checkMsg)
			}
		})
	}
}

func TestMapValues_StringString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]string
		fn       func(string) string
		check    func(map[string]string) bool
		checkMsg string
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(v string) string { return v },
			check:    func(result map[string]string) bool { return result == nil },
			checkMsg: "should return nil",
		},
		{
			name:     "empty map",
			input:    map[string]string{},
			fn:       func(v string) string { return v },
			check:    func(result map[string]string) bool { return len(result) == 0 },
			checkMsg: "should return empty map",
		},
		{
			name: "uppercase values",
			input: map[string]string{
				"greeting": "hello",
				"name":     "world",
			},
			fn: func(v string) string {
				if v == "hello" {
					return "HELLO"
				} else if v == "world" {
					return "WORLD"
				}
				return v
			},
			check: func(result map[string]string) bool {
				return len(result) == 2 &&
					result["greeting"] == "HELLO" &&
					result["name"] == "WORLD"
			},
			checkMsg: "should uppercase values",
		},
		{
			name: "prefix values",
			input: map[string]string{
				"a": "alpha",
				"b": "beta",
			},
			fn: func(v string) string { return "value_" + v },
			check: func(result map[string]string) bool {
				return len(result) == 2 &&
					result["a"] == "value_alpha" &&
					result["b"] == "value_beta"
			},
			checkMsg: "should add prefix to values",
		},
		{
			name: "empty string values",
			input: map[string]string{
				"key1": "",
				"key2": "value",
			},
			fn: func(v string) string { return "x_" + v },
			check: func(result map[string]string) bool {
				return len(result) == 2
			},
			checkMsg: "should handle empty strings",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MapValues(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("MapValues() %s", tt.checkMsg)
			}
		})
	}
}

func TestMapValues_IntInt(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		fn       func(int) int
		check    func(map[string]int) bool
		checkMsg string
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(v int) int { return 0 },
			check:    func(result map[string]int) bool { return result == nil },
			checkMsg: "should return nil",
		},
		{
			name: "negate values",
			input: map[string]int{
				"a":    1,
				"b":    -2,
				"zero": 0,
			},
			fn: func(v int) int { return -v },
			check: func(result map[string]int) bool {
				return len(result) == 3 &&
					result["a"] == -1 &&
					result["b"] == 2 &&
					result["zero"] == 0
			},
			checkMsg: "should negate integer values",
		},
		{
			name: "square values",
			input: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
			fn: func(v int) int { return v * v },
			check: func(result map[string]int) bool {
				return len(result) == 3 &&
					result["one"] == 1 &&
					result["two"] == 4 &&
					result["three"] == 9
			},
			checkMsg: "should square integer values",
		},
		{
			name: "modulo transformation",
			input: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
				"d": 4,
				"e": 5,
			},
			fn: func(v int) int { return v % 3 },
			check: func(result map[string]int) bool {
				return len(result) == 5
			},
			checkMsg: "should apply modulo transformation",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MapValues(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("MapValues() %s", tt.checkMsg)
			}
		})
	}
}

func TestMapValues_FloatString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]float64
		fn       func(float64) string
		check    func(map[string]string) bool
		checkMsg string
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(v float64) string { return "x" },
			check:    func(result map[string]string) bool { return result == nil },
			checkMsg: "should return nil",
		},
		{
			name: "float to string",
			input: map[string]float64{
				"pi":  3.14159,
				"e":   2.71828,
				"phi": 1.61803,
			},
			fn: func(v float64) string { return "value" },
			check: func(result map[string]string) bool {
				return len(result) == 3
			},
			checkMsg: "should convert floats to strings",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MapValues(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("MapValues() %s", tt.checkMsg)
			}
		})
	}
}

func TestMapValues_BoolString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]bool
		fn       func(bool) string
		check    func(map[string]string) bool
		checkMsg string
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(v bool) string { return "x" },
			check:    func(result map[string]string) bool { return result == nil },
			checkMsg: "should return nil",
		},
		{
			name: "bool to string",
			input: map[string]bool{
				"enabled":  true,
				"disabled": false,
			},
			fn: func(v bool) string {
				if v {
					return "yes"
				}
				return "no"
			},
			check: func(result map[string]string) bool {
				return len(result) == 2 &&
					result["enabled"] == "yes" &&
					result["disabled"] == "no"
			},
			checkMsg: "should convert bool to string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MapValues(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("MapValues() %s", tt.checkMsg)
			}
		})
	}
}

func TestMapValues_PreservesKeys(t *testing.T) {
	tests := []struct {
		name     string
		check    func() bool
		checkMsg string
	}{
		{
			name: "keys unchanged",
			check: func() bool {
				input := map[string]int{"a": 1, "b": 2, "c": 3}
				result := lxmaps.MapValues(input, func(v int) int { return v * 10 })
				for k := range input {
					if _, ok := result[k]; !ok {
						return false
					}
				}
				return true
			},
			checkMsg: "all keys should be preserved",
		},
		{
			name: "keys preserved with empty values",
			check: func() bool {
				input := map[string]string{"key1": "", "key2": "x"}
				result := lxmaps.MapValues(input, func(v string) string { return "transformed" })
				return len(result) == 2 && result["key1"] != "" && result["key2"] != ""
			},
			checkMsg: "should preserve keys even with empty values",
		},
		{
			name: "numeric keys preserved",
			check: func() bool {
				input := map[int]int{1: 10, 2: 20, 3: 30}
				result := lxmaps.MapValues(input, func(v int) int { return v / 10 })
				return len(result) == 3 &&
					result[1] == 1 &&
					result[2] == 2 &&
					result[3] == 3
			},
			checkMsg: "should preserve numeric keys",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("MapValues() %s", tt.checkMsg)
			}
		})
	}
}

func TestMapValues_CustomStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	type PersonInfo struct {
		Description string
	}

	tests := []struct {
		name     string
		input    map[string]Person
		fn       func(Person) PersonInfo
		check    func(map[string]PersonInfo) bool
		checkMsg string
	}{
		{
			name: "struct to struct",
			input: map[string]Person{
				"alice": {Name: "Alice", Age: 30},
				"bob":   {Name: "Bob", Age: 25},
			},
			fn: func(p Person) PersonInfo {
				return PersonInfo{Description: p.Name}
			},
			check: func(result map[string]PersonInfo) bool {
				return len(result) == 2 &&
					result["alice"].Description == "Alice" &&
					result["bob"].Description == "Bob"
			},
			checkMsg: "should transform struct values",
		},
		{
			name: "struct with zero values",
			input: map[string]Person{
				"person1": {Name: "", Age: 0},
				"person2": {Name: "Test", Age: 20},
			},
			fn: func(p Person) PersonInfo {
				return PersonInfo{Description: p.Name}
			},
			check: func(result map[string]PersonInfo) bool {
				return len(result) == 2
			},
			checkMsg: "should transform struct with zero values",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MapValues(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("MapValues() %s", tt.checkMsg)
			}
		})
	}
}

func TestMapValues_LargeMap(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		check    func(map[int]int) bool
		checkMsg string
	}{
		{
			name: "1000 entries",
			size: 1000,
			check: func(result map[int]int) bool {
				return len(result) == 1000
			},
			checkMsg: "should handle 1000 entries",
		},
		{
			name: "100 entries",
			size: 100,
			check: func(result map[int]int) bool {
				return len(result) == 100
			},
			checkMsg: "should handle 100 entries",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make(map[int]int)
			for i := 0; i < tt.size; i++ {
				input[i] = i * 10
			}

			result := lxmaps.MapValues(input, func(v int) int {
				return v * 2
			})
			if !tt.check(result) {
				t.Errorf("MapValues() %s, got %d entries", tt.checkMsg, len(result))
			}
		})
	}
}

func TestMapValues_ConsistencyWithSize(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "size preserved",
			check: func() bool {
				input := map[int]int{1: 10, 2: 20, 3: 30}
				result := lxmaps.MapValues(input, func(v int) int { return v * 2 })
				return len(result) == len(input)
			},
		},
		{
			name: "size preserved with transformation",
			check: func() bool {
				input := map[string]string{"a": "x", "b": "y", "c": "z"}
				result := lxmaps.MapValues(input, func(v string) string { return "prefix_" + v })
				return len(result) == len(input)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("MapValues() consistency check failed")
			}
		})
	}
}
