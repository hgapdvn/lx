package lxmaps_test

import (
	"testing"

	"github.com/hgapdvn/lx/maps"
)

func TestOmitBy_StringInt(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]int
		predicate func(k string, v int) bool
		wantNil   bool
		expected  map[string]int
	}{
		{
			name:      "nil map",
			input:     nil,
			predicate: func(k string, v int) bool { return true },
			wantNil:   true,
			expected:  nil,
		},
		{
			name:      "empty map",
			input:     map[string]int{},
			predicate: func(k string, v int) bool { return true },
			wantNil:   false,
			expected:  map[string]int{},
		},
		{
			name:      "single element matches predicate",
			input:     map[string]int{"a": 1},
			predicate: func(k string, v int) bool { return v > 0 },
			wantNil:   false,
			expected:  map[string]int{},
		},
		{
			name:      "single element doesn't match predicate",
			input:     map[string]int{"a": 1},
			predicate: func(k string, v int) bool { return v > 10 },
			wantNil:   false,
			expected:  map[string]int{"a": 1},
		},
		{
			name:      "omit even numbers",
			input:     map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			predicate: func(k string, v int) bool { return v%2 == 0 },
			wantNil:   false,
			expected:  map[string]int{"a": 1, "c": 3},
		},
		{
			name:      "omit by key prefix",
			input:     map[string]int{"prefix_a": 1, "prefix_b": 2, "other": 3},
			predicate: func(k string, v int) bool { return len(k) > 5 },
			wantNil:   false,
			expected:  map[string]int{"other": 3},
		},
		{
			name:      "omit by both key and value",
			input:     map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			predicate: func(k string, v int) bool { return len(k) == 1 && v > 1 },
			wantNil:   false,
			expected:  map[string]int{"a": 1},
		},
		{
			name:      "omit nothing predicate false",
			input:     map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(k string, v int) bool { return false },
			wantNil:   false,
			expected:  map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name:      "omit everything predicate true",
			input:     map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(k string, v int) bool { return true },
			wantNil:   false,
			expected:  map[string]int{},
		},
		{
			name:      "omit negative values",
			input:     map[string]int{"neg1": -5, "zero": 0, "pos1": 5},
			predicate: func(k string, v int) bool { return v < 0 },
			wantNil:   false,
			expected:  map[string]int{"zero": 0, "pos1": 5},
		},
		{
			name:      "omit with unicode keys",
			input:     map[string]int{"こんにちは": 1, "世界": 2, "test": 3},
			predicate: func(k string, v int) bool { return v < 3 },
			wantNil:   false,
			expected:  map[string]int{"test": 3},
		},
		{
			name:      "omit with special chars in keys",
			input:     map[string]int{"!@#": 1, "$%": 2, "normal": 3},
			predicate: func(k string, v int) bool { return v > 1 },
			wantNil:   false,
			expected:  map[string]int{"!@#": 1},
		},
		{
			name:      "omit large values",
			input:     map[string]int{"big": 1000000, "small": 1},
			predicate: func(k string, v int) bool { return v > 100000 },
			wantNil:   false,
			expected:  map[string]int{"small": 1},
		},
		{
			name:      "omit from many entries",
			input:     map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			predicate: func(k string, v int) bool { return v >= 3 },
			wantNil:   false,
			expected:  map[string]int{"a": 1, "b": 2},
		},
		{
			name:      "omit by key condition",
			input:     map[string]int{"apple": 1, "apricot": 2, "banana": 3},
			predicate: func(k string, v int) bool { return k[0] == 'a' },
			wantNil:   false,
			expected:  map[string]int{"banana": 3},
		},
		{
			name:      "omit zero values",
			input:     map[string]int{"zero1": 0, "zero2": 0, "one": 1},
			predicate: func(k string, v int) bool { return v == 0 },
			wantNil:   false,
			expected:  map[string]int{"one": 1},
		},
		{
			name:      "omit preserves zero values not matching predicate",
			input:     map[string]int{"zero": 0, "one": 1, "two": 2},
			predicate: func(k string, v int) bool { return v > 1 },
			wantNil:   false,
			expected:  map[string]int{"zero": 0, "one": 1},
		},
		{
			name:      "omit empty string key",
			input:     map[string]int{"": 42, "a": 1},
			predicate: func(k string, v int) bool { return k == "" },
			wantNil:   false,
			expected:  map[string]int{"a": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.OmitBy(tt.input, tt.predicate)

			if tt.wantNil && got != nil {
				t.Fatalf("OmitBy() returned non-nil map, want nil")
			}
			if !tt.wantNil && got == nil {
				t.Fatalf("OmitBy() returned nil, want non-nil map")
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("OmitBy() returned %d entries, want %d", len(got), len(tt.expected))
			}

			for k, expectedV := range tt.expected {
				gotV, ok := got[k]
				if !ok {
					t.Fatalf("OmitBy() missing key %q in result", k)
				}
				if gotV != expectedV {
					t.Fatalf("OmitBy() for key %q: got %d, want %d", k, gotV, expectedV)
				}
			}

			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("OmitBy() has unexpected key %q in result", k)
				}
			}
		})
	}
}

func TestOmitBy_IntString(t *testing.T) {
	tests := []struct {
		name      string
		input     map[int]string
		predicate func(k int, v string) bool
		wantNil   bool
		expected  map[int]string
	}{
		{
			name:      "nil map",
			input:     nil,
			predicate: func(k int, v string) bool { return true },
			wantNil:   true,
			expected:  nil,
		},
		{
			name:      "empty map",
			input:     map[int]string{},
			predicate: func(k int, v string) bool { return true },
			wantNil:   false,
			expected:  map[int]string{},
		},
		{
			name:      "omit by string length",
			input:     map[int]string{1: "a", 2: "bb", 3: "ccc"},
			predicate: func(k int, v string) bool { return len(v) > 1 },
			wantNil:   false,
			expected:  map[int]string{1: "a"},
		},
		{
			name:      "omit by key value",
			input:     map[int]string{1: "one", 2: "two", 3: "three"},
			predicate: func(k int, v string) bool { return k > 1 },
			wantNil:   false,
			expected:  map[int]string{1: "one"},
		},
		{
			name:      "omit by string content",
			input:     map[int]string{1: "apple", 2: "banana", 3: "apricot"},
			predicate: func(k int, v string) bool { return v[0] == 'a' },
			wantNil:   false,
			expected:  map[int]string{2: "banana"},
		},
		{
			name:      "omit nothing",
			input:     map[int]string{1: "a", 2: "b", 3: "c"},
			predicate: func(k int, v string) bool { return false },
			wantNil:   false,
			expected:  map[int]string{1: "a", 2: "b", 3: "c"},
		},
		{
			name:      "omit everything",
			input:     map[int]string{1: "a", 2: "b", 3: "c"},
			predicate: func(k int, v string) bool { return true },
			wantNil:   false,
			expected:  map[int]string{},
		},
		{
			name:      "omit empty string values",
			input:     map[int]string{1: "", 2: "text", 3: ""},
			predicate: func(k int, v string) bool { return v == "" },
			wantNil:   false,
			expected:  map[int]string{2: "text"},
		},
		{
			name:      "omit by combined conditions",
			input:     map[int]string{1: "a", 2: "bb", 3: "ccc", 4: "dddd"},
			predicate: func(k int, v string) bool { return k > 2 && len(v) > 2 },
			wantNil:   false,
			expected:  map[int]string{1: "a", 2: "bb"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.OmitBy(tt.input, tt.predicate)

			if tt.wantNil && got != nil {
				t.Fatalf("OmitBy() returned non-nil map, want nil")
			}
			if !tt.wantNil && got == nil {
				t.Fatalf("OmitBy() returned nil, want non-nil map")
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("OmitBy() returned %d entries, want %d", len(got), len(tt.expected))
			}

			for k, expectedV := range tt.expected {
				gotV, ok := got[k]
				if !ok {
					t.Fatalf("OmitBy() missing key %q in result", k)
				}
				if gotV != expectedV {
					t.Fatalf("OmitBy() for key %q: got %q, want %q", k, gotV, expectedV)
				}
			}

			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("OmitBy() has unexpected key %q in result", k)
				}
			}
		})
	}
}

func TestOmitBy_BoolFloat(t *testing.T) {
	tests := []struct {
		name      string
		input     map[bool]float64
		predicate func(k bool, v float64) bool
		wantNil   bool
		expected  map[bool]float64
	}{
		{
			name:      "nil map",
			input:     nil,
			predicate: func(k bool, v float64) bool { return true },
			wantNil:   true,
			expected:  nil,
		},
		{
			name:      "omit positive values",
			input:     map[bool]float64{true: 3.14, false: -1.5},
			predicate: func(k bool, v float64) bool { return v > 0 },
			wantNil:   false,
			expected:  map[bool]float64{false: -1.5},
		},
		{
			name:      "omit by key",
			input:     map[bool]float64{true: 1.0, false: 2.0},
			predicate: func(k bool, v float64) bool { return k },
			wantNil:   false,
			expected:  map[bool]float64{false: 2.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.OmitBy(tt.input, tt.predicate)

			if tt.wantNil && got != nil {
				t.Fatalf("OmitBy() returned non-nil map, want nil")
			}
			if !tt.wantNil && got == nil {
				t.Fatalf("OmitBy() returned nil, want non-nil map")
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("OmitBy() returned %d entries, want %d", len(got), len(tt.expected))
			}

			for k, expectedV := range tt.expected {
				gotV, ok := got[k]
				if !ok {
					t.Fatalf("OmitBy() missing key %v in result", k)
				}
				if gotV != expectedV {
					t.Fatalf("OmitBy() for key %v: got %f, want %f", k, gotV, expectedV)
				}
			}

			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("OmitBy() has unexpected key %v in result", k)
				}
			}
		})
	}
}

type Person struct {
	Name string
	Age  int
}

func TestOmitBy_StringStruct(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]Person
		predicate func(k string, v Person) bool
		wantNil   bool
		expected  map[string]Person
	}{
		{
			name:      "nil map",
			input:     nil,
			predicate: func(k string, v Person) bool { return true },
			wantNil:   true,
			expected:  nil,
		},
		{
			name:      "empty map",
			input:     map[string]Person{},
			predicate: func(k string, v Person) bool { return true },
			wantNil:   false,
			expected:  map[string]Person{},
		},
		{
			name: "omit by struct field age",
			input: map[string]Person{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
				"carol": {Name: "Carol", Age: 35},
			},
			predicate: func(k string, v Person) bool { return v.Age > 28 },
			wantNil:   false,
			expected: map[string]Person{
				"alice": {Name: "Alice", Age: 25},
			},
		},
		{
			name: "omit by struct field name",
			input: map[string]Person{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
				"carol": {Name: "Carol", Age: 35},
			},
			predicate: func(k string, v Person) bool { return v.Name[0] == 'A' },
			wantNil:   false,
			expected: map[string]Person{
				"bob":   {Name: "Bob", Age: 30},
				"carol": {Name: "Carol", Age: 35},
			},
		},
		{
			name: "omit by key",
			input: map[string]Person{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
			},
			predicate: func(k string, v Person) bool { return k == "alice" },
			wantNil:   false,
			expected: map[string]Person{
				"bob": {Name: "Bob", Age: 30},
			},
		},
		{
			name: "omit by key and age",
			input: map[string]Person{
				"alice":   {Name: "Alice", Age: 25},
				"bob":     {Name: "Bob", Age: 30},
				"charlie": {Name: "Charlie", Age: 35},
			},
			predicate: func(k string, v Person) bool { return len(k) > 3 && v.Age < 35 },
			wantNil:   false,
			expected: map[string]Person{
				"bob":     {Name: "Bob", Age: 30},
				"charlie": {Name: "Charlie", Age: 35},
			},
		},
		{
			name: "omit nothing",
			input: map[string]Person{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
			},
			predicate: func(k string, v Person) bool { return false },
			wantNil:   false,
			expected: map[string]Person{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
			},
		},
		{
			name: "omit everything",
			input: map[string]Person{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
			},
			predicate: func(k string, v Person) bool { return true },
			wantNil:   false,
			expected:  map[string]Person{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.OmitBy(tt.input, tt.predicate)

			if tt.wantNil && got != nil {
				t.Fatalf("OmitBy() returned non-nil map, want nil")
			}
			if !tt.wantNil && got == nil {
				t.Fatalf("OmitBy() returned nil, want non-nil map")
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("OmitBy() returned %d entries, want %d", len(got), len(tt.expected))
			}

			for k, expectedV := range tt.expected {
				gotV, ok := got[k]
				if !ok {
					t.Fatalf("OmitBy() missing key %q in result", k)
				}
				if gotV != expectedV {
					t.Fatalf("OmitBy() for key %q: got %+v, want %+v", k, gotV, expectedV)
				}
			}

			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("OmitBy() has unexpected key %q in result", k)
				}
			}
		})
	}
}

func TestOmitBy_IntStruct(t *testing.T) {
	tests := []struct {
		name      string
		input     map[int]Person
		predicate func(k int, v Person) bool
		wantNil   bool
		expected  map[int]Person
	}{
		{
			name:      "nil map",
			input:     nil,
			predicate: func(k int, v Person) bool { return true },
			wantNil:   true,
			expected:  nil,
		},
		{
			name:      "empty map",
			input:     map[int]Person{},
			predicate: func(k int, v Person) bool { return true },
			wantNil:   false,
			expected:  map[int]Person{},
		},
		{
			name: "omit by struct age field",
			input: map[int]Person{
				1: {Name: "Alice", Age: 25},
				2: {Name: "Bob", Age: 30},
				3: {Name: "Carol", Age: 35},
			},
			predicate: func(k int, v Person) bool { return v.Age >= 30 },
			wantNil:   false,
			expected: map[int]Person{
				1: {Name: "Alice", Age: 25},
			},
		},
		{
			name: "omit by struct name field",
			input: map[int]Person{
				1: {Name: "Alice", Age: 25},
				2: {Name: "Bob", Age: 30},
				3: {Name: "Carol", Age: 35},
			},
			predicate: func(k int, v Person) bool { return len(v.Name) > 4 },
			wantNil:   false,
			expected: map[int]Person{
				2: {Name: "Bob", Age: 30},
			},
		},
		{
			name: "omit by integer key",
			input: map[int]Person{
				1: {Name: "Alice", Age: 25},
				2: {Name: "Bob", Age: 30},
				3: {Name: "Carol", Age: 35},
			},
			predicate: func(k int, v Person) bool { return k > 1 },
			wantNil:   false,
			expected: map[int]Person{
				1: {Name: "Alice", Age: 25},
			},
		},
		{
			name: "omit by combined key and struct conditions",
			input: map[int]Person{
				1: {Name: "Alice", Age: 25},
				2: {Name: "Bob", Age: 30},
				3: {Name: "Carol", Age: 35},
			},
			predicate: func(k int, v Person) bool { return k > 1 || v.Age < 28 },
			wantNil:   false,
			expected:  map[int]Person{},
		},
		{
			name: "omit nothing",
			input: map[int]Person{
				1: {Name: "Alice", Age: 25},
				2: {Name: "Bob", Age: 30},
			},
			predicate: func(k int, v Person) bool { return false },
			wantNil:   false,
			expected: map[int]Person{
				1: {Name: "Alice", Age: 25},
				2: {Name: "Bob", Age: 30},
			},
		},
		{
			name: "omit everything",
			input: map[int]Person{
				1: {Name: "Alice", Age: 25},
				2: {Name: "Bob", Age: 30},
			},
			predicate: func(k int, v Person) bool { return true },
			wantNil:   false,
			expected:  map[int]Person{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.OmitBy(tt.input, tt.predicate)

			if tt.wantNil && got != nil {
				t.Fatalf("OmitBy() returned non-nil map, want nil")
			}
			if !tt.wantNil && got == nil {
				t.Fatalf("OmitBy() returned nil, want non-nil map")
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("OmitBy() returned %d entries, want %d", len(got), len(tt.expected))
			}

			for k, expectedV := range tt.expected {
				gotV, ok := got[k]
				if !ok {
					t.Fatalf("OmitBy() missing key %d in result", k)
				}
				if gotV != expectedV {
					t.Fatalf("OmitBy() for key %d: got %+v, want %+v", k, gotV, expectedV)
				}
			}

			for k := range got {
				if _, ok := tt.expected[k]; !ok {
					t.Fatalf("OmitBy() has unexpected key %d in result", k)
				}
			}
		})
	}
}
