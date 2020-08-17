package gostrutils

import (
	"testing"
)

func TestIsValidUUID(t *testing.T) {
	type fixture struct {
		input    string
		expected bool
	}

	fixtures := []fixture{
		{
			input:    "51e5fe10-5325-4a32-bce8-7ebe9708c453",
			expected: true,
		},
		{
			input:    "50ee7e88-c29e-4e38-b94b-6937b5a4c8c ",
			expected: false,
		},
		{
			input:    "aeb11973-3be9-47fc-8059-ecdac9ab1887",
			expected: true,
		},
		{
			input:    "abcdefgh-ijkl-mnop-qrst-uvwxyzabcdef",
			expected: false,
		},
		{
			input:    "688831FF-D2D6-4DA9-8676-AC4C9436BC0E",
			expected: true,
		},
		{
			input:    "b74d4ec0-64ab-43ee-b166-fb6db4c6dfb2",
			expected: true,
		},
	}

	for idx, test := range fixtures {
		result := IsValidUUID(test.input)
		if result != test.expected {
			t.Errorf("Expected (%d) %t but got %t", idx, test.expected, result)
		}
	}
}
