package gostrutils

import (
	"fmt"
	"testing"
)

func TestUTF8ToGsm0338(t *testing.T) {
	s := "Hello World"
	gsm := UTF8ToGsm0338(s)
	utf8 := GSM0338ToUTF8(gsm)
	fmt.Printf("word before: %s\nword after gsm: %s\nword after utf8: %s\n", s, gsm, utf8)
}
