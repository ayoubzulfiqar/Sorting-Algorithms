package sort

/*

	-* Tim Sort *-


1. `Tim` function:
   - The `Tim` function is the main sorting function that takes an unsorted array and
     returns the sorted array using TimSort.

   - It first calculates the minimum run size based on the array size using the
    `findMinRun` function.

   - It sorts individual sub-arrays of size `minRun` using the `Insertion` function.
     The insertion sort is used for small sub-arrays as it's efficient for such cases.

   - After sorting individual sub-arrays, the function proceeds to merge them.

   - It initializes a variable `size` to the value of `minRun` and enters a loop to
     merge the sorted sub-arrays. It chooses merge sizes of `size`, `2*size`, `4*size`,
	 and so on until `size` is greater than or equal to the array size.

   - Within the merging loop, it iterates over the sub-arrays, calling the `merging`
     function to merge two sorted sub-arrays into one.

   - The `size` is doubled in each iteration until the entire array is sorted.

   - Finally, the sorted `array` is returned.

2. `findMinRun` function:
   - The `findMinRun` function calculates the minimum run size for TimSort.

   - It uses a constant `minRun` value (set to 32) as a starting point.

   - It iteratively checks if the input `n` is greater than or equal to `minRun`.
     If it is, it sets the least significant bit of `n` to 1 by using a bitwise OR operation.
	 It then right-shifts `n` by one bit.

   - The function returns `n + r`, which ensures that the minimum run size is greater
     than or equal to `minRun`.

3. `min` function:
   - The `min` function is a simple helper function that returns the minimum of two values,
     `a` and `b`.

4. `merging` function:
   - The `merging` function is responsible for merging two sorted sub-arrays, `left` and `right`,
     into a single sorted array.

   - It initializes a `result` slice with a capacity equal to the combined lengths of
     `left` and `right`.

   - It maintains three indices, `leftIndex`, `rightIndex`, and `resultIndex`, to traverse
     the sub-arrays and build the merged result.

   - The function iteratively compares elements from `left` and `right` and appends the
     smaller element to the `result`. The corresponding indices are incremented accordingly.

   - After one of the sub-arrays is exhausted, the remaining elements from the other
     sub-array are directly copied into the `result`.

   - Finally, the `result` slice is copied back into `left` and `right`, and it is returned
     as the merged result.


*/

// Generic Implementation is Slow
func Tim[T Numeric](array []T) []T {
	var n int = len(array)
	var minRun int = findMinRun(n)

	// Sort individual sub-arrays of size minRun
	for start := 0; start < n; start += minRun {
		end := start + minRun
		if end > n {
			end = n
		}
		// We have already Implemented Insertion Sort Long Time ago
		Insertion(array[start:end])
	}

	// Now merge the sorted runs
	var size int = minRun
	for size < n {
		// Choose merges sized sz, sz*2, sz*4, until sz >= n
		for start := 0; start < n; start += 2 * size {
			mid := start + size
			end := min(start+2*size, n)
			merging(array[start:mid], array[mid:end])
		}
		size *= 2
	}
	return array
}

func findMinRun(n int) int {
	const minRun int = 32
	r := 0
	for n >= minRun {
		r |= n & 1
		n >>= 1
	}
	return n + r
}

func min[T Numeric](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

// func insertionSort(array []int) {
// 	var j int

// 	for i := range array {
// 		j = i
// 		for j > 0 && array[j-1] > array[j] {
// 			array[j-1], array[j] = array[j], array[j-1]
// 			j--
// 		}
// 	}
// }

func merging[T Numeric](left []T, right []T) []T {
	// Merge two sorted sub-arrays into one
	var result []T = make([]T, len(left)+len(right))
	leftIndex, rightIndex, resultIndex := 0, 0, 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] <= right[rightIndex] {
			result[resultIndex] = left[leftIndex]
			leftIndex++
		} else {
			result[resultIndex] = right[rightIndex]
			rightIndex++
		}
		resultIndex++
	}

	// Copy remaining elements from either a or b
	for leftIndex < len(left) {
		result[resultIndex] = left[leftIndex]
		leftIndex++
		resultIndex++
	}
	for rightIndex < len(right) {
		result[resultIndex] = right[rightIndex]
		rightIndex++
		resultIndex++
	}

	// Copy tmp back to a and b
	copy(left, result)
	copy(right, result[len(left):])
	return result
}
