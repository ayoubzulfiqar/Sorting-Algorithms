package sort

// Shell Knuth's Sequence
func Knuth[T Numeric](array []T) []T {
	var n int = len(array)

	// Initialize the gap to 1 (start of Knuth's sequence)
	var gap int = 1

	// Generate Knuth's sequence of increments
	for gap < n/3 {
		gap = 3*gap + 1
	}

	// Start with the largest gap and reduce it following Knuth's sequence
	for gap > 0 {
		for i := gap; i < n; i++ {
			// Store the current element for comparison
			var currentElement T = array[i]
			var j int = i

			// Inner loop: Compare and swap elements with a gap
			for j >= gap && array[j-gap] > currentElement {
				// Shift elements to the right by the gap
				array[j] = array[j-gap]
				j -= gap
			}

			// Place the current element in its correct position
			array[j] = currentElement
		}

		// Reduce the gap following Knuth's sequence
		gap /= 3
	}
	return array
}
