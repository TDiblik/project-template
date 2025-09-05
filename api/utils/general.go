package utils

import (
	"math"
	"math/rand/v2"
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
