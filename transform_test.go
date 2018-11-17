package gostrutils

import (
	"testing"
)

func TestCamelCaseToUnderscodeValid(t *testing.T) {
	result := CamelCaseToUnderscode("CamelCase")
	if result != "camel_case" {
		t.Errorf("Expected 'camel_case', got : %s", result)
	}
}

func TestCamelCaseToUnderscodeNumeric(t *testing.T) {
	result := CamelCaseToUnderscode("one_1CamelCase")
	if result != "one_1_camel_case" {
		t.Errorf("Expected 'one_1_camel_case', got: '%s'", result)
	}
}
