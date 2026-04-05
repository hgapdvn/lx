package lxtime_test

import (
	"testing"
	"time"

	lxtime "github.com/hgapdvn/lx/lxtime"
)

func TestIsSameDay_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected bool
	}{
		{
			name:     "same date same time",
			t1:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "same date different times",
			t1:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2026, 4, 4, 23, 59, 59, 0, time.UTC),
			expected: true,
		},
		{
			name:     "different dates",
			t1:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2026, 4, 5, 10, 30, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "one day apart",
			t1:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2026, 4, 5, 9, 59, 59, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsSameDay(tt.t1, tt.t2)
			if result != tt.expected {
				t.Errorf("IsSameDay(%v, %v) = %v, want %v", tt.t1, tt.t2, result, tt.expected)
			}
		})
	}
}

func TestIsSameDay_EdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "midnight to end of day",
			check: func() bool {
				t1 := time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC)
				t2 := time.Date(2026, 4, 4, 23, 59, 59, 999999999, time.UTC)
				return lxtime.IsSameDay(t1, t2)
			},
		},
		{
			name: "just before midnight vs just after midnight",
			check: func() bool {
				t1 := time.Date(2026, 4, 4, 23, 59, 59, 999999999, time.UTC)
				t2 := time.Date(2026, 4, 5, 0, 0, 0, 0, time.UTC)
				return !lxtime.IsSameDay(t1, t2)
			},
		},
		{
			name: "noon to noon same day",
			check: func() bool {
				t1 := time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC)
				t2 := time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC)
				return lxtime.IsSameDay(t1, t2)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsSameDay() check failed")
			}
		})
	}
}

func TestIsSameDay_DifferentMonths(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "end of month to start of next month",
			check: func() bool {
				t1 := time.Date(2026, 4, 30, 10, 30, 0, 0, time.UTC)
				t2 := time.Date(2026, 5, 1, 10, 30, 0, 0, time.UTC)
				return !lxtime.IsSameDay(t1, t2)
			},
		},
		{
			name: "same day of different months",
			check: func() bool {
				t1 := time.Date(2026, 3, 15, 10, 30, 0, 0, time.UTC)
				t2 := time.Date(2026, 4, 15, 10, 30, 0, 0, time.UTC)
				return !lxtime.IsSameDay(t1, t2)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsSameDay() check failed")
			}
		})
	}
}

func TestIsSameDay_DifferentYears(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "end of year to start of next year",
			check: func() bool {
				t1 := time.Date(2025, 12, 31, 10, 30, 0, 0, time.UTC)
				t2 := time.Date(2026, 1, 1, 10, 30, 0, 0, time.UTC)
				return !lxtime.IsSameDay(t1, t2)
			},
		},
		{
			name: "same day of different years",
			check: func() bool {
				t1 := time.Date(2025, 4, 4, 10, 30, 0, 0, time.UTC)
				t2 := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
				return !lxtime.IsSameDay(t1, t2)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsSameDay() check failed")
			}
		})
	}
}

func TestIsSameDay_Timezones(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "same UTC day",
			check: func() bool {
				t1 := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
				t2 := time.Date(2026, 4, 4, 20, 45, 30, 0, time.UTC)
				return lxtime.IsSameDay(t1, t2)
			},
		},
		{
			name: "same day in different timezones",
			check: func() bool {
				est, _ := time.LoadLocation("America/New_York")
				pst, _ := time.LoadLocation("America/Los_Angeles")
				t1 := time.Date(2026, 4, 4, 10, 30, 0, 0, est)
				t2 := time.Date(2026, 4, 4, 10, 30, 0, 0, pst)
				return lxtime.IsSameDay(t1, t2)
			},
		},
		{
			name: "different days due to timezone",
			check: func() bool {
				utc := time.UTC
				est, _ := time.LoadLocation("America/New_York")
				// April 5, 2026 00:01 UTC = April 4, 2026 20:01 EST
				t1 := time.Date(2026, 4, 5, 0, 1, 0, 0, utc)
				t2 := time.Date(2026, 4, 4, 20, 1, 0, 0, est)
				// These are the same instant but different dates in their timezones
				return !lxtime.IsSameDay(t1, t2)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsSameDay() check failed")
			}
		})
	}
}

func TestIsSameDay_Nanoseconds(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "different nanoseconds same day",
			check: func() bool {
				t1 := time.Date(2026, 4, 4, 10, 30, 0, 123456789, time.UTC)
				t2 := time.Date(2026, 4, 4, 10, 30, 0, 987654321, time.UTC)
				return lxtime.IsSameDay(t1, t2)
			},
		},
		{
			name: "nanoseconds irrelevant for day comparison",
			check: func() bool {
				t1 := time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC)
				t2 := time.Date(2026, 4, 4, 23, 59, 59, 999999999, time.UTC)
				return lxtime.IsSameDay(t1, t2)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsSameDay() check failed")
			}
		})
	}
}

func TestIsSameDay_Commutative(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "order doesn't matter - same day",
			check: func() bool {
				t1 := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
				t2 := time.Date(2026, 4, 4, 20, 45, 0, 0, time.UTC)
				return lxtime.IsSameDay(t1, t2) == lxtime.IsSameDay(t2, t1)
			},
		},
		{
			name: "order doesn't matter - different days",
			check: func() bool {
				t1 := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
				t2 := time.Date(2026, 4, 5, 10, 30, 0, 0, time.UTC)
				return lxtime.IsSameDay(t1, t2) == lxtime.IsSameDay(t2, t1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsSameDay() commutativity check failed")
			}
		})
	}
}

func TestIsSameDay_Reflexive(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "time is same day as itself",
			check: func() bool {
				t := time.Date(2026, 4, 4, 10, 30, 45, 123456789, time.UTC)
				return lxtime.IsSameDay(t, t)
			},
		},
		{
			name: "same date different time of day",
			check: func() bool {
				t1 := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
				t2 := time.Date(2026, 4, 4, 15, 45, 30, 999999999, time.UTC)
				return lxtime.IsSameDay(t1, t2) && lxtime.IsSameDay(t1, t1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsSameDay() reflexive check failed")
			}
		})
	}
}

func TestIsSameDay_WithExistingFunctions(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "IsToday implies IsSameDay with now",
			check: func() bool {
				now := time.Now()
				today := time.Now()
				return lxtime.IsSameDay(now, today)
			},
		},
		{
			name: "yesterday and today are different days",
			check: func() bool {
				yesterday := time.Now().AddDate(0, 0, -1)
				today := time.Now()
				return !lxtime.IsSameDay(yesterday, today)
			},
		},
		{
			name: "two times same day have IsSameDay true",
			check: func() bool {
				t1 := time.Date(2026, 4, 4, 10, 0, 0, 0, time.UTC)
				t2 := time.Date(2026, 4, 4, 20, 0, 0, 0, time.UTC)
				return lxtime.IsSameDay(t1, t2)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsSameDay() with existing functions check failed")
			}
		})
	}
}
