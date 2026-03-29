package lxmaps_test

import (
	"testing"

	"github.com/hgapdvn/lx/maps"
)

func TestContainsValue_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		value    int
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			value:    1,
			expected: false,
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			value:    1,
			expected: false,
		},
		{
			name:     "single entry value exists",
			input:    map[string]int{"a": 1},
			value:    1,
			expected: true,
		},
		{
			name:     "single entry value does not exist",
			input:    map[string]int{"a": 1},
			value:    2,
			expected: false,
		},
		{
			name:     "multiple entries value exists",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			value:    2,
			expected: true,
		},
		{
			name:     "multiple entries value does not exist",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			value:    4,
			expected: false,
		},
		{
			name:     "zero value exists",
			input:    map[string]int{"a": 0, "b": 1},
			value:    0,
			expected: true,
		},
		{
			name:     "zero value does not exist",
			input:    map[string]int{"a": 1, "b": 2},
			value:    0,
			expected: false,
		},
		{
			name:     "negative value exists",
			input:    map[string]int{"a": -1, "b": -2},
			value:    -1,
			expected: true,
		},
		{
			name:     "negative value does not exist",
			input:    map[string]int{"a": -1, "b": -2},
			value:    1,
			expected: false,
		},
		{
			name:     "duplicate values in map",
			input:    map[string]int{"a": 5, "b": 5, "c": 5},
			value:    5,
			expected: true,
		},
		{
			name:     "mixed sign values",
			input:    map[string]int{"neg": -5, "zero": 0, "pos": 5},
			value:    0,
			expected: true,
		},
		{
			name:     "large value exists",
			input:    map[string]int{"big": 1000000},
			value:    1000000,
			expected: true,
		},
		{
			name:     "large value does not exist",
			input:    map[string]int{"big": 1000000},
			value:    999999,
			expected: false,
		},
		{
			name:     "many entries value at start",
			input:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			value:    1,
			expected: true,
		},
		{
			name:     "many entries value at end",
			input:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			value:    5,
			expected: true,
		},
		{
			name:     "many entries value in middle",
			input:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			value:    3,
			expected: true,
		},
		{
			name:     "many entries value does not exist",
			input:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			value:    10,
			expected: false,
		},
		{
			name:     "unicode keys different values",
			input:    map[string]int{"こんにちは": 1, "世界": 2},
			value:    2,
			expected: true,
		},
		{
			name:     "min int value",
			input:    map[string]int{"min": -9223372036854775808},
			value:    -9223372036854775808,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ContainsValue(tt.input, tt.value)
			if got != tt.expected {
				t.Fatalf("ContainsValue(%v, %d) = %v, want %v", tt.input, tt.value, got, tt.expected)
			}
		})
	}
}

func TestContainsValue_IntString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		value    string
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			value:    "test",
			expected: false,
		},
		{
			name:     "empty map",
			input:    map[int]string{},
			value:    "test",
			expected: false,
		},
		{
			name:     "single entry value exists",
			input:    map[int]string{1: "one"},
			value:    "one",
			expected: true,
		},
		{
			name:     "single entry value does not exist",
			input:    map[int]string{1: "one"},
			value:    "two",
			expected: false,
		},
		{
			name:     "multiple entries value exists",
			input:    map[int]string{1: "one", 2: "two", 3: "three"},
			value:    "two",
			expected: true,
		},
		{
			name:     "multiple entries value does not exist",
			input:    map[int]string{1: "one", 2: "two", 3: "three"},
			value:    "four",
			expected: false,
		},
		{
			name:     "empty string value exists",
			input:    map[int]string{1: "", 2: "value"},
			value:    "",
			expected: true,
		},
		{
			name:     "empty string value does not exist",
			input:    map[int]string{1: "one", 2: "two"},
			value:    "",
			expected: false,
		},
		{
			name:     "duplicate values in map",
			input:    map[int]string{1: "same", 2: "same", 3: "same"},
			value:    "same",
			expected: true,
		},
		{
			name:     "case sensitive value match",
			input:    map[int]string{1: "Test", 2: "test"},
			value:    "Test",
			expected: true,
		},
		{
			name:     "case sensitive value mismatch",
			input:    map[int]string{1: "Test"},
			value:    "test",
			expected: false,
		},
		{
			name:     "unicode value exists",
			input:    map[int]string{1: "こんにちは", 2: "世界"},
			value:    "こんにちは",
			expected: true,
		},
		{
			name:     "unicode value does not exist",
			input:    map[int]string{1: "こんにちは"},
			value:    "hello",
			expected: false,
		},
		{
			name:     "emoji value exists",
			input:    map[int]string{1: "😊", 2: "👋"},
			value:    "😊",
			expected: true,
		},
		{
			name:     "emoji value does not exist",
			input:    map[int]string{1: "😊"},
			value:    "🚀",
			expected: false,
		},
		{
			name:     "special character value exists",
			input:    map[int]string{1: "!@#", 2: "$%"},
			value:    "!@#",
			expected: true,
		},
		{
			name:     "special character value does not exist",
			input:    map[int]string{1: "!@#"},
			value:    "$%",
			expected: false,
		},
		{
			name:     "whitespace value exists",
			input:    map[int]string{1: " ", 2: "\t"},
			value:    " ",
			expected: true,
		},
		{
			name:     "newline value exists",
			input:    map[int]string{1: "\n", 2: "ok"},
			value:    "\n",
			expected: true,
		},
		{
			name:     "long string value exists",
			input:    map[int]string{1: "long_string_example_0001", 2: "other"},
			value:    "long_string_example_0001",
			expected: true,
		},
		{
			name:     "many entries value exists",
			input:    map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"},
			value:    "c",
			expected: true,
		},
		{
			name:     "many entries value does not exist",
			input:    map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"},
			value:    "f",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ContainsValue(tt.input, tt.value)
			if got != tt.expected {
				t.Fatalf("ContainsValue(%v, %q) = %v, want %v", tt.input, tt.value, got, tt.expected)
			}
		})
	}
}

func TestContainsValue_StructKey(t *testing.T) {
	type Value struct {
		ID   int
		Name string
	}

	tests := []struct {
		name     string
		input    map[string]Value
		value    Value
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			value:    Value{ID: 1, Name: "test"},
			expected: false,
		},
		{
			name:     "empty map",
			input:    map[string]Value{},
			value:    Value{ID: 1, Name: "test"},
			expected: false,
		},
		{
			name:     "struct value exists",
			input:    map[string]Value{"key1": {ID: 1, Name: "Alice"}},
			value:    Value{ID: 1, Name: "Alice"},
			expected: true,
		},
		{
			name:     "struct value does not exist",
			input:    map[string]Value{"key1": {ID: 1, Name: "Alice"}},
			value:    Value{ID: 2, Name: "Bob"},
			expected: false,
		},
		{
			name: "multiple struct values one exists",
			input: map[string]Value{
				"alice": {ID: 1, Name: "Alice"},
				"bob":   {ID: 2, Name: "Bob"},
			},
			value:    Value{ID: 1, Name: "Alice"},
			expected: true,
		},
		{
			name: "multiple struct values one does not exist",
			input: map[string]Value{
				"alice": {ID: 1, Name: "Alice"},
				"bob":   {ID: 2, Name: "Bob"},
			},
			value:    Value{ID: 3, Name: "Charlie"},
			expected: false,
		},
		{
			name:     "duplicate struct values",
			input:    map[string]Value{"a": {ID: 1, Name: "same"}, "b": {ID: 1, Name: "same"}},
			value:    Value{ID: 1, Name: "same"},
			expected: true,
		},
		{
			name:     "struct field mismatch",
			input:    map[string]Value{"key": {ID: 1, Name: "Alice"}},
			value:    Value{ID: 1, Name: "alice"},
			expected: false,
		},
		{
			name:     "unicode struct field match",
			input:    map[string]Value{"key": {ID: 1, Name: "こんにちは"}},
			value:    Value{ID: 1, Name: "こんにちは"},
			expected: true,
		},
		{
			name:     "empty struct field",
			input:    map[string]Value{"key": {ID: 0, Name: ""}},
			value:    Value{ID: 0, Name: ""},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ContainsValue(tt.input, tt.value)
			if got != tt.expected {
				t.Fatalf("ContainsValue(%v, %v) = %v, want %v", tt.input, tt.value, got, tt.expected)
			}
		})
	}
}

func TestContainsValue_BoolValue(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]bool
		value    bool
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			value:    true,
			expected: false,
		},
		{
			name:     "single true value exists",
			input:    map[string]bool{"a": true},
			value:    true,
			expected: true,
		},
		{
			name:     "single false value exists",
			input:    map[string]bool{"a": false},
			value:    false,
			expected: true,
		},
		{
			name:     "true value does not exist",
			input:    map[string]bool{"a": false},
			value:    true,
			expected: false,
		},
		{
			name:     "false value does not exist",
			input:    map[string]bool{"a": true},
			value:    false,
			expected: false,
		},
		{
			name:     "mixed bool values true found",
			input:    map[string]bool{"a": true, "b": false, "c": true},
			value:    true,
			expected: true,
		},
		{
			name:     "mixed bool values false found",
			input:    map[string]bool{"a": true, "b": false, "c": true},
			value:    false,
			expected: true,
		},
		{
			name:     "all true values",
			input:    map[string]bool{"a": true, "b": true, "c": true},
			value:    true,
			expected: true,
		},
		{
			name:     "all false values",
			input:    map[string]bool{"a": false, "b": false, "c": false},
			value:    false,
			expected: true,
		},
		{
			name:     "empty map true",
			input:    map[string]bool{},
			value:    true,
			expected: false,
		},
		{
			name:     "empty map false",
			input:    map[string]bool{},
			value:    false,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ContainsValue(tt.input, tt.value)
			if got != tt.expected {
				t.Fatalf("ContainsValue(%v, %v) = %v, want %v", tt.input, tt.value, got, tt.expected)
			}
		})
	}
}

func TestContainsValue_FloatValue(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]float64
		value    float64
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			value:    3.14,
			expected: false,
		},
		{
			name:     "empty map",
			input:    map[string]float64{},
			value:    3.14,
			expected: false,
		},
		{
			name:     "single value exists",
			input:    map[string]float64{"pi": 3.14},
			value:    3.14,
			expected: true,
		},
		{
			name:     "single value does not exist",
			input:    map[string]float64{"pi": 3.14},
			value:    2.71,
			expected: false,
		},
		{
			name:     "zero value exists",
			input:    map[string]float64{"zero": 0.0},
			value:    0.0,
			expected: true,
		},
		{
			name:     "negative value exists",
			input:    map[string]float64{"neg": -3.14},
			value:    -3.14,
			expected: true,
		},
		{
			name:     "multiple values one exists",
			input:    map[string]float64{"a": 1.5, "b": 2.5, "c": 3.5},
			value:    2.5,
			expected: true,
		},
		{
			name:     "multiple values none exist",
			input:    map[string]float64{"a": 1.5, "b": 2.5, "c": 3.5},
			value:    4.5,
			expected: false,
		},
		{
			name:     "duplicate float values",
			input:    map[string]float64{"a": 3.14, "b": 3.14, "c": 3.14},
			value:    3.14,
			expected: true,
		},
		{
			name:     "very small float value",
			input:    map[string]float64{"tiny": 1.0e-10},
			value:    1.0e-10,
			expected: true,
		},
		{
			name:     "very large float value",
			input:    map[string]float64{"huge": 1.0e10},
			value:    1.0e10,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ContainsValue(tt.input, tt.value)
			if got != tt.expected {
				t.Fatalf("ContainsValue(%v, %f) = %v, want %v", tt.input, tt.value, got, tt.expected)
			}
		})
	}
}

func TestContainsValue_RuneValue(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]rune
		value    rune
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			value:    'a',
			expected: false,
		},
		{
			name:     "empty map",
			input:    map[string]rune{},
			value:    'a',
			expected: false,
		},
		{
			name:     "single char value exists",
			input:    map[string]rune{"letter": 'a'},
			value:    'a',
			expected: true,
		},
		{
			name:     "single char value does not exist",
			input:    map[string]rune{"letter": 'a'},
			value:    'b',
			expected: false,
		},
		{
			name:     "multiple char values",
			input:    map[string]rune{"a": 'x', "b": 'y', "c": 'z'},
			value:    'y',
			expected: true,
		},
		{
			name:     "unicode rune value exists",
			input:    map[string]rune{"japanese": rune('日')},
			value:    rune('日'),
			expected: true,
		},
		{
			name:     "emoji rune value exists",
			input:    map[string]rune{"emoji": rune('😊')},
			value:    rune('😊'),
			expected: true,
		},
		{
			name:     "whitespace rune value",
			input:    map[string]rune{"space": ' ', "tab": '\t'},
			value:    ' ',
			expected: true,
		},
		{
			name:     "zero rune value",
			input:    map[string]rune{"null": 0},
			value:    0,
			expected: true,
		},
		{
			name:     "case sensitive rune",
			input:    map[string]rune{"upper": 'A'},
			value:    'a',
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ContainsValue(tt.input, tt.value)
			if got != tt.expected {
				t.Fatalf("ContainsValue(%v, %q) = %v, want %v", tt.input, tt.value, got, tt.expected)
			}
		})
	}
}

func TestContainsValue_ByteValue(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]byte
		value    byte
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			value:    'a',
			expected: false,
		},
		{
			name:     "empty map",
			input:    map[string]byte{},
			value:    'a',
			expected: false,
		},
		{
			name:     "single byte value exists",
			input:    map[string]byte{"byte": 'a'},
			value:    'a',
			expected: true,
		},
		{
			name:     "single byte value does not exist",
			input:    map[string]byte{"byte": 'a'},
			value:    'b',
			expected: false,
		},
		{
			name:     "multiple byte values",
			input:    map[string]byte{"a": 'x', "b": 'y', "c": 'z'},
			value:    'y',
			expected: true,
		},
		{
			name:     "zero byte value",
			input:    map[string]byte{"zero": 0},
			value:    0,
			expected: true,
		},
		{
			name:     "max byte value",
			input:    map[string]byte{"max": 255},
			value:    255,
			expected: true,
		},
		{
			name:     "duplicate byte values",
			input:    map[string]byte{"a": 'x', "b": 'x', "c": 'x'},
			value:    'x',
			expected: true,
		},
		{
			name:     "case sensitive byte",
			input:    map[string]byte{"upper": 'A'},
			value:    'a',
			expected: false,
		},
		{
			name:     "all bytes",
			input:    map[string]byte{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4},
			value:    3,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ContainsValue(tt.input, tt.value)
			if got != tt.expected {
				t.Fatalf("ContainsValue(%v, %d) = %v, want %v", tt.input, tt.value, got, tt.expected)
			}
		})
	}
}
