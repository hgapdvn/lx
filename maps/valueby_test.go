package lxmaps_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/nthanhhai2909/lx/lxslices"
	"github.com/nthanhhai2909/lx/maps"
)

func TestValueBy_String(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		fn       func(string, int) string
		expects  map[string]string
		expectSz int
	}{
		{"nil map", nil, func(k string, v int) string { return fmt.Sprintf("%s:%d", k, v) }, map[string]string{}, 0},
		{"empty map", map[string]int{}, func(k string, v int) string { return fmt.Sprintf("%s:%d", k, v) }, map[string]string{}, 0},
		{"single entry", map[string]int{"a": 1}, func(k string, v int) string { return fmt.Sprintf("%s:%d", k, v) }, map[string]string{"a": "a:1"}, 1},
		{"two entries", map[string]int{"a": 1, "b": 2}, func(k string, v int) string { return fmt.Sprintf("%s#%d", k, v) }, map[string]string{"a": "a#1", "b": "b#2"}, 2},
		{"three entries", map[string]int{"a": 1, "b": 2, "c": 3}, func(k string, v int) string { return fmt.Sprintf("X:%s|%d", k, v) }, map[string]string{"a": "X:a|1", "b": "X:b|2", "c": "X:c|3"}, 3},
		{"unicode key", map[string]int{"こんにちは": 1}, func(k string, v int) string { return k + ":u" }, map[string]string{"こんにちは": "こんにちは:u"}, 1},
		{"emoji key", map[string]int{"😊": 1}, func(k string, v int) string { return k + "_e" }, map[string]string{"😊": "😊_e"}, 1},
		{"special chars", map[string]int{"!@#": 7}, func(k string, v int) string { return fmt.Sprintf("val:%d", v) }, map[string]string{"!@#": "val:7"}, 1},
		{"whitespace keys", map[string]int{" ": 1, "\t": 2, "\n": 3}, func(k string, v int) string { return fmt.Sprintf("ws:%q", k) }, map[string]string{" ": "ws:\" \"", "\t": "ws:\"\\t\"", "\n": "ws:\"\\n\""}, 3},
		{"numeric string keys", map[string]int{"1": 1, "2": 2, "10": 10}, func(k string, v int) string { return "n:" + k }, map[string]string{"1": "n:1", "2": "n:2", "10": "n:10"}, 3},
		{"case sensitivity", map[string]int{"Go": 1, "go": 2}, func(k string, v int) string { return k }, map[string]string{"Go": "Go", "go": "go"}, 2},
		{"prefix/suffix", map[string]int{"pre_": 1, "_suf": 2, "pre_suf": 3}, func(k string, v int) string { return "p:" + k }, map[string]string{"pre_": "p:pre_", "_suf": "p:_suf", "pre_suf": "p:pre_suf"}, 3},
		{"punctuation", map[string]int{"end.": 1, "start,": 2}, func(k string, v int) string { return k + "!" }, map[string]string{"end.": "end.!", "start,": "start,!"}, 2},
		{"comma semicolon", map[string]int{"a,b": 1, "c;d": 2}, func(k string, v int) string { return k }, map[string]string{"a,b": "a,b", "c;d": "c;d"}, 2},
		{"long keys", map[string]int{"long_key_example_0001": 11, "long_key_example_0002": 22}, func(k string, v int) string { return k + ":ok" }, map[string]string{"long_key_example_0001": "long_key_example_0001:ok", "long_key_example_0002": "long_key_example_0002:ok"}, 2},
		{"many entries", map[string]int{"m1": 1, "m2": 2, "m3": 3, "m4": 4, "m5": 5}, func(k string, v int) string { return fmt.Sprintf("%s|%d", k, v) }, map[string]string{"m1": "m1|1", "m2": "m2|2", "m3": "m3|3", "m4": "m4|4", "m5": "m5|5"}, 5},
		{"zero value", map[string]int{"z": 0, "one": 1}, func(k string, v int) string { return fmt.Sprintf("%d", v) }, map[string]string{"z": "0", "one": "1"}, 2},
		{"negative values", map[string]int{"n1": -1, "n2": -2}, func(k string, v int) string { return fmt.Sprintf("%d", v) }, map[string]string{"n1": "-1", "n2": "-2"}, 2},
		{"mixed small", map[string]int{"x": 7, "y": 8, "z": 9}, func(k string, v int) string { return fmt.Sprintf("%s#%d", k, v) }, map[string]string{"x": "x#7", "y": "y#8", "z": "z#9"}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ValueBy(tt.input, tt.fn)

			if got == nil {
				t.Fatalf("ValueBy returned nil; want non-nil map for input %v", tt.input)
			}

			if len(got) != tt.expectSz {
				t.Fatalf("ValueBy(%v) length = %d; want %d; got=%v", tt.input, len(got), tt.expectSz, got)
			}

			for k, v := range tt.expects {
				if gv, ok := got[k]; !ok {
					t.Fatalf("ValueBy(%v) missing key %v; got=%v", tt.input, k, got)
				} else if gv != v {
					t.Fatalf("ValueBy(%v) for key %v expected value %v; got %v", tt.input, k, v, gv)
				}
			}
		})
	}
}

func TestValueBy_Struct(t *testing.T) {
	type OutV struct {
		Label string
		Num   int
		Flag  bool
	}

	tests := []struct {
		name    string
		input   map[string]int
		fn      func(string, int) OutV
		expects []OutV
	}{
		{"nil input", nil, func(k string, v int) OutV { return OutV{Label: k, Num: v, Flag: v%2 == 0} }, []OutV{}},
		{"empty input", map[string]int{}, func(k string, v int) OutV { return OutV{Label: fmt.Sprintf("k:%s", k), Num: v * 2, Flag: v%2 == 0} }, []OutV{}},
		{"single", map[string]int{"a": 1}, func(k string, v int) OutV { return OutV{Label: fmt.Sprintf("k:%s", k), Num: v * 2, Flag: false} }, []OutV{{Label: "k:a", Num: 2, Flag: false}}},
		{"unicode & emoji", map[string]int{"こんにちは": 1, "😊": 2}, func(k string, v int) OutV { return OutV{Label: k, Num: v, Flag: v > 0} }, []OutV{{Label: "こんにちは", Num: 1, Flag: true}, {Label: "😊", Num: 2, Flag: true}}},
		{"special chars", map[string]int{"!@#": 7, "$%": 8}, func(k string, v int) OutV { return OutV{Label: k + "#", Num: v, Flag: v%2 == 0} }, []OutV{{Label: "!@##", Num: 7, Flag: false}, {Label: "$%#", Num: 8, Flag: true}}},
		{"whitespace name", map[string]int{" ": 1, "\t": 2, "\n": 3}, func(k string, v int) OutV { return OutV{Label: fmt.Sprintf("ws:%q", k), Num: v, Flag: v%2 == 0} }, []OutV{{Label: "ws:\" \"", Num: 1, Flag: false}, {Label: "ws:\"\\t\"", Num: 2, Flag: true}, {Label: "ws:\"\\n\"", Num: 3, Flag: false}}},
		{"long name", map[string]int{"long_name_0001": 11, "long_name_0002": 22}, func(k string, v int) OutV { return OutV{Label: k, Num: v + 1000, Flag: v%10 == 0} }, []OutV{{Label: "long_name_0001", Num: 1011, Flag: false}, {Label: "long_name_0002", Num: 1022, Flag: false}}},
		{"zero and negative", map[string]int{"z": 0, "n": -1}, func(k string, v int) OutV { return OutV{Label: fmt.Sprintf("V:%d", v), Num: v, Flag: v >= 0} }, []OutV{{Label: "V:0", Num: 0, Flag: true}, {Label: "V:-1", Num: -1, Flag: false}}},
		{"many entries", map[string]int{"m1": 1, "m2": 2, "m3": 3, "m4": 4}, func(k string, v int) OutV { return OutV{Label: fmt.Sprintf("L:%s", k), Num: v * 10, Flag: v%2 == 0} }, []OutV{{Label: "L:m1", Num: 10, Flag: false}, {Label: "L:m2", Num: 20, Flag: true}, {Label: "L:m3", Num: 30, Flag: false}, {Label: "L:m4", Num: 40, Flag: true}}},
		{"mixed values", map[string]int{"a": 1, "b": -2, "c": 0}, func(k string, v int) OutV { return OutV{Label: k, Num: v * v, Flag: v%2 == 0} }, []OutV{{Label: "a", Num: 1, Flag: false}, {Label: "b", Num: 4, Flag: true}, {Label: "c", Num: 0, Flag: true}}},
		{"big numbers", map[string]int{"big": 1000000}, func(k string, v int) OutV { return OutV{Label: k, Num: int(math.Abs(float64(v))), Flag: v > 0} }, []OutV{{Label: "big", Num: 1000000, Flag: true}}},
		{"fn uses only key", map[string]int{"k1": 7, "k2": 8}, func(k string, v int) OutV { return OutV{Label: "K:" + k, Num: 0, Flag: false} }, []OutV{{Label: "K:k1", Num: 0, Flag: false}, {Label: "K:k2", Num: 0, Flag: false}}},
		{"fn constant value", map[string]int{"x": 9, "y": 10}, func(k string, v int) OutV { return OutV{Label: "const", Num: 42, Flag: true} }, []OutV{{Label: "const", Num: 42, Flag: true}, {Label: "const", Num: 42, Flag: true}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ValueBy(tt.input, tt.fn)
			if got == nil {
				t.Fatalf("ValueBy returned nil; want non-nil map")
			}

			if len(got) != len(tt.expects) {
				t.Fatalf("expected %d entries; got %d; got=%v", len(tt.expects), len(got), got)
			}

			vals := make([]OutV, 0, len(got))
			for _, v := range got {
				vals = append(vals, v)
			}

			for _, v := range tt.expects {
				if !lxslices.Contains(vals, v) {
					t.Fatalf("expected value %v missing in result %v", v, vals)
				}
			}
		})
	}
}
