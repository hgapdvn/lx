package lxslices

import (
	"reflect"
	"testing"
)

func TestRepeat(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		n        int
		expected []int
	}{
		{
			name:     "repeat positive n",
			value:    5,
			n:        3,
			expected: []int{5, 5, 5},
		},
		{
			name:     "repeat n equals 1",
			value:    42,
			n:        1,
			expected: []int{42},
		},
		{
			name:     "repeat n equals 0",
			value:    10,
			n:        0,
			expected: nil,
		},
		{
			name:     "repeat negative n",
			value:    7,
			n:        -5,
			expected: nil,
		},
		{
			name:     "repeat large n",
			value:    1,
			n:        1000,
			expected: make([]int, 1000),
		},
	}

	// Fill the large n test expected result
	for i := range tests[4].expected {
		tests[4].expected[i] = 1
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Repeat(tt.value, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Repeat(%v, %d) = %v, want %v", tt.value, tt.n, result, tt.expected)
			}
		})
	}
}

func TestRepeat_DifferentTypes(t *testing.T) {
	t.Run("string type", func(t *testing.T) {
		result := Repeat("hello", 3)
		expected := []string{"hello", "hello", "hello"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Repeat() = %v, want %v", result, expected)
		}
	})

	t.Run("float64 type", func(t *testing.T) {
		result := Repeat(3.14, 2)
		expected := []float64{3.14, 3.14}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Repeat() = %v, want %v", result, expected)
		}
	})

	t.Run("struct type", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		p := Person{Name: "John", Age: 30}
		result := Repeat(p, 2)
		expected := []Person{p, p}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Repeat() = %v, want %v", result, expected)
		}
	})

	t.Run("pointer type", func(t *testing.T) {
		val := 42
		ptr := &val
		result := Repeat(ptr, 3)
		if len(result) != 3 {
			t.Errorf("Repeat() length = %d, want 3", len(result))
		}
		for i, p := range result {
			if p != ptr {
				t.Errorf("Repeat()[%d] = %v, want %v", i, p, ptr)
			}
		}
	})
}

func TestRepeatSlice(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		n        int
		expected []int
	}{
		{
			name:     "repeat slice positive n",
			slice:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3, 1, 2, 3, 1, 2, 3},
		},
		{
			name:     "repeat slice n equals 1",
			slice:    []int{5, 6},
			n:        1,
			expected: []int{5, 6},
		},
		{
			name:     "repeat slice n equals 0",
			slice:    []int{1, 2},
			n:        0,
			expected: nil,
		},
		{
			name:     "repeat slice negative n",
			slice:    []int{1, 2},
			n:        -3,
			expected: nil,
		},
		{
			name:     "repeat empty slice",
			slice:    []int{},
			n:        3,
			expected: []int{},
		},
		{
			name:     "repeat nil slice",
			slice:    nil,
			n:        3,
			expected: nil,
		},
		{
			name:     "repeat single element slice",
			slice:    []int{7},
			n:        4,
			expected: []int{7, 7, 7, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RepeatSlice(tt.slice, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("RepeatSlice(%v, %d) = %v, want %v", tt.slice, tt.n, result, tt.expected)
			}
		})
	}
}

func TestRepeatSlice_DifferentTypes(t *testing.T) {
	t.Run("string slice", func(t *testing.T) {
		result := RepeatSlice([]string{"a", "b"}, 2)
		expected := []string{"a", "b", "a", "b"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("RepeatSlice() = %v, want %v", result, expected)
		}
	})

	t.Run("bool slice", func(t *testing.T) {
		result := RepeatSlice([]bool{true, false}, 3)
		expected := []bool{true, false, true, false, true, false}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("RepeatSlice() = %v, want %v", result, expected)
		}
	})
}

func TestRange(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		end      int
		expected []int
	}{
		{
			name:     "range positive numbers",
			start:    1,
			end:      5,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "range from zero",
			start:    0,
			end:      5,
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			name:     "range negative to positive",
			start:    -3,
			end:      3,
			expected: []int{-3, -2, -1, 0, 1, 2},
		},
		{
			name:     "range single element",
			start:    5,
			end:      6,
			expected: []int{5},
		},
		{
			name:     "range start equals end",
			start:    5,
			end:      5,
			expected: nil,
		},
		{
			name:     "range start greater than end",
			start:    10,
			end:      5,
			expected: nil,
		},
		{
			name:     "range negative numbers",
			start:    -5,
			end:      -1,
			expected: []int{-5, -4, -3, -2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Range(tt.start, tt.end)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Range(%d, %d) = %v, want %v", tt.start, tt.end, result, tt.expected)
			}
		})
	}
}

func TestRange_DifferentIntegerTypes(t *testing.T) {
	t.Run("int8 type", func(t *testing.T) {
		result := Range(int8(1), int8(5))
		expected := []int8{1, 2, 3, 4}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Range() = %v, want %v", result, expected)
		}
	})

	t.Run("int64 type", func(t *testing.T) {
		result := Range(int64(10), int64(15))
		expected := []int64{10, 11, 12, 13, 14}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Range() = %v, want %v", result, expected)
		}
	})

	t.Run("uint type", func(t *testing.T) {
		result := Range(uint(0), uint(4))
		expected := []uint{0, 1, 2, 3}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Range() = %v, want %v", result, expected)
		}
	})
}

func TestRangeStep(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		end      int
		step     int
		expected []int
	}{
		{
			name:     "positive step",
			start:    0,
			end:      10,
			step:     2,
			expected: []int{0, 2, 4, 6, 8},
		},
		{
			name:     "positive step not aligned",
			start:    1,
			end:      10,
			step:     3,
			expected: []int{1, 4, 7},
		},
		{
			name:     "negative step",
			start:    10,
			end:      0,
			step:     -2,
			expected: []int{10, 8, 6, 4, 2},
		},
		{
			name:     "negative step not aligned",
			start:    10,
			end:      1,
			step:     -3,
			expected: []int{10, 7, 4},
		},
		{
			name:     "step of 1",
			start:    1,
			end:      5,
			step:     1,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "step of -1",
			start:    5,
			end:      1,
			step:     -1,
			expected: []int{5, 4, 3, 2},
		},
		{
			name:     "step equals 0",
			start:    1,
			end:      5,
			step:     0,
			expected: nil,
		},
		{
			name:     "positive step with start >= end",
			start:    10,
			end:      5,
			step:     2,
			expected: nil,
		},
		{
			name:     "negative step with start <= end",
			start:    5,
			end:      10,
			step:     -2,
			expected: nil,
		},
		{
			name:     "start equals end with positive step",
			start:    5,
			end:      5,
			step:     1,
			expected: nil,
		},
		{
			name:     "start equals end with negative step",
			start:    5,
			end:      5,
			step:     -1,
			expected: nil,
		},
		{
			name:     "large step",
			start:    0,
			end:      100,
			step:     25,
			expected: []int{0, 25, 50, 75},
		},
		{
			name:     "negative range with negative step",
			start:    -1,
			end:      -10,
			step:     -2,
			expected: []int{-1, -3, -5, -7, -9},
		},
		{
			name:     "step larger than range",
			start:    0,
			end:      3,
			step:     10,
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RangeStep(tt.start, tt.end, tt.step)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("RangeStep(%d, %d, %d) = %v, want %v", tt.start, tt.end, tt.step, result, tt.expected)
			}
		})
	}
}

func TestRangeStep_DifferentIntegerTypes(t *testing.T) {
	t.Run("int8 type", func(t *testing.T) {
		result := RangeStep(int8(0), int8(10), int8(3))
		expected := []int8{0, 3, 6, 9}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("RangeStep() = %v, want %v", result, expected)
		}
	})

	t.Run("uint type", func(t *testing.T) {
		result := RangeStep(uint(0), uint(10), uint(2))
		expected := []uint{0, 2, 4, 6, 8}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("RangeStep() = %v, want %v", result, expected)
		}
	})

	t.Run("int64 type", func(t *testing.T) {
		result := RangeStep(int64(5), int64(0), int64(-1))
		expected := []int64{5, 4, 3, 2, 1}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("RangeStep() = %v, want %v", result, expected)
		}
	})
}

// Benchmarks
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(42, 100)
	}
}

func BenchmarkRepeatSlice(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		RepeatSlice(slice, 20)
	}
}

func BenchmarkRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Range(0, 1000)
	}
}

func BenchmarkRangeStep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeStep(0, 1000, 2)
	}
}
