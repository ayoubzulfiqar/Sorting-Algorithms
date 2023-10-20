package sort

import (
	"reflect"
	"testing"
	sorts "github.com/ayoubzulfiqar/Sorting-Algorithms/sort"

)

func TestStoogeSort(t *testing.T) {
	t.Run("Stooge Uint:", func(t *testing.T) {
		data := []uint{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []uint{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testStoogeAssert(t, data, expected)
	})

	t.Run("Stooge Int:", func(t *testing.T) {
		data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testStoogeAssert(t, data, expected)
	})
	t.Run("Stooge Negative Int:", func(t *testing.T) {
		data := []int{3, 1, 4, -2, -5, 9, -1, 6, 5, 3, 5}
		expected := []int{-5, -2, -1, 1, 3, 3, 4, 5, 5, 6, 9}
		testStoogeAssert(t, data, expected)
	})
	t.Run("Stooge Float:", func(t *testing.T) {
		// Test sorting for float64 slices
		data := []float64{3.14, 1.0, 2.71, 0.0}
		expected := []float64{0.0, 1.0, 2.71, 3.14}
		testStoogeAssert(t, data, expected)
	})

	t.Run("Stooge Negative Float:", func(t *testing.T) {
		data := []float64{3.14, 1.0, -2.71, 0.0, -1.5}
		expected := []float64{-2.71, -1.5, 0.0, 1.0, 3.14}
		testStoogeAssert(t, data, expected)
	})

	t.Run("Stooge Edge Case:", func(t *testing.T) {
		// Negative Ints
		testStoogeAssert(t, []int{}, []int{})
		testStoogeAssert(t, []int{-42}, []int{-42})
		testStoogeAssert(t, []int{-1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1})

		// UInts
		testStoogeAssert(t, []uint{}, []uint{})
		testStoogeAssert(t, []uint{42}, []uint{42})
		testStoogeAssert(t, []uint{1, 2, 3, 4, 5}, []uint{1, 2, 3, 4, 5})

		// Float Edge cases
		testStoogeAssert(t, []float64{}, []float64{})
		testStoogeAssert(t, []float64{42.0}, []float64{42.0})
		testStoogeAssert(t, []float64{1.1, 1.1, 1.1, 1.1}, []float64{1.1, 1.1, 1.1, 1.1})
		testStoogeAssert(t, []float64{3.0, 2.0, 1.0}, []float64{1.0, 2.0, 3.0})

	})
}

func testStoogeAssert[N Numeric](t *testing.T, data, expected []N) {
	t.Helper()
	data = sorts.Stooge(data)
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("\nSorting failed. Got %v, expected %v", data, expected)
	}
}

func BenchmarkStooge(b *testing.B) {
	b.Run("Stooge Int", func(b *testing.B) {
		array := generateIntArray(10000)
		benchStoogeAssert(b, array)
	})
	b.Run("Stooge Float", func(b *testing.B) {
		array := generateFloatArray(10000)
		benchStoogeAssert(b, array)
	})
	b.Run("Stooge Negative", func(b *testing.B) {
		array := generateNegativeArray(10000)
		benchStoogeAssert(b, array)
	})
}

func benchStoogeAssert[N Numeric](b *testing.B, array []N) {
	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Stooge(array)
	}
	b.StopTimer()
}
