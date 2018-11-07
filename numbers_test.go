package gostrutils

import (
	"reflect"
	"testing"
)

func TestStrToIntValid(t *testing.T) {
	result := StrToInt("-42", 2)
	if result != -42 {
		t.Errorf("Invalid result value, expected '-42', got %d", result)
	}
}

func TestStrToIntWithDefaultValue(t *testing.T) {
	result := StrToInt("a", -42)
	if result != -42 {
		t.Errorf("Invalid result value, expected '-42', got %d", result)
	}
}

func TestStrToUIntValid(t *testing.T) {
	result := StrToUInt("42", 2)
	if result != 42 {
		t.Errorf("Invalid result value, expected '42', got %d", result)
	}
}

func TestStrToUIntWithDefaultValue(t *testing.T) {
	result := StrToUInt("a", 42)
	if result != 42 {
		t.Errorf("Invalid result value, expected '42', got %d", result)
	}
}

func TestUInt64JoinSimple(t *testing.T) {
	var list []uint64
	list = append(list, 1, 2, 3, 4)
	str := UInt64Join(list, ",")
	if str != "1,2,3,4" {
		t.Errorf("Expected a string of '1,2,3,4', got: '%s'", str)
	}
}

func TestUInt64JoinEmpty(t *testing.T) {
	var list []uint64
	str := UInt64Join(list, ",")
	if str != "" {
		t.Errorf("Expected empty string, got '%s'", str)
	}
}

func TestUInt64JoinSingle(t *testing.T) {
	var list []uint64
	list = append(list, 1)
	str := UInt64Join(list, ",")
	if str != "1" {
		t.Errorf("Expected '1', got '%s'", str)
	}
}

func TestInt64JoinSimple(t *testing.T) {
	var list []int64
	list = append(list, -1, 2, 3, 4)
	str := Int64Join(list, ",")
	if str != "-1,2,3,4" {
		t.Errorf("Expected a string of '-1,2,3,4', got: '%s'", str)
	}
}

func TestInt64JoinEmpty(t *testing.T) {
	var list []int64
	str := Int64Join(list, ",")
	if str != "" {
		t.Errorf("Expected empty string, got '%s'", str)
	}
}

func TestInt64JoinSingle(t *testing.T) {
	var list []int64
	list = append(list, -1)
	str := Int64Join(list, ",")
	if str != "-1" {
		t.Errorf("Expected '-1', got '%s'", str)
	}
}

func TestUin64SplitSimple(t *testing.T) {
	list := Uin64Split("1,2,3,4", ",")
	var expected []uint64
	expected = append(expected, 1, 2, 3, 4)
	if !reflect.DeepEqual(list, expected) {
		t.Errorf("Invalid list was provided, expected list of: %+v, got %+v", expected, list)
	}
}

func TestIn64Split(t *testing.T) {
	list := In64Split("-1,2,3,4", ",")
	var expected []int64
	expected = append(expected, -1, 2, 3, 4)
	if !reflect.DeepEqual(list, expected) {
		t.Errorf("Invalid list was provided, expected list of: %+v, got %+v", expected, list)
	}
}

func TestToFloat32DefaultSimple(t *testing.T) {
	num := ToFloat32Default("3.14", 1.0)
	if num != 3.14 {
		t.Errorf("Expected '3.14', got %f", num)
	}
}

func TestToFloat32DefaultWithDefault(t *testing.T) {
	num := ToFloat32Default(".", 3.14)
	if num != 3.14 {
		t.Errorf("Expected '3.14', got %f", num)
	}
}

func TestToFloat64DefaultSimple(t *testing.T) {
	num := ToFloat6Default("3.14", 1.0)
	if num != 3.14 {
		t.Errorf("Expected '3.14', got %f", num)
	}
}

func TestToFloat64DefaultWithDefault(t *testing.T) {
	num := ToFloat6Default(".", 3.14)
	if num != 3.14 {
		t.Errorf("Expected '3.14', got %f", num)
	}
}
