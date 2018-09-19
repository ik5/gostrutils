package gostrutils

import (
	"math"
	"unicode/utf8"
)

// SmsCalculateSmsFragments calculate based on a string how many SMS messages will it take to deliver a massage
func SmsCalculateSmsFragments(message string) uint64 {
	mbLen := utf8.RuneCountInString(message)
	byeLen := len(message)
	maxMessageLength := 160
	maxMultiMessageLength := 153

	if mbLen != byeLen {
		maxMessageLength = 70
		maxMultiMessageLength = 67
	}

	if mbLen > maxMessageLength {
		return uint64(math.Ceil(float64(mbLen) / float64(maxMultiMessageLength)))
	}

	return 1
}
