package lxmaps_test

import (
	"testing"

	"github.com/hgapdvn/lx/maps"
	"github.com/hgapdvn/lx/slices"
)

func TestUniqKeys_String(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[string]int
		expected []string // unique expected keys, order not assumed
	}{
		{"no args", nil, nil},
		{"single nil map", []map[string]int{nil}, []string{}},
		{"all nil variadic", []map[string]int{nil, nil}, []string{}},
		{"single map", []map[string]int{{"a": 1, "b": 2}}, []string{"a", "b"}},
		{"multi maps duplicated keys", []map[string]int{{"a": 1, "b": 2}, {"b": 3, "c": 4}}, []string{"a", "b", "c"}},
		{"many duplicates", []map[string]int{{"d": 1}, {"d": 2}, {"d": 3}}, []string{"d"}},
		{"mix empty and nil", []map[string]int{nil, {}, {"x": 9}}, []string{"x"}},
		{"unicode and emoji", []map[string]int{{"こんにちは": 1}, {"😊": 2, "こんにちは": 3}}, []string{"こんにちは", "😊"}},
		{"special chars", []map[string]int{{"!@#": 7}, {"$%": 8}}, []string{"!@#", "$%"}},
		{"many maps combined", []map[string]int{{"m1": 1}, {"m2": 2}, {"m3": 3}}, []string{"m1", "m2", "m3"}},
		{"long keys", []map[string]int{{"this_is_a_very_long_key_example_0001": 1}, {"this_is_a_very_long_key_example_0002": 2}}, []string{"this_is_a_very_long_key_example_0001", "this_is_a_very_long_key_example_0002"}},
		{"comma/semicolon mix", []map[string]int{{"a,b": 1}, {"c;d": 2}}, []string{"a,b", "c;d"}},
		{"numeric string keys", []map[string]int{{"1": 1}, {"2": 2}, {"10": 10}}, []string{"1", "2", "10"}},
		{"whitespace keys", []map[string]int{{" ": 1}, {"\t": 2}, {"\n": 3}}, []string{" ", "\t", "\n"}},
		{"case sensitivity", []map[string]int{{"Go": 1}, {"go": 2}}, []string{"Go", "go"}},
		{"prefix/suffix patterns", []map[string]int{{"pre_": 1}, {"_suf": 2}, {"pre_suf": 3}}, []string{"pre_", "_suf", "pre_suf"}},
		{"punctuation", []map[string]int{{"end.": 1}, {"start,": 2}, {"(paren)": 3}}, []string{"end.", "start,", "(paren)"}},
		{"small randomized mix", []map[string]int{{"a": 1, "b": 2}, {"c": 3}, {"b": 4, "d": 5}}, []string{"a", "b", "c", "d"}},
		{"final extras", []map[string]int{{"z1": 1}, {"z2": 2}, {"z3": 3}}, []string{"z1", "z2", "z3"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.UniqKeys(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("UniqKeys(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("UniqKeys(%v) = nil; want non-nil", tt.input)
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("UniqKeys(%v) length = %d; want %d; got=%v", tt.input, len(got), len(tt.expected), got)
			}

			if !lxslices.ContainsAll(got, tt.expected...) {
				t.Fatalf("UniqKeys(%v) missing expected keys; got %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestUniqKeys_Int(t *testing.T) {
	tests := []struct {
		name     string
		input    []map[int]string
		expected []int
	}{
		{"no args", nil, nil},
		{"single nil map", []map[int]string{nil}, []int{}},
		{"all nil variadic", []map[int]string{nil, nil}, []int{}},
		{"multi maps duplicated", []map[int]string{{1: "a", 2: "b"}, {2: "c", 3: "d"}}, []int{1, 2, 3}},
		{"many duplicates", []map[int]string{{5: "x"}, {5: "y"}, {5: "z"}}, []int{5}},
		{"mix empty and nil", []map[int]string{nil, {}, {9: "k"}}, []int{9}},
		{"large keys", []map[int]string{{1000: "a"}, {2000: "b"}, {3000: "c"}}, []int{1000, 2000, 3000}},
		{"negative and zero", []map[int]string{{0: "z"}, {-1: "n"}, {-2: "m"}}, []int{0, -1, -2}},
		{"sequential keys", []map[int]string{{1: "a", 2: "b", 3: "c"}}, []int{1, 2, 3}},
		{"sparse and dense mix", []map[int]string{{1: "a"}, {100: "b"}, {2: "c"}}, []int{1, 100, 2}},
		{"combined many", []map[int]string{{10: "x"}, {11: "y"}, {12: "z"}, {13: "t"}}, []int{10, 11, 12, 13}},
		{"duplicates across many", []map[int]string{{7: "a"}, {7: "b"}, {7: "c"}, {8: "d"}}, []int{7, 8}},
		{"final extras", []map[int]string{{11: "k"}, {22: "l"}}, []int{11, 22}},
		{"randomized small", []map[int]string{{42: "x", 7: "y"}, {7: "z"}}, []int{42, 7}},
		{"edge small", []map[int]string{{-2147483648: "min"}, {2147483647: "max"}}, []int{-2147483648, 2147483647}},
		{"mixed signed", []map[int]string{{-3: "a"}, {3: "b"}}, []int{-3, 3}},
		{"step keys", []map[int]string{{2: "a"}, {4: "b"}, {6: "c"}}, []int{2, 4, 6}},
		{"modulus keys", []map[int]string{{2: "r2"}, {4: "r4"}, {6: "r6"}}, []int{2, 4, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.UniqKeys(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("UniqKeys(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("UniqKeys(%v) = nil; want non-nil", tt.input)
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("UniqKeys(%v) length = %d; want %d; got=%v", tt.input, len(got), len(tt.expected), got)
			}

			if !lxslices.ContainsAll(got, tt.expected...) {
				t.Fatalf("UniqKeys(%v) missing expected keys; got %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestUniqKeys_Struct(t *testing.T) {
	// comparable struct type for keys
	type KS struct {
		ID   int
		Name string
	}

	tests := []struct {
		name     string
		input    []map[KS]int
		expected []KS
	}{
		{"no args", nil, nil},
		{"single nil map", []map[KS]int{nil}, []KS{}},
		{"all nil variadic", []map[KS]int{nil, nil}, []KS{}},
		{"single map", []map[KS]int{{{ID: 1, Name: "a"}: 1, {ID: 2, Name: "b"}: 2}}, []KS{{ID: 1, Name: "a"}, {ID: 2, Name: "b"}}},
		{"multi maps duplicated", []map[KS]int{{{ID: 1, Name: "a"}: 1}, {{ID: 2, Name: "b"}: 2, {ID: 1, Name: "a"}: 3}}, []KS{{ID: 1, Name: "a"}, {ID: 2, Name: "b"}}},
		{"many duplicates", []map[KS]int{{{ID: 5, Name: "x"}: 1}, {{ID: 5, Name: "x"}: 2}, {{ID: 5, Name: "x"}: 3}}, []KS{{ID: 5, Name: "x"}}},
		{"mix empty and nil", []map[KS]int{nil, {}, {{ID: 9, Name: "k"}: 9}}, []KS{{ID: 9, Name: "k"}}},
		{"unicode in name", []map[KS]int{{{ID: 1, Name: "こんにちは"}: 1}, {{ID: 2, Name: "世界"}: 2}}, []KS{{ID: 1, Name: "こんにちは"}, {ID: 2, Name: "世界"}}},
		{"empty name field", []map[KS]int{{{ID: 0, Name: ""}: 1}, {{ID: 1, Name: " "}: 2}}, []KS{{ID: 0, Name: ""}, {ID: 1, Name: " "}}},
		{"large ids", []map[KS]int{{{ID: 1000000, Name: "a"}: 1}, {{ID: 999999, Name: "b"}: 2}}, []KS{{ID: 1000000, Name: "a"}, {ID: 999999, Name: "b"}}},
		{"negative ids", []map[KS]int{{{ID: -1, Name: "n"}: 1}, {{ID: -2, Name: "m"}: 2}}, []KS{{ID: -1, Name: "n"}, {ID: -2, Name: "m"}}},
		{"overlapping maps", []map[KS]int{{{ID: 1, Name: "x"}: 1, {ID: 2, Name: "y"}: 2}, {{ID: 2, Name: "y"}: 3, {ID: 3, Name: "z"}: 4}}, []KS{{ID: 1, Name: "x"}, {ID: 2, Name: "y"}, {ID: 3, Name: "z"}}},
		{"many maps combined", []map[KS]int{{{ID: 1, Name: "m1"}: 1}, {{ID: 2, Name: "m2"}: 2}, {{ID: 3, Name: "m3"}: 3}}, []KS{{ID: 1, Name: "m1"}, {ID: 2, Name: "m2"}, {ID: 3, Name: "m3"}}},
		{"final extras", []map[KS]int{{{ID: 11, Name: "k"}: 1}, {{ID: 22, Name: "l"}: 2}}, []KS{{ID: 11, Name: "k"}, {ID: 22, Name: "l"}}},
		{"random small", []map[KS]int{{{ID: 42, Name: "x"}: 1}, {{ID: 7, Name: "y"}: 2}}, []KS{{ID: 42, Name: "x"}, {ID: 7, Name: "y"}}},
		{"struct with same name different id", []map[KS]int{{{ID: 1, Name: "dup"}: 1}, {{ID: 2, Name: "dup"}: 2}}, []KS{{ID: 1, Name: "dup"}, {ID: 2, Name: "dup"}}},
		{"prefix suffix mix", []map[KS]int{{{ID: 1, Name: "pre_"}: 1}, {{ID: 2, Name: "_suf"}: 2}, {{ID: 3, Name: "pre_suf"}: 3}}, []KS{{ID: 1, Name: "pre_"}, {ID: 2, Name: "_suf"}, {ID: 3, Name: "pre_suf"}}},
		{"special chars in name", []map[KS]int{{{ID: 7, Name: "!@#"}: 1}, {{ID: 8, Name: "$%"}: 2}}, []KS{{ID: 7, Name: "!@#"}, {ID: 8, Name: "$%"}}},
		{"sparse dense mix", []map[KS]int{{{ID: 1, Name: "a"}: 1}, {{ID: 100, Name: "b"}: 2}, {{ID: 2, Name: "c"}: 3}}, []KS{{ID: 1, Name: "a"}, {ID: 100, Name: "b"}, {ID: 2, Name: "c"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.UniqKeys(tt.input...)

			if tt.expected == nil {
				if got != nil {
					t.Fatalf("UniqKeys(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("UniqKeys(%v) = nil; want non-nil", tt.input)
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("UniqKeys(%v) length = %d; want %d; got=%v", tt.input, len(got), len(tt.expected), got)
			}

			if !lxslices.ContainsAll(got, tt.expected...) {
				t.Fatalf("UniqKeys(%v) missing expected keys; got %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
