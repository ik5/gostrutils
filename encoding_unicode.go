package gostrutils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unicode/utf16"
	"unicode/utf8"
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
	if bom == 0 {
		start = 0
	}
	for i := start; i < lb; i += 2 {
		//assuming bom is big endian if 0 returned
		if bom == 0 || bom == 1 {
			u16s[0] = uint16(b[i+1]) + (uint16(b[i]) << 8)
		} else if bom == 2 {
			u16s[0] = uint16(b[i]) + (uint16(b[i+1]) << 8)
		}
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		ret.Write([]byte(string(b8buf[:n])))
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
		return 1
	}

	if b[0] == 0xFF && b[1] == 0xFE {
		return 2
	}

	return 0
}
