package sort

/*



-* The Odd Even Sort (The Brick Sort)*-


- `OddEven` is a generic function that takes a slice of elements of type `T` and returns a sorted
   slice of the same type `T`. The constraint `Numeric` indicates that `T` must be a numeric type (e.g., `int`, `float64`).



- Here, we initialize a boolean variable `sorted` to `false`, which will be used to determine whether
  the array is fully sorted. We also get the length of the input array and store it in `n`.


- This loop continues until the `sorted` variable becomes `true`, indicating that the array is sorted.


- At the beginning of each pass, we assume that the array is already sorted, and if no swaps are
  made during the pass, `sorted` will remain `true`, and the loop will exit.


- This loop iterates over odd indices (`oddIndex`) in the array. It compares and swaps adjacent
  elements at odd indices, following the first step of the Odd-Even Sort algorithm.


- Within the loop, it compares elements at odd indices (`oddIndex`) and their adjacent elements.
  If an element is greater than the one next to it, it swaps them and marks the array as unsorted
  by setting `sorted` to `false`.


- This loop is similar to the previous one but operates on even indices (`evenIndex`) and follows
  the second step of the Odd-Even Sort algorithm.





*/

func OddEven[T Numeric](array []T) []T {
	var sorted bool = false
	var n int = len(array)

	// Continue looping until the array is sorted
	for !sorted {
		sorted = true // Assume the array is sorted initially

		// Odd-Even Sort Step 1: Compare and swap adjacent elements at odd indices
		for oddIndex := 1; oddIndex < n-1; oddIndex += 2 {
			if array[oddIndex] > array[oddIndex+1] {
				array[oddIndex], array[oddIndex+1] = array[oddIndex+1], array[oddIndex]
				sorted = false // Mark as unsorted if a swap occurs
			}
		}

		// Odd-Even Sort Step 2: Compare and swap adjacent elements at even indices
		for evenIndex := 0; evenIndex < n-1; evenIndex += 2 {
			if array[evenIndex] > array[evenIndex+1] {
				array[evenIndex], array[evenIndex+1] = array[evenIndex+1], array[evenIndex]
				sorted = false // Mark as unsorted if a swap occurs
			}
		}
	}

	return array // Return the sorted array
}
