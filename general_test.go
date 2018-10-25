package gostrutils

import (
	"strings"
	"testing"
)

func Test_ValidDetectMimeTypeFromContent(t *testing.T) {
	hello := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}
	mime := DetectMimeTypeFromContent(hello)
	if !strings.HasPrefix(mime, "text/plain") {
		t.Errorf("Bad mime type: %s", mime)
	}
}

func Test_InvalidDetectMimeTypeFromContent(t *testing.T) {
	htmlHead := []byte{0x3C, 0x68, 0x65, 0x61, 0x64, 0x3e}
	mime := DetectMimeTypeFromContent(htmlHead)
	if strings.HasPrefix(mime, "text/plain") {
		t.Errorf("Bad mime type: %s", mime)
	}
}

func Test_StrToPointer(t *testing.T) {
	str := "hello"
	pStr := StrToPointer(str)
	if pStr == nil {
		t.Error("pStr is nil")
	}

	if *pStr != str {
		t.Errorf("pStr (%s) is not equal to str (%s)", *pStr, str)
	}
}
