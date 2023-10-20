package sort

/*



-* The Pigeon Sort *-

 --- Finding Min & Max ---

1. **Function Signature**:
   ```go
   func findMinMax[T Numeric](array []T) (T, T, error)
   ```
   - This function takes an array of generic type `T` as input.
   - It returns two values of type `T`, representing the minimum and maximum values found in the array.
   - If the array is empty, it returns an error.

2. **Error Handling**:
   - The code first checks if the input array is empty (i.e., its length is zero).
   If it is empty, an error is returned with the message "The Array is Empty: Nothing Found."

3. **Initialization**:
   - The code initializes two variables, `maxValue` and `minValue`, both of type `T`,
     with the value of the first element in the input array.
     These variables are used to keep track of the current maximum and minimum values found in the array.

4. **Iteration through the Array**:
   - The code then iterates through each element `num` in the input array using a `for` loop.
   - Inside the loop, it compares `num` with the current `minValue` and `maxValue`.
   - If `num` is smaller than the current `minValue`, it updates `minValue` to be equal to `num`.
   - If `num` is larger than the current `maxValue`, it updates `maxValue` to be equal to `num`.

5. **Result**:
   - After the loop completes, the function returns the final values of `minValue` and `maxValue`,
     which represent the minimum and maximum values found in the input array.
   - If the array was empty, it returns an error indicating that nothing was found.




--- Pigeon Sort ----



1. **Function Signature**:
   ```go
   func Pigeon[T Numeric](array []T) ([]T, error)
   ```
   - This function takes an array of generic type `T` as input.
   - It returns the sorted array and an error (if any).
   - If the input array is empty, it returns an error with the message "The Provided Array is Empty."

2. **Error Handling**:
   - The code first checks if the input array is empty (i.e., its length is zero).
     If it is empty, it returns an error with the message "The Provided Array is Empty."

3. **Finding Min and Max Values**:
   - The code calls the `findMinMax` function to find the minimum (`minValue`)
     and maximum (`maxValue`) values in the input array.
   - If an error occurs while finding the min and max values, it returns that error.

4. **Creating Pigeonholes**:
   - It calculates the `rangeOfValue`, which represents the range of possible values in the input array.
   - It creates a slice called `pigeonHolesBucket` to serve as pigeonholes.
     The length of this slice is determined by the range of values.

5. **Placing Values in Pigeonholes**:
   - The code iterates through the values of the original input array.
   - For each value `num`, it calculates the index where it should be placed in
     the pigeonhole array using `int(num-minValue)`.
   - It increments the value in the corresponding pigeonhole, effectively counting
     how many times each value appears in the original array.

6. **Collecting and Sorting**:
   - The code then collects the values from the pigeonholes in ascending order (from small to large).
   - It uses a loop that iterates through the range of possible values inside the pigeonhole.
   - Within the loop, it collects the smallest value from the current pigeonhole (if any) and places
     it back into the original array.
   - The index variable `index` is used to keep track of the position in the original array where the values are placed.

7. **Result**:
   - After the sorting process is complete, the function returns the sorted `array` and a nil error.



*/

func Pigeon[T Numeric](array []T) []T {
	if len(array) == 0 {
		return array
	}
	// Finding min and max with the help of helper function
	minValue, maxValue := findMinMax(array)

	// Range of possible value - because This determine the Amount of bucket needed for the Collection
	var rangeOfValue T = maxValue - minValue + 1
	// we will create pigeon Holes based on range of value for possible collection of the Values
	pigeonHolesBucket := make([]T, int(rangeOfValue))
	// Iterate through the values of the original array
	for _, num := range array {
		// adding those values into the relative position
		// num-minValue := calculates the index of the pigeonhole where the current element should be placed.
		// and by incrementing ++ we will effectively place the each and every value inside the bucket
		pigeonHolesBucket[int(num-minValue)]++
	}

	/// This part is collecting the element from the Bucket and sorting
	//  it and placing it back to original array as Sorted

	// index is to keep track index at value of the Original Array
	var index int = 0
	// this loop iterate through the range of Values inside the Pigeon Hole in Ascending Order (small to Large)
	for i := 0; i < int(rangeOfValue); i++ {
		// It collect the smallest element from the bucket till pigeon hole is empty
		for pigeonHolesBucket[i] > 0 {
			// array[index] := is the starting position of the Original array
			// i := is the index of the position of the current pigeon Hole Bucket
			// minValue := minimum value in the Original array
			// By adding (i + minValue) minVal to i, we "de-normalize" the value, converting
			// it back to its original value within the range.

			array[index] = minValue + T(i)
			// Increment the Index for preparing the next element
			index++
			// Decrement the values from the bucket to original array till it's empty
			pigeonHolesBucket[i]--
		}
	}

	return array
}
