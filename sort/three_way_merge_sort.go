package sort

func merge[T Numeric](array []T, left []T, middle []T, right []T) []T {
	var i int = 0
	var j int = 0
	var k int = 0
	//  Loop until we've processed all elements in left, middle, and right
	for i < len(left) || j < len(middle) || k < len(right) {
		/*

			i < len(left) checks if there are remaining elements in the left sub-array to consider for merging.

			(j == len(middle) || left[i] <= middle[j]) checks if either the middle sub-array is fully processed
			(j equals its length) or if the current element in the left sub-array is smaller than or equal to
			the current element in the middle sub- array.

			(k == len(right) || left[i] <= right[k]) checks if either the right sub-array is fully processed
			(k equals its length) or if the current element in the left sub-array is smaller than or equal
			to the current element in the right sub-array.


		*/
		if i < len(left) && (j == len(middle) || left[i] <= middle[j]) && (k == len(right) || left[i] <= right[k]) {
			array[i+j+k] = left[i]
			i++
			/*

				j < len(middle) checks if there are remaining elements in the middle sub-array to consider for merging.

				(k == len(right) || middle[j] <= right[k]) checks if either the right sub-array is fully processed
				(k equals its length) or if the current element in the middle sub-array is smaller than or
				equal to the current element in the right sub-array.


			*/
		} else if j < len(middle) && (k == len(right) || middle[j] <= right[k]) {
			array[i+j+k] = middle[j]
			j++
			/*
				if none of the previous conditions are met, it means that the current element in the right sub-array
				is smaller than or equal to the elements in the left and middle sub-arrays, or both the left
				and middle sub-arrays have been fully

				In this case, the current element from the right sub-array is selected to be placed in
				the final merged array. The element is assigned to array[i+j+k], and then k is incremented
				to move on to the next element in the right sub-array.
			*/
		} else {
			array[i+j+k] = right[k]
			// moving on
			k++
		}
	}
	return array
}

func ThreeWayMerge[T Numeric](array []T) []T {
	if len(array) > 1 {
		var mid1 int = len(array) / 3
		var mid2 int = 2 * len(array) / 3
		// Split the input array into three parts: left, middle, and right
		var left []T = make([]T, mid1)
		copy(left, array[:mid1])

		var middle []T = make([]T, mid2-mid1)
		copy(middle, array[mid1:mid2])

		var right []T = make([]T, len(array)-mid2)
		copy(right, array[mid2:])
		// Recursively sort the left, middle, and right sub-arrays
		ThreeWayMerge(left)
		ThreeWayMerge(middle)
		ThreeWayMerge(right)
		// Merge the sorted left, middle, and right sub-arrays into the original array
		array = merge(array, left, middle, right)
	}
	return array
}
