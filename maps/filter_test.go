package lxmaps_test

import (
	"testing"

	"github.com/hgapdvn/lx/maps"
)

func TestFilter_StringInt(t *testing.T) {
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
			expected:  map[string]int{},
		},
		{
			name:      "empty map",
			input:     map[string]int{},
			predicate: func(k string, v int) bool { return true },
			expected:  map[string]int{},
		},
		{
			name:      "single element match",
			input:     map[string]int{"a": 1},
			predicate: func(k string, v int) bool { return v > 0 },
			expected:  map[string]int{"a": 1},
		},
		{
			name:      "single element no match",
			input:     map[string]int{"a": 1},
			predicate: func(k string, v int) bool { return v > 10 },
			expected:  map[string]int{},
		},
		{
			name:      "filter by value even",
			input:     map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			predicate: func(k string, v int) bool { return v%2 == 0 },
			expected:  map[string]int{"b": 2, "d": 4},
		},
		{
			name:      "filter by key prefix",
			input:     map[string]int{"prefix_a": 1, "prefix_b": 2, "other": 3},
			predicate: func(k string, v int) bool { return len(k) > 5 },
			expected:  map[string]int{"prefix_a": 1, "prefix_b": 2},
		},
		{
			name:      "filter by both key and value",
			input:     map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			predicate: func(k string, v int) bool { return len(k) == 1 && v > 1 },
			expected:  map[string]int{"b": 2, "c": 3, "d": 4},
		},
		{
			name:      "match all",
			input:     map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(k string, v int) bool { return true },
			expected:  map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name:      "match none",
			input:     map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(k string, v int) bool { return false },
			expected:  map[string]int{},
		},
		{
			name:      "filter negative values",
			input:     map[string]int{"neg1": -5, "zero": 0, "pos1": 5},
			predicate: func(k string, v int) bool { return v < 0 },
			expected:  map[string]int{"neg1": -5},
		},
		{
			name:      "filter with unicode keys",
			input:     map[string]int{"こんにちは": 1, "世界": 2, "test": 3},
			predicate: func(k string, v int) bool { return v < 3 },
			expected:  map[string]int{"こんにちは": 1, "世界": 2},
		},
		{
			name:      "filter with special chars",
			input:     map[string]int{"!@#": 1, "$%": 2, "normal": 3},
			predicate: func(k string, v int) bool { return v > 1 },
			expected:  map[string]int{"$%": 2, "normal": 3},
		},
		{
			name:      "large values",
			input:     map[string]int{"big": 1000000, "small": 1},
			predicate: func(k string, v int) bool { return v > 100000 },
			expected:  map[string]int{"big": 1000000},
		},
		{
			name:      "many entries",
			input:     map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			predicate: func(k string, v int) bool { return v >= 3 },
			expected:  map[string]int{"c": 3, "d": 4, "e": 5},
		},
		{
			name:      "key based filter",
			input:     map[string]int{"apple": 1, "apricot": 2, "banana": 3},
			predicate: func(k string, v int) bool { return k[0] == 'a' },
			expected:  map[string]int{"apple": 1, "apricot": 2},
		},
		{
			name:      "filter zero values",
			input:     map[string]int{"zero1": 0, "zero2": 0, "one": 1},
			predicate: func(k string, v int) bool { return v == 0 },
			expected:  map[string]int{"zero1": 0, "zero2": 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Filter(tt.input, tt.predicate)

			if len(got) != len(tt.expected) {
				t.Fatalf("Filter() returned %d entries, want %d", len(got), len(tt.expected))
			}

			for k, expectedV := range tt.expected {
				gotV, ok := got[k]
				if !ok {
					t.Fatalf("Filter() missing key %q in result", k)
				}
				if gotV != expectedV {
					t.Fatalf("Filter() for key %q: got %d, want %d", k, gotV, expectedV)
				}
			}

			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("Filter() has unexpected key %q in result", k)
				}
			}
		})
	}
}

func TestFilter_IntString(t *testing.T) {
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
			expected:  map[int]string{},
		},
		{
			name:      "empty map",
			input:     map[int]string{},
			predicate: func(k int, v string) bool { return true },
			expected:  map[int]string{},
		},
		{
			name:      "single element match",
			input:     map[int]string{1: "one"},
			predicate: func(k int, v string) bool { return k > 0 },
			expected:  map[int]string{1: "one"},
		},
		{
			name:      "single element no match",
			input:     map[int]string{1: "one"},
			predicate: func(k int, v string) bool { return k < 0 },
			expected:  map[int]string{},
		},
		{
			name:      "filter by key positive",
			input:     map[int]string{-2: "neg", -1: "neg", 1: "pos", 2: "pos"},
			predicate: func(k int, v string) bool { return k > 0 },
			expected:  map[int]string{1: "pos", 2: "pos"},
		},
		{
			name:      "filter by value length",
			input:     map[int]string{1: "a", 2: "bb", 3: "ccc", 4: "dddd"},
			predicate: func(k int, v string) bool { return len(v) > 2 },
			expected:  map[int]string{3: "ccc", 4: "dddd"},
		},
		{
			name:      "filter by key and value",
			input:     map[int]string{1: "one", 2: "two", 3: "three", 4: "four"},
			predicate: func(k int, v string) bool { return k > 1 && len(v) > 3 },
			expected:  map[int]string{3: "three", 4: "four"},
		},
		{
			name:      "match all",
			input:     map[int]string{1: "a", 2: "b", 3: "c"},
			predicate: func(k int, v string) bool { return true },
			expected:  map[int]string{1: "a", 2: "b", 3: "c"},
		},
		{
			name:      "match none",
			input:     map[int]string{1: "a", 2: "b", 3: "c"},
			predicate: func(k int, v string) bool { return false },
			expected:  map[int]string{},
		},
		{
			name:      "filter even keys",
			input:     map[int]string{1: "odd", 2: "even", 3: "odd", 4: "even"},
			predicate: func(k int, v string) bool { return k%2 == 0 },
			expected:  map[int]string{2: "even", 4: "even"},
		},
		{
			name:      "filter zero key",
			input:     map[int]string{0: "zero", 1: "one", -1: "neg"},
			predicate: func(k int, v string) bool { return k == 0 },
			expected:  map[int]string{0: "zero"},
		},
		{
			name:      "filter unicode values",
			input:     map[int]string{1: "hello", 2: "こんにちは", 3: "world"},
			predicate: func(k int, v string) bool { return len(v) > 5 },
			expected:  map[int]string{2: "こんにちは"},
		},
		{
			name:      "filter emoji values",
			input:     map[int]string{1: "😊", 2: "🚀", 3: "hello"},
			predicate: func(k int, v string) bool { return k < 3 },
			expected:  map[int]string{1: "😊", 2: "🚀"},
		},
		{
			name:      "large keys",
			input:     map[int]string{1000000: "big", 1: "small", 999999: "huge"},
			predicate: func(k int, v string) bool { return k > 10000 },
			expected:  map[int]string{1000000: "big", 999999: "huge"},
		},
		{
			name:      "many entries",
			input:     map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"},
			predicate: func(k int, v string) bool { return k <= 3 },
			expected:  map[int]string{1: "a", 2: "b", 3: "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Filter(tt.input, tt.predicate)

			if len(got) != len(tt.expected) {
				t.Fatalf("Filter() returned %d entries, want %d", len(got), len(tt.expected))
			}

			for k, expectedV := range tt.expected {
				gotV, ok := got[k]
				if !ok {
					t.Fatalf("Filter() missing key %d in result", k)
				}
				if gotV != expectedV {
					t.Fatalf("Filter() for key %d: got %q, want %q", k, gotV, expectedV)
				}
			}

			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("Filter() has unexpected key %d in result", k)
				}
			}
		})
	}
}

func TestFilter_StructValue(t *testing.T) {
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
			name: "filter by struct field",
			input: map[string]User{
				"alice":   {Name: "Alice", Age: 25},
				"bob":     {Name: "Bob", Age: 30},
				"charlie": {Name: "Charlie", Age: 20},
			},
			predicate: func(k string, v User) bool { return v.Age >= 25 },
			expected: map[string]User{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
			},
		},
		{
			name:      "filter empty struct map",
			input:     map[string]User{},
			predicate: func(k string, v User) bool { return v.Age > 0 },
			expected:  map[string]User{},
		},
		{
			name: "filter by key and struct field",
			input: map[string]User{
				"user1":  {Name: "User1", Age: 20},
				"admin1": {Name: "Admin1", Age: 30},
				"admin2": {Name: "Admin2", Age: 35},
			},
			predicate: func(k string, v User) bool { return len(k) >= 6 },
			expected: map[string]User{
				"admin1": {Name: "Admin1", Age: 30},
				"admin2": {Name: "Admin2", Age: 35},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Filter(tt.input, tt.predicate)

			if len(got) != len(tt.expected) {
				t.Fatalf("Filter() returned %d entries, want %d", len(got), len(tt.expected))
			}

			for k, expectedV := range tt.expected {
				gotV, ok := got[k]
				if !ok {
					t.Fatalf("Filter() missing key %q in result", k)
				}
				if gotV != expectedV {
					t.Fatalf("Filter() for key %q: got %+v, want %+v", k, gotV, expectedV)
				}
			}
		})
	}
}
