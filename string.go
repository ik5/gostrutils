package gostrutils

import (
	"strings"
)

// IsEmptyChars return true if after cleaning given chars the string is empty
func IsEmptyChars(s string, chars []rune) bool {
	toMapSet := make(map[rune]bool)
	for _, ch := range chars {
		toMapSet[ch] = true
	}

	result := strings.TrimFunc(s, func(ch rune) bool {
		_, found := toMapSet[ch]
		return !found
	})

	return result == ""
}

// IsEmpty returns true if a string with whitespace only was provided or an
// empty string
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}
