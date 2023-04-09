package wmm

import (
	"fmt"
	"testing"
	"time"
)

func TestDecimalYearsToTime(t *testing.T) {
	ys := []DecimalYear{1995.0, 1996-1.0/365, 1997-1.0/366, 2004.0, 2017.5}
	ts := []time.Time{
		time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1995, 12, 31, 0, 0, 0, 0, time.UTC),
		time.Date(1996, 12, 31, 0, 0, 0, 0, time.UTC),
		time.Date(2004, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2017, 7, 2, 12, 0, 0, 0, time.UTC),
	}
	for i, y := range ys {
		d := y.ToTime()
		testDiff(fmt.Sprin