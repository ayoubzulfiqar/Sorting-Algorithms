package sort

import (
	"math/rand"
)

// shuffle function shuffles the array to mitigate worst-case scenarios.
func shuffle[T Numeric](arr []T) []T {
	// Initialize the random number generator with a seed for consistency.
	rand.New(rand.NewSource(99))

	// Iterate through the array and swap elements randomly to shuffle.
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

//

// Finding the Max Element inside the Array
func findMax[T Numeric](array []T) T {
	var maxIndex T = 0
	// loop through the element of the array
	for i := 1; i < len(array); i++ {
		// if the element is inside the array is larger than the our max Value
		if array[i] > array[int(maxIndex)] {
			// we found new larger element
			maxIndex = T(i)
		}
	}
	return maxIndex

}

//

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
