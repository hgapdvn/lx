package lxmaps_test

import (
	"testing"

	"github.com/hgapdvn/lx/lxmaps"
)

func TestPop_StringInt(t *testing.T) {
	tests := []struct {
		name        string
		input       map[string]int
		keyToPop    string
		expectedVal int
		expectedOk  bool
		expectedLen int
	}{
		{
			name:        "nil map",
			input:       nil,
			keyToPop:    "any",
			expectedVal: 0,
			expectedOk:  false,
			expectedLen: 0,
		},
		{
			name:        "empty map",
			input:       map[string]int{},
			keyToPop:    "a",
			expectedVal: 0,
			expectedOk:  false,
			expectedLen: 0,
		},
		{
			name:        "single entry found",
			input:       map[string]int{"a": 10},
			keyToPop:    "a",
			expectedVal: 10,
			expectedOk:  true,
			expectedLen: 0,
		},
		{
			name:        "single entry not found",
			input:       map[string]int{"a": 10},
			keyToPop:    "b",
			expectedVal: 0,
			expectedOk:  false,
			expectedLen: 1,
		},
		{
			name:        "multiple entries found",
			input:       map[string]int{"a": 1, "b": 2, "c": 3},
			keyToPop:    "b",
			expectedVal: 2,
			expectedOk:  true,
			expectedLen: 2,
		},
		{
			name:        "multiple entries not found",
			input:       map[string]int{"a": 1, "b": 2, "c": 3},
			keyToPop:    "d",
			expectedVal: 0,
			expectedOk:  false,
			expectedLen: 3,
		},
		{
			name:        "pop from first",
			input:       map[string]int{"a": 1, "b": 2, "c": 3},
			keyToPop:    "a",
			expectedVal: 1,
			expectedOk:  true,
			expectedLen: 2,
		},
		{
			name:        "pop from last",
			input:       map[string]int{"a": 1, "b": 2, "c": 3},
			keyToPop:    "c",
			expectedVal: 3,
			expectedOk:  true,
			expectedLen: 2,
		},
		{
			name:        "zero value present",
			input:       map[string]int{"a": 0, "b": 1},
			keyToPop:    "a",
			expectedVal: 0,
			expectedOk:  true,
			expectedLen: 1,
		},
		{
			name:        "negative value",
			input:       map[string]int{"a": -5, "b": 10},
			keyToPop:    "a",
			expectedVal: -5,
			expectedOk:  true,
			expectedLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, ok := lxmaps.Pop(tt.input, tt.keyToPop)
			if val != tt.expectedVal {
				t.Errorf("Pop() value = %d, want %d", val, tt.expectedVal)
			}
			if ok != tt.expectedOk {
				t.Errorf("Pop() ok = %v, want %v", ok, tt.expectedOk)
			}
			if tt.input != nil && len(tt.input) != tt.expectedLen {
				t.Errorf("Pop() map length = %d, want %d", len(tt.input), tt.expectedLen)
			}
		})
	}
}

func TestPop_IntString(t *testing.T) {
	tests := []struct {
		name        string
		input       map[int]string
		keyToPop    int
		expectedVal string
		expectedOk  bool
		expectedLen int
	}{
		{
			name:        "nil map",
			input:       nil,
			keyToPop:    1,
			expectedVal: "",
			expectedOk:  false,
			expectedLen: 0,
		},
		{
			name:        "empty map",
			input:       map[int]string{},
			keyToPop:    1,
			expectedVal: "",
			expectedOk:  false,
			expectedLen: 0,
		},
		{
			name:        "string found",
			input:       map[int]string{1: "one", 2: "two"},
			keyToPop:    1,
			expectedVal: "one",
			expectedOk:  true,
			expectedLen: 1,
		},
		{
			name:        "string not found",
			input:       map[int]string{1: "one", 2: "two"},
			keyToPop:    3,
			expectedVal: "",
			expectedOk:  false,
			expectedLen: 2,
		},
		{
			name:        "empty string value",
			input:       map[int]string{1: "", 2: "value"},
			keyToPop:    1,
			expectedVal: "",
			expectedOk:  true,
			expectedLen: 1,
		},
		{
			name:        "negative key",
			input:       map[int]string{-1: "neg", 0: "zero"},
			keyToPop:    -1,
			expectedVal: "neg",
			expectedOk:  true,
			expectedLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, ok := lxmaps.Pop(tt.input, tt.keyToPop)
			if val != tt.expectedVal {
				t.Errorf("Pop() value = %q, want %q", val, tt.expectedVal)
			}
			if ok != tt.expectedOk {
				t.Errorf("Pop() ok = %v, want %v", ok, tt.expectedOk)
			}
			if tt.input != nil && len(tt.input) != tt.expectedLen {
				t.Errorf("Pop() map length = %d, want %d", len(tt.input), tt.expectedLen)
			}
		})
	}
}

func TestPop_StringBool(t *testing.T) {
	tests := []struct {
		name        string
		input       map[string]bool
		keyToPop    string
		expectedVal bool
		expectedOk  bool
		expectedLen int
	}{
		{
			name:        "nil map",
			input:       nil,
			keyToPop:    "key",
			expectedVal: false,
			expectedOk:  false,
			expectedLen: 0,
		},
		{
			name:        "true value found",
			input:       map[string]bool{"flag": true, "other": false},
			keyToPop:    "flag",
			expectedVal: true,
			expectedOk:  true,
			expectedLen: 1,
		},
		{
			name:        "false value found",
			input:       map[string]bool{"flag": false, "other": true},
			keyToPop:    "flag",
			expectedVal: false,
			expectedOk:  true,
			expectedLen: 1,
		},
		{
			name:        "not found",
			input:       map[string]bool{"a": true},
			keyToPop:    "b",
			expectedVal: false,
			expectedOk:  false,
			expectedLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, ok := lxmaps.Pop(tt.input, tt.keyToPop)
			if val != tt.expectedVal {
				t.Errorf("Pop() value = %v, want %v", val, tt.expectedVal)
			}
			if ok != tt.expectedOk {
				t.Errorf("Pop() ok = %v, want %v", ok, tt.expectedOk)
			}
			if tt.input != nil && len(tt.input) != tt.expectedLen {
				t.Errorf("Pop() map length = %d, want %d", len(tt.input), tt.expectedLen)
			}
		})
	}
}

func TestPop_StringFloat(t *testing.T) {
	tests := []struct {
		name        string
		input       map[string]float64
		keyToPop    string
		expectedVal float64
		expectedOk  bool
		expectedLen int
	}{
		{
			name:        "float found",
			input:       map[string]float64{"pi": 3.14159, "e": 2.71828},
			keyToPop:    "pi",
			expectedVal: 3.14159,
			expectedOk:  true,
			expectedLen: 1,
		},
		{
			name:        "zero float found",
			input:       map[string]float64{"zero": 0.0, "one": 1.0},
			keyToPop:    "zero",
			expectedVal: 0.0,
			expectedOk:  true,
			expectedLen: 1,
		},
		{
			name:        "not found",
			input:       map[string]float64{"pi": 3.14},
			keyToPop:    "e",
			expectedVal: 0.0,
			expectedOk:  false,
			expectedLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, ok := lxmaps.Pop(tt.input, tt.keyToPop)
			if val != tt.expectedVal {
				t.Errorf("Pop() value = %f, want %f", val, tt.expectedVal)
			}
			if ok != tt.expectedOk {
				t.Errorf("Pop() ok = %v, want %v", ok, tt.expectedOk)
			}
			if tt.input != nil && len(tt.input) != tt.expectedLen {
				t.Errorf("Pop() map length = %d, want %d", len(tt.input), tt.expectedLen)
			}
		})
	}
}

func TestPop_CustomStruct(t *testing.T) {
	type Item struct {
		Name  string
		Value int
	}

	tests := []struct {
		name        string
		input       map[string]Item
		keyToPop    string
		expectedVal Item
		expectedOk  bool
		expectedLen int
	}{
		{
			name:        "nil map",
			input:       nil,
			keyToPop:    "a",
			expectedVal: Item{},
			expectedOk:  false,
			expectedLen: 0,
		},
		{
			name: "struct found",
			input: map[string]Item{
				"item1": {Name: "Alice", Value: 10},
				"item2": {Name: "Bob", Value: 20},
			},
			keyToPop:    "item1",
			expectedVal: Item{Name: "Alice", Value: 10},
			expectedOk:  true,
			expectedLen: 1,
		},
		{
			name: "struct not found",
			input: map[string]Item{
				"item1": {Name: "Alice", Value: 10},
			},
			keyToPop:    "item2",
			expectedVal: Item{},
			expectedOk:  false,
			expectedLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, ok := lxmaps.Pop(tt.input, tt.keyToPop)
			if val != tt.expectedVal {
				t.Errorf("Pop() value = %+v, want %+v", val, tt.expectedVal)
			}
			if ok != tt.expectedOk {
				t.Errorf("Pop() ok = %v, want %v", ok, tt.expectedOk)
			}
			if tt.input != nil && len(tt.input) != tt.expectedLen {
				t.Errorf("Pop() map length = %d, want %d", len(tt.input), tt.expectedLen)
			}
		})
	}
}

func TestPop_SequentialPops(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]int
		keysToPop []string
		check     func(map[string]int) bool
	}{
		{
			name:      "pop all entries",
			input:     map[string]int{"a": 1, "b": 2, "c": 3},
			keysToPop: []string{"a", "b", "c"},
			check: func(m map[string]int) bool {
				return len(m) == 0
			},
		},
		{
			name:      "pop some entries",
			input:     map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			keysToPop: []string{"b", "d"},
			check: func(m map[string]int) bool {
				return len(m) == 2 && m["a"] == 1 && m["c"] == 3
			},
		},
		{
			name:      "pop non-existent entries",
			input:     map[string]int{"a": 1, "b": 2},
			keysToPop: []string{"a", "x", "y"},
			check: func(m map[string]int) bool {
				return len(m) == 1 && m["b"] == 2
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, key := range tt.keysToPop {
				lxmaps.Pop(tt.input, key)
			}
			if !tt.check(tt.input) {
				t.Errorf("Pop() sequential pops failed check")
			}
		})
	}
}

func TestPop_StringString(t *testing.T) {
	tests := []struct {
		name        string
		input       map[string]string
		keyToPop    string
		expectedVal string
		expectedOk  bool
		expectedLen int
	}{
		{
			name:        "nil map",
			input:       nil,
			keyToPop:    "key",
			expectedVal: "",
			expectedOk:  false,
			expectedLen: 0,
		},
		{
			name: "string found",
			input: map[string]string{
				"name": "John",
				"city": "NYC",
			},
			keyToPop:    "name",
			expectedVal: "John",
			expectedOk:  true,
			expectedLen: 1,
		},
		{
			name: "empty string value found",
			input: map[string]string{
				"empty": "",
				"full":  "value",
			},
			keyToPop:    "empty",
			expectedVal: "",
			expectedOk:  true,
			expectedLen: 1,
		},
		{
			name:        "key not found",
			input:       map[string]string{"a": "apple"},
			keyToPop:    "b",
			expectedVal: "",
			expectedOk:  false,
			expectedLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, ok := lxmaps.Pop(tt.input, tt.keyToPop)
			if val != tt.expectedVal {
				t.Errorf("Pop() value = %q, want %q", val, tt.expectedVal)
			}
			if ok != tt.expectedOk {
				t.Errorf("Pop() ok = %v, want %v", ok, tt.expectedOk)
			}
			if tt.input != nil && len(tt.input) != tt.expectedLen {
				t.Errorf("Pop() map length = %d, want %d", len(tt.input), tt.expectedLen)
			}
		})
	}
}

func TestPop_IntInt(t *testing.T) {
	tests := []struct {
		name        string
		input       map[int]int
		keyToPop    int
		expectedVal int
		expectedOk  bool
		expectedLen int
	}{
		{
			name:        "nil map",
			input:       nil,
			keyToPop:    1,
			expectedVal: 0,
			expectedOk:  false,
			expectedLen: 0,
		},
		{
			name:        "found",
			input:       map[int]int{1: 10, 2: 20, 3: 30},
			keyToPop:    2,
			expectedVal: 20,
			expectedOk:  true,
			expectedLen: 2,
		},
		{
			name:        "not found",
			input:       map[int]int{1: 10, 2: 20},
			keyToPop:    5,
			expectedVal: 0,
			expectedOk:  false,
			expectedLen: 2,
		},
		{
			name:        "zero key found",
			input:       map[int]int{0: 100, 1: 200},
			keyToPop:    0,
			expectedVal: 100,
			expectedOk:  true,
			expectedLen: 1,
		},
		{
			name:        "negative key found",
			input:       map[int]int{-1: 10, 1: 20},
			keyToPop:    -1,
			expectedVal: 10,
			expectedOk:  true,
			expectedLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, ok := lxmaps.Pop(tt.input, tt.keyToPop)
			if val != tt.expectedVal {
				t.Errorf("Pop() value = %d, want %d", val, tt.expectedVal)
			}
			if ok != tt.expectedOk {
				t.Errorf("Pop() ok = %v, want %v", ok, tt.expectedOk)
			}
			if tt.input != nil && len(tt.input) != tt.expectedLen {
				t.Errorf("Pop() map length = %d, want %d", len(tt.input), tt.expectedLen)
			}
		})
	}
}

func TestPop_MapModification(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		keyToPop string
		check    func(map[string]int) bool
	}{
		{
			name:     "map modified",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			keyToPop: "b",
			check: func(m map[string]int) bool {
				_, exists := m["b"]
				return !exists && len(m) == 2
			},
		},
		{
			name:     "map not modified when key not found",
			input:    map[string]int{"a": 1, "b": 2},
			keyToPop: "x",
			check: func(m map[string]int) bool {
				return len(m) == 2 && m["a"] == 1 && m["b"] == 2
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lxmaps.Pop(tt.input, tt.keyToPop)
			if !tt.check(tt.input) {
				t.Errorf("Pop() map modification check failed")
			}
		})
	}
}
