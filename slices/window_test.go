package lxslices_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/hgapdvn/lx/slices"
)

func TestWindow_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		size     int
		expected [][]int
	}{
		{
			name:     "sliding window of 2",
			slice:    []int{1, 2, 3, 4},
			size:     2,
			expected: [][]int{{1, 2}, {2, 3}, {3, 4}},
		},
		{
			name:     "sliding window of 3",
			slice:    []int{1, 2, 3, 4, 5},
			size:     3,
			expected: [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
		},
		{
			name:     "size equals 1",
			slice:    []int{1, 2, 3},
			size:     1,
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "size equals slice length",
			slice:    []int{1, 2, 3},
			size:     3,
			expected: [][]int{{1, 2, 3}},
		},
		{
			name:     "size larger than slice",
			slice:    []int{1, 2},
			size:     5,
			expected: [][]int{},
		},
		{
			name:     "single element slice",
			slice:    []int{42},
			size:     1,
			expected: [][]int{{42}},
		},
		{
			name:     "single element slice with larger size",
			slice:    []int{42},
			size:     2,
			expected: [][]int{},
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
			result, err := lxslices.Window(tt.slice, tt.size)
			if err != nil {
				t.Errorf("Window() unexpected error = %v", err)
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Window() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestWindow_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		size     int
		expected [][]string
	}{
		{
			name:     "sliding window of 2",
			slice:    []string{"a", "b", "c", "d"},
			size:     2,
			expected: [][]string{{"a", "b"}, {"b", "c"}, {"c", "d"}},
		},
		{
			name:     "sliding window of 3",
			slice:    []string{"a", "b", "c", "d", "e"},
			size:     3,
			expected: [][]string{{"a", "b", "c"}, {"b", "c", "d"}, {"c", "d", "e"}},
		},
		{
			name:     "size equals 1",
			slice:    []string{"a", "b", "c"},
			size:     1,
			expected: [][]string{{"a"}, {"b"}, {"c"}},
		},
		{
			name:     "size equals slice length",
			slice:    []string{"a", "b", "c"},
			size:     3,
			expected: [][]string{{"a", "b", "c"}},
		},
		{
			name:     "size larger than slice",
			slice:    []string{"a", "b"},
			size:     5,
			expected: [][]string{},
		},
		{
			name:     "single element slice",
			slice:    []string{"x"},
			size:     1,
			expected: [][]string{{"x"}},
		},
		{
			name:     "single element slice with larger size",
			slice:    []string{"x"},
			size:     2,
			expected: [][]string{},
		},
		{
			name:     "empty slice",
			slice:    []string{},
			size:     2,
			expected: [][]string{},
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
			result, err := lxslices.Window(tt.slice, tt.size)
			if err != nil {
				t.Errorf("Window() unexpected error = %v", err)
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Window() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestWindow_Struct(t *testing.T) {
	type Item struct{ ID int }
	tests := []struct {
		name     string
		slice    []Item
		size     int
		expected [][]Item
	}{
		{
			name:     "sliding window of 2",
			slice:    []Item{{1}, {2}, {3}, {4}},
			size:     2,
			expected: [][]Item{{{1}, {2}}, {{2}, {3}}, {{3}, {4}}},
		},
		{
			name:     "sliding window of 3",
			slice:    []Item{{1}, {2}, {3}, {4}, {5}},
			size:     3,
			expected: [][]Item{{{1}, {2}, {3}}, {{2}, {3}, {4}}, {{3}, {4}, {5}}},
		},
		{
			name:     "size equals 1",
			slice:    []Item{{1}, {2}, {3}},
			size:     1,
			expected: [][]Item{{{1}}, {{2}}, {{3}}},
		},
		{
			name:     "size equals slice length",
			slice:    []Item{{1}, {2}, {3}},
			size:     3,
			expected: [][]Item{{{1}, {2}, {3}}},
		},
		{
			name:     "size larger than slice",
			slice:    []Item{{1}, {2}},
			size:     5,
			expected: [][]Item{},
		},
		{
			name:     "single element slice",
			slice:    []Item{{42}},
			size:     1,
			expected: [][]Item{{{42}}},
		},
		{
			name:     "single element slice with larger size",
			slice:    []Item{{42}},
			size:     2,
			expected: [][]Item{},
		},
		{
			name:     "empty slice",
			slice:    []Item{},
			size:     2,
			expected: [][]Item{},
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
			result, err := lxslices.Window(tt.slice, tt.size)
			if err != nil {
				t.Errorf("Window() unexpected error = %v", err)
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Window() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestWindow_InvalidSize(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{name: "zero size", size: 0},
		{name: "negative size", size: -1},
		{name: "large negative size", size: -100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := lxslices.Window([]int{1, 2, 3}, tt.size)
			if !errors.Is(err, lxslices.ErrInvalidSize) {
				t.Errorf("Window() error = %v; want ErrInvalidSize", err)
			}
		})
	}
}

func TestWindowFunc_Int(t *testing.T) {
	sum := func(w []int) int {
		s := 0
		for _, v := range w {
			s += v
		}
		return s
	}

	tests := []struct {
		name     string
		slice    []int
		size     int
		fn       func([]int) int
		expected []int
	}{
		{
			name:     "sum of windows of 2",
			slice:    []int{1, 2, 3, 4},
			size:     2,
			fn:       sum,
			expected: []int{3, 5, 7},
		},
		{
			name:     "sum of windows of 3",
			slice:    []int{1, 2, 3, 4, 5},
			size:     3,
			fn:       sum,
			expected: []int{6, 9, 12},
		},
		{
			name:     "size equals 1",
			slice:    []int{10, 20, 30},
			size:     1,
			fn:       sum,
			expected: []int{10, 20, 30},
		},
		{
			name:     "size equals slice length",
			slice:    []int{1, 2, 3},
			size:     3,
			fn:       sum,
			expected: []int{6},
		},
		{
			name:     "size larger than slice",
			slice:    []int{1, 2},
			size:     5,
			fn:       sum,
			expected: []int{},
		},
		{
			name:     "empty slice",
			slice:    []int{},
			size:     2,
			fn:       sum,
			expected: []int{},
		},
		{
			name:     "nil slice",
			slice:    nil,
			size:     2,
			fn:       sum,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := lxslices.WindowFunc(tt.slice, tt.size, tt.fn)
			if err != nil {
				t.Errorf("WindowFunc() unexpected error = %v", err)
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("WindowFunc() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestWindowFunc_StringConcat(t *testing.T) {
	concat := func(w []string) string {
		result := ""
		for _, s := range w {
			result += s
		}
		return result
	}

	tests := []struct {
		name     string
		slice    []string
		size     int
		expected []string
	}{
		{
			name:     "concat windows of 2",
			slice:    []string{"a", "b", "c", "d"},
			size:     2,
			expected: []string{"ab", "bc", "cd"},
		},
		{
			name:     "nil slice",
			slice:    nil,
			size:     2,
			expected: nil,
		},
		{
			name:     "empty slice",
			slice:    []string{},
			size:     1,
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := lxslices.WindowFunc(tt.slice, tt.size, concat)
			if err != nil {
				t.Errorf("WindowFunc() unexpected error = %v", err)
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("WindowFunc() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestWindowFunc_InvalidSize(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{name: "zero size", size: 0},
		{name: "negative size", size: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := lxslices.WindowFunc([]int{1, 2, 3}, tt.size, func(w []int) int { return 0 })
			if !errors.Is(err, lxslices.ErrInvalidSize) {
				t.Errorf("WindowFunc() error = %v; want ErrInvalidSize", err)
			}
		})
	}
}
