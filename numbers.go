package gostrutils

import (
	"strconv"
	"strings"
)

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

// UInt64Join takes a list of uint64 and join them with a seperator
// based on https://play.golang.org/p/KpdkVS1B4s
func UInt64Join(list []uint64, sep string) string {
	length := len(list)

	if length == 0 {
		return ""
	}

	if length == 1 {
		return strconv.FormatUint(list[0], 10)
	}

	s := ""

	for i, item := range list {
		s = s + strconv.FormatUint(item, 10)

		if i < length-1 {
			s = s + sep
		}
	}

	return s
}

// Int64Join takes a list of int64 and join them with a seperator
// based on https://play.golang.org/p/KpdkVS1B4s
func Int64Join(list []int64, sep string) string {
	length := len(list)

	if length == 0 {
		return ""
	}

	if length == 1 {
		return strconv.FormatInt(list[0], 10)
	}

	s := ""

	for i, item := range list {
		s = s + strconv.FormatInt(item, 10)

		if i < length-1 {
			s = s + sep
		}
	}

	return s
}

// Uin64Split get a string with seperator and convert it to a slice of uint64
func Uin64Split(data, sep string) []uint64 {
	fields := strings.Split(data, sep)
	result := make([]uint64, len(fields))

	for i, elem := range fields {
		result[i], _ = strconv.ParseUint(elem, 10, 64)
	}

	return result
}

// In64Split get a string with seperator and convert it to a slice of int64
func In64Split(data, sep string) []int64 {
	fields := strings.Split(data, sep)
	result := make([]int64, len(fields))

	for i, elem := range fields {
		result[i], _ = strconv.ParseInt(elem, 10, 64)
	}

	return result
}
