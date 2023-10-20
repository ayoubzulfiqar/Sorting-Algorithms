package sort

/*

	-* Bucket Sort *-


- The `Bucket` function sorts an array using the Bucket Sort algorithm.
  It's designed to efficiently sort an array of elements, distributing them into
 "buckets" and then combining these buckets to achieve a sorted result.


1. `array`: This is the input array that you want to sort.
    It contains the elements that need to be arranged in ascending order.

2. `minValue` and `maxValue`: These variables are used to find the minimum and maximum
    values within the input array. These values help determine the range of elements to be sorted.

3. `numBuckets`: This variable calculates the number of buckets required for the
    sorting process. It ensures that the buckets cover the entire range of values within the input.

4. `buckets`: This is a two-dimensional slice that represents the buckets where elements
    are temporarily stored during the sorting process.

5. Loop for Distributing Elements:
   - A loop iterates through each element in the input array.
   - For each element, it calculates the index at which it should be placed in
     the `buckets` based on its value relative to `minValue`.
   - The `append` method is used to add the element to the appropriate bucket.

6. Loop for Sorting and Combining Buckets:
   - Another loop iterates through each bucket and checks if it contains any elements.
   - If a bucket has elements, it sorts those elements using a sorting method
     (in this case, it uses the `Insertion` function, but you can use any sorting algorithm you prefer).
   - The `copy` method is used to copy the sorted elements from the bucket back into
     the original `array`. This is done in-place and modifies the original array.
   - The `index` variable keeps track of where the next set of elements from a
     bucket should be placed in the original array. It's incremented by the length of
	 the sorted bucket to ensure elements are placed sequentially.



*/

func Bucket[T Numeric](array []T) []T {
	if len(array) <= 1 {
		return array
	}

	// Find the minimum and maximum values in the array
	var minValue, maxValue T = findMinMax(array)

	// Create buckets
	var numBuckets T = maxValue - minValue + 1
	var buckets [][]T = make([][]T, int(numBuckets))

	// Distribute elements into buckets
	for _, value := range array {
		buckets[int(value-minValue)] = append(buckets[int(value-minValue)], value)
	}

	// Sort each bucket (using any sorting algorithm)
	var index int = 0
	for i := 0; i < int(numBuckets); i++ {
		if len(buckets[i]) > 0 {
			// // Insertion Sort Use any sorting method for the buckets
			Insertion(buckets[i])
			copy(array[index:], buckets[i])
			index += len(buckets[i])

		}
	}
	return array
}
