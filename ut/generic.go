package ut

import (
	"fmt"
	"math/rand"
	"reflect"
	"slices"
	"sort"
	"strings"
	"time"

	"golang.org/x/exp/constraints"
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

func IsFunctionFullyDefined[T comparable](f map[T]T) (bool, *T) {
	// f: set -> set
	set := []T{}
	for key, v := range f {
		set = append(set, key)
		set = append(set, v)
	}
	set = Uniquify(set)
	for _, element := range set {
		if _, ok := f[element]; !ok {
			// fmt.Printf("Function not defined for element: %v\n", element)
			return false, &element
		}
	}
	return true, nil
}

func Min[T constraints.Ordered](a T, b ...T) T {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}
func Min_old[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
func NewMinAndMax[T constraints.Ordered](min, max, value T) (T, T) {
	if min > value {
		return value, max
	}
	if max < value {
		return min, value
	}
	return min, max
}

func GetSorted[T constraints.Ordered](data []T) []T {
	if len(data) == 0 {
		return nil
	}
	x := slices.Clone(data)
	slices.Sort(x)
	return x
}

func Equal2DSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func IsPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func AllAppearEvenTimes(nums []int) bool {
	counts := make(map[int]int)
	for _, n := range nums {
		counts[n]++
	}
	for _, count := range counts {
		if count%2 != 0 {
			return false
		}
	}
	return true
}

type Tuple [2]string

func PrintMapAsCode[K comparable, V any](m map[K]V) {
	var sb strings.Builder

	mapType := reflect.TypeOf(m)
	sb.WriteString(fmt.Sprintf("map[%s]%s{", mapType.Key(), mapType.Elem()))

	// Collect and sort keys (as strings) to ensure deterministic output
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// For known types like string or int, we can sort directly
	switch any(keys[0]).(type) {
	case string:
		sort.Slice(keys, func(i, j int) bool {
			return any(keys[i]).(string) < any(keys[j]).(string)
		})
	case int:
		sort.Slice(keys, func(i, j int) bool {
			return any(keys[i]).(int) < any(keys[j]).(int)
		})
	}

	for i, k := range keys {
		// Format keys as string literals if they're strings
		var keyStr string
		if ks, ok := any(k).(string); ok {
			keyStr = fmt.Sprintf("%q", ks)
		} else {
			keyStr = fmt.Sprintf("%v", k)
		}

		v := m[k]
		sb.WriteString(fmt.Sprintf("%s: %v", keyStr, v))
		if i < len(keys)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("}")
	fmt.Println(sb.String())
}

func ShuffleSlice[T any](s []T) {
	// r := rand.New(rand.NewSource(42))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}
