package sort

/*


	 -* Heap Sort *-

1. The `Heap` function is the main function that performs the Heap Sort algorithm.
   It takes an integer array as input and returns the sorted array.
   The algorithm has two main steps:
   - Step 1: Build a max heap (heapify the array).
   - Step 2: Repeatedly extract elements from the heap (the root of the heap is always the largest element) and rebuild the heap.

2. The `heapify` function is a recursive function that helps maintain the max heap property.
   It takes three parameters: the array, the total number of elements in the heap (`n`),
   and the current element to be considered (`i`).

3. The max heap property ensures that a parent node is always larger than its child nodes,
   which helps in extracting the largest element from the root.

4. The `heapify` function compares the element at index `i` with its left and right child elements.
   If any child element is larger, it swaps them and recursively calls `heapify` on the
   affected subtree to ensure the max heap property is maintained.

5. In the `Heap` function, before each extraction of the largest element,
   the root (largest element) is swapped with the last element in the heap.
   Then, `heapify` is called on the remaining heap to ensure the next largest element is at the root.

6. The process continues until all elements are extracted and the array is sorted.



*/

func Heap[T Numeric](array []T) []T {
	var n int = len(array) // Get the length of the array

	// Step 1: Build a max heap (rearrange the array)
	for i := n/2 - 1; i >= 0; i-- {
		// Call the heapify function to adjust the heap
		array = heapify(array, n, i)
	}

	// Step 2: One by one extract elements from the heap
	for i := n - 1; i >= 0; i-- {
		// Swap the root (largest element) with the last element
		array[0], array[i] = array[i], array[0]

		// Call the heapify function on the reduced heap
		array = heapify(array, i, 0)
	}

	return array // Return the sorted array
}

// Function to heapify a subtree rooted with node i
func heapify[T Numeric](array []T, n, i int) []T {
	var largest int = i     // Initialize the largest element as the root
	var left int = 2*i + 1  // Index of the left child
	var right int = 2*i + 2 // Index of the right child

	// If the left child is larger than the root
	if left < n && array[left] > array[largest] {
		largest = left
	}

	// If the right child is larger than the largest so far
	if right < n && array[right] > array[largest] {
		largest = right
	}

	// If the largest is not the root
	if largest != i {
		// Swap the root and the largest element
		array[i], array[largest] = array[largest], array[i]

		// Recursively call heapify on the affected subtree
		array = heapify(array, n, largest)
	}

	return array // Return the modified array
}
