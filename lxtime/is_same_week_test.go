package lxtime_test

import (
	"fmt"
	"testing"
	"time"

	lxtime "github.com/hgapdvn/lx/lxtime"
)

func TestIsSameWeek_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected bool
	}{
		{
			name:     "same week monday to wednesday",
			t1:       time.Date(2026, 4, 6, 10, 30, 0, 0, time.UTC), // Monday
			t2:       time.Date(2026, 4, 8, 15, 45, 0, 0, time.UTC), // Wednesday
			expected: true,
		},
		{
			name:     "same week monday to sunday",
			t1:       time.Date(2026, 4, 6, 10, 30, 0, 0, time.UTC),   // Monday
			t2:       time.Date(2026, 4, 12, 23, 59, 59, 0, time.UTC), // Sunday
			expected: true,
		},
		{
			name:     "same day same time",
			t1:       time.Date(2026, 4, 6, 10, 30, 0, 0, time.UTC), // Monday
			t2:       time.Date(2026, 4, 6, 10, 30, 0, 0, time.UTC), // Monday
			expected: true,
		},
		{
			name:     "different weeks monday to monday",
			t1:       time.Date(2026, 4, 6, 10, 30, 0, 0, time.UTC),  // Monday of week 1
			t2:       time.Date(2026, 4, 13, 10, 30, 0, 0, time.UTC), // Monday of week 2
			expected: false,
		},
		{
			name:     "different weeks sunday to monday",
			t1:       time.Date(2026, 4, 12, 23, 59, 59, 0, time.UTC), // Sunday of week 1
			t2:       time.Date(2026, 4, 13, 0, 0, 0, 0, time.UTC),    // Monday of week 2
			expected: false,
		},
		{
			name:     "same week different times",
			t1:       time.Date(2026, 4, 6, 1, 0, 0, 0, time.UTC),     // Monday
			t2:       time.Date(2026, 4, 12, 23, 59, 59, 0, time.UTC), // Sunday
			expected: true,
		},
		{
			name:     "across month boundaries same week",
			t1:       time.Date(2026, 3, 30, 10, 30, 0, 0, time.UTC), // Monday (last week of March)
			t2:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),  // Saturday
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsSameWeek(tt.t1, tt.t2)
			if result != tt.expected {
				t.Errorf("IsSameWeek(%v, %v) = %v, want %v", tt.t1, tt.t2, result, tt.expected)
			}
		})
	}
}

func TestIsSameWeek_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected bool
	}{
		{
			name:     "year boundary same week",
			t1:       time.Date(2025, 12, 29, 10, 0, 0, 0, time.UTC), // Monday
			t2:       time.Date(2026, 1, 2, 10, 0, 0, 0, time.UTC),   // Friday
			expected: true,
		},
		{
			name:     "year boundary different weeks",
			t1:       time.Date(2025, 12, 28, 10, 0, 0, 0, time.UTC), // Sunday
			t2:       time.Date(2025, 12, 29, 10, 0, 0, 0, time.UTC), // Monday
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsSameWeek(tt.t1, tt.t2)
			if result != tt.expected {
				t.Errorf("IsSameWeek(%v, %v) = %v, want %v", tt.t1, tt.t2, result, tt.expected)
			}
		})
	}
}

func TestIsSameWeek_Symmetry(t *testing.T) {
	// Test that IsSameWeek is symmetric: if IsSameWeek(a, b) then IsSameWeek(b, a)
	t1 := time.Date(2026, 4, 6, 10, 30, 0, 0, time.UTC)
	t2 := time.Date(2026, 4, 8, 15, 45, 0, 0, time.UTC)

	result1 := lxtime.IsSameWeek(t1, t2)
	result2 := lxtime.IsSameWeek(t2, t1)

	if result1 != result2 {
		t.Errorf("IsSameWeek(t1, t2) = %v but IsSameWeek(t2, t1) = %v, expected symmetry", result1, result2)
	}
}

func ExampleIsSameWeek() {
	t1 := time.Date(2026, 4, 6, 10, 30, 0, 0, time.UTC) // Monday
	t2 := time.Date(2026, 4, 8, 15, 45, 0, 0, time.UTC) // Wednesday
	result := lxtime.IsSameWeek(t1, t2)
	fmt.Println(result)

	t3 := time.Date(2026, 4, 13, 10, 0, 0, 0, time.UTC) // Monday of next week
	result = lxtime.IsSameWeek(t1, t3)
	fmt.Println(result)
	// Output:
	// true
	// false
}
