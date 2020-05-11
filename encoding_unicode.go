package gostrutils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unicode/utf16"
	"unicode/utf8"
)

// BOM Types
const (
	BOMNone = iota
	BOMBE
	BOMLE
)

// EncodeUTF16 get a utf8 string and translate it into a slice of bytes
func EncodeUTF16(s string, bigEndian, addBom bool) []byte {
	r := []rune(s)
	iresult := utf16.Encode(r)
	bytes := []byte{}
	if addBom {
		if bigEndian {
			bytes = append(bytes, 254, 255)
		} else {
			bytes = append(bytes, 255, 254)
		}
	}
	for _, i := range iresult {
		temp := make([]byte, 2)
		if bigEndian {
			binary.BigEndian.PutUint16(temp, i)
		} else {
			binary.LittleEndian.PutUint16(temp, i)
		}
		bytes = append(bytes, temp...)
	}
	return bytes
}

// DecodeUTF16 get a slice of bytes and decode it to UTF-8
func DecodeUTF16(b []byte) (string, error) {

	if len(b)%2 != 0 {
		return "", fmt.Errorf("Must have even length byte slice")
	}

	bom := UTF16Bom(b)
	if bom < 0 {
		return "", fmt.Errorf("Buffer is too small")
	}

	u16s := make([]uint16, 1)
	ret := &bytes.Buffer{}
	b8buf := make([]byte, 4)
	lb := len(b)

	// if there is BOM, we start at 2, otherwise at the beginning
	start := 2
	if bom == BOMNone {
		start = 0
	}
	for i := start; i < lb; i += 2 {
		//assuming bom is big endian if 0 returned
		if bom == BOMNone || bom == BOMBE {
			u16s[0] = uint16(b[i+1]) + (uint16(b[i]) << 8)
		} else if bom == BOMLE {
			u16s[0] = uint16(b[i]) + (uint16(b[i+1]) << 8)
		}
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		_, err := ret.Write([]byte(string(b8buf[:n])))
		if err != nil {
			return "", err
		}
	}

	return ret.String(), nil
}

// UTF16Bom returns 0 for no BOM, 1 for Big Endian and 2 for little endian
// it will return -1 if b is too small for having BOM
func UTF16Bom(b []byte) int8 {
	if len(b) < 2 {
		return -1
	}

	if b[0] == 0xFE && b[1] == 0xFF {
		return BOMBE
	}

	if b[0] == 0xFF && b[1] == 0xFE {
		return BOMLE
	}

	return BOMNone
}

// HexToUTF16Runes takes a hex based string and converts it into UTF16 runes
//
// Such string looks like:
//
//    "\x00H\x00e\x00l\x00l\x00o\x00 \x00W\x00o\x00r\x00l\x00d"
//    "\x05\xe9\x05\xdc\x05\xd5\x05\xdd\x00 \x05\xe2\x05\xd5\x05\xdc\x05\xdd"
//
// Result of the first string is (big endian):
//   [U+0048 'H' U+0065 'e' U+006C 'l' U+006C 'l' U+006F 'o' U+0020 ' ' U+0057 'W' U+006F 'o' U+0072 'r' U+006C 'l' U+0064 'd'] Hello World
//
// Result of the second string is (big endian):
//   [U+05E9 'ש' U+05DC 'ל' U+05D5 'ו' U+05DD 'ם' U+0020 ' ' U+05E2 'ע' U+05D5 'ו' U+05DC 'ל' U+05DD 'ם'] שלום עולם
func HexToUTF16Runes(s string, bigEndian bool) []rune {
	var chars []byte
	var position int
	length := len(s)

	// extract bytes from string
	for {
		if position >= length {
			break
		}
		if s[position] == '\\' {
			position++
			if position == length {
				// Just save the slash and nothing more is needed
				chars = append(chars, s[position-1])
				break
			}
			// save the content of the slash
			chars = append(chars, s[position-1], s[position])
			position++
			continue
		}

		chars = append(chars, s[position])
		position++
		continue
	}

	var runes []rune
	position = 0
	length = len(chars)

	// convert bytes two runes
	for {
		if position >= length-1 {
			break
		}
		aByte := []byte{chars[position], chars[position+1]}
		var aRune uint16
		if bigEndian {
			aRune = binary.BigEndian.Uint16(aByte)
		} else {
			aRune = binary.LittleEndian.Uint16(aByte)
		}
		unicodeChar := utf16.Decode([]uint16{aRune})
		runes = append(runes, unicodeChar...)
		position += 2
	}

	return runes
}
