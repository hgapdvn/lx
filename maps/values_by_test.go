package lxmaps_test

import (
	"sort"
	"testing"

	"github.com/hgapdvn/lx/maps"
	"github.com/hgapdvn/lx/slices"
)

func TestValuesBy_StringInt(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]int
		predicate func(v int) bool
		expected  []int
	}{
		{
			name:      "nil map",
			input:     nil,
			predicate: func(v int) bool { return true },
			expected:  []int{},
		},
		{
			name:      "empty map",
			input:     map[string]int{},
			predicate: func(v int) bool { return true },
			expected:  []int{},
		},
		{
			name:      "single element matches",
			input:     map[string]int{"a": 1},
			predicate: func(v int) bool { return v > 0 },
			expected:  []int{1},
		},
		{
			name:      "single element doesn't match",
			input:     map[string]int{"a": 1},
			predicate: func(v int) bool { return v > 10 },
			expected:  []int{},
		},
		{
			name:      "filter even numbers",
			input:     map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  []int{2, 4},
		},
		{
			name:      "filter positive values",
			input:     map[string]int{"neg": -5, "zero": 0, "pos": 5},
			predicate: func(v int) bool { return v > 0 },
			expected:  []int{5},
		},
		{
			name:      "filter by range",
			input:     map[string]int{"a": 1, "b": 5, "c": 10, "d": 15},
			predicate: func(v int) bool { return v > 3 && v < 12 },
			expected:  []int{5, 10},
		},
		{
			name:      "match nothing",
			input:     map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(v int) bool { return false },
			expected:  []int{},
		},
		{
			name:      "match everything",
			input:     map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(v int) bool { return true },
			expected:  []int{1, 2, 3},
		},
		{
			name:      "filter negative values",
			input:     map[string]int{"neg1": -5, "neg2": -2, "pos": 5},
			predicate: func(v int) bool { return v < 0 },
			expected:  []int{-5, -2},
		},
		{
			name:      "filter zero values",
			input:     map[string]int{"zero1": 0, "one": 1, "zero2": 0},
			predicate: func(v int) bool { return v == 0 },
			expected:  []int{0, 0},
		},
		{
			name:      "large values",
			input:     map[string]int{"big": 1000000, "small": 1},
			predicate: func(v int) bool { return v > 100000 },
			expected:  []int{1000000},
		},
		{
			name:      "many entries",
			input:     map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			predicate: func(v int) bool { return v >= 3 },
			expected:  []int{3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ValuesBy(tt.input, tt.predicate)

			if got == nil && tt.expected != nil {
				t.Fatalf("ValuesBy() returned nil, want non-nil slice")
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("ValuesBy() returned %d values, want %d", len(got), len(tt.expected))
			}

			// Sort both slices for comparison since map iteration order is random
			sort.Ints(got)
			sort.Ints(tt.expected)

			if !lxslices.Equal(got, tt.expected) {
				t.Fatalf("ValuesBy() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestValuesBy_IntString(t *testing.T) {
	tests := []struct {
		name      string
		input     map[int]string
		predicate func(v string) bool
		expected  []string
	}{
		{
			name:      "nil map",
			input:     nil,
			predicate: func(v string) bool { return true },
			expected:  []string{},
		},
		{
			name:      "empty map",
			input:     map[int]string{},
			predicate: func(v string) bool { return true },
			expected:  []string{},
		},
		{
			name:      "filter by string length",
			input:     map[int]string{1: "a", 2: "bb", 3: "ccc"},
			predicate: func(v string) bool { return len(v) > 1 },
			expected:  []string{"bb", "ccc"},
		},
		{
			name:      "filter by string content",
			input:     map[int]string{1: "apple", 2: "banana", 3: "apricot"},
			predicate: func(v string) bool { return v[0] == 'a' },
			expected:  []string{"apple", "apricot"},
		},
		{
			name:      "match nothing",
			input:     map[int]string{1: "a", 2: "b", 3: "c"},
			predicate: func(v string) bool { return false },
			expected:  []string{},
		},
		{
			name:      "match everything",
			input:     map[int]string{1: "a", 2: "b", 3: "c"},
			predicate: func(v string) bool { return true },
			expected:  []string{"a", "b", "c"},
		},
		{
			name:      "filter empty strings",
			input:     map[int]string{1: "", 2: "text", 3: ""},
			predicate: func(v string) bool { return v == "" },
			expected:  []string{"", ""},
		},
		{
			name:      "filter non-empty strings",
			input:     map[int]string{1: "", 2: "text", 3: ""},
			predicate: func(v string) bool { return v != "" },
			expected:  []string{"text"},
		},
		{
			name:      "unicode strings",
			input:     map[int]string{1: "こんにちは", 2: "世界", 3: "test"},
			predicate: func(v string) bool { return len(v) > 6 },
			expected:  []string{"こんにちは"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ValuesBy(tt.input, tt.predicate)

			if got == nil && tt.expected != nil {
				t.Fatalf("ValuesBy() returned nil, want non-nil slice")
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("ValuesBy() returned %d values, want %d", len(got), len(tt.expected))
			}

			// Sort both slices for comparison
			sort.Strings(got)
			sort.Strings(tt.expected)

			if !lxslices.Equal(got, tt.expected) {
				t.Fatalf("ValuesBy() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestValuesBy_BoolFloat(t *testing.T) {
	tests := []struct {
		name      string
		input     map[bool]float64
		predicate func(v float64) bool
		expected  []float64
	}{
		{
			name:      "nil map",
			input:     nil,
			predicate: func(v float64) bool { return true },
			expected:  []float64{},
		},
		{
			name:      "empty map",
			input:     map[bool]float64{},
			predicate: func(v float64) bool { return true },
			expected:  []float64{},
		},
		{
			name:      "filter positive values",
			input:     map[bool]float64{true: 3.14, false: -1.5},
			predicate: func(v float64) bool { return v > 0 },
			expected:  []float64{3.14},
		},
		{
			name:      "filter negative values",
			input:     map[bool]float64{true: 3.14, false: -1.5},
			predicate: func(v float64) bool { return v < 0 },
			expected:  []float64{-1.5},
		},
		{
			name:      "match everything",
			input:     map[bool]float64{true: 1.0, false: 2.0},
			predicate: func(v float64) bool { return true },
			expected:  []float64{1.0, 2.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ValuesBy(tt.input, tt.predicate)

			if got == nil && tt.expected != nil {
				t.Fatalf("ValuesBy() returned nil, want non-nil slice")
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("ValuesBy() returned %d values, want %d", len(got), len(tt.expected))
			}

			// Sort both slices for comparison
			sort.Float64s(got)
			sort.Float64s(tt.expected)

			if !lxslices.Equal(got, tt.expected) {
				t.Fatalf("ValuesBy() = %v, want %v", got, tt.expected)
			}
		})
	}
}

type Product struct {
	Name  string
	Price float64
	Stock int
}

func TestValuesBy_StringStruct(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]Product
		predicate func(v Product) bool
		expected  []Product
	}{
		{
			name:      "nil map",
			input:     nil,
			predicate: func(v Product) bool { return true },
			expected:  []Product{},
		},
		{
			name:      "empty map",
			input:     map[string]Product{},
			predicate: func(v Product) bool { return true },
			expected:  []Product{},
		},
		{
			name: "filter by price",
			input: map[string]Product{
				"apple":  {Name: "Apple", Price: 1.5, Stock: 10},
				"banana": {Name: "Banana", Price: 0.5, Stock: 20},
				"orange": {Name: "Orange", Price: 2.0, Stock: 15},
			},
			predicate: func(v Product) bool { return v.Price > 1.0 },
			expected: []Product{
				{Name: "Apple", Price: 1.5, Stock: 10},
				{Name: "Orange", Price: 2.0, Stock: 15},
			},
		},
		{
			name: "filter by stock",
			input: map[string]Product{
				"apple":  {Name: "Apple", Price: 1.5, Stock: 10},
				"banana": {Name: "Banana", Price: 0.5, Stock: 20},
				"orange": {Name: "Orange", Price: 2.0, Stock: 5},
			},
			predicate: func(v Product) bool { return v.Stock >= 10 },
			expected: []Product{
				{Name: "Apple", Price: 1.5, Stock: 10},
				{Name: "Banana", Price: 0.5, Stock: 20},
			},
		},
		{
			name: "filter by name",
			input: map[string]Product{
				"apple":  {Name: "Apple", Price: 1.5, Stock: 10},
				"banana": {Name: "Banana", Price: 0.5, Stock: 20},
				"orange": {Name: "Orange", Price: 2.0, Stock: 15},
			},
			predicate: func(v Product) bool { return v.Name[0] == 'A' || v.Name[0] == 'B' },
			expected: []Product{
				{Name: "Apple", Price: 1.5, Stock: 10},
				{Name: "Banana", Price: 0.5, Stock: 20},
			},
		},
		{
			name: "filter by combined conditions",
			input: map[string]Product{
				"apple":  {Name: "Apple", Price: 1.5, Stock: 10},
				"banana": {Name: "Banana", Price: 0.5, Stock: 20},
				"orange": {Name: "Orange", Price: 2.0, Stock: 5},
			},
			predicate: func(v Product) bool { return v.Price > 1.0 && v.Stock > 8 },
			expected: []Product{
				{Name: "Apple", Price: 1.5, Stock: 10},
			},
		},
		{
			name: "match nothing",
			input: map[string]Product{
				"apple":  {Name: "Apple", Price: 1.5, Stock: 10},
				"banana": {Name: "Banana", Price: 0.5, Stock: 20},
			},
			predicate: func(v Product) bool { return v.Price > 10.0 },
			expected:  []Product{},
		},
		{
			name: "match everything",
			input: map[string]Product{
				"apple":  {Name: "Apple", Price: 1.5, Stock: 10},
				"banana": {Name: "Banana", Price: 0.5, Stock: 20},
			},
			predicate: func(v Product) bool { return true },
			expected: []Product{
				{Name: "Apple", Price: 1.5, Stock: 10},
				{Name: "Banana", Price: 0.5, Stock: 20},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ValuesBy(tt.input, tt.predicate)

			if got == nil && tt.expected != nil {
				t.Fatalf("ValuesBy() returned nil, want non-nil slice")
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("ValuesBy() returned %d values, want %d", len(got), len(tt.expected))
			}

			// Sort both slices by name for deterministic comparison
			sort.Slice(got, func(i, j int) bool { return got[i].Name < got[j].Name })
			sort.Slice(tt.expected, func(i, j int) bool { return tt.expected[i].Name < tt.expected[j].Name })

			if !lxslices.EqualFunc(got, tt.expected, func(a, b Product) bool {
				return a == b
			}) {
				t.Fatalf("ValuesBy() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestValuesBy_IntStruct(t *testing.T) {
	tests := []struct {
		name      string
		input     map[int]Product
		predicate func(v Product) bool
		expected  []Product
	}{
		{
			name:      "nil map",
			input:     nil,
			predicate: func(v Product) bool { return true },
			expected:  []Product{},
		},
		{
			name:      "empty map",
			input:     map[int]Product{},
			predicate: func(v Product) bool { return true },
			expected:  []Product{},
		},
		{
			name: "filter by price",
			input: map[int]Product{
				1: {Name: "Apple", Price: 1.5, Stock: 10},
				2: {Name: "Banana", Price: 0.5, Stock: 20},
				3: {Name: "Orange", Price: 2.0, Stock: 15},
			},
			predicate: func(v Product) bool { return v.Price > 1.0 },
			expected: []Product{
				{Name: "Apple", Price: 1.5, Stock: 10},
				{Name: "Orange", Price: 2.0, Stock: 15},
			},
		},
		{
			name: "filter by stock level",
			input: map[int]Product{
				1: {Name: "Apple", Price: 1.5, Stock: 5},
				2: {Name: "Banana", Price: 0.5, Stock: 20},
				3: {Name: "Orange", Price: 2.0, Stock: 15},
			},
			predicate: func(v Product) bool { return v.Stock > 10 },
			expected: []Product{
				{Name: "Banana", Price: 0.5, Stock: 20},
				{Name: "Orange", Price: 2.0, Stock: 15},
			},
		},
		{
			name: "single matching struct",
			input: map[int]Product{
				1: {Name: "Apple", Price: 1.5, Stock: 10},
				2: {Name: "Banana", Price: 0.5, Stock: 20},
			},
			predicate: func(v Product) bool { return v.Name == "Apple" },
			expected: []Product{
				{Name: "Apple", Price: 1.5, Stock: 10},
			},
		},
		{
			name: "match nothing",
			input: map[int]Product{
				1: {Name: "Apple", Price: 1.5, Stock: 10},
				2: {Name: "Banana", Price: 0.5, Stock: 20},
			},
			predicate: func(v Product) bool { return false },
			expected:  []Product{},
		},
		{
			name: "match everything",
			input: map[int]Product{
				1: {Name: "Apple", Price: 1.5, Stock: 10},
				2: {Name: "Banana", Price: 0.5, Stock: 20},
			},
			predicate: func(v Product) bool { return true },
			expected: []Product{
				{Name: "Apple", Price: 1.5, Stock: 10},
				{Name: "Banana", Price: 0.5, Stock: 20},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ValuesBy(tt.input, tt.predicate)

			if got == nil && tt.expected != nil {
				t.Fatalf("ValuesBy() returned nil, want non-nil slice")
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("ValuesBy() returned %d values, want %d", len(got), len(tt.expected))
			}

			// Sort both slices by name for deterministic comparison
			sort.Slice(got, func(i, j int) bool { return got[i].Name < got[j].Name })
			sort.Slice(tt.expected, func(i, j int) bool { return tt.expected[i].Name < tt.expected[j].Name })

			if !lxslices.EqualFunc(got, tt.expected, func(a, b Product) bool {
				return a == b
			}) {
				t.Fatalf("ValuesBy() = %v, want %v", got, tt.expected)
			}
		})
	}
}
