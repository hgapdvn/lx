package lxtime_test

import (
	"testing"
	"time"

	lxtime "github.com/hgapdvn/lx/lxtime"
)

func TestWeekOfYear_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		t        time.Time
		expected int
	}{
		{
			name:     "january 1 2026 (thursday)",
			t:        time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "january 5 2026 (monday, first monday)",
			t:        time.Date(2026, 1, 5, 10, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "january 10 2026 (saturday, week 1)",
			t:        time.Date(2026, 1, 10, 10, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "january 12 2026 (monday of week 2)",
			t:        time.Date(2026, 1, 12, 10, 0, 0, 0, time.UTC),
			expected: 2,
		},
		{
			name:     "april 4 2026 (saturday)",
			t:        time.Date(2026, 4, 4, 10, 0, 0, 0, time.UTC),
			expected: 13,
		},
		{
			name:     "december 28 2026 (monday)",
			t:        time.Date(2026, 12, 28, 10, 0, 0, 0, time.UTC),
			expected: 52,
		},
		{
			name:     "december 31 2026 (thursday)",
			t:        time.Date(2026, 12, 31, 10, 0, 0, 0, time.UTC),
			expected: 52,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.WeekOfYear(tt.t)
			if result != tt.expected {
				t.Errorf("WeekOfYear(%v) = %d, want %d", tt.t, result, tt.expected)
			}
		})
	}
}

func TestWeekOfYear_FullYear2026(t *testing.T) {
	// Verify week number progression through a full year
	// 2026 starts on Thursday (week 1), so first Monday is Jan 5
	currentWeek := 1
	currentMonday := time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC)

	for i := 0; i < 53; i++ {
		// Test each day of the week for this week
		for day := 0; day < 7; day++ {
			testDate := currentMonday.AddDate(0, 0, day)

			// Skip if we've gone past the year
			if testDate.Year() > 2026 {
				return
			}

			result := lxtime.WeekOfYear(testDate)
			if result != currentWeek {
				t.Errorf("WeekOfYear(%v) = %d, want %d", testDate, result, currentWeek)
			}
		}

		// Move to next Monday
		currentMonday = currentMonday.AddDate(0, 0, 7)
		currentWeek++

		// Stop if we're past the year
		if currentMonday.Year() > 2026 {
			break
		}
	}
}

func TestWeekOfYear_AllDaysInWeekReturnSameWeekNumber(t *testing.T) {
	// All 7 days in a week should return the same week number
	monday := time.Date(2026, 4, 6, 10, 0, 0, 0, time.UTC)

	expectedWeek := lxtime.WeekOfYear(monday)

	for day := 0; day < 7; day++ {
		testDate := monday.AddDate(0, 0, day)
		result := lxtime.WeekOfYear(testDate)
		if result != expectedWeek {
			t.Errorf("WeekOfYear(%v [day %d of week]) = %d, want %d", testDate, day, result, expectedWeek)
		}
	}
}

func TestWeekOfYear_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		t        time.Time
		expected int
	}{
		{
			name:     "year start before first monday",
			t:        time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "year end after last monday",
			t:        time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
			expected: 52,
		},
		{
			name:     "leap year february 29",
			t:        time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC),
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.WeekOfYear(tt.t)
			if result != tt.expected {
				t.Errorf("WeekOfYear(%v) = %d, want %d", tt.t, result, tt.expected)
			}
		})
	}
}

func TestWeekOfYear_TimeDoesNotMatter(t *testing.T) {
	// Different times on the same day should return the same week
	date := time.Date(2026, 4, 6, 0, 0, 0, 0, time.UTC) // Monday
	times := []time.Time{
		time.Date(2026, 4, 6, 0, 0, 0, 0, time.UTC),
		time.Date(2026, 4, 6, 12, 0, 0, 0, time.UTC),
		time.Date(2026, 4, 6, 23, 59, 59, 0, time.UTC),
	}

	expectedWeek := lxtime.WeekOfYear(date)

	for _, testTime := range times {
		result := lxtime.WeekOfYear(testTime)
		if result != expectedWeek {
			t.Errorf("WeekOfYear(%v) = %d, want %d", testTime, result, expectedWeek)
		}
	}
}
