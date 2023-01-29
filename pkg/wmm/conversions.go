package wmm

import (
	"time"
)

type DecimalYear float64

// DecimalYearsToTime converts an epoch-like float64 year like 2015.0 to a Go time.Time.
// Per document MIL-PRF-89500B Section 3.2, "Time is 