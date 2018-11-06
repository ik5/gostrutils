package gostrutils

import (
	"testing"
)

func TestUTF8ToGsm0338Valid(t *testing.T) {
	s := "Hello World"
	gsm := UTF8ToGsm0338(s)
	if gsm != "Hello World" {
		t.Errorf("Invalid gsm content: '%s', expected 'Hello World'", gsm)
	}
}

func TestUTF8ToGsm0338Valid2(t *testing.T) {
	s := "@"
	gsm := UTF8ToGsm0338(s)
	if gsm != "\x00" {
		t.Errorf("Invalid gsm, expected NULL, got: %s", gsm)
	}
}

func TestUTF8ToGsm0338Invalid(t *testing.T) {
	s := "$"
	gsm := UTF8ToGsm0338(s)
	if gsm == "$" {
		t.Errorf("Invalid gsm convertion, expected '\x22', got '%s' ", gsm)
	}
}

func TestGSM0338ToUTF8Valid(t *testing.T) {
	gsm := UTF8ToGsm0338("Hello World")
	utf8 := GSM0338ToUTF8(gsm)
	if utf8 != "Hello World" {
		t.Errorf("Invalid utf8: '%s'", utf8)
	}
}

func TestGSM0338ToUTF8Valid2(t *testing.T) {
	gsm := UTF8ToGsm0338("$")
	utf8 := GSM0338ToUTF8(gsm)
	if utf8 != "$" {
		t.Errorf("Invalid utf8, expected '$', got '%s'", utf8)
	}
}

func TestGSM0338ToUTF8Invalid(t *testing.T) {
	gsm := "\x00"
	utf8 := GSM0338ToUTF8(gsm)
	if utf8 == "\x00" {
		t.Errorf("Invalid utf8, expected '@' got %s", utf8)
	}
}