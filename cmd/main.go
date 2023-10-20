package main

import (
	"fmt"
	"time"

	sort "github.com/ayoubzulfiqar/Sorting-Algorithms/sort"
)

type Numeric interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func main() {
	ms, us, s := MeasureExecutionTime()
	fmt.Printf("Execution time in milliseconds: %d ms\n", ms)
	fmt.Printf("Execution time in microseconds: %d µs\n", us)
	fmt.Printf("Execution time in seconds: %d s\n", s)

}

// Function to measure execution time
func MeasureExecutionTime() (time.Duration, time.Duration, time.Duration) {
	startTime := time.Now()

	// Call the function you want to measure
	// arr := util.GenerateInt(1000000)
	arr := []int{1, 3, 5, 7, 8, 2, 4, 6, 9}
	// arr := []float64{9.0, 7.0, 5.0, 11.0, 12.0, 2.0, 14.0, 3.0, 10.0, 6.0}
	// arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	fmt.Printf("Unsorted %v\n", arr)

	arr = sort.Redix(arr)
	fmt.Printf("Sorted %v\n", arr)

	endTime := time.Now()

	// Calculate the duration in milliseconds, microseconds, and minutes
	duration := endTime.Sub(startTime)
	durationMilliseconds := duration / time.Millisecond
	durationMicroseconds := duration / time.Microsecond
	durationSeconds := duration / time.Second

	return durationMilliseconds, durationMicroseconds, durationSeconds
}
