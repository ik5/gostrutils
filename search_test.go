package gostrutils

import "testing"

func TestSearchFound(t *testing.T) {
	list := genList()

	found := SearchStringInSlice(list, "hello")
	if !found {
		t.Error("Unable to find 'hello' on the slice")
	}
}

func TestSearchNotFound(t *testing.T) {
	list := genList()

	found := SearchStringInSlice(list, "planet")
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
