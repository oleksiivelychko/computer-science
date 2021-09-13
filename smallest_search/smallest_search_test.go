package smallest_search

import (
	"testing"
)

var slice = []uint{5, 2, 6, 3, 0, 11, 8}

func TestSmallestSearch(t *testing.T) {
	result := SmallestSearch(slice, slice[0])
	if result != 0 {
		t.Errorf("[func SmallestSearch(slice []uint) uint] -> %d != 0", result)
	}
}

func BenchmarkSmallestSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SmallestSearch(slice, slice[0])
	}
}
