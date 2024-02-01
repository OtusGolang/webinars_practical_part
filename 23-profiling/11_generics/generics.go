package generics

import (
	"golang.org/x/exp/constraints"
)

func ConvertMapWithoutGenerics(a map[string]int32) map[string]int64 {
	newMap := make(map[string]int64)

	for key, value := range a {
		newMap[key] = int64(value)
	}

	return newMap
}

func ConvertMapWithGenerics[K comparable, FROM, TO constraints.Integer](in map[K]FROM) map[K]TO {
	nMap := make(map[K]TO, len(in))
	for k, v := range in {
		nMap[k] = TO(v)
	}
	return nMap
}

func GMax[T interface{ string | int }](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func IMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SMax(a, b string) string {
	if a > b {
		return a
	}
	return b
}
