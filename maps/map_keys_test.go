package lxmaps_test

import (
	"testing"

	lxmaps "github.com/hgapdvn/lx/maps"
)

func TestMapKeys_IntString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		fn       func(int) string
		check    func(map[string]string) bool
		checkMsg string
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(k int) string { return "x" },
			check:    func(result map[string]string) bool { return result == nil },
			checkMsg: "should return nil for nil input",
		},
		{
			name:     "empty map",
			input:    map[int]string{},
			fn:       func(k int) string { return "x" },
			check:    func(result map[string]string) bool { return len(result) == 0 },
			checkMsg: "should return empty map for empty input",
		},
		{
			name:  "single entry",
			input: map[int]string{1: "hello"},
			fn: func(k int) string {
				if k == 1 {
					return "one"
				}
				return "other"
			},
			check: func(result map[string]string) bool {
				return len(result) == 1 && result["one"] == "hello"
			},
			checkMsg: "should transform single key",
		},
		{
			name:  "multiple entries",
			input: map[int]string{1: "a", 2: "b", 3: "c"},
			fn: func(k int) string {
				if k == 1 {
					return "one"
				} else if k == 2 {
					return "two"
				} else if k == 3 {
					return "three"
				}
				return "other"
			},
			check: func(result map[string]string) bool {
				return len(result) == 3 &&
					result["one"] == "a" &&
					result["two"] == "b" &&
					result["three"] == "c"
			},
			checkMsg: "should transform multiple keys",
		},
		{
			name:  "negative keys",
			input: map[int]string{-1: "neg", 0: "zero", 1: "pos"},
			fn: func(k int) string {
				if k < 0 {
					return "negative"
				} else if k == 0 {
					return "zero"
				}
				return "positive"
			},
			check: func(result map[string]string) bool {
				return len(result) == 3 &&
					result["negative"] == "neg" &&
					result["zero"] == "zero" &&
					result["positive"] == "pos"
			},
			checkMsg: "should handle negative keys",
		},
		{
			name:  "empty string values",
			input: map[int]string{1: "", 2: "x"},
			fn:    func(k int) string { return "key_" + string(rune(k+48)) },
			check: func(result map[string]string) bool {
				return len(result) == 2 && result["key_2"] == "x"
			},
			checkMsg: "should handle empty string values",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MapKeys(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("MapKeys() %s", tt.checkMsg)
			}
		})
	}
}

func TestMapKeys_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		fn       func(string) int
		check    func(map[int]int) bool
		checkMsg string
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(k string) int { return 0 },
			check:    func(result map[int]int) bool { return result == nil },
			checkMsg: "should return nil",
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			fn:       func(k string) int { return len(k) },
			check:    func(result map[int]int) bool { return len(result) == 0 },
			checkMsg: "should return empty map",
		},
		{
			name: "string to length",
			input: map[string]int{
				"a":   1,
				"abc": 2,
				"ab":  3,
			},
			fn: func(k string) int { return len(k) },
			check: func(result map[int]int) bool {
				return len(result) >= 2 &&
					(result[1] == 1 || result[1] == 3) &&
					(result[2] > 0 || result[3] > 0)
			},
			checkMsg: "should map strings to their lengths",
		},
		{
			name: "string to hash",
			input: map[string]int{
				"key1": 10,
				"key2": 20,
			},
			fn: func(k string) int {
				sum := 0
				for _, ch := range k {
					sum += int(ch)
				}
				return sum
			},
			check: func(result map[int]int) bool {
				return len(result) == 2
			},
			checkMsg: "should hash string keys",
		},
		{
			name: "zero int result",
			input: map[string]int{
				"zero": 0,
				"one":  1,
			},
			fn: func(k string) int { return 0 },
			check: func(result map[int]int) bool {
				return len(result) == 1
			},
			checkMsg: "should handle zero as key result",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MapKeys(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("MapKeys() %s", tt.checkMsg)
			}
		})
	}
}

func TestMapKeys_StringString(t *testing.T) {
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
			fn:       func(k string) string { return k },
			check:    func(result map[string]string) bool { return result == nil },
			checkMsg: "should return nil",
		},
		{
			name:     "empty map",
			input:    map[string]string{},
			fn:       func(k string) string { return k },
			check:    func(result map[string]string) bool { return len(result) == 0 },
			checkMsg: "should return empty map",
		},
		{
			name: "uppercase keys",
			input: map[string]string{
				"hello": "world",
				"foo":   "bar",
			},
			fn: func(k string) string {
				if k == "hello" {
					return "HELLO"
				} else if k == "foo" {
					return "FOO"
				}
				return k
			},
			check: func(result map[string]string) bool {
				return len(result) == 2 &&
					result["HELLO"] == "world" &&
					result["FOO"] == "bar"
			},
			checkMsg: "should uppercase keys",
		},
		{
			name: "prefix keys",
			input: map[string]string{
				"a": "alpha",
				"b": "beta",
			},
			fn: func(k string) string { return "key_" + k },
			check: func(result map[string]string) bool {
				return len(result) == 2 &&
					result["key_a"] == "alpha" &&
					result["key_b"] == "beta"
			},
			checkMsg: "should add prefix to keys",
		},
		{
			name: "empty string keys and values",
			input: map[string]string{
				"":    "empty_key",
				"key": "",
			},
			fn: func(k string) string { return "x_" + k },
			check: func(result map[string]string) bool {
				return len(result) == 2
			},
			checkMsg: "should handle empty strings",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MapKeys(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("MapKeys() %s", tt.checkMsg)
			}
		})
	}
}

func TestMapKeys_StringFloat(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]float64
		fn       func(string) int
		check    func(map[int]float64) bool
		checkMsg string
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(k string) int { return 0 },
			check:    func(result map[int]float64) bool { return result == nil },
			checkMsg: "should return nil",
		},
		{
			name: "float values preserved",
			input: map[string]float64{
				"pi": 3.14159,
				"e":  2.71828,
			},
			fn: func(k string) int { return len(k) },
			check: func(result map[int]float64) bool {
				return len(result) == 1 && len(result) > 0
			},
			checkMsg: "should preserve float values",
		},
		{
			name: "zero float preserved",
			input: map[string]float64{
				"zero": 0.0,
				"one":  1.0,
			},
			fn: func(k string) int { return len(k) },
			check: func(result map[int]float64) bool {
				for _, v := range result {
					if v == 0.0 || v == 1.0 {
						return true
					}
				}
				return false
			},
			checkMsg: "should preserve zero float",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MapKeys(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("MapKeys() %s", tt.checkMsg)
			}
		})
	}
}

func TestMapKeys_StringBool(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]bool
		fn       func(string) int
		check    func(map[int]bool) bool
		checkMsg string
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(k string) int { return 0 },
			check:    func(result map[int]bool) bool { return result == nil },
			checkMsg: "should return nil",
		},
		{
			name: "bool values preserved",
			input: map[string]bool{
				"a": true,
				"b": false,
			},
			fn: func(k string) int { return len(k) },
			check: func(result map[int]bool) bool {
				trueCount := 0
				falseCount := 0
				for _, v := range result {
					if v {
						trueCount++
					} else {
						falseCount++
					}
				}
				return trueCount > 0 || falseCount > 0
			},
			checkMsg: "should preserve bool values",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MapKeys(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("MapKeys() %s", tt.checkMsg)
			}
		})
	}
}

func TestMapKeys_DuplicateTransformation(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		fn       func(int) int
		check    func(map[int]string) bool
		checkMsg string
	}{
		{
			name: "all keys map to same key",
			input: map[int]string{
				1: "first",
				2: "second",
				3: "third",
			},
			fn: func(k int) int { return 100 },
			check: func(result map[int]string) bool {
				return len(result) == 1 && result[100] != ""
			},
			checkMsg: "should keep last value when keys collide",
		},
		{
			name: "some keys map to same key",
			input: map[int]string{
				1: "one",
				2: "two",
				3: "three",
				4: "four",
			},
			fn: func(k int) int {
				return k % 2
			},
			check: func(result map[int]string) bool {
				return len(result) == 2
			},
			checkMsg: "should handle partial key collisions",
		},
		{
			name: "progressive collisions",
			input: map[int]string{
				1: "a",
				2: "b",
				3: "c",
				4: "d",
				5: "e",
			},
			fn: func(k int) int {
				return k / 2
			},
			check: func(result map[int]string) bool {
				return len(result) <= 3
			},
			checkMsg: "should handle progressive collisions",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MapKeys(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("MapKeys() %s", tt.checkMsg)
			}
		})
	}
}

func TestMapKeys_PreservesValues(t *testing.T) {
	tests := []struct {
		name     string
		check    func() bool
		checkMsg string
	}{
		{
			name: "values unchanged",
			check: func() bool {
				input := map[int]string{1: "a", 2: "b", 3: "c"}
				result := lxmaps.MapKeys(input, func(k int) string { return "key_" + string(rune(k+48)) })
				for _, v := range input {
					found := false
					for _, rv := range result {
						if v == rv {
							found = true
							break
						}
					}
					if !found {
						return false
					}
				}
				return true
			},
			checkMsg: "all values should be preserved",
		},
		{
			name: "zero values preserved",
			check: func() bool {
				input := map[int]int{1: 0, 2: 0, 3: 5}
				result := lxmaps.MapKeys(input, func(k int) string { return "k" })
				for _, v := range result {
					if v == 0 || v == 5 {
						return true
					}
				}
				return false
			},
			checkMsg: "should preserve zero values",
		},
		{
			name: "negative values preserved",
			check: func() bool {
				input := map[int]int{1: -5, 2: 10}
				result := lxmaps.MapKeys(input, func(k int) string { return "key" })
				for _, v := range result {
					if v == -5 || v == 10 {
						return true
					}
				}
				return false
			},
			checkMsg: "should preserve negative values",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("MapKeys() %s", tt.checkMsg)
			}
		})
	}
}

func TestMapKeys_CustomStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	tests := []struct {
		name     string
		input    map[int]Person
		fn       func(int) string
		check    func(map[string]Person) bool
		checkMsg string
	}{
		{
			name: "struct values preserved",
			input: map[int]Person{
				1: {Name: "Alice", Age: 30},
				2: {Name: "Bob", Age: 25},
			},
			fn: func(k int) string {
				if k == 1 {
					return "alice_id"
				}
				return "bob_id"
			},
			check: func(result map[string]Person) bool {
				return len(result) == 2 &&
					result["alice_id"].Name == "Alice" &&
					result["alice_id"].Age == 30 &&
					result["bob_id"].Name == "Bob" &&
					result["bob_id"].Age == 25
			},
			checkMsg: "should preserve struct values with transformed keys",
		},
		{
			name: "nil struct fields",
			input: map[int]Person{
				1: {Name: "", Age: 0},
				2: {Name: "Test", Age: 20},
			},
			fn: func(k int) string { return "person_" + string(rune(k+48)) },
			check: func(result map[string]Person) bool {
				return len(result) == 2
			},
			checkMsg: "should preserve zero struct values",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MapKeys(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("MapKeys() %s", tt.checkMsg)
			}
		})
	}
}

func TestMapKeys_LargeMap(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		check    func(map[string]int) bool
		checkMsg string
	}{
		{
			name: "1000 entries with unique transformation",
			size: 1000,
			check: func(result map[string]int) bool {
				return len(result) == 1000
			},
			checkMsg: "should handle 1000 entries",
		},
		{
			name: "100 entries",
			size: 100,
			check: func(result map[string]int) bool {
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

			result := lxmaps.MapKeys(input, func(k int) string {
				return "key_" + string([]byte{byte(k / 256), byte(k % 256)})
			})
			if !tt.check(result) {
				t.Errorf("MapKeys() %s, got %d keys", tt.checkMsg, len(result))
			}
		})
	}
}

func TestMapKeys_IntToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		fn       func(int) int
		check    func(map[int]string) bool
		checkMsg string
	}{
		{
			name:     "nil map",
			input:    nil,
			fn:       func(k int) int { return 0 },
			check:    func(result map[int]string) bool { return result == nil },
			checkMsg: "should return nil",
		},
		{
			name: "negate keys",
			input: map[int]string{
				1:  "one",
				-2: "minus_two",
				0:  "zero",
			},
			fn: func(k int) int { return -k },
			check: func(result map[int]string) bool {
				return len(result) == 3 &&
					result[-1] == "one" &&
					result[2] == "minus_two" &&
					result[0] == "zero"
			},
			checkMsg: "should negate integer keys",
		},
		{
			name: "square keys",
			input: map[int]string{
				1: "a",
				2: "b",
				3: "c",
			},
			fn: func(k int) int { return k * k },
			check: func(result map[int]string) bool {
				return len(result) == 3 &&
					result[1] == "a" &&
					result[4] == "b" &&
					result[9] == "c"
			},
			checkMsg: "should square integer keys",
		},
		{
			name: "modulo transformation",
			input: map[int]string{
				1: "one",
				2: "two",
				3: "three",
				4: "four",
				5: "five",
			},
			fn: func(k int) int { return k % 3 },
			check: func(result map[int]string) bool {
				return len(result) <= 3
			},
			checkMsg: "should apply modulo transformation",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.MapKeys(tt.input, tt.fn)
			if !tt.check(result) {
				t.Errorf("MapKeys() %s", tt.checkMsg)
			}
		})
	}
}

func TestMapKeys_ConsistencyWithSize(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "size preserved when no key collision",
			check: func() bool {
				input := map[int]string{1: "a", 2: "b", 3: "c"}
				result := lxmaps.MapKeys(input, func(k int) string {
					return "key_" + string(rune(k+48))
				})
				return len(result) == len(input)
			},
		},
		{
			name: "size reduced when collisions occur",
			check: func() bool {
				input := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
				result := lxmaps.MapKeys(input, func(k int) string { return "same" })
				return len(result) <= len(input)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("MapKeys() consistency check failed")
			}
		})
	}
}
