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

Here's a table comparing various sorting algorithms based on their complexity in markdown format:

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

## 1. Bingo Sort

First find the Min value inside the array, First Loop through the array from 0 which is starting out index and loop till the
end of unsorted Array then that value will be our min, basically we are sorting based on index so we will called it minIndex
Now inside that loop we will write another loop that iterate through the array again but this time it will be the second value of the array then we will compare second value of the array with the minIndex value if it's smaller than we set that second value at minIndex, after that we will swap the elements, like swapping the position at each and every index

1. **Initialize MinIndex:** Start by initializing a variable called `minIndex` to 0. This will represent the index of the minimum value in the array.

2. **Outer Loop:** Begin an outer loop that will iterate through the array from index 0 (the start of the unsorted portion of the array) to the end of the unsorted portion. This loop will determine the `minIndex` and the element to be swapped.

3. **Inner Loop:** Inside the outer loop, start an inner loop. This loop will iterate through the array again but starting from the second element (index + 1). The goal here is to compare each element with the current `minIndex` value and potentially update `minIndex`.

4. **Compare and Update `minIndex`:** In the inner loop, compare the element at the current index with the element at `minIndex`. If the element is smaller than the element at `minIndex`, update `minIndex` to the current index.

5. **Swap Elements:** After finding the `minIndex`, exit the inner loop. Swap the element at the current outer loop index with the element at `minIndex`. This places the smallest unsorted element in its correct position.

6. **Repeat:** Continue the outer loop, incrementing the index by 1 each time. This will find the next smallest element and place it in the next position in the sorted portion of the array.

7. **Finish:** Once the outer loop has iterated through the entire array, the array will be sorted in ascending order. All elements will be in their correct positions.

This sorting algorithm effectively finds the minimum value in the array based on index positions and swaps elements to sort the array.

```go
func Bingo(arr []int) []int {
	n := len(arr)

	for i := 0; i < n; i++ {
		// Find the index of the minimum unsorted element
		minIndex := findMinUnsorted(arr, i, n)

		// Swap the minimum unsorted element with the current element (i)
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return arr
}

// Helper function to find the index of the minimum unsorted element in the array
func findMinUnsorted(arr []int, startIndex, endIndex int) int {
	minIndex := startIndex

	for i := startIndex + 1; i < endIndex; i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}

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

```go
func Bogo(array []int) []int {
	// while until not sorted we will keep sorting until we run run out of enough life on tis earth
	// to sustain humanity
	for !isSorted(array) {
		// shuffle the array if not sorted
		shuffle(array)
	}
	return array
}

func isSorted(array []int) bool {
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
func shuffle(arr []int) {
	// Initialize the random number generator with a seed for consistency.
	rand.New(rand.NewSource(99))

	// Iterate through the array and swap elements randomly to shuffle.
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
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

1. **Initialize**: Start with an unsorted array of elements that you want to sort. This array can be of any data type, but for simplicity, let's assume you're sorting an integer slice.

2. **Outer Loop**: Use an outer loop that iterates from the beginning of the array to the second-to-last element. You can use a loop variable `i` to keep track of the current index.

3. **Find Minimum**: Inside the outer loop, initialize a variable to store the index of the minimum element. This index starts as `i` because initially, the minimum element is assumed to be at the `i`-th position.

4. **Inner Loop**: Start another loop using a separate variable (e.g., `j`) that iterates from `i+1` to the end of the array. This inner loop is used to find the index of the minimum element in the unsorted part of the array.

5. **Compare**: In the inner loop, compare the element at index `j` with the element at the index stored in the "minimum index" variable. If the element at index `j` is smaller, update the "minimum index" to `j`.

6. **Swap (if necessary)**: After the inner loop, if the "minimum index" is not equal to `i`, it means a smaller element has been found in the unsorted part of the array. Swap the element at index `i` with the element at the "minimum index."

7. **Repeat**: Continue this process by incrementing the outer loop variable `i` and repeating steps 3 to 6. The outer loop will continue until `i` reaches the second-to-last element of the array.

8. **Result**: Once the outer loop completes, your array is sorted in ascending order.

```go
func Selection(array []int) []int {
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

1. **Find the Min and Max Values**

- Begin by finding the minimum and maximum values within the unsorted array. You can do this by iterating through the array and keeping track of the minimum and maximum values.

2. **Calculate the Range of Possible Values**

- Calculate the range of possible values by subtracting the minimum value found in step 1 from the maximum value. This range represents the span of values that can be present in the input array.

3. **Create an Array of Pigeonholes (Buckets)**

- Create an array of pigeonholes (buckets) to accommodate the values within the range of possible values. The number of pigeonholes should be at least equal to the range of values calculated in step 2.

4. **Distribute Values into Pigeonholes**

- Iterate through the unsorted array, and for each element:
  - Calculate its position within the pigeonholes array based on its value within the range. This can be done by subtracting the minimum value from the element's value.
  - Place the element into the corresponding pigeonhole based on this position. If multiple elements have the same value, you can use a data structure (e.g., a linked list) within each pigeonhole to maintain their order.

5. **Collection Phase (Sorting)**

- Begin the collection phase by iterating through the pigeonholes array in order, from the smallest possible value to the largest.
- For each non-empty pigeonhole:
  - Collect the elements back into the original array, maintaining their order.
  - Continue collecting elements from the pigeonhole until it is empty.
- Repeat this process for all pigeonholes in ascending order.

6. **Final Result**

- After completing the collection phase, the original array will be sorted in ascending order.

The Pigeonhole Sort algorithm sorts elements by distributing them into buckets (pigeonholes) based on their values within a known range, and then collecting them back in ascending order. The critical steps involve calculating the range, creating pigeonholes, distributing values, and carefully collecting elements from the pigeonholes to achieve the desired sorting result.

```go
func Pigeon(array []int) ([]int, error) {
	if len(array) == 0 {
		return nil, errors.New("The Provided Array is Empty")
	}
	// Finding min and max with the help of helper function
	minValue, maxValue, minMaxErr := findMinMax(array)
	if minMaxErr != nil {
		return nil, minMaxErr
	}
	// Range of possible value - because This determine the Amount of bucket needed for the Collection
	var rangeOfValue int = maxValue - minValue + 1
	// we will create pigeon Holes based on range of value for possible collection of the Values
	pigeonHolesBucket := make([]int, rangeOfValue)
	// Iterate through the values of the original array
	for _, num := range array {
		// adding those values into the relative position
		// num-minValue := calculates the index of the pigeonhole where the current element should be placed.
		// and by incrementing ++ we will effectively place the each and every value inside the bucket
		pigeonHolesBucket[num-minValue]++
	}

	/// This part is collecting the element from the Bucket and sorting
	//  it and placing it back to original array as Sorted

	// index is to keep track index at value of the Original Array
	var index int = 0
	// this loop iterate through the range of Values inside the Pigeon Hole in Ascending Order (small to Large)
	for i := 0; i < rangeOfValue; i++ {
		// It collect the smallest element from the bucket till pigeon hole is empty
		for pigeonHolesBucket[i] > 0 {
			// array[index] := is the starting position of the Original array
			// i := is the index of the position of the current pigeon Hole Bucket
			// minValue := minimum value in the Original array
			// By adding (i + minValue) minVal to i, we "de-normalize" the value, converting
			// it back to its original value within the range.
			array[index] = i + minValue
			// Increment the Index for preparing the next element
			index++
			// Decrement the values from the bucket to original array till it's empty
			pigeonHolesBucket[i]--
		}
	}

	return array, nil
}

// Finding the min and max value inside the Unsorted Array
func findMinMax(array []int) (int, int, error) {
	// base case - if our array has nothing that It will return Nothing
	if len(array) == 0 {
		return 0, 0, errors.New("The Array is Empty: Nothing Found")
	}

	// Lets Assume that that the First element in the array is the min and also max

	var maxValue int = array[0]
	var minValue int = array[0]
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

	return minValue, maxValue, nil
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

1. **Loop Through the Array**: Begin by setting up a loop to iterate through the unsorted array. This loop will help you process each element in the array.

2. **Select the Current Item**: For each iteration, select the first element in the array as the 'item' that you want to place in its correct sorted position. This element will also serve as your starting point.

3. **Find the Correct Position**: Initialize a 'position' variable to the current starting position (which is the index of the 'item' within the array). Then, loop through the remaining elements in the array, starting from the second element. Compare each element with the 'item' to determine the correct position for 'item.' If an element is smaller than 'item,' increment the 'position' variable.

4. **Skip If Already Sorted**: Check if 'item' is already in its correct sorted position by comparing 'position' with the initial starting position. If they are the same, it means 'item' is already in the right place, so you can move on to the next item.

5. **Find the Next Available Position**: To avoid overwriting existing elements in the array, find the next available position for 'item.' This involves looping through the array until you find a position where 'item' can be placed without overwriting any values.

6. **Swap the Elements**: Once you've identified the correct position and the next available position for 'item,' perform a swap. Swap 'item' with the element at the 'position.' This step effectively moves 'item' to its correct sorted position.

7. **Continue Swapping**: To ensure that 'item' reaches its final sorted position, continue swapping elements until 'item' cycles back to its initial starting position within the array. This step may involve multiple swaps.

```go
func Cycle(array []int) []int {
    // Loop through the unsorted array
    for initialPosition := 0; initialPosition < len(array)-1; initialPosition++ {
        // Select the first element as the 'item' to be placed in its correct position
        item := array[initialPosition]
        // Initialize 'position' as the current starting position for 'item'
        position := initialPosition

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

 1.**Initialization**

- Create an array of elements that you want to sort.

 2. **Swapped Flag**

- Initialize a boolean flag, 'swapped,' as 'true.' This flag will help us track whether any element swaps occurred during a pass.

 3. **Main Loop**

- Enter a loop that continues until 'swapped' is 'false' or the array length is reduced to one. You cannot swap a single-element array.

4. **Calculate Gap**

- Calculate the 'gap' value for this iteration using the 'getGap' function. The gap dynamically shrinks during each pass.

 5. **Reset Swapped Flag**

- Set the 'swapped' flag to 'false' at the start of each pass. We assume no swaps have occurred initially.

 6. **Iterate Through Array**

- Loop through the array, comparing and potentially swapping elements that are 'gap' positions apart.

 7. **Compare and Swap**

- If the element at the current 'index' is greater than the element at 'index+gap,' swap these elements.
- This step involves moving elements toward their correct positions in the sorted order.

 8. **Track Swaps**

- Set the 'swapped' flag to 'true' if any swapping occurred during this pass.

 9. **Gap Adjustment**

- Repeat the process with a smaller gap in the next iteration.
- The 'getGap' function calculates the new gap size, adapting it for better efficiency.

 10. **Repeat Loop**

- Continue the main loop until 'swapped' is 'false' or the array length is reduced to one.

 11. **Sorting Completion**

- After the loop finishes, the array is sorted.

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
func Comb(array []int) {
	// Get the length of the input array
	n := len(array)

	// Set the initial gap to the length of the array
	gap := n

	// Define the shrink factor for adjusting the gap
	shrinkFactor := 1.3

	// Initialize a flag to track whether any swaps occurred
	swapped := true

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
}
```

| Complexity         | Worst Case                              | Average Case                         | Best Case                                | Space Complexity |
|--------------------|----------------------------------------|-------------------------------------|------------------------------------------|------------------|
| Time Complexity    | O(n^2)                                   | O(n^2)                                | O(n*log(n)) [Optimal Gap Sequence]        | O(1)             |
| Explanation        | In the worst case, when the gap is set to 1 and the array is in reverse order or nearly so, Comb Sort becomes inefficient and takes quadratic time. | In most practical cases, Comb Sort's adaptive gap size reduction provides better performance than Bubble Sort but still exhibits quadratic time complexity. | In the best case, when an optimal gap sequence is used and data is partially sorted, Comb Sort approaches a time complexity of n*log(n). Achieving this best-case performance is challenging in practice. | Comb Sort is an in-place sorting algorithm, which means it uses a constant amount of extra memory for variables like the gap and loop indices. |

## Bubble Sort

1. Initialize `n` with the length of the input slice `arr`. `n` represents the number of elements in the slice.

2. Initialize a boolean variable `swapped` to `true`. This variable will be used to check whether any swaps were made in a pass.

3. Enter a `for` loop that continues as long as `swapped` is `true`. This loop ensures that the sorting process continues until no more swaps are made in a pass.

4. Within the loop, set `swapped` to `false`. This flag is initially true and will only be set to false if no swaps are made in the current pass.

5. Start another `for` loop that iterates through the slice from index 1 to `n-1`. We compare adjacent elements and swap them if they are in the wrong order.

6. If the comparison (`arr[i-1] > arr[i]`) is true, swap the elements using a parallel assignment. This is where the actual swapping of elements occurs.

7. Set `swapped` to `true` to indicate that a swap was made in this pass.

8. After the inner loop completes, one pass through the array is done. Now, decrement `n` to reduce the range of elements to consider in the next pass.

```go
func Bubble(array []int) []int {
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

func RecursiveBubble(array []int) []int {
	n := len(array)
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

1. **Function Purpose**: The `Gnome` function is designed to sort an integer array using the Gnome Sort algorithm.

2. **Initialization**:
   - `index`: We start by initializing a variable called `index` to 0. This variable represents our current position within the array.
   - `n`: We find the length of the input array and store it in a variable called `n`.

3. **Looping Through the Array**:
   - The function enters a `for` loop that runs as long as `index` is less than `n`. This loop is the core of the Gnome Sort algorithm.

4. **Checking Conditions**:
   - Inside the loop, we check two conditions:
     - `index == 0`: This condition checks if the current position is at the beginning of the array. If it is, we are already at the starting point.
     - `array[index] >= array[index-1]`: This condition checks if the current element is greater than or equal to the previous element in the array. If this condition is true, it means the elements are in the correct order.

5. **Moving Forward**:
   - If either of the conditions mentioned in the previous step is met, we increment `index` by 1. This effectively moves us to the next position in the array.

6. **Swapping Elements**:
   - If neither of the conditions is met, it implies that the current element is out of order with respect to the previous element. In this case, we perform a swap operation:
     - `array[index]` is swapped with `array[index-1]`, which effectively moves the smaller element one position backward in the array.
     - After the swap, we decrement `index` by 1 to continue checking the previous pair of elements.

7. **Loop Continuation**:
   - The loop continues to iterate until `index` reaches or exceeds the length of the array (`n`). When this happens, it indicates that the entire array is sorted.

8. **Sorted Array Return**:
   - Finally, the function returns the sorted array. The original array is now sorted in place.

```go
func Gnome(array []int) []int {
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

The algorithm operates as follows:

1. In each pass, it compares and swaps elements at even-indexed positions (0, 2, 4, etc.) with their adjacent elements at odd-indexed positions (1, 3, 5, etc.).

2. Then, in the next pass, it reverses the process, comparing and swapping elements at odd-indexed positions with their adjacent elements at even-indexed positions.

3. This process continues until no more swaps are needed in an entire pass, indicating that the array is sorted.

The terms "odd" and "even" in Odd-Even Sort do not refer to the values of the elements but rather to the positions or indices of the elements within the array.

Implementation:

1. It starts by assuming that the array is sorted (`sorted` is initially set to `true`).

2. It enters a loop that continues until the array is sorted. The loop will exit when no swaps are made in an entire pass through both odd and even indices, indicating that the array is sorted.

3. In each pass through the loop, it performs two sorting steps:

   - Step 1: It compares and swaps adjacent elements at odd indices (`oddIndex`) in the array.

   - Step 2: It compares and swaps adjacent elements at even indices (`evenIndex`) in the array.

4. If any swap is made during either of the sorting steps, `sorted` is set to `false` to indicate that the array is not fully sorted.

5. The loop continues until no swaps are made during an entire pass, at which point it exits, and the sorted array is returned.

```go
func OddEven(array []int) []int {
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

1. Initialize variables:
   - `swapped` is a flag to keep track of whether any swaps were made in a pass.
   - `start` represents the starting index of the unsorted portion of the array.
   - `end` represents the last index of the unsorted portion of the array.

2. Enter a loop that continues until no swaps are made in a pass (indicating that the array is sorted).

3. Within the loop:
   - Perform a forward pass:
     - Iterate through the array from `start` to `end-1`.
     - Compare adjacent elements and swap them if they are out of order.
     - Set `swapped` to `true` if any swaps are made.
     - If no swaps are made in this pass, break out of the loop as the array is sorted.

4. Reset the `swapped` flag to `false` and decrement `end` by 1 since the largest element is now at the end.

5. Perform a backward pass: left to right
   - Iterate through the array from `end-1` to `start`.
   - Compare adjacent elements and swap them if they are out of order.
   - Set `swapped` to `true` if any swaps are made.

6. Increment `start` by 1 since the smallest element is now at the beginning.

7. Repeat steps 3-6 until no swaps are made in a pass, indicating that the array is fully sorted.(right to left)

8. Finally, return the sorted array.

```go
func CockTail(array []int) []int {
	// Initialize variables
	var swapped bool = true         // Flag to check if any swaps are made in a pass
	var start int = 0              // Start of the unsorted portion of the array
	var end int = len(array)       // End of the unsorted portion of the array

	// Continue sorting until no swaps are made in a pass
	for swapped {
		swapped = false // Reset the swapped flag at the beginning of each pass

		// Forward pass: Move the largest element to the end (left to right)
		for i := start; i < end-1; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i] // Swap elements if they are out of order
				swapped = true // Mark as swapped
			}
		}

		// If no swaps were made in the forward pass, the array is sorted
		if !swapped {
			break
		}

		swapped = false // Reset the swapped flag for the backward pass
		end = end - 1   // Decrease the end pointer as the largest element is now at the end

		// Backward pass: Move the smallest element to the beginning (right to left)
		for i := end - 1; i >= start; i-- {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i] // Swap elements if they are out of order
				swapped = true // Mark as swapped
			}
		}

		start = start + 1 // Increase the start pointer as the smallest element is now at the beginning
	}

	return array // Return the sorted array
}
```

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

Here's a step-by-step guide to how Strand Sort works without providing code:

1. Initialize two lists: an input list (the unsorted list) and an output list (the sorted list). Initially, the output list is empty.

2. Find the smallest element in the input list and remove it. Place this element into the output list. This element is considered sorted.

3. Iterate through the remaining elements in the input list, and for each element, check if it is smaller than the last element in the output list (the element you just added). If it is, remove it from the input list and add it to the end of the output list.

4. Repeat step 3 until you can't find any more elements to add to the output list. This process is like building a strand of elements within the output list.

5. Once you can't find any more elements to add, combine the output list with the previously sorted list. The output list becomes the new sorted list, and the input list is now the unsorted list.

6. Repeat steps 2 through 5 until the input list is empty. This process effectively keeps building sorted strands of elements and combining them until the entire list is sorted.

7. The algorithm terminates when the input list is empty, and the output list contains all the sorted elements.

```go
func Strand(arr []int) []int {
	// Initialize an empty list to store the sorted elements.
	sorted := []int{}

	// Continue sorting until the input list is empty.
	for len(arr) > 0 {
		// Create a strand with the first element of the input list.
		strand := []int{arr[0]}
		arr = arr[1:]

		// Iterate through the remaining elements in the input list.
		for i := 0; i < len(arr); i++ {
			// If the current element is greater than or equal to the last element in the strand,
			// add it to the strand and remove it from the input list.
			if arr[i] >= strand[len(strand)-1] {
				strand = append(strand, arr[i])
				arr = append(arr[:i], arr[i+1:]...)
				i-- // Adjust the loop index to account for the removed element.
			}
		}

		// Merge the sorted strand into the overall sorted list.
		sorted = strandMerge(sorted, strand)
	}

	// Return the final sorted list.
	return sorted
}

func strandMerge(a, b []int) []int {
	// Create a merged list with enough capacity to hold all elements from a and b.
	result := make([]int, len(a)+len(b))

	// Continue merging until both lists are empty.
	for len(a) > 0 && len(b) > 0 {
		// Compare the first elements of a and b.
		if a[0] > b[0] {
			// Append the smaller element to the merged list and remove it from a.
			result = append(result, a[0])
			a = a[1:]
		} else {
			// Append the smaller element to the merged list and remove it from b.
			result = append(result, b[0])
			b = b[1:]
		}
	}

	// Append any remaining elements from a and b to the merged list (if any).
	result = append(result, a...)
	result = append(result, b...)

	// Return the merged list containing all elements from a and b in sorted order.
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

1. **Pancake Sort Introduction**: Pancake Sort is a sorting algorithm that aims to sort an array of integers in ascending order.

2. **Loop Through Array (Descending)**: The `PancakeSort` function iterates through the array from the end to the beginning. The variable `i` starts from the length of the array and decrements with each iteration, effectively moving towards the left side of the array.

3. **Find the Maximum Element**: In each iteration, the `findMax` function is called to find the index of the maximum element within the current unsorted portion of the array, from index 0 to `i`.

4. **Check and Flip**: If the maximum element is not already at the end of the unsorted portion (index `i-1`), the algorithm performs two flips:
   - The first flip reverses the order of elements from the beginning of the array up to the largest element, effectively moving the largest element to the top of the stack (index 0).
   - The second flip places the largest element in its correct position at the bottom of the sorted portion of the array.

5. **Repeat Until Sorted**: This process continues in each iteration of the loop until the entire array is sorted in ascending order.

6. **Return Sorted Array**: Finally, the function returns the sorted array.

The `findMax` function is responsible for finding the index of the maximum element within a given portion of the array, and the `flip` function reverses elements within a specified range of the array. These helper functions are crucial for the Pancake Sort algorithm to work effectively.

```go
// Pancake Sort is a sorting algorithm that sorts an array of integers in ascending order.
func PancakeSort(array []int) []int {
	// Get the length of the input array
	n := len(array)
	
	// Loop through the array from the end, decrementing `i` to move towards the left
	for i := n; i > 1; i-- {
		// Find the index of the maximum element within the current unsorted portion
		maxIndex := findMax(array[:i])
		
		// If the maximum element is not at the end of the current unsorted portion
		if maxIndex != i-1 {
			// Flip the portion of the array from index 0 to `maxIndex`
			array = flip(array, maxIndex)
			// Flip the portion of the array from index 0 to `i-1`, moving the largest element to its correct position
			array = flip(array, i-1)
		}
	}
	
	// Return the sorted array
	return array
}

// findMax finds the index of the maximum element within the given array.
func findMax(array []int) int {
	maxIndex := 0
	
	// Loop through the elements of the array
	for i := 1; i < len(array); i++ {
		// If the current element is larger than the maximum found so far
		if array[i] > array[maxIndex] {
			// Update the index of the maximum element
			maxIndex = i
		}
	}
	
	// Return the index of the maximum element
	return maxIndex
}

// flip reverses the elements in the slice from index 0 to k, inclusive.
func flip(array []int, k int) []int {
	start := 0
	end := k
	
	// Reverse the elements within the specified range
	for start < end {
		// Swap elements at start and end indices
		array[start], array[end] = array[end], array[start]
		
		// Move the indices towards each other
		start++
		end--
	}
	
	// Return the modified array after the flip operation
	return array
}




// Recursive
func RecursivePanCake(array []int) []int {
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
func recursiveFlip(array []int, n int) []int {
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
// safeAppend safely appends an element of type T to a slice, protected by a mutex.
func safeAppend(arr *[]int, wg *sync.WaitGroup, mutex *sync.Mutex, num int) {
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
func Sleep(array []int) []int {
	// Create a wait group to track the completion of go-routines.
	var wg sync.WaitGroup

	// Create a mutex to ensure safe access to the sorted slice.
	var mutex sync.Mutex

	// Initialize the sorted slice with an initial capacity.
	sorted := make([]int, 0, len(array))

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

The `Insertion` function takes an integer slice called `array` as input and returns the sorted slice. The main goal of this function is to arrange the elements of the input array in ascending order using the Insertion Sort algorithm.

Inside the function, `n` represents the length of the input array, which helps determine the number of elements to be sorted. The loop starts from index 1 because it assumes that the first element at index 0 is already sorted.

For each element in the unsorted portion of the array (starting from index 1), the algorithm uses a variable called `key` to hold the current element being considered. The `key` variable also serves as the position where the element should be inserted if certain conditions are met during the sorting process.

The `sortedIndex` variable keeps track of the index in the sorted portion of the array, which starts at `index - 1`. It represents the position of the first value in the sorted portion, which will be compared with the `key` element.

The core of the algorithm involves comparing the `key` element with the elements in the sorted portion, moving them to the right until the correct position for the `key` element is found. This is achieved through a loop that continues as long as `sortedIndex` is greater than or equal to 0 and the value at the `sortedIndex` is greater than `key`.

Once the correct position for the `key` element is determined, it is inserted into the sorted portion of the array at `sortedIndex + 1`. This process continues for each element in the unsorted portion until the entire array is sorted.

It iterates through the array, comparing and inserting elements one by one, gradually building the sorted portion of the array from left to right. The `key` and `sortedIndex` variables play crucial roles in this process by holding the current element and tracking the position in the sorted portion, respectively.

```go
func Insertion(array []int) []int {
    var n int = len(array)
    // Loop through the elements of the array, starting from index 1,
    // because at position 0, we assume the first element is already in the sorted portion.
    for currentIndex := 1; currentIndex < n; currentIndex++ {
        // key: Holds the current element of the array that we're comparing and potentially inserting.
        var key int = array[currentIndex]
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

// Recursive


func RecursiveInsertion(array []int) []int {
	var n int = len(array)
	if n <= 1 {
		return array
	}
	RecursiveInsertion(array[:n-1])
	var key int = array[n-1]
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

Stooge Sort. It is designed to sort an integer slice (array) in ascending order.

1. **Initialization of `low` and `high`:** The `low` variable is initialized to 0, representing the starting index of the current sub-array, and `high` is initialized to the length of the input array, representing the ending index of the current sub-array.

2. **Checking and Swapping:** There is a conditional check to compare the last element (at index `high-1`) with the first element (at index `low`) of the sub-array. If the last element is smaller than the first element, they are swapped to ensure that the smallest and largest elements are correctly positioned.

3. **Calculating the Length of the SubArray:** The `length` variable is calculated to determine the size of the current sub-array. It's calculated by subtracting `low` from `high`.

4. **Recursive Sorting:** If the `length` of the sub-array is greater than 2 (meaning there are more than 2 elements to sort), the algorithm proceeds with the sorting process. It divides the sub-array into three parts: the first 2/3, the last 2/3, and the first 2/3 again. Then, Stooge Sort is called recursively on these sub-arrays.

5. **Returning the Sorted SubArray:** The function returns the sorted sub-array after all recursive sorting calls are complete.

Stooge Sort is a recursive sorting algorithm that repeatedly divides the array into smaller sub-arrays and sorts them to eventually produce a sorted array. The conditional swapping step at the beginning ensures that the smallest and largest elements are correctly positioned, making it a sorting algorithm with a unique approach.

```go
func Stooge(array []int) []int {
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

Explanation:

- **Space Complexity:** The space complexity of Stooge Sort is O(log n). This represents the amount of additional memory used by the algorithm for the recursive function call stack. The space complexity is relatively low and does not depend on the size of the input array.

- **Time Complexity:** The time complexity of Stooge Sort is not efficient. It can be expressed as O(n^(log 3 / log 1.5)), which is roughly equivalent to O(n^2.7095). This makes Stooge Sort impractical for sorting large datasets because it has a high time complexity.

- **Worst-case time complexity:** The worst-case time complexity of Stooge Sort is O(n^(log 3 / log 1.5)), which is roughly equivalent to O(n^2.7095). This means that in the worst-case scenario, when the array is in the reverse order, the algorithm exhibits very poor performance.

- **Best-case time complexity:** The best-case time complexity is also O(n^(log 3 / log 1.5)), which is the same as the worst-case time complexity. This is because Stooge Sort always performs the same number of comparisons and swaps, regardless of the initial order of the array.

- **Average-case time complexity:** The average-case time complexity of Stooge Sort is also O(n^(log 3 / log 1.5)). Like the best-case scenario, the algorithm's performance remains consistently poor on average for various input data distributions.

Stooge Sort has a poor time complexity in all cases, making it one of the least efficient sorting algorithms. It is typically used for educational purposes to demonstrate sorting algorithms' concepts rather than for practical sorting tasks. For real-world applications, more efficient sorting algorithms like Quick Sort, Merge Sort, or Heap Sort are preferred.

## Merge Sort

1. The `mergeSort` function takes an integer slice as input and sorts it using the Merge Sort algorithm.

2. We start by checking if the input array is nil or contains only one element. If it is, there's nothing to sort, so we return immediately.

3. We calculate the middle index, `mid`, of the array.

4. We create two slices, `left` and `right`, to hold the left and right halves of the array.

5. We copy the elements from the original array into the `left` and `right` slices, splitting the array into two halves.

6. We recursively call `mergeSort` on both the `left` and `right` slices to sort them.

7. We initialize variables `i`, `j`, and `k` to keep track of the indices while merging.

8. We merge the sorted `left` and `right` slices back into the original array. We compare the elements at indices `i` and `j` and place the smaller element in the original array at index `k`. We increment `i` or `j` as necessary and continue until we've merged all elements.

9. Finally, we collect any remaining elements from the `left` and `right` slices and place them in the original array.

```go
// mergeSort sorts the input slice using the Merge Sort algorithm.
func Merge(arr []int) []int {
    // Check if the input array is nil or empty.
    if arr == nil || len(arr) <= 1 {
        return arr
    }

    // Calculate the middle index of the array.
    mid := len(arr) / 2

    // Split the array into two halves: left and right.
    var left []int= make([]int, mid)
    var right []int = make([]int, len(arr)-mid)

    // Copy elements from the original array to the left and right arrays.
    for i := 0; i < mid; i++ {
        left[i] = arr[i]
    }
    for i := mid; i < len(arr); i++ {
        right[i-mid] = arr[i]
    }

    // Recursively sort the left and right halves.
    Merge(left)
    Merge(right)

	var i int = 0
	var j int = 0
	var k int = 0

    // Merge the sorted left and right arrays back into the original array.
    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            arr[k] = left[i]
            i++
        } else {
            arr[k] = right[j]
            j++
        }
        k++
    }

    // Collect any remaining elements from the left and right arrays.
    for i < len(left) {
        arr[k] = left[i]
        i++
        k++
    }
    for j < len(right) {
        arr[k] = right[j]
        j++
        k++
    }
	return arr
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
func Heap(array []int) []int {
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
func heapify(array []int, n, i int) []int {
	largest := i     // Initialize the largest element as the root
	left := 2*i + 1  // Index of the left child
	right := 2*i + 2 // Index of the right child

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

/*

	In-Place Sorting:
	The algorithm sorts the array in-place without using additional memory for temporary arrays.

*/
func Quick(array []int) []int {
	if len(array) < 2 {
		return array
	}

	pivotIndex := partition(array, 0, len(array)-1)

	Quick(array[:pivotIndex])
	Quick(array[pivotIndex+1:])

	return array
}

/*


   Three-Way Partitioning:
   In the partition function, elements equal to the pivot are not repeatedly compared.
   Instead, they are placed on the correct side of the pivot, reducing unnecessary comparisons



*/
func partition(arr []int, low int, high int) int {
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
	var pivot int = arr[high]
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

// Iterative Quick Sort

func Quick(array []int) []int {
	// if the array have only 1 or no element
	if len(array) < 2 {
		panic("Can't be Sorted: Less than Two Element")
	}

	stack := []int{0, len(array) - 1}

	for len(stack) > 0 {
		high := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		low := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		pivotIndex := itPartition(array, low, high)

		if pivotIndex-1 > low {
			stack = append(stack, low)
			stack = append(stack, pivotIndex-1)
		}
		if pivotIndex+1 < high {
			stack = append(stack, pivotIndex+1)
			stack = append(stack, high)
		}
	}
	return array
}

// Iterative Partition Function
func itPartition(array []int, low, high int) int {
	// set the highest like the far most element inside the array as the pivot
	pivot := array[high]
	index := low - 1
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
func merge(array []int, left []int, middle []int, right []int) []int {
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

func ThreeWayMerge(array []int) []int {
	if len(array) > 1 {
		mid1 := len(array) / 3
		mid2 := 2 * len(array) / 3
		// Split the input array into three parts: left, middle, and right
		left := make([]int, mid1)
		copy(left, array[:mid1])

		middle := make([]int, mid2-mid1)
		copy(middle, array[mid1:mid2])

		right := make([]int, len(array)-mid2)
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

### Space AND Time Complexity

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
func ThreeWayQuick(array []int) []int {
	// If the array is already sorted or contains only one element, no sorting is needed.
	if len(array) < 2 {
		return array
	}

	// Shuffle the array to improve performance, mitigating worst-case scenarios.
	shuffle(array)

	// Sort the array using the Three-Way Quick Sort algorithm.
	array = sort(array, 0, len(array)-1)
	return array
}

// sort is a recursive function that sorts the input array in place using Three-Way Quick Sort.
func sort(array []int, low, high int) []int {
	// Base case: If the high index is less than or equal to the low index, the array is sorted.
	if high <= low {
		return array
	}

	// Perform the partitioning step to separate elements into three partitions.
	lessThan, greaterThan := partition(array, low, high)

	// Recursively sort the partitions to the left and right of the pivot.
	sort(array, low, lessThan-1)
	sort(array, greaterThan+1, high)

	return array
}

// partition function partitions the array into three segments and returns the boundaries of equal elements.
func partition(array []int, low, high int) (int, int) {
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

// shuffle function shuffles the array to mitigate worst-case scenarios.
func shuffle(arr []int) {
	// Initialize the random number generator with a seed for consistency.
	rand.New(rand.NewSource(99))

	// Iterate through the array and swap elements randomly to shuffle.
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
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
func Counting(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Find the maximum value in the input array
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	// Initialize the counting array and output array
	counts := make([]int, max+1)
	output := make([]int, len(arr))

	// Count the occurrences of each element
	for _, num := range arr {
		counts[num]++
	}

	// Calculate the cumulative counts
	for i := 1; i <= max; i++ {
		counts[i] += counts[i-1]
	}

	// Build the sorted output array
	for i := len(arr) - 1; i >= 0; i-- {
		output[counts[arr[i]]-1] = arr[i]
		counts[arr[i]]--
	}

	return output
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

In summary, Counting Sort is an efficient linear-time sorting algorithm for sorting data with a limited range of values, making it well-suited for such scenarios.

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
// Function to find the maximum element in the array, considering both positive and negative numbers
func getMax(arr []int) int {
	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

// Function to perform counting sort based on a specific digit place (exp)
func countingSort(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n) // Create an output slice to store the sorted elements
	count := make([]int, 19) // Create a count slice for each digit, accounting for numbers from -9 to 9

	// Count the occurrences of each digit at the current place
	for i := 0; i < n; i++ {
		// Shift the digit range to be from 0 to 18 (for -9 to 9)
		count[(arr[i]/exp)%10+9]++
	}

	// Calculate the positions for each element in the output
	for i := 1; i < 19; i++ {
		count[i] += count[i-1]
	}

	// Place the elements in their sorted order in the output slice
	for i := n - 1; i >= 0; i-- {
		output[count[(arr[i]/exp)%10+9]-1] = arr[i]
		count[(arr[i]/exp)%10+9]--
	}

	// Copy the sorted elements back to the original array
	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}

// Main Radix Sort function
func Redix(arr []int) {
	max := getMax(arr)

	// Iterate through each digit place, from least significant to most significant
	for exp := 1; max/exp > 0; exp *= 10 {
		countingSort(arr, exp)
	}
}

```

### Space & Time Complexity

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
func Bucket(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Find the minimum and maximum values in the array
	minValue, maxValue := findMinMax(arr)

	// Create buckets
	numBuckets := maxValue - minValue + 1
	buckets := make([][]int, numBuckets)

	// Distribute elements into buckets
	for _, val := range arr {
		buckets[val-minValue] = append(buckets[val-minValue], val)
	}

	// Sort each bucket (using any sorting algorithm)
	sorted := []int{}
	for i := 0; i < numBuckets; i++ {
		if len(buckets[i]) > 0 {
			buckets[i] = Insertion(buckets[i]) // Use any sorting method for the buckets
			// - Insertion Sort we already implemented Above
			sorted = append(sorted, buckets[i]...)
		}
	}

	return sorted
}
```

### Space & Time Complexity

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

bitonic sort is most efficient when used with parallel processing or on hardware designed for parallelism, as it inherently has the potential for parallel execution during the sorting and merging phases.

### Normal Implementation - Bitonic Sort

```go
func BitonicSort(array []int, order bool) []int {
	return bitonicSort(array, 0, len(array), order)
}

func bitonicSort(array []int, low, high int, dir bool) []int {
	if high > 1 {
		var mid int = high / 2
		bitonicSort(array, low, mid, !dir)
		bitonicSort(array, low+mid, high-mid, dir)
		array = bitonicMerge(array, low, high, dir)
	}
	return array
}

func bitonicMerge(array []int, low, high int, dir bool) []int {
	if high > 1 {
		mid := greatestPowerOfTwoLessThan(high)
		for index := low; index < low+high-mid; index++ {
			if dir == (array[index] > array[index+mid]) {
				array[index], array[index+mid] = array[index+mid], array[index]
			}
		}
		array = bitonicMerge(array, low, mid, dir)
		array = bitonicMerge(array, low+mid, high-mid, dir)
	}
	return array
}

func greatestPowerOfTwoLessThan(n int) int {
	k := 1
	for k > 0 && k < n {
		k = k << 1
	}
	return k >> 1
}
```

### parallel Implementation

```go

import (
	"runtime"
	"sync"
)

func BitonicSort(arr []int, order bool) []int {
	// Set the maximum number of CPUs to utilize.
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	
	// Start the bitonic sort process with the full array.
	return bitonicSort(arr, 0, len(arr), order)
}

func bitonicSort(array []int, low, high int, order bool) []int {
	// Base case: If there's only one element or none, the array is already sorted.
	if high <= 1 {
		return array
	}
	
	// Calculate the middle point of the array.
	mid := high / 2
	
	// Use a WaitGroup to synchronize go-routines.
	var wg sync.WaitGroup
	wg.Add(2)

	// Start two go-routines to sort the left and right halves concurrently.
	go func() {
		bitonicSort(array, low, mid, !order) // Sort the left half.
		wg.Done()
	}()

	go func() {
		bitonicSort(array, low+mid, high-mid, order) // Sort the right half.
		wg.Done()
	}()

	// Wait for the two halves to be sorted.
	wg.Wait()

	// Merge the sorted halves.
	return bitonicMerge(array, low, high, order)
}

func bitonicMerge(array []int, low, high int, order bool) []int {
	// Base case: If there's only one element or none, the array is already sorted.
	if high <= 1 {
		return array
	}
	
	// Calculate the greatest power of two less than the 'high' value.
	mid := greatestPowerOfTwoLessThan(high)

	// Use a WaitGroup to synchronize go-routines for merging.
	var wg sync.WaitGroup
	wg.Add(2)

	// Compare and swap elements within the bitonic sequence.
	for index := low; index < low+high-mid; index++ {
		if order == (array[index] > array[index+mid]) {
			array[index], array[index+mid] = array[index+mid], array[index]
		}
	}

	// Start two go-routines to merge the left and right halves concurrently.
	go func() {
		bitonicMerge(array, low, mid, order) // Merge the left half.
		wg.Done()
	}()

	go func() {
		bitonicMerge(array, low+mid, high-mid, order) // Merge the right half.
		wg.Done()
	}

	// Wait for the two halves to be merged.
	wg.Wait()

	return array
}

func greatestPowerOfTwoLessThan(high int) int {
	k := 1
	// Calculate the greatest power of two less than 'high'.
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

func TimSort(array []int) []int {
	var n int = len(array)
	var minRun int = findMinRun(n)

	// Sort individual sub-arrays of size minRun
	for start := 0; start < n; start += minRun {
		end := start + minRun
		if end > n {
			end = n
		}
		// We have already Implemented Insertion Sort Long Time ago
		insertion(array[start:end])
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
	const minRun int = 64
	r := 0
	for n >= minRun {
		r |= n & 1
		n >>= 1
	}
	return n + r
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}


func insertion(array []int) []int {
	var j int

	for i := range array {
		j = i
		for j > 0 && array[j-1] > array[j] {
			array[j-1], array[j] = array[j], array[j-1]
			j--
		}
	}
	return array
}

func merging(left []int, right []int) []int {
	// Merge two sorted sub-arrays into one
	var result []int = make([]int, len(left)+len(right))
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
