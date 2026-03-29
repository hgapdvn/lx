package lxmaps_test

import (
	"testing"

	"github.com/hgapdvn/lx/maps"
	"github.com/hgapdvn/lx/slices"
)

func TestGetBy_StringInt(t *testing.T) {
	tests := []struct {
		name          string
		input         map[string]int
		predicate     func(k string, v int) bool
		expectedValue int
		expectedFound bool
	}{
		{
			name:          "nil map",
			input:         nil,
			predicate:     func(k string, v int) bool { return true },
			expectedValue: 0,
			expectedFound: false,
		},
		{
			name:          "empty map",
			input:         map[string]int{},
			predicate:     func(k string, v int) bool { return true },
			expectedValue: 0,
			expectedFound: false,
		},
		{
			name:          "single entry predicate matches",
			input:         map[string]int{"a": 1},
			predicate:     func(k string, v int) bool { return k == "a" },
			expectedValue: 1,
			expectedFound: true,
		},
		{
			name:          "single entry predicate does not match",
			input:         map[string]int{"a": 1},
			predicate:     func(k string, v int) bool { return k == "b" },
			expectedValue: 0,
			expectedFound: false,
		},
		{
			name:          "match by value",
			input:         map[string]int{"a": 1, "b": 2, "c": 3},
			predicate:     func(k string, v int) bool { return v == 2 },
			expectedValue: 2,
			expectedFound: true,
		},
		{
			name:          "match by key and value",
			input:         map[string]int{"a": 1, "b": 2, "c": 3},
			predicate:     func(k string, v int) bool { return k == "c" && v == 3 },
			expectedValue: 3,
			expectedFound: true,
		},
		{
			name:          "no entry matches",
			input:         map[string]int{"a": 1, "b": 2},
			predicate:     func(k string, v int) bool { return v > 10 },
			expectedValue: 0,
			expectedFound: false,
		},
		{
			name:          "zero value matched",
			input:         map[string]int{"a": 0, "b": 1},
			predicate:     func(k string, v int) bool { return v == 0 },
			expectedValue: 0,
			expectedFound: true,
		},
		{
			name:          "negative value matched",
			input:         map[string]int{"a": -1, "b": 2},
			predicate:     func(k string, v int) bool { return v < 0 },
			expectedValue: -1,
			expectedFound: true,
		},
		{
			name:          "empty string key matched",
			input:         map[string]int{"": 42, "a": 1},
			predicate:     func(k string, v int) bool { return k == "" },
			expectedValue: 42,
			expectedFound: true,
		},
		{
			name:          "multiple entries same value all match predicate",
			input:         map[string]int{"x": 5, "y": 5},
			predicate:     func(k string, v int) bool { return v == 5 },
			expectedValue: 5,
			expectedFound: true,
		},
		{
			name:          "predicate false for all",
			input:         map[string]int{"a": 1, "b": 2},
			predicate:     func(k string, v int) bool { return false },
			expectedValue: 0,
			expectedFound: false,
		},
		{
			name:          "predicate true for all single entry",
			input:         map[string]int{"only": 7},
			predicate:     func(k string, v int) bool { return true },
			expectedValue: 7,
			expectedFound: true,
		},
		{
			name:          "unicode keys match one",
			input:         map[string]int{"こんにちは": 1, "世界": 2, "test": 3},
			predicate:     func(k string, v int) bool { return k == "世界" },
			expectedValue: 2,
			expectedFound: true,
		},
		{
			name:          "unicode key match by value",
			input:         map[string]int{"a": 1, "日": 99},
			predicate:     func(k string, v int) bool { return v == 99 },
			expectedValue: 99,
			expectedFound: true,
		},
		{
			name:          "special character keys",
			input:         map[string]int{"!@#": 10, "$%": 20, "normal": 30},
			predicate:     func(k string, v int) bool { return k == "!@#" },
			expectedValue: 10,
			expectedFound: true,
		},
		{
			name:          "large int value",
			input:         map[string]int{"big": 1_000_000, "small": 1},
			predicate:     func(k string, v int) bool { return v == 1_000_000 },
			expectedValue: 1_000_000,
			expectedFound: true,
		},
		{
			name:          "match by key prefix unique",
			input:         map[string]int{"apple": 1, "apricot": 2, "banana": 3},
			predicate:     func(k string, v int) bool { return len(k) >= 6 && k[0] == 'a' },
			expectedValue: 2,
			expectedFound: true,
		},
		{
			name:          "many entries match specific key",
			input:         map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			predicate:     func(k string, v int) bool { return k == "d" },
			expectedValue: 4,
			expectedFound: true,
		},
		{
			name:          "match parity single even value",
			input:         map[string]int{"a": 1, "b": 2, "c": 3},
			predicate:     func(k string, v int) bool { return v%2 == 0 },
			expectedValue: 2,
			expectedFound: true,
		},
		{
			name:          "max int value",
			input:         map[string]int{"x": int(^uint(0) >> 1)},
			predicate:     func(k string, v int) bool { return v > 0 },
			expectedValue: int(^uint(0) >> 1),
			expectedFound: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := lxmaps.GetBy(tt.input, tt.predicate)
			if found != tt.expectedFound {
				t.Errorf("GetBy() found = %v, expected %v", found, tt.expectedFound)
			}
			if value != tt.expectedValue {
				t.Errorf("GetBy() value = %v, expected %v", value, tt.expectedValue)
			}
		})
	}
}

func TestGetBy_IntString(t *testing.T) {
	tests := []struct {
		name          string
		input         map[int]string
		predicate     func(k int, v string) bool
		expectedValue string
		expectedFound bool
	}{
		{
			name:          "nil map",
			input:         nil,
			predicate:     func(k int, v string) bool { return true },
			expectedValue: "",
			expectedFound: false,
		},
		{
			name:          "match by key",
			input:         map[int]string{1: "one", 2: "two"},
			predicate:     func(k int, v string) bool { return k == 2 },
			expectedValue: "two",
			expectedFound: true,
		},
		{
			name:          "match by value length",
			input:         map[int]string{0: "a", 1: "hello", 2: "hi"},
			predicate:     func(k int, v string) bool { return len(v) > 3 },
			expectedValue: "hello",
			expectedFound: true,
		},
		{
			name:          "no match",
			input:         map[int]string{1: "a"},
			predicate:     func(k int, v string) bool { return k == 99 },
			expectedValue: "",
			expectedFound: false,
		},
		{
			name:          "empty map",
			input:         map[int]string{},
			predicate:     func(k int, v string) bool { return true },
			expectedValue: "",
			expectedFound: false,
		},
		{
			name:          "zero key",
			input:         map[int]string{0: "zero", 1: "one"},
			predicate:     func(k int, v string) bool { return k == 0 },
			expectedValue: "zero",
			expectedFound: true,
		},
		{
			name:          "negative key",
			input:         map[int]string{-1: "minus", 1: "plus"},
			predicate:     func(k int, v string) bool { return k < 0 },
			expectedValue: "minus",
			expectedFound: true,
		},
		{
			name:          "empty string value matched",
			input:         map[int]string{1: "", 2: "b"},
			predicate:     func(k int, v string) bool { return v == "" },
			expectedValue: "",
			expectedFound: true,
		},
		{
			name:          "unicode value",
			input:         map[int]string{1: "hello", 2: "こんにちは"},
			predicate:     func(k int, v string) bool { return k == 2 },
			expectedValue: "こんにちは",
			expectedFound: true,
		},
		{
			name:          "emoji value",
			input:         map[int]string{1: "😊", 2: "text"},
			predicate:     func(k int, v string) bool { return k == 1 },
			expectedValue: "😊",
			expectedFound: true,
		},
		{
			name:          "large key",
			input:         map[int]string{1: "small", 1_000_000: "big"},
			predicate:     func(k int, v string) bool { return k == 1_000_000 },
			expectedValue: "big",
			expectedFound: true,
		},
		{
			name:          "filter even keys unique value",
			input:         map[int]string{1: "odd", 2: "even", 3: "odd2"},
			predicate:     func(k int, v string) bool { return k%2 == 0 },
			expectedValue: "even",
			expectedFound: true,
		},
		{
			name:          "predicate true for all single entry",
			input:         map[int]string{42: "answer"},
			predicate:     func(k int, v string) bool { return true },
			expectedValue: "answer",
			expectedFound: true,
		},
		{
			name:          "predicate false for all",
			input:         map[int]string{1: "a", 2: "b"},
			predicate:     func(k int, v string) bool { return false },
			expectedValue: "",
			expectedFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := lxmaps.GetBy(tt.input, tt.predicate)
			if found != tt.expectedFound {
				t.Errorf("GetBy() found = %v, expected %v", found, tt.expectedFound)
			}
			if value != tt.expectedValue {
				t.Errorf("GetBy() value = %v, expected %v", value, tt.expectedValue)
			}
		})
	}
}

func TestGetBy_StringBool(t *testing.T) {
	tests := []struct {
		name          string
		input         map[string]bool
		predicate     func(k string, v bool) bool
		expectedValue bool
		expectedFound bool
	}{
		{
			name:          "nil map",
			input:         nil,
			predicate:     func(k string, v bool) bool { return true },
			expectedValue: false,
			expectedFound: false,
		},
		{
			name:          "empty map",
			input:         map[string]bool{},
			predicate:     func(k string, v bool) bool { return true },
			expectedValue: false,
			expectedFound: false,
		},
		{
			name:          "match true",
			input:         map[string]bool{"a": true, "b": false},
			predicate:     func(k string, v bool) bool { return k == "a" },
			expectedValue: true,
			expectedFound: true,
		},
		{
			name:          "match false value still found",
			input:         map[string]bool{"a": false, "b": true},
			predicate:     func(k string, v bool) bool { return k == "a" },
			expectedValue: false,
			expectedFound: true,
		},
		{
			name:          "match by value true",
			input:         map[string]bool{"x": false, "y": true},
			predicate:     func(k string, v bool) bool { return v },
			expectedValue: true,
			expectedFound: true,
		},
		{
			name:          "match by value false only",
			input:         map[string]bool{"x": true, "y": false},
			predicate:     func(k string, v bool) bool { return !v },
			expectedValue: false,
			expectedFound: true,
		},
		{
			name:          "no match",
			input:         map[string]bool{"a": true},
			predicate:     func(k string, v bool) bool { return k == "missing" },
			expectedValue: false,
			expectedFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := lxmaps.GetBy(tt.input, tt.predicate)
			if found != tt.expectedFound {
				t.Errorf("GetBy() found = %v, expected %v", found, tt.expectedFound)
			}
			if value != tt.expectedValue {
				t.Errorf("GetBy() value = %v, expected %v", value, tt.expectedValue)
			}
		})
	}
}

func TestGetBy_StringStruct(t *testing.T) {
	type user struct {
		Name string
		Age  int
	}

	tests := []struct {
		name          string
		input         map[string]user
		predicate     func(k string, v user) bool
		expectedValue user
		expectedFound bool
	}{
		{
			name:          "nil map",
			input:         nil,
			predicate:     func(k string, v user) bool { return true },
			expectedValue: user{},
			expectedFound: false,
		},
		{
			name:          "match by struct field",
			input:         map[string]user{"alice": {Name: "Alice", Age: 25}, "bob": {Name: "Bob", Age: 30}},
			predicate:     func(k string, v user) bool { return v.Age == 30 },
			expectedValue: user{Name: "Bob", Age: 30},
			expectedFound: true,
		},
		{
			name:          "match by key and field",
			input:         map[string]user{"u1": {Name: "A", Age: 1}, "u2": {Name: "B", Age: 2}},
			predicate:     func(k string, v user) bool { return k == "u1" && v.Name == "A" },
			expectedValue: user{Name: "A", Age: 1},
			expectedFound: true,
		},
		{
			name:          "zero struct matched",
			input:         map[string]user{"empty": {}, "full": {Name: "x", Age: 1}},
			predicate:     func(k string, v user) bool { return v.Name == "" && v.Age == 0 },
			expectedValue: user{},
			expectedFound: true,
		},
		{
			name:          "no match",
			input:         map[string]user{"a": {Name: "A", Age: 1}},
			predicate:     func(k string, v user) bool { return v.Age > 10 },
			expectedValue: user{},
			expectedFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := lxmaps.GetBy(tt.input, tt.predicate)
			if found != tt.expectedFound {
				t.Errorf("GetBy() found = %v, expected %v", found, tt.expectedFound)
			}
			if value != tt.expectedValue {
				t.Errorf("GetBy() value = %+v, expected %+v", value, tt.expectedValue)
			}
		})
	}
}

func TestGetBy_IntString_AmbiguousMatch(t *testing.T) {
	// When several entries satisfy the predicate with different values, map iteration
	// order is undefined; assert found and that the value is one of the matching entries.
	tests := []struct {
		name          string
		input         map[int]string
		predicate     func(k int, v string) bool
		expectedFound bool
		allowedValues []string
	}{
		{
			name: "multiple matches different values",
			input: map[int]string{
				1: "first",
				2: "second",
				3: "third",
			},
			predicate:     func(k int, v string) bool { return k >= 1 && k <= 3 },
			expectedFound: true,
			allowedValues: []string{"first", "second", "third"},
		},
		{
			name: "multiple matches subset by value",
			input: map[int]string{
				1: "a",
				2: "b",
				3: "c",
			},
			predicate:     func(k int, v string) bool { return v == "a" || v == "c" },
			expectedFound: true,
			allowedValues: []string{"a", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := lxmaps.GetBy(tt.input, tt.predicate)
			if found != tt.expectedFound {
				t.Errorf("GetBy() found = %v, want %v", found, tt.expectedFound)
			}
			if !found {
				return
			}
			if !lxslices.Contains(tt.allowedValues, value) {
				t.Errorf("GetBy() value = %q, want one of %v", value, tt.allowedValues)
			}
		})
	}
}
