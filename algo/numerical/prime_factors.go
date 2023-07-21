package numerical

// PrimeFactors returns the prime factors of n > 0 as a map of prime numbers and their powers.
func PrimeFactors(n int) map[int]int {
	if n <= 0 {
		return nil
	}

	factors := make(map[int]int)
	if n == 1 {
		return factors
	}

	count := 0
	for n%2 == 0 {
		n /= 2
		count++
	}
	if count > 0 {
		factors[2] = count
	}

	for i := 3; i <= n; i += 2 {
		count = 0
		for n%i == 0 {
			n /= i
			count++
		}
		if count > 0 {
			factors[i] = count
		}
	}

	return factors
}
