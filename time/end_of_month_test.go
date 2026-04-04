package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/time"
)

func TestEndOfMonth_BasicCases(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "April goes to day 30",
			check: func() bool {
				input := time.Date(2026, 4, 15, 15, 30, 0, 0, time.UTC)
				result := lxtime.EndOfMonth(input)
				return result.Day() == 30 && result.Month() == time.April
			},
		},
		{
			name: "ends at 23:59:59",
			check: func() bool {
				input := time.Date(2026, 4, 15, 15, 30, 0, 0, time.UTC)
				result := lxtime.EndOfMonth(input)
				return result.Hour() == 23 && result.Minute() == 59
			},
		},
		{
			name: "February leap year",
			check: func() bool {
				input := time.Date(2024, 2, 15, 15, 30, 0, 0, time.UTC)
				result := lxtime.EndOfMonth(input)
				return result.Day() == 29
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("EndOfMonth() check failed")
			}
		})
	}
}
