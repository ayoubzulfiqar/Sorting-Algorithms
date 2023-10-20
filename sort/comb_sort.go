package sort

/*


-* The Comb Sort *-



1. **Function Signature:** `func Comb[T Numeric](array []T) []T`

   - `Comb[T Numeric]`: This part indicates that `Comb` is a generic function that can work with
      different numeric data types.
   - `(array []T) []T`: The function takes an input `array` of type `[]T` (a slice of a generic numeric type)
      and returns a sorted slice of the same type.

2. **Variable Initialization:**
   - `n` is initialized to store the length of the input array.
   - `gap` is initialized to the length of the array, representing the gap between compared elements.
   - `shrinkFactor` is set to `1.3` and is used to dynamically adjust the gap during each pass.
   - `swapped` is initialized as `true` and will be used to track whether any swaps have occurred during a pass.

3. **Main Loop:**
   - The function enters a loop that continues until either `gap` becomes greater than `1`
     or no swaps occur during a pass.

4. **Gap Calculation:**
   - Inside the loop, the `gap` value is recalculated using the `shrinkFactor`. It's an integer representation
     of `gap` after dividing it by the `shrinkFactor`.
   - If `gap` becomes less than `1`, it's set to `1` to prevent division by zero.

5. **Swapped Flag Reset:**
   - `swapped` is set to `false` at the beginning of each pass, assuming that no swaps have occurred.

6. **Comparison and Swapping:**
   - A nested loop iterates through the input `array`, comparing elements that are `gap` positions apart.
   - If an element at the current index is greater than the element at `index+gap`, they are swapped.
   - The `swapped` flag is set to `true` to indicate that a swap occurred during this pass.

7. **Repeat Loop:**
   - The outer loop continues until either `gap` becomes greater than `1` or no swaps occur,
     indicating that the array is sorted.

8. **Return Result:**
   - Finally, the sorted array (which is also the input array, modified in place) is returned.



*/

func Comb[T Numeric](array []T) []T {
	// Get the length of the input array
	var n int = len(array)

	// Set the initial gap to the length of the array
	var gap int = n

	// Define the shrink factor for adjusting the gap
	var shrinkFactor float64 = 1.3

	// Initialize a flag to track whether any swaps occurred
	var swapped bool = true

	// Continue looping until the gap is greater than 1 or swaps occurred
	for gap > 1 || swapped {
		// Calculate the gap using the shrink factor
		gap = int(float64(gap) / shrinkFactor)

		// Ensure that the gap is at least 1 to prevent division by zero
		if gap < 1 {
			gap = 1
		}

		// Initialize the swapped flag to false for this pass
		swapped = false

		// Compare and potentially swap elements with the calculated gap
		for index := 0; index < n-gap; index++ {
			if array[index] > array[index+gap] {
				// Swap elements if the condition is met
				array[index], array[index+gap] = array[index+gap], array[index]

				// Set the swapped flag to true to indicate a swap occurred
				swapped = true
			}
		}
	}
	return array
}
