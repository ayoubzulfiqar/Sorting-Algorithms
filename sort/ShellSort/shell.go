package sort

// Original Shell Sort Sequence
func Shell[T Numeric](array []T) []T {
	// Calculate the length of the array
	var length int = len(array)

	// Start with a gap of half the array length
	var gap int = length / 2

	// Continue until the gap is greater than 0 OR reach 1
	for gap > 0 {
		// Iterate through the array from gap to the end
		for i := gap; i < length; i++ {
			// Store the current element in a temporary variable
			var currentElement T = array[i]

			// Initialize position to the current position i
			var position int = i

			// Inner loop: Compare and swap elements with a gap
			for position >= gap && array[position-gap] > currentElement {
				// Shift elements to the right by the gap
				array[position] = array[position-gap]

				// Move to the previous position with the same gap
				position -= gap
			}

			// Place the current element in its correct position
			array[position] = currentElement
		}

		// Reduce the gap size by half
		gap /= 2
	}
	return array
}
