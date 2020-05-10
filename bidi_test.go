package gostrutils

import (
	"testing"

	"golang.org/x/text/language"
)

func TestCountBidiChars(t *testing.T) {
	type wordsInput struct {
		desc     string
		input    string
		expected int
	}

	list := []wordsInput{
		{
			desc:     "Testing English Alphabet",
			input:    "abcdefghijklmnopqrstuvwxyz",
			expected: 0,
		},
		{
			desc:     "Hebrew with punctuation points",
			input:    "עִבְרִית",
			expected: 8,
		},
		{
			desc:     "Arabic with punctuation points",
			input:    " اَلْعَرَبِيَّةُ",
			expected: 15,
		},
		{
			desc:     "Test Hebrew with Arabic, no punctuation points",
			input:    "עברית العربية",
			expected: 12,
		},
		{
			desc:     "Triple language",
			input:    "English, עברית ,العربية",
			expected: 12,
		},
	}

	for idx, words := range list {
		result := CountBidiChars(words.input)

		if result != words.expected {
			t.Errorf("Test #%d. Got: %+v, got: %d", idx, words, result)
		}
	}
}

func TestCountStartBidiLanguage(t *testing.T) {
	type wordsInput struct {
		desc     string
		input    string
		expected int
	}

	list := []wordsInput{
		{
			desc:     "English",
			input:    "a",
			expected: 0,
		},
		{
			desc:     "Symbols",
			input:    "!@#$",
			expected: -1,
		},
		{
			desc:     "Latin and symbols",
			input:    "a12",
			expected: 0,
		},
		{
			desc:     "Hebrew, space and Latin",
			input:    "עברית a",
			expected: 6,
		},
		{
			desc:     "Latin and Hebrew",
			input:    "a עברית",
			expected: 0,
		},
		{
			desc:     "Number, space and Hebrew",
			input:    "100 מאה",
			expected: 7,
		},
		{
			desc:     "Number, space, Hebrew, space and Latin",
			input:    "100 מאה hundred",
			expected: 8,
		},
	}

	for idx, words := range list {
		result := CountStartBidiLanguage(words.input)

		if result != words.expected {
			t.Errorf("Test #%d. Got: %+v, got: %d", idx, words, result)
		}
	}

}

func TestTagIsBidi(t *testing.T) {
	type tagsToTest struct {
		input    language.Tag
		expected bool
	}

	list := []tagsToTest{
		{
			input:    language.Afrikaans,
			expected: false,
		},
		{
			input:    language.Arabic,
			expected: true,
		},
		{
			input:    language.Persian,
			expected: true,
		},
		{
			input:    language.Hebrew,
			expected: true,
		},
		{
			input:    language.Urdu,
			expected: true,
		},
	}

	for idx, tag := range list {
		result := TagIsBidi(tag.input)
		if result != tag.expected {
			t.Errorf("Test %d. Tag: %+v, got %t", idx, tag, result)
		}
	}
}
