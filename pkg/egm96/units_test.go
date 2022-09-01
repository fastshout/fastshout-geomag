package egm96

import (
	"fmt"
	"testing"
)

func TestDegrees(t *testing.T) {
	ds := []float64{59, 30, 20, -12, -89, 0, 0, 0, 0}
	ms := []float64{59, 12, 18, 45, 59, 6, -12, 0, 0}
	ss := []float64{5