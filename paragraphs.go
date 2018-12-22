package gostrutils

// DefaultEllipse is the default char for marking abbreviation
const DefaultEllipse = "…"

// Abbreviate takes 'str' and based on length returns an abbreviate string
// with marks that represents middle of words.
//
// str - The original string
// startAt - Where to start to abbreviate inside a string
// maxLen - The maximum length for a string. It must be at least 4 chatrs
//
//
func Abbreviate(str string, startAt, maxLen int, abbrvChar string) string {
	if str == "" {
		return ""
	}

	if maxLen == 0 {
		return ""
	}

	if abbrvChar == "" {
		abbrvChar = DefaultEllipse
	}

	abbrvCharLen := len(abbrvChar)
	length := len(str)
	if length <= 4 {
		return str
	}

	if maxLen <= abbrvCharLen {
		return str[0:maxLen]
	}

	if length <= maxLen {
		return str
	}

	if startAt > length {
		startAt = length
	}

	if length-startAt < (maxLen - abbrvCharLen) {
		startAt = length - (maxLen - abbrvCharLen)
	}

	if startAt <= 4 {
		return str[0:maxLen-abbrvCharLen-1] + abbrvChar + str[length-1:]
	}
	if (startAt + maxLen - abbrvCharLen) < length {
		abrevStr := Abbreviate(str[startAt:length], (maxLen - abbrvCharLen + 1), maxLen-abbrvCharLen-1, abbrvChar)
		return str[0:1] + abbrvChar + abrevStr
	}
	return str[0:1] + abbrvChar + str[(length-(maxLen-abbrvCharLen)+1):length]
}
