package binary_search

/*
BinarySearch

O(log n)
log2(8)=3 attempts required to find value (operation opposite to exponentiation).

Array must be sortable before.

1. Define middle item.
2. If middle item less than value then search will be in left side of array.
3. If middle item greater than value then search will be in right side of array.
4. Repeat steps 1,2,3.
*/
func BinarySearch(array []uint, target uint) int {

	var bottom uint = 0
	var top = uint(len(array) - 1)

	for bottom < top {

		var middle = bottom + top

		if array[middle] == target {
			return int(middle)
		}

		if array[middle] > target {
			top = middle - 1
		}

		if array[middle] < target {
			bottom = middle + 1
		}
	}

	return -1
}

func BinarySearchRecursion(array []uint, target uint) int {

	var bottom uint = 0
	var top = uint(len(array) - 1)
	var middle = -1

	if bottom < top {
		middle = int(bottom + top)

		if array[middle] == target {
			return middle
		}

		if array[middle] > target {
			top = uint(middle - 1)
		}

		if array[middle] < target {
			bottom = uint(middle + 1)
		}

		middle = BinarySearchRecursion(array[bottom:top+1], target)
	}

	return middle
}
