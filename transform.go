package gostrutils

import "unicode"

// CamelCaseToUnderscore converts CamelCase to camel_case
func CamelCaseToUnderscore(str string) string {
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
		result += string(ch)
	}

	return result
}

// CamelCaseToJavascriptCase convert CamelCase to camelCase
func CamelCaseToJavascriptCase(str string) string {
	var result string
	for idx, ch := range str {
		if idx == 0 {
			result = string(unicode.ToLower(ch))
			continue
		}

		result += string(ch)
	}
	return result
}
