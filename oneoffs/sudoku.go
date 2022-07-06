package oneoffs

import (
	"fmt"
	"log"
	"math"
	"strings"
)

type sudokuGrid struct {
	cells   [][]int
	size    int
	boxSize int
	iters   int
}

func initSudokuGrid(size int) sudokuGrid {
	var cells [][]int
	for i := 0; i < size; i++ {
		row := []int{}
		for j := 0; j < size; j++ {
			row = append(row, 0)
		}
		cells = append(cells, row)
	}

	return sudokuGrid{cells: cells, size: size, boxSize: int(math.Sqrt(float64(size)))}
}
func initSudokuGrid2(values [][]int) sudokuGrid {
	l := len(values)
	return sudokuGrid{cells: values, size: l, boxSize: int(math.Sqrt(float64(l))), iters: 0}
}

type set map[int]struct{}

func newSet(keys ...int) set {
	new := make(set)
	new.add(keys...)
	return new
}
func (x *set) add(keys ...int) {
	temp := *x
	for _, v := range keys {
		temp[v] = struct{}{}
	}
}

func (x *sudokuGrid) solve() {
	if x.isSolution() {
		return
	}
	if x.iters > 20000 {
		log.Fatalf("exceeded max iters")
	}
	if x.iters%1000 == 0 {
		x.printGrid()
	}
	x.iters++
	for i, row := range x.cells {
		for j, v := range row {
			if v == 0 {
				for try := 1; try <= x.size; try++ {
					if x.possible() {
						x.cells[i][j] = try
						x.solve()
						if x.isSolution() {
							x.printGrid()
							return
						}
						x.cells[i][j] = 0
					}
				}
				return
			}
		}
	}
}
func (x *sudokuGrid) possible() bool {
	colValues := newSet()
	for _, row := range x.cells {
		rowValues := newSet()
		for j, v := range row {
			if v == 0 {
				continue
			}
			colKey := 100*j + v
			if _, ok := colValues[colKey]; ok {
				return false
			}
			colValues.add(colKey)
			if _, ok := rowValues[v]; ok {
				return false
			}
			rowValues.add(v)
		}
	}
	for i := 0; i < x.boxSize; i++ {
		for j := 0; j < x.boxSize; j++ {
			boxValues := newSet()
			for ii := 0; ii < x.boxSize; ii++ {
				for jj := 0; jj < x.boxSize; jj++ {
					key := x.cells[x.boxSize*i+ii][x.boxSize*j+jj]
					if key == 0 {
						continue
					}
					if _, ok := boxValues[key]; ok {
						return false
					}
					boxValues.add(key)
				}
			}
		}
	}
	return true
}
func (x *sudokuGrid) isSolution() bool {
	colValues := newSet()
	for _, row := range x.cells {
		rowValues := newSet()
		for j, v := range row {
			if v == 0 {
				return false
			}
			colKey := 100*j + v
			if _, ok := colValues[colKey]; ok {
				return false
			}
			colValues.add(colKey)
			if _, ok := rowValues[v]; ok {
				return false
			}
			rowValues.add(v)
		}
	}
	for i := 0; i < x.boxSize; i++ {
		for j := 0; j < x.boxSize; j++ {
			boxValues := newSet()
			for ii := 0; ii < x.boxSize; ii++ {
				for jj := 0; jj < x.boxSize; jj++ {
					key := x.cells[x.boxSize*i+ii][x.boxSize*j+jj]
					if _, ok := boxValues[key]; ok {
						return false
					}
					boxValues.add(key)
				}
			}
		}
	}
	return true
}
func (x *sudokuGrid) printGrid() {
	s := fmt.Sprintf("iter=%d\n ", x.iters)
	for i, row := range x.cells {
		for j, v := range row {
			bar := ""
			if (j+1)%x.boxSize == 0 && j != x.size-1 {
				bar = "| "
			}
			s += fmt.Sprintf("%d %s", v, bar)
		}
		if (i+1)%x.boxSize == 0 && i != x.size-1 {
			s += "\n"
			for k := 0; k < x.boxSize; k++ {
				s += strings.Repeat("-", x.boxSize*2+1) + "+"
			}
		}
		s = s[:len(s)-1] + "\n "
	}
	fmt.Printf("%v\n", s)
}
