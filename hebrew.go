package gostrutils

import "regexp"

var (
	hebrewPunchRegex = regexp.MustCompile("[\\x{591}-\\x{5C7}]")
)

// ClearHebrewPunctuationPoints removes Hebrew based punctuation points (Nikud)
// from a string. If something went wrong, the original string is returned.
func ClearHebrewPunctuationPoints(s string) string {
	return hebrewPunchRegex.ReplaceAllString(s, "")
}
