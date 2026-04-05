package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/lxtime"
)

func TestStartOfMonth_BasicCases(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "goes to day 1",
			check: func() bool {
				input := time.Date(2026, 4, 15, 15, 30, 0, 0, time.UTC)
				result := lxtime.StartOfMonth(input)
				return result.Day() == 1 && result.Hour() == 0
			},
		},
		{
			name: "preserves month",
			check: func() bool {
				input := time.Date(2026, 4, 15, 15, 30, 0, 0, time.UTC)
				result := lxtime.StartOfMonth(input)
				return result.Month() == time.April
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("StartOfMonth() check failed")
			}
		})
	}
}
