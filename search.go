package gostrutils

// SearchStringInSlice looks for a string inside a slice
func SearchStringInSlice(list []string, needle string) bool {
	f := func(i int) bool {
		return list[i] == needle
	}

	return SliceIndex(len(list), f) > -1
}
