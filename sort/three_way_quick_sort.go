package sort

// ThreeWayQuick performs Three-Way Quick Sort on the input array.
func ThreeWayQuick[T Numeric](array []T) []T {
	// If the array is already sorted or contains only one element, no sorting is needed.
	if len(array) < 2 {
		return array
	}

	// Shuffle the array to improve performance, mitigating worst-case scenarios.
	shuffle(array)

	// Sort the array using the Three-Way Quick Sort algorithm.
	array = sorter(array, 0, len(array)-1)
	return array
}

// sorter is a recursive function that sorts the input array in place using Three-Way Quick Sort.
func sorter[T Numeric](array []T, low, high int) []T {
	// Base case: If the high index is less than or equal to the low index, the array is sorted.
	if high <= low {
		return array
	}

	// Perform the partitioning step to separate elements into three partitions.
	lessThan, greaterThan := partition(array, low, high)

	// Recursively sort the partitions to the left and right of the pivot.
	sorter(array, low, lessThan-1)
	sorter(array, greaterThan+1, high)

	return array
}

// partition function partitions the array into three segments and returns the boundaries of equal elements.
func partition[T Numeric](array []T, low, high int) (int, int) {
	// Initialize the lessThan, iter, and greaterThan pointers and choose the pivot element.
	lessThan, iter, greaterThan := low, low, high
	pivot := array[low]

	// Iterate through the array and place elements into their respective partitions.
	for iter <= greaterThan {
		if array[iter] < pivot {
			// Swap elements that are less than the pivot with the lessThan section.
			array[lessThan], array[iter] = array[iter], array[lessThan]
			lessThan++
			iter++
		} else if array[iter] > pivot {
			// Swap elements that are greater than the pivot with the greaterThan section.
			array[iter], array[greaterThan] = array[greaterThan], array[iter]
			greaterThan--
		} else {
			// Skip elements that are equal to the pivot.
			iter++
		}
	}

	// Return the boundaries of elements equal to the pivot.
	return lessThan, greaterThan

}
