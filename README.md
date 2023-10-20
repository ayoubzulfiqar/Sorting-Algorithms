# Sorting Algorithms

## Table of Content

| Sorting Algorithm          | Difficulty Level     |
|---------------------------|----------------------|
| Bingo Sort Algorithm      | Very Easy            |
| Pigeonhole Sort           | Very Easy            |
| Cycle Sort                | Easy                 |
| Comb Sort                 | Easy                 |
| Bubble Sort               | Easy                 |
| Gnome Sort                | Easy                 |
| Odd-Even Sort / Brick Sort| Easy                 |
| Cocktail Sort             | Easy                 |
| Bitonic Sort              | Moderate             |
| Pancake Sorting           | Moderate             |
| Sleep Sort                | Moderate             |
| Selection Sort            | Moderate             |
| Insertion Sort            | Moderate             |
| Tree Sort                 | Moderate             |
| Stooge Sort               | Moderate             |
| Heap Sort                 | Challenging          |
| ShellSort                 | Challenging          |
| Merge Sort                | Challenging          |
| Quick Sort                | Challenging          |
| TimSort                   | Challenging          |
| 3-way Merge Sort          | Challenging          |
| Counting Sort             | Very Challenging     |
| Radix Sort                | Very Challenging     |
| Bucket Sort               | Very Challenging     |

## Table of Complexity Comparison

| Sorting Algorithm            | Best Case    | Worst Case   | Average Case | Memory | Stable | Method Used              |
| ---------------------------- | ------------ | ------------ | ------------ | ------ | ------ | ------------------------ |
| Selection Sort               | O(n^2)       | O(n^2)       | O(n^2)       | O(1)   | No     | Selection                |
| Bubble Sort                  | O(n)         | O(n^2)       | O(n^2)       | O(1)   | Yes    | Swapping                 |
| Insertion Sort               | O(n)         | O(n^2)       | O(n^2)       | O(1)   | Yes    | Insertion                |
| Merge Sort                   | O(n log n)   | O(n log n)   | O(n log n)   | O(n)   | Yes    | Merging                  |
| Quick Sort                   | O(n log n)   | O(n^2)       | O(n log n)   | O(log n) | No   | Partitioning             |
| Heap Sort                    | O(n log n)   | O(n log n)   | O(n log n)   | O(1)   | No     | Selection (Heapify)      |
| Counting Sort                | O(n + k)     | O(n + k)     | O(n + k)     | O(k)   | Yes    | Counting                 |
| Radix Sort                   | O(nk)        | O(nk)        | O(nk)        | O(n + k) | Yes   | Distribution             |
| Bucket Sort                  | O(n^2)       | O(n^2)       | O(n^2)       | O(n)   | Yes    | Distribution             |
| Bingo Sort Algorithm         | O(n^2)       | O(n^2)       | O(n^2)       | O(n)   | No     | Combining and Sorting    |
| ShellSort                    | O(n log^2 n) | O(n log^2 n) | O(n log^2 n) | O(1)   | No     | Insertion with Gaps      |
| TimSort                      | O(n log n)   | O(n log n)   | O(n log n)   | O(n)   | Yes    | Merging and Insertion    |
| Comb Sort                    | O(n^2)       | O(n^2)       | O(n^2)       | O(1)   | No     | Swapping with Gaps       |
| Pigeonhole Sort              | O(n + N)     | O(n^2)       | O(n + N)     | O(N)   | Yes    | Distribution             |
| Cycle Sort                   | O(n^2)       | O(n^2)       | O(n^2)       | O(1)   | No     | Cyclic Swapping          |
| Cocktail Sort                | O(n)         | O(n^2)       | O(n^2)       | O(1)   | Yes    | Bidirectional Bubble Sort|
| Strand Sort                  | O(n^2)       | O(n^3)       | O(n^2)       | O(n)   | No     | Merging                  |
| Bitonic Sort                 | O(n log^2 n) | O(n log^2 n) | O(n log^2 n) | O(log n) | Yes   | Comparisons and Swaps    |
| Pancake Sorting              | O(n^2)       | O(n^2)       | O(n^2)       | O(1)   | No     | Flipping                 |
| BogoSort or Permutation Sort | O((n+1)!)    | O(∞)         | O((n+1)!)    | O(1)   | No     | Random Permutations      |
| Gnome Sort                   | O(n^2)       | O(n^2)       | O(n^2)       | O(1)   | Yes    | Insertion with BackStep  |
| Sleep Sort – The King of Laziness | O(n)     | O(n)         | O(n)         | O(n)   | No     | Sleep and WakeUp         |
| Stooge Sort                  | O(n^(log 3)) | O(n^(log 3)) | O(n^(log 3)) | O(1)   | No     | Recursive Trisection     |
| Tree Sort                    | O(n log n)   | O(n^2)       | O(n log n)   | O(n)   | Yes    | Binary Search Tree       |
| Odd-Even Sort / Brick Sort   | O(n^2)       | O(n^2)       | O(n^2)       | O(1)   | Yes    | Odd-Even Comparison      |
| 3-way Merge Sort            | O(n log n)   | O(n log n)   | O(n log n)   | O(n)   | Yes    | Merging (3-way)         |
| 3-way Quick Sort            | O(n log n)   | O(n log n)   | O(n log n)   | O(n^2)   | Yes    | Quick (3-way)         |

## Generic

The Numeric interface defined here can be useful when implementing sorting algorithms in Go that need to work with different numeric data types.

By defining a Numeric interface like this that includes the primitive int, float etc types, we can make the Sort function work with any slice of basic numeric data.

When implementing the actual sort algorithm, we don't need to know the specific type, just that it satisfies the Numeric interface by being one of the allowed types.

This allows us to sort ints, floats, uint32 etc all with the same code, without needing type assertions or conversions.

The **`~`** ensures only the exact types are included, so we know slices of int16 or float64 will not be passed in, only the types defined in the interface.

This predictability can make the sorting logic simpler and faster by avoiding type checking.
We also avoid boxing by working with the primitive types directly.

```go
type Numeric interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
```

## 1. Bingo Sort

The `Bingo` function is a generic sorting function that implements the Bingo Sort algorithm.
It takes a slice of generic numeric type `T` as input and sorts it in ascending order.

1. **Input Parameter**:
   - `array`: This is the input slice of type `T` that contains the numeric values you want to sort.

2. **Sorting Process**:
   - The function iterates through the input `array` using a loop.

   - During each iteration, it identifies the index of the minimum unsorted element in the
     remaining unsorted portion of the slice using the `findMinUnsorted` helper function.

   - It then swaps the minimum unsorted element with the current element, effectively moving
     the smallest unsorted element to its correct position in the sorted portion of the slice.

3. **Return Value**:
   - The `Bingo` function returns the sorted slice, which now contains the elements in ascending order.

**Find Min Unsorted Function:**
The `findMinUnsorted` function is a helper function used by the `Bingo` function to find the
index of the minimum unsorted element within a specified range of the slice.

1. **Input Parameters**:
   - `array`: This is the input slice of generic numeric type `T` in which you want to find
      the minimum unsorted element.

   - `startIndex`: The `startIndex` parameter specifies the beginning index of the range
      where you want to search for the minimum unsorted element.

   - `endIndex`: The `endIndex` parameter specifies the end index (exclusive) of the range
      where you want to search for the minimum unsorted element.

2. **Finding the Minimum Unsorted Element**:
   - The function initializes a `minIndex` variable with the `startIndex`, assuming that the
     minimum unsorted element is initially at the `startIndex`.

   - It then iterates through the slice within the specified range (from `startIndex` to `endIndex - 1`).

   - During each iteration, it compares the current element with the element at `minIndex`.
     If the current element is smaller, it updates `minIndex` to the current index.

3. **Return Value**:
   - The `findMinUnsorted` function returns the `minIndex`, which represents the index of the
     minimum unsorted element within the specified range of the slice.

```go
func Bingo[T Numeric](array []T) []T {
	// Get the length of the input slice.
	var end int = len(array)
	// Iterate through the slice to perform the sorting.
	for start := 0; start < end; start++ {
		// Find the index of the minimum unsorted element.
		minIndex := findMinUnsorted(array, start, end)

		// Swap the minimum unsorted element with the current element (i).
		array[start], array[minIndex] = array[minIndex], array[start]
	}

	// Return the sorted slice.
	return array
}

// findMinUnsorted finds the index of the minimum unsorted element in a numeric slice.
// It accepts a slice 'array' of type 'T', a 'start' index, and an 'end' index.
// It returns the index of the minimum unsorted element.


// Because the index of the array is always int so wee return int rather then T
func findMinUnsorted[T Numeric](array []T, start, end int) int {
	// Initialize the index of the minimum unsorted element with 'start'.
	minIndex := start

	// Iterate through the slice from 'start + 1' to 'end'.
	for i := start + 1; i < end; i++ {
		// Compare the current element with the element at 'minIndex'.
		if array[i] < array[minIndex] {
			// If the current element is smaller, update 'minIndex' to the current index.
			minIndex = i
		}
	}

	// Return the index of the minimum unsorted element.
	return minIndex
}
```

### Space & Time Complexity

| Complexity Type            | Complexity      |
|----------------------------|-----------------|
| Time Complexity            | O(n^2)          |
| Best Case Time Complexity  | O(n^2)          |
| Worst Case Time Complexity | O(n^2)          |
| Average Case Time Complexity | O(n^2)        |
| Space Complexity           | O(1) (In-Place) |

Explanation:

- **Time Complexity**: The time complexity of Bingo Sort is O(n^2), where "n" is the number of elements in the array. This is because it involves two nested loops, where the outer loop runs "n" times, and the inner loop runs up to "n" times in the worst case, resulting in a quadratic time complexity.

- **Best Case Time Complexity**: In the best case, the time complexity is still O(n^2) because the algorithm always compares and potentially swaps elements, even if the array is already sorted.

- **Worst Case Time Complexity**: In the worst case, the time complexity is also O(n^2) when the input array is in reverse order, as the algorithm will make a significant number of comparisons and swaps.

- **Average Case Time Complexity**: On average, the time complexity remains O(n^2) because the algorithm does not employ any special strategies to reduce the number of comparisons or swaps.

- **Space Complexity**: The space complexity of Bingo Sort is O(1) because it operates in-place, meaning it does not require additional memory that grows with the size of the input. It only uses a constant amount of additional memory for variables like `minIndex` and `i`.

## Bogo Sort [Permutation Sort || Stupid Sort || Slow Sort || Shotgun Sort || Monkey Sort]

Bogo sort is a highly inefficient and humorously impractical sorting algorithm. It is not a suitable sorting method for practical use, as its average and worst-case time complexity is factorial, which makes it incredibly slow. In Bogo sort, the algorithm repeatedly shuffles the elements in the list randomly and checks if they are sorted. This process continues until the list happens to be sorted by chance.

Bogo sort uses 2 steps to sort elements of the array.

1. It throws the number randomly.
2. Check whether the number is sorted or not.
3. If sorted then return the sorted array.
4. Otherwise it again generate another randomization of the numbers until the array is sorted.

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

```go
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

// shuffle function shuffles the array to mitigate worst-case scenarios.
func shuffle[T Numeric](arr []T) []T {

	// Iterate through the array and swap elements randomly to shuffle.
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
    return arr
}
```

### Space and Time Complexity

| Complexity       | Time                  | Space   |
| ---------------- | --------------------- | ------- |
| **Worst-case**   | O((n+1)!)             | O(n)    |
| **Best-case**    | O(n)                  | O(n)    |
| **Average-case** | O((n+1)!/2)           | O(n)    |

**Time Complexity:**

1. **Worst-case Time Complexity (O((n+1)!)):** The worst-case time complexity for Bogo Sort is extremely high and impractical. It is represented as O((n+1)!), where "n" is the number of elements in the array. This means that the time it takes to sort the array increases factorially with the number of elements. For even moderately-sized arrays, this time complexity makes Bogo Sort prohibitively slow. In the worst case, Bogo Sort may take an extremely long time to sort an array or even never complete due to its randomized nature.

2. **Best-case Time Complexity (O(n)):** The best-case time complexity occurs when, by sheer luck, the input array is already sorted. In this scenario, Bogo Sort has a time complexity of O(n), which is linear. This happens because the algorithm checks if the array is sorted and, in the best case, it immediately recognizes that the array is sorted and terminates.

3. **Average-case Time Complexity (O((n+1)!/2)):** The average-case time complexity of Bogo Sort is also highly inefficient. It is estimated to be approximately O((n+1)!/2), which is still factorial but on average, it takes half the time of the worst case. However, Bogo Sort's randomized nature makes it highly unpredictable, so the actual time it takes to sort an array can vary widely.

**Space Complexity:**

The space complexity of Bogo Sort is relatively simple:

- **Space Complexity (O(n)):** Bogo Sort has a space complexity of O(n), where "n" is the number of elements in the array. This space complexity is primarily due to the storage of the input array and temporary variables used in the shuffling and sorting process. The space required for Bogo Sort is linear and directly proportional to the size of the input array.

In summary, Bogo Sort is a highly impractical sorting algorithm with abysmally high worst-case and average-case time complexities, making it unsuitable for any real-world sorting task. It is used more as a humorous or educational example rather than for practical sorting. More efficient sorting algorithms should be used in real applications.

Please note that these complexities are for theoretical purposes and are based on the random shuffling of the elements. In practice, Bogo Sort is never a reasonable choice for sorting due to its extreme inefficiency.

## Selection Sort

**Loop through the entire array:**
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

```go
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
```

### Time and Space Complexity

| Complexity      | Best       | Worst      | Average    | Space    |
|-----------------|------------|------------|------------|----------|
| Time Complexity | O(n^2)     | O(n^2)     | O(n^2)     | O(1)     |
| Space Complexity| O(1)       | O(1)       | O(1)       | O(1)     |

- **Time Complexity**:
  - Best Case: O(n^2) - Occurs when the array is already sorted in ascending order, and no swaps are needed.
  - Worst Case: O(n^2) - Occurs when the array is sorted in descending order, resulting in the maximum number of comparisons and swaps.
  - Average Case: O(n^2) - On average, the algorithm requires n^2/2 comparisons and n/2 swaps.
  
- **Space Complexity**:
  - The space complexity is constant, O(1), because the algorithm only uses a fixed amount of extra memory for variables, regardless of the input size.

## Pigeon Sort

First is to find the Min and Max value inside the Unsorted Array then Calculate the Range of possible value by subtracting the Min Value inside the Unsorted Array by Max value Calculation of the range is because This determine the Amount of bucket needed for the Collection of Values inside the array then create an bucket array (pigeon holes) adn place the value of the original array inside that bucket then Iterate though the range of values and based on the the bucket hole index we will start adding the value inside the original array because it will be getting sorted starting at the beginning and do it step by step while emptying the pigeon hole bucket until nothing left

1. **Function Signature**:

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

The Pigeonhole Sort algorithm sorts elements by distributing them into buckets (pigeonholes) based on their values within a known range, and then collecting them back in ascending order. The critical steps involve calculating the range, creating pigeonholes, distributing values, and carefully collecting elements from the pigeonholes to achieve the desired sorting result.

```go
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

// Finding the min and max value inside the Unsorted Array
func findMinMax[T Numeric](array []T) (T, T) {
	// base case - if our array has nothing that It will return Nothing
	if len(array) == 0 {
		return T(0), T(0)
	}

	// Lets Assume that that the First element in the array is the min and also max

	var maxValue T = array[0]
	var minValue T = array[0]
	// We iterate though each and every value inside the array and compare it and assign the correct
	// and corresponding value to min and max
	for _, num := range array {
		if num < minValue {
			minValue = num
		}
		if num > maxValue {
			maxValue = num
		}
	}

	return minValue, maxValue
}
```

### Space & Time Complexity

| Complexity Type          | Worst Case    | Best Case     | Average Case  | Space Complexity |
|--------------------------|---------------|---------------|---------------|-------------------|
| Time Complexity          | O(n + N)      | O(n + N)      | O(n + N)      | O(N)              |
| Space Complexity         | O(N)          | O(N)          | O(N)          | O(N)              |

Explanation:

- **Time Complexity**:
  - In the worst case, Pigeonhole Sort has a time complexity of O(n + N), where 'n' is the number of elements in the input array, and 'N' is the range of possible values within the array.
  - In the best case and average case, Pigeonhole Sort still has a time complexity of O(n + N). This is because the distribution and collection phases both depend on the range 'N' and the number of elements 'n' in the array.

- **Space Complexity**:
  - The space complexity of Pigeonhole Sort is O(N) because it requires additional space to store the pigeonholes (buckets), and the number of pigeonholes is determined by the range 'N' of possible values.
  - This space complexity remains the same in both the worst case, best case, and average case.

Keep in mind that Pigeonhole Sort's efficiency heavily depends on the range of values 'N.' If 'N' is significantly larger than 'n' (the number of elements to be sorted), the algorithm may become inefficient due to the large number of empty pigeonholes, resulting in increased space usage and potentially slower sorting times.

## Cycle Sort

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

```go
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
```

### Space & Time Complexity

| Complexity        | Worst Case    | Average Case  | Best Case     |
|-------------------|---------------|---------------|---------------|
| Time Complexity   | O(n^2)        | O(n^2)        | O(n^2)        |
| Space Complexity  | O(1)          | O(1)          | O(1)          |

Explanation:

- The worst-case time complexity is O(n^2) because in the worst case, each element may need to go through multiple swaps to reach its final sorted position.
- The average-case time complexity is also O(n^2) because it typically involves a significant number of comparisons and swaps.
- The best-case time complexity is O(n^2) as well because even if the array is already partially sorted, the algorithm still performs a substantial number of comparisons and swaps.
- The space complexity is constant O(1) because Cycle Sort is an in-place sorting algorithm that doesn't require additional memory allocation that depends on the input size.

## Comb Sort

Comb Sort is a sorting algorithm that combines the principles of Bubble Sort and the concept of dynamically changing gap sizes. This guide will walk you through the key steps of the Comb Sort algorithm, providing explanations along the way.

 1. **Function Signature:** `func Comb[T Numeric](array []T) []T`

- `Comb[T Numeric]`: This part indicates that `Comb` is a generic function that can work with
      different numeric data types.
- `(array []T) []T`: The function takes an input `array` of type `[]T` (a slice of a generic numeric type)
      and returns a sorted slice of the same type.

2. **Variable Initialization:**
   - `n` is initialized to store the length of the input array.
   - `gap` is initialized to the length of the array, representing the gap between compared elements.
   - `shrinkFactor` is set to `1.3` and is used to dynamically adjust the gap during each pass.
   - `swapped` is initialized as `true` and will be used to track whether any swaps have occurred during a pass.

3. **Main Loop:**
   - The function enters a loop that continues until either `gap` becomes greater than `1`
     or no swaps occur during a pass.

4. **Gap Calculation:**
   - Inside the loop, the `gap` value is recalculated using the `shrinkFactor`. It's an integer representation
     of `gap` after dividing it by the `shrinkFactor`.
   - If `gap` becomes less than `1`, it's set to `1` to prevent division by zero.

5. **Swapped Flag Reset:**
   - `swapped` is set to `false` at the beginning of each pass, assuming that no swaps have occurred.

6. **Comparison and Swapping:**
   - A nested loop iterates through the input `array`, comparing elements that are `gap` positions apart.
   - If an element at the current index is greater than the element at `index+gap`, they are swapped.
   - The `swapped` flag is set to `true` to indicate that a swap occurred during this pass.

7. **Repeat Loop:**
   - The outer loop continues until either `gap` becomes greater than `1` or no swaps occur,
     indicating that the array is sorted.

8. **Return Result:**
   - Finally, the sorted array (which is also the input array, modified in place) is returned.

### The Gap Factor

The "gap" is a crucial concept in the Comb Sort algorithm, and it plays a significant role in the algorithm's performance. It determines how far apart elements are compared and swapped during each pass of the algorithm. The gap is important for the following reasons:

**1. Influence on Sorting Efficiency:**

- The gap size directly impacts the algorithm's sorting efficiency. A larger gap allows for faster movement of elements across the array but may result in fewer comparisons and swaps during each pass. A smaller gap, on the other hand, ensures more thorough comparisons but requires more passes to complete the sorting process.

**2. Adaptive Shrinkage:**

- One of the unique features of the Comb Sort algorithm is that the gap size starts large and progressively shrinks with each pass. This adaptive shrinkage is essential for the algorithm's efficiency because it combines the benefits of both large and small gap sizes.
- Initially, with a large gap, the algorithm can quickly move larger elements toward the end of the array, similar to the way Bubble Sort works. This reduces the total number of larger elements that need to be moved later.
- As the gap decreases, the algorithm performs more fine-grained comparisons and swaps, ensuring that smaller elements are placed in their correct positions.

**3. Finding an Optimal Gap:**

- Choosing an optimal gap size is essential for optimizing the performance of Comb Sort. The choice of gap and the shrink factor (typically around 1.3) can significantly affect the sorting speed.
- Experimenting with different gap sizes and shrink factors is often necessary to find the best combination for a specific dataset. Different datasets may benefit from different gap sizes, so it's not a one-size-fits-all parameter.

**4. Time Complexity Considerations:**

- The time complexity of Comb Sort depends on the choice of the gap size. In its worst-case scenario, if the gap is set as 1, Comb Sort degenerates into a Bubble Sort, resulting in a time complexity of O(n^2). However, with a properly chosen gap, it can achieve an average-case time complexity of approximately O(n*log(n)).

```go
func Comb[T Numeric](array []T) []T {
	// Get the length of the input array
	var n int = len(array)

	// Set the initial gap to the length of the array
	var gap int = n

	// Define the shrink factor for adjusting the gap
	var shrinkFactor float64 = 1.3

	// Initialize a flag to track whether any swaps occurred
	var swapped bool = true

	// Continue looping until the gap is greater than 1 or swaps occurred
	for gap > 1 || swapped {
		// Calculate the gap using the shrink factor
		gap = int(float64(gap) / shrinkFactor)

		// Ensure that the gap is at least 1 to prevent division by zero
		if gap < 1 {
			gap = 1
		}

		// Initialize the swapped flag to false for this pass
		swapped = false

		// Compare and potentially swap elements with the calculated gap
		for index := 0; index < n-gap; index++ {
			if array[index] > array[index+gap] {
				// Swap elements if the condition is met
				array[index], array[index+gap] = array[index+gap], array[index]

				// Set the swapped flag to true to indicate a swap occurred
				swapped = true
			}
		}
	}
	return array
}
```

| Complexity         | Worst Case                              | Average Case                         | Best Case                                | Space Complexity |
|--------------------|----------------------------------------|-------------------------------------|------------------------------------------|------------------|
| Time Complexity    | O(n^2)                                   | O(n^2)                                | O(n*log(n)) [Optimal Gap Sequence]        | O(1)             |
| Explanation        | In the worst case, when the gap is set to 1 and the array is in reverse order or nearly so, Comb Sort becomes inefficient and takes quadratic time. | In most practical cases, Comb Sort's adaptive gap size reduction provides better performance than Bubble Sort but still exhibits quadratic time complexity. | In the best case, when an optimal gap sequence is used and data is partially sorted, Comb Sort approaches a time complexity of n*log(n). Achieving this best-case performance is challenging in practice. | Comb Sort is an in-place sorting algorithm, which means it uses a constant amount of extra memory for variables like the gap and loop indices. |

## Bubble Sort

1. The `Bubble` function takes a generic slice `array` as its input.
   It's designed to work with various numeric types represented by the `T Numeric` constraint.

2. It initializes a boolean flag `swapped` to `true`.
   This flag will help determine if any swaps were made during a pass through the array.

3. It calculates the initial length of the input array `n` using the `len` function.

4. The function enters a loop that continues until the `swapped` flag remains `true`.
   This loop represents the main Bubble Sort algorithm.

5. Inside the loop, it assumes that the array is not sorted at the beginning of each pass, so it
   sets `swapped` to `false`.

6. It then enters another loop, starting from the second element (index 1) of the
   array and comparing adjacent elements.

7. If it finds that the previous element (at index `i-1`) is greater than the current
   element (at index `i`), it swaps the elements to put them in the correct order.

8. After each pass through the array, it checks if any swaps were made (`swapped` is `true`).
   If no swaps were made, it means the array is sorted, and the loop exits.

9. At the end of each pass, the largest element is guaranteed to be at the end of the array, so
   it reduces the range (`n`) by one.

10. Finally, it returns the sorted array.

```go
func Bubble[T Numeric](array []T) []T {
	var swapped bool = true
	// entire length of the array
	var n int = len(array)
	// continues Loop until all the elements inside the array is swapped or Sorted
	for swapped {
		// we assume that the array is not sorted
		swapped = false
		// loop thorough the array at position second
		for i := 1; i < n; i++ {
			// array[i-1]: the first element of the array
			// array[i]: second element of the array
			// Checking if first element of the array is greater than the second element than we need to swap
			if array[i-1] > array[i] {
				array[i-1], array[i] = array[i], array[i-1]
				// outer loop will continue till this condition reached
				swapped = true
			}
		}
		// After each pass, the largest element is at the end, so we reduce the range
		n--
	}
	return array
}

// Recursive Bubble

func RecursiveBubble[T Numeric](array []T) []T {
	var n int = len(array)
	if n <= 1 {
		return array
	}
	for i := 1; i < n; i++ {
		if array[i-1] > array[i] {
			array[i-1], array[i] = array[i], array[i-1]
		}
	}
	RecursiveBubble(array[:n-1])
	return array
}
```

### Space & Time Complexity

| Bubble Sort      | Best Case | Average Case | Worst Case  | Space Complexity |
|------------------|---------------------------|------------------------------|---------------------------|------------------|
| Bubble Sort      | O(n)                      | O(n^2)                       | O(n^2)                    | O(1)             |

Explanation:

- Best Case Time Complexity: O(n) - This is when the input array is already sorted, and Bubble Sort makes only one pass to confirm the array is sorted.
- Average Case Time Complexity: O(n^2) - This is the expected time complexity when elements are randomly ordered, and Bubble Sort makes multiple passes.
- Worst Case Time Complexity: O(n^2) - This occurs when the input array is sorted in reverse order, and Bubble Sort makes the maximum number of passes with many swaps.
- Space Complexity: O(1) - Bubble Sort is an in-place sorting algorithm, meaning it doesn't require additional memory proportional to the input size. It has a constant space complexity.

## Gnome Sort (Stupid Sort)

1. **Function Signature**:
   - `func Gnome[T Numeric](array []T) []T`: This function takes a generic array `array` of type `T`,
      where `T` is a numeric type. It returns a sorted array of the same type `T`.

2. **Initialization**:
   - `index`: The function initializes a variable called `index` to 0. This variable
      represents the current
      position within the array.
   - `n`: It calculates the length of the input array and stores it in a variable called `n`.
     This value is used to determine when the sorting process is complete.

3. **Looping Through the Array**:
   - The function enters a `for` loop that continues as long as `index` is less than `n`.
     This loop is the core of the Gnome Sort algorithm.

4. **Checking Conditions**:
   - Inside the loop, there are two conditions checked:
     - `index == 0`: This condition checks if the current position is at the beginning of the array.
	    If `index` is 0, it means we are already at the starting point.
     - `array[index] >= array[index-1]`: This condition checks if the current element
	    is greater than or
	    equal to the previous element in the array. If this condition is true, it means
		the elements are in the correct order.

5. **Moving Forward**:
   - If either of the conditions mentioned in the previous step is met, we increment `index` by 1.
     This effectively moves us to the next position in the array.

6. **Swapping Elements**:
   - If neither of the conditions is met, it implies that the current element is out
     of order with respect to the previous element. In this case,
	 we perform a swap operation:
     - `array[index]` is swapped with `array[index-1]`, effectively moving the smaller element one position
	   backward in the array.
     - After the swap, we decrement `index` by 1 to continue checking the previous pair of elements.

7. **Loop Continuation**:
   - The loop continues to iterate until `index` reaches or exceeds the length of the array (`n`).
     When this happens, it indicates that the entire array is sorted.

8. **Sorted Array Return**:
   - Finally, the function returns the sorted array. The original array is now sorted in place.

```go
func Gnome[T Numeric](array []T) []T {
	// Initialize the starting position within the array.
	var index int = 0

	// Get the length of the input array.
	var n int = len(array)

	// Loop through the array from the start index to the end of the array.
	for index < n {
		// Check if the current position is at the beginning of the array (index == 0)
		// or if the current element is greater than or equal to the previous element.
		if index == 0 || array[index] >= array[index-1] {
			// Move to the next position in the array.
			index = index + 1
		} else {
			// If not, swap the current element with the previous element.
			array[index], array[index-1] = array[index-1], array[index]

			// Move one position backward.
			index = index - 1
		}
	}

	// Return the sorted array.
	return array
}
```

### Space and Time Complexity

| Complexity      | Best Case       | Average Case    | Worst Case      | Space Complexity |
|-----------------|-----------------|-----------------|-----------------|------------------|
| Time Complexity | O(n)            | O(n^2)          | O(n^2)          | O(1)             |
| Space Complexity| O(1)            | O(1)            | O(1)            | O(1)             |

Explanation:

- **Time Complexity**:
  - Best Case: O(n) - The best-case scenario occurs when the input array is already sorted. In this case, Gnome Sort will make a single pass through the array.
  - Average Case: O(n^2) - In the average case, Gnome Sort makes roughly n^2/4 comparisons and swaps.
  - Worst Case: O(n^2) - The worst-case scenario occurs when the input array is sorted in reverse order. Gnome Sort will make the maximum number of comparisons and swaps, resulting in quadratic time complexity.

- **Space Complexity**:
  - O(1) - Gnome Sort is an in-place sorting algorithm, meaning it doesn't require additional memory or data structures proportional to the input size. It has a constant space complexity.

## OddEven Sort (Brick Sort)

- `OddEven` is a generic function that takes a slice of elements of type `T` and returns a sorted
   slice of the same type `T`. The constraint `Numeric` indicates that `T` must be a numeric type (e.g., `int`, `float64`).

- Here, we initialize a boolean variable `sorted` to `false`, which will be used to determine whether
  the array is fully sorted. We also get the length of the input array and store it in `n`.

- This loop continues until the `sorted` variable becomes `true`, indicating that the array is sorted.

- At the beginning of each pass, we assume that the array is already sorted, and if no swaps are
  made during the pass, `sorted` will remain `true`, and the loop will exit.

- This loop iterates over odd indices (`oddIndex`) in the array. It compares and swaps adjacent
  elements at odd indices, following the first step of the Odd-Even Sort algorithm.

- Within the loop, it compares elements at odd indices (`oddIndex`) and their adjacent elements.
  If an element is greater than the one next to it, it swaps them and marks the array as unsorted
  by setting `sorted` to `false`.

- This loop is similar to the previous one but operates on even indices (`evenIndex`) and follows
  the second step of the Odd-Even Sort algorithm.

```go
func OddEven[T Numeric](array []T) []T {
	var sorted bool = false
	var n int = len(array)

	// Continue looping until the array is sorted
	for !sorted {
		sorted = true // Assume the array is sorted initially

		// Odd-Even Sort Step 1: Compare and swap adjacent elements at odd indices
		for oddIndex := 1; oddIndex < n-1; oddIndex += 2 {
			if array[oddIndex] > array[oddIndex+1] {
				array[oddIndex], array[oddIndex+1] = array[oddIndex+1], array[oddIndex]
				sorted = false // Mark as unsorted if a swap occurs
			}
		}

		// Odd-Even Sort Step 2: Compare and swap adjacent elements at even indices
		for evenIndex := 0; evenIndex < n-1; evenIndex += 2 {
			if array[evenIndex] > array[evenIndex+1] {
				array[evenIndex], array[evenIndex+1] = array[evenIndex+1], array[evenIndex]
				sorted = false // Mark as unsorted if a swap occurs
			}
		}
	}

	return array // Return the sorted array
}
```

### Space and Time Complexity

| Complexity          | Worst Case     | Average Case  | Best Case     |
|---------------------|----------------|---------------|---------------|
| Time Complexity     | O(n^2)         | O(n^2)        | O(n)          |
| Space Complexity    | O(1)           | O(1)          | O(1)          |

- **Time Complexity**:
  - Worst Case: In the worst case, when the array is completely unsorted, Odd-Even Sort makes many passes and comparisons. It has a time complexity of O(n^2), where n is the number of elements in the array.
  - Average Case: On average, the number of comparisons and swaps remains high, leading to an average time complexity of O(n^2).
  - Best Case: In the best case, when the array is already sorted, the algorithm performs significantly fewer comparisons and swaps. The best-case time complexity is O(n).

- **Space Complexity**:
  - Regardless of the input data, Odd-Even Sort uses a constant amount of extra space for variables (e.g., `sorted`, `oddIndex`, `evenIndex`, etc.). Therefore, the space complexity is O(1), indicating that it is an in-place sorting algorithm.

## CockTail Sort

1. The provided code is an implementation of the Cocktail Sort algorithm
(also known as Bidirectional Bubble Sort or Shaker Sort).
This sorting algorithm is a variation of the Bubble Sort algorithm and is designed to sort a slice
or array of elements in either ascending or descending order.

- `swapped` is a boolean variable used to keep track of whether any swaps were made during a pass.
   It is initially set to `true` to ensure that the sorting process starts.
- `start` and `end` are integer variables that represent the indices of the first and last
   elements of the unsorted portion of the array, respectively. At the beginning, `start`
   is set to 0, and `end` is set to the length of the input array.

2. The sorting process is performed inside a loop that continues as long as swaps are being
made during a pass.
Initially, `swapped` is set to `true` to ensure that the loop executes at least once.
The purpose of this loop is to continue sorting until no more swaps are necessary,
indicating that the array is sorted.

3. In the forward pass, the code iterates through the unsorted portion of the array from
`start` to `end-1`.
It compares adjacent elements and swaps them if they are out of order (i.e., if `array[i]`
is greater than `array[i+1]`).
This pass moves the largest element to the end of the unsorted portion of the array.

4. After the forward pass, if no swaps were made (`swapped` is `false`), it means that the
array is already
sorted, so the loop can be exited early.

5. The `swapped` flag is reset to `false` for the backward pass, and the `end` pointer is
decreased by 1 because the largest element is now in its correct position at the
end of the unsorted portion.

6. In the backward pass, the code iterates through the unsorted portion of the array in
reverse order (from `end-1` down to `start`). It compares adjacent elements and swaps
them if they are out of order.
This pass moves the smallest element to the beginning of the unsorted portion.

7. After the backward pass, the `start` pointer is increased by 1 because the smallest
element is now in
its correct position at the beginning of the unsorted portion.

8. Finally, when the loop exits (indicating that the entire array is sorted),
the function returns the sorted `array`.

9. Elements are compared and swapped in a bidirectional manner until the array is sorted.
It's a variation of Bubble Sort that can perform better in certain cases, particularly when there
are small elements at both ends of the array.

```go
func CockTail[T Numeric](array []T) []T {
	// Initialize variables
	var swapped bool = true  // Flag to check if any swaps are made in a pass
	var start int = 0        // Start of the unsorted portion of the array
	var end int = len(array) // End of the unsorted portion of the array

	// Continue sorting until no swaps are made in a pass
	for swapped {
		swapped = false // Reset the swapped flag at the beginning of each pass

		// Forward pass: Move the largest element to the end
		for i := start; i < end-1; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i] // Swap elements if they are out of order
				swapped = true                              // Mark as swapped
			}
		}

		// If no swaps were made in the forward pass, the array is sorted
		if !swapped {
			break
		}

		swapped = false // Reset the swapped flag for the backward pass
		end = end - 1   // Decrease the end pointer as the largest element is now at the end

		// Backward pass: Move the smallest element to the beginning
		for i := end - 1; i >= start; i-- {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i] // Swap elements if they are out of order
				swapped = true                              // Mark as swapped
			}
		}

		start = start + 1 // Increase the start pointer as the smallest element is now at the beginning
	}

	return array // Return the sorted array
}
```

### Space and Time Complexity

| Complexity          | Worst Case     | Average Case  | Best Case     |
|---------------------|----------------|---------------|---------------|
| Time Complexity     | O(n^2)         | O(n^2)        | O(n)          |
| Space Complexity    | O(1)           | O(1)          | O(1)          |

- **Time Complexity**:
  - Worst Case: In the worst case, when the input array is completely unsorted, Cocktail Sort makes a large number of comparisons and swaps. Its worst-case time complexity is O(n^2), where n is the number of elements in the array.
  - Average Case: On average, the number of comparisons and swaps remains high, resulting in an average time complexity of O(n^2).
  - Best Case: In the best case, when the input array is already sorted, Cocktail Sort still makes some comparisons but no swaps. The best-case time complexity is O(n).

- **Space Complexity**:
  - Cocktail Sort is an in-place sorting algorithm, meaning it doesn't require additional memory proportional to the input size. Regardless of the input data, it uses a constant amount of extra space for variables. Therefore, the space complexity is O(1), indicating it is a space-efficient algorithm.

## Strand Sort

Strand Sort is a relatively simple sorting algorithm that works by repeatedly taking elements from an unsorted list and placing them into a sorted list. It's not a commonly used sorting algorithm in practice due to its inefficiency, but it's an interesting concept to understand.

1. `Strand` function:
   - The `Strand` function is the main sorting function that takes an unsorted array as input
     and returns the sorted array.

   - It first checks if the input slice has one or fewer elements. If it does, the slice
     is already sorted, and it's returned as is.

   - It initializes a `result` slice with the first element of the input array and removes
     this element from the input slice.

   - The function enters a loop to continue processing the input slice. Inside the loop:
     - It initializes a `subList` slice with the first element of the input and removes this
	   element from the input slice.
     - It iterates through the input slice to find elements that can be added to the `subList`.
	   Elements are added to the `subList` if they are greater than or equal to the last
	   element of the `subList`. The added elements are also removed from the input slice.

   - After constructing the `subList`, it's merged with the `result` using the `sMerge`
     helper function.

   - The loop continues until the input slice is empty.

   - The sorted `result` slice is returned as the final output.

2. `sMerge` function:
   - The `sMerge` function is a helper function used to merge two sorted slices,
     `a` and `b,` into a single sorted slice.

   - It initializes a `result` slice to hold the merged values. The capacity of this
     slice is set to accommodate both input slices.

   - It maintains two indices, `i` and `j`, to traverse the `a` and `b` slices, respectively.

   - The function merges the two slices while maintaining the sorted order.
     It compares elements at `a[i]` and `b[j]` and appends the smaller of the two to the `result` slice. The corresponding index is incremented.

   - After merging, any remaining elements from `a` and `b` are appended to the `result`.

   - The merged `result` slice is returned.

```go
func Strand[T Numeric](array []T) []T {
	// If the input slice has 1 or fewer elements, it's already sorted, so return it as is.
	if len(array) <= 1 {
		return array
	}

	// Initialize the result slice with the first element of the input.
	var result []T = []T{array[0]}
	// Remove the first element from the input slice.
	array = array[1:]

	// Continue processing the input slice.
	for len(array) > 0 {
		// Initialize a subList with the first element of the input.
		var subList []T = []T{array[0]}
		// Remove the first element from the input slice.
		array = array[1:]

		// Iterate through the input slice to find elements that can be added to the subList.
		var i int = 0
		for i < len(array) {
			// If the current element is greater than or equal to the last element of the subList, add it to the subList.
			if array[i] >= subList[len(subList)-1] {
				subList = append(subList, array[i])
				// Remove the added element from the input slice.
				array = append(array[:i], array[i+1:]...)
			} else {
				i++
			}
		}

		// Merge the subList into the result using the sMerge function.
		result = sMerge(result, subList)
	}

	return result
}

// sMerge is a helper function to merge two sorted slices 'a' and 'b' into a single sorted slice.
func sMerge[T Numeric](a []T, b []T) []T {
	// Initialize a result slice to hold the merged values.
	var result []T = make([]T, 0, len(a)+len(b))
	var i, j int = 0, 0

	// Merge the two slices while maintaining the sorted order.
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	// Append any remaining elements from slices 'a' and 'b' to the result.
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)

	return result
}
```

### Space and Time Complexity

| Complexity Type     | Time Complexity          | Space Complexity  |
|---------------------|--------------------------|-------------------|
| Worst-Case          | O(n^2)                   | O(n)              |
| Best-Case           | O(n)                     | O(n)              |
| Average-Case        | O(n^2)                   | O(n)              |

**Time Complexity:**

- Worst-Case Time Complexity (O(n^2)): In the worst-case scenario, when the input list is in reverse order, Strand Sort will exhibit its worst performance. This is because it repeatedly creates new strands, which results in nested loops and excessive merging. As a result, the time complexity becomes quadratic, O(n^2), where "n" is the number of elements in the input list.

- Best-Case Time Complexity (O(n)): In the best-case scenario, when the input list is already sorted or nearly sorted, the algorithm performs more efficiently. It can build strands without extensive merging, resulting in a linear time complexity, O(n).

- Average-Case Time Complexity (O(n^2)): The average-case time complexity of Strand Sort is generally closer to the worst-case scenario due to its nature of frequently building and merging strands. This makes it inefficient for most practical sorting tasks.

**Space Complexity:**

- Space Complexity (O(n)): Strand Sort uses space for two lists, the input list, and the output list (sorted list). The space complexity is linear, O(n), as it stores a copy of the input list and the output list, each with "n" elements. Additionally, some temporary storage is needed for strands and merging.

## PanCake Sort

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

```go
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
```

### Time & Space Complexity

| Complexity      | Worst Case  | Best Case | Average Case | Space Complexity |
|-----------------|-------------|-----------|--------------|------------------|
| Time Complexity | O(n^2)      | O(n^2)    | O(n^2)       | O(1)             |

- **Worst Case**: In the worst case, when the input array is in reverse order (the largest element is at the beginning of the unsorted portion in each step), Pancake Sort requires O(n^2) comparisons and flips to sort the array.

- **Best Case**: The best-case time complexity is also O(n^2). Even if the array is already sorted, Pancake Sort will still require a similar number of comparisons and flips to verify that it's sorted.

- **Average Case**: On average, Pancake Sort performs O(n^2) comparisons and flips. The average case time complexity remains the same as the worst and best cases.

- **Space Complexity**: Pancake Sort has a space complexity of O(1) since it sorts the array in-place, without requiring additional memory allocation proportional to the input size. It uses a constant amount of extra space for variables like loop counters and indices.

## Sleep Sort

Nothing More Useless than the Sleep Sort - it does Nothing. Just Fascinating to look at

Sleep Sort is not guaranteed to work accurately due to the asynchronous nature of go-routines and the unpredictability of when they will wake up. It's a playful experiment rather than a reliable sorting algorithm. The results may vary each time you run the code, and it's not suitable for practical use.

```go
import "sync"

// safeAppend safely appends an element of type T to a slice, protected by a mutex.
func safeAppend[T Numeric](arr *[]T, wg *sync.WaitGroup, mutex *sync.Mutex, num T) {
	// Mark the wait group as done when this goroutine exits.
	defer wg.Done()

	// Lock the mutex to ensure exclusive access to the slice.
	mutex.Lock()

	// Append the element to the slice.
	*arr = append(*arr, num)

	// Unlock the mutex to allow other go-routines to access the slice.
	mutex.Unlock()
}

// Sleep is an implementation of the Sleep Sort algorithm that can work with slices of different numeric types.
func Sleep[T Numeric](array []T) []T {
	// Create a wait group to track the completion of go-routines.
	var wg sync.WaitGroup

	// Create a mutex to ensure safe access to the sorted slice.
	var mutex sync.Mutex

	// Initialize the sorted slice with an initial capacity.
	sorted := make([]T, 0, len(array))

	// Iterate through the input array.
	for _, num := range array {
		// Increment the wait group counter to track the number of active go-routines.
		wg.Add(1)

		// Launch a goroutine to safely append the element to the sorted slice.
		go safeAppend(&sorted, &wg, &mutex, num)
	}

	// Wait for all go-routines to finish before proceeding.
	wg.Wait()

	// Return the sorted slice with initially allocated space removed.
	return sorted
}
```

## Insertion Sort

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

```go
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
```

### Time & Space Complexity

| Algorithm         | Worst Case | Average Case | Best Case | Space Complexity |
|-------------------|-------------------------|---------------------------|-----------------------|-------------------|
| **Insertion Sort** | O(n^2)                  | O(n^2)                    | O(n)                  | O(1)              |

- **Time Complexity (Worst)**: In the worst-case scenario, where the input array is in reverse order, the Insertion Sort algorithm exhibits a time complexity of O(n^2), where "n" is the number of elements in the array. This is because it requires many comparisons and shifts to sort the array.

- **Time Complexity (Average)**: On average, when the input data is randomly ordered, the Insertion Sort algorithm also has a time complexity of O(n^2). It involves a similar number of comparisons and shifts as in the worst-case scenario.

- **Time Complexity (Best)**: In the best-case scenario, when the input array is already nearly sorted or completely sorted, the Insertion Sort algorithm performs significantly better. In this case, the time complexity is O(n), where "n" is the number of elements. This is because only comparisons are needed, and there are no or very few shifts.

- **Space Complexity**: The Insertion Sort algorithm has a space complexity of O(1), indicating that it doesn't require additional memory allocation that scales with the size of the input array. It sorts the array in-place by rearranging elements within the array itself.

## Stooge Sort

1. **Initialization:**
     The function begins by initializing two variables, `low` and `high`, to represent the
	 low and high indices of the entire array. Initially, `low` is set to 0, and `high`
	 is set to the length of the input array.

2. **Check and Swap:**
     It checks if the last element in the array (indexed at `high-1`) is smaller than the
	 first element in the array (indexed at `low`). If this condition is true, it swaps
	 these two elements. This ensures that the smallest element is at the beginning of
	 the array, as required for Stooge Sort.

3. **Calculate the Length:**
     It calculates the length of the current sub-array by subtracting `low` from `high`.
	 This represents the number of elements to be sorted in the current call.

4. **Recursion with a Divide-and-Conquer Approach:**
     If the length of the sub-array is greater than 2, the function applies
	 a divide-and-conquer approach:

   - It calculates one-third of the length and stores it in the variable `third`.

   - It then calls the `Stooge` function recursively on three sub-arrays:
     1. The first 2/3 of the sub-array (from `low` to `high-third`).
     2. The last 2/3 of the sub-array (from `low+third` to `high`).
     3. The first 2/3 of the sub-array again (from `low` to `high-third`).

   This recursive approach sorts smaller sub-arrays and, eventually, the entire array.

5. **Return Sorted Sub-Array:**
     When the base case is reached (sub-arrays with 2 or fewer elements), the function
	 returns the sorted sub-array.

In essence, Stooge Sort works by dividing the array into three parts, sorting the first 2/3 and
last 2/3 separately, and then sorting the first 2/3 again.
This process repeats recursively until the sub-arrays have only 2 or fewer elements,
and it ensures that the array is sorted correctly.

```go
func Stooge[T Numeric](array []T) []T {
	// Initialize the low and high indices for the entire array.
	var low int = 0
	var high int = len(array)

	// Check if the last element is smaller than the first element, swap them if needed.
	if array[high-1] < array[low] {
		temp := array[low]
		array[low] = array[high-1]
		array[high-1] = temp
	}

	// Calculate the length of the current sub-array.
	length := high - low

	// Check if there are more than 2 elements in the sub-array.
	if length > 2 {
		// Calculate one-third of the length.
		third := length / 3

		// Recursively sort the first 2/3 of the sub-array.
		Stooge(array[low : high-third])

		// Recursively sort the last 2/3 of the sub-array.
		Stooge(array[low+third : high])

		// Recursively sort the first 2/3 of the sub-array again.
		Stooge(array[low : high-third])
	}

	// Return the sorted sub-array.
	return array
}
```

### Time & Space Complexity

| Complexity Type        | Time Complexity           |
|------------------------|---------------------------|
| Worst-case            | O(n^(log 3 / log 1.5))     |
| Best-case             | O(n^(log 3 / log 1.5))     |
| Average-case          | O(n^(log 3 / log 1.5))     |
| Space Complexity     | O(log n)       |

- **Space Complexity:** The space complexity of Stooge Sort is O(log n). This represents the amount of additional memory used by the algorithm for the recursive function call stack. The space complexity is relatively low and does not depend on the size of the input array.

- **Time Complexity:** The time complexity of Stooge Sort is not efficient. It can be expressed as O(n^(log 3 / log 1.5)), which is roughly equivalent to O(n^2.7095). This makes Stooge Sort impractical for sorting large datasets because it has a high time complexity.

- **Worst-case time complexity:** The worst-case time complexity of Stooge Sort is O(n^(log 3 / log 1.5)), which is roughly equivalent to O(n^2.7095). This means that in the worst-case scenario, when the array is in the reverse order, the algorithm exhibits very poor performance.

- **Best-case time complexity:** The best-case time complexity is also O(n^(log 3 / log 1.5)), which is the same as the worst-case time complexity. This is because Stooge Sort always performs the same number of comparisons and swaps, regardless of the initial order of the array.

- **Average-case time complexity:** The average-case time complexity of Stooge Sort is also O(n^(log 3 / log 1.5)). Like the best-case scenario, the algorithm's performance remains consistently poor on average for various input data distributions.

Stooge Sort has a poor time complexity in all cases, making it one of the least efficient sorting algorithms. It is typically used for educational purposes to demonstrate sorting algorithms' concepts rather than for practical sorting tasks. For real-world applications, more efficient sorting algorithms like Quick Sort, Merge Sort, or Heap Sort are preferred.

## Merge Sort

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

```go
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
```

### Time & Space Complexity

| Complexity Type | Worst Case  | Average Case | Best Case | Space Complexity |
|-----------------|-------------|--------------|-----------|------------------|
| Time Complexity | O(n log n)  | O(n log n)   | O(n log n)| O(n)            |
| Space Complexity| O(n)        | O(n)         | O(n)      | O(n)            |

- **Time Complexity**:
  - Worst Case: O(n log n) - The worst-case time complexity occurs when we have to repeatedly split and merge the array until it's fully sorted.
  - Average Case: O(n log n) - The average-case time complexity is also O(n log n), making Merge Sort efficient on average.
  - Best Case: O(n log n) - Merge Sort doesn't have a best-case scenario where it performs significantly better because it always divides and conquers.

- **Space Complexity**:
  - Space Complexity: O(n) - Merge Sort has a space complexity of O(n) because it requires additional memory for creating temporary arrays during the merge process. This makes it less memory-efficient compared to some other sorting algorithms.

Merge Sort is known for its stability and consistent performance across different input scenarios, but it may not be the best choice when memory usage is a critical concern due to its space complexity.

## Heap Sort

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

```go
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
```

### Space and Time Complexity

| Complexity Type | Time Complexity | Space Complexity |
|-----------------|-----------------|-----------------|
| Best Case       | O(n log n)      | O(1)            |
| Average Case    | O(n log n)      | O(1)            |
| Worst Case      | O(n log n)      | O(1)            |

- **Time Complexity:**
  - In the best, average, and worst cases, Heap Sort has a time complexity of O(n log n). It offers consistent performance regardless of the input data.
- **Space Complexity:**
  - Heap Sort has a space complexity of O(1) in all cases. It sorts the array in place and does not require additional memory that scales with the input size.

## Quick Sort

Quick Sort is a popular sorting algorithm that follows the divide-and-conquer paradigm. It works by selecting a 'pivot' element from the array and partitioning the other elements into two sub-arrays, according to whether they are less than or greater than the pivot. The sub-arrays are then sorted recursively.

1. **Choose a Pivot**: Select a pivot element from the array. There are different strategies for selecting the pivot, such as choosing the first element, the last element, or a random element.

2. **Partitioning**: Reorder the array so that all elements with values less than the pivot come before all elements greater than the pivot. This is often done in a way that the pivot ends up in its correct sorted position. You'll need to create two sub-arrays, one containing elements less than the pivot and the other containing elements greater than the pivot.

3. **Recursion**: Recursively sort the sub-arrays created in the previous step. Continue this process for both the sub-array of elements less than the pivot and the sub-array of elements greater than the pivot.

4. **Combine**: After all recursive calls return, the smaller and greater sub-arrays will be sorted. You can then combine these arrays with the pivot element to get the final sorted array.

5. **Base Case**: Ensure you have a base case or condition that terminates the recursion, such as when the array has only one or zero elements (already sorted).

```go
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
```

### Space & Time Complexity

| Case       | Time Complexity              | Space Complexity |
|------------|------------------------------|-------------------|
| Best       | O(n log n)                   | O(log n)          |
| Average    | O(n log n)                   | O(log n)          |
| Worst      | O(n^2)                       | O(n)              |

- **Best Case:** Occurs when the pivot chosen at each step divides the array into nearly equal halves, leading to a balanced partitioning. In this case, Quick Sort has a time complexity of O(n log n) and a space complexity of O(log n).

- **Average Case:** On average, Quick Sort performs with a time complexity of O(n log n) and a space complexity of O(log n).

- **Worst Case:** Occurs when the pivot is consistently chosen poorly, causing the partitioning to be highly unbalanced. In this case, Quick Sort may have a time complexity of O(n^2), and it uses O(n) space for the call stack.

## Three Way Merge Sort

The `merge` function is responsible for merging the three sub-arrays—left, middle, and right—into a single sorted array. It does so by comparing elements from these sub-arrays and selecting the smallest one at each step, ensuring the overall array is sorted correctly.

The `ThreeWayMerge` function is the entry point of the sorting process. It first checks if the length of the input array is greater than one, and if so, it proceeds with the sorting. The input array is divided into three parts: `left`, `middle`, and `right`. These sub-arrays are created and populated with the corresponding elements from the original array. The `ThreeWayMerge` function then recursively sorts the `left`, `middle`, and `right` sub-arrays.

The sorted sub-arrays are then merged back together into the original array by calling the `merge` function. This process continues until the entire input array is sorted, and the final sorted array is returned. The code uses recursive calls to sort the sub-arrays and an iterative merging process to ensure the elements are arranged in the correct order, resulting in a fully sorted array.

```go
func merge[T Numeric](array []T, left []T, middle []T, right []T) []T {
	var i int = 0
	var j int = 0
	var k int = 0
	//  Loop until we've processed all elements in left, middle, and right
	for i < len(left) || j < len(middle) || k < len(right) {
		/*

			i < len(left) checks if there are remaining elements in the left sub-array to consider for merging.

			(j == len(middle) || left[i] <= middle[j]) checks if either the middle sub-array is fully processed
			(j equals its length) or if the current element in the left sub-array is smaller than or equal to
			the current element in the middle sub- array.

			(k == len(right) || left[i] <= right[k]) checks if either the right sub-array is fully processed
			(k equals its length) or if the current element in the left sub-array is smaller than or equal
			to the current element in the right sub-array.


		*/
		if i < len(left) && (j == len(middle) || left[i] <= middle[j]) && (k == len(right) || left[i] <= right[k]) {
			array[i+j+k] = left[i]
			i++
			/*

				j < len(middle) checks if there are remaining elements in the middle sub-array to consider for merging.

				(k == len(right) || middle[j] <= right[k]) checks if either the right sub-array is fully processed
				(k equals its length) or if the current element in the middle sub-array is smaller than or
				equal to the current element in the right sub-array.


			*/
		} else if j < len(middle) && (k == len(right) || middle[j] <= right[k]) {
			array[i+j+k] = middle[j]
			j++
			/*
				if none of the previous conditions are met, it means that the current element in the right sub-array
				is smaller than or equal to the elements in the left and middle sub-arrays, or both the left
				and middle sub-arrays have been fully

				In this case, the current element from the right sub-array is selected to be placed in
				the final merged array. The element is assigned to array[i+j+k], and then k is incremented
				to move on to the next element in the right sub-array.
			*/
		} else {
			array[i+j+k] = right[k]
			// moving on
			k++
		}
	}
	return array
}

func ThreeWayMerge[T Numeric](array []T) []T {
	if len(array) > 1 {
		var mid1 int = len(array) / 3
		var mid2 int = 2 * len(array) / 3
		// Split the input array into three parts: left, middle, and right
		var left []T = make([]T, mid1)
		copy(left, array[:mid1])

		var middle []T = make([]T, mid2-mid1)
		copy(middle, array[mid1:mid2])

		var right []T = make([]T, len(array)-mid2)
		copy(right, array[mid2:])
		// Recursively sort the left, middle, and right sub-arrays
		ThreeWayMerge(left)
		ThreeWayMerge(middle)
		ThreeWayMerge(right)
		// Merge the sorted left, middle, and right sub-arrays into the original array
		array = merge(array, left, middle, right)
	}
	return array
}
```

### Space and Time Complexity

| **Case**        | **Time Complexity** | **Space Complexity** |
| --------------- | -------------------- | -------------------- |
| **Worst Case**  | O(n * log n)        | O(n)                 |
| **Best Case**   | O(n * log n)        | O(n)                 |
| **Average Case**| O(n * log n)        | O(n)                 |

**Time Complexities:**

1. **Worst Case:** The worst-case time complexity of three-way merge sort is O(n * log n), where 'n' is the number of elements in the array. This occurs when the array is repeatedly divided into three parts, and each part is recursively sorted and merged. The log n factor arises from the depth of the recursive tree, and the 'n' factor comes from the merging step.

2. **Best Case:** The best-case time complexity is also O(n *log n). In three-way merge sort, the number of comparisons and element movements remains the same regardless of the input order. The 'n* log n' time complexity is a fundamental characteristic of merge sort algorithms.

3. **Average Case:** The average-case time complexity is O(n * log n) as well. In practice, three-way merge sort typically performs with this time complexity. The algorithm evenly divides the array into three parts, resulting in a balanced and efficient sorting process.

**Space Complexity:**

The space complexity of the three-way merge sort algorithm is O(n) in the worst and average cases. This space is primarily required for creating temporary sub-arrays during the merge and the recursive call stack. In the best case, the space complexity is still O(n) because merge sort inherently requires additional space for merging. However, in practice, the space overhead is typically considered reasonable and manageable.

## Three Way Quick Sort

1. `ThreeWayQuick`: The main function that initiates the sorting process. It checks if the array is already sorted or contains only one element. If not, it shuffles the array to improve performance and then calls the `sort` function to perform the sorting.

2. `sort`: A recursive function that sorts the array using Three-Way Quick Sort. It calls the `partition` function to separate elements into three partitions (less than, equal to, and greater than the pivot) and then recursively sorts these partitions.

3. `partition`: This function separates elements into three partitions and returns the boundaries of elements that are equal to the pivot. It uses three pointers (lessThan, iter, and greaterThan) to rearrange elements.

4. `shuffle`: A function that randomly shuffles the elements in the array. This helps avoid worst-case scenarios in Three-Way Quick Sort by introducing randomness when choosing the pivot.

```go
// ThreeWayQuick performs Three-Way Quick Sort on the input array.
func ThreeWayQuick[T Numeric](array []T) []T {
	// If the array is already sorted or contains only one element, no sorting is needed.
	if len(array) < 2 {
		return array
	}

	// Shuffle the array to improve performance, mitigating worst-case scenarios.
	shuffle(array)

	// Sort the array using the Three-Way Quick Sort algorithm.
	array = sorter(array, 0, len(array)-1)
	return array
}

// sorter is a recursive function that sorts the input array in place using Three-Way Quick Sort.
func sorter[T Numeric](array []T, low, high int) []T {
	// Base case: If the high index is less than or equal to the low index, the array is sorted.
	if high <= low {
		return array
	}

	// Perform the partitioning step to separate elements into three partitions.
	lessThan, greaterThan := partition(array, low, high)

	// Recursively sort the partitions to the left and right of the pivot.
	sorter(array, low, lessThan-1)
	sorter(array, greaterThan+1, high)

	return array
}

// partition function partitions the array into three segments and returns the boundaries of equal elements.
func partition[T Numeric](array []T, low, high int) (int, int) {
	// Initialize the lessThan, iter, and greaterThan pointers and choose the pivot element.
	lessThan, iter, greaterThan := low, low, high
	pivot := array[low]

	// Iterate through the array and place elements into their respective partitions.
	for iter <= greaterThan {
		if array[iter] < pivot {
			// Swap elements that are less than the pivot with the lessThan section.
			array[lessThan], array[iter] = array[iter], array[lessThan]
			lessThan++
			iter++
		} else if array[iter] > pivot {
			// Swap elements that are greater than the pivot with the greaterThan section.
			array[iter], array[greaterThan] = array[greaterThan], array[iter]
			greaterThan--
		} else {
			// Skip elements that are equal to the pivot.
			iter++
		}
	}

	// Return the boundaries of elements equal to the pivot.
	return lessThan, greaterThan

}
```

### Space and Time Complexity

| Complexity              | Worst Case       | Best Case        | Average Case     |
|-------------------------|------------------|------------------|------------------|
| **Time Complexity**    | O(n^2)           | O(n log n)       | O(n log n)       |
| **Space Complexity**   | O(log n)         | O(log n)         | O(log n)         |

**Explanation**:

- **Time Complexity**:
  - **Worst Case**: The worst-case time complexity occurs when the pivot selection consistently results in highly imbalanced partitions (e.g., when the array is already sorted in increasing or decreasing order). In this case, the partitioning step doesn't effectively divide the array, leading to O(n^2) time complexity.
  - **Best Case**: The best-case time complexity occurs when the pivot selection is consistently effective, leading to balanced partitions. In this scenario, the time complexity is O(n log n), which is similar to the time complexity of other efficient sorting algorithms.
  - **Average Case**: On average, the Three-Way Quick Sort exhibits O(n log n) time complexity when considering randomized pivot selection, which leads to balanced partitions in most cases.

- **Space Complexity**:
  - The space complexity of Three-Way Quick Sort is O(log n) for the stack space used in the recursive calls. This is due to the partitioning and recursive nature of the algorithm. In the best case, the stack depth is logarithmic, and in the worst case, it can approach linear space usage.

Please note that these complexities can vary depending on the implementation and pivot selection strategy used. The provided complexities are based on the commonly used Lomuto partition scheme and randomized pivot selection for optimal performance.

## Counting Sort

Counting Sort is a non-comparative sorting algorithm that works well for a range of input values, such as integers. It's a stable sorting algorithm, which means that it maintains the relative order of equal elements in the sorted output.

1. **Find the range of input**: Determine the range of values in the input data. You need to know the minimum and maximum values to allocate an appropriate amount of memory for the counting array.

2. **Initialize the counting array**: Create a counting array with a size equal to the range of values. Each element in the counting array will store the count of occurrences of a particular value in the input data. Initialize all the counting array elements to zero.

3. **Count occurrences**: Loop through the input data and, for each element, increment the corresponding index in the counting array. This step counts how many times each unique value appears in the input.

4. **Accumulate counts**: Modify the counting array so that each element at index `i` contains the sum of the counts from indices 0 to `i`. This step helps in determining the final position of each element in the sorted output.

5. **Create the output array**: Create an output array with the same size as the input data.

6. **Populate the output array**: Iterate through the input data in reverse order. For each element, look up its value in the counting array to find its correct position in the output array. Place the element in that position and decrement the count in the counting array to account for the placement.

7. **Repeat step 6 for all elements**: Continue this process for all elements in the input data, working from right to left. This ensures stability in the sorting algorithm, preserving the relative order of elements with the same values.

8. **Final output**: After processing all elements, your output array will be sorted.

9. **Time and space complexity**: Counting Sort has a time complexity of O(n + k), where n is the number of elements in the input and k is the range of input values. It is a linear time sorting algorithm and can be very efficient when the range of values is not significantly larger than the number of elements. The space complexity is O(k) for the counting array.

Keep in mind that Counting Sort is most effective when dealing with integers or a limited range of discrete values. It's not suitable for sorting data with a wide or continuous range of values.

```go
func Counting[T Numeric](array []T) []T {
	// If the array has 0 or 1 elements, it's already sorted
	if len(array) <= 1 {
		return array
	}

	// Find the minimum and maximum values in the array
	var min, max T = findMinMax(array)

	// Create a slice to store counts for each unique element
	var counts []T = make([]T, int(max-min+1))

	// Count the occurrences of each element in the input array
	for _, num := range array {
		counts[int(num-min)]++
	}
	var index int = 0
	// Calculate cumulative counts to determine the correct position of each element
	for i := 0; i < len(counts); i++ {
		for counts[i] > 0 {
			// Place the element at the correct sorted position in the original array
			array[index] = T(i) + min
			counts[i]--
			index++
		}
	}

	return array
}
```

### Time and Space Complexity

| Case      | Time Complexity    | Space Complexity   |
|-----------|--------------------|--------------------|
| Worst     | O(n + k)           | O(n + k)           |
| Best      | O(n + k)           | O(n + k)           |
| Average   | O(n + k)           | O(n + k)           |

- "n" represents the number of elements in the input array.
- "k" represents the range of input values (the difference between the maximum and minimum values).
- Counting Sort has a linear time complexity, O(n + k), and a space complexity of O(n + k) for all cases (worst, best, and average).

**Time Complexity:**
Counting Sort has a linear time complexity of O(n + k), where "n" is the number of elements in the input array, and "k" is the range of input values.

- **Best-case time complexity (O(n + k)):** The best-case scenario occurs when the range of input values (k) is small and relatively constant compared to the number of elements (n). In this case, Counting Sort performs exceptionally well, as it makes only one pass through the input array to count the occurrences and another pass to place the elements in their sorted positions. It's faster than many other sorting algorithms, such as comparison-based algorithms like Quick Sort and Merge Sort.

- **Worst-case time complexity (O(n + k)):** The worst-case scenario also has a linear time complexity. This occurs when the range of input values is significantly larger than the number of elements (k >> n). In this case, the counting array becomes very large, and the algorithm's efficiency is reduced. However, it's still faster than many comparison-based sorting algorithms in the worst case.

- **Average-case time complexity (O(n + k)):** On average, Counting Sort is a linear-time algorithm. It efficiently sorts data with a relatively small range of values compared to the number of elements. This makes it a suitable choice for sorting tasks with such characteristics.

**Space Complexity:**
The space complexity of Counting Sort is also O(n + k).

- The space complexity for the input array itself is O(n) because you need to store the original data.
- The space complexity for the counting array is O(k) since it stores the count of occurrences for each unique value within the range.
- The space complexity for the output array is O(n) because it needs to store the sorted data.

The total space complexity is the sum of these components, which results in O(n + k). In practice, the space requirements depend on the range of values (k), and Counting Sort may use more memory for the counting array when dealing with a larger range.

## Redix Sort

1. `getMax(arr)` function:
   - This function finds the maximum value in the input array `arr`. It is used to determine the number of digits in the largest element, which is necessary for Radix Sort.

2. `countingSort(arr, exp)` function:
   - The `countingSort` function is used to sort the elements of the array `arr` based on a specific digit place, represented by the `exp` parameter.
   - It creates an `output` slice of the same length as the input array to store the sorted elements.
   - It creates a `count` slice of size 10 (since we are dealing with base-10 integers) to keep track of the count of elements with each digit.
   - The function first loops through the input array to count the number of elements for each digit at the current place (e.g., one's place, ten's place, etc.).
   - It then updates the `count` slice to determine the correct position for each element in the `output` slice.
   - Finally, it loops through the input array again to place the elements in their sorted order in the `output` slice.

3. `radixSort(arr)` function:
   - The `radixSort` function is the main sorting function that takes an array as input and sorts it using Radix Sort.
   - It first finds the maximum element in the array using the `getMax` function to determine the number of iterations required for the sorting process.
   - The loop iterates from the least significant digit to the most significant digit, each time calling `countingSort` with the appropriate `exp` value.
   - As the loop progresses, the array is sorted incrementally based on the current digit place. After processing all the digit places, the array is fully sorted.

4. `main` function:
   - In the `main` function, you can see an example array `arr` that you want to sort.
   - It calls the `radixSort` function to sort the array in place.

The key idea behind Radix Sort is to sort the elements by processing individual digits (or character positions) from right to left, and this process is repeated for all the digit places, starting from the least significant digit and moving to the most significant digit. This algorithm is efficient for sorting non-negative integers with a bounded number of digits and can be adapted for other data types as well.

```go
func Redix[T Numeric](array []T) []T {
	var max T = findMax(array) // Find the maximum value in the array

	// Iterate through each digit place, from least significant to most significant
	for exp := 1; max/T(exp) > 0; exp *= 10 {
		array = countingSort(array, T(exp)) // Call countingSort for the current digit place
	}
	return array
}

func countingSort[T Numeric](array []T, exp T) []T {
	var n int = len(array)
	var output []T = make([]T, n) // Create an output slice of the same size as the input
	var count []T = make([]T, 19) // Create a count array with 19 elements for values in the range [-9, 9]

	// Count the occurrences of each digit at the current digit place
	for i := 0; i < n; i++ {
		count[int(array[i]/exp)%10+9]++
	}

	// Calculate the cumulative count for each digit
	for i := 1; i < 19; i++ {
		count[i] += count[i-1]
	}

	// Build the output array by placing elements in the correct order
	for i := n - 1; i >= 0; i-- {
		output[int(count[int(array[i]/exp)%10+9])-1] = array[i]
		count[int(array[i]/exp)%10+9]--
	}

	// Copy the sorted values from the output array back to the input array
	for i := 0; i < n; i++ {
		array[i] = output[i]
	}
	return output
}
```

### Space and Time Complexity

| Case      | Time Complexity | Space Complexity |
|-----------|-----------------|------------------|
| Worst     | O(k * n)        | O(n + k)         |
| Best      | O(k * n)        | O(n + k)         |
| Average   | O(k * n)        | O(n + k)         |

`n` represents the number of elements in the input array, and `k` is the number of digits or characters in the largest element. The time complexity remains the same for all cases because Radix Sort is not affected by the order of elements, making it efficient in all situations. The space complexity includes the space required for the `output` and `count` arrays, both of which are proportional to `n`.

- **Worst Case Time Complexity (O(k * n)):** In the worst case, Radix Sort will have to perform counting sort for each digit place (from the least significant to the most significant) for all `n` elements in the input array. This results in a time complexity of O(k * n), where `k` is the number of digits or characters in the largest element.

- **Best Case Time Complexity (O(k * n)):** Radix Sort's best case is essentially the same as the worst case. Even if the input is nearly sorted, the algorithm will still go through all the digit places, making it relatively consistent.

- **Average Case Time Complexity (O(k * n)):** The average case time complexity is also O(k * n). Radix Sort is not influenced by the order of elements, and it processes each digit place independently.

- **Space Complexity (O(n + k)):** The space complexity for Radix Sort includes the space required for two additional arrays: `output` and `count`. Both of these arrays have sizes proportional to the number of elements in the input array (`n`). Hence, the space complexity is O(n + k), where `n` is the number of elements and `k` is the number of digits or characters in the largest element.

It's important to note that Radix Sort is a stable sorting algorithm, meaning it preserves the relative order of equal elements. Despite its linear time complexity, it is not commonly used for general sorting tasks like quicksort or merge sort. Radix Sort is most suitable for sorting non-negative integers with a bounded number of digits or characters, where its performance is consistently efficient.

## Bucket Sort

Bucket Sort is a sorting algorithm that divides the input data into a finite number of "buckets" or bins, and then each bucket is sorted individually, either using another sorting algorithm or recursively applying the bucket sort. After all the individual buckets are sorted, their contents are concatenated to obtain the final sorted array. Here's a step-by-step guide on how to implement Bucket Sort without providing actual code:

1. **Determine the Range**: To use Bucket Sort effectively, you need to know the range of the input data. This means finding the minimum and maximum values in the input array.

2. **Create Buckets**: Divide the range into a fixed number of buckets. The number of buckets is typically determined by the range of values and the available memory. Each bucket should cover a specific range of values within the overall range.

3. **Distribute Elements into Buckets**: Iterate through the input array, placing each element into its corresponding bucket based on the element's value. This is done by using a simple mapping function, dividing the range of values evenly between the buckets. Elements equal to the minimum value go into the first bucket, elements equal to the maximum value go into the last bucket, and the rest are distributed based on their relative position within the range.

4. **Sort Each Bucket**: For each bucket, sort the elements within that bucket. You can use any sorting algorithm for this step, such as insertion sort, quicksort, or another instance of bucket sort (if you want to recursively apply bucket sort).

5. **Concatenate Buckets**: After all buckets are sorted, concatenate the contents of each bucket in order to create the final sorted array. The elements from the first bucket come first, followed by the elements in the second bucket, and so on.

6. **Return Sorted Array**: The concatenated array is now sorted, and you can return it as the result.

7. **Time Complexity**: The time complexity of bucket sort depends on the sorting algorithm used to sort the individual buckets. The best-case time complexity is O(n) when the input is uniformly distributed across the buckets, and the worst-case time complexity is O(n^2) if all elements are placed in a single bucket. However, on average, it performs well for uniformly distributed data.

8. **Space Complexity**: The space complexity of bucket sort depends on the number of buckets you use. It's typically O(n) for most practical cases.

Keep in mind that the choice of the sorting algorithm for individual buckets can impact the efficiency and stability of the algorithm. Also, if you have a small range of values, you may want to use fewer buckets to avoid wasting memory.

```go
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
```

### Space and Time Complexity

| Complexity       | Best Case      | Average Case   | Worst Case     | Space Complexity |
|------------------|----------------|----------------|----------------|------------------|
| Time Complexity  | O(n+k) (with k as the number of buckets) | O(n^2) | O(n^2) | O(n+k) (for the buckets) |
|                  |               | O(n^2) when all elements end up in a single bucket | O(n^2) when all elements end up in a single bucket |                  |
|                  |               | (assuming a simple sort is used for buckets) | (assuming a simple sort is used for buckets) |                  |

Time complexity in the best case assumes that the elements are uniformly distributed across the buckets, resulting in a linear time complexity. The worst and average cases occur when all elements end up in a single bucket, leading to a quadratic time complexity. The space complexity depends on the number of buckets and the size of the data, resulting in O(n+k) where n is the number of elements and k is the number of buckets.

## Bitonic Sort

Bitonic sort is a comparison-based sorting algorithm that's particularly well-suited for parallel processing. It works by repeatedly dividing the input sequence into two sub-sequences, then sorting those sub-sequences, and finally merging them back together

1. **Understanding Bitonic Sequences**:
   - A bitonic sequence is one that starts off strictly increasing and then becomes strictly decreasing or vice versa.
   - The basic idea of bitonic sort is to divide the input sequence into bitonic sequences, sort them, and then merge them into a single sorted sequence.

2. **Phase 1: Creating Bitonic Sequences**:
   - Split the input sequence into two equal halves.
   - Recursively apply bitonic sort to each half, ensuring that they are in the proper bitonic order.

3. **Phase 2: Sorting Bitonic Sequences**:
   - During this phase, each subsequence is sorted in a specific order, which alternates for each level of recursion.
   - In the first level of recursion, each sequence is sorted in ascending order.
   - In the subsequent levels, you sort sequences in descending order.

4. **Phase 3: Merging Bitonic Sequences**:
   - Merge the two sorted bitonic sequences into a single larger bitonic sequence.
   - This merging operation is performed repeatedly, with the bitonic sequences becoming larger in each iteration until the entire sequence is sorted.

5. **Phase 4: Final Merge**:
   - Merge the remaining two sequences into a single sorted sequence.
   - The final merged sequence is now sorted in ascending order.

6. **Completing the Sort**:
   - The sequence is now fully sorted in ascending order, and the bitonic sort process is complete.

Bitonic sort is most efficient when used with parallel processing or on hardware designed for parallelism, as it inherently has the potential for parallel execution during the sorting and merging phases.

```go
func Bitonic[T Numeric](arr []T, order bool) []T {
	return bitonicSort(arr, 0, len(arr), order)
}

func bitonicSort[T Numeric](array []T, low, high int, order bool) []T {
	if high <= 1 {
		return array
	}

	mid := high / 2

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		bitonicSort(array, low, mid, !order)
		wg.Done()
	}()

	go func() {
		bitonicSort(array, low+mid, high-mid, order)
		wg.Done()
	}()

	wg.Wait()

	return bitonicMerge(array, low, high, order)

}

func bitonicMerge[T Numeric](array []T, low, high int, order bool) []T {

	if high <= 1 {
		return array
	}

	mid := greatestPowerOfTwoLessThan(high)

	var wg sync.WaitGroup
	wg.Add(2)

	for index := low; index < low+high-mid; index++ {
		if order == (array[index] > array[index+mid]) {
			array[index], array[index+mid] = array[index+mid], array[index]
		}
	}

	go func() {
		bitonicMerge(array, low, mid, order)
		wg.Done()
	}()

	go func() {
		bitonicMerge(array, low+mid, high-mid, order)
		wg.Done()
	}()
	wg.Wait()
	return array
}

func greatestPowerOfTwoLessThan(high int) int {
	k := 1
	for k > 0 && k < high {
		k = k << 1
	}
	return k >> 1
}
```

### Space and Time Complexity

| Case       | Time Complexity        | Space Complexity |
|------------|------------------------|-------------------|
| Worst      | O(n * log^2(n))        | O(n)              |
| Average    | O(n * log^2(n))        | O(n)              |
| Best       | O(n * log^2(n))        | O(n)              |

- **Time Complexity**:
  - The worst, average, and best time complexities for Bitonic Sort are all O(n * log^2(n)).
  - The reason for this is that each level of the recursive division and merging process involves O(log^2(n)) operations. The number of levels is also O(log(n)), resulting in a total time complexity of O(n * log^2(n)) for all cases.

- **Space Complexity**:
  - The space complexity for Bitonic Sort is O(n) for all cases.
  - This is because the algorithm typically sorts the input array in-place without requiring additional memory for data structures like extra arrays. However, if additional arrays or buffers are used for parallelization, the space complexity may be higher in practice.

Bitonic Sort is not typically chosen for its practical efficiency, but for its suitability for parallel processing. The O(n * log^2(n)) time complexity is less efficient compared to other sorting algorithms e.g Merge & Quick Sort

## Tim Sort

TimSort is a sorting algorithm that is a hybrid of merge sort and insertion sort algorithm. It is a stable algorithm and works on real-time data. In this, the list that needs to be sorted is first analyzed, and based on that the best approach is selected.

1.Divide the array into blocks known as run
2.The size of a run can either be 32 or 64
3.Sort the elements of every run using insertion sort
4.Merge the sorted runs using the merge sort algorithm
5.Double the size of the merged array after every iteration

Some possible optimizations:

1. Use binary insertion sort for better performance on small arrays
2. Unroll the insertion sort inner loop for performance
3. Use galloping search for faster merging
4. Eliminate the tmp array by merging directly back to the source

**Calculation of minRun:**

We know now that the first step of the TimSort algorithm is to divide the blocks into runs. minRun is the minimum length of each run. It is calculated from the number of elements of the array N.

- It should not be very long as we will implement insertion sort to sort each run and we know that insertion sort works more efficiently for shorter arrays
- However, it should also not be too short as the next step is to merge these runs, shorter runs will result in more number of runs
- It is beneficial if N/minRun is a power of 2 as it will result in the best performance by merge sort.

**Splitting and sorting of runs:**

When we reach this step, we already have two values – N and minRun

- A descending flag is by the comparison between the first 2 items, if there is only 1 item left, then it is set as false.
- Then the other elements are iterated and checked whether they are still in ascending or “strict” descending order until an item that does not follow this order is found.
- After this, we will have a run in either ascending or “strict” descending order. If it is in a “strict” descending order we need to reverse it.
- Then we check to make sure that the length of the run satisfies minRun. If it is smaller, we compensate for the following items to make it achieve the minRun size.

**Merging of runs:**

- At first, we will create a temporary run having the size of the smallest of the merged runs.
- After this, we need to copy the shortest run into the temporary one.
- Now, we will mark the first element of the large and temporary array is as the current position.
- In each following step, we will compare the elements in the large and temporary arrays, and the smaller ones will be moved into a new sorted array.
- We need to repeat the above step until one of the arrays runs out.
- Lastly, we will add All the remaining elements to the end of the new one.

```go
func Tim[T Numeric](array []T) []T {
	var n int = len(array)
	var minRun int = findMinRun(n)

	// Sort individual sub-arrays of size minRun
	for start := 0; start < n; start += minRun {
		end := start + minRun
		if end > n {
			end = n
		}
		// We have already Implemented Insertion Sort Long Time ago
		Insertion(array[start:end])
	}

	// Now merge the sorted runs
	var size int = minRun
	for size < n {
		// Choose merges sized sz, sz*2, sz*4, until sz >= n
		for start := 0; start < n; start += 2 * size {
			mid := start + size
			end := min(start+2*size, n)
			merging(array[start:mid], array[mid:end])
		}
		size *= 2
	}
	return array
}

func findMinRun(n int) int {
	const minRun int = 32
	r := 0
	for n >= minRun {
		r |= n & 1
		n >>= 1
	}
	return n + r
}

func min[T Numeric](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

func merging[T Numeric](left []T, right []T) []T {
	// Merge two sorted sub-arrays into one
	var result []T = make([]T, len(left)+len(right))
	leftIndex, rightIndex, resultIndex := 0, 0, 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] <= right[rightIndex] {
			result[resultIndex] = left[leftIndex]
			leftIndex++
		} else {
			result[resultIndex] = right[rightIndex]
			rightIndex++
		}
		resultIndex++
	}

	// Copy remaining elements from either a or b
	for leftIndex < len(left) {
		result[resultIndex] = left[leftIndex]
		leftIndex++
		resultIndex++
	}
	for rightIndex < len(right) {
		result[resultIndex] = right[rightIndex]
		rightIndex++
		resultIndex++
	}

	// Copy tmp back to a and b
	copy(left, result)
	copy(right, result[len(left):])
	return result
}
```

### Space and Time Complexity

| Complexity       | Time Complexity                  | Space Complexity |
|------------------|----------------------------------|------------------|
| **Worst Case**   | O(n * log(n))                    | O(n)             |
| **Best Case**    | O(n)                             | O(n)             |
| **Average Case** | O(n * log(n))                    | O(n)             |

- **Time Complexity**:
  - In the worst and average cases, Tim Sort has a time complexity of O(n * log(n)), which makes it efficient for sorting large datasets.
  - In the best case, when the data is already partially ordered, the time complexity can reduce to O(n), which is highly efficient.
  
- **Space Complexity**:
  - Tim Sort has a space complexity of O(n) as it requires additional memory for storing temporary data structures during the sorting and merging processes.
