package sort

/*


1. **Merge Sort Algorithm:**
   - Merge Sort is a divide-and-conquer sorting algorithm. It divides the input array into
     smaller halves, recursively sorts these halves, and then merges them back together to
	 produce a sorted array.
   - The primary functions of Merge Sort are splitting the array and merging the sorted halves.

2. **`Merge` Function:**
   - The `Merge` function is the main sorting function. It takes an array of type `T` as
     input and returns the sorted array.

3. **Base Case:**
   - The function starts by checking two base cases:
     - If the input array is `nil` or has only one element, it's already considered
	   sorted, so the function returns the input array as it is.

4. **Splitting the Array:**
   - If the input array has more than one element, it proceeds to split the array into two halves.
   - It calculates the middle index `mid`, which is approximately the midpoint of the array.

5. **Creating Left and Right SubArrays:**
   - The array is divided into two sub-arrays: the left and the right.
   - It creates two separate arrays, `left` and `right`, to hold the elements from these sub-arrays.
   - A loop is used to copy the elements from the original array into the `left` and `right` arrays.

6. **Recursive Calls:**
   - The `Merge` function is recursively called on the `left` and `right` sub-arrays.
     This is where the sorting and splitting process is repeated for each sub-array until
	 the base case is reached.

7. **Merging Sorted SubArrays:**
   - Once the recursive calls return and the sub-arrays are sorted, the function merges
     the two sorted subArrays, `left` and `right`, back into the original array.
   - Three variables, `i`, `j`, and `k`, are used to keep track of the indices for the
     left, right, and merged arrays.

8. **Merging Process:**
   - A `while` loop compares the elements in the `left` and `right` arrays, and the
     smaller of the two is placed in the merged array. This process continues until one
	 of the sub-arrays is exhausted.

9. **Collecting Remaining Elements:**
   - If there are remaining elements in the `left` or `right` sub-arrays, they are
     collected and placed in the merged array.
   - Two separate `for` loops are used for this purpose.

10. **Returning Sorted Array:**
    - Once the `while` loop and `for` loops have completed, the original array
	  is sorted, and the function returns the sorted array.



*/

func PanCake[T Numeric](array []T) []T {
	var n int = len(array)
	// looping at the back of the array then we will decrement as we move towards the left
	for i := n; i > 1; i-- {
		// we will look for the max element inside the array from the backside
		var maxIndex T = findMax(array[:i])
		// if the last element of the array is not the max
		// i-1: represent the last element of the array
		if maxIndex != T(i)-1 {
			// it reverses the order of elements from the beginning of the array up to the largest pancake.
			// The largest pancake will end up at the top of the stack (at index 0).
			array = flip(array, int(maxIndex))
			// It flips the portion of the array from index 0 to i-1.
			// This flip operation effectively moves the largest element to its correct position at
			// the bottom of the sorted portion.
			array = flip(array, i-1)
		}
	}
	return array
}

// This function will reverse the elements in the slice from index 0 to k inclusive.
func flip[T Numeric](array []T, k int) []T {
	// start : starting index of the Array
	var start int = 0
	// end: till at the end index of the array
	var end int = k
	// loop through left to right, if element are smaller than right element
	for start < end {
		// we swap
		array[start], array[end] = array[end], array[start]
		// moving forward toward the right array
		start++
		// moving back to left side of the array
		end--
	}
	return array
}

// Recursive
func RecursivePanCake[T Numeric](array []T) []T {
	var n int = len(array)
	if n == 1 {
		return array
	}
	var maxIndex int = 0
	for i := 0; i < n; i++ {
		if array[i] > array[maxIndex] {
			maxIndex = i
		}
	}
	if maxIndex != 0 {
		array = recursiveFlip(array, maxIndex)
	}
	array = recursiveFlip(array, n-1)
	RecursivePanCake(array[:n-1])
	return array
}
func recursiveFlip[T Numeric](array []T, n int) []T {
	// start : starting index of the Array
	var start int = 0
	// end: till at the end index of the array
	var end int = n
	// loop through left to right, if element are smaller than right element
	for start < end {
		// we swap
		array[start], array[end] = array[end], array[start]
		// moving forward toward the right array
		start++
		// moving back to left side of the array
		end--
	}
	return array
}
