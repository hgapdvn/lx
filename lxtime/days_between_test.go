package lxtime_test

import (
	"testing"
	"time"

	lxtime "github.com/hgapdvn/lx/lxtime"
)

func TestDaysBetween_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected int
	}{
		{
			name:     "same day",
			t1:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2026, 4, 4, 23, 59, 59, 0, time.UTC),
			expected: 0,
		},
		{
			name:     "adjacent days",
			t1:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2026, 4, 5, 10, 30, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "two days apart",
			t1:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2026, 4, 6, 20, 45, 0, 0, time.UTC),
			expected: 2,
		},
		{
			name:     "week apart",
			t1:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2026, 4, 11, 10, 30, 0, 0, time.UTC),
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.DaysBetween(tt.t1, tt.t2)
			if result != tt.expected {
				t.Errorf("DaysBetween(%v, %v) = %d, want %d", tt.t1, tt.t2, result, tt.expected)
			}
		})
	}
}

func TestDaysBetween_Negative(t *testing.T) {
	tests := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected int
	}{
		{
			name:     "reversed same day",
			t1:       time.Date(2026, 4, 4, 23, 59, 59, 0, time.UTC),
			t2:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			expected: 0,
		},
		{
			name:     "reversed adjacent days",
			t1:       time.Date(2026, 4, 5, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			expected: -1,
		},
		{
			name:     "reversed two days apart",
			t1:       time.Date(2026, 4, 6, 20, 45, 0, 0, time.UTC),
			t2:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			expected: -2,
		},
		{
			name:     "reversed week",
			t1:       time.Date(2026, 4, 11, 10, 30, 0, 0, time.UTC),
			t2:       time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC),
			expected: -7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.DaysBetween(tt.t1, tt.t2)
			if result != tt.expected {
				t.Errorf("DaysBetween(%v, %v) = %d, want %d", tt.t1, tt.t2, result, tt.expected)
			}
		})
	}
}

func TestDaysBetween_EdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "midnight to midnight next day",
			check: func() bool {
				t1 := time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC)
				t2 := time.Date(2026, 4, 5, 0, 0, 0, 0, time.UTC)
				return lxtime.DaysBetween(t1, t2) == 1
			},
		},
		{
			name: "just before midnight to just after midnight",
			check: func() bool {
				t1 := time.Date(2026, 4, 4, 23, 59, 59, 999999999, time.UTC)
				t2 := time.Date(2026, 4, 5, 0, 0, 0, 0, time.UTC)
				return lxtime.DaysBetween(t1, t2) == 1
			},
		},
		{
			name: "noon to noon different days",
			check: func() bool {
				t1 := time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC)
				t2 := time.Date(2026, 4, 5, 12, 0, 0, 0, time.UTC)
				return lxtime.DaysBetween(t1, t2) == 1
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("DaysBetween() check failed")
			}
		})
	}
}

func TestDaysBetween_MonthBoundary(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "end of month to start of next month",
			check: func() bool {
				t1 := time.Date(2026, 4, 30, 10, 30, 0, 0, time.UTC)
				t2 := time.Date(2026, 5, 1, 10, 30, 0, 0, time.UTC)
				return lxtime.DaysBetween(t1, t2) == 1
			},
		},
		{
			name: "within month boundary",
			check: func() bool {
				t1 := time.Date(2026, 4, 28, 10, 30, 0, 0, time.UTC)
				t2 := time.Date(2026, 5, 2, 10, 30, 0, 0, time.UTC)
				return lxtime.DaysBetween(t1, t2) == 4
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("DaysBetween() check failed")
			}
		})
	}
}

func TestDaysBetween_YearBoundary(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "end of year to start of next year",
			check: func() bool {
				t1 := time.Date(2025, 12, 31, 10, 30, 0, 0, time.UTC)
				t2 := time.Date(2026, 1, 1, 10, 30, 0, 0, time.UTC)
				return lxtime.DaysBetween(t1, t2) == 1
			},
		},
		{
			name: "across year boundary",
			check: func() bool {
				t1 := time.Date(2025, 12, 28, 10, 30, 0, 0, time.UTC)
				t2 := time.Date(2026, 1, 3, 10, 30, 0, 0, time.UTC)
				return lxtime.DaysBetween(t1, t2) == 6
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("DaysBetween() check failed")
			}
		})
	}
}

func TestDaysBetween_LargeValues(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "30 days apart",
			check: func() bool {
				t1 := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
				t2 := time.Date(2026, 5, 4, 10, 30, 0, 0, time.UTC)
				return lxtime.DaysBetween(t1, t2) == 30
			},
		},
		{
			name: "100 days apart",
			check: func() bool {
				t1 := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
				t2 := time.Date(2026, 4, 11, 0, 0, 0, 0, time.UTC)
				return lxtime.DaysBetween(t1, t2) == 100
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("DaysBetween() check failed")
			}
		})
	}
}

func TestDaysBetween_TimeComponentIgnored(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "time component doesn't affect same day",
			check: func() bool {
				t1 := time.Date(2026, 4, 4, 0, 0, 0, 0, time.UTC)
				t2 := time.Date(2026, 4, 4, 23, 59, 59, 999999999, time.UTC)
				return lxtime.DaysBetween(t1, t2) == 0
			},
		},
		{
			name: "only day difference matters",
			check: func() bool {
				t1 := time.Date(2026, 4, 4, 23, 59, 59, 0, time.UTC)
				t2 := time.Date(2026, 4, 5, 0, 0, 0, 1, time.UTC)
				return lxtime.DaysBetween(t1, t2) == 1
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("DaysBetween() check failed")
			}
		})
	}
}

func TestDaysBetween_AntiCommutative(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "swapping arguments negates result",
			check: func() bool {
				t1 := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
				t2 := time.Date(2026, 4, 7, 20, 45, 0, 0, time.UTC)
				forward := lxtime.DaysBetween(t1, t2)
				backward := lxtime.DaysBetween(t2, t1)
				return forward == -backward
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("DaysBetween() anti-commutative check failed")
			}
		})
	}
}

func TestDaysBetween_WithExistingFunctions(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "today and today have 0 days between",
			check: func() bool {
				t1 := time.Now()
				t2 := time.Now()
				return lxtime.DaysBetween(t1, t2) == 0
			},
		},
		{
			name: "IsSameDay equivalent to DaysBetween == 0",
			check: func() bool {
				t1 := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
				t2 := time.Date(2026, 4, 4, 20, 45, 0, 0, time.UTC)
				isSame := lxtime.IsSameDay(t1, t2)
				daysBetween := lxtime.DaysBetween(t1, t2) == 0
				return isSame == daysBetween
			},
		},
		{
			name: "different days implies non-zero DaysBetween",
			check: func() bool {
				t1 := time.Date(2026, 4, 4, 10, 30, 0, 0, time.UTC)
				t2 := time.Date(2026, 4, 5, 10, 30, 0, 0, time.UTC)
				isSame := lxtime.IsSameDay(t1, t2)
				daysBetween := lxtime.DaysBetween(t1, t2) != 0
				return isSame != daysBetween
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("DaysBetween() with existing functions check failed")
			}
		})
	}
}

func TestDaysBetween_TimezoneFix(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "UTC to JST different timezone correct calculation",
			check: func() bool {
				// This is the example from the bug report
				// t1 = 2026-04-04 10:00 UTC
				// t2 = 2026-04-06 10:00 JST (= 2026-04-06 01:00 UTC)
				jst, _ := time.LoadLocation("Asia/Tokyo")
				t1 := time.Date(2026, 4, 4, 10, 0, 0, 0, time.UTC)
				t2 := time.Date(2026, 4, 6, 10, 0, 0, 0, jst)

				// Should be 2 days, not 1
				return lxtime.DaysBetween(t1, t2) == 2
			},
		},
		{
			name: "EST to PST different timezone",
			check: func() bool {
				est, _ := time.LoadLocation("America/New_York")
				pst, _ := time.LoadLocation("America/Los_Angeles")
				t1 := time.Date(2026, 4, 4, 10, 0, 0, 0, est)
				t2 := time.Date(2026, 4, 6, 10, 0, 0, 0, pst)
				// Both are same dates in calendar terms
				return lxtime.DaysBetween(t1, t2) == 2
			},
		},
		{
			name: "same instant different timezones",
			check: func() bool {
				est, _ := time.LoadLocation("America/New_York")
				pst, _ := time.LoadLocation("America/Los_Angeles")

				// Same instant in time (2026-04-04 14:00 UTC)
				// EST: 2026-04-04 10:00 (UTC-4)
				// PST: 2026-04-04 07:00 (UTC-7)
				t1 := time.Date(2026, 4, 4, 10, 0, 0, 0, est)
				t2 := time.Date(2026, 4, 4, 7, 0, 0, 0, pst)

				// Same day in both timezones
				return lxtime.DaysBetween(t1, t2) == 0
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("DaysBetween() timezone fix check failed")
			}
		})
	}
}
