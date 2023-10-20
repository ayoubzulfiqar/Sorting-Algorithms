package sort

import (
	"reflect"
	"testing"
	sorts "github.com/ayoubzulfiqar/Sorting-Algorithms/sort"

)

func TestStrandSort(t *testing.T) {
	t.Run("Strand Uint:", func(t *testing.T) {
		data := []uint{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []uint{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testStrandAssert(t, data, expected)
	})

	t.Run("Strand Int:", func(t *testing.T) {
		data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testStrandAssert(t, data, expected)
	})
	t.Run("Strand Negative Int:", func(t *testing.T) {
		data := []int{3, 1, 4, -2, -5, 9, -1, 6, 5, 3, 5}
		expected := []int{-5, -2, -1, 1, 3, 3, 4, 5, 5, 6, 9}
		testStrandAssert(t, data, expected)
	})
	t.Run("Strand Float:", func(t *testing.T) {
		// Test sorting for float64 slices
		data := []float64{3.14, 1.0, 2.71, 0.0}
		expected := []float64{0.0, 1.0, 2.71, 3.14}
		testStrandAssert(t, data, expected)
	})

	t.Run("Strand Negative Float:", func(t *testing.T) {
		data := []float64{3.14, 1.0, -2.71, 0.0, -1.5}
		expected := []float64{-2.71, -1.5, 0.0, 1.0, 3.14}
		testStrandAssert(t, data, expected)
	})

	t.Run("Strand Edge Case:", func(t *testing.T) {
		// Negative Ints
		testStrandAssert(t, []int{}, []int{})
		testStrandAssert(t, []int{-42}, []int{-42})
		testStrandAssert(t, []int{-1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1})

		// UInts
		testStrandAssert(t, []uint{}, []uint{})
		testStrandAssert(t, []uint{42}, []uint{42})
		testStrandAssert(t, []uint{1, 2, 3, 4, 5}, []uint{1, 2, 3, 4, 5})

		// Float Edge cases
		testStrandAssert(t, []float64{}, []float64{})
		testStrandAssert(t, []float64{42.0}, []float64{42.0})
		testStrandAssert(t, []float64{1.1, 1.1, 1.1, 1.1}, []float64{1.1, 1.1, 1.1, 1.1})
		testStrandAssert(t, []float64{3.0, 2.0, 1.0}, []float64{1.0, 2.0, 3.0})

	})
}

func testStrandAssert[N Numeric](t *testing.T, data, expected []N) {
	t.Helper()
	data = sorts.Strand(data)
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("\nSorting failed. Got %v, expected %v", data, expected)
	}
}

func BenchmarkStrand(b *testing.B) {
	b.Run("Strand Int", func(b *testing.B) {
		array := generateIntArray(10000)
		benchStrandAssert(b, array)
	})
	b.Run("Strand Float", func(b *testing.B) {
		array := generateFloatArray(10000)
		benchStrandAssert(b, array)
	})
	b.Run("Strand Negative", func(b *testing.B) {
		array := generateNegativeArray(10000)
		benchStrandAssert(b, array)
	})
}

func benchStrandAssert[N Numeric](b *testing.B, array []N) {
	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Strand(array)
	}
	b.StopTimer()
}
