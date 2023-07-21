package numerical

// LCM returns the least common multiple of a and b.
func LCM(a, b int) int {
	// Make a and b non-negative.
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}

	// LCM = (a * b) / gcd(a, b)
	return int((a * b) / GCD(a, b))
}

// LCMArray returns the least common multiple of all the elements in arr. If arr is nil or empty,
// 0 is returned.
func LCMArray(arr []int) int {
	arrLen := len(arr)
	if arrLen == 0 {
		return 0
	}
	if arrLen == 1 {
		return arr[0]
	}

	// ans stores the lcm of elements arr[:i].
	ans := arr[0]

	// Loop from i = 1 to arrLength - 1 and continue updating ans by finding lcm of ans and arr[i].
	for i := 1; i < arrLen; i++ {
		// If any element is 0, lcm is also 0.
		if ans == 0 {
			return 0
		}

		ans = LCM(ans, arr[i])
	}

	return ans
}
