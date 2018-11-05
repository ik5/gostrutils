package gostrutils

import "net/http"

// DetectMimeTypeFromContent takes a slice of bytes, and try to detect the content type that is in use
func DetectMimeTypeFromContent(content []byte) (contentType string) {
	contentType = http.DetectContentType(content)
	return
}

// StrToPointer returns the address of string.
// It is good use when you have a literal string that requires to be converted to pointer
func StrToPointer(str string) *string {
	return &str
}

// PointerToStr converts a nullable string to a normal string
func PointerToStr(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

// ByteToStr converts slice of bytes to string
func ByteToStr(b []byte) string {
	return string(b)
}
