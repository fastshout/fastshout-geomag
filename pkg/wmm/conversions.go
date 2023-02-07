package wmm

import (
	"time"
)

type DecimalYear float64

// DecimalYearsToTime converts an epoch-like float64 year like 2015.0 to a Go time.Time.
// Per document MIL-PRF-89500B Section 3.2, "Time is referenced in decimal years
// (e.g., 15 May 2019 is 2019.367). Note that the day-of-year (DOY) of January