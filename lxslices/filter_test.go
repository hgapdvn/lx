package lxslices_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/nthanhhai2909/lx/lxslices"
)

func TestFind_Int(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		expected  struct {
			value int
			found bool
		}
	}{
		{
			name:      "find at beginning",
			slice:     []int{1, 2, 3},
			predicate: func(v int) bool { return v == 1 },
			expected: struct {
				value int
				found bool
			}{1, true},
		},
		{
			name:      "find in the middle",
			slice:     []int{1, 2, 3},
			predicate: func(v int) bool { return v == 2 },
			expected: struct {
				value int
				found bool
			}{2, true},
		},
		{
			name:  "find at the end",
			slice: []int{1, 2, 3},
			predicate: func(v int) bool {
				return v == 3
			},
			expected: struct {
				value int
				found bool
			}{3, true},
		},
		{
			name:  "duplicate values return the first one",
			slice: []int{1, 2, 3, 3, 3},
			predicate: func(v int) bool {
				return v == 3
			},
			expected: struct {
				value int
				found bool
			}{3, true},
		},
		{
			name:      "no match",
			slice:     []int{1, 2, 3},
			predicate: func(v int) bool { return v == 4 },
			expected: struct {
				value int
				found bool
			}{0, false},
		},
		{
			name:      "empty slice",
			slice:     []int{},
			predicate: func(v int) bool { return v == 0 },
			expected: struct {
				value int
				found bool
			}{0, false},
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(v int) bool { return v == 0 },
			expected: struct {
				value int
				found bool
			}{0, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := lxslices.Find(tt.slice, tt.predicate)
			if value != tt.expected.value || found != tt.expected.found {
				t.Errorf("Find() = (%v, %v); want (%v, %v)", value, found, tt.expected.value, tt.expected.found)
			}
		})
	}
}

func TestFind_String(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		predicate func(string) bool
		expected  struct {
			value string
			found bool
		}
	}{
		{
			name:  "find at beginning",
			slice: []string{"a", "b", "c"},
			predicate: func(s string) bool {
				return s == "a"
			},
			expected: struct {
				value string
				found bool
			}{"a", true},
		},
		{
			name:  "find in the middle",
			slice: []string{"a", "b", "c"},
			predicate: func(s string) bool {
				return s == "b"
			},
			expected: struct {
				value string
				found bool
			}{"b", true},
		},
		{
			name:  "find at the end",
			slice: []string{"a", "b", "c"},
			predicate: func(s string) bool {
				return s == "c"
			},
			expected: struct {
				value string
				found bool
			}{"c", true},
		},
		{
			name:  "duplicate values return the first one",
			slice: []string{"a", "b", "c", "c", "c"},
			predicate: func(s string) bool {
				return s == "c"
			},
			expected: struct {
				value string
				found bool
			}{"c", true},
		},
		{
			name:  "not found",
			slice: []string{"a", "b", "c"},
			predicate: func(s string) bool {
				return s == "d"
			},
			expected: struct {
				value string
				found bool
			}{"", false},
		},
		{
			name:  "empty slice",
			slice: []string{},
			predicate: func(s string) bool {
				return s == "a"
			},
			expected: struct {
				value string
				found bool
			}{"", false},
		},
		{
			name:  "nil slice",
			slice: nil,
			predicate: func(s string) bool {
				return s == "a"
			},
			expected: struct {
				value string
				found bool
			}{value: "", found: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := lxslices.Find(tt.slice, tt.predicate)
			if value != tt.expected.value || found != tt.expected.found {
				t.Errorf("Find() = (%v, %v); want (%v, %v)", value, found, tt.expected.value, tt.expected.found)
			}
		})
	}
}

func TestFind_Struct(t *testing.T) {
	type User struct {
		ID     int
		Name   string
		Active bool
	}
	tests := []struct {
		name      string
		slice     []User
		predicate func(User) bool
		expected  struct {
			value User
			found bool
		}
	}{
		{
			name: "find first active user",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", false},
				{3, "Charlie", true},
			},
			predicate: func(u User) bool { return u.Active },
			expected: struct {
				value User
				found bool
			}{User{1, "Alice", true}, true},
		},
		{
			name: "find first user with name starting with 'A'",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", false},
				{3, "Andrew", true},
			},
			predicate: func(u User) bool { return strings.HasPrefix(u.Name, "A") },
			expected: struct {
				value User
				found bool
			}{User{1, "Alice", true}, true},
		},
		{
			name: "not found user by id",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", false},
				{3, "Charlie", true},
			},
			predicate: func(u User) bool { return u.ID == 4 },
			expected: struct {
				value User
				found bool
			}{User{}, false},
		},
		{
			name: "not found user by name",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", false},
				{3, "Charlie", true},
			},
			predicate: func(user User) bool {
				return user.Name == "Dave"
			},
		},
		{
			name:      "empty slice",
			slice:     []User{},
			predicate: func(u User) bool { return u.Active },
			expected: struct {
				value User
				found bool
			}{User{}, false},
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(u User) bool { return u.Active },
			expected: struct {
				value User
				found bool
			}{User{}, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := lxslices.Find(tt.slice, tt.predicate)
			if value != tt.expected.value || found != tt.expected.found {
				t.Errorf("Find() = (%v, %v); want (%v, %v)", value, found, tt.expected.value, tt.expected.found)
			}
		})
	}
}

func TestFindLast_Int(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		expected  struct {
			value int
			found bool
		}
	}{
		{
			name:      "find at beginning",
			slice:     []int{1, 2, 3},
			predicate: func(v int) bool { return v == 1 },
			expected: struct {
				value int
				found bool
			}{1, true},
		},
		{
			name:      "find in the middle",
			slice:     []int{1, 2, 3},
			predicate: func(v int) bool { return v == 2 },
			expected: struct {
				value int
				found bool
			}{2, true},
		},
		{
			name:  "find at the end",
			slice: []int{1, 2, 3},
			predicate: func(v int) bool {
				return v == 3
			},
			expected: struct {
				value int
				found bool
			}{3, true},
		},
		{
			name:      "last even number",
			slice:     []int{1, 2, 3, 4, 6},
			predicate: func(v int) bool { return v%2 == 0 },
			expected: struct {
				value int
				found bool
			}{6, true},
		},
		{
			name:      "no match",
			slice:     []int{1, 2, 3},
			predicate: func(v int) bool { return v == 4 },
			expected: struct {
				value int
				found bool
			}{0, false},
		},
		{
			name:      "empty slice",
			slice:     []int{},
			predicate: func(v int) bool { return v == 0 },
			expected: struct {
				value int
				found bool
			}{0, false},
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(v int) bool { return v == 0 },
			expected: struct {
				value int
				found bool
			}{0, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := lxslices.FindLast(tt.slice, tt.predicate)
			if value != tt.expected.value || found != tt.expected.found {
				t.Errorf("FindLast() = (%v, %v); want (%v, %v)", value, found, tt.expected.value, tt.expected.found)
			}
		})
	}
}

func TestFindLast_String(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		predicate func(string) bool
		expected  struct {
			value string
			found bool
		}
	}{
		{
			name:  "find at beginning",
			slice: []string{"a", "b", "c"},
			predicate: func(s string) bool {
				return s == "a"
			},
			expected: struct {
				value string
				found bool
			}{"a", true},
		},
		{
			name:  "find in the middle",
			slice: []string{"a", "b", "c"},
			predicate: func(s string) bool {
				return s == "b"
			},
			expected: struct {
				value string
				found bool
			}{"b", true},
		},
		{
			name:  "find at the end",
			slice: []string{"a", "b", "c"},
			predicate: func(s string) bool {
				return s == "c"
			},
			expected: struct {
				value string
				found bool
			}{"c", true},
		},
		{
			name:      "multiple matches return last",
			slice:     []string{"apple", "banana", "apricot", "avocado"},
			predicate: func(s string) bool { return strings.HasPrefix(s, "a") },
			expected: struct {
				value string
				found bool
			}{"avocado", true},
		},
		{
			name:      "not found",
			slice:     []string{"a", "b", "c"},
			predicate: func(s string) bool { return s == "d" },
			expected: struct {
				value string
				found bool
			}{"", false},
		},
		{
			name:      "empty slice",
			slice:     []string{},
			predicate: func(s string) bool { return s == "a" },
			expected: struct {
				value string
				found bool
			}{"", false},
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(s string) bool { return s == "a" },
			expected: struct {
				value string
				found bool
			}{"", false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := lxslices.FindLast(tt.slice, tt.predicate)
			if value != tt.expected.value || found != tt.expected.found {
				t.Errorf("FindLast() = (%v, %v); want (%v, %v)", value, found, tt.expected.value, tt.expected.found)
			}
		})
	}
}

func TestFindLast_Struct(t *testing.T) {
	type User struct {
		ID     int
		Name   string
		Active bool
	}
	tests := []struct {
		name      string
		slice     []User
		predicate func(User) bool
		expected  struct {
			value User
			found bool
		}
	}{
		{
			name: "last active user",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", false},
				{3, "Charlie", true},
			},
			predicate: func(u User) bool { return u.Active },
			expected: struct {
				value User
				found bool
			}{User{3, "Charlie", true}, true},
		},
		{
			name: "last user with name starting with 'A'",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", false},
				{3, "Andrew", true},
				{4, "Alice", true},
			},
			predicate: func(u User) bool { return strings.HasPrefix(u.Name, "A") },
			expected: struct {
				value User
				found bool
			}{User{4, "Alice", true}, true},
		},
		{
			name: "not found user by id",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", false},
				{3, "Charlie", true},
			},
			predicate: func(u User) bool { return u.ID == 4 },
			expected: struct {
				value User
				found bool
			}{User{}, false},
		},
		{
			name:      "empty slice",
			slice:     []User{},
			predicate: func(u User) bool { return u.Active },
			expected: struct {
				value User
				found bool
			}{User{}, false},
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(u User) bool { return u.Active },
			expected: struct {
				value User
				found bool
			}{User{}, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := lxslices.FindLast(tt.slice, tt.predicate)
			if value != tt.expected.value || found != tt.expected.found {
				t.Errorf("FindLast() = (%v, %v); want (%v, %v)", value, found, tt.expected.value, tt.expected.found)
			}
		})
	}
}

func TestFilter_Int(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		expected  []int
	}{
		{
			name:      "filter even numbers",
			slice:     []int{1, 2, 3, 4, 5, 6},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  []int{2, 4, 6},
		},
		{
			name:      "filter odd numbers",
			slice:     []int{1, 2, 3, 4, 5, 6},
			predicate: func(v int) bool { return v%2 != 0 },
			expected:  []int{1, 3, 5},
		},
		{
			name:      "filter greater than 3",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(v int) bool { return v > 3 },
			expected:  []int{4, 5},
		},
		{
			name:      "no matches",
			slice:     []int{1, 2, 3},
			predicate: func(v int) bool { return v > 10 },
			expected:  []int{},
		},
		{
			name:      "all match",
			slice:     []int{2, 4, 6},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  []int{2, 4, 6},
		},
		{
			name:      "empty slice",
			slice:     []int{},
			predicate: func(v int) bool { return v > 0 },
			expected:  []int{},
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(v int) bool { return v > 0 },
			expected:  []int{},
		},
		{
			name:      "predicate always true",
			slice:     []int{1, 2, 3},
			predicate: func(v int) bool { return true },
			expected:  []int{1, 2, 3},
		},
		{
			name:      "predicate always false",
			slice:     []int{1, 2, 3},
			predicate: func(v int) bool { return false },
			expected:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Filter(tt.slice, tt.predicate)
			if len(result) != len(tt.expected) {
				t.Errorf("Filter() length = %v; want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Filter() = %v; want %v", result, tt.expected)
					return
				}
			}
		})
	}
}

func TestFilter_String(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		predicate func(string) bool
		expected  []string
	}{
		{
			name:      "filter strings starting with 'a'",
			slice:     []string{"apple", "banana", "apricot", "cherry"},
			predicate: func(s string) bool { return strings.HasPrefix(s, "a") },
			expected:  []string{"apple", "apricot"},
		},
		{
			name:      "filter strings longer than 5",
			slice:     []string{"cat", "dog", "elephant", "bird"},
			predicate: func(s string) bool { return len(s) > 5 },
			expected:  []string{"elephant"},
		},
		{
			name:      "filter uppercase strings",
			slice:     []string{"HELLO", "world", "GO", "lang"},
			predicate: func(s string) bool { return s == strings.ToUpper(s) },
			expected:  []string{"HELLO", "GO"},
		},
		{
			name:      "no matches",
			slice:     []string{"a", "b", "c"},
			predicate: func(s string) bool { return len(s) > 5 },
			expected:  []string{},
		},
		{
			name:      "all match",
			slice:     []string{"a", "b", "c"},
			predicate: func(s string) bool { return len(s) == 1 },
			expected:  []string{"a", "b", "c"},
		},
		{
			name:      "empty slice",
			slice:     []string{},
			predicate: func(s string) bool { return true },
			expected:  []string{},
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(s string) bool { return true },
			expected:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Filter(tt.slice, tt.predicate)
			if len(result) != len(tt.expected) {
				t.Errorf("Filter() length = %v; want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Filter() = %v; want %v", result, tt.expected)
					return
				}
			}
		})
	}
}

func TestFilter_Struct(t *testing.T) {
	type User struct {
		ID     int
		Name   string
		Active bool
	}

	tests := []struct {
		name      string
		slice     []User
		predicate func(User) bool
		expected  []User
	}{
		{
			name: "filter active users",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", false},
				{3, "Charlie", true},
			},
			predicate: func(u User) bool { return u.Active },
			expected: []User{
				{1, "Alice", true},
				{3, "Charlie", true},
			},
		},
		{
			name: "filter by ID greater than 1",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", false},
				{3, "Charlie", true},
			},
			predicate: func(u User) bool { return u.ID > 1 },
			expected: []User{
				{2, "Bob", false},
				{3, "Charlie", true},
			},
		},
		{
			name: "filter by name starting with 'A'",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", false},
				{3, "Andrew", true},
			},
			predicate: func(u User) bool { return strings.HasPrefix(u.Name, "A") },
			expected: []User{
				{1, "Alice", true},
				{3, "Andrew", true},
			},
		},
		{
			name: "no matches",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", true},
			},
			predicate: func(u User) bool { return !u.Active },
			expected:  []User{},
		},
		{
			name:      "empty slice",
			slice:     []User{},
			predicate: func(u User) bool { return u.Active },
			expected:  []User{},
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(u User) bool { return u.Active },
			expected:  []User{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Filter(tt.slice, tt.predicate)
			if len(result) != len(tt.expected) {
				t.Errorf("Filter() length = %v; want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Filter()[%d] = %v; want %v", i, result[i], tt.expected[i])
					return
				}
			}
		})
	}
}

func TestPartition_Int(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		wantMatch []int
		wantRest  []int
	}{
		{
			name:      "integers by parity",
			slice:     []int{1, 2, 3, 4, 5, 6},
			predicate: func(v int) bool { return v%2 == 0 },
			wantMatch: []int{2, 4, 6},
			wantRest:  []int{1, 3, 5},
		},
		{
			name:      "all match",
			slice:     []int{2, 4, 6},
			predicate: func(v int) bool { return v%2 == 0 },
			wantMatch: []int{2, 4, 6},
			wantRest:  nil,
		},
		{
			name:      "none match",
			slice:     []int{1, 3, 5},
			predicate: func(v int) bool { return v%2 == 0 },
			wantMatch: nil,
			wantRest:  []int{1, 3, 5},
		},
		{
			name:      "empty slice",
			slice:     []int{},
			predicate: func(v int) bool { return true },
			wantMatch: nil,
			wantRest:  nil,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(v int) bool { return true },
			wantMatch: nil,
			wantRest:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match, rest := lxslices.Partition(tt.slice, tt.predicate)

			// Normalize empty/nil slices for DeepEqual comparison if needed,
			// though DeepEqual strictly compares nil vs empty.
			// Our function appends to nil slices initially, so it returns nil if empty.
			if len(match) == 0 && len(tt.wantMatch) == 0 {
				match = nil
				tt.wantMatch = nil
			}
			if len(rest) == 0 && len(tt.wantRest) == 0 {
				rest = nil
				tt.wantRest = nil
			}

			if !reflect.DeepEqual(match, tt.wantMatch) {
				t.Errorf("Partition() match = %v; want %v", match, tt.wantMatch)
			}
			if !reflect.DeepEqual(rest, tt.wantRest) {
				t.Errorf("Partition() rest = %v; want %v", rest, tt.wantRest)
			}
		})
	}
}

func TestPartition_String(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		predicate func(string) bool
		wantMatch []string
		wantRest  []string
	}{
		{
			name:      "strings starting with 'a'",
			slice:     []string{"apple", "banana", "apricot", "cherry"},
			predicate: func(s string) bool { return strings.HasPrefix(s, "a") },
			wantMatch: []string{"apple", "apricot"},
			wantRest:  []string{"banana", "cherry"},
		},
		{
			name:      "all match",
			slice:     []string{"a", "b", "c"},
			predicate: func(s string) bool { return len(s) == 1 },
			wantMatch: []string{"a", "b", "c"},
			wantRest:  nil,
		},
		{
			name:      "none match",
			slice:     []string{"a", "b", "c"},
			predicate: func(s string) bool { return len(s) > 5 },
			wantMatch: nil,
			wantRest:  []string{"a", "b", "c"},
		},
		{
			name:      "empty slice",
			slice:     []string{},
			predicate: func(s string) bool { return true },
			wantMatch: nil,
			wantRest:  nil,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(s string) bool { return true },
			wantMatch: nil,
			wantRest:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match, rest := lxslices.Partition(tt.slice, tt.predicate)

			if len(match) == 0 && len(tt.wantMatch) == 0 {
				match = nil
				tt.wantMatch = nil
			}
			if len(rest) == 0 && len(tt.wantRest) == 0 {
				rest = nil
				tt.wantRest = nil
			}

			if !reflect.DeepEqual(match, tt.wantMatch) {
				t.Errorf("Partition() match = %v; want %v", match, tt.wantMatch)
			}
			if !reflect.DeepEqual(rest, tt.wantRest) {
				t.Errorf("Partition() rest = %v; want %v", rest, tt.wantRest)
			}
		})
	}
}

func TestPartition_Struct(t *testing.T) {
	type User struct {
		ID     int
		Name   string
		Active bool
	}

	tests := []struct {
		name      string
		slice     []User
		predicate func(User) bool
		wantMatch []User
		wantRest  []User
	}{
		{
			name: "partition active users",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", false},
				{3, "Charlie", true},
			},
			predicate: func(u User) bool { return u.Active },
			wantMatch: []User{
				{1, "Alice", true},
				{3, "Charlie", true},
			},
			wantRest: []User{
				{2, "Bob", false},
			},
		},
		{
			name: "partition by starting letter",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", false},
				{3, "Andrew", true},
			},
			predicate: func(u User) bool { return strings.HasPrefix(u.Name, "A") },
			wantMatch: []User{
				{1, "Alice", true},
				{3, "Andrew", true},
			},
			wantRest: []User{
				{2, "Bob", false},
			},
		},
		{
			name: "none match",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", true},
			},
			predicate: func(u User) bool { return !u.Active },
			wantMatch: nil,
			wantRest: []User{
				{1, "Alice", true},
				{2, "Bob", true},
			},
		},
		{
			name:      "empty slice",
			slice:     []User{},
			predicate: func(u User) bool { return u.Active },
			wantMatch: nil,
			wantRest:  nil,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(u User) bool { return u.Active },
			wantMatch: nil,
			wantRest:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match, rest := lxslices.Partition(tt.slice, tt.predicate)

			if len(match) == 0 && len(tt.wantMatch) == 0 {
				match = nil
				tt.wantMatch = nil
			}
			if len(rest) == 0 && len(tt.wantRest) == 0 {
				rest = nil
				tt.wantRest = nil
			}

			if !reflect.DeepEqual(match, tt.wantMatch) {
				t.Errorf("Partition() match = %+v; want %+v", match, tt.wantMatch)
			}
			if !reflect.DeepEqual(rest, tt.wantRest) {
				t.Errorf("Partition() rest = %+v; want %+v", rest, tt.wantRest)
			}
		})
	}
}

func TestAny_Int(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		expected  bool
	}{
		{
			name:      "any even number",
			slice:     []int{1, 2, 3, 4, 5, 6},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  true,
		},
		{
			name:      "any odd number",
			slice:     []int{1, 2, 3, 4, 5, 6},
			predicate: func(v int) bool { return v%2 != 0 },
			expected:  true,
		},
		{
			name:      "no matches",
			slice:     []int{1, 2, 3},
			predicate: func(v int) bool { return v > 10 },
			expected:  false,
		},
		{
			name:      "all match",
			slice:     []int{2, 4, 6},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  true,
		},
		{
			name:      "empty slice",
			slice:     []int{},
			predicate: func(v int) bool { return v > 0 },
			expected:  false,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(v int) bool { return v > 0 },
			expected:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Any(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("Any() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestAny_String(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		predicate func(string) bool
		expected  bool
	}{
		{
			name:      "any strings starting with 'a'",
			slice:     []string{"apple", "banana", "apricot", "cherry"},
			predicate: func(s string) bool { return strings.HasPrefix(s, "a") },
			expected:  true,
		},
		{
			name:      "any strings longer than 5",
			slice:     []string{"cat", "dog", "elephant", "bird"},
			predicate: func(s string) bool { return len(s) > 5 },
			expected:  true,
		},
		{
			name:      "no matches",
			slice:     []string{"a", "b", "c"},
			predicate: func(s string) bool { return len(s) > 5 },
			expected:  false,
		},
		{
			name:      "all match",
			slice:     []string{"a", "b", "c"},
			predicate: func(s string) bool { return len(s) == 1 },
			expected:  true,
		},
		{
			name:      "empty slice",
			slice:     []string{},
			predicate: func(s string) bool { return true },
			expected:  false,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(s string) bool { return true },
			expected:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Any(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("Any() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestAny_Struct(t *testing.T) {
	type User struct {
		ID     int
		Name   string
		Active bool
	}
	tests := []struct {
		name      string
		slice     []User
		predicate func(User) bool
		expected  bool
	}{
		{
			name: "any active users",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", false},
				{3, "Charlie", true},
			},
			predicate: func(u User) bool { return u.Active },
			expected:  true,
		},
		{
			name: "any user with ID greater than 1",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", false},
				{3, "Charlie", true},
			},
			predicate: func(u User) bool { return u.ID > 1 },
			expected:  true,
		},
		{
			name: "no matches",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", true},
			},
			predicate: func(u User) bool { return !u.Active },
			expected:  false,
		},
		{
			name: "all match",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", true},
			},
			predicate: func(u User) bool { return u.Active },
			expected:  true,
		},
		{
			name:      "empty slice",
			slice:     []User{},
			predicate: func(u User) bool { return u.Active },
			expected:  false,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(u User) bool { return u.Active },
			expected:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Any(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("Any() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestAll_Int(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		expected  bool
	}{
		{
			name:      "all even numbers",
			slice:     []int{2, 4, 6},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  true,
		},
		{
			name:      "all odd numbers",
			slice:     []int{1, 3, 5},
			predicate: func(v int) bool { return v%2 != 0 },
			expected:  true,
		},
		{
			name:      "no matches",
			slice:     []int{1, 2, 3},
			predicate: func(v int) bool { return v > 10 },
			expected:  false,
		},
		{
			name:      "all match",
			slice:     []int{2, 4, 6},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  true,
		},
		{
			name:      "empty slice",
			slice:     []int{},
			predicate: func(v int) bool { return v > 0 },
			expected:  true,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(v int) bool { return v > 0 },
			expected:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.All(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("All() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestAll_String(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		predicate func(string) bool
		expected  bool
	}{
		{
			name:      "all strings starting with 'a'",
			slice:     []string{"apple", "apricot"},
			predicate: func(s string) bool { return strings.HasPrefix(s, "a") },
			expected:  true,
		},
		{
			name:      "all strings longer than 5",
			slice:     []string{"my cat", "my dog", "my elephant", "my bird"},
			predicate: func(s string) bool { return len(s) > 5 },
			expected:  true,
		},
		{
			name:      "no matches",
			slice:     []string{"a", "b", "c"},
			predicate: func(s string) bool { return len(s) > 5 },
			expected:  false,
		},
		{
			name:      "all match",
			slice:     []string{"a", "b", "c"},
			predicate: func(s string) bool { return len(s) == 1 },
			expected:  true,
		},
		{
			name:      "empty slice",
			slice:     []string{},
			predicate: func(s string) bool { return true },
			expected:  true,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(s string) bool { return true },
			expected:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.All(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("All() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestAll_Struct(t *testing.T) {
	type User struct {
		ID     int
		Name   string
		Active bool
	}
	tests := []struct {
		name      string
		slice     []User
		predicate func(User) bool
		expected  bool
	}{
		{
			name: "all active users",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", true},
				{3, "Charlie", true},
			},
			predicate: func(u User) bool { return u.Active },
			expected:  true,
		},
		{
			name: "all user with ID greater than 0",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", true},
				{3, "Charlie", true},
			},
			predicate: func(u User) bool { return u.ID > 0 },
			expected:  true,
		},
		{
			name: "no matches",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", false},
			},
			predicate: func(u User) bool { return !u.Active },
			expected:  false,
		},
		{
			name: "all match",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", true},
			},
			predicate: func(u User) bool { return u.Active },
			expected:  true,
		},
		{
			name:      "empty slice",
			slice:     []User{},
			predicate: func(u User) bool { return u.Active },
			expected:  true,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(u User) bool { return u.Active },
			expected:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.All(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("All() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestNone_Int(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		expected  bool
	}{
		{
			name:      "no even numbers",
			slice:     []int{1, 3, 5},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  true,
		},
		{
			name:      "no odd numbers",
			slice:     []int{2, 4, 6},
			predicate: func(v int) bool { return v%2 != 0 },
			expected:  true,
		},
		{
			name:      "no int greater than 10",
			slice:     []int{1, 2, 3},
			predicate: func(v int) bool { return v > 10 },
			expected:  true,
		},
		{
			name:      "all match",
			slice:     []int{2, 4, 6},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  false,
		},
		{
			name:      "empty slice",
			slice:     []int{},
			predicate: func(v int) bool { return v > 0 },
			expected:  true,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(v int) bool { return v > 0 },
			expected:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.None(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("None() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestCount_Int(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		expected  int
	}{
		{
			name:      "count even numbers",
			slice:     []int{1, 2, 3, 4, 5, 6},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  3,
		},
		{
			name:      "count odd numbers",
			slice:     []int{1, 2, 3, 4, 5, 6},
			predicate: func(v int) bool { return v%2 != 0 },
			expected:  3,
		},
		{
			name:      "no matches",
			slice:     []int{1, 2, 3},
			predicate: func(v int) bool { return v > 10 },
			expected:  0,
		},
		{
			name:      "all match",
			slice:     []int{2, 4, 6},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  3,
		},
		{
			name:      "empty slice",
			slice:     []int{},
			predicate: func(v int) bool { return v > 0 },
			expected:  0,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(v int) bool { return v > 0 },
			expected:  0,
		},
		{
			name:      "predicate always true",
			slice:     []int{1, 2, 3},
			predicate: func(v int) bool { return true },
			expected:  3,
		},
		{
			name:      "predicate always false",
			slice:     []int{1, 2, 3},
			predicate: func(v int) bool { return false },
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Count(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("Count() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestCount_String(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		predicate func(string) bool
		expected  int
	}{
		{
			name:      "count strings starting with 'a'",
			slice:     []string{"apple", "banana", "apricot", "avocado"},
			predicate: func(s string) bool { return strings.HasPrefix(s, "a") },
			expected:  3,
		},
		{
			name:      "count strings longer than 5",
			slice:     []string{"cat", "dog", "elephant", "bird"},
			predicate: func(s string) bool { return len(s) > 5 },
			expected:  1,
		},
		{
			name:      "count uppercase strings",
			slice:     []string{"HELLO", "world", "GO", "lang"},
			predicate: func(s string) bool { return s == strings.ToUpper(s) },
			expected:  2,
		},
		{
			name:      "no matches",
			slice:     []string{"a", "b", "c"},
			predicate: func(s string) bool { return len(s) > 5 },
			expected:  0,
		},
		{
			name:      "empty slice",
			slice:     []string{},
			predicate: func(s string) bool { return true },
			expected:  0,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(s string) bool { return true },
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Count(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("Count() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestCount_Struct(t *testing.T) {
	type User struct {
		ID     int
		Name   string
		Active bool
	}
	tests := []struct {
		name      string
		slice     []User
		predicate func(User) bool
		expected  int
	}{
		{
			name: "count active users",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", false},
				{3, "Charlie", true},
			},
			predicate: func(u User) bool { return u.Active },
			expected:  2,
		},
		{
			name: "count users with ID greater than 1",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", false},
				{3, "Charlie", true},
			},
			predicate: func(u User) bool { return u.ID > 1 },
			expected:  2,
		},
		{
			name: "count users with name starting with 'A'",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", false},
				{3, "Andrew", true},
			},
			predicate: func(u User) bool { return strings.HasPrefix(u.Name, "A") },
			expected:  2,
		},
		{
			name: "no matches",
			slice: []User{
				{1, "Alice", true},
				{2, "Bob", true},
			},
			predicate: func(u User) bool { return !u.Active },
			expected:  0,
		},
		{
			name:      "empty slice",
			slice:     []User{},
			predicate: func(u User) bool { return u.Active },
			expected:  0,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(u User) bool { return u.Active },
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Count(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("Count() = %v; want %v", result, tt.expected)
			}
		})
	}
}

// TestFilterNilInput verifies that Filter returns nil for nil input.
func TestFilterNilInput(t *testing.T) {
	result := lxslices.Filter[int](nil, func(v int) bool { return v > 0 })
	if result != nil {
		t.Errorf("Filter(nil, ...) = %v; want nil", result)
	}
}

// TestFilterNoMatchesReturnsEmptyNonNil verifies that a non-nil input with no matching
// elements returns a non-nil empty slice (not nil).
func TestFilterNoMatchesReturnsEmptyNonNil(t *testing.T) {
	result := lxslices.Filter([]int{1, 2, 3}, func(v int) bool { return v > 100 })
	if result == nil {
		t.Error("Filter(non-nil, no-match predicate) = nil; want non-nil empty slice")
	}
	if len(result) != 0 {
		t.Errorf("Filter(non-nil, no-match predicate) length = %d; want 0", len(result))
	}
}

// TestFilterMatches verifies that Filter correctly returns matching elements.
func TestFilterMatches(t *testing.T) {
	result := lxslices.Filter([]int{1, 2, 3, 4, 5}, func(v int) bool { return v%2 == 0 })
	expected := []int{2, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter() = %v; want %v", result, expected)
	}
}
