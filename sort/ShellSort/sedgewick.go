package sort

// Shell Sedgewick's Sequence
func Sedgewick[T Numeric](array []T) []T {
	var n int = len(array)

	// Generate Sedgewick's sequence of increments
	var gap int = 1
	for gap < n/3 {
		gap = 8*gap + 1
	}

	for gap >= 1 {
		for i := gap; i < n; i++ {
			var currentElement T = array[i]
			var j int = i

			for ; j >= gap && array[j-gap] > currentElement; j -= gap {
				array[j] = array[j-gap]
			}

			array[j] = currentElement
		}

		// Calculate the next gap using Sedgewick's sequence
		gap = (gap - 1) / 8
	}
	return array
}
