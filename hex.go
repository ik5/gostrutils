package gostrutils

// IsHex returns true if ch contains a hexadecimal character as a value
func IsHex(ch byte) bool {
	return ('0' <= ch && ch <= '9') ||
		('a' <= ch && ch <= 'f') ||
		('A' <= ch && ch <= 'F')
}

// HexToDec convert hex decimal ASCII value. If ch is not in the range, the
// function will return 0
func HexToDec(ch byte) byte {
	switch { // Jump table optimization?
	case '0' <= ch && ch <= '9':
		return ch - '0'
	case 'a' <= ch && ch <= 'f':
		return ch - 'a' + 10
	case 'A' <= ch && ch <= 'F':
		return ch - 'A' + 10
	}
	return 0
}

// DecToHex convert decimal value to hexadecimal as long as the range is between
// 0 to 15. When value is bigger than 15, 0 will be returned.
func DecToHex(dec byte) byte {
	switch {
	case dec <= 9: // byte always start with 0
		return dec + '0'
	case 10 <= dec && dec <= 15:
		return dec + 'A' - 10
	}
	return 0
}
