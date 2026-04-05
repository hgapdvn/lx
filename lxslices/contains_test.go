package lxslices_test

import (
	"strings"
	"testing"

	"github.com/hgapdvn/lx/lxslices"
)

func TestContains_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		elem     int
		expected bool
	}{
		{
			name:     "element exists in middle",
			slice:    []int{1, 2, 3, 4},
			elem:     3,
			expected: true,
		},
		{
			name:     "element exists at beginning",
			slice:    []int{1, 2, 3, 4},
			elem:     1,
			expected: true,
		},
		{
			name:     "element exists at end",
			slice:    []int{1, 2, 3, 4},
			elem:     4,
			expected: true,
		},
		{
			name:     "element does not exist",
			slice:    []int{1, 2, 3, 4},
			elem:     10,
			expected: false,
		},
		{
			name:     "empty slice",
			slice:    []int{},
			elem:     1,
			expected: false,
		},
		{
			name:     "nil slice",
			slice:    nil,
			elem:     1,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Contains(tt.slice, tt.elem)
			if result != tt.expected {
				t.Errorf("Contains(%v, %v) = %v; want %v",
					tt.slice, tt.elem, result, tt.expected)
			}
		})
	}
}

func TestContains_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		elem     string
		expected bool
	}{
		{"exist at the beginning", []string{"a", "b", "c"}, "a", true},
		{"exist in the middle", []string{"a", "b", "c"}, "b", true},
		{"exist at the end", []string{"a", "b", "c"}, "c", true},
		{"not exist", []string{"a", "b", "c"}, "x", false},
		{"empty slice", []string{}, "a", false},
		{"nil slice", nil, "a", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Contains(tt.slice, tt.elem)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestContains_Struct(t *testing.T) {
	type User struct {
		ID int
	}
	tests := []struct {
		name     string
		slice    []User
		elem     User
		expected bool
	}{
		{
			name:     "exist at the beginning",
			slice:    []User{{1}, {2}, {3}},
			elem:     User{1},
			expected: true,
		},
		{
			name:     "exist in the middle",
			slice:    []User{{1}, {2}, {3}},
			elem:     User{2},
			expected: true,
		},
		{
			name:     "exist at the end",
			slice:    []User{{1}, {2}, {3}},
			elem:     User{3},
			expected: true,
		},
		{
			name:     "struct not found",
			slice:    []User{{1}, {2}, {3}},
			elem:     User{4},
			expected: false,
		},
		{
			name:  "empty slice",
			slice: []User{},
			elem:  User{1},
		},
		{
			name:  "nil slice",
			slice: nil,
			elem:  User{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Contains(tt.slice, tt.elem)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestContainsAny_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		elems    []int
		expected bool
	}{
		{
			name:     "exist at the beginning",
			slice:    []int{1, 2, 3},
			elems:    []int{1},
			expected: true,
		},
		{
			name:     "exist in the middle",
			slice:    []int{1, 2, 3},
			elems:    []int{2},
			expected: true,
		},
		{
			name:     "exist at the end",
			slice:    []int{1, 2, 3},
			elems:    []int{3},
			expected: true,
		},
		{
			name:     "not exist",
			slice:    []int{1, 2, 3},
			elems:    []int{4},
			expected: false,
		},
		{
			name:     "empty slice",
			slice:    []int{},
			elems:    []int{1},
			expected: false,
		},
		{
			name:     "nil slice",
			slice:    nil,
			elems:    []int{1},
			expected: false,
		},
		{
			name:     "empty elems",
			slice:    []int{1, 2, 3},
			elems:    []int{},
			expected: false,
		},
		{
			name:     "nil elems",
			slice:    []int{1, 2, 3},
			elems:    nil,
			expected: false,
		},
		{
			name:     "all elems are in slice",
			slice:    []int{1, 2, 3},
			elems:    []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "all elems are not in slice",
			slice:    []int{1, 2, 3},
			elems:    []int{4, 5, 6},
			expected: false,
		},
		{
			name:     "all elems are in slice, but in different order",
			slice:    []int{1, 2, 3},
			elems:    []int{3, 2, 1},
			expected: true,
		},
		{
			name:     "elems are in slice, but in different order and some are not in slice",
			slice:    []int{1, 2, 3},
			elems:    []int{3, 2, 4},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.ContainsAny(tt.slice, tt.elems...)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestContainsAny_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		elems    []string
		expected bool
	}{
		{
			name:     "exist at the beginning",
			slice:    []string{"a", "b", "c"},
			elems:    []string{"a"},
			expected: true,
		},
		{
			name:     "exist in the middle",
			slice:    []string{"a", "b", "c"},
			elems:    []string{"b"},
			expected: true,
		},
		{
			name:     "exist at the end",
			slice:    []string{"a", "b", "c"},
			elems:    []string{"c"},
			expected: true,
		},
		{
			name:     "not exist",
			slice:    []string{"a", "b", "c"},
			elems:    []string{"d"},
			expected: false,
		},
		{
			name:     "empty slice",
			slice:    []string{},
			elems:    []string{"a"},
			expected: false,
		},
		{
			name:     "nil slice",
			slice:    nil,
			elems:    []string{"a"},
			expected: false,
		},
		{
			name:     "empty elems",
			slice:    []string{"a", "b", "c"},
			elems:    []string{},
			expected: false,
		},
		{
			name:     "nil elems",
			slice:    []string{"a", "b", "c"},
			elems:    nil,
			expected: false,
		},
		{
			name:     "all elems are in slice",
			slice:    []string{"a", "b", "c"},
			elems:    []string{"a", "b", "c"},
			expected: true,
		},
		{
			name:     "all elems are not in slice",
			slice:    []string{"a", "b", "c"},
			elems:    []string{"d", "e", "f"},
			expected: false,
		},
		{
			name:     "all elems are in slice, but in different order",
			slice:    []string{"a", "b", "c"},
			elems:    []string{"c", "b", "a"},
			expected: true,
		},
		{
			name:     "elems are in slice, but in different order and some are not in slice",
			slice:    []string{"a", "b", "c"},
			elems:    []string{"c", "b", "d"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.ContainsAny(tt.slice, tt.elems...)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestContainsAny_Struct(t *testing.T) {
	type User struct {
		ID int
	}
	tests := []struct {
		name     string
		slice    []User
		elems    []User
		expected bool
	}{
		{
			name:     "exist at the beginning",
			slice:    []User{{1}, {2}, {3}},
			elems:    []User{{1}},
			expected: true,
		},
		{
			name:     "exist in the middle",
			slice:    []User{{1}, {2}, {3}},
			elems:    []User{{2}},
			expected: true,
		},
		{
			name:     "exist at the end",
			slice:    []User{{1}, {2}, {3}},
			elems:    []User{{3}},
			expected: true,
		},
		{
			name:     "not exist",
			slice:    []User{{1}, {2}, {3}},
			elems:    []User{{4}},
			expected: false,
		},
		{
			name:     "empty slice",
			slice:    []User{},
			elems:    []User{{1}},
			expected: false,
		},
		{
			name:     "nil slice",
			slice:    nil,
			elems:    []User{{1}},
			expected: false,
		},
		{
			name:     "empty elems",
			slice:    []User{{1}, {2}, {3}},
			elems:    []User{},
			expected: false,
		},
		{
			name:     "nil elems",
			slice:    []User{{1}, {2}, {3}},
			elems:    nil,
			expected: false,
		},
		{
			name:     "all elems are in slice",
			slice:    []User{{1}, {2}, {3}},
			elems:    []User{{1}, {2}, {3}},
			expected: true,
		},
		{
			name:     "all elems are not in slice",
			slice:    []User{{1}, {2}, {3}},
			elems:    []User{{4}, {5}, {6}},
			expected: false,
		},
		{
			name:     "all elems are in slice, but in different order",
			slice:    []User{{1}, {2}, {3}},
			elems:    []User{{3}, {2}, {1}},
			expected: true,
		},
		{
			name:     "elems are in slice, but in different order and some are not in slice",
			slice:    []User{{1}, {2}, {3}},
			elems:    []User{{3}, {2}, {4}},
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.ContainsAny(tt.slice, tt.elems...)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestContainsAll_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		elems    []int
		expected bool
	}{
		{
			name:     "all elems are in slice",
			slice:    []int{1, 2, 3},
			elems:    []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "all elems are not in slice",
			slice:    []int{1, 2, 3},
			elems:    []int{4, 5, 6},
			expected: false,
		},
		{
			name:     "all elems are in slice, but in different order",
			slice:    []int{1, 2, 3},
			elems:    []int{3, 2, 1},
			expected: true,
		},
		{
			name:     "elems are in slice, but in different order and some are not in slice",
			slice:    []int{1, 2, 3},
			elems:    []int{3, 2, 4},
			expected: false,
		},
		{
			name:     "empty slice",
			slice:    []int{},
			elems:    []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "nil slice",
			slice:    nil,
			elems:    []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "empty elems",
			slice:    []int{1, 2, 3},
			elems:    []int{},
			expected: true,
		},
		{
			name:     "nil elems",
			slice:    []int{1, 2, 3},
			elems:    nil,
			expected: true,
		},
		{
			name:     "extra elems are not in slice",
			slice:    []int{1, 2, 3},
			elems:    []int{1, 2, 3, 4},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.ContainsAll(tt.slice, tt.elems...)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestContainsAll_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		elems    []string
		expected bool
	}{
		{
			name:     "all elems are in slice",
			slice:    []string{"a", "b", "c"},
			elems:    []string{"a", "b", "c"},
			expected: true,
		},
		{
			name:     "all elems are not in slice",
			slice:    []string{"a", "b", "c"},
			elems:    []string{"d", "e", "f"},
			expected: false,
		},
		{
			name:     "all elems are in slice, but in different order",
			slice:    []string{"a", "b", "c"},
			elems:    []string{"c", "b", "a"},
			expected: true,
		},
		{
			name:     "elems are in slice, but in different order and some are not in slice",
			slice:    []string{"a", "b", "c"},
			elems:    []string{"c", "b", "d"},
			expected: false,
		},
		{
			name:     "empty slice",
			slice:    []string{},
			elems:    []string{"a", "b", "c"},
			expected: false,
		},
		{
			name:     "nil slice",
			slice:    nil,
			elems:    []string{"a", "b", "c"},
			expected: false,
		},
		{
			name:     "empty elems",
			slice:    []string{"a", "b", "c"},
			elems:    []string{},
			expected: true,
		},
		{
			name:     "nil elems",
			slice:    []string{"a", "b", "c"},
			elems:    nil,
			expected: true,
		},
		{
			name:     "extra elems are not in slice",
			slice:    []string{"a", "b", "c"},
			elems:    []string{"a", "b", "c", "d"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.ContainsAll(tt.slice, tt.elems...)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestContainsAll_Struct(t *testing.T) {
	type User struct {
		ID int
	}
	tests := []struct {
		name     string
		slice    []User
		elems    []User
		expected bool
	}{
		{
			name:     "all elems are in slice",
			slice:    []User{{1}, {2}, {3}},
			elems:    []User{{1}, {2}, {3}},
			expected: true,
		},
		{
			name:     "all elems are not in slice",
			slice:    []User{{1}, {2}, {3}},
			elems:    []User{{4}, {5}, {6}},
			expected: false,
		},
		{
			name:     "all elems are in slice, but in different order",
			slice:    []User{{1}, {2}, {3}},
			elems:    []User{{3}, {2}, {1}},
			expected: true,
		},
		{
			name:     "extra elems are not in slice",
			slice:    []User{{1}, {2}, {3}},
			elems:    []User{{3}, {2}, {4}},
			expected: false,
		},
		{
			name:     "empty slice",
			slice:    []User{},
			elems:    []User{{1}, {2}, {3}},
			expected: false,
		},
		{
			name:     "nil slice",
			slice:    nil,
			elems:    []User{{1}, {2}, {3}},
			expected: false,
		},
		{
			name:     "empty elems",
			slice:    []User{{1}, {2}, {3}},
			elems:    []User{},
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.ContainsAll(tt.slice, tt.elems...)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestContainsFunc_Int(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		expected  bool
	}{
		{
			name:      "match exists",
			slice:     []int{1, 2, 3, 4},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  true,
		},
		{
			name:      "no match",
			slice:     []int{1, 3, 5},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  false,
		},
		{
			name:      "empty slice",
			slice:     []int{},
			predicate: func(v int) bool { return true },
			expected:  false,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(v int) bool { return true },
			expected:  false,
		},
		{
			name:      "predicate always true",
			slice:     []int{1, 2, 3},
			predicate: func(v int) bool { return true },
			expected:  true,
		},
		{
			name:      "predicate always false",
			slice:     []int{1, 2, 3},
			predicate: func(v int) bool { return false },
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.ContainsFunc(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("ContainsFunc(%v) = %v; want %v",
					tt.slice, result, tt.expected)
			}
		})
	}
}

func TestContainsFunc_Struct(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	tests := []struct {
		name      string
		slice     []User
		predicate func(User) bool
		expected  bool
	}{
		{
			name: "find by ID",
			slice: []User{
				{1, "A"},
				{2, "B"},
				{3, "C"},
			},
			predicate: func(u User) bool {
				return u.ID == 2
			},
			expected: true,
		},
		{
			name: "find by Name",
			slice: []User{
				{1, "A"},
				{2, "B"},
				{3, "C"},
			},
			predicate: func(u User) bool {
				return u.Name == "B"
			},
			expected: true,
		},
		{
			name: "find by Name (case insensitive)",
			slice: []User{
				{1, "A"},
				{2, "B"},
				{3, "C"},
			},
			predicate: func(u User) bool {
				return strings.EqualFold(u.Name, "b")
			},
			expected: true,
		},
		{
			name:  "empty slice but predicate always true",
			slice: []User{},
			predicate: func(u User) bool {
				return true
			},
			expected: false,
		},
		{
			name:  "nil slice but predicate always true",
			slice: nil,
			predicate: func(u User) bool {
				return true
			},
		},
		{
			name:  "empty slice but predicate always false",
			slice: []User{},
			predicate: func(u User) bool {
				return false
			},
			expected: false,
		},
		{
			name:  "nil slice but predicate always false",
			slice: nil,
			predicate: func(u User) bool {
				return false
			},
		},
		{
			name: "predicate always false",
			slice: []User{
				{1, "A"},
				{2, "B"},
				{3, "C"},
			},
			predicate: func(u User) bool {
				return false
			},
			expected: false,
		},
		{
			name: "predicate always true",
			slice: []User{
				{1, "A"},
				{2, "B"},
				{3, "C"},
			},
			predicate: func(u User) bool {
				return true
			},
			expected: true,
		},
		{
			name: "not found",
			predicate: func(u User) bool {
				return u.Name == "X"
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.ContainsFunc(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
