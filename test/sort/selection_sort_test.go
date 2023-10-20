package sort

import (
	"reflect"
	"testing"
	sorts "github.com/ayoubzulfiqar/Sorting-Algorithms/sort"

)

func TestSelectionSort(t *testing.T) {
	t.Run("Selection Uint:", func(t *testing.T) {
		data := []uint{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []uint{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testSelectionAssert(t, data, expected)
	})

	t.Run("Selection Int:", func(t *testing.T) {
		data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testSelectionAssert(t, data, expected)
	})
	t.Run("Selection Negative Int:", func(t *testing.T) {
		data := []int{3, 1, 4, -2, -5, 9, -1, 6, 5, 3, 5}
		expected := []int{-5, -2, -1, 1, 3, 3, 4, 5, 5, 6, 9}
		testSelectionAssert(t, data, expected)
	})
	t.Run("Selection Float:", func(t *testing.T) {
		// Test sorting for float64 slices
		data := []float64{3.14, 1.0, 2.71, 0.0}
		expected := []float64{0.0, 1.0, 2.71, 3.14}
		testSelectionAssert(t, data, expected)
	})

	t.Run("Selection Negative Float:", func(t *testing.T) {
		data := []float64{3.14, 1.0, -2.71, 0.0, -1.5}
		expected := []float64{-2.71, -1.5, 0.0, 1.0, 3.14}
		testSelectionAssert(t, data, expected)
	})

	t.Run("Selection Edge Case:", func(t *testing.T) {
		// Negative Ints
		testSelectionAssert(t, []int{}, []int{})
		testSelectionAssert(t, []int{-42}, []int{-42})
		testSelectionAssert(t, []int{-1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1})

		// UInts
		testSelectionAssert(t, []uint{}, []uint{})
		testSelectionAssert(t, []uint{42}, []uint{42})
		testSelectionAssert(t, []uint{1, 2, 3, 4, 5}, []uint{1, 2, 3, 4, 5})

		// Float Edge cases
		testSelectionAssert(t, []float64{}, []float64{})
		testSelectionAssert(t, []float64{42.0}, []float64{42.0})
		testSelectionAssert(t, []float64{1.1, 1.1, 1.1, 1.1}, []float64{1.1, 1.1, 1.1, 1.1})
		testSelectionAssert(t, []float64{3.0, 2.0, 1.0}, []float64{1.0, 2.0, 3.0})

	})
}

func testSelectionAssert[N Numeric](t *testing.T, data, expected []N) {
	t.Helper()
	data = sorts.Selection(data)
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("\nSorting failed. Got %v, expected %v", data, expected)
	}
}

func BenchmarkSelection(b *testing.B) {
	b.Run("Selection Int", func(b *testing.B) {
		array := generateIntArray(10000)
		benchSelectionAssert(b, array)
	})
	b.Run("Selection Float", func(b *testing.B) {
		array := generateFloatArray(10000)
		benchSelectionAssert(b, array)
	})
	b.Run("Selection Negative", func(b *testing.B) {
		array := generateNegativeArray(10000)
		benchSelectionAssert(b, array)
	})
}

func benchSelectionAssert[N Numeric](b *testing.B, array []N) {
	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Selection(array)
	}
	b.StopTimer()
}
