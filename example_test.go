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

func ExampleByteToStr() {
	b := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd'}
	s := gostrutils.ByteToStr(b)
	fmt.Println(s)
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

func ExampleStrToInt64() {
	meaning := gostrutils.StrToInt64("42", 0)
	fmt.Println("Meaning of life universe and everything: ", meaning)
}

func ExampleUInt64Join() {
	list := []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	fibonacci := gostrutils.UInt64Join(list, ", ")
	fmt.Printf("First 10 fibonacci sequence: %s\n", fibonacci)
}

func ExampleUin64Split() {
	fibonacci := "0, 1, 1, 2, 3, 5, 8, 13, 21, 34"
	list := gostrutils.Uin64Split(fibonacci, ", ")
	fmt.Printf("List: %+v\n", list)
}

func ExampleToFloat32Default() {
	pi := gostrutils.ToFloat32Default("3.141592653", 3.14)
	fmt.Printf("Pi: %f\n", pi)
}
