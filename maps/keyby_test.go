package lxmaps_test

import (
	"fmt"
	"testing"

	"github.com/hgapdvn/lx/maps"
	"github.com/hgapdvn/lx/slices"
)

func TestKeyBy_String(t *testing.T) {
	tests := []struct {
		name    string
		input   map[string]int
		fn      func(string, int) string
		expects map[string]int
		extra   []int // allowed values when duplicate-target reduces to single key
	}{
		{"nil map", nil, func(k string, v int) string { return k }, map[string]int{}, nil},
		{"empty map", map[string]int{}, func(k string, v int) string { return k + "_x" }, map[string]int{}, nil},
		{"single entry", map[string]int{"a": 1}, func(k string, v int) string { return k }, map[string]int{"a": 1}, nil},
		{"two entries", map[string]int{"a": 1, "b": 2}, func(k string, v int) string { return k }, map[string]int{"a": 1, "b": 2}, nil},
		{"three entries", map[string]int{"a": 1, "b": 2, "c": 3}, func(k string, v int) string { return fmt.Sprintf("%s#%d", k, v) }, map[string]int{"a#1": 1, "b#2": 2, "c#3": 3}, nil},
		{"duplicate target", map[string]int{"a": 1, "b": 2}, func(k string, v int) string { return "same" }, nil, []int{1, 2}},
		{"unicode keys", map[string]int{"こんにちは": 1, "世界": 2}, func(k string, v int) string { return k + "_u" }, map[string]int{"こんにちは_u": 1, "世界_u": 2}, nil},
		{"emoji keys", map[string]int{"😊": 1, "👋": 2}, func(k string, v int) string { return k + "_e" }, map[string]int{"😊_e": 1, "👋_e": 2}, nil},
		{"special chars", map[string]int{"!@#": 7, "$%": 8}, func(k string, v int) string { return "k:" + k }, map[string]int{"k:!@#": 7, "k:$%": 8}, nil},
		{"whitespace keys", map[string]int{" ": 1, "\t": 2, "\n": 3}, func(k string, v int) string { return fmt.Sprintf("ws:%q", k) }, map[string]int{"ws:\" \"": 1, "ws:\"\\t\"": 2, "ws:\"\\n\"": 3}, nil},
		{"numeric string keys", map[string]int{"1": 1, "2": 2, "10": 10}, func(k string, v int) string { return "n:" + k }, map[string]int{"n:1": 1, "n:2": 2, "n:10": 10}, nil},
		{"case sensitivity", map[string]int{"Go": 1, "go": 2}, func(k string, v int) string { return k }, map[string]int{"Go": 1, "go": 2}, nil},
		{"prefix/suffix", map[string]int{"pre_": 1, "_suf": 2, "pre_suf": 3}, func(k string, v int) string { return "p:" + k }, map[string]int{"p:pre_": 1, "p:_suf": 2, "p:pre_suf": 3}, nil},
		{"punctuation", map[string]int{"end.": 1, "start,": 2, "(paren)": 3}, func(k string, v int) string { return k + "!" }, map[string]int{"end.!": 1, "start,!": 2, "(paren)!": 3}, nil},
		{"comma semicolon", map[string]int{"a,b": 1, "c;d": 2}, func(k string, v int) string { return k }, map[string]int{"a,b": 1, "c;d": 2}, nil},
		{"long keys", map[string]int{"long_key_example_0001": 11, "long_key_example_0002": 22}, func(k string, v int) string { return k + ":ok" }, map[string]int{"long_key_example_0001:ok": 11, "long_key_example_0002:ok": 22}, nil},
		{"many entries", map[string]int{"m1": 1, "m2": 2, "m3": 3, "m4": 4, "m5": 5}, func(k string, v int) string { return fmt.Sprintf("%s|%d", k, v) }, map[string]int{"m1|1": 1, "m2|2": 2, "m3|3": 3, "m4|4": 4, "m5|5": 5}, nil},
		{"zero value", map[string]int{"z": 0, "one": 1}, func(k string, v int) string { return k }, map[string]int{"z": 0, "one": 1}, nil},
		{"negative values", map[string]int{"n1": -1, "n2": -2}, func(k string, v int) string { return k }, map[string]int{"n1": -1, "n2": -2}, nil},
		{"mixed small", map[string]int{"x": 7, "y": 8, "z": 9}, func(k string, v int) string { return fmt.Sprintf("%s#%d", k, v) }, map[string]int{"x#7": 7, "y#8": 8, "z#9": 9}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.KeyBy(tt.input, tt.fn)

			if got == nil {
				t.Fatalf("KeyBy returned nil; want non-nil (empty or filled) for input %v", tt.input)
			}

			if tt.extra != nil {
				if len(got) != 1 {
					t.Fatalf("duplicate-target: expected 1 entry; got=%v", got)
				}
				for _, v := range got {
					if !lxslices.ContainsAny(tt.extra, v) {
						t.Fatalf("duplicate-target: value %v not in allowed set %v", v, tt.extra)
					}
				}
				return
			}

			if len(got) != len(tt.expects) {
				t.Fatalf("KeyBy(%v) length = %d; want %d; got=%v", tt.input, len(got), len(tt.expects), got)
			}

			for k, v := range tt.expects {
				if gv, ok := got[k]; !ok {
					t.Fatalf("KeyBy(%v) missing expected key %v; got=%v", tt.input, k, got)
				} else if gv != v {
					t.Fatalf("KeyBy(%v) for key %v expected value %v; got %v", tt.input, k, v, gv)
				}
			}
		})
	}
}

func TestKeyBy_Struct(t *testing.T) {
	// Comparable struct used as output key
	type K struct {
		ID   int
		Name string
	}

	tests := []struct {
		name    string
		input   map[string]int
		fn      func(string, int) K
		expects []K
	}{
		{"basic struct keys", map[string]int{"a": 1, "b": 2}, func(k string, v int) K { return K{ID: v, Name: k} }, []K{{ID: 1, Name: "a"}, {ID: 2, Name: "b"}}},
		{"duplicate struct key values", map[string]int{"x": 1, "y": 2}, func(k string, v int) K { return K{ID: 0, Name: "dup"} }, []K{{ID: 0, Name: "dup"}}},
		{"empty input", map[string]int{}, func(k string, v int) K { return K{ID: v, Name: k} }, []K{}},
		{"unicode name", map[string]int{"こんにちは": 1, "世界": 2}, func(k string, v int) K { return K{ID: v, Name: k} }, []K{{ID: 1, Name: "こんにちは"}, {ID: 2, Name: "世界"}}},
		{"emoji name", map[string]int{"😊": 1, "👋": 2}, func(k string, v int) K { return K{ID: v, Name: k} }, []K{{ID: 1, Name: "😊"}, {ID: 2, Name: "👋"}}},
		{"special chars name", map[string]int{"!@#": 7, "$%": 8}, func(k string, v int) K { return K{ID: v, Name: k} }, []K{{ID: 7, Name: "!@#"}, {ID: 8, Name: "$%"}}},
		{"whitespace name", map[string]int{" ": 1, "\t": 2}, func(k string, v int) K { return K{ID: v, Name: k} }, []K{{ID: 1, Name: " "}, {ID: 2, Name: "\t"}}},
		{"long name", map[string]int{"long_name_0001": 11, "long_name_0002": 22}, func(k string, v int) K { return K{ID: v, Name: k} }, []K{{ID: 11, Name: "long_name_0001"}, {ID: 22, Name: "long_name_0002"}}},
		{"zero id", map[string]int{"z": 0}, func(k string, v int) K { return K{ID: v, Name: k} }, []K{{ID: 0, Name: "z"}}},
		{"negative id", map[string]int{"n": -1}, func(k string, v int) K { return K{ID: v, Name: k} }, []K{{ID: -1, Name: "n"}}},
		{"many entries", map[string]int{"m1": 1, "m2": 2, "m3": 3, "m4": 4}, func(k string, v int) K { return K{ID: v, Name: k} }, []K{{ID: 1, Name: "m1"}, {ID: 2, Name: "m2"}, {ID: 3, Name: "m3"}, {ID: 4, Name: "m4"}}},
		{"mixed values", map[string]int{"a": 1, "b": -2, "c": 0}, func(k string, v int) K { return K{ID: v, Name: k} }, []K{{ID: 1, Name: "a"}, {ID: -2, Name: "b"}, {ID: 0, Name: "c"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.KeyBy(tt.input, tt.fn)
			if got == nil {
				t.Fatalf("KeyBy returned nil; want non-nil map")
			}

			if len(got) != len(tt.expects) {
				t.Fatalf("expected %d entries; got %d; got=%v", len(tt.expects), len(got), got)
			}

			for _, k := range tt.expects {
				if _, ok := got[k]; !ok {
					t.Fatalf("expected struct key %v missing in result %v", k, got)
				}
			}
		})
	}
}
