package gostrutils

import "bytes"

// GetStringIndexInSlice returns the index of the slice if string was found or -1 if not
func GetStringIndexInSlice(list []string, needle string) int {
	f := func(i int) bool {
		return list[i] == needle
	}

	return SliceIndex(len(list), f)
}

// GetBytesRuneIndexInSlice returns the index of the slice if rune was found or -1
// if not.
// Important, the index returned is of the rune, not of the byte!
func GetBytesRuneIndexInSlice(list []byte, needle rune) int {
	r := bytes.Runes(list)
	f := func(i int) bool {
		return r[i] == needle
	}

	return SliceIndex(len(r), f)
}

// IsStringInSlice looks for a string inside a slice and return true if it exists
func IsStringInSlice(list []string, needle string) bool {
	f := func(i int) bool {
		return list[i] == needle
	}

	return SliceIndex(len(list), f) > -1
}

// IsRuneInByteSlice return true if needle exists in list
func IsRuneInByteSlice(list []byte, needle rune) bool {
	r := bytes.Runes(list)
	f := func(i int) bool {
		return r[i] == needle
	}

	return SliceIndex(len(r), f) > -1
}
