package gostrutils

import "strconv"

// StrToInt convert a string to int64
func StrToInt(str string, def int64) int64 {
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return def
	}
	return result
}

// StrToUInt convert string to uint64
func StrToUInt(str string, def uint64) uint64 {
	result, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return def
	}
	return result
}
