package lxslices_test

import (
	"strings"
	"testing"

	"github.com/hgapdvn/lx/slices"
)

func TestEqual_Int(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		b        []int
		expected bool
	}{
		{
			name:     "equal slices",
			a:        []int{1, 2, 3},
			b:        []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "different lengths",
			a:        []int{1, 2, 3},
			b:        []int{1, 2},
			expected: false,
		},
		{
			name:     "different elements",
			a:        []int{1, 2, 3},
			b:        []int{1, 2, 4},
			expected: false,
		},
		{
			name:     "both empty",
			a:        []int{},
			b:        []int{},
			expected: true,
		},
		{
			name:     "both nil",
			a:        nil,
			b:        nil,
			expected: true,
		},
		{
			name:     "nil vs empty",
			a:        nil,
			b:        []int{},
			expected: false,
		},
		{
			name:     "empty vs nil",
			a:        []int{},
			b:        nil,
			expected: false,
		},
		{
			name:     "first nil",
			a:        nil,
			b:        []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "second nil",
			a:        []int{1, 2, 3},
			b:        nil,
			expected: false,
		},
		{
			name:     "single element equal",
			a:        []int{1},
			b:        []int{1},
			expected: true,
		},
		{
			name:     "single element different",
			a:        []int{1},
			b:        []int{2},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Equal(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Equal(%v, %v) = %v; want %v",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestEqual_String(t *testing.T) {
	tests := []struct {
		name     string
		a        []string
		b        []string
		expected bool
	}{
		{
			name:     "equal slices",
			a:        []string{"a", "b", "c"},
			b:        []string{"a", "b", "c"},
			expected: true,
		},
		{
			name:     "different lengths",
			a:        []string{"a", "b", "c"},
			b:        []string{"a", "b"},
			expected: false,
		},
		{
			name:     "different elements",
			a:        []string{"a", "b", "c"},
			b:        []string{"a", "b", "d"},
			expected: false,
		},
		{
			name:     "both empty",
			a:        []string{},
			b:        []string{},
			expected: true,
		},
		{
			name:     "both nil",
			a:        nil,
			b:        nil,
			expected: true,
		},
		{
			name:     "nil vs empty",
			a:        nil,
			b:        []string{},
			expected: false,
		},
		{
			name:     "empty vs nil",
			a:        []string{},
			b:        nil,
			expected: false,
		},
		{
			name:     "first nil",
			a:        nil,
			b:        []string{"a", "b", "c"},
			expected: false,
		},
		{
			name:     "second nil",
			a:        []string{"a", "b", "c"},
			b:        nil,
			expected: false,
		},
		{
			name:     "single element equal",
			a:        []string{"hello"},
			b:        []string{"hello"},
			expected: true,
		},
		{
			name:     "single element different",
			a:        []string{"hello"},
			b:        []string{"world"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Equal(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Equal(%v, %v) = %v; want %v",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestEqual_Struct(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	tests := []struct {
		name     string
		a        []User
		b        []User
		expected bool
	}{
		{
			name:     "equal structs",
			a:        []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
			b:        []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
			expected: true,
		},
		{
			name:     "different lengths",
			a:        []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
			b:        []User{{1, "Alice"}, {2, "Bob"}},
			expected: false,
		},
		{
			name:     "different structs",
			a:        []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
			b:        []User{{1, "Alice"}, {2, "Bob"}, {4, "David"}},
			expected: false,
		},
		{
			name:     "both empty",
			a:        []User{},
			b:        []User{},
			expected: true,
		},
		{
			name:     "both nil",
			a:        nil,
			b:        nil,
			expected: true,
		},
		{
			name:     "nil vs empty",
			a:        nil,
			b:        []User{},
			expected: false,
		},
		{
			name:     "empty vs nil",
			a:        []User{},
			b:        nil,
			expected: false,
		},
		{
			name:     "first nil",
			a:        nil,
			b:        []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
			expected: false,
		},
		{
			name:     "second nil",
			a:        []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
			b:        nil,
			expected: false,
		},
		{
			name:     "single element equal",
			a:        []User{{1, "Alice"}},
			b:        []User{{1, "Alice"}},
			expected: true,
		},
		{
			name:     "single element different",
			a:        []User{{1, "Alice"}},
			b:        []User{{2, "Bob"}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Equal(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Equal(%v, %v) = %v; want %v",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestEqualFunc_Int(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		b        []int
		eq       func(int, int) bool
		expected bool
	}{
		{
			name:     "equal with default comparison",
			a:        []int{1, 2, 3},
			b:        []int{1, 2, 3},
			eq:       func(a, b int) bool { return a == b },
			expected: true,
		},
		{
			name:     "equal with custom comparison (absolute value)",
			a:        []int{-1, -2, -3},
			b:        []int{1, 2, 3},
			eq:       func(a, b int) bool { return abs(a) == abs(b) },
			expected: true,
		},
		{
			name:     "different with custom comparison",
			a:        []int{1, 2, 3},
			b:        []int{1, 2, 4},
			eq:       func(a, b int) bool { return a == b },
			expected: false,
		},
		{
			name:     "both empty",
			a:        []int{},
			b:        []int{},
			eq:       func(a, b int) bool { return a == b },
			expected: true,
		},
		{
			name:     "both nil",
			a:        nil,
			b:        nil,
			eq:       func(a, b int) bool { return a == b },
			expected: true,
		},
		{
			name:     "nil vs empty",
			a:        nil,
			b:        []int{},
			eq:       func(a, b int) bool { return a == b },
			expected: false,
		},
		{
			name:     "different lengths",
			a:        []int{1, 2, 3},
			b:        []int{1, 2},
			eq:       func(a, b int) bool { return a == b },
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.EqualFunc(tt.a, tt.b, tt.eq)
			if result != tt.expected {
				t.Errorf("EqualFunc(%v, %v) = %v; want %v",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestEqualFunc_String(t *testing.T) {
	tests := []struct {
		name     string
		a        []string
		b        []string
		eq       func(string, string) bool
		expected bool
	}{
		{
			name:     "equal with default comparison",
			a:        []string{"a", "b", "c"},
			b:        []string{"a", "b", "c"},
			eq:       func(a, b string) bool { return a == b },
			expected: true,
		},
		{
			name:     "equal with custom comparison (case insensitive)",
			a:        []string{"hello", "world"},
			b:        []string{"HELLO", "WORLD"},
			eq:       func(a, b string) bool { return strings.EqualFold(a, b) },
			expected: true,
		},
		{
			name:     "different with custom comparison",
			a:        []string{"a", "b", "c"},
			b:        []string{"a", "b", "d"},
			eq:       func(a, b string) bool { return a == b },
			expected: false,
		},
		{
			name:     "both empty",
			a:        []string{},
			b:        []string{},
			eq:       func(a, b string) bool { return a == b },
			expected: true,
		},
		{
			name:     "both nil",
			a:        nil,
			b:        nil,
			eq:       func(a, b string) bool { return a == b },
			expected: true,
		},
		{
			name:     "nil vs empty",
			a:        nil,
			b:        []string{},
			eq:       func(a, b string) bool { return a == b },
			expected: false,
		},
		{
			name:     "different lengths",
			a:        []string{"a", "b", "c"},
			b:        []string{"a", "b"},
			eq:       func(a, b string) bool { return a == b },
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.EqualFunc(tt.a, tt.b, tt.eq)
			if result != tt.expected {
				t.Errorf("EqualFunc(%v, %v) = %v; want %v",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestEqualFunc_Struct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	tests := []struct {
		name     string
		a        []Person
		b        []Person
		eq       func(Person, Person) bool
		expected bool
	}{
		{
			name:     "equal with default comparison",
			a:        []Person{{"Alice", 30}, {"Bob", 25}},
			b:        []Person{{"Alice", 30}, {"Bob", 25}},
			eq:       func(p1, p2 Person) bool { return p1.Name == p2.Name && p1.Age == p2.Age },
			expected: true,
		},
		{
			name:     "equal with custom comparison (name only)",
			a:        []Person{{"Alice", 30}, {"Bob", 25}},
			b:        []Person{{"Alice", 31}, {"Bob", 26}},
			eq:       func(p1, p2 Person) bool { return p1.Name == p2.Name },
			expected: true,
		},
		{
			name:     "different with custom comparison",
			a:        []Person{{"Alice", 30}, {"Bob", 25}},
			b:        []Person{{"Alice", 30}, {"Charlie", 25}},
			eq:       func(p1, p2 Person) bool { return p1.Name == p2.Name },
			expected: false,
		},
		{
			name:     "both empty",
			a:        []Person{},
			b:        []Person{},
			eq:       func(p1, p2 Person) bool { return p1.Name == p2.Name },
			expected: true,
		},
		{
			name:     "both nil",
			a:        nil,
			b:        nil,
			eq:       func(p1, p2 Person) bool { return p1.Name == p2.Name },
			expected: true,
		},
		{
			name:     "nil vs empty",
			a:        nil,
			b:        []Person{},
			eq:       func(p1, p2 Person) bool { return p1.Name == p2.Name },
			expected: false,
		},
		{
			name:     "different lengths",
			a:        []Person{{"Alice", 30}, {"Bob", 25}},
			b:        []Person{{"Alice", 30}},
			eq:       func(p1, p2 Person) bool { return p1.Name == p2.Name },
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.EqualFunc(tt.a, tt.b, tt.eq)
			if result != tt.expected {
				t.Errorf("EqualFunc(%v, %v) = %v; want %v",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestStartsWith_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		prefix   []int
		expected bool
	}{
		{
			name:     "starts with prefix",
			slice:    []int{1, 2, 3, 4},
			prefix:   []int{1, 2},
			expected: true,
		},
		{
			name:     "does not start with prefix",
			slice:    []int{1, 2, 3, 4},
			prefix:   []int{2, 3},
			expected: false,
		},
		{
			name:     "exact match",
			slice:    []int{1, 2, 3},
			prefix:   []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "empty prefix",
			slice:    []int{1, 2, 3},
			prefix:   []int{},
			expected: true,
		},
		{
			name:     "nil prefix",
			slice:    []int{1, 2, 3},
			prefix:   nil,
			expected: true,
		},
		{
			name:     "prefix longer than slice",
			slice:    []int{1, 2},
			prefix:   []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "nil slice with non-empty prefix",
			slice:    nil,
			prefix:   []int{1, 2},
			expected: false,
		},
		{
			name:     "nil slice with empty prefix",
			slice:    nil,
			prefix:   []int{},
			expected: true,
		},
		{
			name:     "empty slice with empty prefix",
			slice:    []int{},
			prefix:   []int{},
			expected: true,
		},
		{
			name:     "single element match",
			slice:    []int{1, 2, 3},
			prefix:   []int{1},
			expected: true,
		},
		{
			name:     "single element no match",
			slice:    []int{1, 2, 3},
			prefix:   []int{2},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.StartsWith(tt.slice, tt.prefix)
			if result != tt.expected {
				t.Errorf("StartsWith(%v, %v) = %v; want %v",
					tt.slice, tt.prefix, result, tt.expected)
			}
		})
	}
}

func TestStartsWith_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		prefix   []string
		expected bool
	}{
		{
			name:     "starts with prefix",
			slice:    []string{"a", "b", "c", "d"},
			prefix:   []string{"a", "b"},
			expected: true,
		},
		{
			name:     "does not start with prefix",
			slice:    []string{"a", "b", "c", "d"},
			prefix:   []string{"b", "c"},
			expected: false,
		},
		{
			name:     "exact match",
			slice:    []string{"a", "b", "c"},
			prefix:   []string{"a", "b", "c"},
			expected: true,
		},
		{
			name:     "empty prefix",
			slice:    []string{"a", "b", "c"},
			prefix:   []string{},
			expected: true,
		},
		{
			name:     "nil prefix",
			slice:    []string{"a", "b", "c"},
			prefix:   nil,
			expected: true,
		},
		{
			name:     "prefix longer than slice",
			slice:    []string{"a", "b"},
			prefix:   []string{"a", "b", "c"},
			expected: false,
		},
		{
			name:     "nil slice with non-empty prefix",
			slice:    nil,
			prefix:   []string{"a", "b"},
			expected: false,
		},
		{
			name:     "nil slice with empty prefix",
			slice:    nil,
			prefix:   []string{},
			expected: true,
		},
		{
			name:     "empty slice with empty prefix",
			slice:    []string{},
			prefix:   []string{},
			expected: true,
		},
		{
			name:     "single element match",
			slice:    []string{"a", "b", "c"},
			prefix:   []string{"a"},
			expected: true,
		},
		{
			name:     "single element no match",
			slice:    []string{"a", "b", "c"},
			prefix:   []string{"b"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.StartsWith(tt.slice, tt.prefix)
			if result != tt.expected {
				t.Errorf("StartsWith(%v, %v) = %v; want %v",
					tt.slice, tt.prefix, result, tt.expected)
			}
		})
	}
}

func TestEndsWith_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		suffix   []int
		expected bool
	}{
		{
			name:     "ends with suffix",
			slice:    []int{1, 2, 3, 4},
			suffix:   []int{3, 4},
			expected: true,
		},
		{
			name:     "does not end with suffix",
			slice:    []int{1, 2, 3, 4},
			suffix:   []int{2, 3},
			expected: false,
		},
		{
			name:     "exact match",
			slice:    []int{1, 2, 3},
			suffix:   []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "empty suffix",
			slice:    []int{1, 2, 3},
			suffix:   []int{},
			expected: true,
		},
		{
			name:     "nil suffix",
			slice:    []int{1, 2, 3},
			suffix:   nil,
			expected: true,
		},
		{
			name:     "suffix longer than slice",
			slice:    []int{3, 4},
			suffix:   []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "nil slice with non-empty suffix",
			slice:    nil,
			suffix:   []int{3, 4},
			expected: false,
		},
		{
			name:     "nil slice with empty suffix",
			slice:    nil,
			suffix:   []int{},
			expected: true,
		},
		{
			name:     "empty slice with empty suffix",
			slice:    []int{},
			suffix:   []int{},
			expected: true,
		},
		{
			name:     "single element match",
			slice:    []int{1, 2, 3},
			suffix:   []int{3},
			expected: true,
		},
		{
			name:     "single element no match",
			slice:    []int{1, 2, 3},
			suffix:   []int{2},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.EndsWith(tt.slice, tt.suffix)
			if result != tt.expected {
				t.Errorf("EndsWith(%v, %v) = %v; want %v",
					tt.slice, tt.suffix, result, tt.expected)
			}
		})
	}
}

func TestEndsWith_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		suffix   []string
		expected bool
	}{
		{
			name:     "ends with suffix",
			slice:    []string{"a", "b", "c", "d"},
			suffix:   []string{"c", "d"},
			expected: true,
		},
		{
			name:     "does not end with suffix",
			slice:    []string{"a", "b", "c", "d"},
			suffix:   []string{"b", "c"},
			expected: false,
		},
		{
			name:     "exact match",
			slice:    []string{"a", "b", "c"},
			suffix:   []string{"a", "b", "c"},
			expected: true,
		},
		{
			name:     "empty suffix",
			slice:    []string{"a", "b", "c"},
			suffix:   []string{},
			expected: true,
		},
		{
			name:     "nil suffix",
			slice:    []string{"a", "b", "c"},
			suffix:   nil,
			expected: true,
		},
		{
			name:     "suffix longer than slice",
			slice:    []string{"c", "d"},
			suffix:   []string{"a", "b", "c", "d"},
			expected: false,
		},
		{
			name:     "nil slice with non-empty suffix",
			slice:    nil,
			suffix:   []string{"c", "d"},
			expected: false,
		},
		{
			name:     "nil slice with empty suffix",
			slice:    nil,
			suffix:   []string{},
			expected: true,
		},
		{
			name:     "empty slice with empty suffix",
			slice:    []string{},
			suffix:   []string{},
			expected: true,
		},
		{
			name:     "single element match",
			slice:    []string{"a", "b", "c"},
			suffix:   []string{"c"},
			expected: true,
		},
		{
			name:     "single element no match",
			slice:    []string{"a", "b", "c"},
			suffix:   []string{"b"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.EndsWith(tt.slice, tt.suffix)
			if result != tt.expected {
				t.Errorf("EndsWith(%v, %v) = %v; want %v",
					tt.slice, tt.suffix, result, tt.expected)
			}
		})
	}
}

func TestEndsWith_Struct(t *testing.T) {
	type Product struct {
		ID   int
		Name string
	}

	tests := []struct {
		name     string
		slice    []Product
		suffix   []Product
		expected bool
	}{
		{
			name:     "ends with suffix",
			slice:    []Product{{1, "Apple"}, {2, "Banana"}, {3, "Cherry"}, {4, "Date"}},
			suffix:   []Product{{3, "Cherry"}, {4, "Date"}},
			expected: true,
		},
		{
			name:     "does not end with suffix",
			slice:    []Product{{1, "Apple"}, {2, "Banana"}, {3, "Cherry"}, {4, "Date"}},
			suffix:   []Product{{2, "Banana"}, {3, "Cherry"}},
			expected: false,
		},
		{
			name:     "exact match",
			slice:    []Product{{1, "Apple"}, {2, "Banana"}, {3, "Cherry"}},
			suffix:   []Product{{1, "Apple"}, {2, "Banana"}, {3, "Cherry"}},
			expected: true,
		},
		{
			name:     "empty suffix",
			slice:    []Product{{1, "Apple"}, {2, "Banana"}, {3, "Cherry"}},
			suffix:   []Product{},
			expected: true,
		},
		{
			name:     "nil suffix",
			slice:    []Product{{1, "Apple"}, {2, "Banana"}, {3, "Cherry"}},
			suffix:   nil,
			expected: true,
		},
		{
			name:     "suffix longer than slice",
			slice:    []Product{{3, "Cherry"}, {4, "Date"}},
			suffix:   []Product{{1, "Apple"}, {2, "Banana"}, {3, "Cherry"}, {4, "Date"}},
			expected: false,
		},
		{
			name:     "nil slice with non-empty suffix",
			slice:    nil,
			suffix:   []Product{{3, "Cherry"}, {4, "Date"}},
			expected: false,
		},
		{
			name:     "nil slice with empty suffix",
			slice:    nil,
			suffix:   []Product{},
			expected: true,
		},
		{
			name:     "empty slice with empty suffix",
			slice:    []Product{},
			suffix:   []Product{},
			expected: true,
		},
		{
			name:     "single element match",
			slice:    []Product{{1, "Apple"}, {2, "Banana"}, {3, "Cherry"}},
			suffix:   []Product{{3, "Cherry"}},
			expected: true,
		},
		{
			name:     "single element no match",
			slice:    []Product{{1, "Apple"}, {2, "Banana"}, {3, "Cherry"}},
			suffix:   []Product{{2, "Banana"}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.EndsWith(tt.slice, tt.suffix)
			if result != tt.expected {
				t.Errorf("EndsWith(%v, %v) = %v; want %v",
					tt.slice, tt.suffix, result, tt.expected)
			}
		})
	}
}

func TestHasPrefix_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		prefix   []int
		expected bool
	}{
		{
			name:     "has prefix",
			slice:    []int{1, 2, 3, 4},
			prefix:   []int{1, 2},
			expected: true,
		},
		{
			name:     "does not have prefix",
			slice:    []int{1, 2, 3, 4},
			prefix:   []int{2, 3},
			expected: false,
		},
		{
			name:     "exact match",
			slice:    []int{1, 2, 3},
			prefix:   []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "empty prefix",
			slice:    []int{1, 2, 3},
			prefix:   []int{},
			expected: true,
		},
		{
			name:     "nil prefix",
			slice:    []int{1, 2, 3},
			prefix:   nil,
			expected: true,
		},
		{
			name:     "prefix longer than slice",
			slice:    []int{1, 2},
			prefix:   []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "nil slice with non-empty prefix",
			slice:    nil,
			prefix:   []int{1, 2},
			expected: false,
		},
		{
			name:     "nil slice with empty prefix",
			slice:    nil,
			prefix:   []int{},
			expected: true,
		},
		{
			name:     "empty slice with empty prefix",
			slice:    []int{},
			prefix:   []int{},
			expected: true,
		},
		{
			name:     "single element match",
			slice:    []int{1, 2, 3},
			prefix:   []int{1},
			expected: true,
		},
		{
			name:     "single element no match",
			slice:    []int{1, 2, 3},
			prefix:   []int{2},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.HasPrefix(tt.slice, tt.prefix)
			if result != tt.expected {
				t.Errorf("HasPrefix(%v, %v) = %v; want %v",
					tt.slice, tt.prefix, result, tt.expected)
			}
			// Verify HasPrefix is consistent with StartsWith
			startsWith := lxslices.StartsWith(tt.slice, tt.prefix)
			if result != startsWith {
				t.Errorf("HasPrefix and StartsWith gave different results: %v vs %v",
					result, startsWith)
			}
		})
	}
}

func TestHasPrefix_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		prefix   []string
		expected bool
	}{
		{
			name:     "has prefix",
			slice:    []string{"a", "b", "c", "d"},
			prefix:   []string{"a", "b"},
			expected: true,
		},
		{
			name:     "does not have prefix",
			slice:    []string{"a", "b", "c", "d"},
			prefix:   []string{"b", "c"},
			expected: false,
		},
		{
			name:     "exact match",
			slice:    []string{"a", "b", "c"},
			prefix:   []string{"a", "b", "c"},
			expected: true,
		},
		{
			name:     "empty prefix",
			slice:    []string{"a", "b", "c"},
			prefix:   []string{},
			expected: true,
		},
		{
			name:     "nil prefix",
			slice:    []string{"a", "b", "c"},
			prefix:   nil,
			expected: true,
		},
		{
			name:     "prefix longer than slice",
			slice:    []string{"a", "b"},
			prefix:   []string{"a", "b", "c"},
			expected: false,
		},
		{
			name:     "nil slice with non-empty prefix",
			slice:    nil,
			prefix:   []string{"a", "b"},
			expected: false,
		},
		{
			name:     "nil slice with empty prefix",
			slice:    nil,
			prefix:   []string{},
			expected: true,
		},
		{
			name:     "empty slice with empty prefix",
			slice:    []string{},
			prefix:   []string{},
			expected: true,
		},
		{
			name:     "single element match",
			slice:    []string{"a", "b", "c"},
			prefix:   []string{"a"},
			expected: true,
		},
		{
			name:     "single element no match",
			slice:    []string{"a", "b", "c"},
			prefix:   []string{"b"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.HasPrefix(tt.slice, tt.prefix)
			if result != tt.expected {
				t.Errorf("HasPrefix(%v, %v) = %v; want %v",
					tt.slice, tt.prefix, result, tt.expected)
			}
			// Verify HasPrefix is consistent with StartsWith
			startsWith := lxslices.StartsWith(tt.slice, tt.prefix)
			if result != startsWith {
				t.Errorf("HasPrefix and StartsWith gave different results: %v vs %v",
					result, startsWith)
			}
		})
	}
}

func TestHasPrefix_Struct(t *testing.T) {
	type Item struct {
		Code string
		Qty  int
	}

	tests := []struct {
		name     string
		slice    []Item
		prefix   []Item
		expected bool
	}{
		{
			name:     "has prefix",
			slice:    []Item{{"A", 10}, {"B", 20}, {"C", 30}},
			prefix:   []Item{{"A", 10}},
			expected: true,
		},
		{
			name:     "does not have prefix",
			slice:    []Item{{"A", 10}, {"B", 20}, {"C", 30}},
			prefix:   []Item{{"B", 20}},
			expected: false,
		},
		{
			name:     "empty prefix",
			slice:    []Item{{"A", 10}},
			prefix:   []Item{},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.HasPrefix(tt.slice, tt.prefix)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
			// Verify HasPrefix is consistent with StartsWith
			startsWith := lxslices.StartsWith(tt.slice, tt.prefix)
			if result != startsWith {
				t.Errorf("HasPrefix and StartsWith gave different results: %v vs %v",
					result, startsWith)
			}
		})
	}
}

func TestHasSuffix_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		suffix   []int
		expected bool
	}{
		{
			name:     "has suffix",
			slice:    []int{1, 2, 3, 4},
			suffix:   []int{3, 4},
			expected: true,
		},
		{
			name:     "does not have suffix",
			slice:    []int{1, 2, 3, 4},
			suffix:   []int{2, 3},
			expected: false,
		},
		{
			name:     "exact match",
			slice:    []int{1, 2, 3},
			suffix:   []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "empty suffix",
			slice:    []int{1, 2, 3},
			suffix:   []int{},
			expected: true,
		},
		{
			name:     "nil suffix",
			slice:    []int{1, 2, 3},
			suffix:   nil,
			expected: true,
		},
		{
			name:     "suffix longer than slice",
			slice:    []int{3, 4},
			suffix:   []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "nil slice with non-empty suffix",
			slice:    nil,
			suffix:   []int{3, 4},
			expected: false,
		},
		{
			name:     "nil slice with empty suffix",
			slice:    nil,
			suffix:   []int{},
			expected: true,
		},
		{
			name:     "empty slice with empty suffix",
			slice:    []int{},
			suffix:   []int{},
			expected: true,
		},
		{
			name:     "single element match",
			slice:    []int{1, 2, 3},
			suffix:   []int{3},
			expected: true,
		},
		{
			name:     "single element no match",
			slice:    []int{1, 2, 3},
			suffix:   []int{2},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.HasSuffix(tt.slice, tt.suffix)
			if result != tt.expected {
				t.Errorf("HasSuffix(%v, %v) = %v; want %v",
					tt.slice, tt.suffix, result, tt.expected)
			}
			// Verify HasSuffix is consistent with EndsWith
			endsWith := lxslices.EndsWith(tt.slice, tt.suffix)
			if result != endsWith {
				t.Errorf("HasSuffix and EndsWith gave different results: %v vs %v",
					result, endsWith)
			}
		})
	}
}

func TestHasSuffix_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		suffix   []string
		expected bool
	}{
		{
			name:     "has suffix",
			slice:    []string{"a", "b", "c", "d"},
			suffix:   []string{"c", "d"},
			expected: true,
		},
		{
			name:     "does not have suffix",
			slice:    []string{"a", "b", "c", "d"},
			suffix:   []string{"b", "c"},
			expected: false,
		},
		{
			name:     "exact match",
			slice:    []string{"a", "b", "c"},
			suffix:   []string{"a", "b", "c"},
			expected: true,
		},
		{
			name:     "empty suffix",
			slice:    []string{"a", "b", "c"},
			suffix:   []string{},
			expected: true,
		},
		{
			name:     "nil suffix",
			slice:    []string{"a", "b", "c"},
			suffix:   nil,
			expected: true,
		},
		{
			name:     "suffix longer than slice",
			slice:    []string{"c", "d"},
			suffix:   []string{"a", "b", "c", "d"},
			expected: false,
		},
		{
			name:     "nil slice with non-empty suffix",
			slice:    nil,
			suffix:   []string{"c", "d"},
			expected: false,
		},
		{
			name:     "nil slice with empty suffix",
			slice:    nil,
			suffix:   []string{},
			expected: true,
		},
		{
			name:     "empty slice with empty suffix",
			slice:    []string{},
			suffix:   []string{},
			expected: true,
		},
		{
			name:     "single element match",
			slice:    []string{"a", "b", "c"},
			suffix:   []string{"c"},
			expected: true,
		},
		{
			name:     "single element no match",
			slice:    []string{"a", "b", "c"},
			suffix:   []string{"b"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.HasSuffix(tt.slice, tt.suffix)
			if result != tt.expected {
				t.Errorf("HasSuffix(%v, %v) = %v; want %v",
					tt.slice, tt.suffix, result, tt.expected)
			}
			// Verify HasSuffix is consistent with EndsWith
			endsWith := lxslices.EndsWith(tt.slice, tt.suffix)
			if result != endsWith {
				t.Errorf("HasSuffix and EndsWith gave different results: %v vs %v",
					result, endsWith)
			}
		})
	}
}

func TestHasSuffix_Struct(t *testing.T) {
	type Item struct {
		Code string
		Qty  int
	}

	tests := []struct {
		name     string
		slice    []Item
		suffix   []Item
		expected bool
	}{
		{
			name:     "has suffix",
			slice:    []Item{{"A", 10}, {"B", 20}, {"C", 30}, {"D", 40}},
			suffix:   []Item{{"C", 30}, {"D", 40}},
			expected: true,
		},
		{
			name:     "does not have suffix",
			slice:    []Item{{"A", 10}, {"B", 20}, {"C", 30}, {"D", 40}},
			suffix:   []Item{{"B", 20}, {"C", 30}},
			expected: false,
		},
		{
			name:     "exact match",
			slice:    []Item{{"A", 10}, {"B", 20}, {"C", 30}},
			suffix:   []Item{{"A", 10}, {"B", 20}, {"C", 30}},
			expected: true,
		},
		{
			name:     "empty suffix",
			slice:    []Item{{"A", 10}, {"B", 20}, {"C", 30}},
			suffix:   []Item{},
			expected: true,
		},
		{
			name:     "nil suffix",
			slice:    []Item{{"A", 10}, {"B", 20}, {"C", 30}},
			suffix:   nil,
			expected: true,
		},
		{
			name:     "suffix longer than slice",
			slice:    []Item{{"C", 30}, {"D", 40}},
			suffix:   []Item{{"A", 10}, {"B", 20}, {"C", 30}, {"D", 40}},
			expected: false,
		},
		{
			name:     "nil slice with non-empty suffix",
			slice:    nil,
			suffix:   []Item{{"C", 30}, {"D", 40}},
			expected: false,
		},
		{
			name:     "nil slice with empty suffix",
			slice:    nil,
			suffix:   []Item{},
			expected: true,
		},
		{
			name:     "empty slice with empty suffix",
			slice:    []Item{},
			suffix:   []Item{},
			expected: true,
		},
		{
			name:     "single element match",
			slice:    []Item{{"A", 10}, {"B", 20}, {"C", 30}},
			suffix:   []Item{{"C", 30}},
			expected: true,
		},
		{
			name:     "single element no match",
			slice:    []Item{{"A", 10}, {"B", 20}, {"C", 30}},
			suffix:   []Item{{"B", 20}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.HasSuffix(tt.slice, tt.suffix)
			if result != tt.expected {
				t.Errorf("HasSuffix(%v, %v) = %v; want %v",
					tt.slice, tt.suffix, result, tt.expected)
			}
			// Verify HasSuffix is consistent with EndsWith
			endsWith := lxslices.EndsWith(tt.slice, tt.suffix)
			if result != endsWith {
				t.Errorf("HasSuffix and EndsWith gave different results: %v vs %v",
					result, endsWith)
			}
		})
	}
}

// Helper function for absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
