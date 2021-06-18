package gostrutils

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	intRegex         = regexp.MustCompile("^[0-9]+$")
	signedIntRegex   = regexp.MustCompile("^-?[0-9]+$")
	floatRegex       = regexp.MustCompile(`^[0-9]+\.[0-9]+$`)
	signedFloatRegex = regexp.MustCompile(`^-?[0-9]+\.[0-9]+$`)
)

// StrToInt64 convert a string to int64
func StrToInt64(str string, def int64) int64 {
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return def
	}
	return result
}

// StrToUInt64 convert string to uint64
func StrToUInt64(str string, def uint64) uint64 {
	result, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return def
	}
	return result
}

// UInt64Join takes a list of uint64 and join them with a separator
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
		s += strconv.FormatUint(item, 10)

		if i < length-1 {
			s += sep
		}
	}

	return s
}

// Int64Join takes a list of int64 and join them with a separator
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
		s += strconv.FormatInt(item, 10)

		if i < length-1 {
			s += sep
		}
	}

	return s
}

// Uin64Split get a string with separator and convert it to a slice of uint64
func Uin64Split(data, sep string) []uint64 {
	fields := strings.Split(data, sep)
	result := make([]uint64, len(fields))

	for i, elem := range fields {
		result[i], _ = strconv.ParseUint(elem, 10, 64)
	}

	return result
}

// In64Split get a string with separator and convert it to a slice of int64
func In64Split(data, sep string) []int64 {
	fields := strings.Split(data, sep)
	result := make([]int64, len(fields))

	for i, elem := range fields {
		result[i], _ = strconv.ParseInt(elem, 10, 64)
	}

	return result
}

// ToFloat32Default convert string to float32 without errors. If error returns, defaultValue is set instead.
func ToFloat32Default(field string, defaultValue float32) float32 {
	result, err := strconv.ParseFloat(field, 32)
	if err != nil {
		return defaultValue
	}

	return float32(result)
}

// ToFloat32 convert string to float32 without errors!
func ToFloat32(field string) float32 {
	return ToFloat32Default(field, 0.0)
}

// ToFloat6Default convert string to float64 without errors. If error returns, default is set instead.
func ToFloat6Default(field string, defaultValue float64) float64 {
	result, err := strconv.ParseFloat(field, 64)
	if err != nil {
		return defaultValue
	}

	return result
}

// ToFloat64 convert string to float64 without errors!
func ToFloat64(field string) float64 {
	return ToFloat6Default(field, 0.0)
}

// IsUInteger returns true if a string is unsigned integer
func IsUInteger(txt string) bool {
	return intRegex.MatchString(txt)
}

// IsInteger returns true if a string is an integer
func IsInteger(txt string) bool {
	return signedIntRegex.MatchString(txt)
}

// IsUFloat returns true if a string is unsigned floating point
func IsUFloat(txt string) bool {
	return floatRegex.MatchString(txt)
}

// IsFloat returns true if a given text is a floating point
func IsFloat(txt string) bool {
	return signedFloatRegex.MatchString(txt)
}

// IsUNumber returns true if a given string is unsigned integer or float
func IsUNumber(txt string) bool {
	return IsUInteger(txt) || IsUFloat(txt)
}

// IsNumber returns true if a given string is integer or float
func IsNumber(txt string) bool {
	return IsInteger(txt) || IsFloat(txt)
}

// IsInRange takes an integer range and look at the string that contains only
// numbers tom make sure it is inside the range
func IsInRange(min, max int64, src string) bool {
	if min > max { // no error so nothing to look for
		return false
	}

	if src == "" {
		return false
	}

	dest, err := strconv.ParseInt(src, 10, 64)
	if err != nil {
		return false
	}

	return dest >= min && dest <= max
}
