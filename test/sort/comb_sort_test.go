package sort

import (
	"reflect"
	"testing"

	sorts "github.com/ayoubzulfiqar/Sorting-Algorithms/sort"

)

func TestCombSort(t *testing.T) {
	t.Run("Comb Uint:", func(t *testing.T) {
		data := []uint{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []uint{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testCombAssert(t, data, expected)
	})

	t.Run("Comb Int:", func(t *testing.T) {
		data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testCombAssert(t, data, expected)
	})
	t.Run("Comb Negative Int:", func(t *testing.T) {
		data := []int{3, 1, 4, -2, -5, 9, -1, 6, 5, 3, 5}
		expected := []int{-5, -2, -1, 1, 3, 3, 4, 5, 5, 6, 9}
		testCombAssert(t, data, expected)
	})
	t.Run("Comb Float:", func(t *testing.T) {
		// Test sorting for float64 slices
		data := []float64{3.14, 1.0, 2.71, 0.0}
		expected := []float64{0.0, 1.0, 2.71, 3.14}
		testCombAssert(t, data, expected)
	})

	t.Run("Comb Negative Float:", func(t *testing.T) {
		data := []float64{3.14, 1.0, -2.71, 0.0, -1.5}
		expected := []float64{-2.71, -1.5, 0.0, 1.0, 3.14}
		testCombAssert(t, data, expected)
	})

	t.Run("Comb Edge Case:", func(t *testing.T) {
		// Negative Ints
		testCombAssert(t, []int{}, []int{})
		testCombAssert(t, []int{-42}, []int{-42})
		testCombAssert(t, []int{-1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1})

		// UInts
		testCombAssert(t, []uint{}, []uint{})
		testCombAssert(t, []uint{42}, []uint{42})
		testCombAssert(t, []uint{1, 2, 3, 4, 5}, []uint{1, 2, 3, 4, 5})

		// Float Edge cases
		testCombAssert(t, []float64{}, []float64{})
		testCombAssert(t, []float64{42.0}, []float64{42.0})
		testCombAssert(t, []float64{1.1, 1.1, 1.1, 1.1}, []float64{1.1, 1.1, 1.1, 1.1})
		testCombAssert(t, []float64{3.0, 2.0, 1.0}, []float64{1.0, 2.0, 3.0})

	})
}

func testCombAssert[N Numeric](t *testing.T, data, expected []N) {
	t.Helper()
	data = sorts.Comb(data)
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("\nSorting failed. Got %v, expected %v", data, expected)
	}
}

func BenchmarkComb(b *testing.B) {
	b.Run("Comb Int", func(b *testing.B) {
		array := generateIntArray(10000)
		benchCombAssert(b, array)
	})
	b.Run("Comb Float", func(b *testing.B) {
		array := generateFloatArray(10000)
		benchCombAssert(b, array)
	})
	b.Run("Comb Negative", func(b *testing.B) {
		array := generateNegativeArray(10000)
		benchCombAssert(b, array)
	})
}

func benchCombAssert[N Numeric](b *testing.B, array []N) {
	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Comb(array)
	}
	b.StopTimer()
}
