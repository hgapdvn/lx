package lxtime

import "time"

// Minutes returns a duration representing n minutes.
//
// Example:
//
//	duration := lxtime.Minutes(30)
//	// duration: 30m0s
func Minutes(n int) time.Duration {
	return time.Duration(n) * time.Minute
}
