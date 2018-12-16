package gostrutils

import (
	"reflect"
	"testing"
)

func TestStrToInt64Valid(t *testing.T) {
	result := StrToInt64("-42", 2)
	if result != -42 {
		t.Errorf("Invalid result value, expected '-42', got %d", result)
	}
}

func TestStrToInt64WithDefaultValue(t *testing.T) {
	result := StrToInt64("a", -42)
	if result != -42 {
		t.Errorf("Invalid result value, expected '-42', got %d", result)
	}
}

func TestStrToUInt64Valid(t *testing.T) {
	result := StrToUInt64("42", 2)
	if result != 42 {
		t.Errorf("Invalid result value, expected '42', got %d", result)
	}
}

func TestStrToUInt64WithDefaultValue(t *testing.T) {
	result := StrToUInt64("a", 42)
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

func TestToFloat32Valid(t *testing.T) {
	num := ToFloat32("3.14")
	if num != 3.14 {
		t.Errorf("Expected '3.14', got '%f'", num)
	}
}

func TestToFloat32Zero(t *testing.T) {
	num := ToFloat32("Hello World")
	if num != 0.0 {
		t.Errorf("Expected '0.0', got '%f'", num)
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

func TestToFloat64Valid(t *testing.T) {
	num := ToFloat64("3.14")
	if num != 3.14 {
		t.Errorf("Expected '3.14', got '%f'", num)
	}
}

func TestToFloat64Zero(t *testing.T) {
	num := ToFloat64("Hello World")
	if num != 0.0 {
		t.Errorf("Expected '0.0', got '%f'", num)
	}
}

func TestIsUIntegerTrue(t *testing.T) {
	txt := "42"
	result := IsUInteger(txt)
	if !result {
		t.Errorf("Expected `%s` to be true", txt)
	}
}

func TestIsUIntegerFalse(t *testing.T) {
	txt := "42 Meaning"
	result := IsUInteger(txt)
	if result {
		t.Errorf("Expected `%s` to be false", txt)
	}
}

func TestIsIntegerTrue(t *testing.T) {
	txt := "-42"
	result := IsInteger(txt)
	if !result {
		t.Errorf("Expected `%s` to be true", txt)
	}
}

func TestIsIntegerFalse(t *testing.T) {
	txt := "-42 Meaning"
	result := IsInteger(txt)
	if result {
		t.Errorf("Expected `%s` to be false", txt)
	}
}

func TestIsUFloatTrue(t *testing.T) {
	txt := "3.14"
	result := IsUFloat(txt)
	if !result {
		t.Errorf("Expected `%s` to be true", txt)
	}
}

func TestIsUFloatFalse(t *testing.T) {
	txt := "-3.14"
	result := IsUFloat(txt)
	if result {
		t.Errorf("Exprected `%s` to be false", txt)
	}
}

func TestIsFloatTrue(t *testing.T) {
	txt := "-3.14"
	result := IsFloat(txt)
	if !result {
		t.Errorf("Exprected `%s` to be true", txt)
	}
}

func TestIsFloatFalse(t *testing.T) {
	txt := "42"
	result := IsFloat(txt)
	if result {
		t.Errorf("Expected `%s` to be true", txt)
	}
}

func TestIsUNumberIntTrue(t *testing.T) {
	txt := "42"
	result := IsUNumber(txt)
	if !result {
		t.Errorf("Expected `%s` to be true", txt)
	}
}

func TestIsUNumberFloatTrue(t *testing.T) {
	txt := "3.14"
	result := IsUNumber(txt)
	if !result {
		t.Errorf("Expected `%s` to be true", txt)
	}
}

func TestUNumberFalse(t *testing.T) {
	txt := "42 meaning"
	result := IsUNumber(txt)
	if result {
		t.Errorf("Expected `%s` to be false", txt)
	}
}

func TestIsNumberIntTrue(t *testing.T) {
	txt := "-42"
	result := IsNumber(txt)
	if !result {
		t.Errorf("Expected `%s` to be true", txt)
	}
}

func TestIsNumberFloatTrue(t *testing.T) {
	txt := "-3.14"
	result := IsNumber(txt)
	if !result {
		t.Errorf("Expected `%s` to be true", txt)
	}
}

func TestNumberFalse(t *testing.T) {
	txt := "42 meaning"
	result := IsNumber(txt)
	if result {
		t.Errorf("Expected `%s` to be false", txt)
	}
}
