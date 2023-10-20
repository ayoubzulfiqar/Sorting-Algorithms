package sort

/*

-* The Gnome Sort (Stupid Sort) *-



1. **Function Signature**:
   - `func Gnome[T Numeric](array []T) []T`: This function takes a generic array `array` of type `T`,
      where `T` is a numeric type. It returns a sorted array of the same type `T`.

2. **Initialization**:
   - `index`: The function initializes a variable called `index` to 0. This variable
      represents the current
      position within the array.
   - `n`: It calculates the length of the input array and stores it in a variable called `n`.
     This value is used to determine when the sorting process is complete.

3. **Looping Through the Array**:
   - The function enters a `for` loop that continues as long as `index` is less than `n`.
     This loop is the core of the Gnome Sort algorithm.

4. **Checking Conditions**:
   - Inside the loop, there are two conditions checked:
     - `index == 0`: This condition checks if the current position is at the beginning of the array.
	    If `index` is 0, it means we are already at the starting point.
     - `array[index] >= array[index-1]`: This condition checks if the current element
	    is greater than or
	    equal to the previous element in the array. If this condition is true, it means
		the elements are in the correct order.

5. **Moving Forward**:
   - If either of the conditions mentioned in the previous step is met, we increment `index` by 1.
     This effectively moves us to the next position in the array.

6. **Swapping Elements**:
   - If neither of the conditions is met, it implies that the current element is out
     of order with respect to the previous element. In this case,
	 we perform a swap operation:
     - `array[index]` is swapped with `array[index-1]`, effectively moving the smaller element one position
	   backward in the array.
     - After the swap, we decrement `index` by 1 to continue checking the previous pair of elements.

7. **Loop Continuation**:
   - The loop continues to iterate until `index` reaches or exceeds the length of the array (`n`).
     When this happens, it indicates that the entire array is sorted.

8. **Sorted Array Return**:
   - Finally, the function returns the sorted array. The original array is now sorted in place.



*/

func Gnome[T Numeric](array []T) []T {
	// Initialize the starting position within the array.
	var index int = 0

	// Get the length of the input array.
	var n int = len(array)

	// Loop through the array from the start index to the end of the array.
	for index < n {
		// Check if the current position is at the beginning of the array (index == 0)
		// or if the current element is greater than or equal to the previous element.
		if index == 0 || array[index] >= array[index-1] {
			// Move to the next position in the array.
			index = index + 1
		} else {
			// If not, swap the current element with the previous element.
			array[index], array[index-1] = array[index-1], array[index]

			// Move one position backward.
			index = index - 1
		}
	}

	// Return the sorted array.
	return array
}
