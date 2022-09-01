package egm96

import (
	"fmt"
	"testing"
)

func TestDegrees(t *testing.T) {
	ds := []float64{59, 30, 20, -12, -89, 0, 0, 0, 0}
	ms := []float64{59, 12, 18, 45, 59, 6, -12, 0, 0}
	ss := []float64{59.999, 46, 31, 12, 1.25, 0, 54, -45, 18}
	dds := []float64{59.999999722, 30.2127777