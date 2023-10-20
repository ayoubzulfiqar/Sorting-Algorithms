package sort

/*


-* The Cycle Sort *-



1. **Loop Through the Array**:
     The code begins by setting up a loop to iterate through the unsorted array.
     This loop is controlled by the `initialPosition` variable and runs from the start of the array
	 to the second-to-last element. It ensures that each element is considered for sorting.

2. **Select the Current Item**:
     For each iteration of the outer loop, the first element of the array is selected as
	 the 'item' that needs to be placed in its correct sorted position.
	 The type `T` is used to represent generic numeric types.

3. **Initialize Variables**: Two variables are initialized:
   - `item T = array[initialPosition]`: `item` stores the value of the current element, which is initially the first element.
   - `position int = initialPosition`: `position` represents the current starting position for the `item`.

4. **Find the Correct Position**:
     A nested loop is used to compare the `item` with elements starting from the second
	 element (`i := initialPosition + 1`) to the end of the array.
	 This loop aims to find the correct position for the `item` by checking if the current
	 element is smaller than the `item`. If a smaller element is found, the `position` is incremented
	 to keep track of where the `item` should be placed.

5. **Skip If Already Sorted**:
     After finding the correct position, the code checks if the `item` is already
     in its correct sorted position by comparing the `position` with the initial starting position (`initialPosition`).
	 If they are the same, it means that the `item` is already in the right place,
	 so the loop continues to the next item using the `continue` statement.
	 This step ensures that smaller elements are already in their correct positions.

6. **Find the Next Available Position**:
     To avoid overwriting existing elements in the array, the code enters another loop to find the
	 next available position for the `item`.
	 It does this by incrementing the `position` until it finds a position where the `item` can be
	 placed without overwriting any values.

7. **Swap the Elements**:
     Once the appropriate position for the `item` is found, the code swaps the `item` with the element at
	 `position`, effectively placing the `item` in its correct sorted position within the array.

8. **Continue Swapping**:
     The code enters another loop to ensure that the `item` reaches its final sorted position within the array.
	 It resets the `position` to the initial starting position and goes through a process similar to
	 steps 4 to 7, finding the next available position for the `item` and swapping it until the `position` returns
	 to the `initialPosition`.

9. **Return the Sorted Array**:
     Finally, the sorted array is returned as the result of the `CycleSort` function.



*/

func Cycle[T Numeric](array []T) []T {
	// Loop through the unsorted array
	for initialPosition := 0; initialPosition < len(array)-1; initialPosition++ {
		// Select the first element as the 'item' to be placed in its correct position
		var item T = array[initialPosition]
		// Initialize 'position' as the current starting position for 'item'
		var position int = initialPosition

		// Find the correct position for the 'item'
		// by comparing it with elements starting from the second element
		for i := initialPosition + 1; i < len(array); i++ {
			// If the current element is smaller than 'item', increment 'position'
			if array[i] < item {
				position++
			}
		}

		// If 'item' is already in its correct position, move on to the next item
		// Ensures that smaller elements are already in their correct positions
		if position == initialPosition {
			continue
		}

		// Find the next available position for 'item' to avoid overwriting elements
		for item == array[position] {
			position++
		}

		// Swap 'item' with the element at 'position', placing 'item' in its correct position
		array[position], item = item, array[position]

		// Continue swapping until 'item' reaches its final sorted position
		for position != initialPosition {
			// Reset 'position' to the initial position
			position = initialPosition

			// Find the next available position for 'item'
			for i := initialPosition + 1; i < len(array); i++ {
				if array[i] < item {
					position++
				}
			}

			// Find the next available position for 'item' to avoid overwriting elements
			for item == array[position] {
				position++
			}

			// Swap 'item' with the element at 'position' to continue its cycle
			array[position], item = item, array[position]
		}
	}
	return array
}
