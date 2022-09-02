package polynomial

import "math"

type legendreFunctionIndex struct {
	n, m int
}
var legendreFunctionCache = make(map[legendreFunctionIndex]Polynomial)

// LegendrePolynomial returns a Polynomial object corresponding to
// the Legendre Polynomial of degree n.
// Once calculated initially, the polynomials are cached for faster future