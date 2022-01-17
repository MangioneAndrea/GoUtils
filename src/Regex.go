package Util

import "regexp"

type Regex string

// LanguageISO639 Regex matching valid languages (Only languages currently translated by JS Intl.format are listed here 30/12/2021)
const LanguageISO639 Regex = "^(af|ak|am|an|ar|as|az|be|bg|bn|bo|br|bs|ca|co|cs|cy|da|de|ee|el|en|eo|es|et|eu|fa|fi|fo|fr|fy|ga|gd|gl|gn|gu|ha|he|hi|hr|ht|hu|hy|ia|id|ig|is|it|iu|ja|jv|ka|kg|kk|km|kn|ko|ku|ky|la|lb|lg|ln|lo|lt|lu|lv|mg|mh|mi|mk|ml|mn|mr|ms|mt|my|nb|ne|nl|nn|no|ny|oc|om|or|pa|pl|ps|pt|qu|rm|rn|ro|ru|rw|sd|si|sk|sl|sm|sn|so|sq|sr|st|su|sv|sw|ta|te|tg|th|ti|tk|tl|tn|to|tr|tt|tw|ty|ug|uk|ur|uz|vi|wa|wo|xh|yi|yo|zh|zu)$"

func (r Regex) Compiled() *regexp.Regexp {
	res, _ := regexp.Compile(string(r))
	return res
}

// AllMatchRegex return whether all string elements of an array match a given regex
func AllMatchRegex(input []string, reg Regex) bool {
	r := reg.Compiled()
	for _, val := range input {
		if !r.MatchString(val) {
			return false
		}
	}
	return true
}
