package gostrutils

import "testing"

func TestClearHebrewPunctuationPointsValid(t *testing.T) {
	hebPunch := "הַפְנָיָה"
	hebNonPunch := "הפניה"

	result := ClearHebrewPunctuationPoints(hebPunch)
	if result != hebNonPunch {
		t.Errorf("Expected '%s' to be '%s', have '%s'", hebPunch, hebNonPunch, result)
	}
}

func TestClearHebrewPunctuationPointsValid2(t *testing.T) {
	hebNonPunch := "הפניה"
	hebNonPunch2 := "הפניה"

	result := ClearHebrewPunctuationPoints(hebNonPunch)
	if result != hebNonPunch2 {
		t.Errorf("Expected '%s' to be '%s', have '%s'", hebNonPunch, hebNonPunch2, result)
	}
}

func TestClearHebrewPunctuationPointsValid3(t *testing.T) {
	eng1 := "hello"
	eng2 := "hello"

	result := ClearHebrewPunctuationPoints(eng1)
	if result != eng2 {
		t.Errorf("Expected '%s' to be '%s', have '%s'", eng1, eng2, result)
	}
}
