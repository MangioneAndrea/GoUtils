package structures

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
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// MapValues Return an array representing the of values of the given map
func MapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))

	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// ForEach Run a function for each element of an array. The parameters are ECMA alike without array (value, index)
// where index and array are optional
func ForEach[V any](arr []V, f func(el V, i int)) {
	for i, v := range arr {
		f(v, i)
	}
}

// Filter an array with the result of a given function. True keeps the element, false discards it
func Filter[V any](arr []V, f func(el V, i int) bool) (result []V) {
	for i, v := range arr {
		if f(v, i) {
			result = append(result, v)
		}
	}
	return result
}

// ArrayMap function equivalent to the javascript .map
func ArrayMap[V any, R any](arr []V, f func(el V) R) []R {
	result := make([]R, len(arr))
	for i, elem := range arr {
		result[i] = f(elem)
	}
	return result
}
