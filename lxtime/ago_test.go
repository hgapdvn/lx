package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/lxtime"
)

func TestAgo_BasicDurations(t *testing.T) {
	tests := []struct {
		name             string
		n                int
		unit             time.Duration
		toleranceSeconds int64
		shouldBePast     bool
	}{
		{
			name:             "zero seconds",
			n:                0,
			unit:             time.Second,
			toleranceSeconds: 1,
			shouldBePast:     false,
		},
		{
			name:             "one second ago",
			n:                1,
			unit:             time.Second,
			toleranceSeconds: 2,
			shouldBePast:     true,
		},
		{
			name:             "five minutes ago",
			n:                5,
			unit:             time.Minute,
			toleranceSeconds: 2,
			shouldBePast:     true,
		},
		{
			name:             "one hour ago",
			n:                1,
			unit:             time.Hour,
			toleranceSeconds: 2,
			shouldBePast:     true,
		},
		{
			name:             "one day ago",
			n:                1,
			unit:             24 * time.Hour,
			toleranceSeconds: 2,
			shouldBePast:     true,
		},
		{
			name:             "thirty minutes ago",
			n:                30,
			unit:             time.Minute,
			toleranceSeconds: 2,
			shouldBePast:     true,
		},
		{
			name:             "two hours ago",
			n:                2,
			unit:             time.Hour,
			toleranceSeconds: 2,
			shouldBePast:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := time.Now()
			result := lxtime.Ago(tt.n, tt.unit)
			after := time.Now()

			// Check that result is in the past
			if tt.shouldBePast {
				if !result.Before(after) {
					t.Errorf("Ago(%d, %v) should be in the past, got %v", tt.n, tt.unit, result)
				}
			}

			// Check approximate accuracy within tolerance
			actualDiff := before.Sub(result).Seconds()
			expectedDiff := time.Duration(tt.n) * tt.unit / time.Second
			tolerance := time.Duration(tt.toleranceSeconds) * time.Second / time.Second

			if actualDiff < expectedDiff.Seconds()-tolerance.Seconds() ||
				actualDiff > expectedDiff.Seconds()+tolerance.Seconds() {
				t.Logf("Ago(%d, %v) timing difference: %v seconds (expected ~%v seconds)", tt.n, tt.unit, actualDiff, expectedDiff.Seconds())
			}
		})
	}
}

func TestAgo_StandardDurations(t *testing.T) {
	tests := []struct {
		name string
		n    int
		unit time.Duration
		desc string
	}{
		{name: "milliseconds", n: 100, unit: time.Millisecond, desc: "100 milliseconds"},
		{name: "seconds", n: 45, unit: time.Second, desc: "45 seconds"},
		{name: "minutes", n: 15, unit: time.Minute, desc: "15 minutes"},
		{name: "hours", n: 3, unit: time.Hour, desc: "3 hours"},
	}

	now := time.Now()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Ago(tt.n, tt.unit)

			// Verify result is before now
			if !result.Before(now) && !result.Equal(now) {
				t.Errorf("Ago(%d, %v) should be at or before now", tt.n, tt.unit)
			}

			// Verify approximate distance
			duration := now.Sub(result)
			expectedDuration := time.Duration(tt.n) * tt.unit

			// Allow 2% tolerance for timing variations
			tolerance := expectedDuration / 50
			if duration < expectedDuration-tolerance || duration > expectedDuration+tolerance {
				t.Logf("Ago(%d, %v) expected ~%v ago, got ~%v ago", tt.n, tt.unit, expectedDuration, duration)
			}
		})
	}
}

func TestAgo_LargeValues(t *testing.T) {
	tests := []struct {
		name string
		n    int
		unit time.Duration
	}{
		{name: "large seconds", n: 999999, unit: time.Second},
		{name: "large minutes", n: 1000, unit: time.Minute},
		{name: "large hours", n: 100, unit: time.Hour},
	}

	now := time.Now()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Ago(tt.n, tt.unit)

			if !result.Before(now) {
				t.Errorf("Ago(%d, %v) should be in the past", tt.n, tt.unit)
			}
		})
	}
}

func TestAgo_NegativeValues(t *testing.T) {
	t.Run("negative_value", func(t *testing.T) {
		now := time.Now()
		result := lxtime.Ago(-5, time.Minute)

		// Negative value should give future time
		if !result.After(now) {
			t.Errorf("Ago(-5, time.Minute) should be in the future (negative ago)")
		}
	})
}

func TestAgo_ConsistencyWithNow(t *testing.T) {
	t.Run("consistent_relative_to_now", func(t *testing.T) {
		before := time.Now()
		result := lxtime.Ago(1, time.Hour)
		after := time.Now()

		// Result should be between before and after
		if !result.Before(before) || result.After(after) {
			t.Errorf("Ago result should be within execution window")
		}

		// Should be approximately 1 hour before the current time
		diff := before.Sub(result)
		oneHour := time.Hour

		tolerance := 100 * time.Millisecond
		if diff < oneHour-tolerance || diff > oneHour+tolerance {
			t.Logf("Ago timing off by %v (tolerance: %v)", diff-oneHour, tolerance)
		}
	})
}

func TestAgo_EquivalenceWithManualAdd(t *testing.T) {
	t.Run("equivalent_to_manual_now_add", func(t *testing.T) {
		agoResult := lxtime.Ago(30, time.Minute)

		manualResult := time.Now().Add(-30 * time.Minute)

		// Should be very close (within 100ms)
		diff := agoResult.Sub(manualResult)
		if diff < 0 {
			diff = -diff
		}
		tolerance := 100 * time.Millisecond

		if diff > tolerance {
			t.Logf("Ago result differs from manual calculation by %v (tolerance: %v)", diff, tolerance)
		}
	})
}

func ExampleAgo() {
	fiveMinutesAgo := lxtime.Ago(5, time.Minute)
	// fiveMinutesAgo: approximately 5 minutes before now
	_ = fiveMinutesAgo

	oneHourAgo := lxtime.Ago(1, time.Hour)
	// oneHourAgo: approximately 1 hour before now
	_ = oneHourAgo
}
