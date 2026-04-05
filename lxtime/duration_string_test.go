package lxtime_test

import (
	"testing"
	"time"

	lxtime "github.com/hgapdvn/lx/lxtime"
)

func TestDurationString_BasicCases(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		expected string
	}{
		{
			name:     "zero duration",
			duration: 0,
			expected: "0 seconds",
		},
		{
			name:     "single second",
			duration: time.Second,
			expected: "1 second",
		},
		{
			name:     "multiple seconds",
			duration: 5 * time.Second,
			expected: "5 seconds",
		},
		{
			name:     "single minute",
			duration: time.Minute,
			expected: "1 minute",
		},
		{
			name:     "multiple minutes",
			duration: 10 * time.Minute,
			expected: "10 minutes",
		},
		{
			name:     "single hour",
			duration: time.Hour,
			expected: "1 hour",
		},
		{
			name:     "multiple hours",
			duration: 3 * time.Hour,
			expected: "3 hours",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.DurationString(tt.duration)
			if result != tt.expected {
				t.Errorf("DurationString(%v) = %q, want %q", tt.duration, result, tt.expected)
			}
		})
	}
}

func TestDurationString_Combined(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		expected string
	}{
		{
			name:     "hours and minutes",
			duration: 2*time.Hour + 3*time.Minute,
			expected: "2 hours, 3 minutes",
		},
		{
			name:     "hours and seconds",
			duration: 1*time.Hour + 30*time.Second,
			expected: "1 hour, 30 seconds",
		},
		{
			name:     "minutes and seconds",
			duration: 5*time.Minute + 45*time.Second,
			expected: "5 minutes, 45 seconds",
		},
		{
			name:     "all three: hours, minutes, seconds",
			duration: 2*time.Hour + 15*time.Minute + 30*time.Second,
			expected: "2 hours, 15 minutes, 30 seconds",
		},
		{
			name:     "hours, minutes, seconds with singular forms",
			duration: 1*time.Hour + 1*time.Minute + 1*time.Second,
			expected: "1 hour, 1 minute, 1 second",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.DurationString(tt.duration)
			if result != tt.expected {
				t.Errorf("DurationString(%v) = %q, want %q", tt.duration, result, tt.expected)
			}
		})
	}
}

func TestDurationString_Submillisecond(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		expected string
	}{
		{
			name:     "millisecond",
			duration: time.Millisecond,
			expected: "1 millisecond",
		},
		{
			name:     "multiple milliseconds",
			duration: 500 * time.Millisecond,
			expected: "500 milliseconds",
		},
		{
			name:     "microsecond",
			duration: time.Microsecond,
			expected: "1 microsecond",
		},
		{
			name:     "multiple microseconds",
			duration: 100 * time.Microsecond,
			expected: "100 microseconds",
		},
		{
			name:     "nanosecond",
			duration: time.Nanosecond,
			expected: "1 nanosecond",
		},
		{
			name:     "multiple nanoseconds",
			duration: 100 * time.Nanosecond,
			expected: "100 nanoseconds",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.DurationString(tt.duration)
			if result != tt.expected {
				t.Errorf("DurationString(%v) = %q, want %q", tt.duration, result, tt.expected)
			}
		})
	}
}

func TestDurationString_Complex(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		expected string
	}{
		{
			name:     "seconds and milliseconds",
			duration: 5*time.Second + 500*time.Millisecond,
			expected: "5 seconds, 500 milliseconds",
		},
		{
			name:     "minutes, seconds, milliseconds",
			duration: 2*time.Minute + 30*time.Second + 250*time.Millisecond,
			expected: "2 minutes, 30 seconds, 250 milliseconds",
		},
		{
			name:     "hours, minutes, seconds, milliseconds",
			duration: 1*time.Hour + 5*time.Minute + 30*time.Second + 100*time.Millisecond,
			expected: "1 hour, 5 minutes, 30 seconds, 100 milliseconds",
		},
		{
			name:     "skip zero components",
			duration: 1*time.Hour + 0*time.Minute + 5*time.Second,
			expected: "1 hour, 5 seconds",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.DurationString(tt.duration)
			if result != tt.expected {
				t.Errorf("DurationString(%v) = %q, want %q", tt.duration, result, tt.expected)
			}
		})
	}
}

func TestDurationString_Negative(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		expected string
	}{
		{
			name:     "negative seconds",
			duration: -5 * time.Second,
			expected: "-5 seconds",
		},
		{
			name:     "negative hours and minutes",
			duration: -(2*time.Hour + 30*time.Minute),
			expected: "-2 hours, 30 minutes",
		},
		{
			name:     "negative with multiple components",
			duration: -(1*time.Hour + 15*time.Minute + 45*time.Second),
			expected: "-1 hour, 15 minutes, 45 seconds",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.DurationString(tt.duration)
			if result != tt.expected {
				t.Errorf("DurationString(%v) = %q, want %q", tt.duration, result, tt.expected)
			}
		})
	}
}

func TestDurationString_PluralForms(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "singular hour",
			check: func() bool {
				return lxtime.DurationString(time.Hour) == "1 hour"
			},
		},
		{
			name: "plural hours",
			check: func() bool {
				return lxtime.DurationString(2*time.Hour) == "2 hours"
			},
		},
		{
			name: "singular minute",
			check: func() bool {
				return lxtime.DurationString(time.Minute) == "1 minute"
			},
		},
		{
			name: "plural minutes",
			check: func() bool {
				return lxtime.DurationString(5*time.Minute) == "5 minutes"
			},
		},
		{
			name: "singular second",
			check: func() bool {
				return lxtime.DurationString(time.Second) == "1 second"
			},
		},
		{
			name: "plural seconds",
			check: func() bool {
				return lxtime.DurationString(10*time.Second) == "10 seconds"
			},
		},
		{
			name: "singular millisecond",
			check: func() bool {
				return lxtime.DurationString(time.Millisecond) == "1 millisecond"
			},
		},
		{
			name: "plural milliseconds",
			check: func() bool {
				return lxtime.DurationString(100*time.Millisecond) == "100 milliseconds"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("DurationString() plural form check failed")
			}
		})
	}
}

func TestDurationString_LargeValues(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		expected string
	}{
		{
			name:     "24 hours",
			duration: 24 * time.Hour,
			expected: "24 hours",
		},
		{
			name:     "multiple hours with various units",
			duration: 10*time.Hour + 59*time.Minute + 59*time.Second + 999*time.Millisecond,
			expected: "10 hours, 59 minutes, 59 seconds, 999 milliseconds",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.DurationString(tt.duration)
			if result != tt.expected {
				t.Errorf("DurationString(%v) = %q, want %q", tt.duration, result, tt.expected)
			}
		})
	}
}

func TestDurationString_SmallValues(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		expected string
	}{
		{
			name:     "single nanosecond",
			duration: time.Nanosecond,
			expected: "1 nanosecond",
		},
		{
			name:     "999 nanoseconds",
			duration: 999 * time.Nanosecond,
			expected: "999 nanoseconds",
		},
		{
			name:     "microseconds and nanoseconds",
			duration: 50*time.Microsecond + 500*time.Nanosecond,
			expected: "50 microseconds, 500 nanoseconds",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxtime.DurationString(tt.duration)
			if result != tt.expected {
				t.Errorf("DurationString(%v) = %q, want %q", tt.duration, result, tt.expected)
			}
		})
	}
}

func TestDurationString_Consistency(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "same duration returns same string",
			check: func() bool {
				d := 2*time.Hour + 30*time.Minute + 45*time.Second
				result1 := lxtime.DurationString(d)
				result2 := lxtime.DurationString(d)
				return result1 == result2
			},
		},
		{
			name: "does not modify input",
			check: func() bool {
				d := 5 * time.Second
				original := d
				_ = lxtime.DurationString(d)
				return d == original
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("DurationString() consistency check failed")
			}
		})
	}
}
