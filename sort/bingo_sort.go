package sort

/*



-* The Bingo Sort *-


**Bingo Function:**
The `Bingo` function is a generic sorting function that implements the Bingo Sort algorithm.
It takes a slice of generic numeric type `T` as input and sorts it in ascending order.

1. **Input Parameter**:
   - `array`: This is the input slice of type `T` that contains the numeric values you want to sort.

2. **Sorting Process**:
   - The function iterates through the input `array` using a loop.

   - During each iteration, it identifies the index of the minimum unsorted element in the
     remaining unsorted portion of the slice using the `findMinUnsorted` helper function.

   - It then swaps the minimum unsorted element with the current element, effectively moving
     the smallest unsorted element to its correct position in the sorted portion of the slice.

3. **Return Value**:
   - The `Bingo` function returns the sorted slice, which now contains the elements in ascending order.

**Find Min Unsorted Function:**
The `findMinUnsorted` function is a helper function used by the `Bingo` function to find the
index of the minimum unsorted element within a specified range of the slice.

1. **Input Parameters**:
   - `array`: This is the input slice of generic numeric type `T` in which you want to find
      the minimum unsorted element.

   - `startIndex`: The `startIndex` parameter specifies the beginning index of the range
      where you want to search for the minimum unsorted element.

   - `endIndex`: The `endIndex` parameter specifies the end index (exclusive) of the range
      where you want to search for the minimum unsorted element.

2. **Finding the Minimum Unsorted Element**:
   - The function initializes a `minIndex` variable with the `startIndex`, assuming that the
     minimum unsorted element is initially at the `startIndex`.

   - It then iterates through the slice within the specified range (from `startIndex` to `endIndex - 1`).

   - During each iteration, it compares the current element with the element at `minIndex`.
     If the current element is smaller, it updates `minIndex` to the current index.

3. **Return Value**:
   - The `findMinUnsorted` function returns the `minIndex`, which represents the index of the
     minimum unsorted element within the specified range of the slice.


*/

// Bingo is a generic function that sorts a slice of numeric values of type 'T'
// using the Bingo Sort algorithm.
func Bingo[T Numeric](array []T) []T {
	// Get the length of the input slice.
	var end int = len(array)
	// Iterate through the slice to perform the sorting.
	for start := 0; start < end; start++ {
		// Find the index of the minimum unsorted element.
		minIndex := findMinUnsorted(array, start, end)

		// Swap the minimum unsorted element with the current element (i).
		array[start], array[minIndex] = array[minIndex], array[start]
	}

	// Return the sorted slice.
	return array
}

// findMinUnsorted finds the index of the minimum unsorted element in a numeric slice.
// It accepts a slice 'array' of type 'T', a 'start' index, and an 'end' index.
// It returns the index of the minimum unsorted element.


// Because the index of the array is always int so wee return int rather then T
func findMinUnsorted[T Numeric](array []T, start, end int) int {
	// Initialize the index of the minimum unsorted element with 'start'.
	minIndex := start

	// Iterate through the slice from 'start + 1' to 'end'.
	for i := start + 1; i < end; i++ {
		// Compare the current element with the element at 'minIndex'.
		if array[i] < array[minIndex] {
			// If the current element is smaller, update 'minIndex' to the current index.
			minIndex = i
		}
	}

	// Return the index of the minimum unsorted element.
	return minIndex
}
