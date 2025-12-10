package utils

import (
	"context"
	"database/sql"
	"log"
	"math"
	"math/rand/v2"
	"net/url"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"

	"github.com/TDiblik/project-template/api/models"
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

func WithSignalCancel(usecase string) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Received shutdown signal, stopping " + usecase + "...")
		cancel()
	}()
	return ctx
}

func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	sb := strings.Builder{}
	for range n {
		sb.WriteByte(letters[rand.IntN(len(letters))])
	}
	return sb.String()
}

func SQLNullStringFromString(value string) models.SQLNullString {
	return SQLNullStringFromStringRef(&value)
}

func SQLNullStringFromStringRef(value *string) models.SQLNullString {
	return models.SQLNullString{
		NullString: sql.NullString{
			String: DerefOrEmpty(value),
			Valid:  IsNotNil(value),
		},
	}
}
