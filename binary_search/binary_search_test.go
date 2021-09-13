package binary_search

import (
	"math/rand"
	"testing"
)

var array = [8]uint{1, 3, 5, 7, 9, 11, 13, 15}

func TestBinarySearch(t *testing.T) {
	index := BinarySearch(array[:], 5)
	if index < 0 {
		t.Errorf("[func BinarySearch(array []uint, target uint) int] -> %d != 2", index)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearch(array[:], uint(rand.Intn(7)))
	}
}

func TestBinarySearchRecursion(t *testing.T) {
	index := BinarySearchRecursion(array[:], 9)
	if index != 4 {
		t.Errorf("[func BinarySearchRecursion(array []uint, target uint) int] -> %d != 4", index)
	}
}

func BenchmarkBinarySearchRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearch(array[:], uint(rand.Intn(7)))
	}
}