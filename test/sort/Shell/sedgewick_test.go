package sort

import (
	"reflect"
	"testing"

	sorts "github.com/ayoubzulfiqar/Sorting-Algorithms/sort/ShellSort"
)

func TestSedgewickSort(t *testing.T) {
	t.Run("Sedgewick Uint:", func(t *testing.T) {
		data := []uint{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []uint{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testSedgewickAssert(t, data, expected)
	})

	t.Run("Sedgewick Int:", func(t *testing.T) {
		data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		testSedgewickAssert(t, data, expected)
	})
	t.Run("Sedgewick Negative Int:", func(t *testing.T) {
		data := []int{3, 1, 4, -2, -5, 9, -1, 6, 5, 3, 5}
		expected := []int{-5, -2, -1, 1, 3, 3, 4, 5, 5, 6, 9}
		testSedgewickAssert(t, data, expected)
	})
	t.Run("Sedgewick Float:", func(t *testing.T) {
		// Test sorting for float64 slices
		data := []float64{3.14, 1.0, 2.71, 0.0}
		expected := []float64{0.0, 1.0, 2.71, 3.14}
		testSedgewickAssert(t, data, expected)
	})

	t.Run("Sedgewick Negative Float:", func(t *testing.T) {
		data := []float64{3.14, 1.0, -2.71, 0.0, -1.5}
		expected := []float64{-2.71, -1.5, 0.0, 1.0, 3.14}
		testSedgewickAssert(t, data, expected)
	})

	t.Run("Sedgewick Edge Case:", func(t *testing.T) {
		// Negative Ints
		testSedgewickAssert(t, []int{}, []int{})
		testSedgewickAssert(t, []int{-42}, []int{-42})
		testSedgewickAssert(t, []int{-1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1})

		// UInts
		testSedgewickAssert(t, []uint{}, []uint{})
		testSedgewickAssert(t, []uint{42}, []uint{42})
		testSedgewickAssert(t, []uint{1, 2, 3, 4, 5}, []uint{1, 2, 3, 4, 5})

		// Float Edge cases
		testSedgewickAssert(t, []float64{}, []float64{})
		testSedgewickAssert(t, []float64{42.0}, []float64{42.0})
		testSedgewickAssert(t, []float64{1.1, 1.1, 1.1, 1.1}, []float64{1.1, 1.1, 1.1, 1.1})
		testSedgewickAssert(t, []float64{3.0, 2.0, 1.0}, []float64{1.0, 2.0, 3.0})

	})
}

func testSedgewickAssert[N Numeric](t *testing.T, data, expected []N) {
	t.Helper()
	data = sorts.Sedgewick(data)
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("\nSorting failed. Got %v, expected %v", data, expected)
	}
}

func BenchmarkSedgewick(b *testing.B) {
	b.Run("Sedgewick Int", func(b *testing.B) {
		array := generateIntArray(10000)
		benchSedgewickAssert(b, array)
	})
	b.Run("Sedgewick Float", func(b *testing.B) {
		array := generateFloatArray(10000)
		benchSedgewickAssert(b, array)
	})
	b.Run("Sedgewick Negative", func(b *testing.B) {
		array := generateNegativeArray(10000)
		benchSedgewickAssert(b, array)
	})
}

func benchSedgewickAssert[N Numeric](b *testing.B, array []N) {
	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Sedgewick(array)
	}
	b.StopTimer()
}
