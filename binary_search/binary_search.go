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
func BinarySearch(slice []int, target int) (bool, int) {

	var bottom int = 0
	var top = len(slice) - 1

	for bottom < top {

		var middle = bottom + top

		if slice[middle] == target {
			return true, middle
		}

		if slice[middle] > target {
			top = middle - 1
		}

		if slice[middle] < target {
			bottom = middle + 1
		}
	}

	return false, 0
}

func BinarySearchRecursion(slice []int, target int) (bool, int) {

	var bottom int = 0
	var top = len(slice) - 1
	var middle = -1

	if bottom < top {
		middle = bottom + top

		if slice[middle] == target {
			return true, middle
		}

		if slice[middle] > target {
			top = middle - 1
		}

		if slice[middle] < target {
			bottom = middle + 1
		}

		_, middle = BinarySearchRecursion(slice[bottom:top+1], target)
	}

	return false, middle
}
