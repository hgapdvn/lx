package lxmaps_test

import (
	"testing"

	lxmaps "github.com/hgapdvn/lx/lxmaps"
)

func TestContainsValueBy_StringInt(t *testing.T) {
	tests := []struct {
		name      string
		m         map[string]int
		predicate func(int) bool
		want      bool
	}{
		{
			name:      "nil map",
			m:         nil,
			predicate: func(v int) bool { return v > 1 },
			want:      false,
		},
		{
			name:      "empty map",
			m:         map[string]int{},
			predicate: func(v int) bool { return v > 1 },
			want:      false,
		},
		{
			name:      "single entry predicate matches",
			m:         map[string]int{"a": 2},
			predicate: func(v int) bool { return v > 1 },
			want:      true,
		},
		{
			name:      "single entry predicate does not match",
			m:         map[string]int{"a": 1},
			predicate: func(v int) bool { return v > 1 },
			want:      false,
		},
		{
			name:      "multiple entries predicate matches first",
			m:         map[string]int{"a": 5, "b": 2, "c": 1},
			predicate: func(v int) bool { return v > 3 },
			want:      true,
		},
		{
			name:      "multiple entries predicate matches middle",
			m:         map[string]int{"a": 1, "b": 5, "c": 2},
			predicate: func(v int) bool { return v > 3 },
			want:      true,
		},
		{
			name:      "multiple entries predicate matches last",
			m:         map[string]int{"a": 1, "b": 2, "c": 5},
			predicate: func(v int) bool { return v > 3 },
			want:      true,
		},
		{
			name:      "multiple entries predicate does not match any",
			m:         map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(v int) bool { return v > 10 },
			want:      false,
		},
		{
			name:      "predicate matches zero value",
			m:         map[string]int{"a": 0, "b": 1},
			predicate: func(v int) bool { return v == 0 },
			want:      true,
		},
		{
			name:      "predicate matches negative value",
			m:         map[string]int{"a": -1, "b": -2},
			predicate: func(v int) bool { return v < 0 },
			want:      true,
		},
		{
			name:      "predicate with multiple matches",
			m:         map[string]int{"a": 5, "b": 5, "c": 5},
			predicate: func(v int) bool { return v == 5 },
			want:      true,
		},
		{
			name:      "predicate with equality check",
			m:         map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(v int) bool { return v == 2 },
			want:      true,
		},
		{
			name:      "predicate with not found",
			m:         map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(v int) bool { return v == 4 },
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ContainsValueBy(tt.m, tt.predicate)
			if got != tt.want {
				t.Errorf("ContainsValueBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsValueBy_StringString(t *testing.T) {
	tests := []struct {
		name      string
		m         map[string]string
		predicate func(string) bool
		want      bool
	}{
		{
			name:      "nil map",
			m:         nil,
			predicate: func(v string) bool { return len(v) > 3 },
			want:      false,
		},
		{
			name:      "empty map",
			m:         map[string]string{},
			predicate: func(v string) bool { return len(v) > 3 },
			want:      false,
		},
		{
			name:      "single entry predicate matches",
			m:         map[string]string{"a": "hello"},
			predicate: func(v string) bool { return len(v) > 3 },
			want:      true,
		},
		{
			name:      "single entry predicate does not match",
			m:         map[string]string{"a": "hi"},
			predicate: func(v string) bool { return len(v) > 3 },
			want:      false,
		},
		{
			name:      "multiple entries predicate matches",
			m:         map[string]string{"a": "hello", "b": "world", "c": "go"},
			predicate: func(v string) bool { return len(v) > 3 },
			want:      true,
		},
		{
			name:      "multiple entries predicate does not match",
			m:         map[string]string{"a": "hi", "b": "go", "c": "ok"},
			predicate: func(v string) bool { return len(v) > 4 },
			want:      false,
		},
		{
			name:      "predicate checks string content",
			m:         map[string]string{"a": "apple", "b": "banana", "c": "cherry"},
			predicate: func(v string) bool { return v == "banana" },
			want:      true,
		},
		{
			name:      "predicate checks contains substring",
			m:         map[string]string{"a": "apple", "b": "apricot", "c": "banana"},
			predicate: func(v string) bool { return len(v) > 0 && v[0] == 'a' },
			want:      true,
		},
		{
			name:      "empty string value",
			m:         map[string]string{"a": "", "b": "hello"},
			predicate: func(v string) bool { return len(v) == 0 },
			want:      true,
		},
		{
			name:      "unicode string values",
			m:         map[string]string{"a": "こんにちは", "b": "世界"},
			predicate: func(v string) bool { return len(v) > 0 },
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ContainsValueBy(tt.m, tt.predicate)
			if got != tt.want {
				t.Errorf("ContainsValueBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsValueBy_IntBool(t *testing.T) {
	tests := []struct {
		name      string
		m         map[int]bool
		predicate func(bool) bool
		want      bool
	}{
		{
			name:      "nil map",
			m:         nil,
			predicate: func(v bool) bool { return v },
			want:      false,
		},
		{
			name:      "empty map",
			m:         map[int]bool{},
			predicate: func(v bool) bool { return v },
			want:      false,
		},
		{
			name:      "single entry true matches",
			m:         map[int]bool{1: true},
			predicate: func(v bool) bool { return v },
			want:      true,
		},
		{
			name:      "single entry false does not match",
			m:         map[int]bool{1: false},
			predicate: func(v bool) bool { return v },
			want:      false,
		},
		{
			name:      "multiple entries with true",
			m:         map[int]bool{1: false, 2: true, 3: false},
			predicate: func(v bool) bool { return v },
			want:      true,
		},
		{
			name:      "multiple entries all false",
			m:         map[int]bool{1: false, 2: false, 3: false},
			predicate: func(v bool) bool { return v },
			want:      false,
		},
		{
			name:      "predicate checks false",
			m:         map[int]bool{1: true, 2: false, 3: true},
			predicate: func(v bool) bool { return !v },
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ContainsValueBy(tt.m, tt.predicate)
			if got != tt.want {
				t.Errorf("ContainsValueBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsValueBy_StringStruct(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	tests := []struct {
		name      string
		m         map[string]User
		predicate func(User) bool
		want      bool
	}{
		{
			name:      "nil map",
			m:         nil,
			predicate: func(u User) bool { return u.Age > 18 },
			want:      false,
		},
		{
			name:      "empty map",
			m:         map[string]User{},
			predicate: func(u User) bool { return u.Age > 18 },
			want:      false,
		},
		{
			name: "single entry predicate matches",
			m: map[string]User{
				"alice": {Name: "Alice", Age: 25},
			},
			predicate: func(u User) bool { return u.Age > 18 },
			want:      true,
		},
		{
			name: "single entry predicate does not match",
			m: map[string]User{
				"bob": {Name: "Bob", Age: 16},
			},
			predicate: func(u User) bool { return u.Age > 18 },
			want:      false,
		},
		{
			name: "multiple entries predicate matches",
			m: map[string]User{
				"alice":   {Name: "Alice", Age: 25},
				"bob":     {Name: "Bob", Age: 16},
				"charlie": {Name: "Charlie", Age: 30},
			},
			predicate: func(u User) bool { return u.Age > 18 },
			want:      true,
		},
		{
			name: "multiple entries predicate does not match any",
			m: map[string]User{
				"alice": {Name: "Alice", Age: 16},
				"bob":   {Name: "Bob", Age: 15},
			},
			predicate: func(u User) bool { return u.Age > 18 },
			want:      false,
		},
		{
			name: "predicate matches by name",
			m: map[string]User{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
			},
			predicate: func(u User) bool { return u.Name == "Bob" },
			want:      true,
		},
		{
			name: "predicate matches by name length",
			m: map[string]User{
				"alice":   {Name: "Alice", Age: 25},
				"bob":     {Name: "Bob", Age: 30},
				"charlie": {Name: "Charlie", Age: 20},
			},
			predicate: func(u User) bool { return len(u.Name) > 5 },
			want:      true,
		},
		{
			name: "predicate with zero age",
			m: map[string]User{
				"alice": {Name: "Alice", Age: 0},
				"bob":   {Name: "Bob", Age: 25},
			},
			predicate: func(u User) bool { return u.Age == 0 },
			want:      true,
		},
		{
			name: "predicate with empty name",
			m: map[string]User{
				"alice": {Name: "", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
			},
			predicate: func(u User) bool { return u.Name == "" },
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ContainsValueBy(tt.m, tt.predicate)
			if got != tt.want {
				t.Errorf("ContainsValueBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
