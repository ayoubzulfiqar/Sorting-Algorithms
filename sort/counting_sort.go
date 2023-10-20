package sort

/*

	-* Counting Sort *-


**Variable Explanations:**

1. `array`: The input array that you want to sort.

2. `min` and `max`: These variables store the minimum and maximum values found in the input array. They are used to determine the range of values to be sorted.

3. `counts`: This slice is used to count the occurrences of each unique element in the input array. Each index of the slice corresponds to a unique element, and the value at that index represents the count of occurrences.

4. `output`: This is the output slice that will store the sorted elements.

**Function Explanation:**

1. **Counting Element Occurrences**:
   - The first loop counts the occurrences of each element in the input array. It uses `int(num-min)` as the index to account for variations in element values.
   - The result is stored in the `counts` slice.

2. **Calculate Cumulative Counts**:
   - The second loop calculates the cumulative counts. It updates each element in the `counts` slice by adding the previous count to the current count.

3. **Place Elements in Sorted Order**:
   - The third loop places each element in its correct sorted position in the `output` slice.
   - It uses the `counts` slice to determine the position and decrement the count accordingly.

The `Counting` function is an implementation of the Counting Sort algorithm,
which is an efficient algorithm for sorting integers or other discrete values.
It's particularly useful when the range of values is relatively small compared
to the number of elements in the input array.



*/

func Counting[T Numeric](array []T) []T {
	// If the array has 0 or 1 elements, it's already sorted
	if len(array) <= 1 {
		return array
	}

	// Find the minimum and maximum values in the array
	var min, max T = findMinMax(array)

	// Create a slice to store counts for each unique element
	var counts []T = make([]T, int(max-min+1))

	// Count the occurrences of each element in the input array
	for _, num := range array {
		counts[int(num-min)]++
	}
	var index int = 0
	// Calculate cumulative counts to determine the correct position of each element
	for i := 0; i < len(counts); i++ {
		for counts[i] > 0 {
			// Place the element at the correct sorted position in the original array
			array[index] = T(i) + min
			counts[i]--
			index++
		}
	}

	return array
}
