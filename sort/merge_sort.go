package sort

/*

	-* Merge Sort *-


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
     - If the input array is `nil` or has only one element, it's already considered sorted,
	   so the function returns the input array as it is.

4. **Splitting the Array:**
   - If the input array has more than one element, it proceeds to split the array into two halves.
   - It calculates the middle index `mid`, which is approximately the midpoint of the array.

5. **Creating Left and Right SubArrays:**
   - The array is divided into two subArrays: the left and the right.
   - It creates two separate arrays, `left` and `right`, to hold the elements from these subArrays.
   - A loop is used to copy the elements from the original array into the `left` and `right` arrays.

6. **Recursive Calls:**
   - The `Merge` function is recursively called on the `left` and `right` subArrays.
    This is where the sorting and splitting process is repeated for each sub-array until
	the base case is reached.

7. **Merging Sorted SubArrays:**
   - Once the recursive calls return and the subArrays are sorted, the function merges
     the two sorted subArrays, `left` and `right`, back into the original array.
   - Three variables, `i`, `j`, and `k`, are used to keep track of the indices for the
     left, right, and merged arrays.

8. **Merging Process:**
   - A `while`(in go it's for) loop compares the elements in the `left` and `right` arrays,
     and the smaller of the two is placed in the merged array. This process continues
	 until one of the sub-arrays is exhausted.

9. **Collecting Remaining Elements:**
   - If there are remaining elements in the `left` or `right` sub-arrays, they are
     collected and placed in the merged array.
   - Two separate `for` loops are used for this purpose.

10. **Returning Sorted Array:**
    - Once the `while(for)` loop and `for` loops have completed, the original array is
	  sorted, and the function returns the sorted array.




*/

func Merge[T Numeric](array []T) []T {
	if array == nil || len(array) <= 1 {
		return array
	}

	if len(array) > 1 {
		var mid int = len(array) / 2

		// Split left part
		var left []T = make([]T, mid)
		for i := 0; i < mid; i++ {
			left[i] = array[i]
		}

		// Split right part
		var right []T = make([]T, len(array)-mid)
		for i := mid; i < len(array); i++ {
			right[i-mid] = array[i]
		}

		Merge(left)
		Merge(right)

		var i int = 0
		var j int = 0
		var k int = 0

		// Merge left and right arrays
		for i < len(left) && j < len(right) {
			if left[i] < right[j] {
				array[k] = left[i]
				i++
			} else {
				array[k] = right[j]
				j++
			}
			k++
		}

		// Collect remaining elements
		for i < len(left) {
			array[k] = left[i]
			i++
			k++
		}

		for j < len(right) {
			array[k] = right[j]
			j++
			k++
		}
	}
	return array
}

// Recursive
func RecursiveMerge[T Numeric](array []T) []T {
	// Base case: If the array has 0 or 1 element, it is already sorted.
	if len(array) <= 1 {
		return array
	}

	// Split the array into two halves
	var middle int = len(array) / 2
	var left []T = array[:middle]
	var right []T = array[middle:]

	// Recursively sort both halves
	left = RecursiveMerge(left)
	right = RecursiveMerge(right)

	// Merge the two sorted halves
	return merger(left, right)
}

func merger[T Numeric](left, right []T) []T {
	var merged []T = make([]T, 0, len(left)+len(right))

	// Initialize pointers for both halves
	i, j := 0, 0

	// Compare and merge elements from both halves
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	// Append any remaining elements from both halves
	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged
}
