package polynomial

// Factorial calculates the factorial of the input integer.
// Doesn't handle negative numbers gracefully, up to user to not pass them.
// Handles up to n=20, beyond that it will overflow.
func Factorial(n int) (z int) {
	if n>1 {
		return n*Factorial(n-1)
	}
	return 1
}

// FactorialRatio calculates the ratio of the factorial of the input integers.
// Useful when dividing a large factorial by a smaller factorial, to fit
// inside an int64.
// Doesn't handle negative or large numbers gracefully, up to user to not pass them.
func FactorialRatio(n, m int) (z int) {
	if n>m {
		return n*FactorialRatio(n-1, m)
	}
	return 1
}

// FactorialRatioFloat calculates the ratio of the factorial of the input integers
// and returns it as a float, to handle large numbers.
// Doesn't handle negative or large nu