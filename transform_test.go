package gostrutils

import (
	"testing"
)

func TestCamelCaseToUnderscoreValid(t *testing.T) {
	result := CamelCaseToUnderscore("CamelCase")
	if result != "camel_case" {
		t.Errorf("Expected 'camel_case', got : %s", result)
	}
}

func TestCamelCaseToUnderscoreNumeric(t *testing.T) {
	result := CamelCaseToUnderscore("one_1CamelCase")
	if result != "one_1_camel_case" {
		t.Errorf("Expected 'one_1_camel_case', got: '%s'", result)
	}
}

func TestCamelCaseToJavascriptCase(t *testing.T) {
	result := CamelCaseToJavascriptCase("CamelCase")
	if result != "camelCase" {
		t.Errorf("Expected 'camelCase', got '%s'", result)
	}
}
