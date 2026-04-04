package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/time"
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
