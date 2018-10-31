package gostrutils

import "testing"

func TestToSQLStringEmptyString(t *testing.T) {
	str := ToSQLString("")
	if str.Valid {
		t.Error("str.Valid cannot be true")
	}
}

func TestToSQLStringNotEmptyString(t *testing.T) {
	str := ToSQLString("Hello World")
	if !str.Valid {
		t.Error("str.Valid cannot be false")
	}
}

func TestToSQLStringNotNullableEmptyString(t *testing.T) {
	str := ToSQLStringNotNullable("")

	if !str.Valid {
		t.Error("str.Valid cannot be false")
	}
}

func TestToSQLStringNotNullableNoneEmptyString(t *testing.T) {
	str := ToSQLStringNotNullable("Hello World")

	if !str.Valid {
		t.Error("str.Valid cannot be false")
	}
}
