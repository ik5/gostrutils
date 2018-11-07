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

func ExampleGSM0338ToUTF8() {
	msg := gostrutils.GSM0338ToUTF8("Please email me to user\x00example.com")
	fmt.Println("Got sms message: ", msg)
}

func ExampleEncodeUTF16() {
	str := gostrutils.EncodeUTF16("Windows Unicode String", true, true)
	fmt.Println(gostrutils.ByteToStr(str))
}

func ExampleDecodeUTF16() {
	// Big endian text of TEST
	text := []byte{
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
	}
	test, err := gostrutils.DecodeUTF16(text)
	if err != nil {
		panic(err)
	}
	fmt.Println("we have 'TEST'? ", test)
}
