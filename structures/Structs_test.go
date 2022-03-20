package structures

import (
	"reflect"
	"testing"
)

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

func TestConvertTo(t *testing.T) {
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
