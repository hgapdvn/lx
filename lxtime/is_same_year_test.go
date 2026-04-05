package lxtime_test

import (
	"fmt"
	"testing"
	"time"

	lxtime "github.com/hgapdvn/lx/lxtime"
)

func TestIsSameYear_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected bool
	}{
		{
			name:     "same year same month same day",
			t1:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "same year january to december",
			t1:       time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			t2:       time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC),
			expected: true,
		},
		{
			name:     "same year different months",
			t1:       time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2026, 6, 15, 10, 30, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "different years",
			t1:       time.Date(2025, 4, 4, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "year boundary",
			t1:       time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC),
			t2:       time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "one year apart",
			t1:       time.Date(2024, 6, 15, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2025, 6, 15, 10, 30, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "many years apart same dates",
			t1:       time.Date(2000, 4, 4, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsSameYear(tt.t1, tt.t2)
			if result != tt.expected {
				t.Errorf("IsSameYear(%v, %v) = %v, want %v", tt.t1, tt.t2, result, tt.expected)
			}
		})
	}
}

func TestIsSameYear_Symmetry(t *testing.T) {
	// Test that IsSameYear is symmetric
	t1 := time.Date(2026, 1, 1, 10, 30, 0, 0, time.UTC)
	t2 := time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)

	result1 := lxtime.IsSameYear(t1, t2)
	result2 := lxtime.IsSameYear(t2, t1)

	if result1 != result2 {
		t.Errorf("IsSameYear(t1, t2) = %v but IsSameYear(t2, t1) = %v, expected symmetry", result1, result2)
	}
}

func TestIsSameYear_Transitivity(t *testing.T) {
	// Test transitivity: if IsSameYear(a, b) and IsSameYear(b, c) then IsSameYear(a, c)
	t1 := time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC)
	t2 := time.Date(2026, 6, 15, 15, 0, 0, 0, time.UTC)
	t3 := time.Date(2026, 12, 31, 20, 0, 0, 0, time.UTC)

	if !lxtime.IsSameYear(t1, t2) || !lxtime.IsSameYear(t2, t3) {
		t.Fatal("Test setup failed")
	}

	if !lxtime.IsSameYear(t1, t3) {
		t.Error("IsSameYear is not transitive")
	}
}

func TestIsSameYear_LeapYearAndNonLeapYear(t *testing.T) {
	tests := []struct {
		name     string
		year1    int
		year2    int
		expected bool
	}{
		{
			name:     "both leap years",
			year1:    2024,
			year2:    2024,
			expected: true,
		},
		{
			name:     "leap year vs non-leap year",
			year1:    2024,
			year2:    2025,
			expected: false,
		},
		{
			name:     "both non-leap years",
			year1:    2025,
			year2:    2026,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t1 := time.Date(tt.year1, 2, 29, 10, 0, 0, 0, time.UTC)
			t2 := time.Date(tt.year2, 2, 28, 10, 0, 0, 0, time.UTC)

			if tt.year1 != 2024 {
				t1 = time.Date(tt.year1, 2, 28, 10, 0, 0, 0, time.UTC)
			}
			if tt.year2 != 2024 {
				t2 = time.Date(tt.year2, 2, 28, 10, 0, 0, 0, time.UTC)
			}

			result := lxtime.IsSameYear(t1, t2)
			if result != tt.expected {
				t.Errorf("IsSameYear(%d, %d) = %v, want %v", tt.year1, tt.year2, result, tt.expected)
			}
		})
	}
}

func ExampleIsSameYear() {
	t1 := time.Date(2026, 1, 1, 10, 30, 0, 0, time.UTC)
	t2 := time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)
	result := lxtime.IsSameYear(t1, t2)
	fmt.Println(result)

	t3 := time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC)
	result = lxtime.IsSameYear(t1, t3)
	fmt.Println(result)
	// Output:
	// true
	// false
}
