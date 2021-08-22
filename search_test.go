package gostrutils

import "testing"

func TestGetStringIndexInSliceFound(t *testing.T) {
	list := genList()
	idx := GetStringIndexInSlice(list, "hello")
	if idx == -1 {
		t.Error("Unable to find 'hello' in slice")
	}
}

func TestGetStringIndexInSliceNotFound(t *testing.T) {
	list := genList()
	idx := GetStringIndexInSlice(list, "planet")
	if idx != -1 {
		t.Error("Unable to find 'hello' on the slice")
	}
}

func TestSearchFound(t *testing.T) {
	list := genList()

	found := IsStringInSlice(list, "hello")
	if !found {
		t.Error("Unable to find 'hello' on the slice")
	}
}

func TestSearchNotFound(t *testing.T) {
	list := genList()

	found := IsStringInSlice(list, "planet")
	if found {
		t.Error("Able to find 'planet' at the list")
	}
}

func genList() []string {
	list := make([]string, 2)
	list[0] = "hello"
	list[1] = "world"

	return list
}

func TestGetBytesRuneIndexInSlice(t *testing.T) {
	type checkList struct {
		str      []byte
		needle   rune
		expected int
	}

	validList := []checkList{
		{
			str:      []byte("hello"),
			needle:   'o',
			expected: 4,
		},
		{
			str:      []byte("שלום עולם"),
			needle:   'ע',
			expected: 5,
		},
		{
			str:      []byte("hello"),
			needle:   'w',
			expected: -1,
		},
		{
			str:      []byte("שלום עולם"),
			needle:   'ז',
			expected: -1,
		},
	}

	t.Run("validList", func(t2 *testing.T) {
		for _, item := range validList {
			result := GetBytesRuneIndexInSlice(item.str, item.needle)

			if result != item.expected {
				t2.Errorf("'%s'['%U'] expected to be at %d but it was located at %d",
					item.str, item.needle, item.expected, result,
				)
			}
		}
	})
}

func TestIsRuneInByteSlice(t *testing.T) {
	type checkList struct {
		str      []byte
		needle   rune
		expected bool
	}

	validList := []checkList{
		{
			str:      []byte("hello"),
			needle:   'o',
			expected: true,
		},
		{
			str:      []byte("שלום עולם"),
			needle:   'ע',
			expected: true,
		},
		{
			str:      []byte("hello"),
			needle:   'w',
			expected: false,
		},
		{
			str:      []byte("שלום עולם"),
			needle:   'ז',
			expected: false,
		},
	}

	t.Run("validList", func(t2 *testing.T) {
		for _, item := range validList {
			result := IsRuneInByteSlice(item.str, item.needle)

			if result != item.expected {
				t2.Errorf("'%s'['%U'] expected to be at %T but %T returned",
					item.str, item.needle, item.expected, result,
				)
			}
		}
	})

}
