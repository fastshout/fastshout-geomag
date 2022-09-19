package polynomial

type Polynomial struct {
	c []float64
}

// NewPolynomial makes a new polynomial object with the specified coefficients.
// e.g. for x^2-1, use NewPolynomial([]float64{-1,0,1}.
func NewPolynomial(c []float64) (p Polynomial) {
	p.c = c
	return p
}

// Coefficients returns the coefficients of the polynomial in a slice.
func (p Polynomial) Coeffici