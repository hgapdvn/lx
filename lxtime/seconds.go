package lxtime

import "time"

// Seconds returns a duration representing n seconds.
//
// Example:
//
//	duration := lxtime.Seconds(45)
//	// duration: 45s
func Seconds(n int) time.Duration {
	return time.Duration(n) * time.Second
}
