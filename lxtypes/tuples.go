package lxtypes

// Pair represents a generic tuple of two values.
// Useful for functions that need to return or work with paired values.
//
// Example:
//
//	p := lxtypes.NewPair(1, "hello")
//	fmt.Println(p.First)   // 1
//	fmt.Println(p.Second)  // "hello"
type Pair[T, U any] struct {
	First  T
	Second U
}

// NewPair creates a new Pair with the given values.
//
// Example:
//
//	p := lxtypes.NewPair(42, "answer")
func NewPair[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{First: first, Second: second}
}

// Values returns the two values as separate return values.
// Useful for unpacking a pair.
//
// Example:
//
//	p := lxtypes.NewPair(1, "hello")
//	x, y := p.Values()
func (p Pair[T, U]) Values() (T, U) {
	return p.First, p.Second
}

// Swap returns a new Pair with First and Second swapped.
//
// Example:
//
//	p := lxtypes.NewPair(1, "hello")
//	swapped := p.Swap()  // Pair[string, int]{"hello", 1}
func (p Pair[T, U]) Swap() Pair[U, T] {
	return Pair[U, T]{First: p.Second, Second: p.First}
}

// MapFirst applies a function to the First value, returning a new Pair.
//
// Example:
//
//	p := lxtypes.NewPair(5, "test")
//	doubled := p.MapFirst(func(n int) int { return n * 2 })
//	// doubled.First == 10, doubled.Second == "test"
func (p Pair[T, U]) MapFirst(fn func(T) T) Pair[T, U] {
	return Pair[T, U]{First: fn(p.First), Second: p.Second}
}

// MapSecond applies a function to the Second value, returning a new Pair.
//
// Example:
//
//	p := lxtypes.NewPair(5, "test")
//	upper := p.MapSecond(func(s string) string { return strings.ToUpper(s) })
//	// upper.First == 5, upper.Second == "TEST"
func (p Pair[T, U]) MapSecond(fn func(U) U) Pair[T, U] {
	return Pair[T, U]{First: p.First, Second: fn(p.Second)}
}

// Triple represents a generic tuple of three values.
//
// Example:
//
//	t := lxtypes.NewTriple(1, "hello", true)
type Triple[T, U, V any] struct {
	First  T
	Second U
	Third  V
}

// NewTriple creates a new Triple with the given values.
//
// Example:
//
//	t := lxtypes.NewTriple(42, "answer", true)
func NewTriple[T, U, V any](first T, second U, third V) Triple[T, U, V] {
	return Triple[T, U, V]{First: first, Second: second, Third: third}
}

// Values returns the three values as separate return values.
// Useful for unpacking a triple.
//
// Example:
//
//	t := lxtypes.NewTriple(1, "hello", true)
//	x, y, z := t.Values()
func (t Triple[T, U, V]) Values() (T, U, V) {
	return t.First, t.Second, t.Third
}

// ToPair returns a Pair containing the first two values, discarding the third.
//
// Example:
//
//	t := lxtypes.NewTriple(1, "hello", true)
//	p := t.ToPair()  // Pair[int, string]{1, "hello"}
func (t Triple[T, U, V]) ToPair() Pair[T, U] {
	return Pair[T, U]{First: t.First, Second: t.Second}
}

// Quad represents a generic tuple of four values.
//
// Example:
//
//	q := lxtypes.NewQuad(1, "hello", true, 3.14)
type Quad[T, U, V, W any] struct {
	First  T
	Second U
	Third  V
	Fourth W
}

// NewQuad creates a new Quad with the given values.
//
// Example:
//
//	q := lxtypes.NewQuad(42, "answer", true, 3.14)
func NewQuad[T, U, V, W any](first T, second U, third V, fourth W) Quad[T, U, V, W] {
	return Quad[T, U, V, W]{First: first, Second: second, Third: third, Fourth: fourth}
}

// Values returns the four values as separate return values.
// Useful for unpacking a quad.
//
// Example:
//
//	q := lxtypes.NewQuad(1, "hello", true, 3.14)
//	w, x, y, z := q.Values()
func (q Quad[T, U, V, W]) Values() (T, U, V, W) {
	return q.First, q.Second, q.Third, q.Fourth
}

// ToPair returns a Pair containing the first two values, discarding the rest.
//
// Example:
//
//	q := lxtypes.NewQuad(1, "hello", true, 3.14)
//	p := q.ToPair()  // Pair[int, string]{1, "hello"}
func (q Quad[T, U, V, W]) ToPair() Pair[T, U] {
	return Pair[T, U]{First: q.First, Second: q.Second}
}

// ToTriple returns a Triple containing the first three values, discarding the fourth.
//
// Example:
//
//	q := lxtypes.NewQuad(1, "hello", true, 3.14)
//	t := q.ToTriple()  // Triple[int, string, bool]{1, "hello", true}
func (q Quad[T, U, V, W]) ToTriple() Triple[T, U, V] {
	return Triple[T, U, V]{First: q.First, Second: q.Second, Third: q.Third}
}

// Tuple5 represents a generic tuple of five values.
//
// Example:
//
//	t := lxtypes.NewTuple5(1, "hello", true, 3.14, []int{1, 2})
type Tuple5[T1, T2, T3, T4, T5 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
}

// NewTuple5 creates a new Tuple5 with the given values.
//
// Example:
//
//	t := lxtypes.NewTuple5(42, "answer", true, 3.14, []int{1, 2})
func NewTuple5[T1, T2, T3, T4, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) Tuple5[T1, T2, T3, T4, T5] {
	return Tuple5[T1, T2, T3, T4, T5]{V1: v1, V2: v2, V3: v3, V4: v4, V5: v5}
}

// Values returns the five values as separate return values.
//
// Example:
//
//	t := lxtypes.NewTuple5(1, "hello", true, 3.14, []int{1, 2})
//	v1, v2, v3, v4, v5 := t.Values()
func (t Tuple5[T1, T2, T3, T4, T5]) Values() (T1, T2, T3, T4, T5) {
	return t.V1, t.V2, t.V3, t.V4, t.V5
}

// Tuple6 represents a generic tuple of six values.
//
// Example:
//
//	t := lxtypes.NewTuple6(1, "hello", true, 3.14, []int{1, 2}, 'a')
type Tuple6[T1, T2, T3, T4, T5, T6 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
}

// NewTuple6 creates a new Tuple6 with the given values.
//
// Example:
//
//	t := lxtypes.NewTuple6(42, "answer", true, 3.14, []int{1, 2}, 'a')
func NewTuple6[T1, T2, T3, T4, T5, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) Tuple6[T1, T2, T3, T4, T5, T6] {
	return Tuple6[T1, T2, T3, T4, T5, T6]{V1: v1, V2: v2, V3: v3, V4: v4, V5: v5, V6: v6}
}

// Values returns the six values as separate return values.
//
// Example:
//
//	t := lxtypes.NewTuple6(1, "hello", true, 3.14, []int{1, 2}, 'a')
//	v1, v2, v3, v4, v5, v6 := t.Values()
func (t Tuple6[T1, T2, T3, T4, T5, T6]) Values() (T1, T2, T3, T4, T5, T6) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6
}

// Tuple7 represents a generic tuple of seven values.
//
// Example:
//
//	t := lxtypes.NewTuple7(1, "hello", true, 3.14, []int{1, 2}, 'a', 100)
type Tuple7[T1, T2, T3, T4, T5, T6, T7 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
}

// NewTuple7 creates a new Tuple7 with the given values.
//
// Example:
//
//	t := lxtypes.NewTuple7(42, "answer", true, 3.14, []int{1, 2}, 'a', 100)
func NewTuple7[T1, T2, T3, T4, T5, T6, T7 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) Tuple7[T1, T2, T3, T4, T5, T6, T7] {
	return Tuple7[T1, T2, T3, T4, T5, T6, T7]{V1: v1, V2: v2, V3: v3, V4: v4, V5: v5, V6: v6, V7: v7}
}

// Values returns the seven values as separate return values.
//
// Example:
//
//	t := lxtypes.NewTuple7(1, "hello", true, 3.14, []int{1, 2}, 'a', 100)
//	v1, v2, v3, v4, v5, v6, v7 := t.Values()
func (t Tuple7[T1, T2, T3, T4, T5, T6, T7]) Values() (T1, T2, T3, T4, T5, T6, T7) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7
}

// Tuple8 represents a generic tuple of eight values.
//
// Example:
//
//	t := lxtypes.NewTuple8(1, "hello", true, 3.14, []int{1, 2}, 'a', 100, "end")
type Tuple8[T1, T2, T3, T4, T5, T6, T7, T8 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T5
	V6 T6
	V7 T7
	V8 T8
}

// NewTuple8 creates a new Tuple8 with the given values.
//
// Example:
//
//	t := lxtypes.NewTuple8(42, "answer", true, 3.14, []int{1, 2}, 'a', 100, "end")
func NewTuple8[T1, T2, T3, T4, T5, T6, T7, T8 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) Tuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{V1: v1, V2: v2, V3: v3, V4: v4, V5: v5, V6: v6, V7: v7, V8: v8}
}

// Values returns the eight values as separate return values.
//
// Example:
//
//	t := lxtypes.NewTuple8(1, "hello", true, 3.14, []int{1, 2}, 'a', 100, "end")
//	v1, v2, v3, v4, v5, v6, v7, v8 := t.Values()
func (t Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Values() (T1, T2, T3, T4, T5, T6, T7, T8) {
	return t.V1, t.V2, t.V3, t.V4, t.V5, t.V6, t.V7, t.V8
}
