package gostrutils

// sliceIndex returns a slice index based on a callback value
func sliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

// SearchStringInSlice looks for a string inside a slice
func SearchStringInSlice(list []string, needle string) bool {
	f := func(i int) bool {
		return list[i] == needle
	}

	return sliceIndex(len(list), f) > -1
}
