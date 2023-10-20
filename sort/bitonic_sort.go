package sort

import (
	"sync"
)

/*
// Bitonic Sort - Sort Any Type (No power of two require)

func BitonicSort[T Numeric](array []T, order bool) []T {
	return bitonicSort(array, 0, len(array), order)
}

func bitonicSort[T Numeric](array []T, low, high int, dir bool) []T {
	if high > 1 {
		var mid int = high / 2
		bitonicSort(array, low, mid, !dir)
		bitonicSort(array, low+mid, high-mid, dir)
		array = bitonicMerge(array, low, high, dir)
	}
	return array
}

func bitonicMerge[T Numeric](array []T, low, high int, dir bool) []T {
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



*/

// ========= Co-Current version

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

// Reading: https://hwlang.de/algorithmen/sortieren/bitonic/oddn.htm
