package gostrutils

import (
	"bytes"
	"testing"
)

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

func TestCopyRange(t *testing.T) {
	type toCheck struct {
		input     string
		from      int
		to        int
		expected  string
		haveError bool
		error     string
	}
	tests := []toCheck{
		{
			input:     "foo",
			from:      3,
			to:        1,
			haveError: true,
			error:     "from is larger then length",
		},
		{
			input:     "foo",
			from:      0,
			to:        5,
			haveError: true,
			error:     "to is bigger then length",
		},
		{
			input:     "foo",
			from:      2,
			to:        4,
			haveError: true,
			error:     "from + to is out of range",
		},
		{
			input:    "foo",
			from:     1,
			to:       3,
			expected: "oo",
		},
		{
			input:    "עברית",
			from:     0,
			to:       3,
			expected: "עבר",
		},
		{
			input:    "עברית",
			from:     1,
			to:       4,
			expected: "ברי",
		},
		{
			// start with 1
			input:    "1עברית",
			from:     1,
			to:       6,
			expected: "עברית",
		},
		{
			// end with 1
			input:    "עברית1",
			from:     0,
			to:       5,
			expected: "עברית",
		},
		{
			input:    "‏1עברית",
			from:     2,
			to:       0,
			expected: "עברית",
		},
	}

	for idx, test := range tests {
		result, err := CopyRange(test.input, test.from, test.to)
		if err != nil {
			if !test.haveError {
				t.Errorf("Unexpected err: %s", err)
				continue
			}
			if err.Error() != test.error {
				t.Errorf("Expected error %d: '%s', got '%s'", idx, test.error, err)
			}
			continue
		}

		if test.expected != result {
			t.Errorf("test %d '%s' (%d): expected '%s'[%d:%d], got :'%s' %x",
				idx, test.input, len(test.input),
				test.expected, test.from, test.to, result, result)
		}
	}

}

func TestKeepByteChars(t *testing.T) {
	type bytesToTest struct {
		buf      []byte
		toKeep   []byte
		expected []byte
	}

	validInputs := []bytesToTest{
		{
			buf:      []byte("42\n"),
			toKeep:   []byte("1234567890"),
			expected: []byte("42"),
		},
		{
			buf:      []byte("שלום עולם | Hello World"),
			toKeep:   []byte("אבגדהוזחטיךכלםמןנסעףפץצקרשת"),
			expected: []byte("שלוםעולם"),
		},
	}

	invalidInputs := []bytesToTest{
		{
			buf:      []byte("42\n"),
			toKeep:   []byte("\r\n"),
			expected: []byte("42"),
		},
		{
			buf:      []byte("שלום עולם | Hello World"),
			toKeep:   []byte("אבגדהוזחטיךכלםמןנסעףפץצקרשת"),
			expected: []byte("שלום עולם"),
		},
	}

	t.Run("valid inputs", func(t2 *testing.T) {
		for _, input := range validInputs {
			result := KeepByteChars(input.buf, input.toKeep)

			if !bytes.Equal(result, input.expected) {
				t2.Errorf("Input '%s' toKeep: '%s', expected: '%s', result: '%s'",
					input.buf, input.toKeep, input.expected, result,
				)
			}
		}
	})

	t.Run("invalid inputs", func(t2 *testing.T) {
		for _, input := range invalidInputs {
			result := KeepByteChars(input.buf, input.toKeep)

			if bytes.Equal(result, input.expected) {
				t2.Errorf("Input '%s' toKeep: '%s', expected: '%s', result: '%s'",
					input.buf, input.toKeep, input.expected, result,
				)
			}
		}

	})
}
