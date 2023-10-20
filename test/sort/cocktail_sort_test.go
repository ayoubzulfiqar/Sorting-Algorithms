package sort

import (
	"reflect"
	"testing"
	sorts "github.com/ayoubzulfiqar/Sorting-Algorithms/sort"

)

func TestCockTailSort(t *testing.T) {
	t.Run("CockTail Uint:", func(t *testing.T) {
		data := []uint{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []uint{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testCockTailAssert(t, data, expected)
	})

	t.Run("CockTail Int:", func(t *testing.T) {
		data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testCockTailAssert(t, data, expected)
	})
	t.Run("CockTail Negative Int:", func(t *testing.T) {
		data := []int{3, 1, 4, -2, -5, 9, -1, 6, 5, 3, 5}
		expected := []int{-5, -2, -1, 1, 3, 3, 4, 5, 5, 6, 9}
		testCockTailAssert(t, data, expected)
	})
	t.Run("CockTail Float:", func(t *testing.T) {
		// Test sorting for float64 slices
		data := []float64{3.14, 1.0, 2.71, 0.0}
		expected := []float64{0.0, 1.0, 2.71, 3.14}
		testCockTailAssert(t, data, expected)
	})

	t.Run("CockTail Negative Float:", func(t *testing.T) {
		data := []float64{3.14, 1.0, -2.71, 0.0, -1.5}
		expected := []float64{-2.71, -1.5, 0.0, 1.0, 3.14}
		testCockTailAssert(t, data, expected)
	})

	t.Run("CockTail Edge Case:", func(t *testing.T) {
		// Negative Ints
		testCockTailAssert(t, []int{}, []int{})
		testCockTailAssert(t, []int{-42}, []int{-42})
		testCockTailAssert(t, []int{-1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1})

		// UInts
		testCockTailAssert(t, []uint{}, []uint{})
		testCockTailAssert(t, []uint{42}, []uint{42})
		testCockTailAssert(t, []uint{1, 2, 3, 4, 5}, []uint{1, 2, 3, 4, 5})

		// Float Edge cases
		testCockTailAssert(t, []float64{}, []float64{})
		testCockTailAssert(t, []float64{42.0}, []float64{42.0})
		testCockTailAssert(t, []float64{1.1, 1.1, 1.1, 1.1}, []float64{1.1, 1.1, 1.1, 1.1})
		testCockTailAssert(t, []float64{3.0, 2.0, 1.0}, []float64{1.0, 2.0, 3.0})

	})
}

func testCockTailAssert[N Numeric](t *testing.T, data, expected []N) {
	t.Helper()
	data = sorts.CockTail(data)
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("\nSorting failed. Got %v, expected %v", data, expected)
	}
}

func BenchmarkCockTail(b *testing.B) {
	b.Run("CockTail Int", func(b *testing.B) {
		array := generateIntArray(10000)
		benchCockTailAssert(b, array)
	})
	b.Run("CockTail Float", func(b *testing.B) {
		array := generateFloatArray(10000)
		benchCockTailAssert(b, array)
	})
	b.Run("CockTail Negative", func(b *testing.B) {
		array := generateNegativeArray(10000)
		benchCockTailAssert(b, array)
	})
}

func benchCockTailAssert[N Numeric](b *testing.B, array []N) {
	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.CockTail(array)
	}
	b.StopTimer()
}
