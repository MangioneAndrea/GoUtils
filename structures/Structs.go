package structures

import (
	"encoding/json"
	"fmt"
)

func ConvertTo[T any](pb any) (res T, e error) {
	bytes, e := json.Marshal(pb)

	if e == nil {
		e = json.Unmarshal(bytes, &res)
	}
	return res, e
}

// StructToMap parse a struct into a map
func StructToMap(in interface{}) (res map[string]interface{}, e error) {
	bytes, e := json.Marshal(in)
	if e != nil {
		return nil, e
	}
	e = json.Unmarshal(bytes, &res)
	return res, e
}

// MapToStruct parse a map into a struct
func MapToStruct[T any](el map[string]interface{}) (T, error) {
	var res T
	j, e := json.Marshal(el)
	if e != nil {
		return res, e
	}
	e = json.Unmarshal(j, &res)
	fmt.Println(res)
	return res, e
}
