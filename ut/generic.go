package ut

import (
	"fmt"
	"reflect"
	"time"
)

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func AssertEqual(expected, actual any) {
	if !reflect.DeepEqual(expected, actual) {
		panic(fmt.Sprintf("Assertion failed: expected %v (%T), got %v (%T)", expected, expected, actual, actual))
	}
}

func Sum[T int | int32 | int64 | float64](s []T) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}

func HasDuplicates[T comparable](items []T) bool {
	seen := make(map[T]struct{}, len(items))
	for _, item := range items {
		if _, found := seen[item]; found {
			return true
		}
		seen[item] = struct{}{}
	}
	return false
}

func RemoveSliceValue[T comparable](slice []T, target T) []T {
	var result []T
	for _, v := range slice {
		if v != target {
			result = append(result, v)
		}
	}
	return result
}

func IsUnique[T comparable](slice []T) bool {
	seen := make(map[T]struct{})
	for _, v := range slice {
		if _, exists := seen[v]; exists {
			return false
		}
		seen[v] = struct{}{}
	}
	return true
}

func ToSet[T comparable](slice []T) map[T]struct{} {
	set := make(map[T]struct{}, len(slice))
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}

func Uniquify[T comparable](slice []T) []T {
	result := []T{}
	for key, _ := range ToSet(slice) {
		result = append(result, key)
	}
	return result
}
