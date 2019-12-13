package gostrutils

import "regexp"

// ClearHebrewPunctuationPoints removes Hebrew based punctuation points (Nikud)
// from a string. If something went wrong, the original string is returned.
func ClearHebrewPunctuationPoints(s string) string {
	re, err := regexp.Compile("[\\x{591}-\\x{5C7}]")
	if err != nil {
		return s
	}

	return re.ReplaceAllString(s, "")
}
