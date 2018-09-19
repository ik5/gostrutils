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

// SearchString go over a string slice and look for the first occurence of a string (case sensitive). If found returns true
func SearchString(list []string, needle string) bool {
	f := func(i int) bool {
		return list[i] == needle
	}

	return sliceIndex(len(list), f) > -1
}
