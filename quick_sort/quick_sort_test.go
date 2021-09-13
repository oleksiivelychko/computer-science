package quick_sort

import (
	"testing"
)

var array = [4]uint{10, 5, 2, 3}

func TestQuickSort(t *testing.T) {
	var sorted = QuickSort(array[:])
	if sorted[0] != 2 {
		t.Errorf("[func QuickSort(slice []uint) []uint] -> %d != 2", sorted[0])
	}
	if sorted[1] != 3 {
		t.Errorf("[func QuickSort(slice []uint) []uint] -> %d != 3", sorted[1])
	}
	if sorted[2] != 5 {
		t.Errorf("[func QuickSort(slice []uint) []uint] -> %d != 5", sorted[2])
	}
	if sorted[3] != 10 {
		t.Errorf("[func QuickSort(slice []uint) []uint] -> %d != 10", sorted[3])
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(array[:])
	}
}