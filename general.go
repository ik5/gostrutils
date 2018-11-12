package gostrutils

import "net/http"

// DetectMimeTypeFromContent takes a slice of bytes, and try to detect the
// content type that is in use.
func DetectMimeTypeFromContent(content []byte) (contentType string) {
	contentType = http.DetectContentType(content)
	return
}

// StrToPointer returns the address of string.
//
// The function can help when there is a need to convert literal string yo be
// have a pointer for that content.
func StrToPointer(str string) *string {
	return &str
}

// PointerToStr converts a pointer string to a normal string
//
// If there is a need for a string to be nil, then usually there is a pointer
// for a string involved. The current function converts nil pointer to empty
// string, or return the string itself
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
