package gostrutils

import (
	"strings"
	"testing"
)

func TestSmsCalculateSmsFragmentsSingleLatin(t *testing.T) {
	count := SmsCalculateSmsFragments("Hello World")
	if count != 1 {
		t.Errorf("Expected '1', got '%d'", count)
	}
}

func TestSmsCalculateSmsFragmentsSingleUTF8(t *testing.T) {
	count := SmsCalculateSmsFragments("Hello World שלום עולם")
	if count != 1 {
		t.Errorf("Expected '1', got '%d'", count)
	}
}

func TestSmsCalculateSmsFragmentsTwoFragmentsLatin(t *testing.T) {
	msg := strings.Repeat("abc", 100)
	count := SmsCalculateSmsFragments(msg)
	if count != 2 {
		t.Errorf("Expected 2 messages, got '%d'", count)
	}
}

func TestSmsCalculateSmsFragmentsTwoFragmentsUtf8(t *testing.T) {
	msg := strings.Repeat("א", 100)
	count := SmsCalculateSmsFragments(msg)
	if count != 2 {
		t.Errorf("Expected 2 messages, got '%d'", count)
	}
}