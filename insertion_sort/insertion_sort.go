package selection_sort

import (
	"github.com/oleksiivelychko/go-helpers/pop"
)

/*InsertionSort

O(n^2)

1. Find the smallest element in array.
2. Put the smallest element into new array.
3. Remove the smallest element from array.
*/
func InsertionSort(slice []uint) []uint {
	var newSlice = make([]uint, 3)

	var length = len(slice)
	for i := 0; i < length; i++ {
		smallestIndex := findSmallestIndex(slice)
		newSlice[i] = slice[smallestIndex]
		slice = pop.UintPop(slice, smallestIndex)
	}

	return newSlice
}

func findSmallestIndex(slice []uint) int {
	var smallestIndex int = 0
	var smallest = slice[0]

	for i := 1; i <= len(slice)-1; i++ {
		if slice[i] < smallest {
			smallestIndex = i
			smallest = slice[i]
		}
	}

	return smallestIndex
}
