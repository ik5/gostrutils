package gostrutils_test

import (
	"fmt"

	"github.com/ik5/gostrutils"
)

func ExampleDetectMimeTypeFromContent() {
	hello := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}

	fmt.Printf("Mime type of %s is: %s", hello, gostrutils.DetectMimeTypeFromContent(hello))
}

func ExampleStrToPointer() {
	hello := "hello"
	addr := gostrutils.StrToPointer(hello)
	fmt.Printf("String: %s, Address: %p\n", hello, addr)
}

func ExamplePointerToStr() {
	empty := gostrutils.PointerToStr(nil)
	hello := "hello"
	pHello := &hello
	newHello := gostrutils.PointerToStr(pHello)
	fmt.Printf("empty: %s, pHello: %s\n", empty, newHello)
}

func sendSmsMessage(msg string) bool {
	return true
}

func ExampleUTF8ToGsm0338() {
	msg := gostrutils.UTF8ToGsm0338("Please email me to user@example.com")
	// send the SMS message here
	if !sendSmsMessage(msg) {
		panic("Cannot send a message")
	}
}
