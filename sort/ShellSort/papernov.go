package sort

// Shell Papernov & Stasevich's increment sequence
func PapernovStasevich[T Numeric](array []T) []T {
	var n int = len(array)

	// Initialize 'k' for Papernov & Stasevich's increment sequence
	var k int = 0
	var gap int = (1 << (k + 1)) - 3*(1<<k) + 1

	for gap >= 1 {
		for i := gap; i < n; i++ {
			// Store the current element for comparison
			var currentElement T = array[i]
			var j int = i

			// Inner loop: Compare and swap elements with the gap
			for ; j >= gap && array[j-gap] > currentElement; j -= gap {
				array[j] = array[j-gap]
			}

			// Place the current element in its correct position
			array[j] = currentElement
		}

		// Calculate the next gap using Papernov & Stasevich's increment formula
		k++
		gap = (1 << (k + 1)) - 3*(1<<k) + 1
	}
	return array
}
