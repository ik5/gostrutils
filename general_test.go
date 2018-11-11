package gostrutils

import (
	"strings"
	"testing"
)

func TestValidDetectMimeTypeFromContent(t *testing.T) {
	hello := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}
	mime := DetectMimeTypeFromContent(hello)
	if !strings.HasPrefix(mime, "text/plain") {
		t.Errorf("Bad mime type: %s", mime)
	}
}

func TestInvalidDetectMimeTypeFromContent(t *testing.T) {
	htmlHead := []byte{0x3C, 0x68, 0x65, 0x61, 0x64, 0x3e}
	mime := DetectMimeTypeFromContent(htmlHead)
	if strings.HasPrefix(mime, "text/plain") {
		t.Errorf("Bad mime type: %s", mime)
	}
}

func TestStrToPointer(t *testing.T) {
	str := "hello"
	pStr := StrToPointer(str)
	if pStr == nil {
		t.Error("pStr is nil")
	}

	if *pStr != str {
		t.Errorf("pStr (%s) is not equal to str (%s)", *pStr, str)
	}
}

func TestPointerToStrEmpty(t *testing.T) {
	var pStr *string
	str := PointerToStr(pStr)
	if str != "" {
		t.Errorf("Expected for empty str, found: '%s'", str)
	}
}

func TestPointerToStr(t *testing.T) {
	origStr := "Hello world"
	pStr := &origStr
	str := PointerToStr(pStr)
	if str != "Hello world" {
		t.Errorf("Expected for '%s', found: '%s'", origStr, str)
	}
}

func TestByteToStr(t *testing.T) {
	helloBytes := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}
	result := ByteToStr(helloBytes)

	if result != "hello" {
		t.Error("ByteToStr did not convert to 'hello'")
	}
}
