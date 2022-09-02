package polynomial

import "math"

type legendreFunctionIndex struct {
	n, m int
}
var legendreFunctionCache = make(map[legendreFunctionIndex]Polynomial)

// LegendrePolynomial returns a Polynomial object correspon