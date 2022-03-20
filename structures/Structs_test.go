package structures

import (
	"reflect"
	"testing"
)

func TestConvertTo(t *testing.T) {
	type A struct {
		ValueA string
		ValueB int
		ValueC bool
	}
	type B struct {
		ValueA string
		ValueB int
		ValueD bool
	}
	input := A{ValueA: "hello", ValueB: 3, ValueC: true}
	expected := B{ValueA: "hello", ValueB: 3, ValueD: false}

	output, err := ConvertTo[B](input)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("ConvertTo() gotRes = %v, want %v", output, expected)
	}
}

func TestMapToStruct(t *testing.T) {
	type Struct struct {
		A int
		B string
		C bool
		d float32 // not exported
	}
	input := map[string]any{
		"A": 1,
		"B": "hello",
		"C": true,
		"d": 45., // not exported,
	}

	res, err := MapToStruct[Struct](input)

	if err != nil {
		t.Errorf("Expected err to be nil, got %s", err)
	}
	if !reflect.DeepEqual(input["A"], res.A) {
		t.Errorf("StructToMap() gotRes = %v, want %v", res.A, input["A"])
	}
	if !reflect.DeepEqual(input["B"], res.B) {
		t.Errorf("StructToMap() gotRes = %v, want %v", res.B, input["B"])
	}
	if !reflect.DeepEqual(input["C"], res.C) {
		t.Errorf("StructToMap() gotRes = %v, want %v", res.C, input["C"])
	}
	if res.d != 0 {
		t.Errorf("Not exported field d was %v, want 0", res.d)
	}
}

func TestStructToMap(t *testing.T) {
	type Struct struct {
		A int
		B string
		C bool
		d float32 // not exported
	}
	input := Struct{
		1, "hello", true, 45.,
	}

	res, err := StructToMap(input)

	if err != nil {
		t.Errorf("Expected err to be nil, got %s", err)
	}
	if res["A"] == input.A {
		t.Errorf("StructToMap() gotRes = %v, want %v", input.A, res["A"])
	}
	if !reflect.DeepEqual(res["B"], input.B) {
		t.Errorf("StructToMap() gotRes = %v, want %v", input.B, res["B"])
	}
	if !reflect.DeepEqual(res["C"], input.C) {
		t.Errorf("StructToMap() gotRes = %v, want %v", input.C, res["C"])
	}
	if res["d"] != nil {
		t.Errorf("Not exported field d was %v, want nill", res["d"])
	}
}
