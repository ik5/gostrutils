package gostrutils

import "testing"

func TestIsASCII(t *testing.T) {
	var inputs = []struct {
		input    string
		expected bool
	}{
		{
			"ABC", true,
		},
		{
			"אבג", false,
		},
		{
			"abc123 אבג", false,
		},
		{
			"", false,
		},
		{
			"abc 123 " + string(rune(130)), false,
		},
	}

	for _, item := range inputs {
		result := IsASCII(item.input)
		if result != item.expected {
			t.Errorf(
				"Have '%s'. Expected %t got %t",
				item.input, item.expected, result,
			)
		}
	}
}

func TestIsANSI(t *testing.T) {
	var inputs = []struct {
		input    string
		expected bool
	}{
		{
			"ABC", true,
		},
		{
			"אבג", false,
		},
		{
			"abc123 אבג", false,
		},
		{
			"", false,
		},
		{
			"abc 123 " + string(rune(130)), true,
		},
	}

	for _, item := range inputs {
		result := IsANSI(item.input)
		if result != item.expected {
			t.Errorf(
				"Have '%s'. Expected %t got %t",
				item.input, item.expected, result,
			)
		}
	}
}

func TestHaveNull(t *testing.T) {
	var inputs = []struct {
		input    string
		expected bool
	}{
		{
			"ABC", false,
		},
		{
			"אבג", false,
		},
		{
			"abc123 אבג", false,
		},
		{
			"", false,
		},
		{
			"abc 123 " + string(rune(130)), false,
		},
		{
			string(rune(0)), true,
		},
		{
			"abc 123" + string(rune(0)), true,
		},
		{
			string(rune(0)) + " abc 123", true,
		},
	}

	for _, item := range inputs {
		result := HaveNULL(item.input)
		if result != item.expected {
			t.Errorf(
				"Have '%s'. Expected %t got %t",
				item.input, item.expected, result,
			)
		}
	}

}
