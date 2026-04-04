package lxtime_test

import (
	"fmt"
	"testing"
	"time"

	lxtime "github.com/hgapdvn/lx/time"
)

func TestIsSameMonth_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected bool
	}{
		{
			name:     "same month same day same time",
			t1:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "same month different days",
			t1:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2026, 4, 30, 23, 59, 59, 0, time.UTC),
			expected: true,
		},
		{
			name:     "same month first and last day",
			t1:       time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
			t2:       time.Date(2026, 4, 30, 23, 59, 59, 0, time.UTC),
			expected: true,
		},
		{
			name:     "different months same year",
			t1:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2026, 5, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "adjacent months",
			t1:       time.Date(2026, 3, 31, 23, 59, 59, 0, time.UTC),
			t2:       time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "different years same month",
			t1:       time.Date(2025, 4, 4, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "january both years",
			t1:       time.Date(2025, 1, 15, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsSameMonth(tt.t1, tt.t2)
			if result != tt.expected {
				t.Errorf("IsSameMonth(%v, %v) = %v, want %v", tt.t1, tt.t2, result, tt.expected)
			}
		})
	}
}

func TestIsSameMonth_AllMonths(t *testing.T) {
	// Test each month of the year
	for month := 1; month <= 12; month++ {
		t.Run("month_"+time.Month(month).String(), func(t *testing.T) {
			t1 := time.Date(2026, time.Month(month), 1, 10, 30, 0, 0, time.UTC)
			t2 := time.Date(2026, time.Month(month), 15, 20, 45, 30, 0, time.UTC)

			result := lxtime.IsSameMonth(t1, t2)
			if !result {
				t.Errorf("IsSameMonth(%v, %v) = %v, want true", t1, t2, result)
			}
		})
	}
}

func TestIsSameMonth_Symmetry(t *testing.T) {
	// Test that IsSameMonth is symmetric
	t1 := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
	t2 := time.Date(2026, 4, 30, 23, 59, 59, 0, time.UTC)

	result1 := lxtime.IsSameMonth(t1, t2)
	result2 := lxtime.IsSameMonth(t2, t1)

	if result1 != result2 {
		t.Errorf("IsSameMonth(t1, t2) = %v but IsSameMonth(t2, t1) = %v, expected symmetry", result1, result2)
	}
}

func TestIsSameMonth_Transitivity(t *testing.T) {
	// Test transitivity: if IsSameMonth(a, b) and IsSameMonth(b, c) then IsSameMonth(a, c)
	t1 := time.Date(2026, 4, 1, 10, 0, 0, 0, time.UTC)
	t2 := time.Date(2026, 4, 15, 15, 0, 0, 0, time.UTC)
	t3 := time.Date(2026, 4, 30, 20, 0, 0, 0, time.UTC)

	if !lxtime.IsSameMonth(t1, t2) || !lxtime.IsSameMonth(t2, t3) {
		t.Fatal("Test setup failed")
	}

	if !lxtime.IsSameMonth(t1, t3) {
		t.Error("IsSameMonth is not transitive")
	}
}

func ExampleIsSameMonth() {
	t1 := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
	t2 := time.Date(2026, 4, 30, 23, 59, 59, 0, time.UTC)
	result := lxtime.IsSameMonth(t1, t2)
	fmt.Println(result)

	t3 := time.Date(2026, 5, 1, 0, 0, 0, 0, time.UTC)
	result = lxtime.IsSameMonth(t1, t3)
	fmt.Println(result)
	// Output:
	// true
	// false
}
