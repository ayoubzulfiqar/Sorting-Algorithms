package sort

import (
	"reflect"
	"testing"
	sorts "github.com/ayoubzulfiqar/Sorting-Algorithms/sort"

)

func TestCycleSort(t *testing.T) {
	t.Run("Cycle Uint:", func(t *testing.T) {
		data := []uint{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []uint{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testCycleAssert(t, data, expected)
	})

	t.Run("Cycle Int:", func(t *testing.T) {
		data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testCycleAssert(t, data, expected)
	})
	t.Run("Cycle Negative Int:", func(t *testing.T) {
		data := []int{3, 1, 4, -2, -5, 9, -1, 6, 5, 3, 5}
		expected := []int{-5, -2, -1, 1, 3, 3, 4, 5, 5, 6, 9}
		testCycleAssert(t, data, expected)
	})
	t.Run("Cycle Float:", func(t *testing.T) {
		// Test sorting for float64 slices
		data := []float64{3.14, 1.0, 2.71, 0.0}
		expected := []float64{0.0, 1.0, 2.71, 3.14}
		testCycleAssert(t, data, expected)
	})

	t.Run("Cycle Negative Float:", func(t *testing.T) {
		data := []float64{3.14, 1.0, -2.71, 0.0, -1.5}
		expected := []float64{-2.71, -1.5, 0.0, 1.0, 3.14}
		testCycleAssert(t, data, expected)
	})

	t.Run("Cycle Edge Case:", func(t *testing.T) {
		// Negative Ints
		testCycleAssert(t, []int{}, []int{})
		testCycleAssert(t, []int{-42}, []int{-42})
		testCycleAssert(t, []int{-1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1})

		// UInts
		testCycleAssert(t, []uint{}, []uint{})
		testCycleAssert(t, []uint{42}, []uint{42})
		testCycleAssert(t, []uint{1, 2, 3, 4, 5}, []uint{1, 2, 3, 4, 5})

		// Float Edge cases
		testCycleAssert(t, []float64{}, []float64{})
		testCycleAssert(t, []float64{42.0}, []float64{42.0})
		testCycleAssert(t, []float64{1.1, 1.1, 1.1, 1.1}, []float64{1.1, 1.1, 1.1, 1.1})
		testCycleAssert(t, []float64{3.0, 2.0, 1.0}, []float64{1.0, 2.0, 3.0})

	})
}

func testCycleAssert[N Numeric](t *testing.T, data, expected []N) {
	t.Helper()
	data = sorts.Cycle(data)
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("\nSorting failed. Got %v, expected %v", data, expected)
	}
}

func BenchmarkCycle(b *testing.B) {
	b.Run("Cycle Int", func(b *testing.B) {
		array := generateIntArray(10000)
		benchCycleAssert(b, array)
	})
	b.Run("Cycle Float", func(b *testing.B) {
		array := generateFloatArray(10000)
		benchCycleAssert(b, array)
	})
	b.Run("Cycle Negative", func(b *testing.B) {
		array := generateNegativeArray(10000)
		benchCycleAssert(b, array)
	})
}

func benchCycleAssert[N Numeric](b *testing.B, array []N) {
	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Cycle(array)
	}
	b.StopTimer()
}
