package sort

import (
	"reflect"
	"testing"
	sorts "github.com/ayoubzulfiqar/Sorting-Algorithms/sort"

)

func TestRedixSort(t *testing.T) {
	t.Run("Redix Uint:", func(t *testing.T) {
		data := []uint{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []uint{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testRedixAssert(t, data, expected)
	})

	t.Run("Redix Int:", func(t *testing.T) {
		data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testRedixAssert(t, data, expected)
	})
	t.Run("Redix Negative Int:", func(t *testing.T) {
		data := []int{3, 1, 4, -2, -5, 9, -1, 6, 5, 3, 5}
		expected := []int{-5, -2, -1, 1, 3, 3, 4, 5, 5, 6, 9}
		testRedixAssert(t, data, expected)
	})
	t.Run("Redix Float:", func(t *testing.T) {
		// Test sorting for float64 slices
		data := []float64{3.14, 1.0, 2.71, 0.0}
		expected := []float64{0.0, 1.0, 2.71, 3.14}
		testRedixAssert(t, data, expected)
	})

	t.Run("Redix Negative Float:", func(t *testing.T) {
		data := []float64{3.14, 1.0, -2.71, 0.0, -1.5}
		expected := []float64{-2.71, -1.5, 0.0, 1.0, 3.14}
		testRedixAssert(t, data, expected)
	})

	t.Run("Redix Edge Case:", func(t *testing.T) {
		// Negative Ints
		testRedixAssert(t, []int{}, []int{})
		testRedixAssert(t, []int{-42}, []int{-42})
		testRedixAssert(t, []int{-1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1})

		// UInts
		testRedixAssert(t, []uint{}, []uint{})
		testRedixAssert(t, []uint{42}, []uint{42})
		testRedixAssert(t, []uint{1, 2, 3, 4, 5}, []uint{1, 2, 3, 4, 5})

		// Float Edge cases
		testRedixAssert(t, []float64{}, []float64{})
		testRedixAssert(t, []float64{42.0}, []float64{42.0})
		testRedixAssert(t, []float64{1.1, 1.1, 1.1, 1.1}, []float64{1.1, 1.1, 1.1, 1.1})
		testRedixAssert(t, []float64{3.0, 2.0, 1.0}, []float64{1.0, 2.0, 3.0})

	})
}

func testRedixAssert[N Numeric](t *testing.T, data, expected []N) {
	t.Helper()
	data = sorts.Redix(data)
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("\nSorting failed. Got %v, expected %v", data, expected)
	}
}

func BenchmarkRedix(b *testing.B) {
	b.Run("Redix Int", func(b *testing.B) {
		array := generateIntArray(10000)
		benchRedixAssert(b, array)
	})
	b.Run("Redix Float", func(b *testing.B) {
		array := generateFloatArray(10000)
		benchRedixAssert(b, array)
	})
	b.Run("Redix Negative", func(b *testing.B) {
		array := generateNegativeArray(10000)
		benchRedixAssert(b, array)
	})
}

func benchRedixAssert[N Numeric](b *testing.B, array []N) {
	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Redix(array)
	}
	b.StopTimer()
}
