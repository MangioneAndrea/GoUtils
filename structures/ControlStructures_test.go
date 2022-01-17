package structures

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestArrayMap_StringToInt(t *testing.T) {
	type V string
	type R int

	want := []R{1, 2}
	got := ArrayMap([]V{"1", "2"}, func(el V) R {
		i, _ := strconv.ParseInt(string(el), 10, 64)
		return R(i)
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ArrayMap() = %v, want %v", got, want)
	}
}
func TestArrayMap_IntToString(t *testing.T) {
	type V int
	type R string

	want := []R{"1", "2"}
	got := ArrayMap([]V{1, 2}, func(el V) R {
		return R(fmt.Sprint(el))
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ArrayMap() = %v, want %v", got, want)
	}
}
func TestMapEntries(t *testing.T) {
	m := map[int]string{
		1:  "hi",
		7:  "hello",
		15: "hola",
	}
	expKeys := []int{1, 7, 15}
	expValues := []string{"hi", "hello", "hola"}
	sort.Ints(expKeys)
	sort.Strings(expValues)

	keys, values := MapEntries(m)
	sort.Ints(keys)
	sort.Strings(values)

	if !reflect.DeepEqual(keys, expKeys) {
		t.Errorf("MapEntries() got = %v, want %v", keys, expKeys)
	}
	if !reflect.DeepEqual(values, expValues) {
		t.Errorf("MapEntries() got1 = %v, want %v", values, expValues)
	}

}

func TestMapKeys(t *testing.T) {
	m := map[int]any{
		1:  "hi",
		7:  "hello",
		15: "hola",
	}
	expKeys := []int{1, 7, 15}
	expValues := []string{"hi", "hello", "hola"}
	sort.Ints(expKeys)
	sort.Strings(expValues)

	keys := MapKeys(m)
	sort.Ints(keys)

	if !reflect.DeepEqual(keys, expKeys) {
		t.Errorf("MapEntries() got = %v, want %v", keys, expKeys)
	}
}

func TestForEach(t *testing.T) {
	arr := []interface{}{1, "hello", map[string]int{}}

	arr2 := make([]interface{}, 3)
	ForEach(arr, func(el interface{}, i int) {
		arr2[i] = el
	})

	if !reflect.DeepEqual(arr[0], arr2[0]) {
		t.Errorf("copy[0] = %v, want %v", arr[0], arr[0])
	}
	if !reflect.DeepEqual(arr[1], arr2[1]) {
		t.Errorf("original[0] = %v, want %v", arr[0], arr[0])
	}
	if !reflect.DeepEqual(arr[2], arr2[2]) {
		t.Errorf("original[0] = %v, want %v", arr[0], arr[0])
	}
}

func TestMapValues(t *testing.T) {
	m := map[int]string{
		1:  "hi",
		7:  "hello",
		15: "hola",
	}
	expValues := []string{"hi", "hello", "hola"}
	sort.Strings(expValues)

	values := MapValues(m)
	sort.Strings(values)

	if !reflect.DeepEqual(values, expValues) {
		t.Errorf("MapEntries() got1 = %v, want %v", values, expValues)
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

func TestUniqueArr_int(t *testing.T) {
	input := []int{0, 0, 0, 1, 2, 2, 3}
	expected := []int{0, 1, 2, 3}

	output := UniqueArr(input)
	if !reflect.DeepEqual(expected, output) {
		t.Errorf("StructToMap() gotRes = %v, want %v", output, expected)
	}
}

func TestUniqueArr_str(t *testing.T) {
	input := []string{"", "hello", "hi", "hello", "hi"}
	expected := []string{"", "hello", "hi"}

	output := UniqueArr(input)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("StructToMap() gotRes = %v, want %v", output, expected)
	}
}
