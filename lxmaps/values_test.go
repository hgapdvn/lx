package lxmaps_test

import (
	"testing"

	"github.com/hgapdvn/lx/lxmaps"
	"github.com/hgapdvn/lx/lxslices"
)

func TestValues_String(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[string]int
		expected []int
	}{
		{"no args", nil, nil},
		{"single nil map", []map[string]int{nil}, []int{}},
		{"empty map", []map[string]int{{}}, []int{}},
		{"single value", []map[string]int{{"a": 1}}, []int{1}},
		{"multiple values small", []map[string]int{{"a": 1, "b": 2, "c": 3}}, []int{1, 2, 3}},
		{"unicode values", []map[string]int{{"k1": 10}, {"k2": 20}}, []int{10, 20}},
		{"emoji values", []map[string]int{{"e": 100}, {"f": 200}}, []int{100, 200}},
		{"long values", []map[string]int{{"long": 123456}, {"other": 654321}}, []int{123456, 654321}},
		{"many values small", []map[string]int{{"k1": 1, "k2": 2, "k3": 3, "k4": 4, "k5": 5}}, []int{1, 2, 3, 4, 5}},
		{"mixed values", []map[string]int{{"x": 7}, {"y": 8}, {"z": 9}}, []int{7, 8, 9}},
		{"zero value", []map[string]int{{"zero": 0}, {"one": 1}}, []int{0, 1}},
		{"negative values", []map[string]int{{"n1": -1}, {"n2": -2}}, []int{-1, -2}},
		{"sparse values", []map[string]int{{"a": 2}, {"b": 100}}, []int{2, 100}},
		{"single large", []map[string]int{{"big": 999}}, []int{999}},
		{"sequential", []map[string]int{{"a": 10}, {"b": 11}, {"c": 12}}, []int{10, 11, 12}},
		{"mixed many", []map[string]int{{"p": 100}, {"q": -1}, {"r": 0}}, []int{100, -1, 0}},
		{"duplicate key across maps", []map[string]int{{"dup": 1}, {"dup": 2}}, []int{1, 2}},
		{"empty + map", []map[string]int{{}, {"x": 3}}, []int{3}},
		{"many maps combined", []map[string]int{{"m1": 1}, {"m2": 2}, {"m3": 3}, {"m4": 4}}, []int{1, 2, 3, 4}},
		{"special chars", []map[string]int{{"!@#": 7}, {"$%": 8}}, []int{7, 8}},
		{"final small", []map[string]int{{"end": 3}, {"stop": 4}}, []int{3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Values(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("Values(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("Values(%v) = nil; want non-nil with length %d", tt.input, len(tt.expected))
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("Values(%v) length = %d; want %d", tt.input, len(got), len(tt.expected))
			}

			for _, exp := range tt.expected {
				if !lxslices.Contains(got, exp) {
					t.Fatalf("Values(%v) missing value %v in result %v", tt.input, exp, got)
				}
			}
		})
	}
}

func TestValues_Int(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[int]string
		expected []string
	}{
		{"no args", nil, nil},
		{"single nil map", []map[int]string{nil}, []string{}},
		{"empty map", []map[int]string{{}}, []string{}},
		{"single value", []map[int]string{{1: "one"}}, []string{"one"}},
		{"multiple values small", []map[int]string{{1: "a", 2: "b", 3: "c"}}, []string{"a", "b", "c"}},
		{"zero value", []map[int]string{{0: "zero"}, {1: "one"}}, []string{"zero", "one"}},
		{"negative values", []map[int]string{{-1: "neg"}, {-2: "neg2"}}, []string{"neg", "neg2"}},
		{"mixed values", []map[int]string{{-5: "a"}, {5: "b"}}, []string{"a", "b"}},
		{"long strings", []map[int]string{{100: "long"}, {200: "other"}}, []string{"long", "other"}},
		{"many values", []map[int]string{{1: "k1", 2: "k2", 3: "k3", 4: "k4", 5: "k5"}}, []string{"k1", "k2", "k3", "k4", "k5"}},
		{"mixed small", []map[int]string{{7: "g"}, {8: "h"}, {9: "i"}}, []string{"g", "h", "i"}},
		{"sparse", []map[int]string{{2: "a"}, {100: "b"}}, []string{"a", "b"}},
		{"single large", []map[int]string{{999: "x"}}, []string{"x"}},
		{"duplicate across maps", []map[int]string{{2: "b"}, {2: "c"}}, []string{"b", "c"}},
		{"empty + map", []map[int]string{{}, {42: "x"}}, []string{"x"}},
		{"final small", []map[int]string{{3: "end"}, {4: "stop"}}, []string{"end", "stop"}},
		{"random small", []map[int]string{{42: "x"}, {7: "y"}}, []string{"x", "y"}},
		{"mixed many", []map[int]string{{100: "x"}, {-1: "y"}, {0: "z"}}, []string{"x", "y", "z"}},
		{"sequential", []map[int]string{{10: "a"}, {11: "b"}, {12: "c"}}, []string{"a", "b", "c"}},
		{"final extras", []map[int]string{{11: "k"}, {22: "l"}}, []string{"k", "l"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Values(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("Values(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("Values(%v) = nil; want non-nil with length %d", tt.input, len(tt.expected))
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("Values(%v) length = %d; want %d", tt.input, len(got), len(tt.expected))
			}

			for _, exp := range tt.expected {
				if !lxslices.Contains(got, exp) {
					t.Fatalf("Values(%v) missing value %q in result %v", tt.input, exp, got)
				}
			}
		})
	}
}

func TestValues_Struct(t *testing.T) {
	// comparable struct type for values
	type VS struct {
		I int
		S string
	}

	tests := []struct {
		name     string
		input    []map[string]VS
		expected []VS
	}{
		{"no args", nil, nil},
		{"single nil map", []map[string]VS{nil}, []VS{}},
		{"empty map", []map[string]VS{{}}, []VS{}},
		{"single value", []map[string]VS{{"a": {I: 1, S: "one"}}}, []VS{{I: 1, S: "one"}}},
		{"multiple values small", []map[string]VS{{"a": {I: 1, S: "a"}, "b": {I: 2, S: "b"}}}, []VS{{I: 1, S: "a"}, {I: 2, S: "b"}}},
		{"duplicate across maps", []map[string]VS{{"dup": {I: 1, S: "a"}}, {"dup": {I: 2, S: "b"}}}, []VS{{I: 1, S: "a"}, {I: 2, S: "b"}}},
		{"many maps combined", []map[string]VS{{"m1": {I: 1, S: "m1"}}, {"m2": {I: 2, S: "m2"}}, {"m3": {I: 3, S: "m3"}}}, []VS{{I: 1, S: "m1"}, {I: 2, S: "m2"}, {I: 3, S: "m3"}}},
		{"unicode values", []map[string]VS{{"u": {I: 10, S: "こんにちは"}}, {"v": {I: 11, S: "世界"}}}, []VS{{I: 10, S: "こんにちは"}, {I: 11, S: "世界"}}},
		{"emoji values", []map[string]VS{{"e": {I: 2, S: "😊"}}, {"f": {I: 3, S: "👋"}}}, []VS{{I: 2, S: "😊"}, {I: 3, S: "👋"}}},
		{"empty+nil mix", []map[string]VS{nil, {}, {"x": {I: 5, S: "X"}}}, []VS{{I: 5, S: "X"}}},
		{"long strings", []map[string]VS{{"long": {I: 12345, S: "longstring..."}}}, []VS{{I: 12345, S: "longstring..."}}},
		{"random small", []map[string]VS{{"x": {I: 42, S: "x"}}, {"y": {I: 7, S: "y"}}}, []VS{{I: 42, S: "x"}, {I: 7, S: "y"}}},
		{"many duplicates", []map[string]VS{{"d": {I: 1, S: "a"}}, {"d": {I: 2, S: "b"}}, {"d": {I: 3, S: "c"}}}, []VS{{I: 1, S: "a"}, {I: 2, S: "b"}, {I: 3, S: "c"}}},
		{"special chars", []map[string]VS{{"!@#": {I: 7, S: "sym"}}, {"$%": {I: 8, S: "sym2"}}}, []VS{{I: 7, S: "sym"}, {I: 8, S: "sym2"}}},
		{"final small", []map[string]VS{{"end": {I: 3, S: "end"}}, {"stop": {I: 4, S: "stop"}}}, []VS{{I: 3, S: "end"}, {I: 4, S: "stop"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Values(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("Values(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("Values(%v) = nil; want non-nil with length %d", tt.input, len(tt.expected))
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("Values(%v) length = %d; want %d", tt.input, len(got), len(tt.expected))
			}

			for _, exp := range tt.expected {
				if !lxslices.Contains(got, exp) {
					t.Fatalf("Values(%v) missing value %v in result %v", tt.input, exp, got)
				}
			}
		})
	}
}
