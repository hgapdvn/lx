package lxmaps_test

import (
	"strings"
	"testing"

	"github.com/hgapdvn/lx/maps"
	"github.com/hgapdvn/lx/slices"
)

func TestKeysBy_StringInt(t *testing.T) {
	tests := []struct {
		name      string
		m         map[string]int
		predicate func(string, int) bool
		expected  []string
	}{
		{
			name:      "nil map",
			m:         nil,
			predicate: func(k string, v int) bool { return v > 1 },
			expected:  []string{},
		},
		{
			name:      "empty map",
			m:         map[string]int{},
			predicate: func(k string, v int) bool { return v > 1 },
			expected:  []string{},
		},
		{
			name:      "single entry matches",
			m:         map[string]int{"a": 2},
			predicate: func(k string, v int) bool { return v > 1 },
			expected:  []string{"a"},
		},
		{
			name:      "single entry does not match",
			m:         map[string]int{"a": 1},
			predicate: func(k string, v int) bool { return v > 1 },
			expected:  []string{},
		},
		{
			name:      "multiple entries some match by value",
			m:         map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(k string, v int) bool { return v > 1 },
			expected:  []string{"b", "c"},
		},
		{
			name:      "multiple entries none match",
			m:         map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(k string, v int) bool { return v > 10 },
			expected:  []string{},
		},
		{
			name:      "multiple entries all match",
			m:         map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(k string, v int) bool { return v > 0 },
			expected:  []string{"a", "b", "c"},
		},
		{
			name:      "predicate uses key",
			m:         map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(k string, v int) bool { return k == "b" },
			expected:  []string{"b"},
		},
		{
			name:      "predicate uses both key and value",
			m:         map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(k string, v int) bool { return k == "a" && v == 1 },
			expected:  []string{"a"},
		},
		{
			name:      "predicate checks key contains substring",
			m:         map[string]int{"apple": 1, "app": 2, "application": 3, "banana": 4},
			predicate: func(k string, v int) bool { return strings.Contains(k, "app") },
			expected:  []string{"apple", "app", "application"},
		},
		{
			name:      "predicate checks value equals key length",
			m:         map[string]int{"a": 1, "bb": 2, "ccc": 3, "dddd": 4},
			predicate: func(k string, v int) bool { return len(k) == v },
			expected:  []string{"a", "bb", "ccc", "dddd"},
		},
		{
			name:      "predicate with negative values",
			m:         map[string]int{"neg": -1, "zero": 0, "pos": 1},
			predicate: func(k string, v int) bool { return v < 0 },
			expected:  []string{"neg"},
		},
		{
			name:      "predicate with zero value",
			m:         map[string]int{"a": 0, "b": 1, "c": 0},
			predicate: func(k string, v int) bool { return v == 0 },
			expected:  []string{"a", "c"},
		},
		{
			name:      "unicode keys matching by value",
			m:         map[string]int{"こんにちは": 5, "世界": 2, "test": 4},
			predicate: func(k string, v int) bool { return v > 3 },
			expected:  []string{"こんにちは", "test"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.KeysBy(tt.m, tt.predicate)

			if got == nil && len(tt.expected) == 0 {
				// both are empty, that's fine
				return
			}

			if len(got) != len(tt.expected) {
				t.Errorf("KeysBy() length = %d, want %d", len(got), len(tt.expected))
				return
			}

			// Order is not guaranteed, so use ContainsAll
			if !lxslices.ContainsAll(got, tt.expected...) {
				t.Errorf("KeysBy() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestKeysBy_StringString(t *testing.T) {
	tests := []struct {
		name      string
		m         map[string]string
		predicate func(string, string) bool
		expected  []string
	}{
		{
			name:      "nil map",
			m:         nil,
			predicate: func(k, v string) bool { return len(v) > 3 },
			expected:  []string{},
		},
		{
			name:      "empty map",
			m:         map[string]string{},
			predicate: func(k, v string) bool { return len(v) > 3 },
			expected:  []string{},
		},
		{
			name:      "single entry matches by value length",
			m:         map[string]string{"a": "hello"},
			predicate: func(k, v string) bool { return len(v) > 3 },
			expected:  []string{"a"},
		},
		{
			name:      "multiple entries filter by value length",
			m:         map[string]string{"a": "hi", "b": "hello", "c": "go", "d": "world"},
			predicate: func(k, v string) bool { return len(v) > 3 },
			expected:  []string{"b", "d"},
		},
		{
			name:      "predicate checks exact value",
			m:         map[string]string{"a": "apple", "b": "banana", "c": "cherry"},
			predicate: func(k, v string) bool { return v == "banana" },
			expected:  []string{"b"},
		},
		{
			name:      "predicate checks value prefix",
			m:         map[string]string{"a": "apple", "b": "apricot", "c": "banana"},
			predicate: func(k, v string) bool { return strings.HasPrefix(v, "a") },
			expected:  []string{"a", "b"},
		},
		{
			name:      "predicate checks key equals value",
			m:         map[string]string{"hello": "hello", "world": "test", "go": "go"},
			predicate: func(k, v string) bool { return k == v },
			expected:  []string{"hello", "go"},
		},
		{
			name:      "predicate case insensitive",
			m:         map[string]string{"a": "Hello", "b": "WORLD", "c": "go"},
			predicate: func(k, v string) bool { return strings.ToLower(v) == "hello" },
			expected:  []string{"a"},
		},
		{
			name:      "empty string value",
			m:         map[string]string{"a": "", "b": "text", "c": ""},
			predicate: func(k, v string) bool { return v == "" },
			expected:  []string{"a", "c"},
		},
		{
			name:      "unicode values",
			m:         map[string]string{"jp": "こんにちは", "cn": "你好", "en": "hello"},
			predicate: func(k, v string) bool { return len(v) > 10 },
			expected:  []string{"jp"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.KeysBy(tt.m, tt.predicate)

			if got == nil && len(tt.expected) == 0 {
				return
			}

			if len(got) != len(tt.expected) {
				t.Errorf("KeysBy() length = %d, want %d", len(got), len(tt.expected))
				return
			}

			if !lxslices.ContainsAll(got, tt.expected...) {
				t.Errorf("KeysBy() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestKeysBy_IntString(t *testing.T) {
	tests := []struct {
		name      string
		m         map[int]string
		predicate func(int, string) bool
		expected  []int
	}{
		{
			name:      "nil map",
			m:         nil,
			predicate: func(k int, v string) bool { return len(v) > 2 },
			expected:  []int{},
		},
		{
			name:      "empty map",
			m:         map[int]string{},
			predicate: func(k int, v string) bool { return len(v) > 2 },
			expected:  []int{},
		},
		{
			name:      "single entry matches",
			m:         map[int]string{1: "hello"},
			predicate: func(k int, v string) bool { return len(v) > 3 },
			expected:  []int{1},
		},
		{
			name:      "multiple entries filter by key",
			m:         map[int]string{1: "a", 2: "b", 3: "c"},
			predicate: func(k int, v string) bool { return k > 1 },
			expected:  []int{2, 3},
		},
		{
			name:      "multiple entries filter by value",
			m:         map[int]string{1: "apple", 2: "go", 3: "banana"},
			predicate: func(k int, v string) bool { return len(v) > 3 },
			expected:  []int{1, 3},
		},
		{
			name:      "predicate uses both key and value",
			m:         map[int]string{1: "one", 2: "two", 3: "three"},
			predicate: func(k int, v string) bool { return k == 2 && v == "two" },
			expected:  []int{2},
		},
		{
			name:      "negative keys",
			m:         map[int]string{-1: "neg", 0: "zero", 1: "pos"},
			predicate: func(k int, v string) bool { return k < 0 },
			expected:  []int{-1},
		},
		{
			name:      "zero key",
			m:         map[int]string{0: "zero", 1: "one"},
			predicate: func(k int, v string) bool { return k == 0 },
			expected:  []int{0},
		},
		{
			name:      "large keys",
			m:         map[int]string{1000000: "million", 999: "small"},
			predicate: func(k int, v string) bool { return k > 10000 },
			expected:  []int{1000000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.KeysBy(tt.m, tt.predicate)

			if got == nil && len(tt.expected) == 0 {
				return
			}

			if len(got) != len(tt.expected) {
				t.Errorf("KeysBy() length = %d, want %d", len(got), len(tt.expected))
				return
			}

			if !lxslices.ContainsAll(got, tt.expected...) {
				t.Errorf("KeysBy() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestKeysBy_StringStruct(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	tests := []struct {
		name      string
		m         map[string]User
		predicate func(string, User) bool
		expected  []string
	}{
		{
			name:      "nil map",
			m:         nil,
			predicate: func(k string, u User) bool { return u.Age > 18 },
			expected:  []string{},
		},
		{
			name:      "empty map",
			m:         map[string]User{},
			predicate: func(k string, u User) bool { return u.Age > 18 },
			expected:  []string{},
		},
		{
			name: "single entry matches by age",
			m: map[string]User{
				"alice": {Name: "Alice", Age: 25},
			},
			predicate: func(k string, u User) bool { return u.Age > 18 },
			expected:  []string{"alice"},
		},
		{
			name: "single entry does not match",
			m: map[string]User{
				"bob": {Name: "Bob", Age: 16},
			},
			predicate: func(k string, u User) bool { return u.Age > 18 },
			expected:  []string{},
		},
		{
			name: "multiple entries filter by age",
			m: map[string]User{
				"alice":   {Name: "Alice", Age: 25},
				"bob":     {Name: "Bob", Age: 16},
				"charlie": {Name: "Charlie", Age: 30},
			},
			predicate: func(k string, u User) bool { return u.Age > 18 },
			expected:  []string{"alice", "charlie"},
		},
		{
			name: "predicate filters by name",
			m: map[string]User{
				"user1": {Name: "Alice", Age: 25},
				"user2": {Name: "Bob", Age: 30},
				"user3": {Name: "Alice", Age: 28},
			},
			predicate: func(k string, u User) bool { return u.Name == "Alice" },
			expected:  []string{"user1", "user3"},
		},
		{
			name: "predicate uses key and struct field",
			m: map[string]User{
				"admin": {Name: "Alice", Age: 30},
				"user":  {Name: "Bob", Age: 25},
				"guest": {Name: "Charlie", Age: 20},
			},
			predicate: func(k string, u User) bool { return k == "admin" && u.Age > 25 },
			expected:  []string{"admin"},
		},
		{
			name: "predicate checks name length",
			m: map[string]User{
				"a": {Name: "Al", Age: 20},
				"b": {Name: "Alexander", Age: 25},
				"c": {Name: "Bob", Age: 30},
			},
			predicate: func(k string, u User) bool { return len(u.Name) > 3 },
			expected:  []string{"b"},
		},
		{
			name: "predicate checks age range",
			m: map[string]User{
				"young":  {Name: "Teen", Age: 15},
				"adult":  {Name: "Adult", Age: 25},
				"senior": {Name: "Elder", Age: 65},
			},
			predicate: func(k string, u User) bool { return u.Age >= 18 && u.Age < 60 },
			expected:  []string{"adult"},
		},
		{
			name: "predicate with zero age",
			m: map[string]User{
				"a": {Name: "Unknown", Age: 0},
				"b": {Name: "Known", Age: 20},
			},
			predicate: func(k string, u User) bool { return u.Age == 0 },
			expected:  []string{"a"},
		},
		{
			name: "all entries match",
			m: map[string]User{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
			},
			predicate: func(k string, u User) bool { return u.Age > 0 },
			expected:  []string{"alice", "bob"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.KeysBy(tt.m, tt.predicate)

			if got == nil && len(tt.expected) == 0 {
				return
			}

			if len(got) != len(tt.expected) {
				t.Errorf("KeysBy() length = %d, want %d", len(got), len(tt.expected))
				return
			}

			if !lxslices.ContainsAll(got, tt.expected...) {
				t.Errorf("KeysBy() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestKeysBy_StringComplexStruct(t *testing.T) {
	type Address struct {
		City    string
		Country string
	}

	type Person struct {
		Name    string
		Age     int
		Address Address
	}

	tests := []struct {
		name      string
		m         map[string]Person
		predicate func(string, Person) bool
		expected  []string
	}{
		{
			name:      "nil map",
			m:         nil,
			predicate: func(k string, p Person) bool { return p.Address.City == "NYC" },
			expected:  []string{},
		},
		{
			name: "filter by nested city",
			m: map[string]Person{
				"john": {Name: "John", Age: 30, Address: Address{City: "NYC", Country: "USA"}},
				"jane": {Name: "Jane", Age: 28, Address: Address{City: "LA", Country: "USA"}},
				"bob":  {Name: "Bob", Age: 35, Address: Address{City: "NYC", Country: "USA"}},
			},
			predicate: func(k string, p Person) bool { return p.Address.City == "NYC" },
			expected:  []string{"john", "bob"},
		},
		{
			name: "filter by age and country",
			m: map[string]Person{
				"john": {Name: "John", Age: 30, Address: Address{City: "NYC", Country: "USA"}},
				"jane": {Name: "Jane", Age: 28, Address: Address{City: "London", Country: "UK"}},
				"bob":  {Name: "Bob", Age: 35, Address: Address{City: "Paris", Country: "France"}},
			},
			predicate: func(k string, p Person) bool { return p.Age > 28 && p.Address.Country == "USA" },
			expected:  []string{"john"},
		},
		{
			name: "filter by name and city length",
			m: map[string]Person{
				"a": {Name: "Alice", Age: 25, Address: Address{City: "NYC", Country: "USA"}},
				"b": {Name: "Bob", Age: 30, Address: Address{City: "LA", Country: "USA"}},
				"c": {Name: "Charlie", Age: 35, Address: Address{City: "London", Country: "UK"}},
			},
			predicate: func(k string, p Person) bool { return len(p.Name) > 4 && len(p.Address.City) > 2 },
			expected:  []string{"a", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.KeysBy(tt.m, tt.predicate)

			if got == nil && len(tt.expected) == 0 {
				return
			}

			if len(got) != len(tt.expected) {
				t.Errorf("KeysBy() length = %d, want %d", len(got), len(tt.expected))
				return
			}

			if !lxslices.ContainsAll(got, tt.expected...) {
				t.Errorf("KeysBy() = %v, want %v", got, tt.expected)
			}
		})
	}
}
