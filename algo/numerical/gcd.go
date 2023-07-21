package numerical

// GCD returns the greatest common divisor of a and b using the basic Euclidean algorithm.
func GCD(a, b int) int {
	// Make a and b non-negative.
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}

	// Make a the larger element. Saves one step in the loop.
	if a < b {
		a, b = b, a
	}

	// While the remainder b is not 0, continue updating a and b according to the basic Euclidean
	// algorithm.
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// GCDArray returns the greatest common divisor of all the elements in arr using the basic
// Euclidean algorithm. If arr is nil or empty, 0 is returned.
func GCDArray(arr []int) int {
	arrLen := len(arr)
	if arrLen == 0 {
		return 0
	}
	if arrLen == 1 {
		return arr[0]
	}

	// ans stores the gcd of elements arr[:i].
	ans := arr[0]

	// Loop from i = 1 to arrLength - 1 and continue updating ans by applying Euclidean algorithm
	// on ans and arr[i].
	for i := 1; i < arrLen; i++ {
		// If any element is 0, gcd is also 0.
		if ans == 0 {
			return 0
		}

		ans = GCD(ans, arr[i])
	}

	return ans
}
