package wmm

import (
	"time"
)

type DecimalYear float64

// DecimalYearsToTime converts an epoch-like float64 year like 2015.0 to a Go time.Time.
// Per document MIL-PRF-89500B Section 3.2, "Time is referenced in decimal years
// (e.g., 15 May 2019 is 2019.367). Note that the day-of-year (DOY) of January 1st is zero
// and December 31st is 364 for non-leap year. For a leap year, DOY of December 31st is 365."
func (y DecimalYear) ToTime() (t time.Time) {
	tYear := int(y)
	yearDays := float64(time.Date(tYear, 12, 31, 0, 0, 0, 0, time