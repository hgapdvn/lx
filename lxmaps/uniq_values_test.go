package lxmaps_test

import (
	"testing"

	"github.com/hgapdvn/lx/lxmaps"
	"github.com/hgapdvn/lx/lxslices"
)

func TestUniqValues_Int(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[string]int
		expected []int
	}{
		{"no args", nil, nil},
		{"single nil map", []map[string]int{nil}, []int{}},
		{"all nil variadic", []map[string]int{nil, nil}, []int{}},
		{"single map", []map[string]int{{"a": 1, "b": 2}}, []int{1, 2}},
		{"multi maps duplicated values", []map[string]int{{"a": 1, "b": 2}, {"c": 2, "d": 3}}, []int{1, 2, 3}},
		{"many duplicates", []map[string]int{{"p": 5}, {"q": 5}, {"r": 5}}, []int{5}},
		{"mix empty and nil", []map[string]int{nil, {}, {"k": 9}}, []int{9}},
		{"negative and zero", []map[string]int{{"z": 0}, {"n": -1}, {"m": -2}}, []int{0, -1, -2}},
		{"large ints", []map[string]int{{"a": 1000000}, {"b": 999999}}, []int{1000000, 999999}},
		{"final extras", []map[string]int{{"x": 11}, {"y": 22}}, []int{11, 22}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.UniqValues(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("UniqValues(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("UniqValues(%v) = nil; want non-nil", tt.input)
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("UniqValues(%v) length = %d; want %d; got=%v", tt.input, len(got), len(tt.expected), got)
			}

			if !lxslices.ContainsAll(got, tt.expected...) {
				t.Fatalf("UniqValues(%v) missing expected values; got %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestUniqValues_Struct(t *testing.T) {
	type VS struct {
		N string
		M int
	}

	tests := []struct {
		name     string
		input    []map[string]VS
		expected []VS
	}{
		{"no args", nil, nil},
		{"single nil map", []map[string]VS{nil}, []VS{}},
		{"all nil variadic", []map[string]VS{nil, nil}, []VS{}},
		{"single map", []map[string]VS{{"a": {N: "one", M: 1}, "b": {N: "two", M: 2}}}, []VS{{N: "one", M: 1}, {N: "two", M: 2}}},
		{"multi maps duplicated values", []map[string]VS{{"a": {N: "x", M: 1}, "b": {N: "y", M: 2}}, {"c": {N: "y", M: 2}, "d": {N: "z", M: 3}}}, []VS{{N: "x", M: 1}, {N: "y", M: 2}, {N: "z", M: 3}}},
		{"many duplicates", []map[string]VS{{"p": {N: "dup", M: 1}}, {"q": {N: "dup", M: 1}}, {"r": {N: "dup", M: 1}}}, []VS{{N: "dup", M: 1}}},
		{"mix empty and nil", []map[string]VS{nil, {}, {"k": {N: "v", M: 9}}}, []VS{{N: "v", M: 9}}},
		{"unicode in struct", []map[string]VS{{"a": {N: "こんにちは", M: 1}}, {"b": {N: "世界", M: 2}}}, []VS{{N: "こんにちは", M: 1}, {N: "世界", M: 2}}},
		{"final extras", []map[string]VS{{"x": {N: "a", M: 1}}, {"y": {N: "b", M: 2}}}, []VS{{N: "a", M: 1}, {N: "b", M: 2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.UniqValues(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("UniqValues(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("UniqValues(%v) = nil; want non-nil", tt.input)
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("UniqValues(%v) length = %d; want %d; got=%v", tt.input, len(got), len(tt.expected), got)
			}

			for _, exp := range tt.expected {
				if !lxslices.Contains(got, exp) {
					t.Fatalf("UniqValues(%v) missing expected value %v; got=%v", tt.input, exp, got)
				}
			}
		})
	}
}

func TestUniqValues_String(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[string]string
		expected []string
	}{
		{"no args", nil, nil},
		{"single nil map", []map[string]string{nil}, []string{}},
		{"all nil variadic", []map[string]string{nil, nil}, []string{}},
		{"single map", []map[string]string{{"a": "one", "b": "two"}}, []string{"one", "two"}},
		{"multi maps duplicated values", []map[string]string{{"a": "one", "b": "two"}, {"c": "two", "d": "three"}}, []string{"one", "two", "three"}},
		{"many duplicates", []map[string]string{{"p": "dup"}, {"q": "dup"}, {"r": "dup"}}, []string{"dup"}},
		{"mix empty and nil", []map[string]string{nil, {}, {"k": "v"}}, []string{"v"}},
		{"unicode", []map[string]string{{"a": "こんにちは"}, {"b": "世界"}}, []string{"こんにちは", "世界"}},
		{"empty string value", []map[string]string{{"a": ""}, {"b": " "}}, []string{"", " "}},
		{"case sensitive", []map[string]string{{"a": "Go"}, {"b": "go"}}, []string{"Go", "go"}},
		{"special chars", []map[string]string{{"x": "!@#"}, {"y": "$%^"}}, []string{"!@#", "$%^"}},
		{"overlapping maps", []map[string]string{{"a": "x", "b": "y"}, {"b": "y", "c": "z"}, {}}, []string{"x", "y", "z"}},
		{"final extras", []map[string]string{{"x": "alpha"}, {"y": "beta"}}, []string{"alpha", "beta"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.UniqValues(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("UniqValues(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("UniqValues(%v) = nil; want non-nil", tt.input)
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("UniqValues(%v) length = %d; want %d; got=%v", tt.input, len(got), len(tt.expected), got)
			}

			if !lxslices.ContainsAll(got, tt.expected...) {
				t.Fatalf("UniqValues(%v) missing expected values; got %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
