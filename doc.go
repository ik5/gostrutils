/*
Package gostrutils contains string utilities that are missing from the main
strings package, as well as other encoding packages arrives by go.

The implementation of the package is set by files based on a subject that belong
to.


Basics

The basic logic of Go, is to use bytes as a non numeric holder.
The basics of the following package is to gain the ability to hold more support
for string based on functions that are missing, while remembering that Go's strings
are UTF-8.

The project itself built with files that holds the subject of what they are doing.

The helpers.go file, holds functions that are not string related functions, but
help creating that string support.

Test coverage

The aim of the following library is to have close to 100% of unit test coverage,
and also examples for all existed functions.

GSM 03.38

GSM 03.38 is a text encoding that is used with sending latin based languages with
SMS.

There are several types of encoding that SMS can support, where one of them is
GSM 03.38.

UTF16

UTF16 is 2 byte encoding mechanism. When a character is a single byte, it is
padded with NULL.

UTF16 contains a representation of characters either as Big or Little endian.
It means that bytes are either appears as `'A', 0x00"`, or `0x00, 'A'`.

In order to understand what type of encoding is in use, and it's endian, there
is a BOM -> Byte Order Mark header that provides such information.

Partial content of UTF16 does not a BOM, but a starting of such, should have one.
It must be the first (two) bytes of the 'string'.

At the past there was a UCS2 encoding that provided the same encoding as UTF16
Big Endian, and because of that, it was deprecated in favor of of UTF16 Big
Endian.

The current support for UTF16, provides three functions:
1. Encoding a go string that is UTF8 into UTF16.
2. Decoding a UTF16 to go's string (UTF8).
3. UTF16 BOM detection.

*/
package gostrutils
