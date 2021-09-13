package selection_sort

import (
	"testing"
)

var array = []uint{18, 21, 11}

func TestSelectionSort(t *testing.T) {
	var sorted = InsertionSort(array[:])

	if sorted[0] != 11 {
		t.Errorf("[func InsertionSort(slice []uint) []uint] -> %d != 11", sorted[0])
	}
	if sorted[1] != 18 {
		t.Errorf("[func InsertionSort(slice []uint) []uint] -> %d != 18", sorted[1])
	}
	if sorted[2] != 21 {
		t.Errorf("[func InsertionSort(slice []uint) []uint] -> %d != 21", sorted[2])
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort(array[:])
	}
}
