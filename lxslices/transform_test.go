package lxslices_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/nthanhhai2909/lx/lxslices"
	"github.com/nthanhhai2909/lx/lxtypes"
)

func TestPartitionN_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		n        int
		expected [][]int
	}{
		{
			name:     "evenly divided",
			slice:    []int{1, 2, 3, 4, 5, 6},
			n:        3,
			expected: [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name:     "unevenly divided (remainder 1)",
			slice:    []int{1, 2, 3, 4, 5, 6, 7},
			n:        3,
			expected: [][]int{{1, 2, 3}, {4, 5}, {6, 7}},
		},
		{
			name:     "unevenly divided (remainder 2)",
			slice:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			n:        3,
			expected: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}},
		},
		{
			name:     "n larger than slice length",
			slice:    []int{1, 2},
			n:        4,
			expected: [][]int{{1}, {2}, {}, {}},
		},
		{
			name:     "n equals slice length",
			slice:    []int{1, 2, 3},
			n:        3,
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "n is 1",
			slice:    []int{1, 2, 3},
			n:        1,
			expected: [][]int{{1, 2, 3}},
		},
		{
			name:     "empty slice",
			slice:    []int{},
			n:        3,
			expected: [][]int{},
		},
		{
			name:     "nil slice",
			slice:    nil,
			n:        3,
			expected: nil,
		},
		{
			name:     "n <= 0",
			slice:    []int{1, 2, 3},
			n:        0,
			expected: [][]int{},
		},
		{
			name:     "n is negative",
			slice:    []int{1, 2, 3},
			n:        -1,
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.PartitionN(tt.slice, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PartitionN(%v, %d) = %v; want %v", tt.slice, tt.n, result, tt.expected)
			}
		})
	}
}

func TestPartitionN_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		n        int
		expected [][]string
	}{
		{
			name:     "evenly divided strings",
			slice:    []string{"a", "b", "c", "d"},
			n:        2,
			expected: [][]string{{"a", "b"}, {"c", "d"}},
		},
		{
			name:     "unevenly divided strings (remainder 1)",
			slice:    []string{"a", "b", "c", "d", "e"},
			n:        2,
			expected: [][]string{{"a", "b", "c"}, {"d", "e"}},
		},
		{
			name:     "n larger than slice length",
			slice:    []string{"a", "b"},
			n:        3,
			expected: [][]string{{"a"}, {"b"}, {}},
		},
		{
			name:     "n is 1",
			slice:    []string{"a", "b", "c"},
			n:        1,
			expected: [][]string{{"a", "b", "c"}},
		},
		{
			name:     "empty slice",
			slice:    []string{},
			n:        2,
			expected: [][]string{},
		},
		{
			name:     "nil slice",
			slice:    nil,
			n:        2,
			expected: nil,
		},
		{
			name:     "n <= 0",
			slice:    []string{"a", "b", "c"},
			n:        0,
			expected: [][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.PartitionN(tt.slice, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PartitionN(%v, %d) = %v; want %v", tt.slice, tt.n, result, tt.expected)
			}
		})
	}
}

func TestPartitionN_Struct(t *testing.T) {
	type Item struct{ ID int }
	tests := []struct {
		name     string
		slice    []Item
		n        int
		expected [][]Item
	}{
		{
			name:     "evenly divided structs",
			slice:    []Item{{1}, {2}, {3}, {4}},
			n:        2,
			expected: [][]Item{{{1}, {2}}, {{3}, {4}}},
		},
		{
			name:     "unevenly divided structs",
			slice:    []Item{{1}, {2}, {3}},
			n:        2,
			expected: [][]Item{{{1}, {2}}, {{3}}},
		},
		{
			name:     "n larger than slice length",
			slice:    []Item{{1}},
			n:        2,
			expected: [][]Item{{{1}}, {}},
		},
		{
			name:     "n is 1",
			slice:    []Item{{1}, {2}},
			n:        1,
			expected: [][]Item{{{1}, {2}}},
		},
		{
			name:     "empty slice",
			slice:    []Item{},
			n:        2,
			expected: [][]Item{},
		},
		{
			name:     "nil slice",
			slice:    nil,
			n:        2,
			expected: nil,
		},
		{
			name:     "n <= 0",
			slice:    []Item{{1}},
			n:        0,
			expected: [][]Item{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.PartitionN(tt.slice, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PartitionN(%v, %d) = %v; want %v", tt.slice, tt.n, result, tt.expected)
			}
		})
	}
}

func TestMap_IntToInt(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		fn       func(int) int
		expected []int
	}{
		{
			name:     "double each number",
			slice:    []int{1, 2, 3, 4},
			fn:       func(v int) int { return v * 2 },
			expected: []int{2, 4, 6, 8},
		},
		{
			name:     "add 10 to each number",
			slice:    []int{1, 2, 3},
			fn:       func(v int) int { return v + 10 },
			expected: []int{11, 12, 13},
		},
		{
			name:     "square each number",
			slice:    []int{1, 2, 3, 4},
			fn:       func(v int) int { return v * v },
			expected: []int{1, 4, 9, 16},
		},
		{
			name:     "empty slice",
			slice:    []int{},
			fn:       func(v int) int { return v * 2 },
			expected: []int{},
		},
		{
			name:     "nil slice",
			slice:    nil,
			fn:       func(v int) int { return v * 2 },
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Map(tt.slice, tt.fn)
			if len(result) != len(tt.expected) {
				t.Errorf("Map() length = %v; want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Map() = %v; want %v", result, tt.expected)
					return
				}
			}
		})
	}
}

func TestMap_IntToString(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		fn       func(int) string
		expected []string
	}{
		{
			name:     "convert numbers to stars",
			slice:    []int{1, 2, 3},
			fn:       func(v int) string { return strings.Repeat("*", v) },
			expected: []string{"*", "**", "***"},
		},
		{
			name:  "convert to even/odd",
			slice: []int{1, 2, 3, 4},
			fn: func(v int) string {
				if v%2 == 0 {
					return "even"
				}
				return "odd"
			},
			expected: []string{"odd", "even", "odd", "even"},
		},
		{
			name:     "empty slice",
			slice:    []int{},
			fn:       func(v int) string { return strings.Repeat("*", v) },
			expected: []string{},
		},
		{
			name:     "nil slice",
			slice:    nil,
			fn:       func(v int) string { return strings.Repeat("*", v) },
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Map(tt.slice, tt.fn)
			if len(result) != len(tt.expected) {
				t.Errorf("Map() length = %v; want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Map() = %v; want %v", result, tt.expected)
					return
				}
			}
		})
	}
}

func TestMap_StringToInt(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		fn       func(string) int
		expected []int
	}{
		{
			name:     "string length",
			slice:    []string{"a", "bb", "ccc", "dddd"},
			fn:       func(s string) int { return len(s) },
			expected: []int{1, 2, 3, 4},
		},
		{
			name:  "count uppercase letters",
			slice: []string{"Hello", "WORLD", "Go"},
			fn: func(s string) int {
				count := 0
				for _, r := range s {
					if r >= 'A' && r <= 'Z' {
						count++
					}
				}
				return count
			},
			expected: []int{1, 5, 1},
		},
		{
			name:     "empty slice",
			slice:    []string{},
			fn:       func(s string) int { return len(s) },
			expected: []int{},
		},
		{
			name:     "nil slice",
			slice:    nil,
			fn:       func(s string) int { return len(s) },
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Map(tt.slice, tt.fn)
			if len(result) != len(tt.expected) {
				t.Errorf("Map() length = %v; want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Map() = %v; want %v", result, tt.expected)
					return
				}
			}
		})
	}
}

func TestMap_StringToString(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		fn       func(string) string
		expected []string
	}{
		{
			name:     "convert to uppercase",
			slice:    []string{"hello", "world", "go"},
			fn:       func(s string) string { return strings.ToUpper(s) },
			expected: []string{"HELLO", "WORLD", "GO"},
		},
		{
			name:     "add prefix",
			slice:    []string{"apple", "banana"},
			fn:       func(s string) string { return "fruit: " + s },
			expected: []string{"fruit: apple", "fruit: banana"},
		},
		{
			name:     "empty slice",
			slice:    []string{},
			fn:       func(s string) string { return strings.ToUpper(s) },
			expected: []string{},
		},
		{
			name:     "nil slice",
			slice:    nil,
			fn:       func(s string) string { return strings.ToUpper(s) },
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Map(tt.slice, tt.fn)
			if len(result) != len(tt.expected) {
				t.Errorf("Map() length = %v; want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Map() = %v; want %v", result, tt.expected)
					return
				}
			}
		})
	}
}

func TestMap_StructTransform(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	type UserDTO struct {
		ID       int
		FullName string
	}

	tests := []struct {
		name     string
		slice    []User
		fn       func(User) UserDTO
		expected []UserDTO
	}{
		{
			name: "convert User to UserDTO",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
			},
			fn: func(u User) UserDTO {
				return UserDTO{ID: u.ID, FullName: "Mr/Ms " + u.Name}
			},
			expected: []UserDTO{
				{1, "Mr/Ms Alice"},
				{2, "Mr/Ms Bob"},
			},
		},
		{
			name:  "empty slice",
			slice: []User{},
			fn: func(u User) UserDTO {
				return UserDTO{ID: u.ID, FullName: u.Name}
			},
			expected: []UserDTO{},
		},
		{
			name:  "nil slice",
			slice: nil,
			fn: func(u User) UserDTO {
				return UserDTO{ID: u.ID, FullName: u.Name}
			},
			expected: []UserDTO{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Map(tt.slice, tt.fn)
			if len(result) != len(tt.expected) {
				t.Errorf("Map() length = %v; want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Map()[%d] = %v; want %v", i, result[i], tt.expected[i])
					return
				}
			}
		})
	}
}

func TestFlatMap(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		slice := []int{1, 2, 3}
		fn := func(n int) []int { return []int{n, n * 2} }
		expected := []int{1, 2, 2, 4, 3, 6}
		result := lxslices.FlatMap(slice, fn)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("FlatMap() = %v; want %v", result, expected)
		}
	})

	t.Run("strings", func(t *testing.T) {
		slice := []string{"a", "b"}
		fn := func(s string) []string { return []string{s, s + s} }
		expected := []string{"a", "aa", "b", "bb"}
		result := lxslices.FlatMap(slice, fn)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("FlatMap() = %v; want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		fn := func(n int) []int { return []int{n} }
		result := lxslices.FlatMap(slice, fn)
		if len(result) != 0 {
			t.Errorf("FlatMap() = %v; want empty slice", result)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var slice []int
		fn := func(n int) []int { return []int{n} }
		result := lxslices.FlatMap(slice, fn)
		if len(result) != 0 {
			t.Errorf("FlatMap() = %v; want empty slice", result)
		}
	})

	t.Run("mapping to empty", func(t *testing.T) {
		slice := []int{1, 2, 3}
		fn := func(n int) []int { return []int{} }
		result := lxslices.FlatMap(slice, fn)
		if len(result) != 0 {
			t.Errorf("FlatMap() = %v; want empty slice", result)
		}
	})
}

func TestForEach(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		slice := []int{1, 2, 3}
		var sum int
		lxslices.ForEach(slice, func(n int) { sum += n })
		if sum != 6 {
			t.Errorf("ForEach() sum = %d; want 6", sum)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		called := false
		lxslices.ForEach(slice, func(n int) { called = true })
		if called {
			t.Errorf("ForEach() should not have been called on empty slice")
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var slice []int
		called := false
		lxslices.ForEach(slice, func(n int) { called = true })
		if called {
			t.Errorf("ForEach() should not have been called on nil slice")
		}
	})
}

func TestForEachIndexed(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		slice := []string{"a", "b", "c"}
		indices := []int{}
		values := []string{}

		lxslices.ForEachIndexed(slice, func(i int, s string) {
			indices = append(indices, i)
			values = append(values, s)
		})

		if !reflect.DeepEqual(indices, []int{0, 1, 2}) {
			t.Errorf("ForEachIndexed() indices = %v; want [0 1 2]", indices)
		}
		if !reflect.DeepEqual(values, []string{"a", "b", "c"}) {
			t.Errorf("ForEachIndexed() values = %v; want [a b c]", values)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		called := false
		lxslices.ForEachIndexed(slice, func(i int, n int) { called = true })
		if called {
			t.Errorf("ForEachIndexed() should not have been called on empty slice")
		}
	})
}

func TestReverse_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected []int
	}{
		{name: "integers", slice: []int{1, 2, 3, 4, 5}, expected: []int{5, 4, 3, 2, 1}},
		{name: "even number of elements", slice: []int{1, 2, 3, 4}, expected: []int{4, 3, 2, 1}},
		{name: "single element", slice: []int{42}, expected: []int{42}},
		{name: "empty slice", slice: []int{}, expected: []int{}},
		{name: "nil slice", slice: nil, expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lxslices.Reverse(tt.slice)
			if !reflect.DeepEqual(tt.slice, tt.expected) {
				t.Errorf("Reverse() = %v; want %v", tt.slice, tt.expected)
			}
		})
	}
}

func TestReverse_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		expected []string
	}{
		{name: "odd number of elements", slice: []string{"a", "b", "c"}, expected: []string{"c", "b", "a"}},
		{name: "even number of elements", slice: []string{"a", "b", "c", "d"}, expected: []string{"d", "c", "b", "a"}},
		{name: "single element", slice: []string{"go"}, expected: []string{"go"}},
		{name: "empty slice", slice: []string{}, expected: []string{}},
		{name: "nil slice", slice: nil, expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lxslices.Reverse(tt.slice)
			if !reflect.DeepEqual(tt.slice, tt.expected) {
				t.Errorf("Reverse() = %v; want %v", tt.slice, tt.expected)
			}
		})
	}
}

func TestReverse_Struct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	tests := []struct {
		name     string
		slice    []Person
		expected []Person
	}{
		{
			name: "odd number of elements",
			slice: []Person{
				{"Alice", 30},
				{"Bob", 25},
				{"Charlie", 35},
			},
			expected: []Person{
				{"Charlie", 35},
				{"Bob", 25},
				{"Alice", 30},
			},
		},
		{
			name: "even number of elements",
			slice: []Person{
				{"Alice", 30},
				{"Bob", 25},
				{"Charlie", 35},
				{"Dave", 40},
			},
			expected: []Person{
				{"Dave", 40},
				{"Charlie", 35},
				{"Bob", 25},
				{"Alice", 30},
			},
		},
		{
			name:     "single element",
			slice:    []Person{{"Alice", 30}},
			expected: []Person{{"Alice", 30}},
		},
		{name: "empty slice", slice: []Person{}, expected: []Person{}},
		{name: "nil slice", slice: nil, expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lxslices.Reverse(tt.slice)
			if !reflect.DeepEqual(tt.slice, tt.expected) {
				t.Errorf("Reverse() = %v; want %v", tt.slice, tt.expected)
			}
		})
	}
}

func TestGroupBy(t *testing.T) {
	t.Run("integers by parity", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		expected := map[string][]int{
			"even": {2, 4},
			"odd":  {1, 3, 5},
		}
		result := lxslices.GroupBy(slice, func(n int) string {
			if n%2 == 0 {
				return "even"
			}
			return "odd"
		})
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("GroupBy() = %v; want %v", result, expected)
		}
	})

	t.Run("strings by length", func(t *testing.T) {
		slice := []string{"a", "bb", "ccc", "dd", "e"}
		expected := map[int][]string{
			1: {"a", "e"},
			2: {"bb", "dd"},
			3: {"ccc"},
		}
		result := lxslices.GroupBy(slice, func(s string) int {
			return len(s)
		})
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("GroupBy() = %v; want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		expected := map[string][]int{}
		result := lxslices.GroupBy(slice, func(n int) string {
			if n%2 == 0 {
				return "even"
			}
			return "odd"
		})
		if len(result) != 0 {
			t.Errorf("GroupBy() = %v; want empty map", result)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("GroupBy() = %v; want %v", result, expected)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var slice []int
		expected := map[string][]int{}
		result := lxslices.GroupBy(slice, func(n int) string {
			if n%2 == 0 {
				return "even"
			}
			return "odd"
		})
		if len(result) != 0 {
			t.Errorf("GroupBy() = %v; want empty map", result)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("GroupBy() = %v; want %v", result, expected)
		}
	})

	t.Run("structs by property", func(t *testing.T) {
		type Person struct {
			Name   string
			Active bool
		}
		slice := []Person{
			{"Alice", true},
			{"Bob", false},
			{"Charlie", true},
		}
		expected := map[bool][]Person{
			true:  {{"Alice", true}, {"Charlie", true}},
			false: {{"Bob", false}},
		}
		result := lxslices.GroupBy(slice, func(p Person) bool {
			return p.Active
		})
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("GroupBy() = %v; want %v", result, expected)
		}
	})
}

func TestAssociateBy_Struct(t *testing.T) {
	type User struct {
		ID   int
		Name string
		Age  int
	}
	tests := []struct {
		name      string
		slice     []User
		fn        func(User) int
		expected  map[int]User
		expectErr bool
	}{
		{
			name: "user group by id",
			slice: []User{
				{1, "Alice", 25},
				{2, "Bob", 30},
				{3, "Charlie", 35},
			},
			fn: func(u User) int { return u.ID },
			expected: map[int]User{
				1: {1, "Alice", 25},
				2: {2, "Bob", 30},
				3: {3, "Charlie", 35},
			},
			expectErr: false,
		},
		{
			name: "duplicate keys error",
			slice: []User{
				{1, "Alice", 25},
				{1, "Bob", 30},
			},
			fn:        func(u User) int { return u.ID },
			expected:  nil,
			expectErr: true,
		},
		{
			name:      "empty slice",
			slice:     []User{},
			fn:        func(u User) int { return u.ID },
			expected:  map[int]User{},
			expectErr: false,
		},
		{
			name:      "nil slice",
			slice:     nil,
			fn:        func(u User) int { return u.ID },
			expected:  map[int]User{},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := lxslices.AssociateBy(tt.slice, tt.fn)
			if tt.expectErr {
				if err == nil {
					t.Errorf("AssociateBy() error = nil; want error")
				}
			} else {
				if err != nil {
					t.Errorf("AssociateBy() unexpected error = %v", err)
				}
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("AssociateBy() = %v; want %v", result, tt.expected)
				}
			}
		})
	}
}

func TestAssociateBy_Int(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		fn        func(int) string
		expected  map[string]int
		expectErr bool
	}{
		{
			name:      "duplicate keys",
			slice:     []int{1, 2, 3, 4, 5},
			fn:        func(n int) string { return "even" },
			expected:  nil,
			expectErr: true,
		},
		{
			name:      "unique keys",
			slice:     []int{1, 2, 3},
			fn:        func(n int) string { return fmt.Sprintf("key-%d", n) },
			expected:  map[string]int{"key-1": 1, "key-2": 2, "key-3": 3},
			expectErr: false,
		},
		{
			name:      "nil slice",
			slice:     nil,
			fn:        func(n int) string { return "even" },
			expected:  map[string]int{},
			expectErr: false,
		},
		{
			name:      "empty slice",
			slice:     []int{},
			fn:        func(n int) string { return "even" },
			expected:  map[string]int{},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := lxslices.AssociateBy(tt.slice, tt.fn)
			if tt.expectErr {
				if err == nil {
					t.Errorf("AssociateBy() error = nil; want error")
				}
			} else {
				if err != nil {
					t.Errorf("AssociateBy() unexpected error = %v", err)
				}
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("AssociateBy() = %v; want %v", result, tt.expected)
				}
			}
		})
	}
}

func TestConcat_Int(t *testing.T) {
	tests := []struct {
		name     string
		slices   [][]int
		expected []int
	}{
		{name: "no slices", slices: nil, expected: nil},
		{name: "all nil", slices: [][]int{nil, nil}, expected: nil},
		{name: "empty non-nil", slices: [][]int{{}, {}}, expected: []int{}},
		{name: "concat multiple", slices: [][]int{{1, 2}, {3}, {4, 5}}, expected: []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.Concat(tt.slices...)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Concat(%v) = %v; want %v", tt.slices, res, tt.expected)
			}
		})
	}
}

func TestConcat_String(t *testing.T) {
	tests := []struct {
		name     string
		slices   [][]string
		expected []string
	}{
		{name: "concat strings", slices: [][]string{{"a"}, {"b", "c"}}, expected: []string{"a", "b", "c"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.Concat(tt.slices...)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Concat(%v) = %v; want %v", tt.slices, res, tt.expected)
			}
		})
	}
}

func TestConcat_Struct(t *testing.T) {
	type Node struct {
		ID   int
		Name string
	}
	tests := []struct {
		name     string
		slices   [][]Node
		expected []Node
	}{
		{name: "concat nodes", slices: [][]Node{{{1, "a"}}, {{2, "b"}}}, expected: []Node{{1, "a"}, {2, "b"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.Concat(tt.slices...)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Concat(%v) = %v; want %v", tt.slices, res, tt.expected)
			}
		})
	}
}

func TestZip_IntString(t *testing.T) {
	// int and string
	a := []int{1, 2, 3}
	b := []string{"a", "b"}
	got := lxslices.Zip(a, b)
	want := []lxtypes.Pair[int, string]{{1, "a"}, {2, "b"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Zip(%v,%v) = %v; want %v", a, b, got, want)
	}
}

func TestZip_StringString_EmptyNil(t *testing.T) {
	// both nil -> nil
	if got := lxslices.Zip[string, string](nil, nil); got != nil {
		t.Fatalf("Zip(nil,nil) = %v; want nil", got)
	}
	// both empty non-nil -> empty slice
	a := []string{}
	b := []string{}
	if got := lxslices.Zip(a, b); got == nil || len(got) != 0 {
		t.Fatalf("Zip(empty,empty) = %v; want empty non-nil slice", got)
	}
}

func TestZip_Struct(t *testing.T) {
	type A struct{ X int }
	type B struct{ Y string }
	a := []A{{1}, {2}}
	b := []B{{"one"}, {"two"}}
	got := lxslices.Zip(a, b)
	want := []lxtypes.Pair[A, B]{{First: A{1}, Second: B{"one"}}, {First: A{2}, Second: B{"two"}}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Zip(structs) = %v; want %v", got, want)
	}
}

func TestUnzip_IntString(t *testing.T) {
	pairs := []lxtypes.Pair[int, string]{{1, "a"}, {2, "b"}}
	a, b := lxslices.Unzip(pairs)
	if !reflect.DeepEqual(a, []int{1, 2}) || !reflect.DeepEqual(b, []string{"a", "b"}) {
		t.Fatalf("Unzip(%v) = (%v, %v); want (%v, %v)", pairs, a, b, []int{1, 2}, []string{"a", "b"})
	}
}

func TestUnzip_EmptyNil(t *testing.T) {
	// nil -> nil,nil
	a, b := lxslices.Unzip[any, any](nil)
	if a != nil || b != nil {
		t.Fatalf("Unzip(nil) = (%v, %v); want (nil, nil)", a, b)
	}
	// empty non-nil -> empty slices
	pairs := []lxtypes.Pair[int, int]{}
	a2, b2 := lxslices.Unzip(pairs)
	if a2 == nil || b2 == nil || len(a2) != 0 || len(b2) != 0 {
		t.Fatalf("Unzip(empty) = (%v, %v); want empty non-nil slices", a2, b2)
	}
}

// ---- end moved tests ----

func TestCopy_Int(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{name: "nil slice", slice: nil, want: nil},
		{name: "empty slice", slice: []int{}, want: []int{}},
		{name: "single element", slice: []int{42}, want: []int{42}},
		{name: "multiple values", slice: []int{1, 2, 3}, want: []int{1, 2, 3}},
		{name: "large slice", slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// snapshot original for later comparisons
			orig := append([]int(nil), tt.slice...)

			got := lxslices.Copy(tt.slice)

			// Check values match
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Copy(%v) = %v; want %v", tt.slice, got, tt.want)
			}

			// Check nil vs empty distinction
			if tt.slice == nil && got != nil {
				t.Fatalf("Copy(nil) should return nil, got %v", got)
			}
			if tt.slice != nil && len(tt.slice) == 0 && got == nil {
				t.Fatalf("Copy(empty non-nil) should return empty non-nil slice, got nil")
			}

			// Further independence checks only make sense for non-nil and non-empty
			if tt.slice != nil && len(tt.slice) > 0 {
				// backing array independence: compare addresses of first element
				origAddr := &tt.slice[0]
				gotAddr := &got[0]
				if origAddr == gotAddr {
					t.Fatalf("Copy must allocate a new backing array; addresses equal: %p", origAddr)
				}

				// 1) Mutate the original and ensure copy is unaffected
				tt.slice[0] = tt.slice[0] + 100
				if reflect.DeepEqual(got, tt.slice) {
					t.Fatalf("Copy should produce independent slice; original change reflected in copy")
				}
				// restore original from snapshot for the next checks
				tt.slice = append([]int(nil), orig...)

				// 2) Mutate the copy and ensure original is unaffected
				got[0] = got[0] + 200
				if reflect.DeepEqual(got, tt.slice) {
					t.Fatalf("Modifying copy should not affect original slice")
				}

				// 3) Ensure append operations are independent (appending to copy should not change original)
				originalLen := len(tt.slice)
				originalCap := cap(tt.slice)
				gotCopy := append([]int(nil), got...)
				got = append(got, 999)
				if len(tt.slice) != originalLen {
					t.Fatalf("Appending to copy affected original length: want %d got %d", originalLen, len(tt.slice))
				}
				if cap(tt.slice) != originalCap {
					t.Fatalf("Appending to copy affected original capacity: want %d got %d", originalCap, cap(tt.slice))
				}
				// restore got
				got = gotCopy
			}
		})
	}
}

func TestCopy_String(t *testing.T) {
	tests := []struct {
		name  string
		slice []string
		want  []string
	}{
		{name: "nil slice", slice: nil, want: nil},
		{name: "empty slice", slice: []string{}, want: []string{}},
		{name: "values", slice: []string{"a", "b"}, want: []string{"a", "b"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orig := append([]string(nil), tt.slice...)

			got := lxslices.Copy(tt.slice)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Copy(%v) = %v; want %v", tt.slice, got, tt.want)
			}

			if tt.slice != nil && len(tt.slice) > 0 {
				// mutate original
				tt.slice[0] = tt.slice[0] + "!"
				if reflect.DeepEqual(got, tt.slice) {
					t.Fatalf("Copy should produce independent slice; original change reflected in copy")
				}
				// restore
				tt.slice = append([]string(nil), orig...)

				// mutate copy
				got[0] = got[0] + "?"
				if reflect.DeepEqual(got, tt.slice) {
					t.Fatalf("Modifying copy should not affect original slice")
				}
			}
		})
	}
}

func TestCopy_Struct(t *testing.T) {
	type Node struct{ ID int }
	tests := []struct {
		name  string
		slice []Node
		want  []Node
	}{
		{name: "nil slice", slice: nil, want: nil},
		{name: "empty slice", slice: []Node{}, want: []Node{}},
		{name: "values", slice: []Node{{1}, {2}}, want: []Node{{1}, {2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orig := append([]Node(nil), tt.slice...)

			got := lxslices.Copy(tt.slice)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Copy(%v) = %v; want %v", tt.slice, got, tt.want)
			}

			if tt.slice != nil && len(tt.slice) > 0 {
				// mutate original element
				tt.slice[0].ID += 10
				if reflect.DeepEqual(got, tt.slice) {
					t.Fatalf("Copy should produce independent slice; original change reflected in copy")
				}
				// restore
				tt.slice = append([]Node(nil), orig...)

				// mutate copy element
				got[0].ID += 20
				if reflect.DeepEqual(got, tt.slice) {
					t.Fatalf("Modifying copy should not affect original slice")
				}
			}
		})
	}
}

// Clone is an alias for Copy; test the alias behavior separately
func TestClone_Int(t *testing.T) {
	src := []int{1, 2, 3}
	cl := lxslices.Clone(src)
	if !reflect.DeepEqual(cl, src) {
		t.Fatalf("Clone(%v) = %v; want %v", src, cl, src)
	}
	// mutate source and ensure clone unaffected
	src[0] = 999
	if cl[0] == src[0] {
		t.Fatalf("Clone should be independent of source backing array")
	}

	// mutate clone and ensure source unaffected
	cl[1] = 555
	if cl[1] == src[1] {
		t.Fatalf("Modifying clone should not affect source")
	}
}

func TestClone_String(t *testing.T) {
	src := []string{"x", "y"}
	cl := lxslices.Clone(src)
	if !reflect.DeepEqual(cl, src) {
		t.Fatalf("Clone(%v) = %v; want %v", src, cl, src)
	}
	src[1] = "z"
	if cl[1] == src[1] {
		t.Fatalf("Clone should be independent of source backing array")
	}
	cl[0] = "q"
	if cl[0] == src[0] {
		t.Fatalf("Modifying clone should not affect source")
	}
}

func TestClone_Struct(t *testing.T) {
	type Node struct{ ID int }
	src := []Node{{1}, {2}}
	cl := lxslices.Clone(src)
	if !reflect.DeepEqual(cl, src) {
		t.Fatalf("Clone(%v) = %v; want %v", src, cl, src)
	}
	src[0].ID = 77
	if cl[0].ID == src[0].ID {
		t.Fatalf("Clone should be independent of source backing array")
	}
	cl[1].ID = 88
	if cl[1].ID == src[1].ID {
		t.Fatalf("Modifying clone should not affect source")
	}
}

func TestChunk_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		size     int
		expected [][]int
	}{
		{
			name:     "perfectly divisible",
			slice:    []int{1, 2, 3, 4, 5, 6},
			size:     2,
			expected: [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name:     "not perfectly divisible",
			slice:    []int{1, 2, 3, 4, 5},
			size:     2,
			expected: [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name:     "size larger than slice",
			slice:    []int{1, 2, 3},
			size:     10,
			expected: [][]int{{1, 2, 3}},
		},
		{
			name:     "size equals slice length",
			slice:    []int{1, 2, 3},
			size:     3,
			expected: [][]int{{1, 2, 3}},
		},
		{
			name:     "empty slice",
			slice:    []int{},
			size:     2,
			expected: [][]int{},
		},
		{
			name:     "nil slice",
			slice:    nil,
			size:     2,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Chunk(tt.slice, tt.size)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Chunk() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestChunk_PanicNegativeZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Chunk did not panic on zero size")
		}
	}()
	lxslices.Chunk([]int{1, 2, 3}, 0)
}

func TestChunk_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		size     int
		expected [][]string
	}{
		{
			name:     "chunk strings",
			slice:    []string{"a", "b", "c"},
			size:     2,
			expected: [][]string{{"a", "b"}, {"c"}},
		},
		{
			name:     "empty slice",
			slice:    []string{},
			size:     2,
			expected: [][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Chunk(tt.slice, tt.size)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Chunk() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestChunk_Struct(t *testing.T) {
	type Item struct{ ID int }
	tests := []struct {
		name     string
		slice    []Item
		size     int
		expected [][]Item
	}{
		{
			name:     "chunk structs",
			slice:    []Item{{1}, {2}, {3}, {4}},
			size:     2,
			expected: [][]Item{{{1}, {2}}, {{3}, {4}}},
		},
		{
			name:     "single element chunk",
			slice:    []Item{{1}, {2}},
			size:     1,
			expected: [][]Item{{{1}}, {{2}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Chunk(tt.slice, tt.size)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Chunk() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestSplitAt_Int(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		index     int
		expected1 []int
		expected2 []int
	}{
		{
			name:      "split in middle",
			slice:     []int{1, 2, 3, 4},
			index:     2,
			expected1: []int{1, 2},
			expected2: []int{3, 4},
		},
		{
			name:      "split at 0",
			slice:     []int{1, 2, 3, 4},
			index:     0,
			expected1: nil,
			expected2: []int{1, 2, 3, 4},
		},
		{
			name:      "split at negative index",
			slice:     []int{1, 2, 3, 4},
			index:     -1,
			expected1: nil,
			expected2: []int{1, 2, 3, 4},
		},
		{
			name:      "split at end",
			slice:     []int{1, 2, 3, 4},
			index:     4,
			expected1: []int{1, 2, 3, 4},
			expected2: nil,
		},
		{
			name:      "split beyond end",
			slice:     []int{1, 2, 3, 4},
			index:     5,
			expected1: []int{1, 2, 3, 4},
			expected2: nil,
		},
		{
			name:      "empty slice",
			slice:     []int{},
			index:     0,
			expected1: nil,
			expected2: []int{},
		},
		{
			name:      "nil slice",
			slice:     nil,
			index:     2,
			expected1: nil,
			expected2: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result1, result2 := lxslices.SplitAt(tt.slice, tt.index)
			if !reflect.DeepEqual(result1, tt.expected1) {
				t.Errorf("SplitAt() result1 = %v; want %v", result1, tt.expected1)
			}
			if !reflect.DeepEqual(result2, tt.expected2) {
				t.Errorf("SplitAt() result2 = %v; want %v", result2, tt.expected2)
			}
		})
	}
}

func TestSplitAt_String(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		index     int
		expected1 []string
		expected2 []string
	}{
		{
			name:      "split at 0",
			slice:     []string{"a", "b", "c"},
			index:     0,
			expected1: nil,
			expected2: []string{"a", "b", "c"},
		},
		{
			name:      "split at negative index",
			slice:     []string{"a", "b", "c"},
			index:     -1,
			expected1: nil,
			expected2: []string{"a", "b", "c"},
		},
		{
			name:      "split at end",
			slice:     []string{"a", "b", "c"},
			index:     3,
			expected1: []string{"a", "b", "c"},
			expected2: nil,
		},
		{
			name:      "split beyond end",
			slice:     []string{"a", "b", "c"},
			index:     4,
			expected1: []string{"a", "b", "c"},
			expected2: nil,
		},
		{
			name:      "empty slice",
			slice:     []string{},
			index:     0,
			expected1: nil,
			expected2: []string{},
		},
		{
			name:      "nil slice",
			slice:     nil,
			index:     1,
			expected1: nil,
			expected2: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result1, result2 := lxslices.SplitAt(tt.slice, tt.index)
			if !reflect.DeepEqual(result1, tt.expected1) {
				t.Errorf("SplitAt() result1 = %v; want %v", result1, tt.expected1)
			}
			if !reflect.DeepEqual(result2, tt.expected2) {
				t.Errorf("SplitAt() result2 = %v; want %v", result2, tt.expected2)
			}
		})
	}
}

func TestSplitAt_Struct(t *testing.T) {
	type Item struct{ ID int }
	tests := []struct {
		name      string
		slice     []Item
		index     int
		expected1 []Item
		expected2 []Item
	}{
		{
			name:      "split in middle",
			slice:     []Item{{1}, {2}, {3}, {4}},
			index:     3,
			expected1: []Item{{1}, {2}, {3}},
			expected2: []Item{{4}},
		},
		{
			name:      "split at 0",
			slice:     []Item{{1}, {2}},
			index:     0,
			expected1: nil,
			expected2: []Item{{1}, {2}},
		},
		{
			name:      "split at negative index",
			slice:     []Item{{1}, {2}},
			index:     -1,
			expected1: nil,
			expected2: []Item{{1}, {2}},
		},
		{
			name:      "split at end",
			slice:     []Item{{1}, {2}},
			index:     2,
			expected1: []Item{{1}, {2}},
			expected2: nil,
		},
		{
			name:      "split beyond end",
			slice:     []Item{{1}},
			index:     2,
			expected1: []Item{{1}},
			expected2: nil,
		},
		{
			name:      "empty slice",
			slice:     []Item{},
			index:     0,
			expected1: nil,
			expected2: []Item{},
		},
		{
			name:      "nil slice",
			slice:     nil,
			index:     1,
			expected1: nil,
			expected2: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result1, result2 := lxslices.SplitAt(tt.slice, tt.index)
			if !reflect.DeepEqual(result1, tt.expected1) {
				t.Errorf("SplitAt() result1 = %v; want %v", result1, tt.expected1)
			}
			if !reflect.DeepEqual(result2, tt.expected2) {
				t.Errorf("SplitAt() result2 = %v; want %v", result2, tt.expected2)
			}
		})
	}
}
