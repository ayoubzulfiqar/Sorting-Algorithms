package sort

/*


-* The Bubble Sort *-




1. The `Bubble` function takes a generic slice `array` as its input.
   It's designed to work with various numeric types represented by the `T Numeric` constraint.

2. It initializes a boolean flag `swapped` to `true`.
   This flag will help determine if any swaps were made during a pass through the array.

3. It calculates the initial length of the input array `n` using the `len` function.

4. The function enters a loop that continues until the `swapped` flag remains `true`.
   This loop represents the main Bubble Sort algorithm.

5. Inside the loop, it assumes that the array is not sorted at the beginning of each pass, so it
   sets `swapped` to `false`.

6. It then enters another loop, starting from the second element (index 1) of the
   array and comparing adjacent elements.

7. If it finds that the previous element (at index `i-1`) is greater than the current
   element (at index `i`), it swaps the elements to put them in the correct order.

8. After each pass through the array, it checks if any swaps were made (`swapped` is `true`).
   If no swaps were made, it means the array is sorted, and the loop exits.

9. At the end of each pass, the largest element is guaranteed to be at the end of the array, so
   it reduces the range (`n`) by one.

10. Finally, it returns the sorted array.



*/

func Bubble[T Numeric](array []T) []T {
	var swapped bool = true
	// entire length of the array
	var n int = len(array)
	// continues Loop until all the elements inside the array is swapped or Sorted
	for swapped {
		// we assume that the array is not sorted
		swapped = false
		// loop thorough the array at position second
		for i := 1; i < n; i++ {
			// array[i-1]: the first element of the array
			// array[i]: second element of the array
			// Checking if first element of the array is greater than the second element than we need to swap
			if array[i-1] > array[i] {
				array[i-1], array[i] = array[i], array[i-1]
				// outer loop will continue till this condition reached
				swapped = true
			}
		}
		// After each pass, the largest element is at the end, so we reduce the range
		n--
	}
	return array
}

/*



-* The Recursive Bubble Sort *-




1. The function `RecursiveBubble` is a recursive sorting function that takes a
   generic slice `array` as its input, and it's designed to work with various
   numeric types, indicated by the `T Numeric` constraint.

2. It calculates the initial length of the input array `n` using the `len` function.

3. It checks for the base case of the recursion: if the length of the array `n`
   is less than or equal to 1, it means the array is already sorted (or empty),
   so it returns the array as is. This is the stopping condition for the recursion.

4. Inside the recursive function, it enters a loop that starts from the second element (index 1) of
   the array and compares adjacent elements.

5. If it finds that the previous element (at index `i-1`) is greater than the current
   element (at index `i`), it swaps the elements to put them in the correct order,
   just like in the standard Bubble Sort.

6. After the loop, it makes a recursive call to `RecursiveBubble` with a slice of the
   array from the beginning to the second-to-last element (effectively reducing the
   array size by one element). This recursive call continues the sorting process on the
   remaining unsorted portion of the array.

7. The recursion continues until the base case is met (when `n` is less than or equal to 1),
   and each recursive call helps sort a smaller portion of the array.

8. Finally, the function returns the sorted array.





*/


func RecursiveBubble[T Numeric](array []T) []T {
	var n int = len(array)
	if n <= 1 {
		return array
	}
	for i := 1; i < n; i++ {
		if array[i-1] > array[i] {
			array[i-1], array[i] = array[i], array[i-1]
		}
	}
	RecursiveBubble(array[:n-1])
	return array
}
