package sort

import (
	"reflect"
	"testing"
	sorts "github.com/ayoubzulfiqar/Sorting-Algorithms/sort/ShellSort"

)

func TestPapernovStasevichSort(t *testing.T) {
	t.Run("PapernovStasevich Uint:", func(t *testing.T) {
		data := []uint{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []uint{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testPapernovStasevichAssert(t, data, expected)
	})

	t.Run("PapernovStasevich Int:", func(t *testing.T) {
		data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testPapernovStasevichAssert(t, data, expected)
	})
	t.Run("PapernovStasevich Negative Int:", func(t *testing.T) {
		data := []int{3, 1, 4, -2, -5, 9, -1, 6, 5, 3, 5}
		expected := []int{-5, -2, -1, 1, 3, 3, 4, 5, 5, 6, 9}
		testPapernovStasevichAssert(t, data, expected)
	})
	t.Run("PapernovStasevich Float:", func(t *testing.T) {
		// Test sorting for float64 slices
		data := []float64{3.14, 1.0, 2.71, 0.0}
		expected := []float64{0.0, 1.0, 2.71, 3.14}
		testPapernovStasevichAssert(t, data, expected)
	})

	t.Run("PapernovStasevich Negative Float:", func(t *testing.T) {
		data := []float64{3.14, 1.0, -2.71, 0.0, -1.5}
		expected := []float64{-2.71, -1.5, 0.0, 1.0, 3.14}
		testPapernovStasevichAssert(t, data, expected)
	})

	t.Run("PapernovStasevich Edge Case:", func(t *testing.T) {
		// Negative Ints
		testPapernovStasevichAssert(t, []int{}, []int{})
		testPapernovStasevichAssert(t, []int{-42}, []int{-42})
		testPapernovStasevichAssert(t, []int{-1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1})

		// UInts
		testPapernovStasevichAssert(t, []uint{}, []uint{})
		testPapernovStasevichAssert(t, []uint{42}, []uint{42})
		testPapernovStasevichAssert(t, []uint{1, 2, 3, 4, 5}, []uint{1, 2, 3, 4, 5})

		// Float Edge cases
		testPapernovStasevichAssert(t, []float64{}, []float64{})
		testPapernovStasevichAssert(t, []float64{42.0}, []float64{42.0})
		testPapernovStasevichAssert(t, []float64{1.1, 1.1, 1.1, 1.1}, []float64{1.1, 1.1, 1.1, 1.1})
		testPapernovStasevichAssert(t, []float64{3.0, 2.0, 1.0}, []float64{1.0, 2.0, 3.0})

	})
}

func testPapernovStasevichAssert[N Numeric](t *testing.T, data, expected []N) {
	t.Helper()
	data = sorts.PapernovStasevich(data)
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("\nSorting failed. Got %v, expected %v", data, expected)
	}
}

func BenchmarkPapernovStasevich(b *testing.B) {
	b.Run("PapernovStasevich Int", func(b *testing.B) {
		array := generateIntArray(10000)
		benchPapernovStasevichAssert(b, array)
	})
	b.Run("PapernovStasevich Float", func(b *testing.B) {
		array := generateFloatArray(10000)
		benchPapernovStasevichAssert(b, array)
	})
	b.Run("PapernovStasevich Negative", func(b *testing.B) {
		array := generateNegativeArray(10000)
		benchPapernovStasevichAssert(b, array)
	})
}

func benchPapernovStasevichAssert[N Numeric](b *testing.B, array []N) {
	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.PapernovStasevich(array)
	}
	b.StopTimer()
}
