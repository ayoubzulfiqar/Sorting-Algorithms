package sort

/*


	 -* Insertion Sort *-


1. **Insertion Sort Algorithm:**
   - Insertion Sort is a simple sorting algorithm that works by dividing the array into two
     portions: a sorted portion and an unsorted portion. It iterates through the unsorted
	 portion, comparing elements and shifting them to their correct positions in the sorted portion.
   - It starts with the assumption that the first element (at index 0) is already sorted.
     It then iterates through the rest of the array, one element at a time, and inserts each
	 element into its correct position within the sorted portion.

2. **`Insertion` Function:**
   - The `Insertion` function is the main sorting function. It takes an array of type `T`
     as input and returns the sorted array.
   - It begins by determining the length of the array (`n`).

3. **For Loop (Iteration through Unsorted Portion):**
   - A `for` loop starts from the second element (index 1) and goes up to the end of the array (`n`).
   - The loop iterates through the unsorted portion of the array, one element at a time.

4. **Key and Sorted Index:**
   - Inside the loop, it defines two important variables:
     - `key`: This variable holds the current element that is being compared and potentially
	   inserted into the sorted portion.
     - `sortedIndex`: This variable represents the index of the last element in the sorted
	    portion of the array. It's used to compare with the `key`.

5. **Comparing and Shifting:**
   - The algorithm compares the `key` with elements in the sorted portion and shifts
     them to the right until the correct position for the `key` is found.
   - It does this by using a `while` loop, which continues as long as `sortedIndex` is
     greater than or equal to 0 (meaning we haven't reached the beginning of the sorted portion)
	 and the element at the `sortedIndex` is greater than the `key`.
   - In each iteration, it shifts the element at `sortedIndex` to the right
     (i.e., assigns the value of `array[sortedIndex]` to `array[sortedIndex+1]`) to
	 make room for the `key`. It then decrements `sortedIndex` by 1 to check the next element.

6. **Insertion:**
   - Once the correct position for the `key` is found, the `key` is inserted into its
    correct position within the sorted portion. This is done by assigning the value
	of `key` to `array[sortedIndex+1]`.

7. **Repeat and Return:**
   - The loop continues until all elements in the unsorted portion are processed,
     and the entire array is sorted.
   - Finally, the sorted array is returned by the `Insertion` function.



*/

func Insertion[T Numeric](array []T) []T {
	var n int = len(array)
	// Loop through the elements of the array, starting from index 1,
	// because at position 0, we assume the first element is already in the sorted portion.
	for currentIndex := 1; currentIndex < n; currentIndex++ {
		// key: Holds the current element of the array that we're comparing and potentially inserting.
		var key T = array[currentIndex]
		// sortedIndex: Represents the index of the last element in the sorted portion of the array,
		// which will be compared with the key.
		var sortedIndex int = currentIndex - 1

		// Compare the key with elements in the sorted portion and shift them to the right
		// until we find the correct position for the key.
		for sortedIndex >= 0 && array[sortedIndex] > key {
			array[sortedIndex+1] = array[sortedIndex]
			sortedIndex = sortedIndex - 1
		}
		// Insert the key into its correct position within the sorted portion.
		array[sortedIndex+1] = key
	}
	return array
}

func RecursiveInsertion[T Numeric](array []T) []T {
	var n int = len(array)
	if n <= 1 {
		return array
	}
	RecursiveInsertion(array[:n-1])
	var key T = array[n-1]
	var j int = n - 2
	for j >= 0 && array[j] > key {
		array[j+1] = array[j]
		j--
	}
	array[j+1] = key
	return array
}
