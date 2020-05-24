package gostrutils

import "testing"

func TestIsEmptyFuncs(t *testing.T) {
	type toCheck struct {
		desc     string
		input    string
		chars    string
		expected bool
	}
	result := t.Run("IsEmptyChars", func(t *testing.T) {
		input := []toCheck{
			{
				desc:     "input of 'a', char '-', expected: false",
				input:    "aaaaaa",
				chars:    "-",
				expected: false,
			},
			{
				desc:     "input of '-', char '-', expected: true",
				input:    "-----------",
				chars:    "-",
				expected: true,
			},
			{
				desc:     "input: empty string, chars: ' ', expected: true",
				input:    "",
				chars:    " ",
				expected: true,
			},
			{
				desc:     "input of '-a', chars: '-', expected: false",
				input:    "-------------a",
				chars:    "-",
				expected: false,
			},
			{
				desc:     "input of '-aaaaa' chars: '-a', expected: true",
				input:    "------------a",
				chars:    "-a",
				expected: true,
			},
		}

		for idx, test := range input {
			result := IsEmptyChars(test.input, []rune(test.chars))
			if result != test.expected {
				t.Errorf("Tested #%d %+v got %t\n", idx, test, result)
			}
		}
	})

	if !result {
		t.Errorf("Invalid IsEmptyChars test")
	}

	result = t.Run("IsEmpty", func(t *testing.T) {
		input := []toCheck{
			{
				desc:     "input 'whitespace' expected: true",
				input:    "\r\n      \t",
				expected: true,
			},
			{
				desc:     "input: ' ', expected: true",
				input:    "         ",
				expected: true,
			},
			{
				desc:     "input 'a', expected: false",
				input:    "aaaaaaaa",
				expected: false,
			},
			{
				desc:     "input ' a' expected: false",
				input:    "        aaaaaaaaaaaa",
				expected: false,
			},
		}

		for idx, test := range input {
			result := IsEmpty(test.input)
			if result != test.expected {
				t.Errorf("Tested #%d %+v got %t", idx, test, result)
			}
		}

	})

	if !result {
		t.Errorf("Invalid IsEmptyChars test")
	}
}

func TestTruncate(t *testing.T) {
	type toCheck struct {
		input    string
		length   int
		expected string
	}

	tests := []toCheck{
		{
			input:    "",
			length:   0,
			expected: "",
		},
		{
			input:    "Hello World",
			length:   100,
			expected: "Hello World",
		},
		{
			input:    "simple English",
			length:   4,
			expected: "simp",
		},
		{
			input:    "עברית קשה שפה",
			length:   5,
			expected: "עברית",
		},
		{
			input:    "מעורב mixed",
			length:   6,
			expected: "מעורב ",
		},
		{
			input:    "mixed מעורב",
			length:   5,
			expected: "mixed",
		},
	}

	for idx, test := range tests {
		result := Truncate(test.input, test.length)

		if test.expected != result {
			t.Errorf("test %d '%s' (%d): expected '%s'[%d], got :'%s' %x",
				idx, test.input, len(test.input),
				test.expected, test.length, result, result)
		}
	}
}
