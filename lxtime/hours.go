package lxtime

import "time"

// Hours returns a duration representing n hours.
//
// Example:
//
//	duration := lxtime.Hours(3)
//	// duration: 3h0m0s
func Hours(n int) time.Duration {
	return time.Duration(n) * time.Hour
}
