package sort

/*

This Bitonic Sort is Concurrent and fast
But the Problem is that it require the len(array) power of 2
Which means if array is

arr := []int{1,2,3,4,5,6,7,8,9}      the length is 8
so 8 can equally divided by 2 it mean It will Work

if

arr := []int{0,1,2,3,4,5,6,7,8,9}      the length is 9

the 9/2 is not totally divided able so this will not sort correctly

*/

import (
	"sync"
)

// thresHold is used to decide whether to run concurrently.
// const thresHold = 1 << 14

// Bitonic Sort - Concurrent - Only Work on Power of Two
func Bitonic_Sort(array []int, ord bool) []int {
	return bitonic_Sort(array, ord, len(array))
}

func bitonic_Sort(array []int, ord bool, ln int) []int {
	if ln <= 1 {
		return array
	}
	mid := ln >> 1
	if mid >= 1<<14 {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			bitonic_Sort(array[:mid], true, mid)
			wg.Done()
		}()
		go func() {
			bitonic_Sort(array[mid:], false, mid)
			wg.Done()
		}()
		wg.Wait()
	} else {
		bitonic_Sort(array[:mid], true, mid)
		bitonic_Sort(array[mid:], false, mid)
	}
	return bitonic_Merge(array, ord, ln)
}

func bitonic_Merge(array []int, ord bool, ln int) []int {
	if ln <= 1 {
		return array
	}
	compareAndSwap(array, ord, ln)
	mid := ln >> 1
	if mid >= 1<<14 {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			bitonic_Merge(array[:mid], ord, mid)
			wg.Done()
		}()
		go func() {
			bitonic_Merge(array[mid:], ord, mid)
			wg.Done()
		}()
		wg.Wait()
	} else {
		bitonic_Merge(array[:mid], ord, mid)
		bitonic_Merge(array[mid:], ord, mid)
	}
	return array
}

func compareAndSwap(array []int, ord bool, ln int) {
	mid := ln >> 1
	for i := 0; i < mid; i++ {
		peer := mid ^ i
		if (array[i] > array[peer]) == ord {
			array[i] = array[i] ^ array[peer]
			array[peer] = array[i] ^ array[peer]
			array[i] = array[i] ^ array[peer]
		}
	}
}

//===== Another Implementation - Non-Concurrent

// `len(x)` must be a power of 2.
func RecursiveBitonic(array []int, ord bool) {
	recursiveBitonicSort(array, ord, len(array))
}

func recursiveBitonicSort(array []int, ord bool, ln int) {
	if ln <= 1 {
		return
	}
	mid := ln >> 1
	recursiveBitonicSort(array[:mid], true, mid)
	recursiveBitonicSort(array[mid:], false, mid)
	recursiveBitonicMerge(array, ord, ln)
}

func recursiveBitonicMerge(array []int, ord bool, ln int) {
	if ln <= 1 {
		return
	}
	compareAndSwap(array, ord, ln)
	mid := ln >> 1
	recursiveBitonicMerge(array[:mid], ord, mid)
	recursiveBitonicMerge(array[mid:], ord, mid)
}
