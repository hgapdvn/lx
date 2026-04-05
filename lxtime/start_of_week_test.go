package lxtime_test

import (
	"testing"
	"time"

	lxtime "github.com/hgapdvn/lx/lxtime"
)

func TestStartOfWeek_BasicCases(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "Monday at start",
			check: func() bool {
				input := time.Date(2026, 4, 6, 15, 30, 0, 0, time.UTC)
				result := lxtime.StartOfWeek(input)
				return result.Weekday() == time.Monday && result.Hour() == 0
			},
		},
		{
			name: "Wednesday goes to Monday",
			check: func() bool {
				input := time.Date(2026, 4, 8, 15, 30, 0, 0, time.UTC)
				result := lxtime.StartOfWeek(input)
				expected := time.Date(2026, 4, 6, 0, 0, 0, 0, time.UTC)
				return result.Equal(expected)
			},
		},
		{
			name: "Sunday goes to previous Monday",
			check: func() bool {
				input := time.Date(2026, 4, 5, 15, 30, 0, 0, time.UTC)
				result := lxtime.StartOfWeek(input)
				return result.Weekday() == time.Monday
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("StartOfWeek() check failed")
			}
		})
	}
}

func TestStartOfWeek_TimezoneCases(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "EST preserves timezone",
			check: func() bool {
				est, _ := time.LoadLocation("America/New_York")
				// 2026-04-08 (Wednesday) in EST
				input := time.Date(2026, 4, 8, 10, 30, 0, 0, est)
				result := lxtime.StartOfWeek(input)
				// Should be Monday 2026-04-06 in EST
				expected := time.Date(2026, 4, 6, 0, 0, 0, 0, est)
				return result.Equal(expected)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("StartOfWeek() timezone check failed")
			}
		})
	}
}
