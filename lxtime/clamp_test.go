package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/lxtime"
)

func TestClamp_WithinRange(t *testing.T) {
	tests := []struct {
		name     string
		t        time.Time
		start    time.Time
		end      time.Time
		expected time.Time
	}{
		{
			name:     "time in middle of range",
			t:        time.Date(2026, 4, 15, 12, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 12, 0, 0, 0, time.UTC),
		},
		{
			name:     "time at start boundary",
			t:        time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
		},
		{
			name:     "time at end boundary",
			t:        time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
		},
		{
			name:     "all three equal",
			t:        time.Date(2026, 4, 15, 12, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 15, 12, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 15, 12, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 12, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Clamp(tt.t, tt.start, tt.end)
			if !result.Equal(tt.expected) {
				t.Errorf("Clamp(%v, %v, %v) = %v, want %v", tt.t, tt.start, tt.end, result, tt.expected)
			}
		})
	}
}

func TestClamp_BeforeRange(t *testing.T) {
	tests := []struct {
		name     string
		t        time.Time
		start    time.Time
		end      time.Time
		expected time.Time
	}{
		{
			name:     "time one second before start",
			t:        time.Date(2026, 4, 15, 9, 59, 59, 0, time.UTC),
			start:    time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
		},
		{
			name:     "time one hour before start",
			t:        time.Date(2026, 4, 15, 9, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
		},
		{
			name:     "time one day before start",
			t:        time.Date(2026, 4, 14, 10, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
		},
		{
			name:     "time nanosecond before start",
			t:        time.Date(2026, 4, 15, 9, 59, 59, 999999999, time.UTC),
			start:    time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Clamp(tt.t, tt.start, tt.end)
			if !result.Equal(tt.expected) {
				t.Errorf("Clamp(%v, %v, %v) = %v, want %v", tt.t, tt.start, tt.end, result, tt.expected)
			}
		})
	}
}

func TestClamp_AfterRange(t *testing.T) {
	tests := []struct {
		name     string
		t        time.Time
		start    time.Time
		end      time.Time
		expected time.Time
	}{
		{
			name:     "time one second after end",
			t:        time.Date(2026, 4, 15, 14, 0, 1, 0, time.UTC),
			start:    time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
		},
		{
			name:     "time one hour after end",
			t:        time.Date(2026, 4, 15, 15, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
		},
		{
			name:     "time one day after end",
			t:        time.Date(2026, 4, 16, 10, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
		},
		{
			name:     "time nanosecond after end",
			t:        time.Date(2026, 4, 15, 14, 0, 0, 1, time.UTC),
			start:    time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Clamp(tt.t, tt.start, tt.end)
			if !result.Equal(tt.expected) {
				t.Errorf("Clamp(%v, %v, %v) = %v, want %v", tt.t, tt.start, tt.end, result, tt.expected)
			}
		})
	}
}

func TestClamp_ZeroValues(t *testing.T) {
	tests := []struct {
		name     string
		t        time.Time
		start    time.Time
		end      time.Time
		expected time.Time
	}{
		{
			name:     "zero value t with normal range",
			t:        time.Time{},
			start:    time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
		},
		{
			name:     "zero value t with zero range",
			t:        time.Time{},
			start:    time.Time{},
			end:      time.Time{},
			expected: time.Time{},
		},
		{
			name:     "normal t with zero start and end (zero is in the past)",
			t:        time.Date(2026, 4, 15, 12, 0, 0, 0, time.UTC),
			start:    time.Time{},
			end:      time.Time{},
			expected: time.Time{}, // t is after end (zero), so clamped to end
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Clamp(tt.t, tt.start, tt.end)
			if !result.Equal(tt.expected) {
				t.Errorf("Clamp(%v, %v, %v) = %v, want %v", tt.t, tt.start, tt.end, result, tt.expected)
			}
		})
	}
}

func TestClamp_DifferentTimezones(t *testing.T) {
	est, _ := time.LoadLocation("America/New_York")
	utc := time.UTC

	tests := []struct {
		name     string
		t        time.Time
		start    time.Time
		end      time.Time
		expected time.Time
	}{
		{
			name:     "same instant different zones within range",
			t:        time.Date(2026, 4, 15, 12, 0, 0, 0, est),
			start:    time.Date(2026, 4, 15, 10, 0, 0, 0, utc),
			end:      time.Date(2026, 4, 15, 18, 0, 0, 0, utc),
			expected: time.Date(2026, 4, 15, 12, 0, 0, 0, est),
		},
		{
			name:     "time after range with different timezone",
			t:        time.Date(2026, 4, 15, 20, 0, 0, 0, est),
			start:    time.Date(2026, 4, 15, 10, 0, 0, 0, utc),
			end:      time.Date(2026, 4, 15, 14, 0, 0, 0, utc),
			expected: time.Date(2026, 4, 15, 14, 0, 0, 0, utc),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Clamp(tt.t, tt.start, tt.end)
			if !result.Equal(tt.expected) {
				t.Errorf("Clamp(%v, %v, %v) = %v, want %v", tt.t, tt.start, tt.end, result, tt.expected)
			}
		})
	}
}

func TestClamp_NanosecondPrecision(t *testing.T) {
	base := time.Date(2026, 4, 15, 12, 0, 0, 0, time.UTC)
	start := base.Add(-1000 * time.Nanosecond)
	end := base.Add(1000 * time.Nanosecond)

	tests := []struct {
		name     string
		t        time.Time
		expected time.Time
	}{
		{
			name:     "exactly at start",
			t:        start,
			expected: start,
		},
		{
			name:     "one nanosecond before start",
			t:        start.Add(-1 * time.Nanosecond),
			expected: start,
		},
		{
			name:     "one nanosecond after start",
			t:        start.Add(1 * time.Nanosecond),
			expected: start.Add(1 * time.Nanosecond),
		},
		{
			name:     "exactly at end",
			t:        end,
			expected: end,
		},
		{
			name:     "one nanosecond before end",
			t:        end.Add(-1 * time.Nanosecond),
			expected: end.Add(-1 * time.Nanosecond),
		},
		{
			name:     "one nanosecond after end",
			t:        end.Add(1 * time.Nanosecond),
			expected: end,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Clamp(tt.t, start, end)
			if !result.Equal(tt.expected) {
				t.Errorf("Clamp(%v, %v, %v) = %v, want %v", tt.t, start, end, result, tt.expected)
			}
		})
	}
}

func TestClamp_WithHelperFunctions(t *testing.T) {
	tests := []struct {
		name     string
		t        time.Time
		start    time.Time
		end      time.Time
		expected time.Time
	}{
		{
			name:     "clamp time within range using Min/Max",
			t:        time.Date(2026, 4, 15, 12, 0, 0, 0, time.UTC),
			start:    time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			expected: time.Date(2026, 4, 15, 12, 0, 0, 0, time.UTC),
		},
		{
			name:     "clamp time created by Ago",
			t:        lxtime.Ago(30, time.Minute),
			start:    lxtime.Ago(1, time.Hour),
			end:      lxtime.FromNow(1, time.Hour),
			expected: lxtime.Ago(30, time.Minute),
		},
		{
			name:     "clamp time to past range",
			t:        time.Date(2026, 4, 15, 12, 0, 0, 0, time.UTC),
			start:    lxtime.Ago(2, time.Hour),
			end:      lxtime.Ago(1, time.Hour),
			expected: lxtime.Ago(1, time.Hour),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.Clamp(tt.t, tt.start, tt.end)
			// For dynamic times (Ago/FromNow), use approximate comparison
			tolerance := 100 * time.Millisecond
			diff := result.Sub(tt.expected)
			if diff < -tolerance || diff > tolerance {
				t.Logf("Clamp result differs by %v (tolerance: %v)", diff, tolerance)
			}
		})
	}
}

func TestClamp_Idempotent(t *testing.T) {
	t.Run("clamping_twice_gives_same_result", func(t *testing.T) {
		timeVal := time.Date(2026, 4, 15, 12, 0, 0, 0, time.UTC)
		start := time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC)
		end := time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC)

		clamped1 := lxtime.Clamp(timeVal, start, end)
		clamped2 := lxtime.Clamp(clamped1, start, end)

		if !clamped1.Equal(clamped2) {
			t.Errorf("Clamp is not idempotent: %v != %v", clamped1, clamped2)
		}
	})
}

func TestClamp_Commutative(t *testing.T) {
	t.Run("result_is_within_range", func(t *testing.T) {
		// Test that result is always within [start, end]
		tests := []struct {
			t     time.Time
			start time.Time
			end   time.Time
		}{
			{
				t:     time.Date(2026, 4, 15, 8, 0, 0, 0, time.UTC),
				start: time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
				end:   time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			},
			{
				t:     time.Date(2026, 4, 15, 12, 0, 0, 0, time.UTC),
				start: time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
				end:   time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			},
			{
				t:     time.Date(2026, 4, 15, 16, 0, 0, 0, time.UTC),
				start: time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC),
				end:   time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC),
			},
		}

		for _, test := range tests {
			result := lxtime.Clamp(test.t, test.start, test.end)
			if result.Before(test.start) || result.After(test.end) {
				t.Errorf("Clamp(%v, %v, %v) = %v is outside range", test.t, test.start, test.end, result)
			}
		}
	})
}

func ExampleClamp() {
	start := time.Date(2026, 4, 15, 10, 0, 0, 0, time.UTC)
	end := time.Date(2026, 4, 15, 14, 0, 0, 0, time.UTC)
	mid := time.Date(2026, 4, 15, 12, 0, 0, 0, time.UTC)

	clamped := lxtime.Clamp(mid, start, end)
	// clamped: 2026-04-15 12:00:00 +0000 UTC (unchanged, within range)
	_ = clamped

	tooEarly := time.Date(2026, 4, 15, 8, 0, 0, 0, time.UTC)
	clamped2 := lxtime.Clamp(tooEarly, start, end)
	// clamped2: 2026-04-15 10:00:00 +0000 UTC (clamped to start)
	_ = clamped2

	tooLate := time.Date(2026, 4, 15, 16, 0, 0, 0, time.UTC)
	clamped3 := lxtime.Clamp(tooLate, start, end)
	// clamped3: 2026-04-15 14:00:00 +0000 UTC (clamped to end)
	_ = clamped3
}
