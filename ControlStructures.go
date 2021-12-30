package Util

import (
	"encoding/json"
)

// UniqueArr Return an array with unique values
func UniqueArr[T comparable](input []T) (res []T) {
	m := make(map[T]bool)
	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			res = append(res, val)
		}
	}
	return res
}

// MapEntries Return two arrays representing the set of keys and values of the given map
func MapEntries[K comparable, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(m))
	values := make([]V, 0, len(m))

	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

// MapKeys Return an array representing the set of keys of the given map
func MapKeys[K comparable](m map[K]any) []K {
	keys := make([]K, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// MapValues Return an array representing the of values of the given map
func MapValues[V any](m map[interface{}]V) []V {
	values := make([]V, 0, len(m))

	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// ArrayMap function equivalent to the javascript .map
func ArrayMap[V any, R any](arr []V, f func(el V) R) []R {
	result := make([]R, len(arr))
	for i, elem := range arr {
		result[i] = f(elem)
	}
	return result
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
func MapToStruct[T any](el map[string]interface{}) (res T, e error) {
	j, e := json.Marshal(el)
	if e != nil {
		return T{}, e
	}
	e = json.Unmarshal(j, &res)
	return res, e
}
