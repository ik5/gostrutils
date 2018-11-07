package gostrutils

import "testing"

func TestIsIsraeliPhoneNumberValid(t *testing.T) {
	if !IsIsraeliPhoneNumber("031234567") {
		t.Error("Expected a valid Israeli phone number, but go false")
	}
}

func TestIsIsraeliPhoneNumberInvalid(t *testing.T) {
	if IsIsraeliPhoneNumber("5551234222") {
		t.Error("Expected a non valid Israeli phone number, but go true")
	}
}

func TestIsE164Valid(t *testing.T) {
	if !IsE164("+15551234222") {
		t.Error("Expected a valid e164 number, got false")
	}
}

func TestIsE164Invalid(t *testing.T) {
	if !IsE164("15551234222") {
		t.Error("Expected an invalid e164 number, got true")
	}
}

func TestCleanPhoneNumber(t *testing.T) {
	num := CleanPhoneNumber("+1-555-1234-222")
	if num != "15551234222" {
		t.Errorf("Expected '15551234222', got '%s'", num)
	}
}
