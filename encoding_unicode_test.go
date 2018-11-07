package gostrutils

import (
	"testing"
)
func TestDecodeUTF16(t *testing.T) {
	b := []byte{
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

	utf8, err := DecodeUTF16(b)
	if err != nil {
		t.Errorf("Error decoding: %s", err)
	}

	if utf8 != "TEST㑬" {
		t.Errorf("Expected 'TEST㑬', got: '%s'", utf8)
	}
}