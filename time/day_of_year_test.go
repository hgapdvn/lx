package lxtime_test

import (
	"fmt"
	"testing"
	"time"

	lxtime "github.com/hgapdvn/lx/time"
)

func TestDayOfYear_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		t        time.Time
		expected int
	}{
		{
			name:     "january 1st",
			t:        time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "january 31st",
			t:        time.Date(2026, 1, 31, 10, 0, 0, 0, time.UTC),
			expected: 31,
		},
		{
			name:     "february 1st non-leap year",
			t:        time.Date(2026, 2, 1, 10, 0, 0, 0, time.UTC),
			expected: 32,
		},
		{
			name:     "march 1st non-leap year",
			t:        time.Date(2026, 3, 1, 10, 0, 0, 0, time.UTC),
			expected: 60,
		},
		{
			name:     "april 4 2026",
			t:        time.Date(2026, 4, 4, 10, 0, 0, 0, time.UTC),
			expected: 94,
		},
		{
			name:     "june 30 (middle of year)",
			t:        time.Date(2026, 6, 30, 10, 0, 0, 0, time.UTC),
			expected: 181,
		},
		{
			name:     "december 31st non-leap year",
			t:        time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC),
			expected: 365,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.DayOfYear(tt.t)
			if result != tt.expected {
				t.Errorf("DayOfYear(%v) = %d, want %d", tt.t, result, tt.expected)
			}
		})
	}
}

func TestDayOfYear_LeapYear(t *testing.T) {
	tests := []struct {
		name     string
		t        time.Time
		expected int
	}{
		{
			name:     "leap year february 28",
			t:        time.Date(2024, 2, 28, 10, 0, 0, 0, time.UTC),
			expected: 59,
		},
		{
			name:     "leap year february 29",
			t:        time.Date(2024, 2, 29, 10, 0, 0, 0, time.UTC),
			expected: 60,
		},
		{
			name:     "leap year march 1",
			t:        time.Date(2024, 3, 1, 10, 0, 0, 0, time.UTC),
			expected: 61,
		},
		{
			name:     "leap year december 31",
			t:        time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC),
			expected: 366,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.DayOfYear(tt.t)
			if result != tt.expected {
				t.Errorf("DayOfYear(%v) = %d, want %d", tt.t, result, tt.expected)
			}
		})
	}
}

func TestDayOfYear_TimeComponentsIgnored(t *testing.T) {
	// Different times on the same day should return the same day of year
	baseDate := time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC)
	times := []time.Time{
		time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC),
		time.Date(2026, 4, 4, 23, 59, 59, 999999999, time.UTC),
	}

	expectedDay := lxtime.DayOfYear(baseDate)

	for _, testTime := range times {
		result := lxtime.DayOfYear(testTime)
		if result != expectedDay {
			t.Errorf("DayOfYear(%v) = %d, want %d", testTime, result, expectedDay)
		}
	}
}

func TestDayOfYear_SequentialDays(t *testing.T) {
	// Test that day of year increments properly through months
	currentDay := 1
	for month := 1; month <= 12; month++ {
		// Days in each month for 2026 (non-leap year)
		daysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
		for day := 1; day <= daysInMonth[month-1]; day++ {
			testDate := time.Date(2026, time.Month(month), day, 10, 0, 0, 0, time.UTC)
			result := lxtime.DayOfYear(testDate)
			if result != currentDay {
				t.Errorf("DayOfYear(%v) = %d, want %d", testDate, result, currentDay)
			}
			currentDay++
		}
	}
}

func TestDayOfYear_AllYears(t *testing.T) {
	// Test various years to ensure consistency
	tests := []struct {
		year    int
		isLeap  bool
		lastDay int
	}{
		{2020, true, 366},  // Leap year
		{2024, true, 366},  // Leap year
		{2025, false, 365}, // Non-leap year
		{2026, false, 365}, // Non-leap year
		{2027, false, 365}, // Non-leap year
		{2028, true, 366},  // Leap year
	}

	for _, tt := range tests {
		t.Run("year_"+string(rune(tt.year)), func(t *testing.T) {
			lastDay := time.Date(tt.year, 12, 31, 23, 59, 59, 0, time.UTC)
			result := lxtime.DayOfYear(lastDay)
			if result != tt.lastDay {
				t.Errorf("DayOfYear(Dec 31, %d) = %d, want %d", tt.year, result, tt.lastDay)
			}
		})
	}
}

func TestDayOfYear_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		t        time.Time
		expected int
	}{
		{
			name:     "unix epoch january 1 1970",
			t:        time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "year 2000 january 1 (leap year)",
			t:        time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "year 2000 december 31 (leap year)",
			t:        time.Date(2000, 12, 31, 0, 0, 0, 0, time.UTC),
			expected: 366,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.DayOfYear(tt.t)
			if result != tt.expected {
				t.Errorf("DayOfYear(%v) = %d, want %d", tt.t, result, tt.expected)
			}
		})
	}
}

func ExampleDayOfYear() {
	t := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
	day := lxtime.DayOfYear(t)
	fmt.Println(day)

	t2 := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	day2 := lxtime.DayOfYear(t2)
	fmt.Println(day2)
	// Output:
	// 94
	// 1
}
