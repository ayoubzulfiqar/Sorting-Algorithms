package sort

/*


-* Shell Sort *-

Sequences:

The interval between the elements is reduced based on the sequence used.
Some of the optimal sequences that can be used in the shell sort algorithm are:

1. Shell's original sequence: N/2 , N/4 , …, 1
2. Knuth's increments: 1, 4, 13, …, (3k - 1) / 2
3. Sedgewick's increments: 1, 8, 23, 77, 281, 1073, 4193, 16577...4j+1+ 3·2j+ 1
4. Hibbard's increments: 1, 3, 7, 15, 31, 63, 127, 255, 511…
5. Papernov & Stasevich increment: 1, 3, 5, 9, 17, 33, 65,...
6. Pratt: 1, 2, 3, 4, 6, 9, 8, 12, 18, 27, 16, 24, 36, 54, 81....


Note: The performance of the shell sort depends on the type of sequence used for a given input array.





*/

// With Original Shell Sort Sequence
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

// Knuth's Sequence
func ShellKnuth[T Numeric](array []T) []T {
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

// Sedgewick's Sequence

func ShellSedgewick[T Numeric](array []T) []T {
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

// Hibbard's Sequence
func ShellHibbard[T Numeric](array []T) []T {
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

func ShellPS[T Numeric](array []T) []T {
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

// Pratt's sequence
func ShellPratt[T Numeric](array []T) []T {
	var n int = len(array)

	// Generate Pratt's sequence of increments
	var gaps []int = []int{1}
	var i int = 0

	for {
		powerTwo := 1 << i
		if powerTwo > n {
			break
		}

		j := 0
		for {
			gap := powerTwo * (3 << j)
			if gap > n {
				break
			}
			gaps = append(gaps, gap)
			j++
		}
		i++
	}

	for _, gap := range gaps {
		for i := gap; i < n; i++ {
			var currentElement T = array[i]
			var j int = i

			for ; j >= gap && array[j-gap] > currentElement; j -= gap {
				array[j] = array[j-gap]
			}

			array[j] = currentElement
		}
	}
	return array
}
