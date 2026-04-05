package lxslices_test

import (
	"math"
	"testing"

	"github.com/hgapdvn/lx/lxslices"
)

func TestIndex_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		elem     int
		expected int
	}{
		{
			name:     "element at beginning",
			slice:    []int{10, 20, 30},
			elem:     10,
			expected: 0,
		},
		{
			name:     "element in middle",
			slice:    []int{10, 20, 30},
			elem:     20,
			expected: 1,
		},
		{
			name:     "element at end",
			slice:    []int{10, 20, 30},
			elem:     30,
			expected: 2,
		},
		{
			name:     "element not found",
			slice:    []int{10, 20, 30},
			elem:     40,
			expected: -1,
		},
		{
			name:     "empty slice",
			slice:    []int{},
			elem:     1,
			expected: -1,
		},
		{
			name:     "nil slice",
			slice:    nil,
			elem:     1,
			expected: -1,
		},
		{
			name:     "duplicate elements returns first index",
			slice:    []int{5, 7, 5, 9},
			elem:     5,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Index(tt.slice, tt.elem)
			if result != tt.expected {
				t.Errorf("Index(%v, %v) = %d; want %d",
					tt.slice, tt.elem, result, tt.expected)
			}
		})
	}
}

func TestIndex_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		elem     string
		expected int
	}{
		{
			name:     "element in middle",
			slice:    []string{"apple", "banana", "cherry"},
			elem:     "banana",
			expected: 1,
		},
		{
			name:     "element at beginning",
			slice:    []string{"apple", "banana", "cherry"},
			elem:     "apple",
			expected: 0,
		},
		{
			name:     "element at end",
			slice:    []string{"apple", "banana", "cherry"},
			elem:     "cherry",
			expected: 2,
		},
		{
			name:     "element not found",
			slice:    []string{"apple", "banana", "cherry"},
			elem:     "orange",
			expected: -1,
		},
		{
			name:     "empty slice",
			slice:    []string{},
			elem:     "apple",
			expected: -1,
		},
		{
			name:     "nil slice",
			slice:    nil,
			elem:     "apple",
			expected: -1,
		},
		{
			name:     "duplicate elements returns first index",
			slice:    []string{"go", "java", "go", "rust"},
			elem:     "go",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Index(tt.slice, tt.elem)
			if result != tt.expected {
				t.Errorf("Index(%v, %q) = %d; want %d",
					tt.slice, tt.elem, result, tt.expected)
			}
		})
	}
}

func TestIndex_Struct(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	tests := []struct {
		name     string
		slice    []User
		elem     User
		expected int
	}{
		{
			name: "element at beginning",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
			},
			elem:     User{1, "Alice"},
			expected: 0,
		},
		{
			name: "element in middle",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{3, "Charlie"},
			},
			elem:     User{2, "Bob"},
			expected: 1,
		},
		{
			name: "element at end",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
			},
			elem:     User{2, "Bob"},
			expected: 1,
		},
		{
			name: "element not found",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
			},
			elem:     User{3, "Charlie"},
			expected: -1,
		},
		{
			name:     "empty slice",
			slice:    []User{},
			elem:     User{1, "Alice"},
			expected: -1,
		},
		{
			name:     "nil slice",
			slice:    nil,
			elem:     User{1, "Alice"},
			expected: -1,
		},
		{
			name: "duplicate returns first index",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{1, "Alice"},
			},
			elem:     User{1, "Alice"},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Index(tt.slice, tt.elem)
			if result != tt.expected {
				t.Errorf("Index(%v, %+v) = %d; want %d",
					tt.slice, tt.elem, result, tt.expected)
			}
		})
	}
}

func TestIndexFunc_Int(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		expected  int
	}{
		{
			name:      "match in middle",
			slice:     []int{1, 3, 4, 7},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  2, // 4
		},
		{
			name:      "match at beginning",
			slice:     []int{2, 3, 5},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  0,
		},
		{
			name:      "match at end",
			slice:     []int{1, 3, 6},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  2,
		},
		{
			name:      "no match",
			slice:     []int{1, 3, 5},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  -1,
		},
		{
			name:      "empty slice",
			slice:     []int{},
			predicate: func(v int) bool { return true },
			expected:  -1,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(v int) bool { return true },
			expected:  -1,
		},
		{
			name:      "multiple matches returns first",
			slice:     []int{1, 4, 6, 8},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.IndexFunc(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("IndexFunc(%v) = %d; want %d",
					tt.slice, result, tt.expected)
			}
		})
	}
}

func TestIndexFunc_String(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		predicate func(string) bool
		expected  int
	}{
		{
			name:      "match in middle",
			slice:     []string{"apple", "banana", "cherry"},
			predicate: func(v string) bool { return v == "banana" },
			expected:  1,
		},
		{
			name:      "match at beginning",
			slice:     []string{"apple", "banana", "cherry"},
			predicate: func(v string) bool { return v == "apple" },
			expected:  0,
		},
		{
			name:      "match at end",
			slice:     []string{"apple", "banana", "cherry"},
			predicate: func(v string) bool { return v == "cherry" },
			expected:  2,
		},
		{
			name:      "no match",
			slice:     []string{"apple", "banana", "cherry"},
			predicate: func(v string) bool { return len(v) == 4 },
			expected:  -1,
		},
		{
			name:      "empty slice",
			slice:     []string{},
			predicate: func(v string) bool { return true },
			expected:  -1,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(v string) bool { return true },
			expected:  -1,
		},
		{
			name:      "multiple matches returns first",
			slice:     []string{"go", "java", "go", "rust"},
			predicate: func(v string) bool { return len(v) == 2 },
			expected:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.IndexFunc(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("IndexFunc(%v) = %d; want %d",
					tt.slice, result, tt.expected)
			}
		})
	}
}

func TestIndexFunc_Struct(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	tests := []struct {
		name      string
		slice     []User
		predicate func(User) bool
		expected  int
	}{
		{
			name: "match by ID",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{3, "Charlie"},
			},
			predicate: func(u User) bool {
				return u.ID == 2
			},
			expected: 1,
		},
		{
			name: "no match",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{3, "Charlie"},
			},
			predicate: func(u User) bool {
				return u.Name == "David"
			},
			expected: -1,
		},
		{
			name:      "empty slice",
			slice:     []User{},
			predicate: func(u User) bool { return true },
			expected:  -1,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(u User) bool { return true },
			expected:  -1,
		},
		{
			name: "multiple matches returns first",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{3, "Alice"},
			},
			predicate: func(u User) bool {
				return u.Name == "Alice"
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.IndexFunc(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("IndexFunc(%v) = %d; want %d",
					tt.slice, result, tt.expected)
			}
		})
	}
}

func TestLastIndex_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		elem     int
		expected int
	}{
		{
			name:     "element at beginning",
			slice:    []int{10, 20, 30},
			elem:     10,
			expected: 0,
		},
		{
			name:     "element in middle",
			slice:    []int{10, 20, 30},
			elem:     20,
			expected: 1,
		},
		{
			name:     "element at end",
			slice:    []int{10, 20, 30},
			elem:     30,
			expected: 2,
		},
		{
			name:     "element not found",
			slice:    []int{10, 20, 30},
			elem:     40,
			expected: -1,
		},
		{
			name:     "empty slice",
			slice:    []int{},
			elem:     1,
			expected: -1,
		},
		{
			name:     "nil slice",
			slice:    nil,
			elem:     1,
			expected: -1,
		},
		{
			name:     "duplicate elements returns last index",
			slice:    []int{5, 7, 5, 9},
			elem:     5,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.LastIndex(tt.slice, tt.elem)
			if result != tt.expected {
				t.Errorf("LastIndex(%v, %v) = %d; want %d",
					tt.slice, tt.elem, result, tt.expected)
			}
		})
	}
}

func TestLastIndex_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		elem     string
		expected int
	}{
		{
			name:     "element at beginning",
			slice:    []string{"apple", "banana", "cherry"},
			elem:     "apple",
			expected: 0,
		},
		{
			name:     "element in middle",
			slice:    []string{"apple", "banana", "cherry"},
			elem:     "banana",
			expected: 1,
		},
		{
			name:     "element at end",
			slice:    []string{"apple", "banana", "cherry"},
			elem:     "cherry",
			expected: 2,
		},
		{
			name:     "element not found",
			slice:    []string{"apple", "banana", "cherry"},
			elem:     "orange",
			expected: -1,
		},
		{
			name:     "empty slice",
			slice:    []string{},
			elem:     "apple",
			expected: -1,
		},
		{
			name:     "nil slice",
			slice:    nil,
			elem:     "apple",
			expected: -1,
		},
		{
			name:     "duplicate elements returns last index",
			slice:    []string{"go", "java", "go", "rust"},
			elem:     "go",
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.LastIndex(tt.slice, tt.elem)
			if result != tt.expected {
				t.Errorf("LastIndex(%v, %q) = %d; want %d",
					tt.slice, tt.elem, result, tt.expected)
			}
		})
	}
}

func TestLastIndex_Struct(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}
	tests := []struct {
		name     string
		slice    []User
		elem     User
		expected int
	}{
		{
			name: "element at beginning",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{3, "Charlie"},
			},
			elem:     User{1, "Alice"},
			expected: 0,
		},
		{
			name: "element in middle",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{3, "Charlie"},
			},
			elem:     User{2, "Bob"},
			expected: 1,
		},
		{
			name: "element at end",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{3, "Charlie"},
			},
			elem:     User{3, "Charlie"},
			expected: 2,
		},
		{
			name: "element not found",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{3, "Charlie"},
			},
			elem:     User{4, "David"},
			expected: -1,
		},
		{
			name:     "empty slice",
			slice:    []User{},
			elem:     User{1, "Alice"},
			expected: -1,
		},
		{
			name:     "nil slice",
			slice:    nil,
			elem:     User{1, "Alice"},
			expected: -1,
		},
		{
			name: "duplicate elements returns last index",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{1, "Alice"},
			},
			elem:     User{1, "Alice"},
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.LastIndex(tt.slice, tt.elem)
			if result != tt.expected {
				t.Errorf("LastIndex(%v, %+v) = %d; want %d",
					tt.slice, tt.elem, result, tt.expected)
			}
		})
	}
}

func TestLastIndexFunc_Int(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		expected  int
	}{
		{
			name:      "match in middle",
			slice:     []int{1, 3, 4, 7},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  2,
		},
		{
			name:      "match at beginning",
			slice:     []int{2, 3, 5},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  0,
		},
		{
			name:      "match at end",
			slice:     []int{1, 3, 6},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  2,
		},
		{
			name:      "no match",
			slice:     []int{1, 3, 5},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  -1,
		},
		{
			name:      "empty slice",
			slice:     []int{},
			predicate: func(v int) bool { return true },
			expected:  -1,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(v int) bool { return true },
			expected:  -1,
		},
		{
			name:      "multiple matches returns last",
			slice:     []int{1, 4, 6, 8},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.LastIndexFunc(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("LastIndexFunc(%v) = %d; want %d",
					tt.slice, result, tt.expected)
			}
		})
	}
}

func TestLastIndexFunc_String(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		predicate func(string) bool
		expected  int
	}{
		{
			name:      "match in middle",
			slice:     []string{"apple", "banana", "cherry"},
			predicate: func(v string) bool { return v == "banana" },
			expected:  1,
		},
		{
			name:      "match at beginning",
			slice:     []string{"apple", "banana", "cherry"},
			predicate: func(v string) bool { return v == "apple" },
			expected:  0,
		},
		{
			name:      "match at end",
			slice:     []string{"apple", "banana", "cherry"},
			predicate: func(v string) bool { return v == "cherry" },
			expected:  2,
		},
		{
			name:      "no match",
			slice:     []string{"apple", "banana", "cherry"},
			predicate: func(v string) bool { return len(v) == 4 },
			expected:  -1,
		},
		{
			name:      "empty slice",
			slice:     []string{},
			predicate: func(v string) bool { return true },
			expected:  -1,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(v string) bool { return true },
			expected:  -1,
		},
		{
			name:      "multiple matches returns last",
			slice:     []string{"go", "java", "go", "rust"},
			predicate: func(v string) bool { return v == "go" },
			expected:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.LastIndexFunc(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("LastIndexFunc(%v) = %d; want %d",
					tt.slice, result, tt.expected)
			}
		})
	}
}

func TestLastIndexFunc_Struct(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}
	tests := []struct {
		name      string
		slice     []User
		predicate func(User) bool
		expected  int
	}{
		{
			name: "match by ID",
			slice: []User{
				{1, "Alice"},
			},
			predicate: func(u User) bool {
				return u.ID == 1
			},
			expected: 0,
		},
		{
			name: "no match",
			slice: []User{
				{1, "Alice"},
			},
			predicate: func(u User) bool {
				return u.Name == "Bob"
			},
			expected: -1,
		},
		{
			name:      "empty slice",
			slice:     []User{},
			predicate: func(u User) bool { return true },
			expected:  -1,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(u User) bool { return true },
			expected:  -1,
		},
		{
			name: "multiple matches returns last",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{3, "Alice"},
			},
			predicate: func(u User) bool {
				return u.Name == "Alice"
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.LastIndexFunc(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("LastIndexFunc(%v) = %d; want %d",
					tt.slice, result, tt.expected)
			}
		})
	}
}

func TestMinIndex_Int(t *testing.T) {
	tests := []struct {
		name        string
		slice       []int
		expectedIdx int
		expectedVal int
		expectedOk  bool
	}{
		{
			name:        "min at beginning",
			slice:       []int{1, 2, 3},
			expectedIdx: 0,
			expectedVal: 1,
			expectedOk:  true,
		},
		{
			name:        "min in middle",
			slice:       []int{3, 1, 2},
			expectedIdx: 1,
			expectedVal: 1,
			expectedOk:  true,
		},
		{
			name:        "min at end",
			slice:       []int{3, 2, 1},
			expectedIdx: 2,
			expectedVal: 1,
			expectedOk:  true,
		},
		{
			name:        "duplicate minima returns first index",
			slice:       []int{2, 1, 1, 3},
			expectedIdx: 1,
			expectedVal: 1,
			expectedOk:  true,
		},
		{
			name:        "empty slice",
			slice:       []int{},
			expectedIdx: -1,
			expectedVal: 0,
			expectedOk:  false,
		},
		{
			name:        "nil slice",
			slice:       nil,
			expectedIdx: -1,
			expectedVal: 0,
			expectedOk:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx, ok := lxslices.MinIndex(tt.slice)
			if ok != tt.expectedOk {
				t.Errorf("MinIndex() ok = %v; want %v", ok, tt.expectedOk)
				return
			}
			if ok {
				if idx != tt.expectedIdx {
					t.Errorf("MinIndex() idx = %v; want %v", idx, tt.expectedIdx)
					return
				}
				if tt.slice[idx] != tt.expectedVal {
					t.Errorf("MinIndex() value = %v; want %v", tt.slice[idx], tt.expectedVal)
				}
			} else {
				if idx != -1 {
					t.Errorf("MinIndex() idx = %v; want -1 for empty slice", idx)
				}
			}
		})
	}
}

func TestMinIndex_String(t *testing.T) {
	tests := []struct {
		name        string
		slice       []string
		expectedIdx int
		expectedVal string
		expectedOk  bool
	}{
		{"min string at beginning", []string{"a", "b", "c"}, 0, "a", true},
		{"min string in middle", []string{"c", "a", "b"}, 1, "a", true},
		{"duplicate minima returns first index", []string{"b", "a", "a"}, 1, "a", true},
		{"empty slice", []string{}, -1, "", false},
		{"nil slice", nil, -1, "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx, ok := lxslices.MinIndex(tt.slice)
			if ok != tt.expectedOk {
				t.Errorf("MinIndex() ok = %v; want %v", ok, tt.expectedOk)
				return
			}
			if ok {
				if idx != tt.expectedIdx {
					t.Errorf("MinIndex() idx = %v; want %v", idx, tt.expectedIdx)
					return
				}
				if tt.slice[idx] != tt.expectedVal {
					t.Errorf("MinIndex() value = %v; want %v", tt.slice[idx], tt.expectedVal)
				}
			} else {
				if idx != -1 {
					t.Errorf("MinIndex() idx = %v; want -1 for empty slice", idx)
				}
			}
		})
	}
}

func TestMinIndex_Float64(t *testing.T) {
	tests := []struct {
		name        string
		slice       []float64
		expectedIdx int
		expectedVal float64
		expectedOk  bool
		expectNaN   bool
	}{
		{"min at beginning", []float64{1.0, 2.0, 3.0}, 0, 1.0, true, false},
		{"min in middle", []float64{3.0, 1.0, 2.0}, 1, 1.0, true, false},
		{"min at end", []float64{3.0, 2.0, 1.0}, 2, 1.0, true, false},
		{"duplicate minima returns first index", []float64{2.0, 1.0, 1.0, 3.0}, 1, 1.0, true, false},
		{"empty slice", []float64{}, -1, 0.0, false, false},
		{"nil slice", nil, -1, 0.0, false, false},
		{"NaN first returns NaN", []float64{math.NaN(), 1.0, 2.0}, 0, 0.0, true, true},
		{"NaN later ignored", []float64{1.0, math.NaN(), 2.0}, 0, 1.0, true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx, ok := lxslices.MinIndex(tt.slice)
			if ok != tt.expectedOk {
				t.Errorf("MinIndex() ok = %v; want %v", ok, tt.expectedOk)
				return
			}
			if ok {
				if idx != tt.expectedIdx {
					t.Errorf("MinIndex() idx = %v; want %v", idx, tt.expectedIdx)
					return
				}
				val := tt.slice[idx]
				if tt.expectNaN {
					if !math.IsNaN(val) {
						t.Errorf("MinIndex() value = %v; want NaN", val)
					}
				} else {
					if val != tt.expectedVal {
						t.Errorf("MinIndex() value = %v; want %v", val, tt.expectedVal)
					}
				}
			} else {
				if idx != -1 {
					t.Errorf("MinIndex() idx = %v; want -1 for empty slice", idx)
				}
			}
		})
	}
}

func TestMaxIndex_Int(t *testing.T) {
	tests := []struct {
		name        string
		slice       []int
		expectedIdx int
		expectedVal int
		expectedOk  bool
	}{
		{
			name:        "max at beginning",
			slice:       []int{3, 2, 1},
			expectedIdx: 0,
			expectedVal: 3,
			expectedOk:  true,
		},
		{
			name:        "max in middle",
			slice:       []int{1, 3, 2},
			expectedIdx: 1,
			expectedVal: 3,
			expectedOk:  true,
		},
		{
			name:        "max at end",
			slice:       []int{1, 2, 5},
			expectedIdx: 2,
			expectedVal: 5,
			expectedOk:  true,
		},
		{
			name:        "duplicate maxima returns first index",
			slice:       []int{5, 3, 5, 2},
			expectedIdx: 0,
			expectedVal: 5,
			expectedOk:  true,
		},
		{
			name:        "empty slice",
			slice:       []int{},
			expectedIdx: -1,
			expectedVal: 0,
			expectedOk:  false,
		},
		{
			name:        "nil slice",
			slice:       nil,
			expectedIdx: -1,
			expectedVal: 0,
			expectedOk:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx, ok := lxslices.MaxIndex(tt.slice)
			if ok != tt.expectedOk {
				t.Errorf("MaxIndex() ok = %v; want %v", ok, tt.expectedOk)
				return
			}
			if ok {
				if idx != tt.expectedIdx {
					t.Errorf("MaxIndex() idx = %v; want %v", idx, tt.expectedIdx)
					return
				}
				if tt.slice[idx] != tt.expectedVal {
					t.Errorf("MaxIndex() value = %v; want %v", tt.slice[idx], tt.expectedVal)
				}
			} else {
				if idx != -1 {
					t.Errorf("MaxIndex() idx = %v; want -1 for empty slice", idx)
				}
			}
		})
	}
}

func TestMaxIndex_Float64(t *testing.T) {
	tests := []struct {
		name        string
		slice       []float64
		expectedIdx int
		expectedVal float64
		expectedOk  bool
		expectNaN   bool
	}{
		{"max at beginning", []float64{3.0, 2.0, 1.0}, 0, 3.0, true, false},
		{"max in middle", []float64{1.0, 3.0, 2.0}, 1, 3.0, true, false},
		{"max at end", []float64{1.0, 2.0, 5.0}, 2, 5.0, true, false},
		{"duplicate maxima returns first index", []float64{5.0, 3.0, 5.0, 2.0}, 0, 5.0, true, false},
		{"empty slice", []float64{}, -1, 0.0, false, false},
		{"nil slice", nil, -1, 0.0, false, false},
		{"NaN first returns NaN", []float64{math.NaN(), 1.0, 2.0}, 0, 0.0, true, true},
		{"NaN later ignored", []float64{1.0, math.NaN(), 2.0}, 2, 2.0, true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx, ok := lxslices.MaxIndex(tt.slice)
			if ok != tt.expectedOk {
				t.Errorf("MaxIndex() ok = %v; want %v", ok, tt.expectedOk)
				return
			}
			if ok {
				if idx != tt.expectedIdx {
					t.Errorf("MaxIndex() idx = %v; want %v", idx, tt.expectedIdx)
					return
				}
				val := tt.slice[idx]
				if tt.expectNaN {
					if !math.IsNaN(val) {
						t.Errorf("MaxIndex() value = %v; want NaN", val)
					}
				} else {
					if val != tt.expectedVal {
						t.Errorf("MaxIndex() value = %v; want %v", val, tt.expectedVal)
					}
				}
			} else {
				if idx != -1 {
					t.Errorf("MaxIndex() idx = %v; want -1 for empty slice", idx)
				}
			}
		})
	}
}

func TestMaxIndex_String(t *testing.T) {
	tests := []struct {
		name        string
		slice       []string
		expectedIdx int
		expectedVal string
		expectedOk  bool
	}{
		{"max string at beginning", []string{"z", "y", "x"}, 0, "z", true},
		{"max string in middle", []string{"a", "z", "b"}, 1, "z", true},
		{"duplicate maxima returns first index", []string{"z", "a", "z"}, 0, "z", true},
		{"empty slice", []string{}, -1, "", false},
		{"nil slice", nil, -1, "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx, ok := lxslices.MaxIndex(tt.slice)
			if ok != tt.expectedOk {
				t.Errorf("MaxIndex() ok = %v; want %v", ok, tt.expectedOk)
				return
			}
			if ok {
				if idx != tt.expectedIdx {
					t.Errorf("MaxIndex() idx = %v; want %v", idx, tt.expectedIdx)
					return
				}
				if tt.slice[idx] != tt.expectedVal {
					t.Errorf("MaxIndex() value = %v; want %v", tt.slice[idx], tt.expectedVal)
				}
			} else {
				if idx != -1 {
					t.Errorf("MaxIndex() idx = %v; want -1 for empty slice", idx)
				}
			}
		})
	}
}

func TestFirst(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var s []int
			v, ok := lxslices.First(s)
			if ok || v != 0 {
				t.Fatalf("First(nil) = (%v, %v); want (0, false)", v, ok)
			}
		})

		t.Run("empty", func(t *testing.T) {
			s := []int{}
			v, ok := lxslices.First(s)
			if ok || v != 0 {
				t.Fatalf("First(empty) = (%v, %v); want (0, false)", v, ok)
			}
		})

		t.Run("non-empty", func(t *testing.T) {
			s := []int{7, 8, 9}
			v, ok := lxslices.First(s)
			if !ok || v != 7 {
				t.Fatalf("First(%v) = (%v, %v); want (7, true)", s, v, ok)
			}
		})
	})

	t.Run("string", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var s []string
			v, ok := lxslices.First(s)
			if ok || v != "" {
				t.Fatalf("First(nil) = (%v, %v); want (\"\", false)", v, ok)
			}
		})

		t.Run("empty", func(t *testing.T) {
			s := []string{}
			v, ok := lxslices.First(s)
			if ok || v != "" {
				t.Fatalf("First(empty) = (%v, %v); want (\"\", false)", v, ok)
			}
		})

		t.Run("non-empty", func(t *testing.T) {
			s := []string{"a", "b"}
			v, ok := lxslices.First(s)
			if !ok || v != "a" {
				t.Fatalf("First(%v) = (%v, %v); want (\"a\", true)", s, v, ok)
			}
		})
	})

	t.Run("struct", func(t *testing.T) {
		type P struct{ A int }

		t.Run("nil", func(t *testing.T) {
			var s []P
			v, ok := lxslices.First(s)
			if ok || v != (P{}) {
				t.Fatalf("First(nil) = (%v, %v); want (%v, false)", v, ok, P{})
			}
		})

		t.Run("empty", func(t *testing.T) {
			s := []P{}
			v, ok := lxslices.First(s)
			if ok || v != (P{}) {
				t.Fatalf("First(empty) = (%v, %v); want (%v, false)", v, ok, P{})
			}
		})

		t.Run("non-empty", func(t *testing.T) {
			s := []P{{1}, {2}}
			v, ok := lxslices.First(s)
			if !ok || v != (P{1}) {
				t.Fatalf("First(%v) = (%v, %v); want (%v, true)", s, v, ok, P{1})
			}
		})
	})
}

func TestLast(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var s []int
			v, ok := lxslices.Last(s)
			if ok || v != 0 {
				t.Fatalf("Last(nil) = (%v, %v); want (0, false)", v, ok)
			}
		})

		t.Run("empty", func(t *testing.T) {
			s := []int{}
			v, ok := lxslices.Last(s)
			if ok || v != 0 {
				t.Fatalf("Last(empty) = (%v, %v); want (0, false)", v, ok)
			}
		})

		t.Run("non-empty", func(t *testing.T) {
			s := []int{7, 8, 9}
			v, ok := lxslices.Last(s)
			if !ok || v != 9 {
				t.Fatalf("Last(%v) = (%v, %v); want (9, true)", s, v, ok)
			}
		})
	})

	t.Run("string", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			var s []string
			v, ok := lxslices.Last(s)
			if ok || v != "" {
				t.Fatalf("Last(nil) = (%v, %v); want (\"\", false)", v, ok)
			}
		})

		t.Run("empty", func(t *testing.T) {
			s := []string{}
			v, ok := lxslices.Last(s)
			if ok || v != "" {
				t.Fatalf("Last(empty) = (%v, %v); want (\"\", false)", v, ok)
			}
		})

		t.Run("non-empty", func(t *testing.T) {
			s := []string{"a", "b"}
			v, ok := lxslices.Last(s)
			if !ok || v != "b" {
				t.Fatalf("Last(%v) = (%v, %v); want (\"b\", true)", s, v, ok)
			}
		})
	})

	t.Run("struct", func(t *testing.T) {
		type P struct{ A int }

		t.Run("nil", func(t *testing.T) {
			var s []P
			v, ok := lxslices.Last(s)
			if ok || v != (P{}) {
				t.Fatalf("Last(nil) = (%v, %v); want (%v, false)", v, ok, P{})
			}
		})

		t.Run("empty", func(t *testing.T) {
			s := []P{}
			v, ok := lxslices.Last(s)
			if ok || v != (P{}) {
				t.Fatalf("Last(empty) = (%v, %v); want (%v, false)", v, ok, P{})
			}
		})

		t.Run("non-empty", func(t *testing.T) {
			s := []P{{1}, {2}}
			v, ok := lxslices.Last(s)
			if !ok || v != (P{2}) {
				t.Fatalf("Last(%v) = (%v, %v); want (%v, true)", s, v, ok, P{2})
			}
		})
	})
}

func TestGet(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		t.Run("in range", func(t *testing.T) {
			s := []int{10, 11, 12}
			v, ok := lxslices.Get(s, 1)
			if !ok || v != 11 {
				t.Fatalf("Get(%v, 1) = (%v, %v); want (11, true)", s, v, ok)
			}
		})

		t.Run("out of range negative", func(t *testing.T) {
			s := []int{10, 11, 12}
			v, ok := lxslices.Get(s, -1)
			if ok || v != 0 {
				t.Fatalf("Get(%v, -1) = (%v, %v); want (0, false)", s, v, ok)
			}
		})

		t.Run("out of range too large", func(t *testing.T) {
			s := []int{10, 11, 12}
			v, ok := lxslices.Get(s, 3)
			if ok || v != 0 {
				t.Fatalf("Get(%v, 3) = (%v, %v); want (0, false)", s, v, ok)
			}
		})
	})

	t.Run("string", func(t *testing.T) {
		t.Run("in range", func(t *testing.T) {
			s := []string{"x", "y", "z"}
			v, ok := lxslices.Get(s, 2)
			if !ok || v != "z" {
				t.Fatalf("Get(%v, 2) = (%v, %v); want (\"z\", true)", s, v, ok)
			}
		})

		t.Run("out of range negative", func(t *testing.T) {
			s := []string{"x", "y", "z"}
			v, ok := lxslices.Get(s, -1)
			if ok || v != "" {
				t.Fatalf("Get(%v, -1) = (%v, %v); want (\"\", false)", s, v, ok)
			}
		})

		t.Run("out of range too large", func(t *testing.T) {
			s := []string{"x", "y", "z"}
			v, ok := lxslices.Get(s, 3)
			if ok || v != "" {
				t.Fatalf("Get(%v, 3) = (%v, %v); want (\"\", false)", s, v, ok)
			}
		})
	})

	t.Run("struct", func(t *testing.T) {
		type P struct{ A int }

		t.Run("in range", func(t *testing.T) {
			s := []P{{5}, {6}, {7}}
			v, ok := lxslices.Get(s, 1)
			if !ok || v != (P{6}) {
				t.Fatalf("Get(%v, 1) = (%v, %v); want (%v, true)", s, v, ok, P{6})
			}
		})

		t.Run("out of range negative", func(t *testing.T) {
			s := []P{{5}, {6}, {7}}
			v, ok := lxslices.Get(s, -1)
			if ok || v != (P{}) {
				t.Fatalf("Get(%v, -1) = (%v, %v); want (%v, false)", s, v, ok, P{})
			}
		})

		t.Run("out of range too large", func(t *testing.T) {
			s := []P{{5}, {6}, {7}}
			v, ok := lxslices.Get(s, 3)
			if ok || v != (P{}) {
				t.Fatalf("Get(%v, 3) = (%v, %v); want (%v, false)", s, v, ok, P{})
			}
		})
	})
}

func TestBinarySearch_Int(t *testing.T) {
	tests := []struct {
		name          string
		slice         []int
		target        int
		expectedIndex int
	}{
		{
			name:          "find element at beginning",
			slice:         []int{1, 3, 5, 7, 9},
			target:        1,
			expectedIndex: 0,
		},
		{
			name:          "find element in middle",
			slice:         []int{1, 3, 5, 7, 9},
			target:        5,
			expectedIndex: 2,
		},
		{
			name:          "find element at end",
			slice:         []int{1, 3, 5, 7, 9},
			target:        9,
			expectedIndex: 4,
		},
		{
			name:          "element not found - less than min",
			slice:         []int{1, 3, 5, 7, 9},
			target:        0,
			expectedIndex: -1,
		},
		{
			name:          "element not found - greater than max",
			slice:         []int{1, 3, 5, 7, 9},
			target:        10,
			expectedIndex: -1,
		},
		{
			name:          "element not found - in range",
			slice:         []int{1, 3, 5, 7, 9},
			target:        4,
			expectedIndex: -1,
		},
		{
			name:          "single element - found",
			slice:         []int{42},
			target:        42,
			expectedIndex: 0,
		},
		{
			name:          "single element - not found",
			slice:         []int{42},
			target:        1,
			expectedIndex: -1,
		},
		{
			name:          "two elements - find first",
			slice:         []int{1, 2},
			target:        1,
			expectedIndex: 0,
		},
		{
			name:          "two elements - find second",
			slice:         []int{1, 2},
			target:        2,
			expectedIndex: 1,
		},
		{
			name:          "empty slice",
			slice:         []int{},
			target:        1,
			expectedIndex: -1,
		},
		{
			name:          "nil slice",
			slice:         nil,
			target:        1,
			expectedIndex: -1,
		},
		{
			name:          "large sorted slice",
			slice:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			target:        8,
			expectedIndex: 7,
		},
		{
			name:          "negative numbers",
			slice:         []int{-10, -5, 0, 5, 10},
			target:        -5,
			expectedIndex: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx := lxslices.BinarySearch(tt.slice, tt.target)
			if idx != tt.expectedIndex {
				t.Errorf("BinarySearch(%v, %d) = %d; want %d",
					tt.slice, tt.target, idx, tt.expectedIndex)
			}
		})
	}
}

func TestBinarySearch_String(t *testing.T) {
	tests := []struct {
		name          string
		slice         []string
		target        string
		expectedIndex int
	}{
		{
			name:          "find element at beginning",
			slice:         []string{"apple", "banana", "cherry", "date", "elderberry"},
			target:        "apple",
			expectedIndex: 0,
		},
		{
			name:          "find element in middle",
			slice:         []string{"apple", "banana", "cherry", "date", "elderberry"},
			target:        "cherry",
			expectedIndex: 2,
		},
		{
			name:          "find element at end",
			slice:         []string{"apple", "banana", "cherry", "date", "elderberry"},
			target:        "elderberry",
			expectedIndex: 4,
		},
		{
			name:          "element not found - less than min",
			slice:         []string{"apple", "banana", "cherry", "date", "elderberry"},
			target:        "aardvark",
			expectedIndex: -1,
		},
		{
			name:          "element not found - greater than max",
			slice:         []string{"apple", "banana", "cherry", "date", "elderberry"},
			target:        "zebra",
			expectedIndex: -1,
		},
		{
			name:          "element not found - in range",
			slice:         []string{"alpha", "beta", "gamma"},
			target:        "delta",
			expectedIndex: -1,
		},
		{
			name:          "single element - found",
			slice:         []string{"hello"},
			target:        "hello",
			expectedIndex: 0,
		},
		{
			name:          "single element - not found",
			slice:         []string{"hello"},
			target:        "world",
			expectedIndex: -1,
		},
		{
			name:          "two elements - find first",
			slice:         []string{"alpha", "beta"},
			target:        "alpha",
			expectedIndex: 0,
		},
		{
			name:          "two elements - find second",
			slice:         []string{"alpha", "beta"},
			target:        "beta",
			expectedIndex: 1,
		},
		{
			name:          "empty slice",
			slice:         []string{},
			target:        "test",
			expectedIndex: -1,
		},
		{
			name:          "nil slice",
			slice:         nil,
			target:        "test",
			expectedIndex: -1,
		},
		{
			name:          "large sorted slice",
			slice:         []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o"},
			target:        "h",
			expectedIndex: 7,
		},
		{
			name:          "size equals slice length",
			slice:         []string{"alpha", "beta", "gamma"},
			target:        "gamma",
			expectedIndex: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx := lxslices.BinarySearch(tt.slice, tt.target)
			if idx != tt.expectedIndex {
				t.Errorf("BinarySearch(%v, %q) = %d; want %d",
					tt.slice, tt.target, idx, tt.expectedIndex)
			}
		})
	}
}

func TestBinarySearch_Float(t *testing.T) {
	tests := []struct {
		name          string
		slice         []float64
		target        float64
		expectedIndex int
	}{
		{
			name:          "find element at beginning",
			slice:         []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			target:        1.1,
			expectedIndex: 0,
		},
		{
			name:          "find element in middle",
			slice:         []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			target:        3.3,
			expectedIndex: 2,
		},
		{
			name:          "find element at end",
			slice:         []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			target:        5.5,
			expectedIndex: 4,
		},
		{
			name:          "element not found - less than min",
			slice:         []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			target:        0.5,
			expectedIndex: -1,
		},
		{
			name:          "element not found - greater than max",
			slice:         []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			target:        6.0,
			expectedIndex: -1,
		},
		{
			name:          "element not found - in range",
			slice:         []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			target:        3.0,
			expectedIndex: -1,
		},
		{
			name:          "single element - found",
			slice:         []float64{42.5},
			target:        42.5,
			expectedIndex: 0,
		},
		{
			name:          "single element - not found",
			slice:         []float64{42.5},
			target:        1.0,
			expectedIndex: -1,
		},
		{
			name:          "two elements - find first",
			slice:         []float64{1.5, 2.5},
			target:        1.5,
			expectedIndex: 0,
		},
		{
			name:          "two elements - find second",
			slice:         []float64{1.5, 2.5},
			target:        2.5,
			expectedIndex: 1,
		},
		{
			name:          "empty slice",
			slice:         []float64{},
			target:        1.0,
			expectedIndex: -1,
		},
		{
			name:          "nil slice",
			slice:         nil,
			target:        1.0,
			expectedIndex: -1,
		},
		{
			name:          "large sorted slice",
			slice:         []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1, 11.1, 12.2, 13.3, 14.4, 15.5},
			target:        8.8,
			expectedIndex: 7,
		},
		{
			name:          "negative floats",
			slice:         []float64{-10.5, -5.5, 0.0, 5.5, 10.5},
			target:        -5.5,
			expectedIndex: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx := lxslices.BinarySearch(tt.slice, tt.target)
			if idx != tt.expectedIndex {
				t.Errorf("BinarySearch(%v, %f) = %d; want %d",
					tt.slice, tt.target, idx, tt.expectedIndex)
			}
		})
	}
}

func TestBinarySearchFunc_Struct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
		{"Diana", 40},
		{"Eve", 45},
	}

	// Comparator for searching by Age
	ageComparator := func(target, elem Person) int {
		if target.Age < elem.Age {
			return -1
		}
		if target.Age > elem.Age {
			return 1
		}
		return 0
	}

	tests := []struct {
		name          string
		slice         []Person
		target        Person
		expectedIndex int
	}{
		{
			name:          "find person by age - middle",
			slice:         people,
			target:        Person{"", 35},
			expectedIndex: 2,
		},
		{
			name:          "find person by age - beginning",
			slice:         people,
			target:        Person{"", 25},
			expectedIndex: 0,
		},
		{
			name:          "find person by age - end",
			slice:         people,
			target:        Person{"", 45},
			expectedIndex: 4,
		},
		{
			name:          "person not found - age too low",
			slice:         people,
			target:        Person{"", 20},
			expectedIndex: -1,
		},
		{
			name:          "person not found - age too high",
			slice:         people,
			target:        Person{"", 50},
			expectedIndex: -1,
		},
		{
			name:          "person not found - age in range",
			slice:         people,
			target:        Person{"", 33},
			expectedIndex: -1,
		},
		{
			name:          "empty slice",
			slice:         []Person{},
			target:        Person{"", 30},
			expectedIndex: -1,
		},
		{
			name:          "nil slice",
			slice:         nil,
			target:        Person{"", 30},
			expectedIndex: -1,
		},
		{
			name:          "single element - found",
			slice:         []Person{{"Bob", 30}},
			target:        Person{"", 30},
			expectedIndex: 0,
		},
		{
			name:          "single element - not found",
			slice:         []Person{{"Bob", 30}},
			target:        Person{"", 25},
			expectedIndex: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx := lxslices.BinarySearchFunc(tt.slice, tt.target, ageComparator)
			if idx != tt.expectedIndex {
				t.Errorf("BinarySearchFunc(%v, %v, ageComparator) = %d; want %d",
					tt.slice, tt.target, idx, tt.expectedIndex)
			}
		})
	}
}

func TestBinarySearchFunc_IntWithCustomComparator(t *testing.T) {
	// Test with descending order
	descendingInts := []int{90, 70, 50, 30, 10}

	// Comparator for descending order
	descendingComparator := func(target, elem int) int {
		if target > elem {
			return -1
		}
		if target < elem {
			return 1
		}
		return 0
	}

	tests := []struct {
		name          string
		slice         []int
		target        int
		expectedIndex int
	}{
		{
			name:          "find in descending order - middle",
			slice:         descendingInts,
			target:        50,
			expectedIndex: 2,
		},
		{
			name:          "find in descending order - beginning",
			slice:         descendingInts,
			target:        90,
			expectedIndex: 0,
		},
		{
			name:          "find in descending order - end",
			slice:         descendingInts,
			target:        10,
			expectedIndex: 4,
		},
		{
			name:          "not found in descending order",
			slice:         descendingInts,
			target:        40,
			expectedIndex: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx := lxslices.BinarySearchFunc(tt.slice, tt.target, descendingComparator)
			if idx != tt.expectedIndex {
				t.Errorf("BinarySearchFunc(%v, %d, descendingComparator) = %d; want %d",
					tt.slice, tt.target, idx, tt.expectedIndex)
			}
		})
	}
}

func TestBinarySearchFunc_StringWithReverseComparator(t *testing.T) {
	// Test with reverse alphabetical order
	reverseStrings := []string{"zebra", "yankee", "xray", "whiskey", "alpha"}

	// Comparator for reverse alphabetical order
	reverseComparator := func(target, elem string) int {
		if target > elem {
			return -1
		}
		if target < elem {
			return 1
		}
		return 0
	}

	tests := []struct {
		name          string
		slice         []string
		target        string
		expectedIndex int
	}{
		{
			name:          "find in reverse order",
			slice:         reverseStrings,
			target:        "xray",
			expectedIndex: 2,
		},
		{
			name:          "find at beginning",
			slice:         reverseStrings,
			target:        "zebra",
			expectedIndex: 0,
		},
		{
			name:          "find at end",
			slice:         reverseStrings,
			target:        "alpha",
			expectedIndex: 4,
		},
		{
			name:          "not found",
			slice:         reverseStrings,
			target:        "bravo",
			expectedIndex: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx := lxslices.BinarySearchFunc(tt.slice, tt.target, reverseComparator)
			if idx != tt.expectedIndex {
				t.Errorf("BinarySearchFunc(%v, %q, reverseComparator) = %d; want %d",
					tt.slice, tt.target, idx, tt.expectedIndex)
			}
		})
	}
}
