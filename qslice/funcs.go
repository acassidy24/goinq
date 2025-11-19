package qslice

// add byReference version
func Where[T any](s []T, predicate func(T) bool) []T {
	ret := []T{}
	for _, v := range s {
		if predicate(v) {
			ret = append(ret, v)
		}
	}
	return ret
}
