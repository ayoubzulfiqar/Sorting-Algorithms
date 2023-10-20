package sort

/*

	-* Stooge Sort *-


1. **Initialization:**
     The function begins by initializing two variables, `low` and `high`, to represent the
	 low and high indices of the entire array. Initially, `low` is set to 0, and `high`
	 is set to the length of the input array.

2. **Check and Swap:**
     It checks if the last element in the array (indexed at `high-1`) is smaller than the
	 first element in the array (indexed at `low`). If this condition is true, it swaps
	 these two elements. This ensures that the smallest element is at the beginning of
	 the array, as required for Stooge Sort.

3. **Calculate the Length:**
     It calculates the length of the current sub-array by subtracting `low` from `high`.
	 This represents the number of elements to be sorted in the current call.

4. **Recursion with a Divide-and-Conquer Approach:**
     If the length of the sub-array is greater than 2, the function applies
	 a divide-and-conquer approach:

   - It calculates one-third of the length and stores it in the variable `third`.

   - It then calls the `Stooge` function recursively on three sub-arrays:
     1. The first 2/3 of the sub-array (from `low` to `high-third`).
     2. The last 2/3 of the sub-array (from `low+third` to `high`).
     3. The first 2/3 of the sub-array again (from `low` to `high-third`).

   This recursive approach sorts smaller sub-arrays and, eventually, the entire array.

5. **Return Sorted Sub-Array:**
     When the base case is reached (sub-arrays with 2 or fewer elements), the function
	 returns the sorted sub-array.

In essence, Stooge Sort works by dividing the array into three parts, sorting the first 2/3 and
last 2/3 separately, and then sorting the first 2/3 again.
This process repeats recursively until the sub-arrays have only 2 or fewer elements,
and it ensures that the array is sorted correctly.




*/

func Stooge[T Numeric](array []T) []T {
	// Initialize the low and high indices for the entire array.
	var low int = 0
	var high int = len(array)

	// Check if the last element is smaller than the first element, swap them if needed.
	if array[high-1] < array[low] {
		temp := array[low]
		array[low] = array[high-1]
		array[high-1] = temp
	}

	// Calculate the length of the current sub-array.
	length := high - low

	// Check if there are more than 2 elements in the sub-array.
	if length > 2 {
		// Calculate one-third of the length.
		third := length / 3

		// Recursively sort the first 2/3 of the sub-array.
		Stooge(array[low : high-third])

		// Recursively sort the last 2/3 of the sub-array.
		Stooge(array[low+third : high])

		// Recursively sort the first 2/3 of the sub-array again.
		Stooge(array[low : high-third])
	}

	// Return the sorted sub-array.
	return array
}
