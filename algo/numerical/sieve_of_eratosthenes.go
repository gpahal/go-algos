package numerical

// SieveOfEratosthenes finds all prime numbers up to any given limit. It returns a bool slice of
// length `limit` and its nth index is true if n is NOT a prime and false if it is. 0 and 1 are not
// primes.
//
// NOTE: limit should be > 0 and <= 2 ^ 24 (around 16.7 million), otherwise nil is returned.
func SieveOfEratosthenes(limit int) []bool {
	// NOTE: 1 << n = 2 ^ n (as long as 2 ^ n is within the limits of what an int value can hold)
	if limit <= 0 || limit > (1<<24) {
		return nil
	}

	// sieve is the slice that stores primality values. By default, all numbers are prime ie.
	// false.
	sieve := make([]bool, limit+1)

	// Marks 0 and 1 as NOT prime.
	sieve[0] = true
	sieve[1] = true

	// Mark multiples of 2 as composite or not prime.
	for m := 4; m <= limit; m += 2 {
		sieve[m] = true
	}

	// Loop over all odd numbers starting from 3. 2 is being treated as a special case as it is the
	// only prime that is even. This allows us to skip all even numbers in the loop.
	for p := 3; p <= limit; p += 2 {
		if !sieve[p] {
			// Mark numbers p(p), p(p+2), p(p+4), ... as composite or not prime. NOTE: p+1, p+3 are
			// even and hence p(p+1), p(p+3), ... have already been marked as they are multiples of
			// 2.
			for m := p * p; m <= limit; m += 2 * p {
				sieve[m] = true
			}
		}
	}

	return sieve
}

// SieveOfEratosthenesDynamic returns a generator function which when called the nth time returns
// the nth prime number. Unlike the normal SieveOfEratosthenes function, this function does not
// allocate a huge slice at initialization. That is why it is has the word dynamic in its name.
func SieveOfEratosthenesDynamic() func() int {
	// composites is a temporary map of composites (only odd numbers) to their factors.
	composites := make(map[int][]int)

	// Start with the number 2.
	p := 2

	return func() int {
		// 2 is being treated as a special case. Simply, set p to 3 and return 2.
		if p == 2 {
			p = 3
			return 2
		}

		// returnVal stores the current value of p as it needs to be returned at the end.
		returnVal := p

		// Append p to composites[p*p] as p is a factor of p*p.
		pSquared := p * p
		factors, ok := composites[pSquared]
		if ok {
			composites[pSquared] = append(factors, p)
		} else {
			composites[pSquared] = []int{p}
		}

		var factorMultiple int
		var tmpFactors []int

		// Increment by 2 to analyze the next potential prime number. As p >= 3 here, even numbers
		// cannot be prime, thus, we are incrementing by 2.
		p += 2
		for {
			// If current p is not composite => p is prime. This will be returned in the next
			// invocation of the generator function.
			factors, ok = composites[p]
			if !ok {
				break
			}

			// Loop over all the factors of composite p.
			for _, factor := range factors {
				// Append factor as a factor of p + 2*factor. NOTE: p + factor, p + 3*factor, ...
				// are even and are not analyzed at all. NOTE: When we analyze p + 2*factor, we
				// will ultimately also analyze p + 4*factor and hence only one number ie.
				// p + 2*factor is added as a composite in this step.
				factorMultiple = p + 2*factor
				tmpFactors, ok = composites[factorMultiple]
				if ok {
					composites[factorMultiple] = append(tmpFactors, factor)
				} else {
					composites[factorMultiple] = []int{factor}
				}
			}

			// Delete p from composites as it is no longer useful and deletion will save space.
			delete(composites, p)
			p += 2
		}

		return returnVal
	}
}
