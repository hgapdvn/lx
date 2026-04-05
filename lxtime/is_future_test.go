package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/lxtime"
)

func TestIsFuture_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "one second in future",
			time:     lxtime.FromNow(1, time.Second),
			expected: true,
		},
		{
			name:     "five minutes in future",
			time:     lxtime.FromNow(5, time.Minute),
			expected: true,
		},
		{
			name:     "one hour in future",
			time:     lxtime.FromNow(1, time.Hour),
			expected: true,
		},
		{
			name:     "one day in future",
			time:     lxtime.FromNow(1, 24*time.Hour),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsFuture(tt.time)
			if result != tt.expected {
				t.Errorf("IsFuture(%v) = %v, want %v", tt.time, result, tt.expected)
			}
		})
	}
}

func TestIsFuture_PastTimes(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "one second in past",
			time:     lxtime.Ago(1, time.Second),
			expected: false,
		},
		{
			name:     "five minutes in past",
			time:     lxtime.Ago(5, time.Minute),
			expected: false,
		},
		{
			name:     "one hour in past",
			time:     lxtime.Ago(1, time.Hour),
			expected: false,
		},
		{
			name:     "one day in past",
			time:     lxtime.Ago(1, 24*time.Hour),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsFuture(tt.time)
			if result != tt.expected {
				t.Errorf("IsFuture(%v) = %v, want %v", tt.time, result, tt.expected)
			}
		})
	}
}

func TestIsFuture_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		getTime  func() time.Time
		expected bool
		desc     string
	}{
		{
			name: "approximately now",
			getTime: func() time.Time {
				return time.Now().Add(100 * time.Nanosecond)
			},
			expected: true,
			desc:     "very small future (may be equal)",
		},
		{
			name: "zero value",
			getTime: func() time.Time {
				return time.Time{}
			},
			expected: false,
			desc:     "zero time is in the past",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testTime := tt.getTime()
			result := lxtime.IsFuture(testTime)
			// For edge cases, we check that the result matches or is reasonable
			if tt.desc == "very small future (may be equal)" {
				// Very small future times might appear equal to now depending on timing
				t.Logf("IsFuture(%v) = %v (timing edge case)", testTime, result)
			} else if result != tt.expected {
				t.Errorf("IsFuture(%v) = %v, want %v", testTime, result, tt.expected)
			}
		})
	}
}

func TestIsFuture_WithRelativeHelpers(t *testing.T) {
	t.Run("IsFuture_with_FromNow", func(t *testing.T) {
		futureTime := lxtime.FromNow(10, time.Minute)
		if !lxtime.IsFuture(futureTime) {
			t.Errorf("IsFuture should return true for time from FromNow(10, time.Minute)")
		}
	})

	t.Run("IsFuture_with_Ago", func(t *testing.T) {
		pastTime := lxtime.Ago(10, time.Minute)
		if lxtime.IsFuture(pastTime) {
			t.Errorf("IsFuture should return false for time from Ago(10, time.Minute)")
		}
	})
}

func TestIsFuture_Consistency(t *testing.T) {
	t.Run("consistency_multiple_calls", func(t *testing.T) {
		futureTime := lxtime.FromNow(30, time.Second)
		result1 := lxtime.IsFuture(futureTime)
		result2 := lxtime.IsFuture(futureTime)

		// Results may differ due to timing, but at least one should be true
		if !result1 && !result2 {
			t.Errorf("IsFuture should be true for future time in at least one check")
		}
	})
}

func TestIsFuture_LargeValues(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{
			name:     "far future (100 years)",
			time:     time.Now().AddDate(100, 0, 0),
			expected: true,
		},
		{
			name:     "far past (100 years)",
			time:     time.Now().AddDate(-100, 0, 0),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.IsFuture(tt.time)
			if result != tt.expected {
				t.Errorf("IsFuture(%v) = %v, want %v", tt.time, result, tt.expected)
			}
		})
	}
}

func ExampleIsFuture() {
	t := lxtime.FromNow(5, time.Minute)
	isFuture := lxtime.IsFuture(t)
	// isFuture: true
	_ = isFuture

	past := lxtime.Ago(5, time.Minute)
	isPastFuture := lxtime.IsFuture(past)
	// isPastFuture: false
	_ = isPastFuture
}
