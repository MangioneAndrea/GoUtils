package console

import (
	"bytes"
	"testing"
)

func TestLog(t *testing.T) {
	Log("hello")
	out = new(bytes.Buffer)

	myStr := "hello"

	Log(myStr)

	result := out.(*bytes.Buffer).String()

	if result != myStr {
		t.Errorf("Log %s, expected %s", result, myStr)
	}
}
