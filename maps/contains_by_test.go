package lxmaps_test

import (
	"testing"

	lxmaps "github.com/hgapdvn/lx/maps"
)

func TestContainsBy_StringInt(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int
		fn   func(string, int) bool
		want bool
	}{
		{
			name: "empty map",
			m:    map[string]int{},
			fn:   func(k string, v int) bool { return k == "a" && v == 1 },
			want: false,
		},
		{
			name: "single entry with matching key and value",
			m:    map[string]int{"a": 1},
			fn:   func(k string, v int) bool { return k == "a" && v == 1 },
			want: true,
		},
		{
			name: "single entry with matching key but different value",
			m:    map[string]int{"a": 1},
			fn:   func(k string, v int) bool { return k == "a" && v == 2 },
			want: false,
		},
		{
			name: "single entry with different key but matching value",
			m:    map[string]int{"a": 1},
			fn:   func(k string, v int) bool { return k == "b" && v == 1 },
			want: false,
		},
		{
			name: "single entry with matching either key or value",
			m:    map[string]int{"a": 1},
			fn:   func(k string, v int) bool { return k == "a" || v == 1 },
			want: true,
		},
		{
			name: "multiple entries with matching key and value",
			m:    map[string]int{"a": 1, "b": 2, "c": 3},
			fn:   func(k string, v int) bool { return k == "a" && v == 1 },
			want: true,
		},
		{
			name: "multiple entries with matching key but different value",
			m:    map[string]int{"a": 1, "b": 2, "c": 3},
			fn:   func(k string, v int) bool { return k == "a" && v == 2 },
			want: false,
		},
		{
			name: "multiple entries with different key but matching value",
			m:    map[string]int{"a": 1, "b": 2, "c": 3},
			fn:   func(k string, v int) bool { return k == "d" && v == 3 },
			want: false,
		},
		{
			name: "multiple entries with matching either key or value",
			m:    map[string]int{"a": 1, "b": 2, "c": 3},
			fn:   func(k string, v int) bool { return k == "a" || v == 3 },
			want: true,
		},
		{
			name: "multiple entries with matching more than one key or value",
			m:    map[string]int{"a": 1, "b": 2, "c": 3},
			fn:   func(k string, v int) bool { return k == "a" || v == 2 || k == "c" },
			want: true,
		},
		{
			name: "multiple entries but no match",
			m:    map[string]int{"a": 1, "b": 2, "c": 3},
			fn:   func(k string, v int) bool { return k == "d" && v == 4 },
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ContainsBy(tt.m, tt.fn)
			if got != tt.want {
				t.Errorf("ContainsBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsBy_StringString(t *testing.T) {
	testes := []struct {
		name string
		m    map[string]string
		fn   func(string, string) bool
		want bool
	}{
		{
			name: "empty map",
			m:    map[string]string{},
			fn:   func(k string, v string) bool { return k == "a" && v == "b" },
			want: false,
		},
		{
			name: "single entry with matching key and value",
			m:    map[string]string{"a": "b"},
			fn:   func(k string, v string) bool { return k == "a" && v == "b" },
			want: true,
		},
		{
			name: "single entry with matching key but different value",
			m:    map[string]string{"a": "b"},
			fn:   func(k string, v string) bool { return k == "a" && v == "c" },
			want: false,
		},
		{
			name: "single entry with different key but matching value",
			m:    map[string]string{"a": "b"},
			fn:   func(k string, v string) bool { return k == "c" && v == "b" },
			want: false,
		},
		{
			name: "single entry with matching either key or value",
			m:    map[string]string{"a": "b"},
			fn:   func(k string, v string) bool { return k == "a" || v == "b" },
			want: true,
		},
		{
			name: "multiple entries with matching key and value",
			m:    map[string]string{"a": "b", "c": "d"},
			fn:   func(k string, v string) bool { return k == "a" && v == "b" },
			want: true,
		},
		{
			name: "multiple entries with matching key but different value",
			m:    map[string]string{"a": "b", "c": "d"},
			fn:   func(k string, v string) bool { return k == "a" && v == "c" },
			want: false,
		},
		{
			name: "multiple entries with different key but matching value",
			m:    map[string]string{"a": "b", "c": "d"},
			fn:   func(k string, v string) bool { return k == "c" && v == "b" },
			want: false,
		},
		{
			name: "multiple entries with matching either key or value",
			m:    map[string]string{"a": "b", "c": "d"},
			fn:   func(k string, v string) bool { return k == "a" || v == "b" },
			want: true,
		},
		{
			name: "multiple entries with matching more than one key or value",
			m:    map[string]string{"a": "b", "c": "d"},
			fn:   func(k string, v string) bool { return k == "a" || v == "c" },
			want: true,
		},
		{
			name: "multiple entries but no match",
			m:    map[string]string{"a": "b", "c": "d"},
			fn:   func(k string, v string) bool { return k == "e" && v == "f" },
			want: false,
		},
	}
	for _, tt := range testes {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ContainsBy(tt.m, tt.fn)
			if got != tt.want {
				t.Errorf("ContainsBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsBy_StringStruct(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	testes := []struct {
		name  string
		input map[string]User
		fn    func(string, User) bool
		want  bool
	}{
		{
			name:  "empty map",
			input: map[string]User{},
			want:  false,
		},
		{
			name: "single entry with matching key and value",
			input: map[string]User{
				"alice": {Name: "Alice", Age: 25},
			},
			fn:   func(k string, v User) bool { return k == "alice" && v.Name == "Alice" && v.Age == 25 },
			want: true,
		},
		{
			name: "single entry with matching key but different value",
			input: map[string]User{
				"alice": {Name: "Alice", Age: 25},
			},
			fn:   func(k string, v User) bool { return k == "alice" && v.Name == "Alice" && v.Age == 30 },
			want: false,
		},
		{
			name: "single entry with different key but matching value",
			input: map[string]User{
				"alice": {Name: "Alice", Age: 25},
			},
			fn:   func(k string, v User) bool { return k == "bob" && v.Name == "Alice" && v.Age == 25 },
			want: false,
		},
		{
			name: "single entry with matching either key or value",
			input: map[string]User{
				"alice": {Name: "Alice", Age: 25},
			},
			fn:   func(k string, v User) bool { return k == "alice" || v.Name == "Alice" },
			want: true,
		},
		{
			name: "multiple entries with matching key and value",
			input: map[string]User{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
				"kuro":  {Name: "Kuro", Age: 16},
			},
			fn:   func(k string, v User) bool { return k == "alice" && v.Name == "Alice" && v.Age == 25 },
			want: true,
		},
		{
			name: "multiple entries with matching key but different value",
			input: map[string]User{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
				"kuro":  {Name: "Kuro", Age: 16},
			},
			fn:   func(k string, v User) bool { return k == "alice" && v.Name == "Alice" && v.Age == 30 },
			want: false,
		},
		{
			name: "multiple entries with different key but matching value",
			input: map[string]User{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
				"kuro":  {Name: "Kuro", Age: 16},
			},
			fn:   func(k string, v User) bool { return k == "bob" && v.Name == "Alice" && v.Age == 30 },
			want: false,
		},
		{
			name: "multiple entries with matching either key or value",
			input: map[string]User{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
				"kuro":  {Name: "Kuro", Age: 16},
			},
			fn:   func(k string, v User) bool { return k == "alice" || v.Name == "Alice" },
			want: true,
		},
		{
			name: "multiple entries with matching more than one key or value",
			input: map[string]User{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
				"kuro":  {Name: "Kuro", Age: 16},
			},
			fn:   func(k string, v User) bool { return k == "alice" || v.Name == "Bob" },
			want: true,
		},
		{
			name: "multiple entries but no match",
			input: map[string]User{
				"alice": {Name: "Alice", Age: 25},
				"bob":   {Name: "Bob", Age: 30},
				"kuro":  {Name: "Kuro", Age: 16},
			},
			fn:   func(k string, v User) bool { return k == "charlie" && v.Name == "Charlie" },
			want: false,
		},
	}
	for _, tt := range testes {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.ContainsBy(tt.input, tt.fn)
			if got != tt.want {
				t.Errorf("ContainsBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
