package kolman

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test_algorithmX(t *testing.T) {
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		{"algorithmX", algorithmXTest},
		// {"isSolution", isSolutionTest},
		// {"removeColumn", removeColumnTest},
		// {"removeRow", removeRowTest},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}

func algorithmXTest(t *testing.T) {
	tests := []struct {
		name   string
		sets   map[int][]int
		expect []int
	}{
		// {"2x2 not disjoint -> error", map[int][]int{
		// 	0: {1, 0},
		// 	1: {1, 1},
		// }, []int{}},
		// {"2x2", map[int][]int{
		// 	0: {1, 0},
		// 	1: {0, 1},
		// }, []int{0, 1}},
		// {"2.6 ex1, 10x6", map[int][]int{
		// 	0: {1, 1, 0, 0, 1, 0, 0, 0, 0, 0},
		// 	1: {0, 0, 1, 0, 1, 0, 0, 1, 1, 0},
		// 	2: {1, 0, 0, 0, 0, 0, 0, 1, 0, 1},
		// 	3: {0, 0, 0, 1, 0, 0, 1, 0, 0, 1},
		// 	4: {0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		// 	5: {1, 1, 0, 0, 0, 1, 0, 0, 0, 0},
		// }, []int{1, 3, 5}},
		// {"3x3 simple", map[int][]int{
		// 	0: {1, 0, 1},
		// 	1: {1, 0, 0},
		// 	2: {0, 1, 0},
		// }, []int{0, 2}},
		// {"3x3 requires backtrack", map[int][]int{
		// 	0: {0, 0, 1},
		// 	1: {1, 0, 0},
		// 	2: {0, 1, 1},
		// }, []int{1, 2}},

		// {"3x4", map[int][]int{
		// 	0: {1, 0, 1},
		// 	1: {0, 1, 0},
		// 	2: {1, 0, 0},
		// 	3: {0, 1, 1},
		// }, []int{0, 1}},
		// {"3x4", map[int][]int{
		// 	0: {1, 0, 0},
		// 	1: {0, 1, 0},
		// 	2: {1, 0, 0},
		// 	3: {0, 0, 1},
		// }, []int{0, 1, 3}},
		// {"4x4", map[int][]int{
		// 	0: {1, 0, 0, 0},
		// 	1: {0, 1, 0, 0},
		// 	2: {0, 0, 1, 0},
		// 	3: {0, 0, 1, 1},
		// }, []int{0, 1, 3}},
		// {"7x6", map[int][]int{
		// 	0: {1, 1, 0, 0, 1, 0, 0},
		// 	1: {1, 1, 0, 0, 0, 0, 0},
		// 	2: {1, 0, 0, 0, 1, 1, 0},
		// 	3: {0, 0, 0, 1, 0, 1, 1},
		// 	4: {0, 0, 1, 1, 1, 0, 1},
		// 	5: {0, 0, 1, 0, 1, 0, 0},
		// }, []int{1, 3, 5}},
	}
	for i, tt := range tests {
		x := exactCover{sets: tt.sets}
		err := x.algorithmX()
		if err != nil {
			t.Error(err)
		}
		sort.Slice(x.coverIndexes, func(i2, j int) bool {
			return x.coverIndexes[i2] < x.coverIndexes[j]
		})
		if reflect.DeepEqual(x.coverIndexes, tt.expect) {
			fmt.Printf("SUCCESS %d: %s", i, tt.name)
		} else {
			t.Errorf("FAIL %d: expected: %v, got: %v", i, tt.expect, x.coverIndexes)
		}
	}
}
func isSolutionTest(t *testing.T) {
	tests := []struct {
		sets         map[int][]int
		coverIndexes []int
		expect       bool
	}{
		{map[int][]int{
			0: {1, 0, 1},
			1: {0, 1, 0},
			2: {1, 0, 0},
		}, []int{0, 1}, true},
		{map[int][]int{
			0: {1, 0, 1},
			1: {0, 1, 0},
			2: {1, 0, 0},
		}, []int{0, 1, 2}, false},
		{map[int][]int{
			0: {1, 0, 1},
			1: {0, 1, 0},
			2: {1, 0, 0},
		}, []int{1, 2}, false},
	}
	for i, tt := range tests {
		x := exactCover{sets: tt.sets, coverIndexes: tt.coverIndexes}

		if x.isSolution() != tt.expect {
			t.Errorf("FAIL %d", i)
		}
	}
}
func removeColumnTest(t *testing.T) {
	tests := []struct {
		rmCol  int
		sets   map[int][]int
		expect map[int][]int
	}{
		{1, map[int][]int{
			0: {1, 0, 1},
			1: {0, 1, 0},
			2: {1, 0, 0},
		}, map[int][]int{
			0: {1, 1},
			1: {0, 0},
			2: {1, 0},
		},
		},
		{2, map[int][]int{
			0: {1, 0, 1},
			1: {0, 1, 0},
			2: {1, 0, 0},
		}, map[int][]int{
			0: {1, 0},
			1: {0, 1},
			2: {1, 0},
		}},
		{0, map[int][]int{
			0: {1, 0, 1},
			1: {0, 1, 0},
			2: {1, 0, 0},
		}, map[int][]int{
			0: {0, 1},
			1: {1, 0},
			2: {0, 0},
		}},
		{2, map[int][]int{
			0: {1, 1, 0, 0, 1, 0, 0},
			1: {1, 1, 0, 0, 0, 0, 0},
			2: {1, 0, 0, 0, 1, 1, 0},
			3: {0, 0, 0, 1, 0, 1, 1},
			4: {0, 0, 1, 1, 1, 0, 1},
			5: {0, 0, 1, 0, 1, 0, 0},
		}, map[int][]int{
			0: {1, 1, 0, 1, 0, 0},
			1: {1, 1, 0, 0, 0, 0},
			2: {1, 0, 0, 1, 1, 0},
			3: {0, 0, 1, 0, 1, 1},
			4: {0, 0, 1, 1, 0, 1},
			5: {0, 0, 0, 1, 0, 0},
		}},
	}
	for i, tt := range tests {
		x := exactCover{sets: tt.sets}
		x.removeColumn(tt.rmCol)
		if !reflect.DeepEqual(x.sets, tt.expect) {
			t.Errorf("FAIL %d", i)
		}
	}
}
func removeRowTest(t *testing.T) {
	tests := []struct {
		rmCol  int
		m      [][]int
		expect [][]int
	}{
		{0, [][]int{
			{1, 0, 1},
			{0, 1, 0},
			{1, 0, 0},
		}, [][]int{
			{0, 1, 0},
			{1, 0, 0},
		}},
		{1, [][]int{
			{1, 0, 1},
			{0, 1, 0},
			{1, 0, 0},
		}, [][]int{
			{1, 0, 1},
			{1, 0, 0},
		}},
		{2, [][]int{
			{1, 0, 1},
			{0, 1, 0},
			{1, 0, 0},
		}, [][]int{
			{1, 0, 1},
			{0, 1, 0},
		}},
	}
	for i, tt := range tests {
		got := removeRow(tt.rmCol, tt.m)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("FAIL %d: expected: %v, got: %v", i, tt.expect, got)
		}
	}
}
