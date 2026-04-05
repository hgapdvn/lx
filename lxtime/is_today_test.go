package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/lxtime"
)

func TestIsToday_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "current time is today",
			time:     time.Now(),
			expected: true,
		},
		{
			name:     "today at noon",
			time:     time.Now().Truncate(24 * time.Hour).Add(12 * time.Hour),
			expected: true,
		},
		{
			name:     "yesterday is not today",
			time:     time.Now().AddDate(0, 0, -1),
			expected: false,
		},
		{
			name:     "tomorrow is not today",
			time:     time.Now().AddDate(0, 0, 1),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsToday(tt.time)
			if result != tt.expected {
				t.Errorf("IsToday() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsToday_DifferentTimes(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "midnight of today is today",
			check: func() bool {
				today := time.Now().Truncate(24 * time.Hour)
				return lxtime.IsToday(today)
			},
		},
		{
			name: "just past midnight of tomorrow is not today",
			check: func() bool {
				justTomorrow := time.Now().Truncate(24 * time.Hour).Add(24 * time.Hour)
				return !lxtime.IsToday(justTomorrow)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsToday() check failed")
			}
		})
	}
}

func TestIsToday_DifferentDates(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "1 day ago is not today",
			check: func() bool {
				return !lxtime.IsToday(time.Now().AddDate(0, 0, -1))
			},
		},
		{
			name: "7 days ago is not today",
			check: func() bool {
				return !lxtime.IsToday(time.Now().AddDate(0, 0, -7))
			},
		},
		{
			name: "30 days ago is not today",
			check: func() bool {
				return !lxtime.IsToday(time.Now().AddDate(0, 0, -30))
			},
		},
		{
			name: "1 day in future is not today",
			check: func() bool {
				return !lxtime.IsToday(time.Now().AddDate(0, 0, 1))
			},
		},
		{
			name: "7 days in future is not today",
			check: func() bool {
				return !lxtime.IsToday(time.Now().AddDate(0, 0, 7))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsToday() check failed")
			}
		})
	}
}

func TestIsToday_DifferentTimezones(t *testing.T) {
	utc := time.UTC
	est, _ := time.LoadLocation("America/New_York")
	pst, _ := time.LoadLocation("America/Los_Angeles")

	// Use a time in the middle of the day to avoid timezone boundary issues
	now := time.Now()
	midday := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, utc)

	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "today in UTC is today",
			check: func() bool {
				testTime := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, utc)
				return lxtime.IsToday(testTime)
			},
		},
		{
			name: "today in EST is today",
			check: func() bool {
				testTime := midday.In(est)
				return lxtime.IsToday(testTime)
			},
		},
		{
			name: "today in PST is today",
			check: func() bool {
				testTime := midday.In(pst)
				return lxtime.IsToday(testTime)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsToday() check failed")
			}
		})
	}
}

func TestIsToday_EdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "today with nanoseconds is today",
			check: func() bool {
				t := time.Now().Add(123456789 * time.Nanosecond)
				return lxtime.IsToday(t)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsToday() check failed")
			}
		})
	}
}
