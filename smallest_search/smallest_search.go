package smallest_search

func SmallestSearch(slice []uint, smallest uint) uint {
	if len(slice) == 0 {
		return smallest
	}

	var current = slice[len(slice)-1]

	if smallest > current {
		smallest = current
	}

	return SmallestSearch(slice[:len(slice)-1], smallest)
}
