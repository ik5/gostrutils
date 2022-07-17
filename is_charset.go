package gostrutils

// IsASCII returns true if the entire string contains only ASCII characters.
// On empty string, the function returns false.
func IsASCII(str string) bool {
	if str == "" {
		return false
	}

	for _, ch := range str {
		if ch > 127 {
			return false
		}
	}

	return true
}

// IsANSI returns true if the entire string contains only ANSI characters.
// On empty string, the function returns false.
func IsANSI(str string) bool {
	if str == "" {
		return false
	}
	for _, ch := range str {
		if ch > 255 {
			return false
		}
	}

	return true
}

// HaveNULL look for null char, and return true if found at first place.
// Empty str returns false.
func HaveNULL(str string) bool {
	if str == "" {
		return false
	}

	for _, ch := range str {
		if ch == 0 {
			return true
		}
	}

	return false
}
