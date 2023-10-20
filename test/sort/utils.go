package sort

import (
	"math/rand"
)

// Helper function to generate a random array of a given size
func generateIntArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(10000) // Adjust the upper bound as needed
	}
	return arr
}

func generateFloatArray(size int) []float64 {
	arr := make([]float64, size)
	rangeSize := 50.0 - 10.0

	for i := 0; i < size; i++ {
		arr[i] = rangeSize + rand.Float64()*545*62
	}

	return arr
}

func generateNegativeArray(size int) []int {

	data := make([]int, size)
	for i := 0; i < size; i++ {
		// Generate a random negative integer
		data[i] = -rand.Intn(1 << 30) // Use a large number to ensure negative values
	}

	return data
}

type Numeric interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
