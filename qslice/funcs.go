package qslice

func Where[T any](slice []T, predicate func(T) bool) []T {
	ret := []T{}
	for _, v := range slice {
		if predicate(v) {
			ret = append(ret, v)
		}
	}
	return ret
}
