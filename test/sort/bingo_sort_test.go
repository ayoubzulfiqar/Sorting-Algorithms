package sort

import (
	"reflect"
	"testing"

	sorts "github.com/ayoubzulfiqar/Sorting-Algorithms/sort"
)

func TestBingoSort(t *testing.T) {
	t.Run("BingoUint:", func(t *testing.T) {
		data := []uint{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []uint{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testBingoAssert(t, data, expected)
	})

	t.Run("Bingo Int:", func(t *testing.T) {
		data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testBingoAssert(t, data, expected)
	})
	t.Run("Bingo Negative Int:", func(t *testing.T) {
		data := []int{3, 1, 4, -2, -5, 9, -1, 6, 5, 3, 5}
		expected := []int{-5, -2, -1, 1, 3, 3, 4, 5, 5, 6, 9}
		testBingoAssert(t, data, expected)
	})
	t.Run("Bingo Float:", func(t *testing.T) {
		// Test sorting for float64 slices
		data := []float64{3.14, 1.0, 2.71, 0.0}
		expected := []float64{0.0, 1.0, 2.71, 3.14}
		testBingoAssert(t, data, expected)
	})

	t.Run("Bingo Negative Float:", func(t *testing.T) {
		data := []float64{3.14, 1.0, -2.71, 0.0, -1.5}
		expected := []float64{-2.71, -1.5, 0.0, 1.0, 3.14}
		testBingoAssert(t, data, expected)
	})

	t.Run("Bingo Edge Case:", func(t *testing.T) {
		// Negative Ints
		testBingoAssert(t, []int{}, []int{})
		testBingoAssert(t, []int{-42}, []int{-42})
		testBingoAssert(t, []int{-1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1})

		// UInts
		testBingoAssert(t, []uint{}, []uint{})
		testBingoAssert(t, []uint{42}, []uint{42})
		testBingoAssert(t, []uint{1, 2, 3, 4, 5}, []uint{1, 2, 3, 4, 5})

		// Float Edge cases
		testBingoAssert(t, []float64{}, []float64{})
		testBingoAssert(t, []float64{42.0}, []float64{42.0})
		testBingoAssert(t, []float64{1.1, 1.1, 1.1, 1.1}, []float64{1.1, 1.1, 1.1, 1.1})
		testBingoAssert(t, []float64{3.0, 2.0, 1.0}, []float64{1.0, 2.0, 3.0})

	})
}

func testBingoAssert[N Numeric](t *testing.T, data, expected []N) {
	t.Helper()
	data = sorts.Bingo(data)
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("\nSorting failed. Got %v, expected %v", data, expected)
	}
}

func BenchmarkBingo(b *testing.B) {
	b.Run("Bingo Int", func(b *testing.B) {
		array := generateIntArray(10000)
		benchBingoAssert(b, array)
	})
	b.Run("Bingo Float", func(b *testing.B) {
		array := generateFloatArray(10000)
		benchBingoAssert(b, array)
	})
	b.Run("Bingo Negative", func(b *testing.B) {
		array := generateNegativeArray(10000)
		benchBingoAssert(b, array)
	})
}

func benchBingoAssert[N Numeric](b *testing.B, array []N) {
	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Bingo(array)
	}
	b.StopTimer()
}
