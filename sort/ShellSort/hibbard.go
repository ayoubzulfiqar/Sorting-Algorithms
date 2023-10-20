package sort

// Shell Hibbard's Sequence
func Hibbard[T Numeric](array []T) []T {
	var n int = len(array)

	// Initialize the increment 'k' to 1 (start of Hibbard's sequence)
	var k int = 1

	// Calculate the initial gap using Hibbard's sequence
	var gap int = (1 << k) - 1

	// Continue until the gap is less than the array length
	for gap < n {
		// Iterate through the array with the current gap
		for i := gap; i < n; i++ {
			// Store the current element for comparison
			var currentElement T = array[i]
			var j int = i

			// Inner loop: Compare and swap elements with the gap
			for j >= gap && array[j-gap] > currentElement {
				// Shift elements to the right by the gap
				array[j] = array[j-gap]
				j -= gap
			}

			// Place the current element in its correct position
			array[j] = currentElement
		}

		// Calculate the next gap following Hibbard's sequence
		k++
		gap = (1 << k) - 1
	}
	return array
}
