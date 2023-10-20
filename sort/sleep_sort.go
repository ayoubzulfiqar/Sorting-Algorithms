package sort

import "sync"

// safeAppend safely appends an element of type T to a slice, protected by a mutex.
func safeAppend[T Numeric](arr *[]T, wg *sync.WaitGroup, mutex *sync.Mutex, num T) {
	// Mark the wait group as done when this goroutine exits.
	defer wg.Done()

	// Lock the mutex to ensure exclusive access to the slice.
	mutex.Lock()

	// Append the element to the slice.
	*arr = append(*arr, num)

	// Unlock the mutex to allow other goroutines to access the slice.
	mutex.Unlock()
}

// Sleep is an implementation of the Sleep Sort algorithm that can work with slices of different numeric types.
func Sleep[T Numeric](array []T) []T {
	// Create a wait group to track the completion of goroutines.
	var wg sync.WaitGroup

	// Create a mutex to ensure safe access to the sorted slice.
	var mutex sync.Mutex

	// Initialize the sorted slice with an initial capacity.
	sorted := make([]T, 0, len(array))

	// Iterate through the input array.
	for _, num := range array {
		// Increment the wait group counter to track the number of active goroutines.
		wg.Add(1)

		// Launch a goroutine to safely append the element to the sorted slice.
		go safeAppend(&sorted, &wg, &mutex, num)
	}

	// Wait for all goroutines to finish before proceeding.
	wg.Wait()

	// Return the sorted slice with initially allocated space removed.
	return sorted
}
