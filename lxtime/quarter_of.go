package lxtime

import "time"

// QuarterOf returns the quarter of the year for the given time.
// Returns 1 for Q1 (Jan-Mar), 2 for Q2 (Apr-Jun), 3 for Q3 (Jul-Sep), 4 for Q4 (Oct-Dec).
//
// Example:
//
//	t := time.Date(2026, 4, 15, 0, 0, 0, 0, time.UTC)
//	quarter := lxtime.QuarterOf(t)
//	// quarter: 2
//
//	t2 := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
//	quarter2 := lxtime.QuarterOf(t2)
//	// quarter2: 1
func QuarterOf(t time.Time) int {
	month := t.Month()
	// Quarter 1: Jan (1), Feb (2), Mar (3)
	// Quarter 2: Apr (4), May (5), Jun (6)
	// Quarter 3: Jul (7), Aug (8), Sep (9)
	// Quarter 4: Oct (10), Nov (11), Dec (12)
	return (int(month)-1)/3 + 1
}
