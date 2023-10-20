package sort

/*

	-* Bogo Sort *-


1. `func Bogo[T Numeric](array []T) []T`:
   - This is a function named `Bogo` that takes a generic type `T` and a slice of
     type `T` named `array` as its input.
   - `Numeric` is used as a type constraint to ensure that the type `T` must be
      numeric (e.g., `int`, `float64`, etc.).

   - The purpose of this function is to perform a "BogoSort" operation
     on the input array, which is a highly inefficient sorting algorithm.
	 BogoSort generates random permutations of the input array until it happens
	 to stumble upon a sorted permutation, which can take a very long time.

   - Inside the function, there's a `for` loop that continues until the `isBogoSorted`
     function returns `true`, indicating that the array is sorted. Within the loop,
	 the array is repeatedly shuffled using the `shuffle` function.
   - Once a sorted permutation is found, the function returns the sorted array.

2. `func isBogoSorted[T Numeric](array []T) bool`:
   - This is a function named `isBogoSorted` that takes a generic type `T` and a slice
     of type `T` named `array` as its input.

   - It also has the `Numeric` type constraint to ensure that `T` is a numeric type.

   - The purpose of this function is to check if the input array is sorted
     or not (in ascending order).

   - It does so by iterating through the elements of the array in a `for` loop.
     If it finds any pair of adjacent elements where the element at the current index
	 is smaller than the element at the previous index, it returns `false`, indicating that
	 the array is not sorted.

   - If it completes the loop without finding such a pair, it returns `true`,
     indicating that the array is sorted.


*/
func Bogo[T Numeric](array []T) []T {
	// while until not sorted we will keep sorting until we run run out of enough life on tis earth
	// to sustain humanity
	for !isBogoSorted(array) {
		// shuffle the array if not sorted
		array = shuffle(array)
	}
	return array
}

func isBogoSorted[T Numeric](array []T) bool {
	// looping through the elements of the array
	for i := 1; i < len(array); i++ {
		// if the element at index 0 is greater than the next one
		if array[i-1] > array[i] {
			// means not sorted
			return false
		}
	}
	// else sorted [almost]
	return true
}
