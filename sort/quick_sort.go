package sort

import "math/rand"

/*

	In-Place Sorting:
	The algorithm sorts the array in-place without using additional memory for temporary arrays.

*/
func RecursiveQuick[T Numeric](array []T) []T {
	if len(array) < 2 {
		return array
	}

	pivotIndex := twoPartition(array, 0, len(array)-1)

	RecursiveQuick(array[:pivotIndex])
	RecursiveQuick(array[pivotIndex+1:])

	return array
}

/*


   Three-Way Partitioning:
   In the twoPartition function, elements equal to the pivot are not repeatedly compared.
   Instead, they are placed on the correct side of the pivot, reducing unnecessary comparisons



*/
func twoPartition[T Numeric](arr []T, low int, high int) int {
	// Selecting Random Pivot: Why?

	/*

		Randomized Pivot Selection:
		Instead of choosing the pivot as the first or last element,we select a random pivot.
		This helps to mitigate the worst-case scenarios where the input data is already partially sorted.

		Tail Recursion Elimination:
		The code avoids unnecessary recursion on the smaller sub-array by using a while loop for the tail end of the recursion.

	*/

	var pivotIndex int = rand.Intn(high - low + 1)
	// Swapping the Pivot Index with  highest Element inside the array
	arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]
	// setting the highest element inside the array as pivot
	var pivot T = arr[high]
	var index int = low - 1

	for i := low; i < high; i++ {
		if arr[i] <= pivot {
			index++
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	// Place the pivot element in its correct position
	arr[index+1], arr[high] = arr[high], arr[index+1]

	return index + 1
}

// Iterative
func Quick[T Numeric](array []T) []T {
	// if the array have only 1 or no element
	if len(array) <= 2 {
		return array
	}

	stack := []T{0, T(len(array) - 1)}

	for len(stack) > 0 {
		high := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		var low T = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		pivotIndex := itPartition(array, int(low), int(high))

		if pivotIndex-1 > int(low) {
			stack = append(stack, low)
			stack = append(stack, T(pivotIndex)-1)
		}
		if pivotIndex+1 < int(high) {
			stack = append(stack, T(pivotIndex)+1)
			stack = append(stack, high)
		}
	}
	return array
}

// Iterative Partition Function
func itPartition[T Numeric](array []T, low, high int) int {
	// set the highest like the far most element inside the array as the pivot
	var pivot T = array[high]
	var index int = low - 1
	// loop through the array start to end
	for j := low; j < high; j++ {
		// if the element inside the array is the smaller than we will swap
		if array[j] <= pivot {
			index++
			// swap the element at the correct index
			array[index], array[j] = array[j], array[index]
		}
	}
	// swapping the sorted index and placing it to the right position
	array[index+1], array[high] = array[high], array[index+1]
	return index + 1
}
