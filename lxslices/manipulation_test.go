package lxslices_test

import (
	"reflect"
	"testing"

	"github.com/hgapdvn/lx/lxslices"
)

// Tests for Append
func TestAppend_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		elems    []int
		expected []int
	}{
		{name: "append to nil", slice: nil, elems: []int{1, 2}, expected: []int{1, 2}},
		{name: "append to non-empty", slice: []int{3}, elems: []int{4}, expected: []int{3, 4}},
		{name: "append none", slice: []int{5}, elems: []int{}, expected: []int{5}},
		{name: "append to empty non-nil", slice: []int{}, elems: []int{7}, expected: []int{7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.Append(tt.slice, tt.elems...)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Append(%v, %v) = %v; want %v", tt.slice, tt.elems, res, tt.expected)
			}
		})
	}
}

func TestAppend_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		elems    []string
		expected []string
	}{
		{name: "append to nil", slice: nil, elems: []string{"a"}, expected: []string{"a"}},
		{name: "append multiple", slice: []string{"b"}, elems: []string{"c", "d"}, expected: []string{"b", "c", "d"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.Append(tt.slice, tt.elems...)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Append(%v, %v) = %v; want %v", tt.slice, tt.elems, res, tt.expected)
			}
		})
	}
}

func TestAppend_Struct(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}
	tests := []struct {
		name     string
		slice    []User
		elem     User
		expected []User
	}{
		{name: "append single", slice: nil, elem: User{1, "A"}, expected: []User{{1, "A"}}},
		{name: "append to existing", slice: []User{{2, "B"}}, elem: User{3, "C"}, expected: []User{{2, "B"}, {3, "C"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.Append(tt.slice, tt.elem)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Append(%v, %v) = %v; want %v", tt.slice, tt.elem, res, tt.expected)
			}
		})
	}
}

func TestPrepend_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		elems    []int
		expected []int
	}{
		{name: "prepend to nil", slice: nil, elems: []int{1, 2}, expected: []int{1, 2}},
		{name: "prepend single", slice: []int{3}, elems: []int{0}, expected: []int{0, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.Prepend(tt.slice, tt.elems...)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Prepend(%v, %v) = %v; want %v", tt.slice, tt.elems, res, tt.expected)
			}
		})
	}
}

func TestPrepend_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		elems    []string
		expected []string
	}{
		{name: "prepend", slice: []string{"b"}, elems: []string{"a"}, expected: []string{"a", "b"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.Prepend(tt.slice, tt.elems...)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Prepend(%v, %v) = %v; want %v", tt.slice, tt.elems, res, tt.expected)
			}
		})
	}
}

func TestPrepend_Struct(t *testing.T) {
	type Item struct {
		K int
		V string
	}
	tests := []struct {
		name     string
		slice    []Item
		elem     Item
		expected []Item
	}{
		{name: "prepend struct", slice: []Item{{2, "b"}}, elem: Item{1, "a"}, expected: []Item{{1, "a"}, {2, "b"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.Prepend(tt.slice, tt.elem)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Prepend(%v, %v) = %v; want %v", tt.slice, tt.elem, res, tt.expected)
			}
		})
	}
}

func TestInsert_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		index    int
		elem     int
		expected []int
	}{
		{name: "insert middle", slice: []int{10, 30}, index: 1, elem: 20, expected: []int{10, 20, 30}},
		{name: "insert at 0", slice: []int{10}, index: 0, elem: 5, expected: []int{5, 10}},
		{name: "insert zero value at index zero", slice: []int{1, 2, 3}, index: 0, elem: 0, expected: []int{0, 1, 2, 3}},
		{name: "insert OOB append", slice: []int{1, 2}, index: 5, elem: 3, expected: []int{1, 2, 3}},
		{name: "insert negative treated as prepend", slice: []int{2, 3}, index: -5, elem: 1, expected: []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.Insert(tt.slice, tt.index, tt.elem)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Insert(%v,%d,%v) = %v; want %v", tt.slice, tt.index, tt.elem, res, tt.expected)
			}
		})
	}
}

func TestInsert_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		index    int
		elem     string
		expected []string
	}{
		{name: "insert", slice: []string{"a", "c"}, index: 1, elem: "b", expected: []string{"a", "b", "c"}},
		{name: "insert at beginning with negative index", slice: []string{"x"}, index: -1, elem: "y", expected: []string{"y", "x"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.Insert(tt.slice, tt.index, tt.elem)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Insert(%v,%d,%v) = %v; want %v", tt.slice, tt.index, tt.elem, res, tt.expected)
			}
		})
	}
}

func TestInsert_Struct(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}
	tests := []struct {
		name     string
		slice    []User
		index    int
		elem     User
		expected []User
	}{
		{name: "insert struct", slice: []User{{1, "A"}, {3, "C"}}, index: 1, elem: User{2, "B"}, expected: []User{{1, "A"}, {2, "B"}, {3, "C"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.Insert(tt.slice, tt.index, tt.elem)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Insert(%v,%d,%v) = %v; want %v", tt.slice, tt.index, tt.elem, res, tt.expected)
			}
		})
	}
}

// Tests for Remove (first occurrence)
func TestRemove_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		value    int
		expected []int
	}{
		{name: "remove present", slice: []int{1, 2, 3, 2}, value: 2, expected: []int{1, 3, 2}},
		{name: "remove absent", slice: []int{1, 2}, value: 9, expected: []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.Remove(tt.slice, tt.value)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Remove(%v,%v) = %v; want %v", tt.slice, tt.value, res, tt.expected)
			}
		})
	}
}

func TestRemove_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		value    string
		expected []string
	}{
		{name: "remove", slice: []string{"x", "y", "x"}, value: "x", expected: []string{"y", "x"}},
		{name: "remove absent", slice: []string{"a", "b"}, value: "z", expected: []string{"a", "b"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.Remove(tt.slice, tt.value)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Remove(%v,%v) = %v; want %v", tt.slice, tt.value, res, tt.expected)
			}
		})
	}
}

func TestRemove_Struct(t *testing.T) {
	type Item struct {
		ID    int
		Label string
	}
	tests := []struct {
		name     string
		slice    []Item
		value    Item
		expected []Item
	}{
		{name: "remove struct", slice: []Item{{1, "A"}, {2, "B"}, {1, "A"}}, value: Item{1, "A"}, expected: []Item{{2, "B"}, {1, "A"}}},
		{name: "remove absent", slice: []Item{{1, "A"}}, value: Item{9, "Z"}, expected: []Item{{1, "A"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.Remove(tt.slice, tt.value)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Remove(%v,%v) = %v; want %v", tt.slice, tt.value, res, tt.expected)
			}
		})
	}
}

func TestRemoveAt_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		index    int
		expected []int
	}{
		{name: "remove middle", slice: []int{1, 2, 3}, index: 1, expected: []int{1, 3}},
		{name: "remove OOB", slice: []int{1, 2}, index: 5, expected: []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.RemoveAt(tt.slice, tt.index)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("RemoveAt(%v,%d) = %v; want %v", tt.slice, tt.index, res, tt.expected)
			}
		})
	}
}

func TestRemoveAt_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		index    int
		expected []string
	}{
		{name: "remove first", slice: []string{"a", "b"}, index: 0, expected: []string{"b"}},
		{name: "remove OOB returns original", slice: []string{"a"}, index: 5, expected: []string{"a"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.RemoveAt(tt.slice, tt.index)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("RemoveAt(%v,%d) = %v; want %v", tt.slice, tt.index, res, tt.expected)
			}
		})
	}
}

func TestRemoveAt_Struct(t *testing.T) {
	type Pair struct {
		A int
		B string
	}
	tests := []struct {
		name     string
		slice    []Pair
		index    int
		expected []Pair
	}{
		{name: "remove struct", slice: []Pair{{1, "x"}, {2, "y"}}, index: 0, expected: []Pair{{2, "y"}}},
		{name: "remove OOB returns original", slice: []Pair{{1, "x"}}, index: 3, expected: []Pair{{1, "x"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.RemoveAt(tt.slice, tt.index)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("RemoveAt(%v,%d) = %v; want %v", tt.slice, tt.index, res, tt.expected)
			}
		})
	}
}

func TestRemoveAll_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		value    int
		expected []int
	}{
		{name: "remove all", slice: []int{1, 2, 1, 3}, value: 1, expected: []int{2, 3}},
		{name: "none removed", slice: []int{4}, value: 9, expected: []int{4}},
		{name: "empty slice", slice: []int{}, value: 1, expected: []int{}},
		{name: "nil slice", slice: nil, value: 1, expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.RemoveAll(tt.slice, tt.value)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("RemoveAll(%v,%v) = %v; want %v", tt.slice, tt.value, res, tt.expected)
			}
		})
	}
}

func TestRemoveAll_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		value    string
		expected []string
	}{
		{name: "remove all", slice: []string{"a", "b", "a"}, value: "a", expected: []string{"b"}},
		{name: "empty slice", slice: []string{}, value: "a", expected: []string{}},
		{name: "nil slice", slice: nil, value: "a", expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.RemoveAll(tt.slice, tt.value)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("RemoveAll(%v,%v) = %v; want %v", tt.slice, tt.value, res, tt.expected)
			}
		})
	}
}

func TestRemoveAll_Struct(t *testing.T) {
	type Node struct {
		ID    int
		Label string
	}
	tests := []struct {
		name     string
		slice    []Node
		value    Node
		expected []Node
	}{
		{name: "remove nodes", slice: []Node{{1, "A"}, {1, "A"}, {2, "B"}}, value: Node{1, "A"}, expected: []Node{{2, "B"}}},
		{name: "empty slice", slice: []Node{}, value: Node{1, "A"}, expected: []Node{}},
		{name: "nil slice", slice: nil, value: Node{1, "A"}, expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.RemoveAll(tt.slice, tt.value)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("RemoveAll(%v,%v) = %v; want %v", tt.slice, tt.value, res, tt.expected)
			}
		})
	}
}

func TestRemoveFunc_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		pred     func(int) bool
		expected []int
	}{
		{name: "remove evens", slice: []int{1, 2, 3, 4}, pred: func(v int) bool { return v%2 == 0 }, expected: []int{1, 3}},
		{name: "nil slice", slice: nil, pred: func(v int) bool { return v%2 == 0 }, expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.RemoveFunc(tt.slice, tt.pred)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("RemoveFunc(%v) = %v; want %v", tt.slice, res, tt.expected)
			}
		})
	}
}

func TestRemoveFunc_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		pred     func(string) bool
		expected []string
	}{
		{name: "remove long", slice: []string{"apple", "pear"}, pred: func(s string) bool { return len(s) > 4 }, expected: []string{"pear"}},
		{name: "empty slice", slice: []string{}, pred: func(s string) bool { return true }, expected: []string{}},
		{name: "nil slice", slice: nil, pred: func(s string) bool { return true }, expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.RemoveFunc(tt.slice, tt.pred)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("RemoveFunc(%v) = %v; want %v", tt.slice, res, tt.expected)
			}
		})
	}
}

func TestRemoveFunc_Struct(t *testing.T) {
	type Rec struct {
		ID  int
		Tag string
	}
	tests := []struct {
		name     string
		slice    []Rec
		pred     func(Rec) bool
		expected []Rec
	}{
		{name: "remove tag A", slice: []Rec{{1, "A"}, {2, "B"}, {3, "A"}}, pred: func(r Rec) bool { return r.Tag == "A" }, expected: []Rec{{2, "B"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.RemoveFunc(tt.slice, tt.pred)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("RemoveFunc(%v) = %v; want %v", tt.slice, res, tt.expected)
			}
		})
	}
}

// Tests for RemoveDuplicates (Unique)
func TestRemoveDuplicates_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected []int
	}{
		{name: "unique ints", slice: []int{2, 2, 1, 3, 1}, expected: []int{2, 1, 3}},
		{name: "empty", slice: []int{}, expected: []int{}},
		{name: "nil slice", slice: nil, expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.RemoveDuplicates(tt.slice)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("RemoveDuplicates(%v) = %v; want %v", tt.slice, res, tt.expected)
			}
		})
	}
}

func TestRemoveDuplicates_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		expected []string
	}{
		{name: "unique strings", slice: []string{"a", "a", "b"}, expected: []string{"a", "b"}},
		{name: "empty", slice: []string{}, expected: []string{}},
		{name: "nil slice", slice: nil, expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.RemoveDuplicates(tt.slice)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("RemoveDuplicates(%v) = %v; want %v", tt.slice, res, tt.expected)
			}
		})
	}
}

func TestRemoveDuplicates_Struct(t *testing.T) {
	type Node struct {
		ID   int
		Name string
	}
	tests := []struct {
		name     string
		slice    []Node
		expected []Node
	}{
		{name: "unique structs", slice: []Node{{1, "A"}, {1, "A"}, {2, "B"}}, expected: []Node{{1, "A"}, {2, "B"}}},
		{name: "empty", slice: []Node{}, expected: []Node{}},
		{name: "nil slice", slice: nil, expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.RemoveDuplicates(tt.slice)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("RemoveDuplicates(%v) = %v; want %v", tt.slice, res, tt.expected)
			}
		})
	}
}

// Tests for Replace
func TestReplace_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		old      int
		new      int
		expected []int
	}{
		{name: "replace ints", slice: []int{1, 2, 1}, old: 1, new: 9, expected: []int{9, 2, 9}},
		{name: "replace none", slice: []int{5}, old: 0, new: 1, expected: []int{5}},
		{name: "empty", slice: []int{}, old: 1, new: 2, expected: []int{}},
		{name: "nil slice", slice: nil, old: 1, new: 2, expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.Replace(tt.slice, tt.old, tt.new)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Replace(%v,%v,%v) = %v; want %v", tt.slice, tt.old, tt.new, res, tt.expected)
			}
		})
	}
}

func TestReplace_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		old      string
		new      string
		expected []string
	}{
		{name: "replace strings", slice: []string{"go", "py", "go"}, old: "go", new: "golang", expected: []string{"golang", "py", "golang"}},
		{name: "empty slice", slice: []string{}, old: "x", new: "y", expected: []string{}},
		{name: "nil slice", slice: nil, old: "x", new: "y", expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.Replace(tt.slice, tt.old, tt.new)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Replace(%v,%v,%v) = %v; want %v", tt.slice, tt.old, tt.new, res, tt.expected)
			}
		})
	}
}

func TestReplace_Struct(t *testing.T) {
	type Pair struct {
		ID int
		S  string
	}
	tests := []struct {
		name     string
		slice    []Pair
		old      Pair
		new      Pair
		expected []Pair
	}{
		{name: "replace struct", slice: []Pair{{1, "A"}, {2, "B"}, {1, "A"}}, old: Pair{1, "A"}, new: Pair{9, "Z"}, expected: []Pair{{9, "Z"}, {2, "B"}, {9, "Z"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.Replace(tt.slice, tt.old, tt.new)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Replace(%v,%v,%v) = %v; want %v", tt.slice, tt.old, tt.new, res, tt.expected)
			}
		})
	}
}

// Tests for ReplaceAt
func TestReplaceAt_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		index    int
		new      int
		expected []int
	}{
		{name: "replace at", slice: []int{1, 2, 3}, index: 1, new: 8, expected: []int{1, 8, 3}},
		{name: "OOB returns original", slice: []int{1}, index: 5, new: 9, expected: []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.ReplaceAt(tt.slice, tt.index, tt.new)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("ReplaceAt(%v,%d,%v) = %v; want %v", tt.slice, tt.index, tt.new, res, tt.expected)
			}
		})
	}
}

func TestReplaceAt_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		index    int
		new      string
		expected []string
	}{
		{name: "replace at", slice: []string{"a", "b"}, index: 0, new: "z", expected: []string{"z", "b"}},
		{name: "OOB negative index", slice: []string{"a"}, index: -1, new: "x", expected: []string{"a"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.ReplaceAt(tt.slice, tt.index, tt.new)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("ReplaceAt(%v,%d,%v) = %v; want %v", tt.slice, tt.index, tt.new, res, tt.expected)
			}
		})
	}
}

func TestReplaceAt_Struct(t *testing.T) {
	type Node struct {
		ID   int
		Name string
	}
	tests := []struct {
		name     string
		slice    []Node
		index    int
		new      Node
		expected []Node
	}{
		{name: "replace at", slice: []Node{{1, "A"}, {2, "B"}}, index: 0, new: Node{9, "Z"}, expected: []Node{{9, "Z"}, {2, "B"}}},
		{name: "OOB returns original", slice: []Node{{1, "A"}}, index: 5, new: Node{8, "X"}, expected: []Node{{1, "A"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lxslices.ReplaceAt(tt.slice, tt.index, tt.new)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("ReplaceAt(%v,%d,%v) = %v; want %v", tt.slice, tt.index, tt.new, res, tt.expected)
			}
		})
	}
}

func TestRotateLeft_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		k        int
		expected []int
	}{
		{name: "rotate 2", slice: []int{1, 2, 3, 4, 5}, k: 2, expected: []int{3, 4, 5, 1, 2}},
		{name: "k > len", slice: []int{1, 2, 3, 4, 5}, k: 7, expected: []int{3, 4, 5, 1, 2}},
		{name: "k == 0 returns same", slice: []int{1, 2, 3}, k: 0, expected: []int{1, 2, 3}},
		{name: "k == len returns same", slice: []int{1, 2, 3}, k: 3, expected: []int{1, 2, 3}},
		{name: "negative k rotates right", slice: []int{1, 2, 3, 4}, k: -1, expected: []int{4, 1, 2, 3}},
		{name: "empty slice", slice: []int{}, k: 3, expected: []int{}},
		{name: "nil slice", slice: nil, k: 2, expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lxslices.RotateLeft(tt.slice, tt.k)
			if !reflect.DeepEqual(tt.slice, tt.expected) {
				t.Errorf("RotateLeft resulted in %v; want %v", tt.slice, tt.expected)
			}
		})
	}
}

func TestRotateLeft_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		k        int
		expected []string
	}{
		{name: "rotate 1", slice: []string{"a", "b", "c", "d"}, k: 1, expected: []string{"b", "c", "d", "a"}},
		{name: "k > len", slice: []string{"a", "b", "c", "d"}, k: 5, expected: []string{"b", "c", "d", "a"}},
		{name: "negative k rotates right", slice: []string{"a", "b", "c", "d"}, k: -1, expected: []string{"d", "a", "b", "c"}},
		{name: "empty", slice: []string{}, k: 2, expected: []string{}},
		{name: "nil slice", slice: nil, k: 1, expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lxslices.RotateLeft(tt.slice, tt.k)
			if !reflect.DeepEqual(tt.slice, tt.expected) {
				t.Errorf("RotateLeft resulted in %v; want %v", tt.slice, tt.expected)
			}
		})
	}
}

func TestRotateLeft_Struct(t *testing.T) {
	type Node struct{ ID int }
	tests := []struct {
		name     string
		slice    []Node
		k        int
		expected []Node
	}{
		{name: "rotate 1", slice: []Node{{1}, {2}, {3}}, k: 1, expected: []Node{{2}, {3}, {1}}},
		{name: "k > len", slice: []Node{{1}, {2}, {3}}, k: 4, expected: []Node{{2}, {3}, {1}}},
		{name: "negative k rotates right", slice: []Node{{1}, {2}, {3}}, k: -1, expected: []Node{{3}, {1}, {2}}},
		{name: "empty", slice: []Node{}, k: 2, expected: []Node{}},
		{name: "nil slice", slice: nil, k: 1, expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lxslices.RotateLeft(tt.slice, tt.k)
			if !reflect.DeepEqual(tt.slice, tt.expected) {
				t.Errorf("RotateLeft resulted in %v; want %v", tt.slice, tt.expected)
			}
		})
	}
}

func TestRotateRight_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		k        int
		expected []int
	}{
		{name: "rotate right 2", slice: []int{1, 2, 3, 4, 5}, k: 2, expected: []int{4, 5, 1, 2, 3}},
		{name: "k == 0 returns same", slice: []int{7, 8}, k: 0, expected: []int{7, 8}},
		{name: "k == len returns same", slice: []int{7, 8}, k: 2, expected: []int{7, 8}},
		{name: "empty", slice: []int{}, k: 1, expected: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lxslices.RotateRight(tt.slice, tt.k)
			if !reflect.DeepEqual(tt.slice, tt.expected) {
				t.Errorf("RotateRight resulted in %v; want %v", tt.slice, tt.expected)
			}
		})
	}
}

func TestRotateRight_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		k        int
		expected []string
	}{
		{name: "rotate right 1", slice: []string{"a", "b", "c", "d"}, k: 1, expected: []string{"d", "a", "b", "c"}},
		{name: "k > len", slice: []string{"a", "b", "c", "d"}, k: 5, expected: []string{"d", "a", "b", "c"}},
		{name: "negative k rotates left", slice: []string{"a", "b", "c", "d"}, k: -1, expected: []string{"b", "c", "d", "a"}},
		{name: "k == 0 returns same", slice: []string{"x", "y"}, k: 0, expected: []string{"x", "y"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lxslices.RotateRight(tt.slice, tt.k)
			if !reflect.DeepEqual(tt.slice, tt.expected) {
				t.Errorf("RotateRight resulted in %v; want %v", tt.slice, tt.expected)
			}
		})
	}
}

func TestRotateRight_Struct(t *testing.T) {
	type Node struct{ ID int }
	tests := []struct {
		name     string
		slice    []Node
		k        int
		expected []Node
	}{
		{name: "rotate right 1", slice: []Node{{1}, {2}, {3}}, k: 1, expected: []Node{{3}, {1}, {2}}},
		{name: "k > len", slice: []Node{{1}, {2}, {3}}, k: 4, expected: []Node{{3}, {1}, {2}}},
		{name: "negative k rotates left", slice: []Node{{1}, {2}, {3}}, k: -1, expected: []Node{{2}, {3}, {1}}},
		{name: "single element", slice: []Node{{9}}, k: 10, expected: []Node{{9}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lxslices.RotateRight(tt.slice, tt.k)
			if !reflect.DeepEqual(tt.slice, tt.expected) {
				t.Errorf("RotateRight resulted in %v; want %v", tt.slice, tt.expected)
			}
		})
	}
}
