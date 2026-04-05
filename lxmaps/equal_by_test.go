package lxmaps_test

import (
	"strings"
	"testing"

	lxmaps "github.com/hgapdvn/lx/lxmaps"
)

func TestEqualBy_StringInt(t *testing.T) {
	tests := []struct {
		name     string
		m1       map[string]int
		m2       map[string]int
		eq       func(v1, v2 int) bool
		expected bool
	}{
		{
			name:     "both nil",
			m1:       nil,
			m2:       nil,
			eq:       func(v1, v2 int) bool { return v1 == v2 },
			expected: true,
		},
		{
			name:     "nil vs empty",
			m1:       nil,
			m2:       map[string]int{},
			eq:       func(v1, v2 int) bool { return v1 == v2 },
			expected: false,
		},
		{
			name:     "empty vs nil",
			m1:       map[string]int{},
			m2:       nil,
			eq:       func(v1, v2 int) bool { return v1 == v2 },
			expected: false,
		},
		{
			name:     "both empty",
			m1:       map[string]int{},
			m2:       map[string]int{},
			eq:       func(v1, v2 int) bool { return v1 == v2 },
			expected: true,
		},
		{
			name:     "single entry equal",
			m1:       map[string]int{"a": 1},
			m2:       map[string]int{"a": 1},
			eq:       func(v1, v2 int) bool { return v1 == v2 },
			expected: true,
		},
		{
			name:     "single entry not equal",
			m1:       map[string]int{"a": 1},
			m2:       map[string]int{"a": 2},
			eq:       func(v1, v2 int) bool { return v1 == v2 },
			expected: false,
		},
		{
			name:     "multiple entries equal",
			m1:       map[string]int{"a": 1, "b": 2, "c": 3},
			m2:       map[string]int{"c": 3, "a": 1, "b": 2},
			eq:       func(v1, v2 int) bool { return v1 == v2 },
			expected: true,
		},
		{
			name:     "multiple entries different length",
			m1:       map[string]int{"a": 1, "b": 2},
			m2:       map[string]int{"a": 1, "b": 2, "c": 3},
			eq:       func(v1, v2 int) bool { return v1 == v2 },
			expected: false,
		},
		{
			name:     "missing key in m2",
			m1:       map[string]int{"a": 1, "b": 2},
			m2:       map[string]int{"a": 1},
			eq:       func(v1, v2 int) bool { return v1 == v2 },
			expected: false,
		},
		{
			name:     "with custom equality - absolute diff",
			m1:       map[string]int{"a": 10, "b": 20},
			m2:       map[string]int{"a": 12, "b": 22},
			eq:       func(v1, v2 int) bool { return (v1-v2) < 5 && (v2-v1) < 5 },
			expected: true,
		},
		{
			name:     "with custom equality - not matching",
			m1:       map[string]int{"a": 10, "b": 20},
			m2:       map[string]int{"a": 20, "b": 20},
			eq:       func(v1, v2 int) bool { return (v1-v2) < 5 && (v2-v1) < 5 },
			expected: false,
		},
		{
			name:     "zero values equal",
			m1:       map[string]int{"z": 0},
			m2:       map[string]int{"z": 0},
			eq:       func(v1, v2 int) bool { return v1 == v2 },
			expected: true,
		},
		{
			name:     "negative values equal",
			m1:       map[string]int{"x": -1, "y": -2},
			m2:       map[string]int{"y": -2, "x": -1},
			eq:       func(v1, v2 int) bool { return v1 == v2 },
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.EqualBy(tt.m1, tt.m2, tt.eq)
			if got != tt.expected {
				t.Errorf("EqualBy() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestEqualBy_StringString(t *testing.T) {
	tests := []struct {
		name     string
		m1       map[string]string
		m2       map[string]string
		eq       func(v1, v2 string) bool
		expected bool
	}{
		{
			name:     "both nil",
			m1:       nil,
			m2:       nil,
			eq:       func(v1, v2 string) bool { return v1 == v2 },
			expected: true,
		},
		{
			name:     "both empty",
			m1:       map[string]string{},
			m2:       map[string]string{},
			eq:       func(v1, v2 string) bool { return v1 == v2 },
			expected: true,
		},
		{
			name:     "single entry equal",
			m1:       map[string]string{"a": "hello"},
			m2:       map[string]string{"a": "hello"},
			eq:       func(v1, v2 string) bool { return v1 == v2 },
			expected: true,
		},
		{
			name:     "single entry not equal",
			m1:       map[string]string{"a": "hello"},
			m2:       map[string]string{"a": "world"},
			eq:       func(v1, v2 string) bool { return v1 == v2 },
			expected: false,
		},
		{
			name:     "multiple entries equal",
			m1:       map[string]string{"a": "hello", "b": "world"},
			m2:       map[string]string{"b": "world", "a": "hello"},
			eq:       func(v1, v2 string) bool { return v1 == v2 },
			expected: true,
		},
		{
			name:     "case insensitive equality",
			m1:       map[string]string{"a": "Hello", "b": "World"},
			m2:       map[string]string{"b": "world", "a": "hello"},
			eq:       func(v1, v2 string) bool { return strings.EqualFold(v1, v2) },
			expected: true,
		},
		{
			name:     "case insensitive not matching",
			m1:       map[string]string{"a": "Hello"},
			m2:       map[string]string{"a": "Goodbye"},
			eq:       func(v1, v2 string) bool { return strings.EqualFold(v1, v2) },
			expected: false,
		},
		{
			name:     "length based equality",
			m1:       map[string]string{"a": "abc", "b": "defg"},
			m2:       map[string]string{"a": "xyz", "b": "wxyz"},
			eq:       func(v1, v2 string) bool { return len(v1) == len(v2) },
			expected: true,
		},
		{
			name:     "empty string values equal",
			m1:       map[string]string{"a": "", "b": "text"},
			m2:       map[string]string{"a": "", "b": "text"},
			eq:       func(v1, v2 string) bool { return v1 == v2 },
			expected: true,
		},
		{
			name:     "unicode strings equal",
			m1:       map[string]string{"a": "こんにちは", "b": "世界"},
			m2:       map[string]string{"b": "世界", "a": "こんにちは"},
			eq:       func(v1, v2 string) bool { return v1 == v2 },
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.EqualBy(tt.m1, tt.m2, tt.eq)
			if got != tt.expected {
				t.Errorf("EqualBy() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestEqualBy_StringStruct(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	tests := []struct {
		name     string
		m1       map[string]User
		m2       map[string]User
		eq       func(u1, u2 User) bool
		expected bool
	}{
		{
			name:     "both nil",
			m1:       nil,
			m2:       nil,
			eq:       func(u1, u2 User) bool { return u1 == u2 },
			expected: true,
		},
		{
			name:     "nil vs empty",
			m1:       nil,
			m2:       map[string]User{},
			eq:       func(u1, u2 User) bool { return u1 == u2 },
			expected: false,
		},
		{
			name:     "both empty",
			m1:       map[string]User{},
			m2:       map[string]User{},
			eq:       func(u1, u2 User) bool { return u1 == u2 },
			expected: true,
		},
		{
			name: "single entry struct equal",
			m1: map[string]User{
				"alice": {Name: "Alice", Age: 25},
			},
			m2: map[string]User{
				"alice": {Name: "Alice", Age: 25},
			},
			eq:       func(u1, u2 User) bool { return u1 == u2 },
			expected: true,
		},
		{
			name: "single entry struct different name",
			m1: map[string]User{
				"alice": {Name: "Alice", Age: 25},
			},
			m2: map[string]User{
				"alice": {Name: "Alicia", Age: 25},
			},
			eq:       func(u1, u2 User) bool { return u1 == u2 },
			expected: false,
		},
		{
			name: "single entry struct different age",
			m1: map[string]User{
				"alice": {Name: "Alice", Age: 25},
			},
			m2: map[string]User{
				"alice": {Name: "Alice", Age: 26},
			},
			eq:       func(u1, u2 User) bool { return u1 == u2 },
			expected: false,
		},
		{
			name: "multiple entries struct equal",
			m1: map[string]User{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
			},
			m2: map[string]User{
				"bob":   {Name: "Bob", Age: 30},
				"alice": {Name: "Alice", Age: 25},
			},
			eq:       func(u1, u2 User) bool { return u1 == u2 },
			expected: true,
		},
		{
			name: "multiple entries one struct different",
			m1: map[string]User{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
			},
			m2: map[string]User{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 31},
			},
			eq:       func(u1, u2 User) bool { return u1 == u2 },
			expected: false,
		},
		{
			name: "zero value struct",
			m1: map[string]User{
				"empty": {},
			},
			m2: map[string]User{
				"empty": {},
			},
			eq:       func(u1, u2 User) bool { return u1 == u2 },
			expected: true,
		},
		{
			name: "zero value vs populated",
			m1: map[string]User{
				"empty": {},
			},
			m2: map[string]User{
				"empty": {Name: "X", Age: 0},
			},
			eq:       func(u1, u2 User) bool { return u1 == u2 },
			expected: false,
		},
		{
			name: "custom equality - age tolerance",
			m1: map[string]User{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
			},
			m2: map[string]User{
				"alice": {Name: "Alice", Age: 26},
				"bob":   {Name: "Bob", Age: 31},
			},
			eq: func(u1, u2 User) bool {
				return u1.Name == u2.Name && (u1.Age-u2.Age < 2 && u2.Age-u1.Age < 2)
			},
			expected: true,
		},
		{
			name: "custom equality - name only",
			m1: map[string]User{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
			},
			m2: map[string]User{
				"alice": {Name: "Alice", Age: 100},
				"bob":   {Name: "Bob", Age: 50},
			},
			eq: func(u1, u2 User) bool {
				return u1.Name == u2.Name
			},
			expected: true,
		},
		{
			name: "custom equality - name only not matching",
			m1: map[string]User{
				"alice": {Name: "Alice", Age: 25},
			},
			m2: map[string]User{
				"alice": {Name: "Alicia", Age: 25},
			},
			eq: func(u1, u2 User) bool {
				return u1.Name == u2.Name
			},
			expected: false,
		},
		{
			name: "empty name equal",
			m1: map[string]User{
				"a": {Name: "", Age: 20},
			},
			m2: map[string]User{
				"a": {Name: "", Age: 20},
			},
			eq:       func(u1, u2 User) bool { return u1 == u2 },
			expected: true,
		},
		{
			name: "case insensitive name comparison",
			m1: map[string]User{
				"alice": {Name: "ALICE", Age: 25},
				"bob":   {Name: "BOB", Age: 30},
			},
			m2: map[string]User{
				"bob":   {Name: "bob", Age: 30},
				"alice": {Name: "alice", Age: 25},
			},
			eq: func(u1, u2 User) bool {
				return strings.EqualFold(u1.Name, u2.Name) && u1.Age == u2.Age
			},
			expected: true,
		},
		{
			name: "missing key in m2",
			m1: map[string]User{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
			},
			m2: map[string]User{
				"alice": {Name: "Alice", Age: 25},
			},
			eq:       func(u1, u2 User) bool { return u1 == u2 },
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.EqualBy(tt.m1, tt.m2, tt.eq)
			if got != tt.expected {
				t.Errorf("EqualBy() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestEqualBy_StringComplexStruct(t *testing.T) {
	type Address struct {
		Street string
		City   string
	}

	type Person struct {
		Name    string
		Age     int
		Address Address
	}

	tests := []struct {
		name     string
		m1       map[string]Person
		m2       map[string]Person
		eq       func(p1, p2 Person) bool
		expected bool
	}{
		{
			name: "complex struct equal",
			m1: map[string]Person{
				"john": {
					Name: "John",
					Age:  30,
					Address: Address{
						Street: "123 Main St",
						City:   "New York",
					},
				},
			},
			m2: map[string]Person{
				"john": {
					Name: "John",
					Age:  30,
					Address: Address{
						Street: "123 Main St",
						City:   "New York",
					},
				},
			},
			eq:       func(p1, p2 Person) bool { return p1 == p2 },
			expected: true,
		},
		{
			name: "complex struct different nested field",
			m1: map[string]Person{
				"john": {
					Name: "John",
					Age:  30,
					Address: Address{
						Street: "123 Main St",
						City:   "New York",
					},
				},
			},
			m2: map[string]Person{
				"john": {
					Name: "John",
					Age:  30,
					Address: Address{
						Street: "123 Main St",
						City:   "Boston",
					},
				},
			},
			eq:       func(p1, p2 Person) bool { return p1 == p2 },
			expected: false,
		},
		{
			name: "multiple complex structs equal",
			m1: map[string]Person{
				"john": {
					Name: "John",
					Age:  30,
					Address: Address{
						Street: "123 Main St",
						City:   "New York",
					},
				},
				"jane": {
					Name: "Jane",
					Age:  28,
					Address: Address{
						Street: "456 Oak Ave",
						City:   "Boston",
					},
				},
			},
			m2: map[string]Person{
				"jane": {
					Name: "Jane",
					Age:  28,
					Address: Address{
						Street: "456 Oak Ave",
						City:   "Boston",
					},
				},
				"john": {
					Name: "John",
					Age:  30,
					Address: Address{
						Street: "123 Main St",
						City:   "New York",
					},
				},
			},
			eq:       func(p1, p2 Person) bool { return p1 == p2 },
			expected: true,
		},
		{
			name: "custom equality - compare name and city only",
			m1: map[string]Person{
				"john": {
					Name: "John",
					Age:  30,
					Address: Address{
						Street: "123 Main St",
						City:   "New York",
					},
				},
			},
			m2: map[string]Person{
				"john": {
					Name: "John",
					Age:  35,
					Address: Address{
						Street: "789 Pine Rd",
						City:   "New York",
					},
				},
			},
			eq: func(p1, p2 Person) bool {
				return p1.Name == p2.Name && p1.Address.City == p2.Address.City
			},
			expected: true,
		},
		{
			name: "custom equality city mismatch",
			m1: map[string]Person{
				"john": {
					Name: "John",
					Age:  30,
					Address: Address{
						Street: "123 Main St",
						City:   "New York",
					},
				},
			},
			m2: map[string]Person{
				"john": {
					Name: "John",
					Age:  30,
					Address: Address{
						Street: "456 Oak Ave",
						City:   "Boston",
					},
				},
			},
			eq: func(p1, p2 Person) bool {
				return p1.Name == p2.Name && p1.Address.City == p2.Address.City
			},
			expected: false,
		},
		{
			name: "zero value complex struct",
			m1: map[string]Person{
				"empty": {},
			},
			m2: map[string]Person{
				"empty": {},
			},
			eq:       func(p1, p2 Person) bool { return p1 == p2 },
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.EqualBy(tt.m1, tt.m2, tt.eq)
			if got != tt.expected {
				t.Errorf("EqualBy() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestEqualBy_IntSlice(t *testing.T) {
	tests := []struct {
		name     string
		m1       map[int][]int
		m2       map[int][]int
		eq       func(s1, s2 []int) bool
		expected bool
	}{
		{
			name:     "both nil",
			m1:       nil,
			m2:       nil,
			eq:       func(s1, s2 []int) bool { return len(s1) == len(s2) },
			expected: true,
		},
		{
			name: "slices equal length",
			m1: map[int][]int{
				1: {1, 2, 3},
			},
			m2: map[int][]int{
				1: {4, 5, 6},
			},
			eq: func(s1, s2 []int) bool {
				if len(s1) != len(s2) {
					return false
				}
				return len(s1) == 3
			},
			expected: true,
		},
		{
			name: "slices different length",
			m1: map[int][]int{
				1: {1, 2, 3},
			},
			m2: map[int][]int{
				1: {1, 2},
			},
			eq: func(s1, s2 []int) bool {
				return len(s1) == len(s2)
			},
			expected: false,
		},
		{
			name: "slices with same elements",
			m1: map[int][]int{
				1: {1, 2, 3},
				2: {4, 5},
			},
			m2: map[int][]int{
				2: {4, 5},
				1: {1, 2, 3},
			},
			eq: func(s1, s2 []int) bool {
				if len(s1) != len(s2) {
					return false
				}
				for i := range s1 {
					if s1[i] != s2[i] {
						return false
					}
				}
				return true
			},
			expected: true,
		},
		{
			name: "nil slices equal",
			m1: map[int][]int{
				1: nil,
			},
			m2: map[int][]int{
				1: nil,
			},
			eq: func(s1, s2 []int) bool {
				if s1 == nil && s2 == nil {
					return true
				}
				if len(s1) != len(s2) {
					return false
				}
				for i := range s1 {
					if s1[i] != s2[i] {
						return false
					}
				}
				return true
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.EqualBy(tt.m1, tt.m2, tt.eq)
			if got != tt.expected {
				t.Errorf("EqualBy() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
