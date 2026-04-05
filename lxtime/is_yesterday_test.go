package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/lxtime"
)

func TestIsYesterday_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "yesterday is yesterday",
			time:     time.Now().AddDate(0, 0, -1),
			expected: true,
		},
		{
			name:     "yesterday at noon",
			time:     time.Now().AddDate(0, 0, -1).Truncate(24 * time.Hour).Add(12 * time.Hour),
			expected: true,
		},
		{
			name:     "today is not yesterday",
			time:     time.Now(),
			expected: false,
		},
		{
			name:     "two days ago is not yesterday",
			time:     time.Now().AddDate(0, 0, -2),
			expected: false,
		},
		{
			name:     "tomorrow is not yesterday",
			time:     time.Now().AddDate(0, 0, 1),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsYesterday(tt.time)
			if result != tt.expected {
				t.Errorf("IsYesterday() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsYesterday_DifferentTimes(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "midnight of yesterday is yesterday",
			check: func() bool {
				yesterday := time.Now().AddDate(0, 0, -1).Truncate(24 * time.Hour)
				return lxtime.IsYesterday(yesterday)
			},
		},
		{
			name: "almost midnight of today is not yesterday",
			check: func() bool {
				almostToday := time.Now().Truncate(24 * time.Hour).Add(24*time.Hour - 1*time.Nanosecond)
				return !lxtime.IsYesterday(almostToday)
			},
		},
		{
			name: "just after midnight of today is not yesterday",
			check: func() bool {
				justToday := time.Now().Truncate(24 * time.Hour)
				return !lxtime.IsYesterday(justToday)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsYesterday() check failed")
			}
		})
	}
}

func TestIsYesterday_DifferentDates(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "2 days ago is not yesterday",
			check: func() bool {
				return !lxtime.IsYesterday(time.Now().AddDate(0, 0, -2))
			},
		},
		{
			name: "7 days ago is not yesterday",
			check: func() bool {
				return !lxtime.IsYesterday(time.Now().AddDate(0, 0, -7))
			},
		},
		{
			name: "today is not yesterday",
			check: func() bool {
				return !lxtime.IsYesterday(time.Now())
			},
		},
		{
			name: "tomorrow is not yesterday",
			check: func() bool {
				return !lxtime.IsYesterday(time.Now().AddDate(0, 0, 1))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsYesterday() check failed")
			}
		})
	}
}

func TestIsYesterday_EdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "yesterday with nanoseconds is yesterday",
			check: func() bool {
				t := time.Now().AddDate(0, 0, -1).Add(123456789 * time.Nanosecond)
				return lxtime.IsYesterday(t)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsYesterday() check failed")
			}
		})
	}
}
