package sort

/*


-* The CockTail Sort (Shaker Sort) *-


The provided code is an implementation of the Cocktail Sort algorithm
(also known as Bidirectional Bubble Sort or Shaker Sort).
This sorting algorithm is a variation of the Bubble Sort algorithm and is designed to sort a slice
or array of elements in either ascending or descending order.


- `swapped` is a boolean variable used to keep track of whether any swaps were made during a pass.
   It is initially set to `true` to ensure that the sorting process starts.
- `start` and `end` are integer variables that represent the indices of the first and last
   elements of the unsorted portion of the array, respectively. At the beginning, `start`
   is set to 0, and `end` is set to the length of the input array.

The sorting process is performed inside a loop that continues as long as swaps are being
made during a pass.
Initially, `swapped` is set to `true` to ensure that the loop executes at least once.
The purpose of this loop is to continue sorting until no more swaps are necessary,
indicating that the array is sorted.


In the forward pass, the code iterates through the unsorted portion of the array from
`start` to `end-1`.
It compares adjacent elements and swaps them if they are out of order (i.e., if `array[i]`
is greater than `array[i+1]`).
This pass moves the largest element to the end of the unsorted portion of the array.

After the forward pass, if no swaps were made (`swapped` is `false`), it means that the
array is already
sorted, so the loop can be exited early.

The `swapped` flag is reset to `false` for the backward pass, and the `end` pointer is
decreased by 1 because the largest element is now in its correct position at the
end of the unsorted portion.

In the backward pass, the code iterates through the unsorted portion of the array in
reverse order (from `end-1` down to `start`). It compares adjacent elements and swaps
them if they are out of order.
This pass moves the smallest element to the beginning of the unsorted portion.

After the backward pass, the `start` pointer is increased by 1 because the smallest
element is now in
its correct position at the beginning of the unsorted portion.

Finally, when the loop exits (indicating that the entire array is sorted),
the function returns the sorted `array`.

Elements are compared and swapped in a bidirectional manner until the array is sorted.
It's a variation of Bubble Sort that can perform better in certain cases, particularly when there
are small elements at both ends of the array.


*/

func CockTail[T Numeric](array []T) []T {
	// Initialize variables
	var swapped bool = true  // Flag to check if any swaps are made in a pass
	var start int = 0        // Start of the unsorted portion of the array
	var end int = len(array) // End of the unsorted portion of the array

	// Continue sorting until no swaps are made in a pass
	for swapped {
		swapped = false // Reset the swapped flag at the beginning of each pass

		// Forward pass: Move the largest element to the end
		for i := start; i < end-1; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i] // Swap elements if they are out of order
				swapped = true                              // Mark as swapped
			}
		}

		// If no swaps were made in the forward pass, the array is sorted
		if !swapped {
			break
		}

		swapped = false // Reset the swapped flag for the backward pass
		end = end - 1   // Decrease the end pointer as the largest element is now at the end

		// Backward pass: Move the smallest element to the beginning
		for i := end - 1; i >= start; i-- {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i] // Swap elements if they are out of order
				swapped = true                              // Mark as swapped
			}
		}

		start = start + 1 // Increase the start pointer as the smallest element is now at the beginning
	}

	return array // Return the sorted array
}
