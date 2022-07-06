package oneoffs

import (
	"fmt"
	"testing"
)

func Test_sudoku(t *testing.T) {
	initTestData()
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		// {"initSudokuGrid", initSudokuGridTest},
		{"isSolution", isSolutionTest},
		// {"solve", solveTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}

func initSudokuGridTest(t *testing.T) {
	tests := []struct {
		size int
	}{
		{4},
		{9},
	}
	for _, tt := range tests {
		got := initSudokuGrid(tt.size)
		got.printGrid()
	}
}
func isSolutionTest(t *testing.T) {
	tests := []struct {
		values [][]int
		want   bool
	}{
		{testData["9x9_solution_a"], true},
		{testData["9x9_a_2"], false},
		// {testData["9x9_a"], false},
		// {testData["4x4_a"], false},
		// {testData["4x4_solved_a"], true},
		// {testData["4x4_invalid_box"], false},
		// {testData["4x4_invalid_row"], false},
		// {testData["4x4_invalid_col"], false},
	}

	for i, tt := range tests {
		x := initSudokuGrid2(tt.values)
		got := x.isSolution()
		if got != tt.want {
			t.Errorf("FAIL %d: want: %t, got: %t", i, tt.want, got)
			x.printGrid()
		}

	}
}
func solveTest(t *testing.T) {
	tests := []struct {
		values [][]int
	}{

		// {testData["4x4_solved_a"]},
		// {testData["4x4_a"]},
		{testData["9x9_a"]},
	}

	for _, tt := range tests {
		x := initSudokuGrid2(tt.values)
		x.solve()
		x.printGrid()
		fmt.Println(x.isSolution())

	}
}

var testData map[string][][]int

func initTestData() {
	x := make(map[string][][]int)
	x["4x4_solved_a"] = [][]int{
		{1, 4, 2, 3},
		{2, 3, 1, 4},
		{3, 2, 4, 1},
		{4, 1, 3, 2}}
	x["4x4_a"] = [][]int{
		{1, 0, 0, 3},
		{2, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 0}}
	x["4x4_invalid_box"] = [][]int{
		{1, 2, 4, 3},
		{2, 3, 1, 4},
		{3, 2, 4, 1},
		{4, 1, 3, 2}}
	x["4x4_invalid_row"] = [][]int{
		{2, 4, 2, 3},
		{1, 3, 1, 4},
		{3, 2, 4, 1},
		{4, 1, 3, 2}}
	x["4x4_invalid_col"] = [][]int{
		{1, 4, 2, 3},
		{2, 3, 4, 1},
		{3, 2, 4, 1},
		{4, 1, 3, 2}}
	x["9x9_a"] = [][]int{
		{0, 0, 0, 0, 0, 6, 0, 8, 0},
		{0, 4, 3, 9, 0, 0, 1, 0, 7},
		{7, 2, 0, 1, 5, 0, 9, 0, 0},
		{1, 0, 0, 2, 0, 0, 6, 0, 0},
		{3, 0, 0, 0, 0, 4, 0, 0, 0},
		{0, 0, 0, 0, 6, 8, 0, 0, 5},
		{0, 0, 5, 0, 0, 0, 0, 7, 0},
		{0, 0, 0, 0, 2, 0, 0, 6, 1},
		{0, 0, 7, 6, 0, 0, 2, 0, 0}}
	x["9x9_solution_a"] = [][]int{
		{5, 9, 1, 7, 4, 6, 3, 8, 2},
		{6, 4, 3, 9, 8, 2, 1, 5, 7},
		{7, 2, 8, 1, 5, 3, 9, 4, 6},
		{1, 5, 4, 2, 7, 9, 6, 3, 8},
		{3, 8, 6, 5, 1, 4, 7, 2, 9},
		{9, 7, 2, 3, 6, 8, 4, 1, 5},
		{2, 6, 5, 4, 9, 1, 8, 7, 3},
		{4, 3, 9, 8, 2, 7, 5, 6, 1},
		{8, 1, 7, 6, 3, 5, 2, 9, 4}}
	x["9x9_a_2"] = [][]int{
		{5, 9, 1, 7, 4, 6, 3, 8, 2},
		{6, 4, 3, 9, 8, 2, 1, 5, 7},
		{7, 2, 8, 1, 5, 3, 9, 4, 6},
		{1, 5, 4, 2, 7, 9, 6, 3, 8},
		{3, 8, 6, 5, 1, 4, 7, 2, 9},
		{9, 7, 2, 3, 6, 8, 4, 1, 5},
		{2, 6, 5, 4, 9, 1, 8, 7, 3},
		{4, 3, 9, 8, 2, 7, 5, 6, 1},
		{8, 1, 7, 6, 3, 5, 2, 9, 0}}
	testData = x
}
