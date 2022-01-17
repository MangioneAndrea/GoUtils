package regex

import "testing"

func TestAllMatchRegex_ReturnFalseIfSomeElementDoNotMatch(t *testing.T) {
	const myRegex Regex = "^hello$"

	arr := []string{"hello", "hello", "hi"}

	res := AllMatchRegex(arr, myRegex)

	if res {
		t.Errorf("AllMatchRegex expected false for %s with regex %s", arr, myRegex)
	}
}
func TestAllMatchRegexChecksAllItems(t *testing.T) {
	const myRegex Regex = "^hello$"

	arr := []string{"hello", "hello", "hello"}

	res := AllMatchRegex(arr, myRegex)

	if !res {
		t.Errorf("AllMatchRegex expected true for %s with regex %s", arr, myRegex)
	}
}
