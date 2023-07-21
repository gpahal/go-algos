package numerical

// Fibonacci return the nth fibonacci number. NOTE: Fibonacci(0) = 0 and Fibonacci(1) = 1.
func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	// pprev stores the prev to prev fibonacci value. Initially pprev = Fibonacci(0).
	pprev := 0
	// prev stores the previous fibonacci value. Initially prev = Fibonacci(1).
	prev := 1

	// Apply the Fibonacci formula (ie. F(k-1), F(k) = F(k), F(k-1) + F(k)) n-1 times on pprev and
	// prev to get Fibonacci(n).
	for i := n - 1; i > 0; i-- {
		pprev, prev = prev, pprev+prev
	}

	return prev
}
