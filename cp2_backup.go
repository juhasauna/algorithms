package temp

import (
	"fmt"
	"log"
	"reflect"
)

func algorithmX(a [][]int) []int {
	if len(a) == 1 {
		return []int{}
	}
	m := a

	lSet := []int{}
	colOnesCount := []int{}

	for i := 0; i < len(m); i++ {
		isZeroCol := true
		for j := 0; j < len(m[0]); j++ {
			if len(colOnesCount) == j {
				colOnesCount = append(colOnesCount, 0)
			}
			if m[i][j] == 1 {
				isZeroCol = false
				colOnesCount[j]++
			}
		}
		if isZeroCol {
			fmt.Println("no solution since there is a zero column")
			return lSet
		}
	}
	// a)
	chosenCol := 0
	minOnesCount := colOnesCount[0]
	for i, v := range colOnesCount {
		if v < minOnesCount {
			minOnesCount = v
			chosenCol = i
		}
	}
	fmt.Println("chosenCol", chosenCol)
	// b)
	chosenRow := 0
	for i, row := range m {
		if row[chosenCol] == 1 {
			chosenRow = i
			lSet = append(lSet, chosenRow)
			break
		}
	}
	fmt.Println("chosenRow", chosenRow)
	// c) Eliminate superfluous rows

	chosenRowValues := []int{}
	for _, v := range m[chosenRow] {
		chosenRowValues = append(chosenRowValues, v)
	}

	rowsToRemove := []int{}
	for i := 0; i < len(m); i++ {
		if i == chosenRow {
			continue
		}
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] == chosenRowValues[j] && m[i][j] == 1 {
				rowsToRemove = append(rowsToRemove, i)
				break
			}
		}
	}
	fmt.Printf("rows to remove %v\n", rowsToRemove)
	for i := len(rowsToRemove) - 1; i > 0; i-- {
		m = removeRow(rowsToRemove[i], m)
	}

	// d) Eliminate superfluous cols
	colsToRemove := []int{}
	for i, v := range chosenRowValues {
		if v == 1 {
			colsToRemove = append(colsToRemove, i)
		}
	}

	for i := len(colsToRemove) - 1; i > 0; i-- {
		m = removeColumn(colsToRemove[i], m)
	}

	// Remove chosenRow
	m = append(m[:chosenRow], m[chosenRow+1:]...)
	if reflect.DeepEqual(a, m) {
		fmt.Println("final matrix", m)
		return lSet
	}
	fmt.Printf("next iteration starting, lSet: %v, m: %v\n", lSet, m)
	return append(lSet, algorithmX(m)...)
	// TODO: save original matrix indexes.
}

func removeColumn(c int, m [][]int) [][]int {
	if c > len(m[0])-1 {
		log.Fatalf("invalid remove column input, column too big: %d", c)
	}
	if c < 0 {
		log.Fatalf("invalid remove column input, column too small: %d", c)
	}
	for i, v := range m {
		v = append(v[:c], v[c+1:]...)
		m[i] = v
	}
	return m
}

func removeRow(r int, m [][]int) [][]int {
	if r > len(m)-1 {
		log.Fatalf("invalid remove row input, row too big: %d", r)
	}
	if r < 0 {
		log.Fatalf("invalid remove row input, row too small: %d", r)
	}

	return append(m[:r], m[r+1:]...)
}
