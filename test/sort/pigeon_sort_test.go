package sort

import (
	"reflect"
	"testing"
	sorts "github.com/ayoubzulfiqar/Sorting-Algorithms/sort"

)

func TestPigeonSort(t *testing.T) {
	t.Run("Pigeon Uint:", func(t *testing.T) {
		data := []uint{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []uint{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testPigeonAssert(t, data, expected)
	})

	t.Run("Pigeon Int:", func(t *testing.T) {
		data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testPigeonAssert(t, data, expected)
	})
	t.Run("Pigeon Negative Int:", func(t *testing.T) {
		data := []int{3, 1, 4, -2, -5, 9, -1, 6, 5, 3, 5}
		expected := []int{-5, -2, -1, 1, 3, 3, 4, 5, 5, 6, 9}
		testPigeonAssert(t, data, expected)
	})
	t.Run("Pigeon Float:", func(t *testing.T) {
		// Test sorting for float64 slices
		data := []float64{3.14, 1.0, 2.71, 0.0}
		expected := []float64{0.0, 1.0, 2.71, 3.14}
		testPigeonAssert(t, data, expected)
	})

	t.Run("Pigeon Negative Float:", func(t *testing.T) {
		data := []float64{3.14, 1.0, -2.71, 0.0, -1.5}
		expected := []float64{-2.71, -1.5, 0.0, 1.0, 3.14}
		testPigeonAssert(t, data, expected)
	})

	t.Run("Pigeon Edge Case:", func(t *testing.T) {
		// Negative Ints
		testPigeonAssert(t, []int{}, []int{})
		testPigeonAssert(t, []int{-42}, []int{-42})
		testPigeonAssert(t, []int{-1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1})

		// UInts
		testPigeonAssert(t, []uint{}, []uint{})
		testPigeonAssert(t, []uint{42}, []uint{42})
		testPigeonAssert(t, []uint{1, 2, 3, 4, 5}, []uint{1, 2, 3, 4, 5})

		// Float Edge cases
		testPigeonAssert(t, []float64{}, []float64{})
		testPigeonAssert(t, []float64{42.0}, []float64{42.0})
		testPigeonAssert(t, []float64{1.1, 1.1, 1.1, 1.1}, []float64{1.1, 1.1, 1.1, 1.1})
		testPigeonAssert(t, []float64{3.0, 2.0, 1.0}, []float64{1.0, 2.0, 3.0})

	})
}

func testPigeonAssert[N Numeric](t *testing.T, data, expected []N) {
	t.Helper()
	data = sorts.Pigeon(data)
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("\nSorting failed. Got %v, expected %v", data, expected)
	}
}

func BenchmarkPigeon(b *testing.B) {
	b.Run("Pigeon Int", func(b *testing.B) {
		array := generateIntArray(10000)
		benchPigeonAssert(b, array)
	})
	b.Run("Pigeon Float", func(b *testing.B) {
		array := generateFloatArray(10000)
		benchPigeonAssert(b, array)
	})
	b.Run("Pigeon Negative", func(b *testing.B) {
		array := generateNegativeArray(10000)
		benchPigeonAssert(b, array)
	})
}

func benchPigeonAssert[N Numeric](b *testing.B, array []N) {
	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Pigeon(array)
	}
	b.StopTimer()
}
