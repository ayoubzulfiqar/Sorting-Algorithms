package sort

// Shell Pratt's sequence
func Pratt[T Numeric](array []T) []T {
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
