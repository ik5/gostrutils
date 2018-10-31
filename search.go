package gostrutils

// GetStringIndexInSlice returns the index of the slice if string was found or -1 if not
func GetStringIndexInSlice(list []string, needle string) int {
	f := func(i int) bool {
		return list[i] == needle
	}

	return SliceIndex(len(list), f)
}

// IsStringInSlice looks for a string inside a slice and return true if it exists
func IsStringInSlice(list []string, needle string) bool {
	return GetStringIndexInSlice(list, needle) > -1
}
