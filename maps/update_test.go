package lxmaps_test

import (
	"testing"

	lxmaps "github.com/hgapdvn/lx/maps"
)

func TestUpdate_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		key      string
		fn       func(int, bool) int
		check    func(map[string]int) bool
		checkMsg string
	}{
		{
			name:     "nil map",
			input:    nil,
			key:      "a",
			fn:       func(v int, exists bool) int { return 100 },
			check:    func(result map[string]int) bool { return len(result) == 1 && result["a"] == 100 },
			checkMsg: "should create map and add entry",
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			key:      "a",
			fn:       func(v int, exists bool) int { return 100 },
			check:    func(result map[string]int) bool { return len(result) == 1 && result["a"] == 100 },
			checkMsg: "should add entry to empty map",
		},
		{
			name:  "update existing key",
			input: map[string]int{"a": 1, "b": 2},
			key:   "a",
			fn: func(v int, exists bool) int {
				if exists {
					return v + 10
				}
				return 100
			},
			check:    func(result map[string]int) bool { return len(result) == 2 && result["a"] == 11 && result["b"] == 2 },
			checkMsg: "should increment existing value",
		},
		{
			name:  "insert new key",
			input: map[string]int{"a": 1},
			key:   "b",
			fn: func(v int, exists bool) int {
				if exists {
					return v + 10
				}
				return 100
			},
			check:    func(result map[string]int) bool { return len(result) == 2 && result["a"] == 1 && result["b"] == 100 },
			checkMsg: "should insert new key with default",
		},
		{
			name:  "update with conditional logic",
			input: map[string]int{"x": 5},
			key:   "x",
			fn: func(v int, exists bool) int {
				if !exists {
					return 0
				}
				if v < 10 {
					return v * 2
				}
				return v
			},
			check:    func(result map[string]int) bool { return result["x"] == 10 },
			checkMsg: "should double value less than 10",
		},
		{
			name:  "check exists boolean for existing key",
			input: map[string]int{"key": 42},
			key:   "key",
			fn: func(v int, exists bool) int {
				if exists && v == 42 {
					return 99
				}
				return -1
			},
			check:    func(result map[string]int) bool { return result["key"] == 99 },
			checkMsg: "should receive true for existing key",
		},
		{
			name:  "check exists boolean for new key",
			input: map[string]int{"other": 1},
			key:   "new",
			fn: func(v int, exists bool) int {
				if !exists && v == 0 {
					return 555
				}
				return -1
			},
			check:    func(result map[string]int) bool { return result["new"] == 555 },
			checkMsg: "should receive false and zero value for new key",
		},
		{
			name:  "update to zero",
			input: map[string]int{"a": 10},
			key:   "a",
			fn:    func(v int, exists bool) int { return 0 },
			check: func(result map[string]int) bool {
				val, ok := result["a"]
				return ok && val == 0
			},
			checkMsg: "should update to zero value",
		},
		{
			name:     "update negative to positive",
			input:    map[string]int{"neg": -5},
			key:      "neg",
			fn:       func(v int, exists bool) int { return -v },
			check:    func(result map[string]int) bool { return result["neg"] == 5 },
			checkMsg: "should negate value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Update(tt.input, tt.key, tt.fn)
			if !tt.check(result) {
				t.Errorf("Update() %s", tt.checkMsg)
			}
		})
	}
}

func TestUpdate_IntString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		key      int
		fn       func(string, bool) string
		check    func(map[int]string) bool
		checkMsg string
	}{
		{
			name:     "nil map string",
			input:    nil,
			key:      1,
			fn:       func(v string, exists bool) string { return "hello" },
			check:    func(result map[int]string) bool { return len(result) == 1 && result[1] == "hello" },
			checkMsg: "should create map with string value",
		},
		{
			name:  "append to existing string",
			input: map[int]string{1: "hello"},
			key:   1,
			fn: func(v string, exists bool) string {
				if exists {
					return v + " world"
				}
				return "default"
			},
			check:    func(result map[int]string) bool { return result[1] == "hello world" },
			checkMsg: "should append to existing string",
		},
		{
			name:  "set default for missing key",
			input: map[int]string{1: "exists"},
			key:   2,
			fn: func(v string, exists bool) string {
				if exists {
					return v
				}
				return "default"
			},
			check:    func(result map[int]string) bool { return result[2] == "default" },
			checkMsg: "should use default for missing key",
		},
		{
			name:  "uppercase existing",
			input: map[int]string{1: "test"},
			key:   1,
			fn: func(v string, exists bool) string {
				if !exists {
					return v
				}
				if len(v) > 0 {
					return "UPDATED"
				}
				return v
			},
			check:    func(result map[int]string) bool { return result[1] == "UPDATED" },
			checkMsg: "should uppercase existing",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Update(tt.input, tt.key, tt.fn)
			if !tt.check(result) {
				t.Errorf("Update() %s", tt.checkMsg)
			}
		})
	}
}

func TestUpdate_StringBool(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]bool
		key      string
		fn       func(bool, bool) bool
		check    func(map[string]bool) bool
		checkMsg string
	}{
		{
			name:     "toggle existing true to false",
			input:    map[string]bool{"flag": true},
			key:      "flag",
			fn:       func(v bool, exists bool) bool { return !v },
			check:    func(result map[string]bool) bool { return result["flag"] == false },
			checkMsg: "should toggle true to false",
		},
		{
			name:     "toggle existing false to true",
			input:    map[string]bool{"flag": false},
			key:      "flag",
			fn:       func(v bool, exists bool) bool { return !v },
			check:    func(result map[string]bool) bool { return result["flag"] == true },
			checkMsg: "should toggle false to true",
		},
		{
			name:  "set new bool",
			input: map[string]bool{"a": true},
			key:   "b",
			fn: func(v bool, exists bool) bool {
				if exists {
					return v
				}
				return true
			},
			check:    func(result map[string]bool) bool { return result["b"] == true },
			checkMsg: "should set new bool to true",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Update(tt.input, tt.key, tt.fn)
			if !tt.check(result) {
				t.Errorf("Update() %s", tt.checkMsg)
			}
		})
	}
}

func TestUpdate_CustomStruct(t *testing.T) {
	type Item struct {
		Name  string
		Count int
	}

	tests := []struct {
		name     string
		input    map[string]Item
		key      string
		fn       func(Item, bool) Item
		check    func(map[string]Item) bool
		checkMsg string
	}{
		{
			name: "update existing struct",
			input: map[string]Item{
				"item1": {Name: "apple", Count: 5},
			},
			key: "item1",
			fn: func(v Item, exists bool) Item {
				if exists {
					v.Count += 10
				}
				return v
			},
			check: func(result map[string]Item) bool {
				return result["item1"].Name == "apple" && result["item1"].Count == 15
			},
			checkMsg: "should increment count",
		},
		{
			name: "insert new struct",
			input: map[string]Item{
				"item1": {Name: "apple", Count: 5},
			},
			key: "item2",
			fn: func(v Item, exists bool) Item {
				if !exists {
					return Item{Name: "banana", Count: 3}
				}
				return v
			},
			check: func(result map[string]Item) bool {
				return result["item2"].Name == "banana" && result["item2"].Count == 3
			},
			checkMsg: "should create new struct",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Update(tt.input, tt.key, tt.fn)
			if !tt.check(result) {
				t.Errorf("Update() %s", tt.checkMsg)
			}
		})
	}
}

func TestUpdate_Chaining(t *testing.T) {
	tests := []struct {
		name     string
		check    func() bool
		checkMsg string
	}{
		{
			name: "chain multiple updates",
			check: func() bool {
				m := map[string]int{"a": 1}
				lxmaps.Update(m, "a", func(v int, exists bool) int { return v + 1 })
				lxmaps.Update(m, "a", func(v int, exists bool) int { return v + 1 })
				lxmaps.Update(m, "a", func(v int, exists bool) int { return v + 1 })
				return m["a"] == 4
			},
			checkMsg: "should chain updates correctly",
		},
		{
			name: "update returns modified map",
			check: func() bool {
				m := map[string]int{"x": 10}
				originalLen := len(m)
				result := lxmaps.Update(m, "x", func(v int, exists bool) int { return v * 2 })
				return result != nil && len(result) == originalLen && m["x"] == 20
			},
			checkMsg: "should return map reference with modifications",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("Update() %s", tt.checkMsg)
			}
		})
	}
}

func TestUpdate_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		check    func() bool
		checkMsg string
	}{
		{
			name: "update preserves other entries",
			check: func() bool {
				m := map[string]int{"a": 1, "b": 2, "c": 3}
				lxmaps.Update(m, "b", func(v int, exists bool) int { return 99 })
				return len(m) == 3 && m["a"] == 1 && m["b"] == 99 && m["c"] == 3
			},
			checkMsg: "should preserve other entries",
		},
		{
			name: "update multiple different keys",
			check: func() bool {
				m := map[string]int{}
				lxmaps.Update(m, "x", func(v int, exists bool) int { return 100 })
				lxmaps.Update(m, "y", func(v int, exists bool) int { return 200 })
				lxmaps.Update(m, "z", func(v int, exists bool) int { return 300 })
				return len(m) == 3 && m["x"] == 100 && m["y"] == 200 && m["z"] == 300
			},
			checkMsg: "should handle multiple updates",
		},
		{
			name: "update with side effects",
			check: func() bool {
				m := map[string]int{"count": 0}
				calls := 0
				lxmaps.Update(m, "count", func(v int, exists bool) int {
					calls++
					return v + calls
				})
				return m["count"] == 1 && calls == 1
			},
			checkMsg: "should call function exactly once",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("Update() %s", tt.checkMsg)
			}
		})
	}
}

func TestUpdate_StringString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]string
		key      string
		fn       func(string, bool) string
		check    func(map[string]string) bool
		checkMsg string
	}{
		{
			name: "concatenate strings",
			input: map[string]string{
				"greeting": "hello",
			},
			key: "greeting",
			fn: func(v string, exists bool) string {
				if exists {
					return v + " world"
				}
				return "default"
			},
			check:    func(result map[string]string) bool { return result["greeting"] == "hello world" },
			checkMsg: "should concatenate",
		},
		{
			name:     "set empty string",
			input:    map[string]string{"key": "value"},
			key:      "key",
			fn:       func(v string, exists bool) string { return "" },
			check:    func(result map[string]string) bool { return result["key"] == "" },
			checkMsg: "should set empty string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxmaps.Update(tt.input, tt.key, tt.fn)
			if !tt.check(result) {
				t.Errorf("Update() %s", tt.checkMsg)
			}
		})
	}
}
