package lxmaps_test

import (
	"reflect"
	"testing"

	"github.com/hgapdvn/lx/maps"
)

func TestPickBy_StringInt(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]int
		predicate func(k string, v int) bool
		expected  map[string]int
	}{
		{
			name:      "nil map",
			input:     nil,
			predicate: func(k string, v int) bool { return true },
			expected:  nil,
		},
		{
			name:      "empty map",
			input:     map[string]int{},
			predicate: func(k string, v int) bool { return true },
			expected:  map[string]int{},
		},
		{
			name:      "single entry predicate matches",
			input:     map[string]int{"a": 1},
			predicate: func(k string, v int) bool { return k == "a" },
			expected:  map[string]int{"a": 1},
		},
		{
			name:      "single entry predicate does not match",
			input:     map[string]int{"a": 1},
			predicate: func(k string, v int) bool { return k == "b" },
			expected:  map[string]int{},
		},
		{
			name:      "match by value",
			input:     map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(k string, v int) bool { return v == 2 },
			expected:  map[string]int{"b": 2},
		},
		{
			name:      "match by key and value",
			input:     map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(k string, v int) bool { return k == "c" && v == 3 },
			expected:  map[string]int{"c": 3},
		},
		{
			name:      "no entry matches",
			input:     map[string]int{"a": 1, "b": 2},
			predicate: func(k string, v int) bool { return v > 10 },
			expected:  map[string]int{},
		},
		{
			name:      "zero value matched",
			input:     map[string]int{"a": 0, "b": 1},
			predicate: func(k string, v int) bool { return v == 0 },
			expected:  map[string]int{"a": 0},
		},
		{
			name:      "negative value matched",
			input:     map[string]int{"a": -1, "b": 2},
			predicate: func(k string, v int) bool { return v < 0 },
			expected:  map[string]int{"a": -1},
		},
		{
			name:      "empty string key matched",
			input:     map[string]int{"": 42, "a": 1},
			predicate: func(k string, v int) bool { return k == "" },
			expected:  map[string]int{"": 42},
		},
		{
			name:      "multiple entries same value all match predicate",
			input:     map[string]int{"x": 5, "y": 5},
			predicate: func(k string, v int) bool { return v == 5 },
			expected:  map[string]int{"x": 5, "y": 5},
		},
		{
			name:      "predicate false for all",
			input:     map[string]int{"a": 1, "b": 2},
			predicate: func(k string, v int) bool { return false },
			expected:  map[string]int{},
		},
		{
			name:      "predicate true for all single entry",
			input:     map[string]int{"only": 7},
			predicate: func(k string, v int) bool { return true },
			expected:  map[string]int{"only": 7},
		},
		{
			name:      "unicode keys match one",
			input:     map[string]int{"こんにちは": 1, "世界": 2, "test": 3},
			predicate: func(k string, v int) bool { return k == "世界" },
			expected:  map[string]int{"世界": 2},
		},
		{
			name:      "unicode key match by value",
			input:     map[string]int{"a": 1, "日": 99},
			predicate: func(k string, v int) bool { return v == 99 },
			expected:  map[string]int{"日": 99},
		},
		{
			name:      "special character keys",
			input:     map[string]int{"!@#": 10, "$%": 20, "normal": 30},
			predicate: func(k string, v int) bool { return k == "!@#" },
			expected:  map[string]int{"!@#": 10},
		},
		{
			name:      "emoji keys",
			input:     map[string]int{"😊": 1, "🚀": 2, "plain": 3},
			predicate: func(k string, v int) bool { return k == "😊" },
			expected:  map[string]int{"😊": 1},
		},
		{
			name:      "case sensitive keys",
			input:     map[string]int{"A": 1, "a": 2},
			predicate: func(k string, v int) bool { return k == "A" },
			expected:  map[string]int{"A": 1},
		},
		{
			name:      "large int value",
			input:     map[string]int{"big": 1_000_000, "small": 1},
			predicate: func(k string, v int) bool { return v == 1_000_000 },
			expected:  map[string]int{"big": 1_000_000},
		},
		{
			name:      "negative value",
			input:     map[string]int{"a": -1, "b": 2},
			predicate: func(k string, v int) bool { return v < 0 },
			expected:  map[string]int{"a": -1},
		},
		{
			name:      "empty string key",
			input:     map[string]int{"": 42, "a": 1},
			predicate: func(k string, v int) bool { return k == "" },
			expected:  map[string]int{"": 42},
		},
		{
			name:      "multiple entries same value all match predicate",
			input:     map[string]int{"x": 5, "y": 5},
			predicate: func(k string, v int) bool { return v == 5 },
			expected:  map[string]int{"x": 5, "y": 5},
		},
		{
			name:      "predicate false for all",
			input:     map[string]int{"a": 1, "b": 2},
			predicate: func(k string, v int) bool { return false },
			expected:  map[string]int{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := lxmaps.PickBy(test.input, test.predicate)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestPickBy_IntString(t *testing.T) {
	tests := []struct {
		name      string
		input     map[int]string
		predicate func(k int, v string) bool
		expected  map[int]string
	}{
		{
			name:      "nil map",
			input:     nil,
			predicate: func(k int, v string) bool { return true },
			expected:  nil,
		},
		{
			name:      "empty map",
			input:     map[int]string{},
			predicate: func(k int, v string) bool { return true },
			expected:  map[int]string{},
		},
		{
			name:      "single entry predicate matches",
			input:     map[int]string{1: "a"},
			predicate: func(k int, v string) bool { return k == 1 },
			expected:  map[int]string{1: "a"},
		},
		{
			name:      "match by value",
			input:     map[int]string{1: "a", 2: "b", 3: "c"},
			predicate: func(k int, v string) bool { return v == "b" },
			expected:  map[int]string{2: "b"},
		},
		{
			name:      "match by key and value",
			input:     map[int]string{1: "a", 2: "b", 3: "c"},
			predicate: func(k int, v string) bool { return k == 3 && v == "c" },
			expected:  map[int]string{3: "c"},
		},
		{
			name:      "no entry matches",
			input:     map[int]string{1: "a", 2: "b"},
			predicate: func(k int, v string) bool { return v > "c" },
			expected:  map[int]string{},
		},
		{
			name:      "zero value matched",
			input:     map[int]string{0: "zero", 1: "one"},
			predicate: func(k int, v string) bool { return v == "zero" },
			expected:  map[int]string{0: "zero"},
		},
		{
			name:      "negative value matched",
			input:     map[int]string{0: "zero", 1: "one"},
			predicate: func(k int, v string) bool { return k < 1 },
			expected:  map[int]string{0: "zero"},
		},
		{
			name:      "empty string key matched",
			input:     map[int]string{0: "zero", 1: "one"},
			predicate: func(k int, v string) bool { return k == 0 },
			expected:  map[int]string{0: "zero"},
		},
		{
			name:      "multiple entries same value all match predicate",
			input:     map[int]string{0: "five", 1: "five"},
			predicate: func(k int, v string) bool { return v == "five" },
			expected:  map[int]string{0: "five", 1: "five"},
		},
		{
			name:      "predicate false for all",
			input:     map[int]string{0: "one", 1: "two"},
			predicate: func(k int, v string) bool { return false },
			expected:  map[int]string{},
		},
		{
			name:      "predicate true for all single entry",
			input:     map[int]string{0: "only"},
			predicate: func(k int, v string) bool { return true },
			expected:  map[int]string{0: "only"},
		},
		{
			name:      "unicode keys match one",
			input:     map[int]string{0: "a", 1: "b", 2: "c"},
			predicate: func(k int, v string) bool { return k == 1 },
			expected:  map[int]string{1: "b"},
		},
		{
			name:      "unicode key match by value",
			input:     map[int]string{0: "a", 1: "b"},
			predicate: func(k int, v string) bool { return v == "b" },
			expected:  map[int]string{1: "b"},
		},
		{
			name:      "case sensitive keys",
			input:     map[int]string{0: "A", 1: "a"},
			predicate: func(k int, v string) bool { return k == 0 },
			expected:  map[int]string{0: "A"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := lxmaps.PickBy(test.input, test.predicate)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestPickBy_StringStruct(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	tests := []struct {
		name      string
		input     map[string]User
		predicate func(k string, v User) bool
		expected  map[string]User
	}{
		{
			name:      "nil map",
			input:     nil,
			predicate: func(k string, v User) bool { return true },
			expected:  nil,
		},
		{
			name:      "empty map",
			input:     map[string]User{},
			predicate: func(k string, v User) bool { return true },
			expected:  map[string]User{},
		},
		{
			name:      "single entry predicate matches",
			input:     map[string]User{"a": {Name: "a", Age: 1}},
			predicate: func(k string, v User) bool { return k == "a" },
			expected:  map[string]User{"a": {Name: "a", Age: 1}},
		},
		{
			name:      "single entry predicate does not match",
			input:     map[string]User{"a": {Name: "a", Age: 1}},
			predicate: func(k string, v User) bool { return k == "b" },
			expected:  map[string]User{},
		},
		{
			name:      "match by value",
			input:     map[string]User{"a": {Name: "a", Age: 1}, "b": {Name: "b", Age: 2}, "c": {Name: "c", Age: 3}},
			predicate: func(k string, v User) bool { return v.Name == "b" },
			expected:  map[string]User{"b": {Name: "b", Age: 2}},
		},
		{
			name:      "match by key and value",
			input:     map[string]User{"a": {Name: "a", Age: 1}, "b": {Name: "b", Age: 2}, "c": {Name: "c", Age: 3}},
			predicate: func(k string, v User) bool { return k == "c" && v.Name == "c" },
			expected:  map[string]User{"c": {Name: "c", Age: 3}},
		},
		{
			name:      "no entry matches",
			input:     map[string]User{"a": {Name: "a", Age: 1}, "b": {Name: "b", Age: 2}},
			predicate: func(k string, v User) bool { return v.Age > 10 },
			expected:  map[string]User{},
		},
		{
			name:      "zero value matched",
			input:     map[string]User{"a": {Name: "a", Age: 0}, "b": {Name: "b", Age: 1}},
			predicate: func(k string, v User) bool { return v.Age == 0 },
			expected:  map[string]User{"a": {Name: "a", Age: 0}},
		},
		{
			name:      "negative value matched",
			input:     map[string]User{"a": {Name: "a", Age: -1}, "b": {Name: "b", Age: 2}},
			predicate: func(k string, v User) bool { return v.Age < 0 },
			expected:  map[string]User{"a": {Name: "a", Age: -1}},
		},
		{
			name:      "empty string key matched",
			input:     map[string]User{"": {Name: "a", Age: 1}, "b": {Name: "b", Age: 2}},
			predicate: func(k string, v User) bool { return k == "" },
			expected:  map[string]User{"": {Name: "a", Age: 1}},
		},
		{
			name:      "multiple entries same value all match predicate",
			input:     map[string]User{"x": {Name: "x", Age: 5}, "y": {Name: "y", Age: 5}},
			predicate: func(k string, v User) bool { return v.Age == 5 },
			expected:  map[string]User{"x": {Name: "x", Age: 5}, "y": {Name: "y", Age: 5}},
		},
		{
			name:      "predicate false for all",
			input:     map[string]User{"a": {Name: "a", Age: 1}, "b": {Name: "b", Age: 2}},
			predicate: func(k string, v User) bool { return false },
			expected:  map[string]User{},
		},
		{
			name:      "predicate true for all single entry",
			input:     map[string]User{"only": {Name: "only", Age: 7}},
			predicate: func(k string, v User) bool { return true },
			expected:  map[string]User{"only": {Name: "only", Age: 7}},
		},
		{
			name:      "unicode keys match one",
			input:     map[string]User{"こんにちは": {Name: "こんにちは", Age: 1}, "世界": {Name: "世界", Age: 2}, "test": {Name: "test", Age: 3}},
			predicate: func(k string, v User) bool { return k == "世界" },
			expected:  map[string]User{"世界": {Name: "世界", Age: 2}},
		},
		{
			name:      "unicode key match by value",
			input:     map[string]User{"a": {Name: "a", Age: 1}, "日": {Name: "日", Age: 99}},
			predicate: func(k string, v User) bool { return v.Age == 99 },
			expected:  map[string]User{"日": {Name: "日", Age: 99}},
		},
		{
			name:      "case sensitive keys",
			input:     map[string]User{"A": {Name: "A", Age: 1}, "a": {Name: "a", Age: 2}},
			predicate: func(k string, v User) bool { return k == "A" },
			expected:  map[string]User{"A": {Name: "A", Age: 1}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := lxmaps.PickBy(test.input, test.predicate)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func BenchmarkPickBy(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	predicate := func(k string, v int) bool { return v > 1 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lxmaps.PickBy(m, predicate)
	}
}
