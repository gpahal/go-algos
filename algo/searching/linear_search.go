package searching

// LinearSearch linearly searches the slice for the first occurrence of key and returns the index.
// If the key is not found, -1 is returned.
func LinearSearch(arr []int, key int) int {
	// Loop over all elements and return the index if the key is found.
	for i, el := range arr {
		if el == key {
			return i
		}
	}

	return -1
}
