package utils

import (
	"math"
	"math/rand/v2"
	"net/url"
	"path"
)

// WARNING: This fucks up the order of the array tho
func RemoveIndexFromArrayFast[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// Bigger bias != bias towards higher numbers (https://gamedev.stackexchange.com/a/116875)
func GetBiasedRandom(min, max int64, bias float64) int64 {
	return int64(math.Floor(float64(min) + (float64(max)-float64(min))*math.Pow(float64(rand.Float64()), float64(bias))))
}

// JoinUrl joins a base URL with additional path segments safely
func JoinUrl(base string, paths ...string) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	all := []string{u.Path}
	all = append(all, paths...)
	u.Path = path.Join(all...)

	return u.String(), nil
}

// Calls JoinUrl, but panics on errors, use with causion
func JoinUrlOrPanic(base string, paths ...string) string {
	url, err := JoinUrl(base, paths...)
	if err != nil {
		panic(err)
	}
	return url
}

func DerefOrEmpty[T any](val *T) T {
	if val == nil {
		var empty T
		return empty
	}
	return *val
}

func IsNotNil[T any](val *T) bool {
	return val != nil
}
