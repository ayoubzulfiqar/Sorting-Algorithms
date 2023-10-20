package sort

/*


	-* Selection Sort *-


1. **Loop through the entire array:**
     The function starts by looping through the input array
     from the first element (index 0) to the second-to-last element (index `len(array)-1`).
     It iterates through the array, considering each element in turn.

2. **Assume the current element as the minimum:**
    For each iteration, it assumes the element at the current index as the minimum
	element and stores its index in `minIndex`.

3. **Search for a smaller element:**
     It then starts another inner loop, looping through the rest of the array,
	 starting from the element immediately after the current index (`i + 1`).
	 In this inner loop, it compares each element to the element assumed as the minimum.

4. **Finding a smaller element:**
     If an element is found that is smaller than the assumed minimum (i.e., `array[j]`
	 is less than `array[minIndex]`), it updates `minIndex` to the index of the new,
	 smaller element. This inner loop continues until the end of the array.

5. **Swap if necessary:**
     After the inner loop completes, if `minIndex` is not the same as the original assumed
	 minimum (i.e., `minIndex != i`), it means that a smaller element was found in the
	 inner loop. In that case, the function swaps the elements at `i` and `minIndex`,
	 putting the smallest element in its correct position at the beginning of the unsorted
	 portion of the array.

6. **Repeat for the next element:**
     The outer loop then moves to the next element (index `i+1`), assuming that the elements
	 before it are already sorted. The process is repeated, and the next smallest element is
	 selected and placed in the appropriate position.

7. **Repeat until the array is sorted:**
     This process continues, with each iteration placing the next smallest element in the
	 correct position. The outer loop runs until it reaches the second-to-last element,
	 ensuring that the entire array is sorted.

8. **Return the sorted array:**
     Once the outer loop completes, the function returns the sorted array.


*/

func Selection[T Numeric](array []T) []T {
	// looping through the whole array
	for i := 0; i < len(array)-1; i++ {
		// assuming the element in the array is min
		var minIndex int = i
		// looping though the array and checking step by step one at a time and compare it if it is minimum
		for j := i + 1; j < len(array); j++ {
			// element at position at j is less than like minimum to the previous value
			for array[j] < array[minIndex] {
				// we set it as our next min
				minIndex = j
			}
		}
		// if the are not same
		if minIndex != i {
			// than Swap if necessary
			array[i], array[minIndex] = array[minIndex], array[i]
		}
	}
	return array
}
