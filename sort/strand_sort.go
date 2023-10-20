package sort

/*


	-* Strand Sort *-

1. `Strand` function:
   - The `Strand` function is the main sorting function that takes an unsorted array as input
     and returns the sorted array.

   - It first checks if the input slice has one or fewer elements. If it does, the slice
     is already sorted, and it's returned as is.

   - It initializes a `result` slice with the first element of the input array and removes
     this element from the input slice.

   - The function enters a loop to continue processing the input slice. Inside the loop:
     - It initializes a `subList` slice with the first element of the input and removes this
	   element from the input slice.
     - It iterates through the input slice to find elements that can be added to the `subList`.
	   Elements are added to the `subList` if they are greater than or equal to the last
	   element of the `subList`. The added elements are also removed from the input slice.

   - After constructing the `subList`, it's merged with the `result` using the `sMerge`
     helper function.

   - The loop continues until the input slice is empty.

   - The sorted `result` slice is returned as the final output.

2. `sMerge` function:
   - The `sMerge` function is a helper function used to merge two sorted slices,
     `a` and `b,` into a single sorted slice.

   - It initializes a `result` slice to hold the merged values. The capacity of this
     slice is set to accommodate both input slices.

   - It maintains two indices, `i` and `j`, to traverse the `a` and `b` slices, respectively.

   - The function merges the two slices while maintaining the sorted order.
     It compares elements at `a[i]` and `b[j]` and appends the smaller of the two to the `result` slice. The corresponding index is incremented.

   - After merging, any remaining elements from `a` and `b` are appended to the `result`.

   - The merged `result` slice is returned.




*/
// Strand is a sorting algorithm that takes an unsorted integer slice and returns a sorted slice.
func Strand[T Numeric](array []T) []T {
	// If the input slice has 1 or fewer elements, it's already sorted, so return it as is.
	if len(array) <= 1 {
		return array
	}

	// Initialize the result slice with the first element of the input.
	var result []T = []T{array[0]}
	// Remove the first element from the input slice.
	array = array[1:]

	// Continue processing the input slice.
	for len(array) > 0 {
		// Initialize a subList with the first element of the input.
		var subList []T = []T{array[0]}
		// Remove the first element from the input slice.
		array = array[1:]

		// Iterate through the input slice to find elements that can be added to the subList.
		var i int = 0
		for i < len(array) {
			// If the current element is greater than or equal to the last element of the subList, add it to the subList.
			if array[i] >= subList[len(subList)-1] {
				subList = append(subList, array[i])
				// Remove the added element from the input slice.
				array = append(array[:i], array[i+1:]...)
			} else {
				i++
			}
		}

		// Merge the subList into the result using the sMerge function.
		result = sMerge(result, subList)
	}

	return result
}

// sMerge is a helper function to merge two sorted slices 'a' and 'b' into a single sorted slice.
func sMerge[T Numeric](a []T, b []T) []T {
	// Initialize a result slice to hold the merged values.
	var result []T = make([]T, 0, len(a)+len(b))
	var i, j int = 0, 0

	// Merge the two slices while maintaining the sorted order.
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	// Append any remaining elements from slices 'a' and 'b' to the result.
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)

	return result
}
