package gostrutils

// CountBidiChars counts how many chars belong to a bi-directional language:
// Arabic (and sub language), Hebrew, Urdu and Yiddish
func CountBidiChars(s string) int {
	result := 0

	for _, ch := range s {
		if // Hebrew
		(ch >= 0x590 && ch <= 0x5FF) ||
			// Arabic
			(ch >= 0x600 && ch <= 0x6FF) ||
			(ch >= 0x750 && ch <= 0x77F) ||
			(ch >= 0x8A0 && ch <= 0x8FF) ||
			// Syriac
			(ch >= 0x700 && ch <= 0x74F) ||
			// Thaana
			(ch >= 0x780 && ch <= 0x7BF) ||
			// NKO and Samaritan
			(ch >= 0x7C0 && ch <= 0x83E) ||
			(ch >= 0x840 && ch <= 0x85E) {
			result++
		}
	}

	return result
}

// CountStartBidiLanguage returns number of chars from the beginning.
// If chars are not A bidi range, and not up to 0x40 char, then 0 will be
// returned.
// If only chars up to (including) char 0x40 (@) -1 will be returned
//
// Important: result is both bidi chars and up to 0x40 count from the beginning
func CountStartBidiLanguage(s string) int {
	result := 0
	nunLatin := 0

	for _, ch := range s {
		if // Hebrew
		(ch >= 0x590 && ch <= 0x5FF) ||
			// Arabic
			(ch >= 0x600 && ch <= 0x6FF) ||
			(ch >= 0x750 && ch <= 0x77F) ||
			(ch >= 0x8A0 && ch <= 0x8FF) ||
			// Syriac
			(ch >= 0x700 && ch <= 0x74F) ||
			// Thaana
			(ch >= 0x780 && ch <= 0x7BF) ||
			// NKO and Samaritan
			(ch >= 0x7C0 && ch <= 0x83E) ||
			(ch >= 0x840 && ch <= 0x85E) {
			result++
			continue
		}

		if ch >= 0x00 && ch <= 0x40 {
			nunLatin++
			continue
		}
		break
	}

	if result == 0 && nunLatin > 0 {
		return -1
	}

	return result + nunLatin
}
