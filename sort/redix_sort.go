package sort

/*

	 -* Redix Sort *-

The Redix function is the main Radix Sort function.
It first finds the maximum value in the input array to determine the number
of digits in the largest element. Then, it iterates through each digit
place (from least significant to most significant), calling the countingSort
function with the current digit place.

The countingSort function performs the counting sort for a specific digit place (given by exp).
It initializes an output slice of the same size as the input and a count array with
19 elements to count occurrences of values in the range [-9, 9].

It counts the occurrences of each digit at the current digit place by iterating
through the input array.

It calculates the cumulative count to determine the positions of values in the output array.

It builds the sorted output array by placing elements in the correct order based on the
current digit place.

Finally, it copies the sorted values from the output array back to the
input array, and the sorted slice is returned.


*/

func Redix[T Numeric](array []T) []T {
	var max T = findMax(array) // Find the maximum value in the array

	// Iterate through each digit place, from least significant to most significant
	for exp := 1; max/T(exp) > 0; exp *= 10 {
		array = countingSort(array, T(exp)) // Call countingSort for the current digit place
	}
	return array
}

func countingSort[T Numeric](array []T, exp T) []T {
	var n int = len(array)
	var output []T = make([]T, n) // Create an output slice of the same size as the input
	var count []T = make([]T, 19) // Create a count array with 19 elements for values in the range [-9, 9]

	// Count the occurrences of each digit at the current digit place
	for i := 0; i < n; i++ {
		count[int(array[i]/exp)%10+9]++
	}

	// Calculate the cumulative count for each digit
	for i := 1; i < 19; i++ {
		count[i] += count[i-1]
	}

	// Build the output array by placing elements in the correct order
	for i := n - 1; i >= 0; i-- {
		output[int(count[int(array[i]/exp)%10+9])-1] = array[i]
		count[int(array[i]/exp)%10+9]--
	}

	// Copy the sorted values from the output array back to the input array
	for i := 0; i < n; i++ {
		array[i] = output[i]
	}
	return output
}
