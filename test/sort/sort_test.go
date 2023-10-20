package sort

// Checking Speed of Every Sorting Algorithm _ Let's Get Naughty

// 10 Thousand := 10000
// 1 Million := 1000000
// 10 Million := 10000000
// 100 Million := 100000000

import (
	"testing"
	sorts "github.com/ayoubzulfiqar/Sorting-Algorithms/sort"

)

const tenTH int = 10000

func BenchmarkBingoSort(b *testing.B) {
	array := generateIntArray(tenTH)
	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Bingo(array)
	}
	b.StopTimer()
}

func BenchmarkBogoSort(b *testing.B) {
	array := generateIntArray(tenTH)
	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Bogo(array)
	}
	b.StopTimer()
}

func BenchmarkBubbleSort(b *testing.B) {
	array := generateIntArray(tenTH)

	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Bubble(array)
	}
	b.StopTimer()
}

func BenchmarkBucketSort(b *testing.B) {
	array := generateIntArray(tenTH)
	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Bucket(array)
	}
	b.StopTimer()
}

func BenchmarkCockTailSort(b *testing.B) {
	array := generateIntArray(tenTH)

	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.CockTail(array)
	}
	b.StopTimer()
}
func BenchmarkCombSort(b *testing.B) {
	array := generateIntArray(tenTH)

	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Comb(array)
	}
	b.StopTimer()
}
func BenchmarkCountingSort(b *testing.B) {
	array := generateIntArray(tenTH)

	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Counting(array)
	}
	b.StopTimer()
}
func BenchmarkCycleSort(b *testing.B) {
	array := generateIntArray(tenTH)

	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Cycle(array)
	}
	b.StopTimer()
}

func BenchmarkGnomeSort(b *testing.B) {
	array := generateIntArray(tenTH)

	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Gnome(array)
	}
	b.StopTimer()
}

func BenchmarkHeapSort(b *testing.B) {
	array := generateIntArray(tenTH)

	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Heap(array)
	}
	b.StopTimer()
}
func BenchmarkInsertionSort(b *testing.B) {
	array := generateIntArray(tenTH)

	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Insertion(array)
	}
	b.StopTimer()
}

func BenchmarkMergeSort(b *testing.B) {
	array := generateIntArray(tenTH)

	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Merge(array)
	}
	b.StopTimer()
}

func BenchmarkOddEvenSort(b *testing.B) {
	array := generateIntArray(tenTH)

	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.OddEven(array)
	}
	b.StopTimer()
}
func BenchmarkPanCakeSort(b *testing.B) {
	array := generateIntArray(tenTH)

	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.PanCake(array)
	}
	b.StopTimer()
}

func BenchmarkPigeonSort(b *testing.B) {
	array := generateIntArray(tenTH)

	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Pigeon(array)
	}
	b.StopTimer()
}

func BenchmarkRedixSort(b *testing.B) {
	array := generateIntArray(tenTH)

	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Redix(array)
	}
	b.StopTimer()
}
func BenchmarkSelectionSort(b *testing.B) {
	array := generateIntArray(tenTH)

	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Selection(array)
	}
	b.StopTimer()
}

func BenchmarkStoogeSort(b *testing.B) {
	array := generateIntArray(tenTH)

	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Stooge(array)
	}
	b.StopTimer()
}
func BenchmarkStrandSort(b *testing.B) {
	array := generateIntArray(tenTH)

	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Strand(array)
	}
	b.StopTimer()
}

func BenchmarkThreeWayMergeSort(b *testing.B) {
	array := generateIntArray(tenTH)

	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.ThreeWayMerge(array)
	}
	b.StopTimer()
}

func BenchmarkThreeWayQuickSort(b *testing.B) {
	array := generateIntArray(tenTH)

	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.ThreeWayQuick(array)
	}
	b.StopTimer()
}
func BenchmarkTimSort(b *testing.B) {
	array := generateIntArray(tenTH)

	b.Helper()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		array = sorts.Tim(array)
	}
	b.StopTimer()
}
