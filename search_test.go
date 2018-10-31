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
