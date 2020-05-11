package gostrutils

import "testing"

func TestAbbreviateEmpty(t *testing.T) {
	str := Abbreviate("", 0, 20, "")
	if str != "" {
		t.Errorf("Expected empty string, got '%s'", str)
	}
}

func TestAbbreviateFour(t *testing.T) {
	str := Abbreviate("four", 0, 20, "")
	if str != "four" {
		t.Errorf("Expected the return of 'four', gor '%s'", str)
	}
}

func TestAbbreviateAbbrvChar(t *testing.T) {
	str := Abbreviate("12345", 0, 5, ".....")
	if str != "12345" {
		t.Errorf("Expected '12345', got %s", str)
	}
}

func TestAbbreviateValidLength(t *testing.T) {
	str := Abbreviate("1234567890", 0, 10, "")
	if str != "1234567890" {
		t.Errorf("Expected '1234567890', got '%s'", str)
	}
}

func TestAbbreviateOffestLength(t *testing.T) {
	str := Abbreviate("1234567890", 11, 10, "")
	if str != "1234567890" {
		t.Errorf("Expected '1234567890', got '%s'", str)
	}
}

func TestAbbreviateOffestLowerThenFour(t *testing.T) {
	str := Abbreviate("1234567890", 0, 5, "")
	if str != "1…0" {
		t.Errorf("Expected '1…0', got '%s'", str)
	}
}

func TestAbbreviateOffestHigherThenFour(t *testing.T) {
	str := Abbreviate("1234567890", 5, 5, "")
	if str != "1…6" {
		t.Errorf("Expected '1…6', got '%s'", str)
	}
}

func TestAbbreviateMaxLength(t *testing.T) {
	str := Abbreviate("1234567890", 1, 0, "")
	if str != "" {
		t.Errorf("Expected empty string, got '%s'", str)
	}
}

func TestAbbreviateStartLength(t *testing.T) {
	str := Abbreviate("1234567890", 12, 5, DefaultEllipse)
	expected := "1…0"

	if str != expected {
		t.Errorf("Expected '%s' got '%s'", expected, str)
	}
}
