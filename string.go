package gostrutils

import (
	"errors"
	"strings"
	"unicode/utf8"
)

// IsEmptyChars return true if after cleaning given chars the string is empty
func IsEmptyChars(s string, chars []rune) bool {
	toMapSet := make(map[rune]bool)
	for _, ch := range chars {
		toMapSet[ch] = true
	}

	result := strings.TrimFunc(s, func(ch rune) bool {
		_, found := toMapSet[ch]
		return found
	})

	return result == ""
}

// IsEmpty returns true if a string with whitespace only was provided or an
// empty string
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// Truncate a string to smaller parts, UTF8 multi-byte wise.
//
// When using s[:3] to truncate to 3 chars, it will not take runes but rather
// bytes. A multi-byte char (rune) can contain between one to 4 bytes, but not
// all of them will be returned.
//
// The following function will return the number of runes, and not the number of
// bytes inside a string.
func Truncate(s string, length int) string {
	if s == "" {
		return s
	}

	runeLength := utf8.RuneCountInString(s)
	if length >= runeLength {
		return s
	}

	result := []rune{}
	for len(result) != length {
		r, size := utf8.DecodeRuneInString(s)
		result = append(result, r)
		s = s[size:]
	}

	return string(result)
}

// CopyRange returns a copy of a string based on start and end of a rune
// for multi-byte chars instead of byte based chars (ASCII).
// 'to' is the amount of chars interested plus one.
// for example, for a string 'foo' extracting 'oo' by from: 1 and to 3.
func CopyRange(src string, from, to int) (string, error) {
	if src == "" {
		return src, nil
	}
	tmp := []rune(src)

	runeLength := len(tmp)
	if to <= 0 {
		to = runeLength
	}

	if from >= runeLength {
		return "", errors.New("from is larger then length")
	}

	if to > runeLength+1 {
		return "", errors.New("to is bigger then length")
	}

	if (from + to) > runeLength+2 {
		return "", errors.New("from + to is out of range")
	}

	return string(tmp[from:to]), nil
}
