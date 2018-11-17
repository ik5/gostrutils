package gostrutils

import "unicode"

// CamelCaseToUnderscode converts CamelCase to camel_case
func CamelCaseToUnderscode(str string) string {
	var result string
	for idx, ch := range str {
		// first letter will just be lowered
		if idx == 0 {
			result = string(unicode.ToLower(ch))
			continue
		}

		// anywhere else
		if unicode.IsUpper(ch) {
			result = result + "_" + string(unicode.ToLower(ch))
			continue
		}

		// nothing to see here, just accept it
		result = result + string(ch)
	}

	return result
}
