package qmap

func Where[K comparable, V any](m map[K]V, predicate func(K, V) bool) map[K]V {
	ret := map[K]V{}
	for k, v := range m {
		if predicate(k, v) {
			ret[k] = v
		}
	}
	return ret
}
