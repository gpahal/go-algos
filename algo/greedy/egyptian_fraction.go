package greedy

// EgyptianFraction takes a positive fraction (num = numerator and den = denominator) less than 1
// and returns it's representation as a sum of unique unit fractions. A fraction is a unit fraction
// if it's numerator is 1 and denominator is a positive integer. This function returns a slice
// representing the denominators of such a representation in increasing order.
//
// # Example
//
//     ef := greedy.EgyptianFraction(2, 3)
//     // ef = []int{2, 6} as 2/3 = 1/2 + 1/6
//
// If num < 0 or den <= 0 or num >= den, nil is returned.
func EgyptianFraction(num, den int) []int {
	if num < 0 || den <= 0 || num >= den {
		return nil
	}
	if num == 0 {
		return []int{}
	}

	seq := make([]int, 0)
	for {
		if den%num == 0 {
			seq = append(seq, den/num)
			break
		}

		newDen := (den / num) + 1
		seq = append(seq, newDen)
		num, den = num*newDen-den, den*newDen
	}

	return seq
}
