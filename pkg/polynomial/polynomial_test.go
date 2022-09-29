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

func TestFactorial(t *testing.T) {
	var (
		ns = []int{20, 19, 5, 3, 4, 0, 1}
		zs = []int{2432902008176640000, 121645100408832000, 120, 6, 24, 1, 1}
	)

	for i:=0; i<len(ns); i++ {
		z := Factorial(ns[i])
		testDiff(fmt.Sprintf("%d!", ns[i]), float64(z), float64(zs[i]), eps, t)
	}
}

// FactorialRatioFloat needs to calculate up to 24!
func TestFactorialRatioFloat(t *testing.T) {
	var (
		ns = []int{6, 6, 6, 6, 3, 3, 1, 24}
		ms = []int{2, 3, 1, 0, 3, 2, 1, 0}
		zs = []float64{360, 120, 720, 720, 1, 3, 1, 620448401733239439360000}
	)

	for i:=0; i<len(ns); i++ {
		z := FactorialRatioFloat(ns[i], ms[i])
		testDiff(fmt.Sprintf("%d!/%d!", ns[i], ms[i]), z, zs[i], eps, t)
	}
}

func TestEvaluate(t *testing.T) {
	var (
		cs = [][]float64{
			{-1, 0, 2},
			{0.5, -1, 1, 2},
		}
		xs = []float64{2, 0.5, -1.5}
		ys = [][]float64{
			{7, 18.5},
			{-0.5, 0.5},
			{3.5, -2.5},
		}
	)

	for i:=0; i<len(cs); i++ {
		for j:=0; j<len(xs); j++ {
			p := NewPolynomial(cs[i])
			y := p.Evaluate(xs[j])
			testDiff(fmt.Sp