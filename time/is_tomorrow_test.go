package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/time"
)

func TestIsTomorrow_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "tomorrow is tomorrow",
			time:     time.Now().AddDate(0, 0, 1),
			expected: true,
		},
		{
			name:     "tomorrow at midnight",
			time:     time.Now().AddDate(0, 0, 1).Truncate(24 * time.Hour),
			expected: true,
		},
		{
			name:     "tomorrow at noon",
			time:     time.Now().AddDate(0, 0, 1).Truncate(24 * time.Hour).Add(12 * time.Hour),
			expected: true,
		},
		{
			name:     "today is not tomorrow",
			time:     time.Now(),
			expected: false,
		},
		{
			name:     "two days from now is not tomorrow",
			time:     time.Now().AddDate(0, 0, 2),
			expected: false,
		},
		{
			name:     "yesterday is not tomorrow",
			time:     time.Now().AddDate(0, 0, -1),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsTomorrow(tt.time)
			if result != tt.expected {
				t.Errorf("IsTomorrow() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsTomorrow_DifferentTimes(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "midnight of tomorrow is tomorrow",
			check: func() bool {
				tomorrow := time.Now().AddDate(0, 0, 1).Truncate(24 * time.Hour)
				return lxtime.IsTomorrow(tomorrow)
			},
		},
		{
			name: "just after midnight of today is not tomorrow",
			check: func() bool {
				justToday := time.Now().Truncate(24 * time.Hour)
				return !lxtime.IsTomorrow(justToday)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsTomorrow() check failed")
			}
		})
	}
}

func TestIsTomorrow_DifferentDates(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "2 days from now is not tomorrow",
			check: func() bool {
				return !lxtime.IsTomorrow(time.Now().AddDate(0, 0, 2))
			},
		},
		{
			name: "7 days from now is not tomorrow",
			check: func() bool {
				return !lxtime.IsTomorrow(time.Now().AddDate(0, 0, 7))
			},
		},
		{
			name: "today is not tomorrow",
			check: func() bool {
				return !lxtime.IsTomorrow(time.Now())
			},
		},
		{
			name: "yesterday is not tomorrow",
			check: func() bool {
				return !lxtime.IsTomorrow(time.Now().AddDate(0, 0, -1))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsTomorrow() check failed")
			}
		})
	}
}

func TestIsTomorrow_EdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "tomorrow with nanoseconds is tomorrow",
			check: func() bool {
				t := time.Now().AddDate(0, 0, 1).Add(123456789 * time.Nanosecond)
				return lxtime.IsTomorrow(t)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("IsTomorrow() check failed")
			}
		})
	}
}
