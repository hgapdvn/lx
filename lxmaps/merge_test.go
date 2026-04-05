package lxmaps_test

import (
	"testing"

	"github.com/hgapdvn/lx/lxmaps"
)

func TestMerge_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[string]int
		expected map[string]int
	}{
		{
			name:     "no args",
			input:    nil,
			expected: map[string]int{},
		},
		{
			name:     "single nil map",
			input:    []map[string]int{nil},
			expected: map[string]int{},
		},
		{
			name:     "single empty map",
			input:    []map[string]int{{}},
			expected: map[string]int{},
		},
		{
			name:     "single map with one entry",
			input:    []map[string]int{{"a": 1}},
			expected: map[string]int{"a": 1},
		},
		{
			name:     "single map with multiple entries",
			input:    []map[string]int{{"a": 1, "b": 2, "c": 3}},
			expected: map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name: "two maps no overlap",
			input: []map[string]int{
				{"a": 1, "b": 2},
				{"c": 3, "d": 4},
			},
			expected: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
		},
		{
			name: "two maps with overlap last wins",
			input: []map[string]int{
				{"a": 1, "b": 2},
				{"b": 20, "c": 3},
			},
			expected: map[string]int{"a": 1, "b": 20, "c": 3},
		},
		{
			name: "three maps progressive override",
			input: []map[string]int{
				{"a": 1, "b": 2},
				{"b": 20, "c": 3},
				{"c": 30, "d": 4},
			},
			expected: map[string]int{"a": 1, "b": 20, "c": 30, "d": 4},
		},
		{
			name: "multiple maps same key override",
			input: []map[string]int{
				{"key": 1},
				{"key": 2},
				{"key": 3},
			},
			expected: map[string]int{"key": 3},
		},
		{
			name: "empty maps mixed",
			input: []map[string]int{
				{},
				{"a": 1},
				{},
				{"b": 2},
				{},
			},
			expected: map[string]int{"a": 1, "b": 2},
		},
		{
			name: "nil and empty maps",
			input: []map[string]int{
				nil,
				{},
				{"a": 1},
				nil,
			},
			expected: map[string]int{"a": 1},
		},
		{
			name: "unicode keys",
			input: []map[string]int{
				{"こんにちは": 1, "世界": 2},
				{"こんにちは": 10, "foo": 3},
			},
			expected: map[string]int{"こんにちは": 10, "世界": 2, "foo": 3},
		},
		{
			name: "emoji keys",
			input: []map[string]int{
				{"😊": 1, "👋": 2},
				{"😊": 10, "🚀": 3},
			},
			expected: map[string]int{"😊": 10, "👋": 2, "🚀": 3},
		},
		{
			name: "special character keys",
			input: []map[string]int{
				{"!@#": 1, "$%": 2},
				{"!@#": 10, "&*": 3},
			},
			expected: map[string]int{"!@#": 10, "$%": 2, "&*": 3},
		},
		{
			name: "whitespace keys",
			input: []map[string]int{
				{" ": 1, "\t": 2},
				{" ": 10, "\n": 3},
			},
			expected: map[string]int{" ": 10, "\t": 2, "\n": 3},
		},
		{
			name: "numeric string keys",
			input: []map[string]int{
				{"1": 1, "2": 2},
				{"2": 20, "3": 3},
			},
			expected: map[string]int{"1": 1, "2": 20, "3": 3},
		},
		{
			name: "case sensitive keys",
			input: []map[string]int{
				{"Go": 1, "go": 2},
				{"Go": 10, "GO": 3},
			},
			expected: map[string]int{"Go": 10, "go": 2, "GO": 3},
		},
		{
			name: "long string keys",
			input: []map[string]int{
				{"long_key_example_0001": 1, "long_key_example_0002": 2},
				{"long_key_example_0001": 10, "long_key_example_0003": 3},
			},
			expected: map[string]int{"long_key_example_0001": 10, "long_key_example_0002": 2, "long_key_example_0003": 3},
		},
		{
			name: "negative values",
			input: []map[string]int{
				{"a": -1, "b": -2},
				{"b": -20, "c": -3},
			},
			expected: map[string]int{"a": -1, "b": -20, "c": -3},
		},
		{
			name: "zero values",
			input: []map[string]int{
				{"a": 0, "b": 1},
				{"b": 0, "c": 2},
			},
			expected: map[string]int{"a": 0, "b": 0, "c": 2},
		},
		{
			name: "many maps",
			input: []map[string]int{
				{"a": 1},
				{"b": 2},
				{"c": 3},
				{"d": 4},
				{"e": 5},
			},
			expected: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
		},
		{
			name: "large values",
			input: []map[string]int{
				{"big1": 1000000},
				{"big2": 999999},
				{"big1": 2000000},
			},
			expected: map[string]int{"big1": 2000000, "big2": 999999},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Merge(tt.input...)

			if len(got) != len(tt.expected) {
				t.Fatalf("Merge() returned %d entries, want %d; got=%v, want=%v", len(got), len(tt.expected), got, tt.expected)
			}

			for k, expectedV := range tt.expected {
				gotV, ok := got[k]
				if !ok {
					t.Fatalf("Merge() missing key %q in result; got=%v", k, got)
				}
				if gotV != expectedV {
					t.Fatalf("Merge() for key %q: got %d, want %d", k, gotV, expectedV)
				}
			}

			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("Merge() has unexpected key %q in result; got=%v", k, got)
				}
			}
		})
	}
}

func TestMerge_IntString(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[int]string
		expected map[int]string
	}{
		{
			name:     "no args",
			input:    nil,
			expected: map[int]string{},
		},
		{
			name:     "single nil map",
			input:    []map[int]string{nil},
			expected: map[int]string{},
		},
		{
			name:     "single empty map",
			input:    []map[int]string{},
			expected: map[int]string{},
		},
		{
			name:     "single map with one entry",
			input:    []map[int]string{{1: "one"}},
			expected: map[int]string{1: "one"},
		},
		{
			name:     "single map with multiple entries",
			input:    []map[int]string{{1: "one", 2: "two", 3: "three"}},
			expected: map[int]string{1: "one", 2: "two", 3: "three"},
		},
		{
			name: "two maps no overlap",
			input: []map[int]string{
				{1: "one", 2: "two"},
				{3: "three", 4: "four"},
			},
			expected: map[int]string{1: "one", 2: "two", 3: "three", 4: "four"},
		},
		{
			name: "two maps with overlap last wins",
			input: []map[int]string{
				{1: "one", 2: "two"},
				{2: "TWO", 3: "three"},
			},
			expected: map[int]string{1: "one", 2: "TWO", 3: "three"},
		},
		{
			name: "three maps progressive override",
			input: []map[int]string{
				{1: "one", 2: "two"},
				{2: "TWO", 3: "three"},
				{3: "THREE", 4: "four"},
			},
			expected: map[int]string{1: "one", 2: "TWO", 3: "THREE", 4: "four"},
		},
		{
			name: "multiple maps same key override",
			input: []map[int]string{
				{100: "first"},
				{100: "second"},
				{100: "third"},
			},
			expected: map[int]string{100: "third"},
		},
		{
			name: "empty maps mixed",
			input: []map[int]string{
				{},
				{1: "one"},
				{},
				{2: "two"},
				{},
			},
			expected: map[int]string{1: "one", 2: "two"},
		},
		{
			name: "nil and empty maps",
			input: []map[int]string{
				nil,
				{},
				{1: "one"},
				nil,
			},
			expected: map[int]string{1: "one"},
		},
		{
			name: "negative keys",
			input: []map[int]string{
				{-1: "neg", -2: "neg2"},
				{-1: "NEG", 1: "pos"},
			},
			expected: map[int]string{-1: "NEG", -2: "neg2", 1: "pos"},
		},
		{
			name: "zero key",
			input: []map[int]string{
				{0: "zero"},
				{0: "ZERO", 1: "one"},
			},
			expected: map[int]string{0: "ZERO", 1: "one"},
		},
		{
			name: "unicode values",
			input: []map[int]string{
				{1: "こんにちは", 2: "世界"},
				{1: "HELLO", 3: "foo"},
			},
			expected: map[int]string{1: "HELLO", 2: "世界", 3: "foo"},
		},
		{
			name: "emoji values",
			input: []map[int]string{
				{1: "😊", 2: "👋"},
				{1: "🚀", 3: "hello"},
			},
			expected: map[int]string{1: "🚀", 2: "👋", 3: "hello"},
		},
		{
			name: "empty string values",
			input: []map[int]string{
				{1: "", 2: "value"},
				{2: "", 3: "three"},
			},
			expected: map[int]string{1: "", 2: "", 3: "three"},
		},
		{
			name: "whitespace string values",
			input: []map[int]string{
				{1: " ", 2: "\t"},
				{1: "\n", 3: "ok"},
			},
			expected: map[int]string{1: "\n", 2: "\t", 3: "ok"},
		},
		{
			name: "long string values",
			input: []map[int]string{
				{1: "long_string_example_0001", 2: "long_string_example_0002"},
				{1: "LONG_STRING_EXAMPLE_0001", 3: "long_string_example_0003"},
			},
			expected: map[int]string{1: "LONG_STRING_EXAMPLE_0001", 2: "long_string_example_0002", 3: "long_string_example_0003"},
		},
		{
			name: "large keys",
			input: []map[int]string{
				{1000000: "million", 999999: "almostmillion"},
				{1000000: "MILLION", 1000001: "morethammillion"},
			},
			expected: map[int]string{1000000: "MILLION", 999999: "almostmillion", 1000001: "morethammillion"},
		},
		{
			name: "many maps",
			input: []map[int]string{
				{1: "a"},
				{2: "b"},
				{3: "c"},
				{4: "d"},
				{5: "e"},
			},
			expected: map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Merge(tt.input...)

			if len(got) != len(tt.expected) {
				t.Fatalf("Merge() returned %d entries, want %d; got=%v, want=%v", len(got), len(tt.expected), got, tt.expected)
			}

			for k, expectedV := range tt.expected {
				gotV, ok := got[k]
				if !ok {
					t.Fatalf("Merge() missing key %d in result; got=%v", k, got)
				}
				if gotV != expectedV {
					t.Fatalf("Merge() for key %d: got %q, want %q", k, gotV, expectedV)
				}
			}

			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("Merge() has unexpected key %d in result; got=%v", k, got)
				}
			}
		})
	}
}

func TestMerge_StructValue(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	tests := []struct {
		name     string
		input    []map[string]User
		expected map[string]User
	}{
		{
			name:     "no args",
			input:    nil,
			expected: map[string]User{},
		},
		{
			name:     "single map with struct values",
			input:    []map[string]User{{"alice": {Name: "Alice", Age: 25}}},
			expected: map[string]User{"alice": {Name: "Alice", Age: 25}},
		},
		{
			name: "two maps with struct values",
			input: []map[string]User{
				{"alice": {Name: "Alice", Age: 25}, "bob": {Name: "Bob", Age: 30}},
				{"charlie": {Name: "Charlie", Age: 20}, "alice": {Name: "Alice2", Age: 26}},
			},
			expected: map[string]User{
				"alice":   {Name: "Alice2", Age: 26},
				"bob":     {Name: "Bob", Age: 30},
				"charlie": {Name: "Charlie", Age: 20},
			},
		},
		{
			name: "multiple maps override",
			input: []map[string]User{
				{"user1": {Name: "User1", Age: 10}},
				{"user1": {Name: "User1Updated", Age: 20}},
				{"user1": {Name: "User1Final", Age: 30}},
			},
			expected: map[string]User{
				"user1": {Name: "User1Final", Age: 30},
			},
		},
		{
			name: "empty and nil with struct values",
			input: []map[string]User{
				nil,
				{},
				{"user": {Name: "User", Age: 15}},
			},
			expected: map[string]User{
				"user": {Name: "User", Age: 15},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Merge(tt.input...)

			if len(got) != len(tt.expected) {
				t.Fatalf("Merge() returned %d entries, want %d", len(got), len(tt.expected))
			}

			for k, expectedV := range tt.expected {
				gotV, ok := got[k]
				if !ok {
					t.Fatalf("Merge() missing key %q in result", k)
				}
				if gotV != expectedV {
					t.Fatalf("Merge() for key %q: got %+v, want %+v", k, gotV, expectedV)
				}
			}
		})
	}
}

func TestMerge_InterfaceValue(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name:     "no args",
			input:    nil,
			expected: map[string]interface{}{},
		},
		{
			name: "mixed interface types",
			input: []map[string]interface{}{
				{"str": "hello", "int": 42, "bool": true},
				{"float": 3.14, "str": "world", "nil": nil},
			},
			expected: map[string]interface{}{
				"str":   "world",
				"int":   42,
				"bool":  true,
				"float": 3.14,
				"nil":   nil,
			},
		},
		{
			name: "interface value override",
			input: []map[string]interface{}{
				{"key": "first"},
				{"key": 42},
				{"key": true},
			},
			expected: map[string]interface{}{
				"key": true,
			},
		},
		{
			name: "nil interface values",
			input: []map[string]interface{}{
				{"a": 1, "b": nil},
				{"b": 2, "c": nil},
			},
			expected: map[string]interface{}{
				"a": 1,
				"b": 2,
				"c": nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Merge(tt.input...)

			if len(got) != len(tt.expected) {
				t.Fatalf("Merge() returned %d entries, want %d", len(got), len(tt.expected))
			}

			for k, expectedV := range tt.expected {
				gotV, ok := got[k]
				if !ok {
					t.Fatalf("Merge() missing key %q in result", k)
				}
				if gotV != expectedV {
					t.Fatalf("Merge() for key %q: got %v, want %v", k, gotV, expectedV)
				}
			}
		})
	}
}
