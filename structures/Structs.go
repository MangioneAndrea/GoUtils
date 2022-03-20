package structures

import (
	"encoding/json"
)

func ConvertTo[T any](pb any) (res T, e error) {
	bytes, e := json.Marshal(pb)

	if e == nil {
		e = json.Unmarshal(bytes, &res)
	}
	return res, e
}
