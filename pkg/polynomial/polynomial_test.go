package polynomial

import (
	"fmt"
	"testing"
)

const eps = 1e-6

func testDiff(name string, actual, expected float64, eps float64, t *testing.T) {
	if actual - expected > -eps && actual - expected < eps {
		t.Logf("%s correct: expected %8.4f, got %8.4f", name, expected, actual)
		return
	}
	t.Errorf("%s incorrect: expected %8.4f, got %8.4f", name, expected, actual)
}

func TestPow(t *testing.T) {
	var (
		xs = []float64{2.0, 0.5, 1.0, 3.14, 10}
		ns = []int{5, 3, 4, 0, -3}
		ys = []float64{32, 0.125, 1, 1, 0.001}
	)

	for i:=0; i<len(xs); i++ {
		y := Pow(xs[i], ns[i])
		testDiff("Pow", y, ys[i], eps, t)
	}
}

func TestFa