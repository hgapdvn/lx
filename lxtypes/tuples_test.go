package lxtypes_test

import (
	"strings"
	"testing"

	"github.com/hgapdvn/lx/lxtypes"
)

// ============================================================================
// Pair Tests
// ============================================================================

func TestNewPair(t *testing.T) {
	tests := []struct {
		name   string
		first  int
		second string
		want   lxtypes.Pair[int, string]
	}{
		{
			name:   "basic pair",
			first:  42,
			second: "answer",
			want:   lxtypes.Pair[int, string]{First: 42, Second: "answer"},
		},
		{
			name:   "zero values",
			first:  0,
			second: "",
			want:   lxtypes.Pair[int, string]{First: 0, Second: ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxtypes.NewPair(tt.first, tt.second)
			if got.First != tt.want.First || got.Second != tt.want.Second {
				t.Errorf("NewPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPair_Values(t *testing.T) {
	p := lxtypes.NewPair(10, "test")
	first, second := p.Values()

	if first != 10 {
		t.Errorf("Values() first = %v, want 10", first)
	}
	if second != "test" {
		t.Errorf("Values() second = %v, want test", second)
	}
}

func TestPair_Swap(t *testing.T) {
	tests := []struct {
		name  string
		input lxtypes.Pair[int, string]
		want  lxtypes.Pair[string, int]
	}{
		{
			name:  "swap pair",
			input: lxtypes.NewPair(42, "hello"),
			want:  lxtypes.Pair[string, int]{First: "hello", Second: 42},
		},
		{
			name:  "swap zero values",
			input: lxtypes.NewPair(0, ""),
			want:  lxtypes.Pair[string, int]{First: "", Second: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Swap()
			if got.First != tt.want.First || got.Second != tt.want.Second {
				t.Errorf("Swap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPair_MapFirst(t *testing.T) {
	p := lxtypes.NewPair(5, "test")
	double := func(n int) int { return n * 2 }

	got := p.MapFirst(double)

	if got.First != 10 {
		t.Errorf("MapFirst() First = %v, want 10", got.First)
	}
	if got.Second != "test" {
		t.Errorf("MapFirst() Second = %v, want test", got.Second)
	}
}

func TestPair_MapSecond(t *testing.T) {
	p := lxtypes.NewPair(5, "test")
	upper := func(s string) string { return strings.ToUpper(s) }

	got := p.MapSecond(upper)

	if got.First != 5 {
		t.Errorf("MapSecond() First = %v, want 5", got.First)
	}
	if got.Second != "TEST" {
		t.Errorf("MapSecond() Second = %v, want TEST", got.Second)
	}
}

func TestPair_ZeroValue(t *testing.T) {
	var zero lxtypes.Pair[int, string]
	if zero.First != 0 || zero.Second != "" {
		t.Errorf("zero value Pair = %v, want {0, \"\"}", zero)
	}
}

func TestPair_DifferentTypes(t *testing.T) {
	// Test with various type combinations
	t.Run("int and bool", func(t *testing.T) {
		p := lxtypes.NewPair(42, true)
		if p.First != 42 || p.Second != true {
			t.Errorf("Pair[int, bool] = %v, want {42, true}", p)
		}
	})

	t.Run("string and float", func(t *testing.T) {
		p := lxtypes.NewPair("pi", 3.14)
		if p.First != "pi" || p.Second != 3.14 {
			t.Errorf("Pair[string, float64] = %v, want {\"pi\", 3.14}", p)
		}
	})

	t.Run("struct types", func(t *testing.T) {
		type Person struct{ Name string }
		type Age struct{ Years int }
		p := lxtypes.NewPair(Person{"Alice"}, Age{30})
		if p.First.Name != "Alice" || p.Second.Years != 30 {
			t.Errorf("Pair[Person, Age] = %v, want {{\"Alice\"}, {30}}", p)
		}
	})
}

// ============================================================================
// Triple Tests
// ============================================================================

func TestNewTriple(t *testing.T) {
	tests := []struct {
		name   string
		first  int
		second string
		third  bool
		want   lxtypes.Triple[int, string, bool]
	}{
		{
			name:   "basic triple",
			first:  42,
			second: "answer",
			third:  true,
			want:   lxtypes.Triple[int, string, bool]{First: 42, Second: "answer", Third: true},
		},
		{
			name:   "zero values",
			first:  0,
			second: "",
			third:  false,
			want:   lxtypes.Triple[int, string, bool]{First: 0, Second: "", Third: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxtypes.NewTriple(tt.first, tt.second, tt.third)
			if got.First != tt.want.First || got.Second != tt.want.Second || got.Third != tt.want.Third {
				t.Errorf("NewTriple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTriple_Values(t *testing.T) {
	tr := lxtypes.NewTriple(10, "test", true)
	first, second, third := tr.Values()

	if first != 10 {
		t.Errorf("Values() first = %v, want 10", first)
	}
	if second != "test" {
		t.Errorf("Values() second = %v, want test", second)
	}
	if third != true {
		t.Errorf("Values() third = %v, want true", third)
	}
}

func TestTriple_ToPair(t *testing.T) {
	tr := lxtypes.NewTriple(42, "hello", true)
	got := tr.ToPair()

	want := lxtypes.Pair[int, string]{First: 42, Second: "hello"}
	if got.First != want.First || got.Second != want.Second {
		t.Errorf("ToPair() = %v, want %v", got, want)
	}
}

func TestTriple_ZeroValue(t *testing.T) {
	var zero lxtypes.Triple[int, string, bool]
	if zero.First != 0 || zero.Second != "" || zero.Third != false {
		t.Errorf("zero value Triple = %v, want {0, \"\", false}", zero)
	}
}

func TestTriple_DifferentTypes(t *testing.T) {
	t.Run("numeric types", func(t *testing.T) {
		tr := lxtypes.NewTriple(1, 2.5, 3.14159)
		if tr.First != 1 || tr.Second != 2.5 || tr.Third != 3.14159 {
			t.Errorf("Triple[int, float32, float64] = %v", tr)
		}
	})

	t.Run("mixed types", func(t *testing.T) {
		type Point struct{ X, Y int }
		tr := lxtypes.NewTriple("coord", Point{10, 20}, []int{1, 2, 3})
		if tr.First != "coord" || tr.Second.X != 10 || len(tr.Third) != 3 {
			t.Errorf("Triple[string, Point, []int] = %v", tr)
		}
	})
}

// ============================================================================
// Quad Tests
// ============================================================================

func TestNewQuad(t *testing.T) {
	tests := []struct {
		name   string
		first  int
		second string
		third  bool
		fourth float64
		want   lxtypes.Quad[int, string, bool, float64]
	}{
		{
			name:   "basic quad",
			first:  42,
			second: "answer",
			third:  true,
			fourth: 3.14,
			want:   lxtypes.Quad[int, string, bool, float64]{First: 42, Second: "answer", Third: true, Fourth: 3.14},
		},
		{
			name:   "zero values",
			first:  0,
			second: "",
			third:  false,
			fourth: 0.0,
			want:   lxtypes.Quad[int, string, bool, float64]{First: 0, Second: "", Third: false, Fourth: 0.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxtypes.NewQuad(tt.first, tt.second, tt.third, tt.fourth)
			if got.First != tt.want.First || got.Second != tt.want.Second ||
				got.Third != tt.want.Third || got.Fourth != tt.want.Fourth {
				t.Errorf("NewQuad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuad_Values(t *testing.T) {
	q := lxtypes.NewQuad(10, "test", true, 3.14)
	first, second, third, fourth := q.Values()

	if first != 10 {
		t.Errorf("Values() first = %v, want 10", first)
	}
	if second != "test" {
		t.Errorf("Values() second = %v, want test", second)
	}
	if third != true {
		t.Errorf("Values() third = %v, want true", third)
	}
	if fourth != 3.14 {
		t.Errorf("Values() fourth = %v, want 3.14", fourth)
	}
}

func TestQuad_ToPair(t *testing.T) {
	q := lxtypes.NewQuad(42, "hello", true, 3.14)
	got := q.ToPair()

	want := lxtypes.Pair[int, string]{First: 42, Second: "hello"}
	if got.First != want.First || got.Second != want.Second {
		t.Errorf("ToPair() = %v, want %v", got, want)
	}
}

func TestQuad_ToTriple(t *testing.T) {
	q := lxtypes.NewQuad(42, "hello", true, 3.14)
	got := q.ToTriple()

	want := lxtypes.Triple[int, string, bool]{First: 42, Second: "hello", Third: true}
	if got.First != want.First || got.Second != want.Second || got.Third != want.Third {
		t.Errorf("ToTriple() = %v, want %v", got, want)
	}
}

func TestQuad_ZeroValue(t *testing.T) {
	var zero lxtypes.Quad[int, string, bool, float64]
	if zero.First != 0 || zero.Second != "" || zero.Third != false || zero.Fourth != 0.0 {
		t.Errorf("zero value Quad = %v, want {0, \"\", false, 0.0}", zero)
	}
}

func TestQuad_DifferentTypes(t *testing.T) {
	t.Run("all different types", func(t *testing.T) {
		q := lxtypes.NewQuad(42, "hello", true, 3.14)
		if q.First != 42 || q.Second != "hello" || q.Third != true || q.Fourth != 3.14 {
			t.Errorf("Quad = %v", q)
		}
	})

	t.Run("complex types", func(t *testing.T) {
		type User struct{ Name string }
		q := lxtypes.NewQuad(
			[]int{1, 2, 3},
			map[string]int{"a": 1},
			User{"Alice"},
			func() string { return "test" },
		)
		if len(q.First) != 3 || q.Second["a"] != 1 || q.Third.Name != "Alice" || q.Fourth() != "test" {
			t.Errorf("Quad with complex types failed")
		}
	})
}

// ============================================================================
// Tuple5 Tests
// ============================================================================

func TestNewTuple5(t *testing.T) {
	tests := []struct {
		name string
		v1   int
		v2   string
		v3   bool
		v4   float64
		v5   []int
	}{
		{
			name: "basic tuple5",
			v1:   1,
			v2:   "two",
			v3:   true,
			v4:   4.0,
			v5:   []int{5, 6},
		},
		{
			name: "zero values",
			v1:   0,
			v2:   "",
			v3:   false,
			v4:   0.0,
			v5:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxtypes.NewTuple5(tt.v1, tt.v2, tt.v3, tt.v4, tt.v5)
			if got.V1 != tt.v1 || got.V2 != tt.v2 || got.V3 != tt.v3 || got.V4 != tt.v4 {
				t.Errorf("NewTuple5() = %v, want {%v, %v, %v, %v, %v}", got, tt.v1, tt.v2, tt.v3, tt.v4, tt.v5)
			}
		})
	}
}

func TestTuple5_Values(t *testing.T) {
	tup := lxtypes.NewTuple5(1, "two", true, 4.0, []int{5})
	v1, v2, v3, v4, v5 := tup.Values()

	if v1 != 1 {
		t.Errorf("Values() v1 = %v, want 1", v1)
	}
	if v2 != "two" {
		t.Errorf("Values() v2 = %v, want two", v2)
	}
	if v3 != true {
		t.Errorf("Values() v3 = %v, want true", v3)
	}
	if v4 != 4.0 {
		t.Errorf("Values() v4 = %v, want 4.0", v4)
	}
	if len(v5) != 1 || v5[0] != 5 {
		t.Errorf("Values() v5 = %v, want [5]", v5)
	}
}

func TestTuple5_DifferentTypes(t *testing.T) {
	tup := lxtypes.NewTuple5(42, "answer", true, 3.14, []string{"a", "b"})
	if tup.V1 != 42 || tup.V2 != "answer" || tup.V3 != true || tup.V4 != 3.14 || len(tup.V5) != 2 {
		t.Errorf("Tuple5 with different types failed")
	}
}

// ============================================================================
// Tuple6 Tests
// ============================================================================

func TestNewTuple6(t *testing.T) {
	tests := []struct {
		name string
		v1   int
		v2   string
		v3   bool
		v4   float64
		v5   []int
		v6   rune
	}{
		{
			name: "basic tuple6",
			v1:   1,
			v2:   "two",
			v3:   true,
			v4:   4.0,
			v5:   []int{5},
			v6:   'a',
		},
		{
			name: "zero values",
			v1:   0,
			v2:   "",
			v3:   false,
			v4:   0.0,
			v5:   nil,
			v6:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxtypes.NewTuple6(tt.v1, tt.v2, tt.v3, tt.v4, tt.v5, tt.v6)
			if got.V1 != tt.v1 || got.V2 != tt.v2 || got.V3 != tt.v3 || got.V4 != tt.v4 || got.V6 != tt.v6 {
				t.Errorf("NewTuple6() failed")
			}
		})
	}
}

func TestTuple6_Values(t *testing.T) {
	tup := lxtypes.NewTuple6(1, "two", true, 4.0, []int{5}, 'a')
	v1, v2, v3, v4, v5, v6 := tup.Values()

	if v1 != 1 || v2 != "two" || v3 != true || v4 != 4.0 || len(v5) != 1 || v6 != 'a' {
		t.Errorf("Tuple6.Values() failed")
	}
}

// ============================================================================
// Tuple7 Tests
// ============================================================================

func TestNewTuple7(t *testing.T) {
	tup := lxtypes.NewTuple7(1, "two", true, 4.0, []int{5}, 'a', byte(7))
	if tup.V1 != 1 || tup.V2 != "two" || tup.V3 != true || tup.V4 != 4.0 || tup.V6 != 'a' || tup.V7 != 7 {
		t.Errorf("NewTuple7() failed")
	}
}

func TestTuple7_Values(t *testing.T) {
	tup := lxtypes.NewTuple7(1, "two", true, 4.0, []int{5}, 'a', byte(7))
	v1, v2, v3, v4, v5, v6, v7 := tup.Values()

	if v1 != 1 || v2 != "two" || v3 != true || v4 != 4.0 || len(v5) != 1 || v6 != 'a' || v7 != 7 {
		t.Errorf("Tuple7.Values() failed")
	}
}

// ============================================================================
// Tuple8 Tests
// ============================================================================

func TestNewTuple8(t *testing.T) {
	tup := lxtypes.NewTuple8(1, "two", true, 4.0, []int{5}, 'a', byte(7), uint(8))
	if tup.V1 != 1 || tup.V2 != "two" || tup.V3 != true || tup.V4 != 4.0 || tup.V6 != 'a' || tup.V7 != 7 || tup.V8 != 8 {
		t.Errorf("NewTuple8() failed")
	}
}

func TestTuple8_Values(t *testing.T) {
	tup := lxtypes.NewTuple8(1, "two", true, 4.0, []int{5}, 'a', byte(7), uint(8))
	v1, v2, v3, v4, v5, v6, v7, v8 := tup.Values()

	if v1 != 1 || v2 != "two" || v3 != true || v4 != 4.0 || len(v5) != 1 || v6 != 'a' || v7 != 7 || v8 != 8 {
		t.Errorf("Tuple8.Values() failed")
	}
}

func TestTuple8_ZeroValue(t *testing.T) {
	var tup lxtypes.Tuple8[int, string, bool, float64, []int, rune, byte, uint]
	if tup.V1 != 0 || tup.V2 != "" || tup.V3 != false || tup.V4 != 0.0 || tup.V5 != nil || tup.V6 != 0 || tup.V7 != 0 || tup.V8 != 0 {
		t.Errorf("Tuple8 zero value failed")
	}
}

// ============================================================================
// Integration Tests
// ============================================================================

func TestTupleChaining(t *testing.T) {
	// Test converting between tuple types
	q := lxtypes.NewQuad(1, "hello", true, 3.14)
	tr := q.ToTriple()
	p := tr.ToPair()

	if p.First != 1 || p.Second != "hello" {
		t.Errorf("Chained conversion failed: got %v", p)
	}
}

func TestTupleInSlice(t *testing.T) {
	// Test using tuples in slices (common use case)
	pairs := []lxtypes.Pair[int, string]{
		lxtypes.NewPair(1, "one"),
		lxtypes.NewPair(2, "two"),
		lxtypes.NewPair(3, "three"),
	}

	if len(pairs) != 3 {
		t.Errorf("len(pairs) = %v, want 3", len(pairs))
	}
	if pairs[1].First != 2 || pairs[1].Second != "two" {
		t.Errorf("pairs[1] = %v, want {2, \"two\"}", pairs[1])
	}
}

func TestTupleInMap(t *testing.T) {
	// Test using tuples as map values
	m := map[string]lxtypes.Pair[int, bool]{
		"a": lxtypes.NewPair(1, true),
		"b": lxtypes.NewPair(2, false),
	}

	if m["a"].First != 1 || m["a"].Second != true {
		t.Errorf("m[\"a\"] = %v, want {1, true}", m["a"])
	}
}
