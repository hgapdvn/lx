package lxslices_test

import (
	"reflect"
	"testing"

	"github.com/hgapdvn/lx/lxslices"
	"github.com/hgapdvn/lx/lxtypes"
)

func TestTake_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		n        int
		expected []int
	}{
		{
			name:     "take first 3 elements",
			slice:    []int{1, 2, 3, 4, 5},
			n:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "take more than length",
			slice:    []int{1, 2},
			n:        5,
			expected: []int{1, 2},
		},
		{
			name:     "take exact length",
			slice:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "take 0 elements",
			slice:    []int{1, 2, 3},
			n:        0,
			expected: []int{},
		},
		{
			name:     "take negative count",
			slice:    []int{1, 2, 3},
			n:        -1,
			expected: []int{},
		},
		{
			name:     "take from empty slice",
			slice:    []int{},
			n:        3,
			expected: []int{},
		},
		{
			name:     "take from nil slice",
			slice:    nil,
			n:        3,
			expected: nil,
		},
		{
			name:     "take single element",
			slice:    []int{1, 2, 3, 4, 5},
			n:        1,
			expected: []int{1},
		},
		{
			name:     "take all but one",
			slice:    []int{1, 2, 3, 4, 5},
			n:        4,
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Take(tt.slice, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Take(%v, %d) = %v; want %v",
					tt.slice, tt.n, result, tt.expected)
			}
		})
	}
}

func TestTake_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		n        int
		expected []string
	}{
		{
			name:     "take first 2 elements",
			slice:    []string{"a", "b", "c", "d"},
			n:        2,
			expected: []string{"a", "b"},
		},
		{
			name:     "take more than length",
			slice:    []string{"x", "y"},
			n:        10,
			expected: []string{"x", "y"},
		},
		{
			name:     "take exact length",
			slice:    []string{"a", "b", "c"},
			n:        3,
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "take 0 elements",
			slice:    []string{"a", "b"},
			n:        0,
			expected: []string{},
		},
		{
			name:     "take negative count",
			slice:    []string{"a", "b"},
			n:        -5,
			expected: []string{},
		},
		{
			name:     "take from empty slice",
			slice:    []string{},
			n:        2,
			expected: []string{},
		},
		{
			name:     "take from nil slice",
			slice:    nil,
			n:        2,
			expected: nil,
		},
		{
			name:     "take single element",
			slice:    []string{"hello", "world"},
			n:        1,
			expected: []string{"hello"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Take(tt.slice, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Take(%v, %d) = %v; want %v",
					tt.slice, tt.n, result, tt.expected)
			}
		})
	}
}

func TestTake_Struct(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	tests := []struct {
		name     string
		slice    []User
		n        int
		expected []User
	}{
		{
			name:     "take first 2 structs",
			slice:    []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}, {4, "Dave"}, {5, "Eve"}},
			n:        2,
			expected: []User{{1, "Alice"}, {2, "Bob"}},
		},
		{
			name:     "take more than length",
			slice:    []User{{1, "Alice"}, {2, "Bob"}},
			n:        5,
			expected: []User{{1, "Alice"}, {2, "Bob"}},
		},
		{
			name:     "take exact length",
			slice:    []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
			n:        3,
			expected: []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
		},
		{
			name:     "take 0 structs",
			slice:    []User{{1, "Alice"}, {2, "Bob"}},
			n:        0,
			expected: []User{},
		},
		{
			name:     "take negative count",
			slice:    []User{{1, "Alice"}, {2, "Bob"}},
			n:        -1,
			expected: []User{},
		},
		{
			name:     "take from empty slice",
			slice:    []User{},
			n:        2,
			expected: []User{},
		},
		{
			name:     "take from nil slice",
			slice:    nil,
			n:        2,
			expected: nil,
		},
		{
			name:     "take single struct",
			slice:    []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
			n:        1,
			expected: []User{{1, "Alice"}},
		},
		{
			name:     "take all but one",
			slice:    []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}, {4, "Dave"}},
			n:        3,
			expected: []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Take(tt.slice, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Take(%v, %d) = %v; want %v",
					tt.slice, tt.n, result, tt.expected)
			}
		})
	}
}

func TestTakeLast_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		n        int
		expected []int
	}{
		{
			name:     "take last 3 elements",
			slice:    []int{1, 2, 3, 4, 5},
			n:        3,
			expected: []int{3, 4, 5},
		},
		{
			name:     "take last more than length",
			slice:    []int{1, 2},
			n:        5,
			expected: []int{1, 2},
		},
		{
			name:     "take last exact length",
			slice:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "take last 0 elements",
			slice:    []int{1, 2, 3},
			n:        0,
			expected: []int{},
		},
		{
			name:     "take last negative count",
			slice:    []int{1, 2, 3},
			n:        -1,
			expected: []int{},
		},
		{
			name:     "take last from empty slice",
			slice:    []int{},
			n:        3,
			expected: []int{},
		},
		{
			name:     "take last from nil slice",
			slice:    nil,
			n:        3,
			expected: nil,
		},
		{
			name:     "take last single element",
			slice:    []int{1, 2, 3, 4, 5},
			n:        1,
			expected: []int{5},
		},
		{
			name:     "take last all but one",
			slice:    []int{1, 2, 3, 4, 5},
			n:        4,
			expected: []int{2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.TakeLast(tt.slice, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TakeLast(%v, %d) = %v; want %v",
					tt.slice, tt.n, result, tt.expected)
			}
		})
	}
}

func TestTakeLast_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		n        int
		expected []string
	}{
		{
			name:     "take last 2 elements",
			slice:    []string{"a", "b", "c", "d"},
			n:        2,
			expected: []string{"c", "d"},
		},
		{
			name:     "take last more than length",
			slice:    []string{"x", "y"},
			n:        10,
			expected: []string{"x", "y"},
		},
		{
			name:     "take last exact length",
			slice:    []string{"a", "b", "c"},
			n:        3,
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "take last 0 elements",
			slice:    []string{"a", "b"},
			n:        0,
			expected: []string{},
		},
		{
			name:     "take last negative count",
			slice:    []string{"a", "b"},
			n:        -3,
			expected: []string{},
		},
		{
			name:     "take last from empty slice",
			slice:    []string{},
			n:        2,
			expected: []string{},
		},
		{
			name:     "take last from nil slice",
			slice:    nil,
			n:        2,
			expected: nil,
		},
		{
			name:     "take last single element",
			slice:    []string{"hello", "world"},
			n:        1,
			expected: []string{"world"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.TakeLast(tt.slice, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TakeLast(%v, %d) = %v; want %v",
					tt.slice, tt.n, result, tt.expected)
			}
		})
	}
}

func TestTakeLast_Struct(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	tests := []struct {
		name     string
		slice    []User
		n        int
		expected []User
	}{
		{
			name:     "take last 2 structs",
			slice:    []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}, {4, "Dave"}, {5, "Eve"}},
			n:        2,
			expected: []User{{4, "Dave"}, {5, "Eve"}},
		},
		{
			name:     "take last more than length",
			slice:    []User{{1, "Alice"}, {2, "Bob"}},
			n:        5,
			expected: []User{{1, "Alice"}, {2, "Bob"}},
		},
		{
			name:     "take last exact length",
			slice:    []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
			n:        3,
			expected: []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
		},
		{
			name:     "take last 0 structs",
			slice:    []User{{1, "Alice"}, {2, "Bob"}},
			n:        0,
			expected: []User{},
		},
		{
			name:     "take last negative count",
			slice:    []User{{1, "Alice"}, {2, "Bob"}},
			n:        -1,
			expected: []User{},
		},
		{
			name:     "take last from empty slice",
			slice:    []User{},
			n:        2,
			expected: []User{},
		},
		{
			name:     "take last from nil slice",
			slice:    nil,
			n:        2,
			expected: nil,
		},
		{
			name:     "take last single struct",
			slice:    []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
			n:        1,
			expected: []User{{3, "Charlie"}},
		},
		{
			name:     "take last all but one",
			slice:    []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}, {4, "Dave"}},
			n:        3,
			expected: []User{{2, "Bob"}, {3, "Charlie"}, {4, "Dave"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.TakeLast(tt.slice, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TakeLast(%v, %d) = %v; want %v",
					tt.slice, tt.n, result, tt.expected)
			}
		})
	}
}

func TestTakeWhile_Int(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate lxtypes.Predicate[int]
		expected  []int
	}{
		{
			name:      "take while even",
			slice:     []int{2, 4, 6, 7, 8},
			predicate: func(n int) bool { return n%2 == 0 },
			expected:  []int{2, 4, 6},
		},
		{
			name:      "take while less than 5",
			slice:     []int{1, 2, 3, 5, 6},
			predicate: func(n int) bool { return n < 5 },
			expected:  []int{1, 2, 3},
		},
		{
			name:      "predicate false on first element",
			slice:     []int{10, 2, 3},
			predicate: func(n int) bool { return n < 5 },
			expected:  []int{},
		},
		{
			name:      "predicate true for all",
			slice:     []int{1, 2, 3},
			predicate: func(n int) bool { return n < 10 },
			expected:  []int{1, 2, 3},
		},
		{
			name:      "empty slice",
			slice:     []int{},
			predicate: func(n int) bool { return true },
			expected:  []int{},
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(n int) bool { return true },
			expected:  nil,
		},
		{
			name:      "take while positive",
			slice:     []int{1, 2, 3, -1, 4},
			predicate: func(n int) bool { return n > 0 },
			expected:  []int{1, 2, 3},
		},
		{
			name:      "take while greater than 10",
			slice:     []int{15, 20, 25, 5, 30},
			predicate: func(n int) bool { return n > 10 },
			expected:  []int{15, 20, 25},
		},
		{
			name:      "single element matches",
			slice:     []int{5},
			predicate: func(n int) bool { return n < 10 },
			expected:  []int{5},
		},
		{
			name:      "single element doesn't match",
			slice:     []int{15},
			predicate: func(n int) bool { return n < 10 },
			expected:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.TakeWhile(tt.slice, tt.predicate)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TakeWhile(%v) = %v; want %v",
					tt.slice, result, tt.expected)
			}
		})
	}
}

func TestTakeWhile_String(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		predicate lxtypes.Predicate[string]
		expected  []string
	}{
		{
			name:      "take while starts with 'a'",
			slice:     []string{"apple", "avocado", "banana", "apricot"},
			predicate: func(s string) bool { return len(s) > 0 && s[0] == 'a' },
			expected:  []string{"apple", "avocado"},
		},
		{
			name:      "predicate false on first",
			slice:     []string{"banana", "apple"},
			predicate: func(s string) bool { return len(s) > 0 && s[0] == 'a' },
			expected:  []string{},
		},
		{
			name:      "predicate true for all",
			slice:     []string{"a", "b", "c"},
			predicate: func(s string) bool { return len(s) == 1 },
			expected:  []string{"a", "b", "c"},
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(s string) bool { return true },
			expected:  nil,
		},
		{
			name:      "take while length less than 5",
			slice:     []string{"cat", "dog", "elephant", "ant"},
			predicate: func(s string) bool { return len(s) < 5 },
			expected:  []string{"cat", "dog"},
		},
		{
			name:      "empty slice",
			slice:     []string{},
			predicate: func(s string) bool { return true },
			expected:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.TakeWhile(tt.slice, tt.predicate)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TakeWhile(%v) = %v; want %v",
					tt.slice, result, tt.expected)
			}
		})
	}
}

func TestTakeWhile_Struct(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	tests := []struct {
		name      string
		slice     []User
		predicate lxtypes.Predicate[User]
		expected  []User
	}{
		{
			name:      "take while ID less than 3",
			slice:     []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}, {4, "Dave"}},
			predicate: func(u User) bool { return u.ID < 3 },
			expected:  []User{{1, "Alice"}, {2, "Bob"}},
		},
		{
			name:      "take while ID even",
			slice:     []User{{2, "Alice"}, {4, "Bob"}, {5, "Charlie"}, {6, "Dave"}},
			predicate: func(u User) bool { return u.ID%2 == 0 },
			expected:  []User{{2, "Alice"}, {4, "Bob"}},
		},
		{
			name:      "predicate false on first",
			slice:     []User{{5, "Alice"}, {2, "Bob"}},
			predicate: func(u User) bool { return u.ID < 3 },
			expected:  []User{},
		},
		{
			name:      "predicate true for all",
			slice:     []User{{1, "Alice"}, {2, "Bob"}},
			predicate: func(u User) bool { return u.ID < 10 },
			expected:  []User{{1, "Alice"}, {2, "Bob"}},
		},
		{
			name:      "empty slice",
			slice:     []User{},
			predicate: func(u User) bool { return true },
			expected:  []User{},
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(u User) bool { return true },
			expected:  nil,
		},
		{
			name:      "take while name length less than 5",
			slice:     []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
			predicate: func(u User) bool { return len(u.Name) < 5 },
			expected:  []User{},
		},
		{
			name:      "take while name starts with A",
			slice:     []User{{1, "Alice"}, {2, "Amy"}, {3, "Bob"}},
			predicate: func(u User) bool { return len(u.Name) > 0 && u.Name[0] == 'A' },
			expected:  []User{{1, "Alice"}, {2, "Amy"}},
		},
		{
			name:      "single element matches",
			slice:     []User{{1, "Alice"}},
			predicate: func(u User) bool { return u.ID < 10 },
			expected:  []User{{1, "Alice"}},
		},
		{
			name:      "single element doesn't match",
			slice:     []User{{15, "Alice"}},
			predicate: func(u User) bool { return u.ID < 10 },
			expected:  []User{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.TakeWhile(tt.slice, tt.predicate)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TakeWhile(%v) = %v; want %v",
					tt.slice, result, tt.expected)
			}
		})
	}
}

func TestDrop_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		n        int
		expected []int
	}{
		{
			name:     "drop first 2 elements",
			slice:    []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{3, 4, 5},
		},
		{
			name:     "drop more than length",
			slice:    []int{1, 2},
			n:        5,
			expected: []int{},
		},
		{
			name:     "drop exact length",
			slice:    []int{1, 2, 3},
			n:        3,
			expected: []int{},
		},
		{
			name:     "drop 0 elements",
			slice:    []int{1, 2, 3},
			n:        0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "drop negative count",
			slice:    []int{1, 2, 3},
			n:        -1,
			expected: []int{1, 2, 3},
		},
		{
			name:     "drop from empty slice",
			slice:    []int{},
			n:        3,
			expected: []int{},
		},
		{
			name:     "drop from nil slice",
			slice:    nil,
			n:        3,
			expected: nil,
		},
		{
			name:     "drop single element",
			slice:    []int{1, 2, 3, 4, 5},
			n:        1,
			expected: []int{2, 3, 4, 5},
		},
		{
			name:     "drop all but one",
			slice:    []int{1, 2, 3, 4, 5},
			n:        4,
			expected: []int{5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Drop(tt.slice, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Drop(%v, %d) = %v; want %v",
					tt.slice, tt.n, result, tt.expected)
			}
		})
	}
}

func TestDrop_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		n        int
		expected []string
	}{
		{
			name:     "drop first 2 elements",
			slice:    []string{"a", "b", "c", "d"},
			n:        2,
			expected: []string{"c", "d"},
		},
		{
			name:     "drop more than length",
			slice:    []string{"x", "y"},
			n:        10,
			expected: []string{},
		},
		{
			name:     "drop exact length",
			slice:    []string{"a", "b", "c"},
			n:        3,
			expected: []string{},
		},
		{
			name:     "drop 0 elements",
			slice:    []string{"a", "b"},
			n:        0,
			expected: []string{"a", "b"},
		},
		{
			name:     "drop negative count",
			slice:    []string{"a", "b"},
			n:        -2,
			expected: []string{"a", "b"},
		},
		{
			name:     "drop from empty slice",
			slice:    []string{},
			n:        2,
			expected: []string{},
		},
		{
			name:     "drop from nil slice",
			slice:    nil,
			n:        2,
			expected: nil,
		},
		{
			name:     "drop single element",
			slice:    []string{"hello", "world"},
			n:        1,
			expected: []string{"world"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Drop(tt.slice, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Drop(%v, %d) = %v; want %v",
					tt.slice, tt.n, result, tt.expected)
			}
		})
	}
}

func TestDrop_Struct(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	tests := []struct {
		name     string
		slice    []User
		n        int
		expected []User
	}{
		{
			name:     "drop first 2 structs",
			slice:    []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}, {4, "Dave"}, {5, "Eve"}},
			n:        2,
			expected: []User{{3, "Charlie"}, {4, "Dave"}, {5, "Eve"}},
		},
		{
			name:     "drop more than length",
			slice:    []User{{1, "Alice"}, {2, "Bob"}},
			n:        5,
			expected: []User{},
		},
		{
			name:     "drop exact length",
			slice:    []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
			n:        3,
			expected: []User{},
		},
		{
			name:     "drop 0 structs",
			slice:    []User{{1, "Alice"}, {2, "Bob"}},
			n:        0,
			expected: []User{{1, "Alice"}, {2, "Bob"}},
		},
		{
			name:     "drop negative count",
			slice:    []User{{1, "Alice"}, {2, "Bob"}},
			n:        -1,
			expected: []User{{1, "Alice"}, {2, "Bob"}},
		},
		{
			name:     "drop from empty slice",
			slice:    []User{},
			n:        2,
			expected: []User{},
		},
		{
			name:     "drop from nil slice",
			slice:    nil,
			n:        2,
			expected: nil,
		},
		{
			name:     "drop single struct",
			slice:    []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
			n:        1,
			expected: []User{{2, "Bob"}, {3, "Charlie"}},
		},
		{
			name:     "drop all but one",
			slice:    []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}, {4, "Dave"}},
			n:        3,
			expected: []User{{4, "Dave"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Drop(tt.slice, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Drop(%v, %d) = %v; want %v",
					tt.slice, tt.n, result, tt.expected)
			}
		})
	}
}

func TestDropLast_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		n        int
		expected []int
	}{
		{
			name:     "drop last 2 elements",
			slice:    []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{1, 2, 3},
		},
		{
			name:     "drop last more than length",
			slice:    []int{1, 2},
			n:        5,
			expected: []int{},
		},
		{
			name:     "drop last exact length",
			slice:    []int{1, 2, 3},
			n:        3,
			expected: []int{},
		},
		{
			name:     "drop last 0 elements",
			slice:    []int{1, 2, 3},
			n:        0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "drop last negative count",
			slice:    []int{1, 2, 3},
			n:        -1,
			expected: []int{1, 2, 3},
		},
		{
			name:     "drop last from empty slice",
			slice:    []int{},
			n:        3,
			expected: []int{},
		},
		{
			name:     "drop last from nil slice",
			slice:    nil,
			n:        3,
			expected: nil,
		},
		{
			name:     "drop last single element",
			slice:    []int{1, 2, 3, 4, 5},
			n:        1,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "drop last all but one",
			slice:    []int{1, 2, 3, 4, 5},
			n:        4,
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.DropLast(tt.slice, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("DropLast(%v, %d) = %v; want %v",
					tt.slice, tt.n, result, tt.expected)
			}
		})
	}
}

func TestDropLast_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		n        int
		expected []string
	}{
		{
			name:     "drop last 2 elements",
			slice:    []string{"a", "b", "c", "d"},
			n:        2,
			expected: []string{"a", "b"},
		},
		{
			name:     "drop last more than length",
			slice:    []string{"x", "y"},
			n:        10,
			expected: []string{},
		},
		{
			name:     "drop last exact length",
			slice:    []string{"a", "b", "c"},
			n:        3,
			expected: []string{},
		},
		{
			name:     "drop last 0 elements",
			slice:    []string{"a", "b"},
			n:        0,
			expected: []string{"a", "b"},
		},
		{
			name:     "drop last negative count",
			slice:    []string{"a", "b"},
			n:        -3,
			expected: []string{"a", "b"},
		},
		{
			name:     "drop last from empty slice",
			slice:    []string{},
			n:        2,
			expected: []string{},
		},
		{
			name:     "drop last from nil slice",
			slice:    nil,
			n:        2,
			expected: nil,
		},
		{
			name:     "drop last single element",
			slice:    []string{"hello", "world"},
			n:        1,
			expected: []string{"hello"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.DropLast(tt.slice, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("DropLast(%v, %d) = %v; want %v",
					tt.slice, tt.n, result, tt.expected)
			}
		})
	}
}

func TestDropLast_Struct(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	tests := []struct {
		name     string
		slice    []User
		n        int
		expected []User
	}{
		{
			name:     "drop last 2 structs",
			slice:    []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}, {4, "Dave"}, {5, "Eve"}},
			n:        2,
			expected: []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
		},
		{
			name:     "drop last more than length",
			slice:    []User{{1, "Alice"}, {2, "Bob"}},
			n:        5,
			expected: []User{},
		},
		{
			name:     "drop last exact length",
			slice:    []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
			n:        3,
			expected: []User{},
		},
		{
			name:     "drop last 0 structs",
			slice:    []User{{1, "Alice"}, {2, "Bob"}},
			n:        0,
			expected: []User{{1, "Alice"}, {2, "Bob"}},
		},
		{
			name:     "drop last negative count",
			slice:    []User{{1, "Alice"}, {2, "Bob"}},
			n:        -1,
			expected: []User{{1, "Alice"}, {2, "Bob"}},
		},
		{
			name:     "drop last from empty slice",
			slice:    []User{},
			n:        2,
			expected: []User{},
		},
		{
			name:     "drop last from nil slice",
			slice:    nil,
			n:        2,
			expected: nil,
		},
		{
			name:     "drop last single struct",
			slice:    []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
			n:        1,
			expected: []User{{1, "Alice"}, {2, "Bob"}},
		},
		{
			name:     "drop last all but one",
			slice:    []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}, {4, "Dave"}},
			n:        3,
			expected: []User{{1, "Alice"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.DropLast(tt.slice, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("DropLast(%v, %d) = %v; want %v",
					tt.slice, tt.n, result, tt.expected)
			}
		})
	}
}

func TestDropWhile_Int(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate lxtypes.Predicate[int]
		expected  []int
	}{
		{
			name:      "drop while even",
			slice:     []int{2, 4, 6, 7, 8},
			predicate: func(n int) bool { return n%2 == 0 },
			expected:  []int{7, 8},
		},
		{
			name:      "drop while less than 5",
			slice:     []int{1, 2, 3, 5, 6},
			predicate: func(n int) bool { return n < 5 },
			expected:  []int{5, 6},
		},
		{
			name:      "predicate false on first element",
			slice:     []int{10, 2, 3},
			predicate: func(n int) bool { return n < 5 },
			expected:  []int{10, 2, 3},
		},
		{
			name:      "predicate true for all",
			slice:     []int{1, 2, 3},
			predicate: func(n int) bool { return n < 10 },
			expected:  []int{},
		},
		{
			name:      "empty slice",
			slice:     []int{},
			predicate: func(n int) bool { return true },
			expected:  []int{},
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(n int) bool { return true },
			expected:  nil,
		},
		{
			name:      "drop while positive",
			slice:     []int{1, 2, 3, -1, 4},
			predicate: func(n int) bool { return n > 0 },
			expected:  []int{-1, 4},
		},
		{
			name:      "drop while less than 10",
			slice:     []int{5, 7, 15, 20},
			predicate: func(n int) bool { return n < 10 },
			expected:  []int{15, 20},
		},
		{
			name:      "single element matches",
			slice:     []int{5},
			predicate: func(n int) bool { return n < 10 },
			expected:  []int{},
		},
		{
			name:      "single element doesn't match",
			slice:     []int{15},
			predicate: func(n int) bool { return n < 10 },
			expected:  []int{15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.DropWhile(tt.slice, tt.predicate)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("DropWhile(%v) = %v; want %v",
					tt.slice, result, tt.expected)
			}
		})
	}
}

func TestDropWhile_String(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		predicate lxtypes.Predicate[string]
		expected  []string
	}{
		{
			name:      "drop while starts with 'a'",
			slice:     []string{"apple", "avocado", "banana", "apricot"},
			predicate: func(s string) bool { return len(s) > 0 && s[0] == 'a' },
			expected:  []string{"banana", "apricot"},
		},
		{
			name:      "predicate false on first",
			slice:     []string{"banana", "apple"},
			predicate: func(s string) bool { return len(s) > 0 && s[0] == 'a' },
			expected:  []string{"banana", "apple"},
		},
		{
			name:      "predicate true for all",
			slice:     []string{"a", "b", "c"},
			predicate: func(s string) bool { return len(s) == 1 },
			expected:  []string{},
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(s string) bool { return true },
			expected:  nil,
		},
		{
			name:      "drop while length less than 5",
			slice:     []string{"cat", "dog", "elephant", "ant"},
			predicate: func(s string) bool { return len(s) < 5 },
			expected:  []string{"elephant", "ant"},
		},
		{
			name:      "empty slice",
			slice:     []string{},
			predicate: func(s string) bool { return true },
			expected:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.DropWhile(tt.slice, tt.predicate)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("DropWhile(%v) = %v; want %v",
					tt.slice, result, tt.expected)
			}
		})
	}
}

func TestDropWhile_Struct(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	tests := []struct {
		name      string
		slice     []User
		predicate lxtypes.Predicate[User]
		expected  []User
	}{
		{
			name:      "drop while ID less than 3",
			slice:     []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}, {4, "Dave"}},
			predicate: func(u User) bool { return u.ID < 3 },
			expected:  []User{{3, "Charlie"}, {4, "Dave"}},
		},
		{
			name:      "drop while ID even",
			slice:     []User{{2, "Alice"}, {4, "Bob"}, {5, "Charlie"}, {6, "Dave"}},
			predicate: func(u User) bool { return u.ID%2 == 0 },
			expected:  []User{{5, "Charlie"}, {6, "Dave"}},
		},
		{
			name:      "predicate false on first",
			slice:     []User{{5, "Alice"}, {2, "Bob"}},
			predicate: func(u User) bool { return u.ID < 3 },
			expected:  []User{{5, "Alice"}, {2, "Bob"}},
		},
		{
			name:      "predicate true for all",
			slice:     []User{{1, "Alice"}, {2, "Bob"}},
			predicate: func(u User) bool { return u.ID < 10 },
			expected:  []User{},
		},
		{
			name:      "empty slice",
			slice:     []User{},
			predicate: func(u User) bool { return true },
			expected:  []User{},
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(u User) bool { return true },
			expected:  nil,
		},
		{
			name:      "drop while name length less than 5",
			slice:     []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}},
			predicate: func(u User) bool { return len(u.Name) >= 5 },
			expected:  []User{{2, "Bob"}, {3, "Charlie"}},
		},
		{
			name:      "drop while name starts with A",
			slice:     []User{{1, "Alice"}, {2, "Amy"}, {3, "Bob"}},
			predicate: func(u User) bool { return len(u.Name) > 0 && u.Name[0] == 'A' },
			expected:  []User{{3, "Bob"}},
		},
		{
			name:      "single element matches",
			slice:     []User{{1, "Alice"}},
			predicate: func(u User) bool { return u.ID < 10 },
			expected:  []User{},
		},
		{
			name:      "single element doesn't match",
			slice:     []User{{15, "Alice"}},
			predicate: func(u User) bool { return u.ID < 10 },
			expected:  []User{{15, "Alice"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.DropWhile(tt.slice, tt.predicate)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("DropWhile(%v) = %v; want %v",
					tt.slice, result, tt.expected)
			}
		})
	}
}

// Test immutability - functions should not modify original slice
func TestTake_NoMutation(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	originalCopy := []int{1, 2, 3, 4, 5}
	result := lxslices.Take(original, 3)

	// Modify the result
	if len(result) > 0 {
		result[0] = 999
	}

	// Original should be unchanged
	if !reflect.DeepEqual(original, originalCopy) {
		t.Error("Take modified the original slice")
	}
}

func TestDrop_NoMutation(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	originalCopy := []int{1, 2, 3, 4, 5}
	result := lxslices.Drop(original, 2)

	// Modify the result
	if len(result) > 0 {
		result[0] = 999
	}

	// Original should be unchanged
	if !reflect.DeepEqual(original, originalCopy) {
		t.Error("Drop modified the original slice")
	}
}

func TestTakeLast_NoMutation(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	originalCopy := []int{1, 2, 3, 4, 5}
	result := lxslices.TakeLast(original, 3)

	// Modify the result
	if len(result) > 0 {
		result[0] = 999
	}

	// Original should be unchanged
	if !reflect.DeepEqual(original, originalCopy) {
		t.Error("TakeLast modified the original slice")
	}
}

func TestDropLast_NoMutation(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	originalCopy := []int{1, 2, 3, 4, 5}
	result := lxslices.DropLast(original, 2)

	// Modify the result
	if len(result) > 0 {
		result[0] = 999
	}

	// Original should be unchanged
	if !reflect.DeepEqual(original, originalCopy) {
		t.Error("DropLast modified the original slice")
	}
}

func TestTakeWhile_NoMutation(t *testing.T) {
	original := []int{2, 4, 6, 7, 8}
	originalCopy := []int{2, 4, 6, 7, 8}
	result := lxslices.TakeWhile(original, func(n int) bool { return n%2 == 0 })

	// Modify the result
	if len(result) > 0 {
		result[0] = 999
	}

	// Original should be unchanged
	if !reflect.DeepEqual(original, originalCopy) {
		t.Error("TakeWhile modified the original slice")
	}
}

func TestDropWhile_NoMutation(t *testing.T) {
	original := []int{2, 4, 6, 7, 8}
	originalCopy := []int{2, 4, 6, 7, 8}
	result := lxslices.DropWhile(original, func(n int) bool { return n%2 == 0 })

	// Modify the result
	if len(result) > 0 {
		result[0] = 999
	}

	// Original should be unchanged
	if !reflect.DeepEqual(original, originalCopy) {
		t.Error("DropWhile modified the original slice")
	}
}
