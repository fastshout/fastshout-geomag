package polynomial

// Factorial calculates the factorial of the input integer.
// Doesn't handle negative numbers gracefully, up to user to not pass them.
// Handles up to n=20, beyond that it will overflow.
func Factorial(n int) (z int) {
	if n>1 {