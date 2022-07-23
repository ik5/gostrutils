package gostrutils

import "testing"

func TestIsHex(t *testing.T) {
	var toTest = []struct {
		value    byte
		expected bool
	}{
		{
			'a', true,
		},
		{
			'A', true,
		},
		{
			'0', true,
		},
		{
			'g', false,
		},
	}

	for _, item := range toTest {
		result := IsHex(item.value)
		if result != item.expected {
			t.Errorf(
				"Checked '%s', expected '%t' but got '%t'",
				string(item.value), item.expected, result,
			)
		}
	}
}

func TestHexToNumeric(t *testing.T) {
	var toTest = []struct {
		value    byte
		expected byte
	}{
		{'a', 10},
		{'A', 10},
		{'f', 15},
		{'F', 15},
		{'2', 2},
		{'0', 0},
		{'9', 9},
		{'z', 0},
		{'Z', 0},
		{'!', 0},
		{'@', 0},
	}

	for _, item := range toTest {
		result := HexToDec(item.value)
		if item.expected != result {
			t.Errorf(
				"Tested '%s', expected '%d', but got '%d'",
				string(item.value), item.expected, result,
			)
		}
	}
}

func TestDescToHex(t *testing.T) {
	var toTest = []struct {
		value    byte
		expected byte
	}{
		{10, 'A'},
		{10, 'A'},
		{15, 'F'},
		{15, 'F'},
		{2, '2'},
		{0, '0'},
		{9, '9'},
		{16, 0},
		{20, 0},
		{90, 0},
		{127, 0},
	}

	for _, item := range toTest {
		result := DecToHex(item.value)
		if item.expected != result {
			t.Errorf(
				"Tested '%d', expected '%s', but got '%s'",
				item.value, string(item.expected), string(result),
			)
		}
	}

}
