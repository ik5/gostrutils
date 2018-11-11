package gostrutils

import (
	"reflect"
	"testing"
)

func TestEncodeUTF16BigEndian(t *testing.T) {
	b := utf16ByteSliceBigEndian()
	result := EncodeUTF16("TEST㑬", true, true)
	if !reflect.DeepEqual(b, result) {
		t.Errorf("Result and expected result are different:\n%+v\n%+v", b, result)
	}
}

func TestEncodeUTF16LittleEndian(t *testing.T) {
	b := utf16ByteSliceLittleEndian()
	result := EncodeUTF16("TEST㑬", false, true)
	if !reflect.DeepEqual(b, result) {
		t.Errorf("Result and expected result are different:\n%+v\n%+v", b, result)
	}
}

func TestDecodeUTF16(t *testing.T) {
	b := utf16ByteSliceLittleEndian()
	utf8, err := DecodeUTF16(b)
	if err != nil {
		t.Errorf("Error decoding: %s", err)
	}

	if utf8 != "TEST㑬" {
		t.Errorf("Expected 'TEST㑬', got: '%s'", utf8)
	}
}

func TestDecodeUTF16BigEndian(t *testing.T) {
	b := []byte{0xff, 0xfe, 'A', 0x00}
	str, err := DecodeUTF16(b)
	if str != "A" {
		t.Errorf("Invalid string result, expected 'A' string, got: '%s'", str)
	}

	if err != nil {
		t.Errorf("Have error: %s", err)
	}
}

func TestUTF16BomInvalid(t *testing.T) {
	b := []byte{0xff}
	result := UTF16Bom(b)
	if result >= 0 {
		t.Errorf("Expected '-1', got '%d'", result)
	}
}

func TestUTF16BomBigEndian(t *testing.T) {
	b := []byte{0xfe, 0xff}
	result := UTF16Bom(b)
	if result != 1 {
		t.Errorf("Expected '1' got '%d'", result)
	}
}

func TestUTF8BomLittleEndian(t *testing.T) {
	b := []byte{0xff, 0xfe}
	result := UTF16Bom(b)
	if result != 2 {
		t.Errorf("Expected '2' got '%d'", result)
	}
}

func TestUTF8BomNoBOM(t *testing.T) {
	b := []byte{'A', 0x00, 'B', 0x00}
	result := UTF16Bom(b)
	if result != 0 {
		t.Errorf("Expected '0', got '%d'", result)
	}
}

func utf16ByteSliceLittleEndian() []byte {
	return []byte{
		0xff, // BOM
		0xfe, // BOM
		'T',
		0x00,
		'E',
		0x00,
		'S',
		0x00,
		'T',
		0x00,
		0x6C,
		0x34,
	}
}

func utf16ByteSliceBigEndian() []byte {
	return []byte{
		0xfe, // BOM
		0xff, // BOM
		0x00,
		'T',
		0x00,
		'E',
		0x00,
		'S',
		0x00,
		'T',
		0x34,
		0x6C,
	}

}
