package lxmaps_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/lxmaps"
)

func TestContains_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		key      string
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			key:      "key",
			expected: false,
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			key:      "key",
			expected: false,
		},
		{
			name:     "single entry key exists",
			input:    map[string]int{"a": 1},
			key:      "a",
			expected: true,
		},
		{
			name:     "single entry key does not exist",
			input:    map[string]int{"a": 1},
			key:      "b",
			expected: false,
		},
		{
			name:     "multiple entries key exists",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			key:      "b",
			expected: true,
		},
		{
			name:     "multiple entries key does not exist",
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			key:      "d",
			expected: false,
		},
		{
			name:     "key with zero value",
			input:    map[string]int{"a": 0},
			key:      "a",
			expected: false,
		},
		{
			name:     "key with negative value",
			input:    map[string]int{"a": -1},
			key:      "a",
			expected: true,
		},
		{
			name:     "key with positive value",
			input:    map[string]int{"a": 1},
			key:      "a",
			expected: true,
		},
		{
			name:     "unicode key exists",
			input:    map[string]int{"こんにちは": 1, "世界": 2},
			key:      "こんにちは",
			expected: true,
		},
		{
			name:     "unicode key does not exist",
			input:    map[string]int{"こんにちは": 1},
			key:      "foo",
			expected: false,
		},
		{
			name:     "emoji key exists",
			input:    map[string]int{"😊": 1, "👋": 2},
			key:      "😊",
			expected: true,
		},
		{
			name:     "emoji key does not exist",
			input:    map[string]int{"😊": 1},
			key:      "🚀",
			expected: false,
		},
		{
			name:     "special character key exists",
			input:    map[string]int{"!@#": 1, "$%": 2},
			key:      "!@#",
			expected: true,
		},
		{
			name:     "special character key does not exist",
			input:    map[string]int{"!@#": 1},
			key:      "$%",
			expected: false,
		},
		{
			name:     "whitespace key exists",
			input:    map[string]int{" ": 1, "\t": 2},
			key:      " ",
			expected: true,
		},
		{
			name:     "whitespace key does not exist",
			input:    map[string]int{" ": 1},
			key:      "\t",
			expected: false,
		},
		{
			name:     "newline key exists",
			input:    map[string]int{"\n": 1},
			key:      "\n",
			expected: true,
		},
		{
			name:     "numeric string key exists",
			input:    map[string]int{"1": 1, "2": 2},
			key:      "1",
			expected: true,
		},
		{
			name:     "numeric string key does not exist",
			input:    map[string]int{"1": 1},
			key:      "2",
			expected: false,
		},
		{
			name:     "case sensitive key exists",
			input:    map[string]int{"Go": 1},
			key:      "Go",
			expected: true,
		},
		{
			name:     "case sensitive key does not exist",
			input:    map[string]int{"Go": 1},
			key:      "go",
			expected: false,
		},
		{
			name:     "long key exists",
			input:    map[string]int{"long_key_example_0001": 1},
			key:      "long_key_example_0001",
			expected: true,
		},
		{
			name:     "long key does not exist",
			input:    map[string]int{"long_key_example_0001": 1},
			key:      "long_key_example_0002",
			expected: false,
		},
		{
			name:     "many entries key exists",
			input:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			key:      "c",
			expected: true,
		},
		{
			name:     "many entries key does not exist",
			input:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			key:      "f",
			expected: false,
		},
		{
			name:     "empty string key exists",
			input:    map[string]int{"": 1},
			key:      "",
			expected: true,
		},
		{
			name:     "empty string key does not exist",
			input:    map[string]int{"a": 1},
			key:      "",
			expected: false,
		},
		{
			name:     "large value key exists",
			input:    map[string]int{"big": 1000000},
			key:      "big",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Contains(tt.input, tt.key)
			if got != tt.expected {
				t.Fatalf("Contains(%v, %q) = %v, want %v", tt.input, tt.key, got, tt.expected)
			}
		})
	}
}

func TestContains_IntString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		key      int
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			key:      1,
			expected: false,
		},
		{
			name:     "empty map",
			input:    map[int]string{},
			key:      1,
			expected: false,
		},
		{
			name:     "single entry key exists",
			input:    map[int]string{1: "one"},
			key:      1,
			expected: true,
		},
		{
			name:     "single entry key does not exist",
			input:    map[int]string{1: "one"},
			key:      2,
			expected: false,
		},
		{
			name:     "multiple entries key exists",
			input:    map[int]string{1: "one", 2: "two", 3: "three"},
			key:      2,
			expected: true,
		},
		{
			name:     "multiple entries key does not exist",
			input:    map[int]string{1: "one", 2: "two", 3: "three"},
			key:      4,
			expected: false,
		},
		{
			name:     "key with empty string value",
			input:    map[int]string{1: ""},
			key:      1,
			expected: false,
		},
		{
			name:     "key with non-empty string value",
			input:    map[int]string{1: "one"},
			key:      1,
			expected: true,
		},
		{
			name:     "zero key exists",
			input:    map[int]string{0: "zero"},
			key:      0,
			expected: true,
		},
		{
			name:     "zero key does not exist",
			input:    map[int]string{1: "one"},
			key:      0,
			expected: false,
		},
		{
			name:     "negative key exists",
			input:    map[int]string{-1: "neg"},
			key:      -1,
			expected: true,
		},
		{
			name:     "negative key does not exist",
			input:    map[int]string{-1: "neg"},
			key:      1,
			expected: false,
		},
		{
			name:     "large key exists",
			input:    map[int]string{1000000: "million"},
			key:      1000000,
			expected: true,
		},
		{
			name:     "large key does not exist",
			input:    map[int]string{1000000: "million"},
			key:      999999,
			expected: false,
		},
		{
			name:     "mixed sign keys exist",
			input:    map[int]string{-5: "neg", 0: "zero", 5: "pos"},
			key:      0,
			expected: true,
		},
		{
			name:     "many entries key exists",
			input:    map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"},
			key:      3,
			expected: true,
		},
		{
			name:     "many entries key does not exist",
			input:    map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"},
			key:      6,
			expected: false,
		},
		{
			name:     "unicode value key exists",
			input:    map[int]string{1: "こんにちは"},
			key:      1,
			expected: true,
		},
		{
			name:     "emoji value key exists",
			input:    map[int]string{1: "😊"},
			key:      1,
			expected: true,
		},
		{
			name:     "whitespace value key exists",
			input:    map[int]string{1: " "},
			key:      1,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Contains(tt.input, tt.key)
			if got != tt.expected {
				t.Fatalf("Contains(%v, %d) = %v, want %v", tt.input, tt.key, got, tt.expected)
			}
		})
	}
}

func TestContains_StructKey(t *testing.T) {
	type Key struct {
		ID   int
		Name string
	}

	tests := []struct {
		name     string
		input    map[Key]string
		key      Key
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			key:      Key{ID: 1, Name: "test"},
			expected: false,
		},
		{
			name:     "empty map",
			input:    map[Key]string{},
			key:      Key{ID: 1, Name: "test"},
			expected: false,
		},
		{
			name:     "struct key exists",
			input:    map[Key]string{{ID: 1, Name: "alice"}: "value1"},
			key:      Key{ID: 1, Name: "alice"},
			expected: true,
		},
		{
			name:     "struct key does not exist",
			input:    map[Key]string{{ID: 1, Name: "alice"}: "value1"},
			key:      Key{ID: 2, Name: "bob"},
			expected: false,
		},
		{
			name: "multiple struct keys one exists",
			input: map[Key]string{
				{ID: 1, Name: "alice"}: "value1",
				{ID: 2, Name: "bob"}:   "value2",
			},
			key:      Key{ID: 1, Name: "alice"},
			expected: true,
		},
		{
			name: "multiple struct keys one does not exist",
			input: map[Key]string{
				{ID: 1, Name: "alice"}: "value1",
				{ID: 2, Name: "bob"}:   "value2",
			},
			key:      Key{ID: 3, Name: "charlie"},
			expected: false,
		},
		{
			name:     "struct key with empty value",
			input:    map[Key]string{{ID: 1, Name: "test"}: ""},
			key:      Key{ID: 1, Name: "test"},
			expected: false,
		},
		{
			name:     "struct key with non-empty value",
			input:    map[Key]string{{ID: 1, Name: "test"}: "value"},
			key:      Key{ID: 1, Name: "test"},
			expected: true,
		},
		{
			name:     "unicode struct field exists",
			input:    map[Key]string{{ID: 1, Name: "こんにちは"}: "value"},
			key:      Key{ID: 1, Name: "こんにちは"},
			expected: true,
		},
		{
			name:     "case sensitive struct field",
			input:    map[Key]string{{ID: 1, Name: "Test"}: "value"},
			key:      Key{ID: 1, Name: "test"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Contains(tt.input, tt.key)
			if got != tt.expected {
				t.Fatalf("Contains(%v, %v) = %v, want %v", tt.input, tt.key, got, tt.expected)
			}
		})
	}
}

func TestContains_InterfaceValue(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		key      string
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			key:      "key",
			expected: false,
		},
		{
			name:     "key with nil value",
			input:    map[string]interface{}{"key": nil},
			key:      "key",
			expected: false,
		},
		{
			name:     "key with string value",
			input:    map[string]interface{}{"key": "value"},
			key:      "key",
			expected: true,
		},
		{
			name:     "key with int value",
			input:    map[string]interface{}{"key": 42},
			key:      "key",
			expected: true,
		},
		{
			name:     "key with bool value",
			input:    map[string]interface{}{"key": true},
			key:      "key",
			expected: true,
		},
		{
			name:     "key with float value",
			input:    map[string]interface{}{"key": 3.14},
			key:      "key",
			expected: true,
		},
		{
			name:     "key with empty string value",
			input:    map[string]interface{}{"key": ""},
			key:      "key",
			expected: false,
		},
		{
			name:     "key with zero int value",
			input:    map[string]interface{}{"key": 0},
			key:      "key",
			expected: false,
		},
		{
			name:     "key with false bool value",
			input:    map[string]interface{}{"key": false},
			key:      "key",
			expected: false,
		},
		{
			name:     "key with slice value",
			input:    map[string]interface{}{"key": []int{1, 2, 3}},
			key:      "key",
			expected: true,
		},
		{
			name:     "key with empty slice value",
			input:    map[string]interface{}{"key": []int{}},
			key:      "key",
			expected: false,
		},
		{
			name:     "key with map value",
			input:    map[string]interface{}{"key": map[string]string{"a": "b"}},
			key:      "key",
			expected: true,
		},
		{
			name:     "key with empty map value",
			input:    map[string]interface{}{"key": map[string]string{}},
			key:      "key",
			expected: false,
		},
		{
			name:     "multiple keys with mixed values",
			input:    map[string]interface{}{"a": "value", "b": 42, "c": nil},
			key:      "a",
			expected: true,
		},
		{
			name:     "multiple keys with mixed values key not found",
			input:    map[string]interface{}{"a": "value", "b": 42},
			key:      "c",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Contains(tt.input, tt.key)
			if got != tt.expected {
				t.Fatalf("Contains(%v, %q) = %v, want %v", tt.input, tt.key, got, tt.expected)
			}
		})
	}
}

func TestContains_StringPointer(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	tests := []struct {
		name     string
		input    map[string]*Person
		key      string
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			key:      "alice",
			expected: false,
		},
		{
			name:     "empty map",
			input:    map[string]*Person{},
			key:      "alice",
			expected: false,
		},
		{
			name:     "key with non-nil pointer value",
			input:    map[string]*Person{"alice": {Name: "Alice", Age: 25}},
			key:      "alice",
			expected: true,
		},
		{
			name:     "key with nil pointer value",
			input:    map[string]*Person{"alice": nil},
			key:      "alice",
			expected: false,
		},
		{
			name:     "multiple entries key with non-nil pointer",
			input:    map[string]*Person{"alice": {Name: "Alice", Age: 25}, "bob": nil},
			key:      "alice",
			expected: true,
		},
		{
			name:     "multiple entries key with nil pointer",
			input:    map[string]*Person{"alice": {Name: "Alice", Age: 25}, "bob": nil},
			key:      "bob",
			expected: false,
		},
		{
			name:     "multiple entries key does not exist",
			input:    map[string]*Person{"alice": {Name: "Alice", Age: 25}, "bob": nil},
			key:      "charlie",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Contains(tt.input, tt.key)
			if got != tt.expected {
				t.Fatalf("Contains(%v, %q) = %v, want %v", tt.input, tt.key, got, tt.expected)
			}
		})
	}
}

func TestContains_BoolValue(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]bool
		key      string
		expected bool
	}{
		{
			name:     "nil map",
			input:    nil,
			key:      "key",
			expected: false,
		},
		{
			name:     "key with true value",
			input:    map[string]bool{"key": true},
			key:      "key",
			expected: true,
		},
		{
			name:     "key with false value",
			input:    map[string]bool{"key": false},
			key:      "key",
			expected: false,
		},
		{
			name:     "multiple entries true value",
			input:    map[string]bool{"a": true, "b": false},
			key:      "a",
			expected: true,
		},
		{
			name:     "multiple entries false value",
			input:    map[string]bool{"a": true, "b": false},
			key:      "b",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Contains(tt.input, tt.key)
			if got != tt.expected {
				t.Fatalf("Contains(%v, %q) = %v, want %v", tt.input, tt.key, got, tt.expected)
			}
		})
	}
}
