package sort

import (
	"reflect"
	"testing"
	sorts "github.com/ayoubzulfiqar/Sorting-Algorithms/sort/ShellSort"

)

func TestPrattSort(t *testing.T) {
	t.Run("Pratt Uint:", func(t *testing.T) {
		data := []uint{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []uint{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testPrattAssert(t, data, expected)
	})

	t.Run("Pratt Int:", func(t *testing.T) {
		data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testPrattAssert(t, data, expected)
	})
	t.Run("Pratt Negative Int:", func(t *testing.T) {
		data := []int{3, 1, 4, -2, -5, 9, -1, 6, 5, 3, 5}
		expected := []int{-5, -2, -1, 1, 3, 3, 4, 5, 5, 6, 9}
		testPrattAssert(t, data, expected)
	})
	t.Run("Pratt Float:", func(t *testing.T) {
		// Test sorting for float64 slices
		data := []float64{3.14, 1.0, 2.71, 0.0}
		expected := []float64{0.0, 1.0, 2.71, 3.14}
		testPrattAssert(t, data, expected)
	})

	t.Run("Pratt Negative Float:", func(t *testing.T) {
		data := []float64{3.14, 1.0, -2.71, 0.0, -1.5}
		expected := []float64{-2.71, -1.5, 0.0, 1.0, 3.14}
		testPrattAssert(t, data, expected)
	})

	t.Run("Pratt Edge Case:", func(t *testing.T) {
		// Negative Ints
		testPrattAssert(t, []int{}, []int{})
		testPrattAssert(t, []int{-42}, []int{-42})
		testPrattAssert(t, []int{-1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1})

		// UInts
		testPrattAssert(t, []uint{}, []uint{})
		testPrattAssert(t, []uint{42}, []uint{42})
		testPrattAssert(t, []uint{1, 2, 3, 4, 5}, []uint{1, 2, 3, 4, 5})

		// Float Edge cases
		testPrattAssert(t, []float64{}, []float64{})
		testPrattAssert(t, []float64{42.0}, []float64{42.0})
		testPrattAssert(t, []float64{1.1, 1.1, 1.1, 1.1}, []float64{1.1, 1.1, 1.1, 1.1})
		testPrattAssert(t, []float64{3.0, 2.0, 1.0}, []float64{1.0, 2.0, 3.0})

	})
}

func testPrattAssert[N Numeric](t *testing.T, data, expected []N) {
	t.Helper()
	data = sorts.Pratt(data)
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("\nSorting failed. Got %v, expected %v", data, expected)
	}
}

func BenchmarkPratt(b *testing.B) {
	b.Run("Pratt Int", func(b *testing.B) {
		array := generateIntArray(10000)
		benchPrattAssert(b, array)
	})
	b.Run("Pratt Float", func(b *testing.B) {
		array := generateFloatArray(10000)
		benchPrattAssert(b, array)
	})
	b.Run("Pratt Negative", func(b *testing.B) {
		array := generateNegativeArray(10000)
		benchPrattAssert(b, array)
	})
}

func benchPrattAssert[N Numeric](b *testing.B, array []N) {
	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Pratt(array)
	}
	b.StopTimer()
}
